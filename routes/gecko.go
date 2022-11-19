package routes

import (
	"minotor/config"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func GetCoinsPrice(c *gin.Context) {
	CoinsPrice, err := thirdapp.GetCurrencyValue(config.Cfg.CoinList...)
	if err != nil {
		c.String(500, "Unable to get coins price")
	}
	for key, _ := range config.Cfg.CoinList {
		var TmpGecko = data.GeckoCoin{}
		TmpGecko.Coin = config.Cfg.CoinList[key]
		ResponseCoin := (*CoinsPrice)[config.Cfg.CoinList[key]]
		TmpGecko.USD = ResponseCoin["usd"]
		TmpGecko.EUR = ResponseCoin["eur"]
		TmpGecko.Timestamp = time.Now().Format(time.RFC3339)
		tmpGeckoJson, _ := json.Marshal(TmpGecko)
		es.Bulk("defi-coins", string(tmpGeckoJson))

	}
	c.String(200, "OK")
}
