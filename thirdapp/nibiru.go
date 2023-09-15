package thirdapp

import (
	"fmt"
	"io/ioutil"
	"minotor/utils"
)

func GetNibiruValidators() (int, []byte) {

	resp, err := utils.DoRequest("GET", "http://192.168.1.46:1317/cosmos/staking/v1beta1/validators", nil)
	if err != nil {
		return resp.StatusCode, []byte(fmt.Sprintf("%s error on GetNodesStats", err))
	}
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
