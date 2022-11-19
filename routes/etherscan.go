package routes

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
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
		es.Bulk("defi-balance", string(tmpBalanceJson))

	}
	c.String(200, "OK")

}

func GetLastBlock(c *gin.Context) {
	block := redis.GetFromToRedis(0, "LastBlock")
	if block == "" {
		blockInt := thirdapp.GetLastBlock()
		redis.WriteToRedis(0, "LastBlock", strconv.Itoa(blockInt), "mid")
	}
	c.String(200, "OK")

}

func GetLastTransaction(c *gin.Context) {
	block := redis.GetFromToRedis(0, "LastBlock")
	if block == "" {
		blockInt := thirdapp.GetLastBlock()
		block = strconv.Itoa(blockInt)
		redis.WriteToRedis(0, "LastBlock", block, "mid")

	}
	for key, _ := range config.Wtw.Adress {
		tx := thirdapp.GetLastTx(block, config.Wtw.Adress[key])
		for resultKey, _ := range tx.Result {
			ts, _ := strconv.Atoi(tx.Result[resultKey].TimeStamp)
			tx.Result[resultKey].Timestamp = time.Unix(int64(ts), 0).Format(time.RFC3339)
			txRawJson, _ := json.Marshal(tx.Result[resultKey])
			es.Bulk("defi-tx", string(txRawJson))
		}
	}
	c.String(200, "OK")

}
