package thirdapp

import (
	"encoding/json"
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
