package thirdapp

import (
	"2miner-monitoring/config"
	"github.com/nanmu42/etherscan-api"
	"math/big"
	"strconv"
	"time"
)

type AccountBalance struct {
	Account string  `json:"account"`
	Balance *BigInt `json:"balance"`
}

type BigInt big.Int

func procEtherscanClient() *etherscan.Client {
	return etherscan.New(etherscan.Mainnet, config.Cfg.APITokenEtherscan)
}

func GetAccountBalance(walletId string) *big.Int {
	client := procEtherscanClient()
	balance, err := client.AccountBalance(walletId)
	if err != nil {
		panic(err)
	}
	return balance.Int()
}

func GetTokenBalance(contractAdress, walletId string) *big.Int {
	client := procEtherscanClient()
	tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")
	if err != nil {
		panic(err)
	}
	return tokenBalance.Int()
}

func GetMultiAccountBalance(walletIds []string) []etherscan.AccountBalance {
	client := procEtherscanClient()
	balance, err := client.MultiAccountBalance(walletIds...)
	if err != nil {
		panic(err)
	}
	return balance
}

func GetLastBlock() int {
	client := procEtherscanClient()
	block, err := client.BlockNumber(time.Now().Unix(), "before")
	if err != nil {
		panic(err)
	}
	return block
}

func GetLastTx(blockstring, wallet string) []etherscan.NormalTx {
	client := procEtherscanClient()
	endBlock, _ := strconv.Atoi(blockstring)
	startBlock := endBlock - 276
	tx, err := client.NormalTxByAddress(wallet, &startBlock, &endBlock, 1, 1, true)
	if err != nil {
		panic(err)
	}
	return tx
}
