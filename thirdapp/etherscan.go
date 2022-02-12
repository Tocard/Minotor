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

func GetAccountBalance(walletId string) *big.Int {
	// check account balance
	balance, err := client.AccountBalance(walletId)
	if err != nil {
		panic(err)
	}
	return balance.Int()
}

func GetTokenBalance(contractAdress, walletId string) *big.Int {
	// check account balance
	tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")
	if err != nil {
		panic(err)
	}
	return tokenBalance.Int()
}

func GetMultiAccountBalance(walletIds []string) []etherscan.AccountBalance {
	// check account balance
	balance, err := client.MultiAccountBalance(walletIds...)
	if err != nil {
		panic(err)
	}
	return balance
}
