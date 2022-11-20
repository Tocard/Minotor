package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/config"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"minotor/utils"
	"net/http"
	"time"
)

func GetCosmosTokens(c *gin.Context) {
	var CosmosTokensByte [][]byte
	CosmosTokens := data.CosmosTokens{}
	Now := time.Now().Format(time.RFC3339)

	code, Body := thirdapp.GetCosmosTokens()
	err := json.Unmarshal(Body, &CosmosTokens)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetCosmosTokens", err))
		return
	}
	for _, CosmosToken := range CosmosTokens {
		CosmosToken.Timestamp = Now
		CosmosTokenJson, _ := json.Marshal(CosmosToken)
		CosmosTokensByte = append(CosmosTokensByte, CosmosTokenJson)
	}
	es.BulkData("minotor-cosmos-token", CosmosTokensByte)
	c.String(code, string(Body))
}

func GetCosmosMarket(c *gin.Context) {
	var GeckoAdvanceCoinsByte [][]byte
	GeckoAdvanceCoins := data.GeckoAdvanceCoins{}
	Now := time.Now().Format(time.RFC3339)

	Body := thirdapp.GetCoinsMarket()
	json.Unmarshal(Body, &GeckoAdvanceCoins)
	//	if err != nil { #TODO: understand why this error is triggered
	//	c.String(500, fmt.Sprintf("%s error on GetCosmosTokens", err.Error()))
	//	return
	//}
	for _, CosmosToken := range GeckoAdvanceCoins {
		CosmosToken.Timestamp = Now
		CosmosTokenJson, _ := json.Marshal(CosmosToken)
		GeckoAdvanceCoinsByte = append(GeckoAdvanceCoinsByte, CosmosTokenJson)
	}
	es.BulkData("minotor-cosmos-market", GeckoAdvanceCoinsByte)
	c.String(200, string(Body))
}

//delegation: staking/delegators/cosmos10nhs553a5g2ve7mh72je7f7zeeg9az7qdmglsj/delegations
// account URL de requête:
// unbounding URL de requête: staking/delegators/cosmos10nhs553a5g2ve7mh72je7f7zeeg9az7qdmglsj/unbonding_delegations
// price URL de requête: https://proxy.atomscan.com/prices

func WrapAllCosmosEndpoint(c *gin.Context) {
	wallet := c.Param("wallet")
	url := fmt.Sprintf("%s:%d/cosmos/GetBalance/%s", config.Cfg.APIAdress, config.Cfg.APIPort, wallet)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()
	url = fmt.Sprintf("%s:%d/cosmos/GetDelegation/%s", config.Cfg.APIAdress, config.Cfg.APIPort, wallet)
	resp, err = http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()
	c.String(200, "Call made")
}

func GetCosmosWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	var CosmosWallet [][]byte

	_, balance, coin := thirdapp.GetCosmosBalance(wallet)
	Balance := data.CosmosBalance{}
	err := json.Unmarshal(balance, &Balance)
	if err != nil {
		log.Println(err.Error())
	}
	for _, Res := range Balance.Result {
		Res.Timestamp = time.Now().Format(time.RFC3339)
		Res.Wallet = wallet
		Res.GovCoin = coin
		Res.Height = Balance.Height
		Res.Factor = data.GetFactor(coin)
		ResJson, _ := json.Marshal(Res)
		CosmosWallet = append(CosmosWallet, ResJson)
	}
	es.BulkData("cosmos-balance", CosmosWallet)
	c.String(200, string(balance))
}

func GetCosmosBounding(c *gin.Context) {
	wallet := c.Param("wallet")
	var CosmosBounding [][]byte

	_, balance, coin := thirdapp.GetCosmosBounding(wallet)
	Balance := data.CosmosDelegation{}
	err := json.Unmarshal(balance, &Balance)
	log.Println(string(balance))
	if err != nil {
		log.Println(err.Error())
	}
	for _, Res := range Balance.Result {
		Res.Timestamp = time.Now().Format(time.RFC3339)
		Res.Wallet = wallet
		Res.GovCoin = coin
		Res.Denom = Res.Balance.Denom
		Res.Height = Balance.Height
		Res.Factor = data.GetFactor(coin)
		ResJson, _ := json.Marshal(Res)
		CosmosBounding = append(CosmosBounding, ResJson)
	}
	es.BulkData("cosmos-delegation", CosmosBounding)
	c.String(200, string(balance))
}
