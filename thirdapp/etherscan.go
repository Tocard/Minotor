package thirdapp

import (
	"2miner-monitoring/config"
	"github.com/nanmu42/etherscan-api"
	"math/big"
)

type AccountBalance struct {
	Account string  `json:"account"`
	Balance *BigInt `json:"balance"`
}

type BigInt big.Int

var client = etherscan.New(etherscan.Mainnet, config.Cfg.APITokenEtherscan)

func getEtherBalance(walletId string) *big.Int {
	// check account balance
	balance, err := client.AccountBalance(walletId)
	if err != nil {
		panic(err)
	}
	return balance.Int()
}
