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
	"strings"
	"time"

	"github.com/jecoz/lit"
	"github.com/jecoz/lit/bibtex"
	"golang.org/x/net/publicsuffix"
)

const (
	searchEndpoint   = "https://api.elsevier.com/content/search/scopus"
	abstractEndpoint = "https://api.elsevier.com/content/abstract/eid/%s"
)

const (
	KeyLinkAbstract    = "link_abstract"
	KeyEid             = "eid"
	KeyIssn            = "issn"
	KeyDOI             = "doi"
	KeyPageRange       = "page_range"
	KeyVolume          = "volume"
	KeyPublicationName = "publication_name"
	KeyArticleNumber   = "article_number"
	KeyAggregationType = "aggregation_type"
	KeySubtype         = "subtype"
	KeyCitedByCount    = "cited_by_count"
	KeyAffiliation     = "affiliation"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func (c Client) DefaultPerPage() int {
	return 25
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

type openAffiliation struct {
	Name    string `json:"affilname"`
	City    string `json:"affiliation-city"`
	Country string `json:"affiliation-country"`
}

type openSearchEntry struct {
	Title           string `json:"dc:title"`
	Eid             string `json:"eid"`
	CoverDateRaw    string `json:"prism:coverDate"`
	Creator         string `json:"dc:creator"`
	Issn            string `json:"prism:issn"`
	DOI             string `json:"prism:doi"`
	PageRange       string `json:"prism:pageRange"`
	Volume          string `json:"prism:volume"`
	PublicationName string `json:"prism:publicationName"`
	ArticleNumber   string `json:"article-number"`
	AggregationType string `json:"prism:aggregationType"`
	Subtype         string `json:"subtype"`
	CitedByCount    string `json:"citedby-count"`

	Links        []openSearchLink  `json:"link"`
	Affiliations []openAffiliation `json:"affiliation"`
}

func (e openSearchEntry) CoverDate() (time.Time, error) {
	t, err := time.Parse("2006-01-02", e.CoverDateRaw)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse cover date: %w", err)
	}
	return t, nil
}

func (e openSearchEntry) Affiliation() string {
	affiliations := []string{}
	for _, v := range e.Affiliations {
		fields := []string{
			v.Name,
			v.City,
			v.Country,
		}
		stripped := make([]string, 0, len(fields))
		for _, f := range fields {
			if len(f) > 0 {
				stripped = append(stripped, f)
			}
		}
		if len(stripped) == 0 {
			continue
		}
		affiliations = append(affiliations, strings.Join(stripped, ", "))
	}
	if len(affiliations) == 0 {
		return ""
	}
	return strings.Join(affiliations, "; ")
}

func (e openSearchEntry) Values() map[string]string {
	links := make(map[string]string)
	for _, v := range e.Links {
		links[v.Tag] = v.Ref
	}

	return map[string]string{
		KeyLinkAbstract:    links["scopus"],
		KeyEid:             e.Eid,
		KeyIssn:            e.Issn,
		KeyDOI:             e.DOI,
		KeyPageRange:       e.PageRange,
		KeyVolume:          e.Volume,
		KeyPublicationName: e.PublicationName,
		KeyArticleNumber:   e.ArticleNumber,
		KeyAggregationType: e.AggregationType,
		KeySubtype:         e.Subtype,
		KeyCitedByCount:    e.CitedByCount,
		KeyAffiliation:     e.Affiliation(),
	}
}

type openSearchResult struct {
	Total   string            `json:"opensearch:totalResults"`
	Entries []openSearchEntry `json:"entry"`
}

func mapPublications(entries []openSearchEntry) ([]lit.Publication, error) {
	pubs := make([]lit.Publication, len(entries))
	for i, v := range entries {
		coverDate, err := v.CoverDate()
		if err != nil {
			return pubs, fmt.Errorf("search result %d, %s: %w", i, v.Eid, err)
		}

		pubs[i] = lit.Publication{
			Title:     v.Title,
			CoverDate: coverDate,
			Creator:   v.Creator,
			Values:    v.Values(),
		}
	}
	return pubs, nil
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
	pubs, err := mapPublications(result.Entries)
	if err != nil {
		return lit.Response{}, err
	}
	return lit.Response{
		Req:        req,
		Literature: pubs,
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

func (c Client) GetRateLimit() time.Duration {
	// https://dev.elsevier.com/api_key_settings.html
	return time.Millisecond * 1000 / time.Duration(6)
}

func (c Client) ConcurrencyLimit() int {
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
	body, err := c.GetLink(ctx, p.Values[KeyLinkAbstract])
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

func (c Client) GetName() string {
	return "Scopus by ELSEVIER"
}

func makeEntry(p lit.Publication) bibtex.Entry {
	return bibtex.Entry{
		Title:  p.Title,
		Author: p.Creator,
		Year:   p.CoverDate.Year(),
	}
}

func makeArticle(p lit.Publication) bibtex.Reference {
	var pageRange *string
	var volume *string

	if r, ok := p.Values[KeyPageRange]; ok {
		pageRange = &r
	}
	if v, ok := p.Values[KeyVolume]; ok {
		volume = &v
	}

	return bibtex.Article{
		Entry:     makeEntry(p),
		Journal:   p.Values[KeyPublicationName],
		Volume:    volume,
		PageRange: pageRange,
	}
}

func makeMisc(p lit.Publication) bibtex.Reference {
	note := fmt.Sprintf("%q", p.Values)
	return bibtex.Misc{
		Entry: makeEntry(p),
		Note:  &note,
	}
}

func (c Client) ToBibTeX(p lit.Publication) bibtex.Reference {
	// TODO: up to now, I just saw Journal publications with my queries. I
	// expect other types to show up. When that happens, the data stored in
	// the notes of the misc entry will become useful.
	switch p.Values[KeyAggregationType] {
	case "Journal":
		return makeArticle(p)
	default:
		return makeMisc(p)
	}
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
