package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jecoz/edb"
	"github.com/jecoz/lit"
	"github.com/jecoz/lit/bibtex"
)

type MockClient struct {
	maxLit int

	maxLitErr error
	litErr    error

	requestCount int
	pubsCount    int
}

func (c *MockClient) DefaultPerPage() int {
	return 25
}

func (c *MockClient) GetName() string {
	return "mock client"
}

func (c *MockClient) GetRateLimit() time.Duration {
	return time.Millisecond * 1000 / time.Duration(100)
}

func (c *MockClient) GetLiterature(ctx context.Context, r lit.Request) (lit.Response, error) {
	start := r.Page * r.PerPage
	size := c.maxLit - start
	if size < 0 {
		return lit.Response{}, fmt.Errorf("request out of bounds: %d over %d", start, c.maxLit)
	}
	if size > r.PerPage {
		size = r.PerPage
	}

	pubs := make([]lit.Publication, size)
	for i := 0; i < size; i++ {
		pubs[i] = lit.Publication{
			Title:     fmt.Sprintf("pub #%d", i+start),
			CoverDate: time.Now(),
			Creator:   "Ciuck Taylor",
		}
	}

	c.pubsCount += len(pubs)
	c.requestCount++
	return lit.Response{
		Req:        r,
		Literature: pubs,
	}, c.litErr
}

func (c *MockClient) GetMaxLiterature(context.Context, lit.Request) (int, error) {
	return c.maxLit, c.maxLitErr
}

func (c *MockClient) GetAbstract(context.Context, lit.Publication) (lit.Abstract, error) {
	return lit.Abstract{}, fmt.Errorf("get abstract: not implemented")
}

func (c *MockClient) ReferenceLink(lit.Publication) string {
	return "https://nowhere.com"
}

type MockFile struct {
	*bytes.Buffer
}

func (f MockFile) Close() error { return nil }

func mockDb(q string, maxLit int) (*edb.Db, func()) {
	f, err := os.CreateTemp("", "lit")
	if err != nil {
		panic(err)
	}
	db := edb.New(f)
	if err := db.Append(&edb.Event{
		Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Issuer: "testing",
		Scope:  "lit",
		Action: "set_query",
		Data:   []string{q, fmt.Sprintf("%d", maxLit)},
	}); err != nil {
		panic(err)
	}
	return db, func() {
		f.Close()
		os.Remove(f.Name())
	}
}

func mockProgram(t *testing.T, db *edb.Db, client lit.Library) (*tea.Program, io.Writer) {
	inr, inw := io.Pipe()
	p, err := Program(db, client,
		tea.WithoutRenderer(),
		tea.WithoutCatchPanics(),
		tea.WithInput(inr),
		tea.WithOutput(io.Discard),
	)
	if err != nil {
		t.Fatal(err)
	}
	return p, inw
}

func TestMainExitWithMaxLitErr(t *testing.T) {
	t.Parallel()
	maxLit := 2
	maxLitErr := fmt.Errorf("max lit error: unable to do it")
	client := &MockClient{
		maxLit:    maxLit,
		maxLitErr: maxLitErr,
	}
	db, cleanup := mockDb("some q", maxLit)
	defer cleanup()

	t.Run("", func(t *testing.T) {
		if _, err := Program(db, client,
			tea.WithoutRenderer(),
			tea.WithoutCatchPanics(),
		); !errors.Is(err, maxLitErr) {
			t.Fatalf("unexpected error stored in model: want %q, have %q", maxLitErr, err)
		}
	})
}

func TestMainExitWithLitErr(t *testing.T) {
	t.Parallel()
	maxLit := 76
	litErr := fmt.Errorf("lit error: unable to do it")
	client := &MockClient{
		maxLit: maxLit,
		litErr: litErr,
	}
	db, cleanup := mockDb("some q", maxLit)
	defer cleanup()

	t.Run("", func(t *testing.T) {
		p, _ := mockProgram(t, db, client)
		go func() {
			<-time.After(client.timeout())
			p.Quit()
		}()
		i, err := p.StartReturningModel()
		if err != nil {
			t.Fatal(err)
		}
		m := i.(model)
		if !m.done {
			t.Fatalf("program is not done yet")
		}
		if !errors.Is(m.err, litErr) {
			t.Fatalf("unexpected error stored in model: want %q, have %q", litErr, m.err)
		}
	})
}

func (c MockClient) timeout() time.Duration {
	ms := float64(c.maxLit) / float64(c.GetRateLimit().Milliseconds())
	return time.Millisecond * time.Duration(int(ms)*10)
}

func (c MockClient) ToBibTeX(p lit.Publication) bibtex.Reference {
	return bibtex.Misc{
		Entry: bibtex.Entry{
			Title:  p.Title,
			Author: p.Creator,
			Year:   p.CoverDate.Year(),
		},
	}
}

func TestMain(t *testing.T) {
	t.Parallel()
	maxLit := 776
	client := &MockClient{
		maxLit: maxLit,
	}
	db, cleanup := mockDb("some q", maxLit)
	defer cleanup()

	t.Run("", func(t *testing.T) {
		p, _ := mockProgram(t, db, client)
		go func() {
			<-time.After(client.timeout())
			p.Quit()
		}()
		i, err := p.StartReturningModel()
		if err != nil {
			t.Fatal(err)
		}
		m := i.(model)
		if !m.done {
			t.Fatalf("program is not done yet")
		}
		if m.err != nil {
			t.Fatal(err)
		}

		if client.pubsCount != client.maxLit {
			t.Fatalf("pubs count: have %d, want %d", client.pubsCount, client.maxLit)
		}
		expectedRequestCount := int(math.Ceil(float64(client.maxLit) / float64(client.DefaultPerPage())))
		if client.requestCount != expectedRequestCount {
			t.Fatalf("request count: have %d, want %d", client.pubsCount, expectedRequestCount)
		}
	})
}
