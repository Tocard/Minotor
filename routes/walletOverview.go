package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"net/http"
	"time"
)

func PostWalletsOverview(c *gin.Context) {
	WalletOverview := data.WalletOverviewPure{}
	var WalletOverviewJson [][]byte

	if err := c.BindJSON(&WalletOverview); err != nil {
		return
	}
	timer := time.Now().Format(time.RFC3339)
	GeckoAdvanceCoins := data.GeckoAdvanceCoins{}
	GeckoAdvanceCoins2 := data.GeckoAdvanceCoins{}
	Body, Body2 := thirdapp.Get2CoinsMarket()
	json.Unmarshal(Body, &GeckoAdvanceCoins.GeckoAdvanceCoins)
	json.Unmarshal(Body2, &GeckoAdvanceCoins2.GeckoAdvanceCoins)
	Coin := make(map[string]data.GeckoAdvanceCoin)
	for _, Token := range GeckoAdvanceCoins.GeckoAdvanceCoins {
		Coin[Token.Symbol] = Token
	}
	for _, Token := range GeckoAdvanceCoins2.GeckoAdvanceCoins {
		Coin[Token.Symbol] = Token
	}
	for _, Raw := range WalletOverview.WalletOverview {
		Raw.Timestamp = timer
		Raw.Pseudo = WalletOverview.Pseudo
		if Raw.Price != 0.0 {
			Coin[Raw.Token] = data.GeckoAdvanceCoin{CurrentPrice: Raw.Price}
		}
		Raw.Value = Coin[Raw.Token].CurrentPrice * Raw.Amount
		rawJson, _ := json.Marshal(Raw)
		WalletOverviewJson = append(WalletOverviewJson, rawJson)
	}
	es.BulkData("minotor-wallet-overview", WalletOverviewJson)
	c.IndentedJSON(http.StatusCreated, WalletOverviewJson)
}
