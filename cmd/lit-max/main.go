package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jecoz/lit"
	"github.com/jecoz/lit/log"
	"github.com/jecoz/lit/scopus"
)

var (
	scopusKey    = os.Getenv("SCOPUS_API_KEY")
	defaultQuery = os.Getenv("QUERY")
)

var (
	queryString = flag.String("q", defaultQuery, "Query string")
)

func main() {
	flag.Parse()

	query := *queryString
	if query == "" {
		log.Fatalf("no query string provided")
	}

	client := scopus.NewClient(scopusKey)

	max, err := lit.GetMaxLiterature(context.Background(), client, lit.Request{
		Query: *queryString,
	})
	if err != nil {
		log.Fatale(err)
	}
	fmt.Fprintf(os.Stdout, "query=%q\n", query)
	fmt.Fprintf(os.Stdout, "results_count=%d\n", max)
}
