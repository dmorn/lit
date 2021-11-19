package main

import (
	"context"
	"flag"
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

	pubChan := lit.GetLiterature(context.Background(), client, lit.Request{
		Query: *queryString,
	})

	received := 0
	for pub := range pubChan.Chan {
		received++
		log.Event("main", log.Measurement{
			"received_count": received,
			"left_count":     pubChan.Total - received,
			"total_count":    pubChan.Total,
		}, nil)

		if err := pub.WriteTo(os.Stdout); err != nil {
			log.Fatale(err)
		}
	}
	if err := pubChan.Err; err != nil {
		log.Fatale(err)
	}
}
