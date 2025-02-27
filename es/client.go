package es

import (
	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"log"
	"minotor/config"
	"os"
	"time"
)

var (
	EsClient *elasticsearch.Client
)

func Connection() {
	retryBackoff := backoff.NewExponentialBackOff()
	cert, err := ioutil.ReadFile(config.Cfg.CaPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", config.Cfg.CaPath, err)
	}
	EsClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     config.Cfg.ElasticsearchHosts,
		Username:      config.Cfg.ElasticsearchUser,
		Password:      config.Cfg.ElasticsearchPassword,
		CACert:        cert,
		Logger:        &elastictransport.TextLogger{Output: os.Stdout},
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		}, MaxRetries: 5,
	})

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
}
