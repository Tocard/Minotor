package es

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"strings"
)

func Bulk(index, data string) {
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        EsClient,
		Index:         index,
		NumWorkers:    4,
		FlushBytes:    10000000,
		FlushInterval: 60,
	})
	err := indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action: "create",
			Body:   strings.NewReader(data),

			// OnSuccess is the optional callback for each successful operation
			OnSuccess: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem,
			) {
				fmt.Printf("[%d] %s test/%s", res.Status, res.Result, item.DocumentID)
			},

			// OnFailure is the optional callback for each failed operation
			OnFailure: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem, err error,
			) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		},
	)
	if err != nil {
		log.Printf("Unexpected error: %s\n", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := indexer.Close(context.Background()); err != nil {
		log.Printf("Unexpected error: %s\n", err)
	}

	// Report the indexer statistics
	//
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		log.Printf("Indexed [%d] documents with [%d] errors\n", stats.NumFlushed, stats.NumFailed)
	} else {
		log.Printf("Successfully indexed [%d] documents", stats.NumFlushed)
	}

	// For optimal performance, consider using a third-party package for JSON decoding and HTTP transport.
	//
	// For more information, examples and benchmarks, see:
	//
	// --> https://github.com/elastic/go-elasticsearch/tree/master/_examples/bulk
	//indexer.Close(context.Background())

}
