package handlers

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ScrapHashrateNo(c *gin.Context) {
	thirdapp.RunCrawler()
	JsonCardResult, _ := json.Marshal(data.CardsResult)
	es.Bulk("2miners-hashrateNo", string(JsonCardResult))
	c.String(200, "OK")

}
