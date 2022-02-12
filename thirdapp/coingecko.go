package thirdapp

import (
	"fmt"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
	"net/http"
	"time"
)

func procGeckoClient() *gecko.Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return gecko.NewClient(httpClient)
}

func GetCoinList() {
	cg := procGeckoClient()
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}

func GetCurrencyValue(coins ...string) (*map[string]map[string]float32, error) {
	cg := procGeckoClient()
	vc := []string{"usd", "eur"}
	sp, err := cg.SimplePrice(coins, vc)
	if err != nil {
		return nil, err
	}
	return sp, nil
}
