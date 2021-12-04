package main

import (
	"bytes"
	"errors"
	"testing"
	//"io"
	"os"
	//"math"
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jecoz/edb"
	"github.com/jecoz/lit"
)

type MockClient struct {
	t      *testing.T
	maxCC  int
	maxLit int

	maxLitErr error
	litErr    error

	requestCount int
	pubsCount    int
}

func (c *MockClient) GetName() string {
	return "mock client"
}

func (c *MockClient) ConcurrencyLimit() int {
	max := c.maxCC
	if max < 1 {
		max = 1
	}
	return max
}

func (c *MockClient) GetLiterature(ctx context.Context, r lit.Request) (lit.Response, error) {
	c.t.Logf("mock: get literature called: %+v", r)
	c.t.Logf("mock: get literature called (error): %v", c.litErr)
	start := r.Page * r.PerPage
	size := c.maxLit - start
	if size < 0 {
		return lit.Response{}, fmt.Errorf("request out of bounds: %d over %d", start, c.maxLit)
	}
	if size > r.MaxPage {
		size = r.MaxPage
	}

	pubs := make([]lit.Publication, size)
	for i := 0; i < size; i++ {
		pubs[i] = lit.Publication{
			Title:        fmt.Sprintf("pub #%d", i+start),
			Eid:          "eid",
			Issn:         "issn",
			CoverDate:    time.Now(),
			Creator:      "Ciuck Taylor",
			LinkAbstract: "missing.com",
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

func TestMainExitWithMaxLitErr(t *testing.T) {
	maxLit := 2
	maxLitErr := fmt.Errorf("max lit error: unable to do it")
	client := &MockClient{
		t:         t,
		maxCC:     4,
		maxLit:    maxLit,
		maxLitErr: maxLitErr,
	}
	db, cleanup := mockDb("some q", maxLit)
	defer cleanup()

	t.Run("", func(t *testing.T) {
		_, err := Program(db, client,
			tea.WithoutRenderer(),
			tea.WithoutCatchPanics(),
		)
		if !errors.Is(err, maxLitErr) {
			t.Fatalf("unexpected error stored in model: want %q, have %q", maxLitErr, err)
		}
	})
}

func TestMainExitWithLitErr(t *testing.T) {
	maxLit := 76
	litErr := fmt.Errorf("lit error: unable to do it")
	client := &MockClient{
		t:      t,
		maxCC:  4,
		maxLit: maxLit,
		litErr: litErr,
	}
	db, cleanup := mockDb("some q", maxLit)
	defer cleanup()

	p, _ := Program(db, client,
		tea.WithoutRenderer(),
		tea.WithoutCatchPanics(),
	)

	t.Run("", func(t *testing.T) {
		go func() {
			<-time.After(500 * time.Millisecond)
			p.Quit()
		}()
		i, _ := p.StartReturningModel()
		m := i.(model)
		if !errors.Is(m.err, litErr) {
			t.Fatalf("unexpected error stored in model: want %q, have %q", litErr, m.err)
		}
	})
}

/*
func TestMain(t *testing.T) {
	client := &MockClient{
		t: t,
		maxCC: 4,
		maxLit: 100 * 4,
	}

	db := edb.New(f)
	if err := db.Append(&edb.Event{
		Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Issuer: "testing",
		Scope:  "lit",
		Action: "set_query",
		Data:   []string{"a query", fmt.Sprintf("%d", client.maxLit)},
	}); err != nil {
		t.Fatal(err)
	}

	inputR, inputW := io.Pipe()
	outputR, outputW := io.Pipe()
	p, err := Program(db, client,
		tea.WithInput(inputR),
		tea.WithOutput(outputW),
		tea.WithoutRenderer(),
		tea.WithoutCatchPanics(),
	)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		<-time.After(1*time.Second)
		fmt.Fprintf(inputW, "q")
	}()
	go func() {
		io.Copy(io.Discard, outputR)
	}()

	if _, err := p.StartReturningModel(); err != nil {
		t.Fatal(err)
	}

	if client.pubsCount != client.maxLit {
		t.Fatalf("pubs count: have %d, want %d", client.pubsCount, client.maxLit)
	}
	expectedRequestCount := int(math.Ceil(float64(client.maxLit) / float64(lit.DefaultPerPage)))
	if client.requestCount != expectedRequestCount {
		t.Fatalf("request count: have %d, want %d", client.pubsCount, expectedRequestCount)
	}
}
*/
