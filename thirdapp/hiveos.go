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

func HiveosGetFarms() string {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		log.Printf("%s error on HiveosGetFarms", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HiveosGetWorkers(farmrId int) string {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/workers", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		log.Printf("%s error on HiveosGetWorkers", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HiveosGetWorker(farmrId, workerId int) string {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/workers/%d", config.Cfg.HiveosUrl, farmrId, workerId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		log.Printf("%s error on HiveosGetWorker", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HiveosGetOc(farmrId int) string {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/oc", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		log.Printf("%s error on HiveosGetOc", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
