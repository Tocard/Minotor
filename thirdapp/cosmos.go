package thirdapp

import (
	"2miner-monitoring/utils"
	"fmt"
	"io/ioutil"
)

func GetCosmosTokens() (int, []byte) {

	resp, err := utils.DoRequest("GET", "https://api-osmosis.imperator.co/tokens/v2/all", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetCosmosTokens", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
