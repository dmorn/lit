package scopus

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/jecoz/lit"
	"golang.org/x/net/publicsuffix"
)

const (
	searchEndpoint   = "https://api.elsevier.com/content/search/scopus"
	abstractEndpoint = "https://api.elsevier.com/content/abstract/eid/%s"
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
	Eid   string           `json:"eid"`
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
			Eid:   v.Eid,
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

func (c Client) GetLink(ctx context.Context, link string) (io.ReadCloser, error) {
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-ELS-APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return resp.Body, nil
}

func ParseAbstract(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(data)

	rgx := regexp.MustCompile("<section +.*id=\"abstractSection\".*>")
	index := rgx.FindReaderIndex(buf)
	if index == nil {
		return "", fmt.Errorf("abstractSection not found within input reader")
	}

	buf.Seek(int64(index[0]), io.SeekStart)

	d := xml.NewDecoder(buf)
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity

	max := 50
	for i := 0; i < max; i++ {
		t, err := d.Token()
		if err != nil {
			return "", fmt.Errorf("find abstract char data: %w", err)
		}
		if se, ok := t.(xml.StartElement); ok && se.Name.Local == "p" {
			next, err := d.Token()
			if err != nil {
				panic(err)
			}
			cd, ok := next.(xml.CharData)
			if !ok {
				return "", fmt.Errorf("expected to find char data after paragraph, found %T", next)
			}
			return string(cd), nil
		}
	}
	return "", fmt.Errorf("could not find abstract char data within abstract section")
}

func (c Client) GetAbstract(ctx context.Context, p lit.Publication) (lit.Abstract, error) {
	body, err := c.GetLink(ctx, p.Links["scopus"])
	if err != nil {
		return lit.Abstract{}, err
	}
	defer body.Close()

	text, err := ParseAbstract(body)
	if err != nil {
		return lit.Abstract{}, err
	}
	return lit.Abstract{
		Text: text,
	}, nil
}

func NewClient(apiKey string) Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    15 * time.Second,
		DisableCompression: false,
	}

	// Used for folling scopus links. Remember the default CheckRedirect
	// function allows at most 10 redirects to be followed.
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: tr,
		Jar:       jar,
	}

	return Client{
		apiKey:     apiKey,
		httpClient: client,
	}
}
