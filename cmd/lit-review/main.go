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
	scopusKey = os.Getenv("SCOPUS_API_KEY")
)

func main() {
	flag.Parse()

	pubs, err := lit.ReadLiterature(os.Stdin)
	if err != nil {
		log.Fatale(err)
	}

	log.Info("Publications read: %d\n", len(pubs))

	client := scopus.NewClient(scopusKey)
	ctx := context.Background()

	for _, pub := range pubs {
		if err := pub.GetAbstract(ctx, client); err != nil {
			log.Fatale(err)
		}
		if err := pub.Abstract.WriteTo(os.Stdout); err != nil {
			log.Fatale(err)
		}
	}
}
