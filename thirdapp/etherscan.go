package thirdapp

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"encoding/json"
	"fmt"
	"github.com/nanmu42/etherscan-api"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
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

func GetLastTx(blockstring, wallet string) data.Tx {
	endBlock, _ := strconv.Atoi(blockstring)
	startBlock := endBlock - 276
	client := http.Client{
		Timeout: 180 * time.Second,
	}
	url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&page=1&offset=10000&sort=asc&apikey=%s", wallet, startBlock, endBlock, config.Cfg.APITokenEtherscan)
	resp, err := client.Get(url)
	tx := data.Tx{}
	if err != nil {
		log.Printf("%s error on GetLAstTx", err)
		return tx
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &tx)
	return tx
}
