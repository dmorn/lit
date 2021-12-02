package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jecoz/edb"
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
	edbPath     = flag.String("edb", "lit.edb", "Event database file. Everything will be stored here.")
)

func Main() error {
	query := *queryString
	if query == "" {
		return fmt.Errorf("no query string provided")
	}

	os.Remove(*edbPath)
	db, err := edb.Open(*edbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	client := scopus.NewClient(scopusKey)
	pubChan := lit.GetLiterature(context.Background(), client, lit.Request{
		Query: *queryString,
	})

	if err := db.Append(&edb.Event{
		Id:     fmt.Sprintf("query"),
		Issuer: "lit",
		Scope:  "lit",
		Action: "set_query",
		Data:   []string{*queryString},
	}); err != nil {
		return err
	}

	received := 0
	for pub := range pubChan.Chan {
		received++
		log.Event("main", log.Measurement{
			"received_count": received,
			"left_count":     pubChan.Total - received,
			"total_count":    pubChan.Total,
		}, nil)
		data, err := pub.Marshal()
		if err != nil {
			return err
		}
		if err := db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", received),
			Issuer: "scopus",
			Scope:  "lit",
			Action: "add_lit",
			Data:   []string{data},
		}); err != nil {
			return err
		}
	}
	return pubChan.Err
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
