package routes

import (
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetCosmosTokens(c *gin.Context) {
	var CosmosTokensByte [][]byte
	CosmosTokens := data.CosmosTokens{}
	Now := time.Now().Format(time.RFC3339)

	code, Body := thirdapp.GetCosmosTokens()
	err := json.Unmarshal(Body, &CosmosTokens)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetCosmosTokens", err))
		return
	}
	for _, CosmosToken := range CosmosTokens {
		CosmosToken.Timestamp = Now
		CosmosTokenJson, _ := json.Marshal(CosmosToken)
		CosmosTokensByte = append(CosmosTokensByte, CosmosTokenJson)
	}
	es.BulkData("minotor-cosmos-token", CosmosTokensByte)
	c.String(code, string(Body))
}

func GetCosmosMarket(c *gin.Context) {
	var GeckoAdvanceCoinsByte [][]byte
	GeckoAdvanceCoins := data.GeckoAdvanceCoins{}
	Now := time.Now().Format(time.RFC3339)

	Body := thirdapp.GetCoinsMarket()
	json.Unmarshal(Body, &GeckoAdvanceCoins)
	//	if err != nil { #TODO: understand why this error is triggered
	//	c.String(500, fmt.Sprintf("%s error on GetCosmosTokens", err.Error()))
	//	return
	//}
	for _, CosmosToken := range GeckoAdvanceCoins {
		CosmosToken.Timestamp = Now
		CosmosTokenJson, _ := json.Marshal(CosmosToken)
		GeckoAdvanceCoinsByte = append(GeckoAdvanceCoinsByte, CosmosTokenJson)
	}
	es.BulkData("minotor-cosmos-market", GeckoAdvanceCoinsByte)
	c.String(200, string(Body))
}
