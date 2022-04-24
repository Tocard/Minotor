package routes

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/thirdapp"
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
