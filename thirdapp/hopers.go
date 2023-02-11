package thirdapp

import (
	"fmt"
	"io"
	"minotor/config"
	"minotor/utils"
)

func GetHoppersBalance() (int, []byte) {
	resp, err := utils.DoRequest("GET", config.Cfg.HopersUrl, nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetHopersBalance", err))
	}
	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, body
}
