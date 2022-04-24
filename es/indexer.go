package es

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"time"
)

var (
	Indexers = make(map[string]esutil.BulkIndexer)
)

func CreateIndexer(indexName string) {

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,        // The default index name
		Client:        EsClient,         // The Elasticsearch client
		NumWorkers:    8,                // The number of worker goroutines
		FlushBytes:    int(5e+6),        // The flush threshold in bytes
		FlushInterval: 60 * time.Second, // The periodic flush interval
		// OnFlushStart is called when flush is trigger
		OnFlushStart: func(context.Context) context.Context {
			log.Println("Flush is trigger")
			return nil
		},
		OnFlushEnd: func(context.Context) {
			log.Println("End Flush period")
		},
		OnError: func(ctx context.Context, err error) {

			log.Printf("Error on indexer %s\n", err.Error())
		},
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	Indexers[indexName] = bi
}
