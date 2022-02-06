package es

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/utils"
	"context"
	"encoding/json"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	r      map[string]interface{}
	wg     sync.WaitGroup
	client *elasticsearch.Client
	m1     = regexp.MustCompile(`workers\.(.*)\.hr`)
	m2     = regexp.MustCompile(`workers\.(.*)\.hr2`)
	m3     = regexp.MustCompile(`workers\.(.*)\.lastBeat`)
	m4     = regexp.MustCompile(`workers\.(.*)\.offline`)
	m5     = regexp.MustCompile(`workers\.(.*)\.rhr`)
	m6     = regexp.MustCompile(`workers\.(.*)\.sharesInvalid`)
	m7     = regexp.MustCompile(`workers\.(.*)\.sharesStale`)
	m8     = regexp.MustCompile(`workers\.(.*)\.sharesValid`)
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

func ExtractWorkerInfo(workers map[string]interface{}, wallet string) {
	var WorkerArray []data.Worker
	for key, _ := range workers {
		miner := workers[key].(map[string]interface{})
		tmpMiner := data.Worker{}
		tmpMiner.Name = key
		for minerKey, value := range miner {
			if minerKey == "hr" {
				tmpMiner.Hr = value.(float64)
			} else if minerKey == "offline" {
				tmpMiner.Offline = value.(bool)
			} else if minerKey == "Hr2" {
				tmpMiner.Hr2 = value.(float64)
			} else if minerKey == "lastBeat" {
				tmpMiner.LastBeat = value.(float64)
			} else if minerKey == "sharesValid" {
				tmpMiner.SharesValid = value.(float64)
			} else if minerKey == "sharesInvalid" {
				tmpMiner.SharesInvalid = value.(float64)
			} else if minerKey == "sharesStale" {
				tmpMiner.SharesStale = value.(float64)
			}
			tmpMiner.Wallet = wallet
			tmpMiner.Timestamp = time.Now().Format(time.RFC3339)
		}
		tmpMinerJson, _ := json.Marshal(tmpMiner)
		Bulk("2miners-worker", string(tmpMinerJson))

		WorkerArray = append(WorkerArray, tmpMiner)
	}
}

func ExtractSimpleField(esBulk *data.MinerInfo, json map[string]interface{}, wallet string) {
	esBulk.TwoMiners.Wallet = wallet
	esBulk.Timestamp = time.Now().Format(time.RFC3339)
	esBulk.Two4Hnumreward = json["24hnumreward"].(float64)
	esBulk.Two4Hreward = json["24hreward"].(float64)
	esBulk.APIVersion = json["apiVersion"].(float64)
	if json["allowedMaxPayout"] != nil {
		esBulk.AllowedMaxPayout = json["allowedMaxPayout"].(int64)
	}
	if json["allowedMinPayout"] != nil {
		esBulk.AllowedMinPayout = json["allowedMinPayout"].(int)
	}
	if json["defaultMinPayout"] != nil {
		esBulk.DefaultMinPayout = json["defaultMinPayout"].(int)
	}
	if json["ipHint"] != nil {
		esBulk.IPHint = json["ipHint"].(string)
	}
	if json["ipWorkerName"] != nil {
		esBulk.IPWorkerName = json["ipWorkerName"].(string)
	}
	if json["minPayout"] != nil {
		esBulk.MinPayout = json["minPayout"].(int)
	}
	esBulk.CurrentHashrate = json["currentHashrate"].(float64)
	esBulk.CurrentLuck, _ = strconv.ParseFloat(json["currentLuck"].(string), 64)
	esBulk.Hashrate = json["hashrate"].(float64)
	esBulk.PageSize = json["pageSize"].(float64)
	esBulk.UpdatedAt = json["updatedAt"].(float64)
	esBulk.WorkersOffline = json["workersOffline"].(float64)
	esBulk.WorkersOnline = json["workersOnline"].(float64)
	esBulk.WorkersTotal = json["workersTotal"].(float64)

}

func ParseJson(bulkData data.MinerStat) string {
	bulk, _ := ioutil.ReadAll(bulkData.Json)
	var result map[string]interface{}
	err := json.Unmarshal(bulk, &result)
	utils.HandleHttpError(err)
	var EsBulk data.MinerInfo
	ExtractSimpleField(&EsBulk, result, bulkData.Wallet)
	ExtractWorkerInfo(result["workers"].(map[string]interface{}), bulkData.Wallet)
	EsBulkJson, err := json.Marshal(EsBulk)
	if err != nil {
		panic(err)
	}
	return string(EsBulkJson)
}

func PrepareDataForEs(bulkData data.MinerStat) string {
	//WalletInserted := SendAsJson(bulkData)
	return ParseJson(bulkData)

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
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
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
