package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/jecoz/lit"
	"github.com/jecoz/lit/scopus"
)

var (
	scopusKey    = os.Getenv("SCOPUS_API_KEY")
	defaultQuery = os.Getenv("QUERY")
)

var (
	queryString = flag.String("q", defaultQuery, "Query string")
	queryPath   = flag.String("Q", "", "Query file path")
)

func main() {
	flag.Parse()

	query := *queryString
	if query == "" {
		log.Fatalf("no query string provided")
	}

	client := scopus.NewClient(scopusKey)

	pubs := []lit.Publication{}
	pubChan := lit.SearchLiterature(context.Background(), client, lit.Request{
		Query: *queryString,
	})
	for pub := range pubChan.Chan {
		pubs = append(pubs, pub)
	}
	if err := pubChan.Err; err != nil {
		log.Fatal(err)
	}

	p, err := json.Marshal(pubs)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, p, "", "\t"); err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(os.Stdout)
}
