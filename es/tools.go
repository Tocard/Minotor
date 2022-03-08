package es

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"context"
	"encoding/json"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	wg     sync.WaitGroup
	client *elasticsearch.Client
)

func Connection() {
	cert, err := ioutil.ReadFile(config.Cfg.CaPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", config.Cfg.CaPath, err)
	}
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Cfg.ElasticsearchHosts,
		Username:  config.Cfg.ElasticsearchUser,
		Password:  config.Cfg.ElasticsearchPassword,
		CACert:    cert,
		Logger:    &elastictransport.TextLogger{Output: os.Stdout},
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

func Write(index string) {
	for i := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"@timestamp" : "`)
			b.WriteString(data.GetTimestampAsString(data.Clock.Unix()))
			b.WriteString(`"}`)
			log.Printf(b.String())
			// Set up the request object.
			req := esapi.IndexRequest{
				Index: index,
				Body:  strings.NewReader(b.String()),
			}

			// Perform the request with the client.

			res, err := req.Do(context.Background(), client)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document %s", res.Status(), res.String())
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))
}

func Bulk(index, data string) {
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     client,
		Index:      index,
		NumWorkers: 1,
	})
	indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action: "create",
			Body:   strings.NewReader(data),
			// OnFailure is the optional callback for each failed operation
			OnFailure: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem, err error,
			) {
				if err != nil {
					log.Fatalf("ERROR: %s on index %s with data %s", err, index, data)
				} else {
					log.Fatalf("ERROR: %s: %s on index %s with data %s", res.Error.Type, res.Error.Reason, index, data)
				}
			},
		})
	indexer.Close(context.Background())

}

//func Bulk(index, bulkData string) {
//	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
//		Client:     client,
//		Index:      index,
//		NumWorkers: 1,
//	})
//	indexer.Add(
//		context.Background(),
//		esutil.BulkIndexerItem{
//			Action: "create",
//			Body:   strings.NewReader(bulkData),
//		})
//	indexer.Close(context.Background())
//
//}
