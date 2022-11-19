package thirdapp

import (
	"minotor/utils"
	"fmt"
	"io/ioutil"
	"log"
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

func GetNodesOverview() (int, []byte) {
	resp, err := utils.DoRequest("GET", "https://api.runonflux.io/daemon/listzelnodes", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetBlocStats", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}

func GetZelNodeStatus(ip string) (int, []byte) {
	url := fmt.Sprintf("http://%s:%d/daemon/getzelnodestatus", ip, 16126)
	log.Println(url)
	resp, err := utils.DoRequest("GET", url, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetZelNodeStatus", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
