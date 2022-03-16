package thirdapp

import (
	"2miner-monitoring/utils"
	"fmt"
	"io/ioutil"
)

func GetNodesStats() (int, []byte) {

	resp, err := utils.DoRequest("GET", "https://api.runonflux.io/daemon/getzelnodecount", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetNodesStats", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}

func GetBlocStats() (int, []byte) {
	resp, err := utils.DoRequest("GET", "https://explorer.runonflux.io/api/statistics/total", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetBlocStats", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
