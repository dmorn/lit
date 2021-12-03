package lit

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"
	"unicode"
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

func (a Abstract) WriteTo(w io.Writer) error {
	return writeTo(w, a)
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

	*Abstract `json:"abstract,omitempty"`
	*Review   `json:"review,omitempty"`
}

func takeMaxRunes(s string, n int) string {
	s = strings.TrimSpace(s)
	r := strings.NewReader(s)
	take := 0
	for i := 0; i < n; i++ {
		ru, size, err := r.ReadRune()
		if err != nil || unicode.IsSpace(ru) {
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
	return json.Unmarshal([]byte(data), p)
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

func (p Publication) WriteTo(w io.Writer) error {
	return writeTo(w, p)
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
	// Chan of publications. Closed when no more will be delivered.
	Chan <-chan Publication

	// Once the Chan is open, Total tells the number of publications to
	// expect from it.
	Total int

	// Err is only available after Chan was closed.
	Err error
}

type Request struct {
	Query string

	Page    int
	PerPage int
	MaxPage int
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
	ConcurrencyLimit() int
	GetLiterature(context.Context, Request) (Response, error)
	GetMaxLiterature(context.Context, Request) (int, error)
	GetAbstract(context.Context, Publication) (Abstract, error)
}

func searchLoop(ctx context.Context, lib Library, req Request, pubChan chan<- Publication) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var onceErr error
	var once sync.Once

	cc := make(chan struct{}, lib.ConcurrencyLimit())

	for i := 0; i < req.MaxPage; i++ {
		go func(i int) {
			select {
			case cc <- struct{}{}:
				defer func() { <-cc }()
			case <-ctx.Done():
				return
			}

			req.Page = i
			resp, err := lib.GetLiterature(ctx, req)
			if err != nil {
				once.Do(func() {
					onceErr = fmt.Errorf("get literature: page %d: %w", i, err)
					cancel()
				})
				return
			}
			for _, pub := range resp.Literature {
				pubChan <- pub
			}
		}(i)
	}
	for i := 0; i < cap(cc); i++ {
		cc <- struct{}{}
	}

	return onceErr
}

func GetLiterature(ctx context.Context, lib Library, req Request) *PublicationChan {
	pubChan := make(chan Publication)
	pc := &PublicationChan{
		Chan: pubChan,
	}

	if req.PerPage <= 0 {
		req.PerPage = DefaultPerPage
	}

	maxItems, err := lib.GetMaxLiterature(ctx, req)
	if err != nil {
		pc.Err = err
		close(pubChan)
		return pc
	}

	req.MaxPage = int(math.Ceil(float64(maxItems) / float64(req.PerPage)))
	pc.Total = maxItems

	go func() {
		defer close(pubChan)
		pc.Err = searchLoop(ctx, lib, req, pubChan)
	}()

	return pc
}

func GetMaxLiterature(ctx context.Context, lib Library, req Request) (int, error) {
	return lib.GetMaxLiterature(ctx, req)
}

func ReadLiterature(r io.Reader) ([]Publication, error) {
	scan := bufio.NewScanner(r)
	acc := []Publication{}
	for scan.Scan() {
		var p Publication
		if err := json.Unmarshal(scan.Bytes(), &p); err != nil {
			return acc, fmt.Errorf("unexpected publication line: %w", err)
		}
		acc = append(acc, p)
	}
	if err := scan.Err(); err != nil {
		return acc, err
	}
	return acc, nil
}
