package lit

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
	"unicode"

	"golang.org/x/sync/errgroup"
)

const (
	DefaultPerPage = 25
)

type Abstract struct {
	Text string `json:"text"`
}

func (a Abstract) GetText() string {
	return strings.ReplaceAll(a.Text, "\n", "")
}

type Review struct {
	IsAccepted    bool   `json:"is_accepted"`
	IsHighlighted bool   `json:"is_highlighted"`
	RejectReason  string `json:"reject_reason"`
}

type Publication struct {
	Title     string    `json:"title"`
	Eid       string    `json:"eid"`
	Issn      string    `json:"issn"`
	CoverDate time.Time `json:"cover_date"`
	Creator   string    `json:"creator"`

	LinkAbstract string `json:"link_abstract"`

	Aggregation string `json:"prism:aggregationType"`

	*Abstract `json:"abstract,omitempty"`
	*Review   `json:"review,omitempty"`
}

func takeMaxRunes(s string, n int) string {
	s = strings.TrimSpace(s)
	r := strings.NewReader(s)
	take := 0
	for i := 0; i < n; i++ {
		ru, size, err := r.ReadRune()
		if err != nil || !unicode.IsLetter(ru) {
			break
		}
		take += size
	}
	return string([]byte(s)[:take])
}

func (p Publication) CreatorShort() string {
	return takeMaxRunes(p.Creator, 3)
}

func (p Publication) TitleShort() string {
	return takeMaxRunes(p.Title, 5)
}

func (p Publication) BibId() string {
	return fmt.Sprintf("%s%d%s", p.CreatorShort(), p.CoverDate.Year(), p.TitleShort())
}

func (p Publication) Marshal() (string, error) {
	var buf bytes.Buffer
	if err := writeTo(&buf, &p); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (p *Publication) Unmarshal(data string) error {
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	return json.NewDecoder(r).Decode(p)
}

func writeTo(w io.Writer, i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}
	baseEnc := base64.NewEncoder(base64.StdEncoding, w)
	if _, err := baseEnc.Write(data); err != nil {
		return err
	}
	return baseEnc.Close()
}

func (p *Publication) GetAbstract(ctx context.Context, lib Library) error {
	abs, err := lib.GetAbstract(ctx, *p)
	if err != nil {
		return err
	}
	p.Abstract = &abs
	return nil
}

type PublicationChan struct {
	// queue of publications. Closed when no more will be delivered.
	queue chan Publication

	// Once the Chan is open, max tells the number of publications to
	// expect from it.
	max int

	// err is only available after queue was closed.
	err error
}

func (c *PublicationChan) Recv() <-chan Publication {
	return c.queue
}

func (c *PublicationChan) Send(p Publication) error {
	c.queue <- p
	return nil
}

func (c *PublicationChan) Max() int {
	return c.max
}

func (c *PublicationChan) CloseWithError(err error) {
	c.err = err
	close(c.queue)
}

func (c *PublicationChan) Err() error {
	return c.err
}

func NewPublicationChan(max, queueLen int) *PublicationChan {
	return &PublicationChan{
		queue: make(chan Publication, queueLen),
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
	Req        Request
	Literature []Publication
}

func (r Response) Len() int {
	return len(r.Literature)
}

func (r Response) IsEmpty() bool {
	return r.Len() == 0
}

type Library interface {
	GetName() string
	GetRateLimit() time.Duration
	GetLiterature(context.Context, Request) (Response, error)
	GetMaxLiterature(context.Context, Request) (int, error)
	GetAbstract(context.Context, Publication) (Abstract, error)
}

func searchLoop(ctx context.Context, pubChan *PublicationChan, lib Library, req Request) {
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
			for _, pub := range resp.Literature {
				pubChan.Send(pub)
			}
			return nil
		})
	}
	pubChan.CloseWithError(g.Wait())
}

func GetLiterature(ctx context.Context, pubChan *PublicationChan, lib Library, req Request) {
	if req.PerPage <= 0 {
		req.PerPage = DefaultPerPage
	}
	req.MaxResults = pubChan.Max()

	searchLoop(ctx, pubChan, lib, req)
}

func GetMaxLiterature(ctx context.Context, lib Library, req Request) (int, error) {
	return lib.GetMaxLiterature(ctx, req)
}
