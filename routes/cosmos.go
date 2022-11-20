package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/config"
	"minotor/data"
	"minotor/db"
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

func WrapAllCosmosEndpoint(c *gin.Context) {

	url := fmt.Sprintf("%s:%d/cosmos/GetBalance", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	url = fmt.Sprintf("%s:%d/cosmos/GetDelegation", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err = http.Get(url)
	utils.HandleHttpError(err)
	url = fmt.Sprintf("%s:%d/cosmos/GetUnDelegation", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err = http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()
	c.String(200, "Call made")
}

func GetCosmosWallet(c *gin.Context) {
	var CosmosWallet [][]byte

	Wallets, dbErr := db.GetAllWallets()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, Wallet := range Wallets {
		_, balance, coin := thirdapp.GetCosmosBalance(Wallet.Wallet)
		Balance := data.CosmosBalance{}
		err := json.Unmarshal(balance, &Balance)
		if err != nil {
			log.Println(err.Error())
		}
		for _, Res := range Balance.Result {
			Res.Timestamp = time.Now().Format(time.RFC3339)
			Res.Wallet = Wallet.Wallet
			Res.GovCoin = coin
			Res.Height = Balance.Height
			Res.Factor = data.GetFactor(coin)
			ResJson, _ := json.Marshal(Res)
			CosmosWallet = append(CosmosWallet, ResJson)
		}
	}
	es.BulkData("cosmos-balance", CosmosWallet)
	c.String(200, "Harvest all wallet adresse")
}

func GetCosmosBounding(c *gin.Context) {
	var CosmosBounding [][]byte

	Wallets, dbErr := db.GetAllWallets()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, Wallet := range Wallets {
		_, balance, coin := thirdapp.GetCosmosBounding(Wallet.Wallet)
		Balance := data.CosmosDelegation{}
		err := json.Unmarshal(balance, &Balance)
		if err != nil {
			log.Println(err.Error())
		}
		for _, Res := range Balance.Result {
			Res.Timestamp = time.Now().Format(time.RFC3339)
			Res.Wallet = Wallet.Wallet
			Res.GovCoin = coin
			Res.Denom = Res.Balance.Denom
			Res.Height = Balance.Height
			Res.Factor = data.GetFactor(coin)
			ResJson, _ := json.Marshal(Res)
			CosmosBounding = append(CosmosBounding, ResJson)
		}
	}
	es.BulkData("cosmos-delegation", CosmosBounding)
	c.String(200, "Harvest all bound")
}

func GetCosmosUnBounding(c *gin.Context) {
	var CosmosUnBounding [][]byte

	Wallets, dbErr := db.GetAllWallets()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, Wallet := range Wallets {
		_, balance, coin := thirdapp.GetCosmosUnBounding(Wallet.Wallet)
		Balance := data.CosmosUnDelegation{}
		err := json.Unmarshal(balance, &Balance)
		log.Println(string(balance))
		if err != nil {
			log.Println(err.Error())
		}
		for _, Res := range Balance.Result {
			for _, Entry := range Res.Entries {
				Entry.Timestamp = time.Now().Format(time.RFC3339)
				Entry.Wallet = Wallet.Wallet
				Entry.GovCoin = coin
				Entry.Height = Balance.Height
				Entry.Factor = data.GetFactor(coin)
				EntryJson, _ := json.Marshal(Entry)
				CosmosUnBounding = append(CosmosUnBounding, EntryJson)
			}
		}
	}
	es.BulkData("cosmos-undelegation", CosmosUnBounding)
	c.String(200, "Harvest all bound")
}

func RegisterWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	Wallet := db.NewWallet(wallet)
	err := Wallet.Save()
	if err != nil {
		resp := fmt.Sprintf("something went wrong while registered %s", wallet)
		c.String(503, resp)

	} else {
		resp := fmt.Sprintf("wallet %s fully registered", wallet)
		c.String(201, resp)
	}
}

func UnRegisterWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	Wallet, err := db.GetWalletByAdresses(wallet)
	if err != nil {
		resp := fmt.Sprintf("Wallet %s is not registered", wallet)
		c.String(404, resp)
	}
	err = Wallet.Delete()
	if err != nil {
		resp := fmt.Sprintf("Unable to delete wallet %s, contact admin", wallet)
		c.String(503, resp)
	} else {
		resp := fmt.Sprintf("wallet %s succefully removed", wallet)
		c.String(200, resp)
	}
}
