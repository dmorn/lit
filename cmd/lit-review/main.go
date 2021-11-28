package main

import (
	"context"
	"flag"
	"io"
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
	pub := pubs[0]

	body, err := client.GetLink(context.Background(), pub.Links["scopus"])
	if err != nil {
		log.Fatale(err)
	}
	defer body.Close()
	io.Copy(os.Stdout, body)
}
