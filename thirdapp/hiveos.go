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

func HiveosRefreshToken() (int, string) {
	Redistoken := redis.GetFromToRedis(0, "token")
	url := fmt.Sprintf("%s/auth/refresh", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("POST", url, nil, Redistoken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosRefreshToken", err)
	}
	token := data.HiveosToken{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	redis.WriteToRedis(0, "token", token.AccessToken, "long")
	return 200, "Auth Hiveos OK"
}

func HiveosGetAuthToken() (int, string) {
	url := fmt.Sprintf("%s/auth/login", config.Cfg.HiveosUrl)
	payload := data.HiveosAuth{Login: config.Cfg.MinotorHiveOsUser, Password: config.Cfg.MinotorHiveOsPass}
	log.Printf("payload = %s", payload)
	resp, err := utils.DoRequest("POST", url, payload)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetAuthToken", err)
	}
	token := data.HiveosToken{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	log.Printf("body = %s", body)
	log.Printf("token = %s", token)
	redis.WriteToRedis(0, "token", token.AccessToken, "long")
	log.Printf("TOKEN = %s", redis.GetFromToRedis(0, "token"))
	return 200, "Renew Hiveos Auth OK"
}

func HiveosGetFarms() (int, string) {
	url := fmt.Sprintf("%s/farms", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("GET", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetFarms", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func HiveosGetWorkers(farmrId int) (int, string) {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/workers", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetWorkers", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func HiveosGetWorker(farmrId, workerId int) (int, string) {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/workers/%d", config.Cfg.HiveosUrl, farmrId, workerId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetWorker", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func HiveosGetOc(farmrId int) (int, string) {
	Redistoken := redis.GetFromToRedis(0, "token")
	if Redistoken == "" {
		HiveosRefreshToken()
		Redistoken = redis.GetFromToRedis(0, "token")
	}
	url := fmt.Sprintf("%s/farms/%d/oc", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, Redistoken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetOc", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}
