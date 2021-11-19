package lit

import (
	"context"
	"fmt"
	"io"
	"math"
	"sync"

	"github.com/jecoz/lit/log"
)

const (
	DefaultPerPage = 25
)

const (
	TagTitle = "#pub:title[%d]:%s"
	TagLink  = "#pub:link[%d]:%s[%d]:%s"
)

type Publication struct {
	Title string
	Links map[string]string
}

func (p Publication) WriteTo(w io.Writer) error {
	fmt.Fprintf(w, TagTitle+"\n", len(p.Title), p.Title)
	for k, v := range p.Links {
		fmt.Fprintf(w, TagLink+"\n", len(k), k, len(v), v)
	}
	fmt.Fprintln(w, "")
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
	ConcurrencyLimit() int
	GetLiterature(context.Context, Request) (Response, error)
	GetMaxLiterature(context.Context, Request) (int, error)
}

func searchLoop(ctx context.Context, lib Library, req Request, pubChan chan<- Publication) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var onceErr error
	var once sync.Once

	cc := make(chan struct{}, lib.ConcurrencyLimit())

	for i := 0; i < req.MaxPage; i++ {
		go func(i int) {
			cc <- struct{}{}
			defer func() { <-cc }()

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

func SearchLiterature(ctx context.Context, lib Library, req Request) *PublicationChan {
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

	log.Event("lit.SearchLiterature", log.Measurement{
		"per_page":  req.PerPage,
		"max_page":  req.MaxPage,
		"max_items": pc.Total,
	}, log.Metadata{
		"query": req.Query,
	})

	go func() {
		defer close(pubChan)
		pc.Err = searchLoop(ctx, lib, req, pubChan)
	}()

	return pc
}
