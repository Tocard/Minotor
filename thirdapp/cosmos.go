package thirdapp

import (
	"fmt"
	"io/ioutil"
	"minotor/data"
	"minotor/utils"
)

func GetCosmosTokens() (int, []byte) {

	resp, err := utils.DoRequest("GET", "https://api-osmosis.imperator.co/tokens/v2/all", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetCosmosTokens", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}

func GetCosmosBalance(wallet string) (int, []byte, string) {

	baseUrl, coin := data.GetTokenUrl(wallet)
	if baseUrl == "Not yet supported" {
		return 501, []byte(baseUrl), ""
	}
	url := fmt.Sprintf("%s/bank/balances/%s", baseUrl, wallet)
	resp, err := utils.DoRequest("GET", url, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetCosmosBalance", err)), ""
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body, coin
}

func GetCosmosBounding(wallet string) (int, []byte, string) {
	baseUrl, coin := data.GetTokenUrl(wallet)
	if baseUrl == "Not yet supported" {
		return 501, []byte(baseUrl), ""
	}
	url := fmt.Sprintf("%s/staking/delegators/%s/delegations", baseUrl, wallet)
	resp, err := utils.DoRequest("GET", url, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetCosmosBalance", err)), ""
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body, coin
}

func GetCosmosUnBounding(wallet string) (int, []byte, string) {
	baseUrl, coin := data.GetTokenUrl(wallet)
	if baseUrl == "Not yet supported" {
		return 501, []byte(baseUrl), ""
	}
	url := fmt.Sprintf("%s/staking/delegators/%s/unbonding_delegations", baseUrl, wallet)
	resp, err := utils.DoRequest("GET", url, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetCosmosBalance", err)), ""
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body, coin
}
