package engine

import (
	"2miner-monitoring/data"
	"log"
	"net/http"
	"time"
)

func AsyncGet(urls map[string]string) []data.MinerStat {
	ch := make(chan data.MinerStat)
	defer close(ch)
	var responses []data.MinerStat

	for wallet, url := range urls {

		go func(wallet string, u string) {
			localTime := time.Now()
			log.Printf("AsyncGet start %s at %s", u, localTime)
			client := http.Client{
				Timeout: 180 * time.Second,
			}
			resp, err := client.Get(u)
			if err != nil {
				log.Fatalf("ERROR: during async  %s", err)

			}
			//defer resp.Body.Close()
			ch <- data.MinerStat{
				Json:   resp.Body,
				Wallet: wallet,
			}
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
