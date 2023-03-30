package data

import (
	"fmt"
	"minotor/config"
	"strings"
	"time"
)

func GetTokenUrl(wallet string) (string, string) {
	tokenType := ""
	for _, coin := range config.Cfg.TokenWatcher {
		if strings.Contains(wallet, coin) {
			for _, url := range config.Cfg.UrlTokenWatcher {
				if coin != "cosmos" {
					tokenType = fmt.Sprintf("%s-lcd", coin)
				} else {
					tokenType = "cosmos"
				}
				if strings.Contains(url, tokenType) {
					return url, coin
				}
			}
		}
	}
	return "Not yet supported", ""
}

func GetFactor(coin string) float64 {
	switch coin {
	case "cosmos":
		return 0.000001
	case "osmo":
		return 0.000001
	case "evmos":
		return 0.000000000000000001
	}
	return 0
}

type CosmosDelegation struct {
	Height string `json:"height"`
	Result []struct {
		Height     string  `json:"height"`
		Timestamp  string  `json:"@timestamp"`
		GovCoin    string  `json:"gov_coin"`
		Wallet     string  `json:"wallet"`
		Factor     float64 `json:"factor"`
		Denom      string  `json:"denom"`
		Delegation struct {
			DelegatorAddress string  `json:"delegator_address"`
			ValidatorAddress string  `json:"validator_address"`
			Shares           float64 `json:"shares,string"`
		} `json:"delegation"`
		Balance struct {
			Denom  string  `json:"denom"`
			Amount float64 `json:"amount,string"`
		} `json:"balance"`
	} `json:"result"`
}

type CosmosBalance struct {
	Height string `json:"height"`
	Result []struct {
		Timestamp string  `json:"@timestamp"`
		Wallet    string  `json:"wallet"`
		Height    string  `json:"height"`
		GovCoin   string  `json:"gov_coin"`
		Factor    float64 `json:"factor"`
		Denom     string  `json:"denom"`
		Amount    float64 `json:"amount,string"`
	} `json:"result"`
}

type CosmosUnDelegation struct {
	Height string `json:"height"`
	Result []struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Entries          []struct {
			CreationHeight string    `json:"creation_height"`
			CompletionTime time.Time `json:"completion_time"`
			InitialBalance float64   `json:"initial_balance,string"`
			Balance        float64   `json:"balance,string"`
			Coin           string    `json:"coin"`
			Wallet         string    `json:"wallet"`
			Timestamp      string    `json:"@timestamp"`
			Height         string    `json:"height"`
			GovCoin        string    `json:"gov_coin"`
			Factor         float64   `json:"factor"`
		} `json:"entries"`
	} `json:"result"`
}
