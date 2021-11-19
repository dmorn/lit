package scopus

import (
	"context"
	"fmt"

	"github.com/jecoz/lit"
)

type Client struct {
	apiKey string
}

func (c Client) GetLiterature(ctx context.Context, req lit.Request) (lit.Response, error) {
	return lit.Response{}, fmt.Errorf("not implemented")
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
	}
}
