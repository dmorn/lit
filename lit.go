package lit

import (
	"context"
	"fmt"
)

type Publication struct {
	Title    string
	Abstract string
}

type PublicationChan struct {
	// Chan of publications. Closed when no more will be delivered.
	Chan <-chan Publication

	// Err is only available after Chan was closed.
	Err error
}

type Request struct {
	Query  string
	Cursor string
}

type Response struct {
	Literature []Publication
	NextCursor string
}

func (r Response) IsEmpty() bool {
	return len(r.Literature) == 0
}

type Library interface {
	GetLiterature(context.Context, Request) (Response, error)
}

func SearchLiterature(ctx context.Context, lib Library, req Request) *PublicationChan {
	pubChan := make(chan Publication)
	pc := &PublicationChan{
		Chan: pubChan,
	}

	go func() {
		defer close(pubChan)
		for {
			resp, err := lib.GetLiterature(ctx, req)
			if err != nil {
				pc.Err = fmt.Errorf("get literature: %w", err)
				return
			}
			if resp.IsEmpty() {
				return
			}
			for _, pub := range resp.Literature {
				pubChan <- pub
			}
			req.Cursor = resp.NextCursor
		}
	}()

	return pc
}
