package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"time"
)

func GetHopersBalance(c *gin.Context) {
	HopersBalance := data.Hopers{}
	code, Body := thirdapp.GetHoppersBalance()
	HopersBalance.Timestamp = time.Now().Format(time.RFC3339)
	err := json.Unmarshal(Body, &HopersBalance)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetCosmosTokens", err))
		return
	}
	HopersBalanceJson, _ := json.Marshal(HopersBalance)
	es.Bulk("minotor-cosmos-hopers", string(HopersBalanceJson))
	c.String(code, string(Body))
}
