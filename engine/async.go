package engine

import (
	"2miner-monitoring/data"
	"log"
	"net/http"
	"time"
)

func AsyncGet(urls map[string]string) []data.MinerStat {
	ch := make(chan data.MinerStat)
	var responses []data.MinerStat

	for wallet, url := range urls {

		go func(wallet string, u string) {
			localTime := time.Now()
			log.Printf("AsyncGet start %s at %s", u, localTime)
			resp, _ := http.Get(u)
			//MinerInfo := &data.MinerInfo{}
			//d := json.NewDecoder(resp.Body)
			//_ = d.Decode(MinerInfo)
			//MinerInfo.Wallet = wallet
			//MinerInfo.Timestamp = time.Now().Format(time.RFC3339)

			ch <- data.MinerStat{
				Json:   resp.Body,
				Wallet: wallet,
			}
			//ch <- MinerInfo
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
