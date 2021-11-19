package scopus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jecoz/lit"
)

const (
	searchEndpoint = "https://api.elsevier.com/content/search/scopus"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func newSearchRequest(ctx context.Context, apiKey string, src lit.Request) *http.Request {
	u, err := url.Parse(searchEndpoint)
	if err != nil {
		panic(err)
	}
	q := u.Query()

	q.Set("query", src.Query)
	q.Set("count", fmt.Sprintf("%d", src.PerPage))
	q.Set("start", fmt.Sprintf("%d", src.PerPage*src.Page))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-ELS-APIKey", apiKey)
	req.Header.Set("Accept", "application/json")

	return req
}

func extractError(r *http.Response) error {
	if msg := r.Header.Get("X-Els-Status"); msg != "" {
		return fmt.Errorf("%s", msg)
	}

	return fmt.Errorf("%s", r.Status)
}

type openSearchLink struct {
	Tag string `json:"@ref"`
	Ref string `json:"@href"`
}

type openSearchEntry struct {
	Title string           `json:"dc:title"`
	Links []openSearchLink `json:"link"`
}

type openSearchResult struct {
	Total   string            `json:"opensearch:totalResults"`
	Entries []openSearchEntry `json:"entry"`
}

func mapPublications(entries []openSearchEntry) []lit.Publication {
	pubs := make([]lit.Publication, len(entries))
	for i, v := range entries {
		links := make(map[string]string)
		for _, w := range v.Links {
			links[w.Tag] = w.Ref
		}
		pubs[i] = lit.Publication{
			Title: v.Title,
			Links: links,
		}
	}
	return pubs
}

func (c Client) getSearchResults(ctx context.Context, req lit.Request) (*openSearchResult, error) {
	search := newSearchRequest(ctx, c.apiKey, req)
	resp, err := c.httpClient.Do(search)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, extractError(resp)
	}
	defer resp.Body.Close()

	var p struct {
		Result *openSearchResult `json:"search-results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	return p.Result, nil
}

func (c Client) GetLiterature(ctx context.Context, req lit.Request) (lit.Response, error) {
	result, err := c.getSearchResults(ctx, req)
	if err != nil {
		return lit.Response{}, err
	}

	return lit.Response{
		Req:        req,
		Literature: mapPublications(result.Entries),
	}, nil
}

func (c Client) GetMaxLiterature(ctx context.Context, req lit.Request) (int, error) {
	result, err := c.getSearchResults(ctx, req)
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(result.Total)
	if err != nil {
		return 0, fmt.Errorf("unexpected result field: %w", err)
	}
	return n, nil
}

func (c Client) ConcurrencyLimit() int {
	// https://dev.elsevier.com/api_key_settings.html
	return 6
}

func NewClient(apiKey string) Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    15 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{
		Transport: tr,
	}

	return Client{
		apiKey:     apiKey,
		httpClient: client,
	}
}
