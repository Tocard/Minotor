package thirdapp

import (
	"fmt"
	"io/ioutil"
	"minotor/utils"
)

func GetStreamStatsFromBruberScan(addr string) (int, []byte) {
	url := fmt.Sprintf("https://brubeckscan.app/api/nodes/%s", addr)
	resp, err := utils.DoRequest("GET", url, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetStreamStatsFromBruberScan", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
