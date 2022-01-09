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

func (c Client) GetLiterature(ctx context.Context, req lit.Request) (lit.Response, error) {
	search := newSearchRequest(ctx, c.apiKey, req)
	resp, err := c.httpClient.Do(search)
	if err != nil {
		return lit.Response{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return lit.Response{}, extractError(resp)
	}
	defer resp.Body.Close()

	var p struct {
		Results struct {
			Entries []json.RawMessage `json:"entry"`
		} `json:"search-results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return lit.Response{}, err
	}
	blobs := make([]lit.Blob, len(p.Results.Entries))
	for i, v := range p.Results.Entries {
		blobs[i] = lit.Blob(v)
	}
	return lit.Response{
		Req:   req,
		Blobs: blobs,
	}, nil
}

func (c Client) GetMaxLiterature(ctx context.Context, req lit.Request) (int, error) {
	search := newSearchRequest(ctx, c.apiKey, req)
	resp, err := c.httpClient.Do(search)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, extractError(resp)
	}
	defer resp.Body.Close()

	var p struct {
		Results struct {
			Total string `json:"opensearch:totalResults"`
		} `json:"search-results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(p.Results.Total)
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

func (c Client) PrettyPrint(b lit.Blob, dst *bytes.Buffer) error {
	return json.Indent(dst, []byte(b), "", "\t")
}

func (c Client) ParsePublication(b lit.Blob) (lit.Publication, error) {
	var entry openSearchEntry
	if err := json.Unmarshal([]byte(b), &entry); err != nil {
		return lit.Publication{}, err
	}
	coverDate, err := entry.CoverDate()
	if err != nil {
		return lit.Publication{}, fmt.Errorf("search result %s: %w", entry.Eid, err)
	}

	return lit.Publication{
		Title:     entry.Title,
		CoverDate: coverDate,
		Creator:   entry.Creator,
		Values:    entry.Values(),
	}, nil
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

func getStringPtr(p lit.Publication, key string) *string {
	val, ok := p.Values[key]
	if !ok {
		return nil
	}
	if len(val) == 0 {
		return nil
	}
	return &val
}

func makeEntry(p lit.Publication) bibtex.Entry {
	doi := getStringPtr(p, KeyDOI)
	issn := getStringPtr(p, KeyIssn)
	url := getStringPtr(p, KeyLinkAbstract)

	var abstract *string
	if abs := p.Abstract; abs != nil {
		abstract = &(abs.Text)
	}

	var keywords *string
	if k := p.Keywords; k != nil {
		text := k.Text()
		keywords = &text
	}

	var reason *string
	if rev := p.Review; rev != nil && !rev.IsAccepted {
		reason = &(rev.RejectReason)
	}

	return bibtex.Entry{
		Title:        p.Title,
		Author:       p.Creator,
		Year:         p.CoverDate.Year(),
		DOI:          doi,
		Issn:         issn,
		Url:          url,
		Abstract:     abstract,
		Keywords:     keywords,
		RejectReason: reason,
	}
}

func makeArticle(p lit.Publication) bibtex.Reference {
	pageRange := getStringPtr(p, KeyPageRange)
	volume := getStringPtr(p, KeyVolume)

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

func makeInProceedings(p lit.Publication) bibtex.Reference {
	// TODO: add publisher
	// var publisher *string

	pageRange := getStringPtr(p, KeyPageRange)
	address := getStringPtr(p, KeyAffiliation)

	return bibtex.InProceedings{
		Entry:     makeEntry(p),
		BookTitle: p.Values[KeyPublicationName],
		Address:   address,
		PageRange: pageRange,
	}
}

func makeInCollection(p lit.Publication) bibtex.Reference {
	// TODO: add publisher
	// var publisher *string

	pageRange := getStringPtr(p, KeyPageRange)
	address := getStringPtr(p, KeyAffiliation)

	return bibtex.InCollection{
		Entry:     makeEntry(p),
		BookTitle: p.Values[KeyPublicationName],
		Address:   address,
		PageRange: pageRange,
	}
}

func makeBook(p lit.Publication) bibtex.Reference {
	// TODO: add publisher
	// var publisher *string

	address := getStringPtr(p, KeyAffiliation)

	return bibtex.Book{
		Entry:   makeEntry(p),
		Address: address,
	}
}

func (c Client) ToBibTeX(p lit.Publication) bibtex.Reference {
	switch p.Values[KeyAggregationType] {
	case "Journal", "Trade Journal":
		return makeArticle(p)
	case "Conference Proceeding":
		return makeInProceedings(p)
	case "Book":
		return makeBook(p)
	case "Book Series":
		return makeInCollection(p)
	default:
		return makeMisc(p)
	}
}

func (c Client) ReferenceLink(p lit.Publication) string {
	u, err := url.Parse(p.Values[KeyLinkAbstract])
	if err != nil {
		// No worries
		panic(err)
	}
	q := u.Query()
	q.Set("access_token", c.apiKey)
	u.RawQuery = q.Encode()
	return u.String()
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
