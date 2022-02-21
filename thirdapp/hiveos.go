package thirdapp

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/redis"
	"2miner-monitoring/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func HiveosGetAuthToken() {
	url := fmt.Sprintf("%s/auth/login", config.Cfg.HiveosUrl)
	payload := data.HiveosAuth{Login: "test", Password: "test"}
	resp, err := utils.DoRequest("POST", url, payload)
	if err != nil {
		log.Printf("%s error on HiveosGetAuthToken", err)
	}
	token := data.HiveosToken{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	redis.WriteToRedis(0, "token", token.AccessToken, "long")
}

func HiveosRefreshToken() {
	Redistoken := redis.GetFromToRedis(0, "token")
	url := fmt.Sprintf("%s/auth/refresh", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("POST", url, nil, Redistoken)
	if err != nil {
		log.Printf("%s error on HiveosRefreshToken", err)
	}
	token := data.HiveosToken{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	redis.WriteToRedis(0, "token", token.AccessToken, "long")

}
