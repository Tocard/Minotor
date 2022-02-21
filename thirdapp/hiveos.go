package thirdapp

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/redis"
	"2miner-monitoring/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func HiveosRefreshToken() (int, string) {
	url := fmt.Sprintf("%s/auth/refresh", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("POST", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosRefreshToken", err)
	}
	token := data.HiveosToken{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	redis.WriteToRedis(0, "token", token.AccessToken, "long")
	return 200, "Auth Hiveos OK"
}

func HiveosGetFarms() (int, []byte) {
	url := fmt.Sprintf("%s/farms", config.Cfg.HiveosUrl)
	resp, err := utils.DoRequest("GET", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on HiveosGetFarms", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body
}

func HiveosGetWorkers(farmrId int) (int, string) {
	url := fmt.Sprintf("%s/farms/%d/workers", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetWorkers", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func HiveosGetWorker(farmrId, workerId int) (int, string) {
	url := fmt.Sprintf("%s/farms/%d/workers/%d", config.Cfg.HiveosUrl, farmrId, workerId)
	resp, err := utils.DoRequest("GET", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetWorker", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func HiveosGetOc(farmrId int) (int, string) {
	url := fmt.Sprintf("%s/farms/%d/oc", config.Cfg.HiveosUrl, farmrId)
	resp, err := utils.DoRequest("GET", url, nil, config.Cfg.MinotorHiveosToken)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("%s error on HiveosGetOc", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}
