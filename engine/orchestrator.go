package engine

import (
	"fmt"
	"log"
	"minotor/config"
	"minotor/db"
	"minotor/utils"
	"net/http"
)

func HarvestCoinPrice() {
	url := fmt.Sprintf("%s:%d/coins/price", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HarvestComsosWallet() {
	Wallets, dbErr := db.GetAllWallets()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, Wallet := range Wallets {
		url := fmt.Sprintf("%s:%d/cosmos/wrapper/%s", config.Cfg.APIAdress, config.Cfg.APIPort, Wallet.Wallet)
		resp, err := http.Get(url)
		utils.HandleHttpError(err)
		defer resp.Body.Close()
	}
	log.Println("Harvested", Wallets)
}

func FluxNodeRentability() {
	url := fmt.Sprintf("%s:%d/flux/calcul_nodes_rentability", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func FluxNodesOverview() {
	url := fmt.Sprintf("%s:%d/flux/flux_nodes_overview", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetCosmosTokens() {
	url := fmt.Sprintf("%s:%d/cosmos/get_tokens", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetCosmosMarket() {
	url := fmt.Sprintf("%s:%d/cosmos/get_market", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
