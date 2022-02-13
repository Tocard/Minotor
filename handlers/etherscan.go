package handlers

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func GetWalletsBalance(c *gin.Context) {
	balances := thirdapp.GetMultiAccountBalance(config.Wtw.Adress)
	for key, _ := range balances {
		tmpBalance := data.Balance{}
		tmpBalance.Balance = balances[key].Balance
		tmpBalance.Wallet = balances[key].Account
		tmpBalance.Timestamp = time.Now().Format(time.RFC3339)
		tmpBalanceJson, _ := json.Marshal(tmpBalance)
		es.Bulk("2miners-balance", string(tmpBalanceJson))

	}
	c.String(200, "OK")

}
