package thirdapp

import (
	"encoding/json"
	"fmt"
	gecko "github.com/superoo7/go-gecko/v3"
	geckoTypes "github.com/superoo7/go-gecko/v3/types"
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

func GetCoinsMarket() []byte {
	cg := procGeckoClient()
	vsCurrency := "eur"
	var ids []string
	perPage := 250
	page := 1
	pcp := geckoTypes.PriceChangePercentageObject
	priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
	order := geckoTypes.OrderTypeObject.MarketCapDesc
	market, err := cg.CoinsMarket(vsCurrency, ids, order, perPage, page, false, priceChangePercentage)
	if err != nil {
		log.Fatal(err)
	}
	u, err := json.Marshal(*market)
	return u
}

func Get2CoinsMarket() ([]byte, []byte) {
	cg := procGeckoClient()
	vsCurrency := "eur"
	var ids []string
	pcp := geckoTypes.PriceChangePercentageObject
	priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
	order := geckoTypes.OrderTypeObject.MarketCapDesc
	market, err := cg.CoinsMarket(vsCurrency, ids, order, 250, 1, false, priceChangePercentage)
	if err != nil {
		log.Fatal(err)
	}
	u, err := json.Marshal(*market)
	market2, err := cg.CoinsMarket(vsCurrency, ids, order, 250, 2, false, priceChangePercentage)
	if err != nil {
		log.Fatal(err)
	}
	u2, err := json.Marshal(*market2)
	return u, u2
}

func GetCoin() {
	cg := procGeckoClient()
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}
