package es

import (
	"2miner-monitoring/config"
	"context"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	client   *elasticsearch.Client
	Indexers = make(map[string]esutil.BulkIndexer, 20)
)

func Connection() {
	cert, err := ioutil.ReadFile(config.Cfg.CaPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", config.Cfg.CaPath, err)
	}
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     config.Cfg.ElasticsearchHosts,
		Username:      config.Cfg.ElasticsearchUser,
		Password:      config.Cfg.ElasticsearchPassword,
		CACert:        cert,
		Logger:        &elastictransport.TextLogger{Output: os.Stdout},
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff:  func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
		MaxRetries:    5,
	})

	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}

func Health() {
	res, err := client.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}
	log.Println(res)
}

//func Write(index string) {
//	for i := range []string{"Test One", "Test Two"} {
//		wg.Add(1)
//
//		go func(i int) {
//			defer wg.Done()
//
//			// Build the request body.
//			var b strings.Builder
//			b.WriteString(`{"@timestamp" : "`)
//			b.WriteString(data.GetTimestampAsString(data.Clock.Unix()))
//			b.WriteString(`"}`)
//			log.Printf(b.String())
//			// Set up the request object.
//			req := esapi.IndexRequest{
//				Index: index,
//				Body:  strings.NewReader(b.String()),
//			}
//
//			// Perform the request with the client.
//
//			res, err := req.Do(context.Background(), client)
//			if err != nil {
//				log.Fatalf("Error getting response: %s", err)
//			}
//			defer res.Body.Close()
//
//			if res.IsError() {
//				log.Printf("[%s] Error indexing document %s", res.Status(), res.String())
//			} else {
//				// Deserialize the response into a map.
//				var r map[string]interface{}
//				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//					log.Printf("Error parsing the response body: %s", err)
//				} else {
//					// Print the response status and indexed document version.
//					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
//				}
//			}
//		}(i)
//	}
//	wg.Wait()
//
//	log.Println(strings.Repeat("-", 37))
//}

func Bulk(index, data string) {
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        client,
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
		log.Fatalf("Unexpected error: %s", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := indexer.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	//
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		log.Fatalf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
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

func EsSearch() *esapi.Response {
	body := `{
  "query": {
    "exists":{ 
      "field":"id"
      }
  },
  "aggs": {
    "by_id": {
      "terms": {
        "field": "id"
        , "size": 1
      }
    }
  }
}`
	indx := []string{"2miners-hiveos-farm"}
	req := esapi.SearchRequest{
		Index:   indx,
		Body:    strings.NewReader(body),
		Size:    esapi.IntPtr(25),
		Pretty:  true,
		Timeout: 100,
	}
	Esdata, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Unexpected error when getting a response: %s", err)
	} else {
		log.Printf("%s", Esdata)
		return Esdata
	}
	return nil
}

func CreatebulkIndexer() {
	Indexs := []string{"2miners-flux-node", "2miners-balance", "2miners-tx", "flux-node-overview", "2miners-coins",
		"2miners-hiveos-hashrate-coin", "2miners-hiveos-hashrate", "2miners-hiveos-farm", "2miners-hiveos-flightsheet",
		"2miners-hiveos-gpu", "2miners-hiveos-gpu-total-info", "2miners-hiveos-worker", "2miners-worker",
		"2miners-data", "2miners-sumreward", "2miners-reward", "2miners-payment", "2miners-stat", "2miners-poolstat",
		"2miners-hashrate_no"}

	for _, index := range Indexs {
		bulkindexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Client:        client, // The Elasticsearch client
			Index:         index,  // The default index name
			NumWorkers:    4,      // The number of worker goroutines (default: number of CPUs)
			FlushBytes:    5e+6,   // The flush threshold in bytes (default: 5M)
			FlushInterval: 60,
		})
		if err != nil {
			log.Fatalf("Error creating the indexer: %s", err)
		}
		Indexers[index] = bulkindexer
		log.Printf("Added new indexer %s\n to list %s", bulkindexer, Indexers)
	}
}

func NewBulk(index, data string) {
	log.Printf("###List indexer on Bulk %s\n", Indexers)
	if Indexers[index] == nil {
		return
	}
	log.Printf("***Indexer on Bulk %s\n", Indexers[index])
	err := Indexers[index].Add(
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
		log.Fatalf("Unexpected error: %s", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := Indexers[index].Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	//
	stats := Indexers[index].Stats()
	if stats.NumFailed > 0 {
		log.Fatalf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
	} else {
		log.Printf("Successfully indexed [%d] documents", stats.NumFlushed)
	}

	// For optimal performance, consider using a third-party package for JSON decoding and HTTP transport.
	//
	// For more information, examples and benchmarks, see:
	//
	// --> https://github.com/elastic/go-elasticsearch/tree/master/_examples/bulk
}
