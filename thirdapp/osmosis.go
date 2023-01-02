package thirdapp

import (
	"fmt"
	"io/ioutil"
	"minotor/utils"
)

func GetAllPool() (int, []byte) {
	resp, err := utils.DoRequest("GET", "https://api-osmosis.imperator.co/pools/v2beta3/all", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetAllPool", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
