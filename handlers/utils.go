package handlers

import (
	"2miner-monitoring/config"
	"2miner-monitoring/redis"
	"2miner-monitoring/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Health(c *gin.Context) {
	c.String(200, "We are Alive")
}

func RequestStorage(c *gin.Context) (map[string]interface{}, string) {
	wallet := c.Param("wallet")
	redisResult := redis.GetFromToRedis(1, wallet)
	if redisResult == "" {

		client := http.Client{
			Timeout: 180 * time.Second,
		}
		url := fmt.Sprintf("%s/accounts/%s", config.Cfg.TwoMinersURL, wallet)
		resp, err := client.Get(url)
		defer resp.Body.Close()
		if err != nil {
			log.Fatalf("ERROR: during fetch %s  %s", url, err)

		}
		bulk, _ := ioutil.ReadAll(resp.Body)
		redis.WriteToRedis(1, wallet, string(bulk), "short")
		var result map[string]interface{}
		err = json.Unmarshal(bulk, &result)
		utils.HandleHttpError(err)
		return result, wallet
	} else {
		var result map[string]interface{}
		err := json.Unmarshal([]byte(redisResult), &result)
		utils.HandleHttpError(err)
		return result, wallet
	}
}
