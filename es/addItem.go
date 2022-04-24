package es

import (
	"bytes"
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"sync/atomic"
)

func AddIemIndexer(bi esutil.BulkIndexer, data []byte, countSuccessful *uint64) {

	err := bi.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action: "create",
			Body:   bytes.NewReader(data),

			// OnSuccess is called for each successful operation
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
				atomic.AddUint64(countSuccessful, 1)
			},

			// OnFailure is called for each failed operation
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		},
	)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
}
