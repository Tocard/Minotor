package engine

import (
	"2miner-monitoring/data"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func AsyncGet(urls map[string]string) []*data.MinerInfo {
	ch := make(chan *data.MinerInfo)
	var responses []*data.MinerInfo

	for wallet, url := range urls {

		go func(wallet string, u string) {
			localTime := time.Now()
			log.Printf("AsyncGet start %s at %s", u, localTime)
			resp, _ := http.Get(u)
			MinerInfo := &data.MinerInfo{}
			d := json.NewDecoder(resp.Body)
			//d := msgpack.NewDecoder(resp.Body)
			_ = d.Decode(MinerInfo)
			//d.Decode(MinerInfo)
			MinerInfo.Wallet = wallet
			MinerInfo.Timestamp = time.Now().Format(time.RFC3339)
			ch <- MinerInfo
			duration := time.Since(localTime)
			log.Printf("AsyncGet stop at %s. Duration :%f secondes", u, duration.Seconds())
		}(wallet, url)
	}

loop:
	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == len(urls) {
				break loop
			}
		}
	}
	return responses
}
