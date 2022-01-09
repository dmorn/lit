package lit

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jecoz/lit/bibtex"
	"golang.org/x/sync/errgroup"
)

func marshal(s string) (string, error) {
	var buf bytes.Buffer
	baseEnc := base64.NewEncoder(base64.StdEncoding, &buf)
	if _, err := baseEnc.Write([]byte(s)); err != nil {
		return "", err
	}
	if err := baseEnc.Close(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func unmarshal(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Blob []byte

func (b Blob) Marshal() (string, error) {
	return marshal(string(b))
}

func (b *Blob) Unmarshal(data string) error {
	d, err := unmarshal(data)
	if err != nil {
		return err
	}
	*b = Blob(d)
	return nil
}

type Abstract struct {
	Text string `json:"text"`
}

func (a Abstract) GetText() string {
	return strings.ReplaceAll(a.Text, "\n", "")
}

func (a Abstract) Marshal() (string, error) {
	return marshal(a.Text)
}

func (a *Abstract) Unmarshal(data string) error {
	d, err := unmarshal(data)
	if err != nil {
		return err
	}
	a.Text = d
	return nil
}

type Review struct {
	IsAccepted    bool   `json:"is_accepted"`
	IsHighlighted bool   `json:"is_highlighted"`
	RejectReason  string `json:"reject_reason"`
}

func (r Review) Marshal() (string, error) {
	d, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return marshal(string(d))
}

func (r *Review) Unmarshal(data string) error {
	d, err := unmarshal(data)
	if err != nil {
		return err
	}
	return json.NewDecoder(strings.NewReader(d)).Decode(r)
}

type Keywords struct {
	Values []string `json:"values"`
}

func (k Keywords) Text() string {
	return strings.Join(k.Values, ", ")
}

func (k *Keywords) Parse(input string) {
	values := strings.Split(input, ",")
	trimmed := make([]string, len(values))
	for i, v := range values {
		trimmed[i] = strings.TrimSpace(v)
	}
	k.Values = trimmed
}

func (k Keywords) Marshal() (string, error) {
	d, err := json.Marshal(k)
	if err != nil {
		return "", err
	}
	return marshal(string(d))
}

func (k *Keywords) Unmarshal(data string) error {
	d, err := unmarshal(data)
	if err != nil {
		return err
	}
	return json.NewDecoder(strings.NewReader(d)).Decode(k)
}

type Publication struct {
	Title     string    `json:"title"`
	CoverDate time.Time `json:"cover_date"`
	Creator   string    `json:"creator"`

	// Misc stuff used by clients to accomplish Library interface.
	Values map[string]string `json:"values"`

	*Abstract `json:"abstract,omitempty"`
	*Review   `json:"review,omitempty"`
	*Keywords `json:"keywords,omitempty"`
}

func (p *Publication) GetAbstract(ctx context.Context, lib Library) error {
	abs, err := lib.GetAbstract(ctx, *p)
	if err != nil {
		return err
	}
	p.Abstract = &abs
	return nil
}

type BlobChan struct {
	// queue of publications. Closed when no more will be delivered.
	queue chan Blob

	// Once the Chan is open, max tells the number of publications to
	// expect from it.
	max int

	// err is only available after queue was closed.
	err error
}

func (c *BlobChan) Recv() <-chan Blob {
	return c.queue
}

func (c *BlobChan) Send(p Blob) error {
	c.queue <- p
	return nil
}

func (c *BlobChan) Max() int {
	return c.max
}

func (c *BlobChan) CloseWithError(err error) {
	c.err = err
	close(c.queue)
}

func (c *BlobChan) Err() error {
	return c.err
}

func NewBlobChan(max, queueLen int) *BlobChan {
	return &BlobChan{
		queue: make(chan Blob, queueLen),
		max:   max,
	}
}

type Request struct {
	Query string

	Page       int
	PerPage    int
	MaxResults int
}

func (r Request) RoundsNeeded() int {
	return int(math.Ceil(float64(r.MaxResults) / float64(r.PerPage)))
}

func (r Request) CloneWithPage(p int) Request {
	o := new(Request)
	*o = r
	o.Page = p
	return *o
}

func (r Request) String() string {
	from := r.Page * r.PerPage
	to := from + r.PerPage
	if to > r.MaxResults {
		to = r.MaxResults
	}
	return fmt.Sprintf("{q=%q page=%d from=%d to=%d max=%d}", r.Query, r.Page, from, to, r.MaxResults)
}

type Response struct {
	Req   Request
	Blobs []Blob
}

func (r Response) Len() int {
	return len(r.Blobs)
}

func (r Response) IsEmpty() bool {
	return r.Len() == 0
}

type Library interface {
	GetName() string
	GetRateLimit() time.Duration
	DefaultPerPage() int

	GetLiterature(context.Context, Request) (Response, error)
	GetMaxLiterature(context.Context, Request) (int, error)
	ParsePublication(Blob) (Publication, error)
	PrettyPrint(Blob, *bytes.Buffer) error
	GetAbstract(context.Context, Publication) (Abstract, error)

	ToBibTeX(Publication) bibtex.Reference
	ReferenceLink(Publication) string
}

func searchLoop(ctx context.Context, blobChan *BlobChan, lib Library, req Request) {
	requests := make(chan Request, req.RoundsNeeded())
	for i := 0; i < req.RoundsNeeded(); i++ {
		requests <- req.CloneWithPage(i)
	}
	close(requests)

	limiter := time.Tick(lib.GetRateLimit())

	g, ctx := errgroup.WithContext(ctx)
	for r := range requests {
		<-limiter
		r := r
		g.Go(func() error {
			resp, err := lib.GetLiterature(ctx, r)
			if err != nil {
				return fmt.Errorf("get literature: page %d: %w", r.Page, err)
			}
			for _, blob := range resp.Blobs {
				blobChan.Send(blob)
			}
			return nil
		})
	}
	blobChan.CloseWithError(g.Wait())
}

func GetLiterature(ctx context.Context, blobChan *BlobChan, lib Library, req Request) {
	req.MaxResults = blobChan.Max()
	if req.PerPage <= 0 {
		req.PerPage = lib.DefaultPerPage()
	}

	searchLoop(ctx, blobChan, lib, req)
}
