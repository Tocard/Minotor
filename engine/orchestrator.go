package engine

import (
	"fmt"
	"minotor/config"
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

	url := fmt.Sprintf("%s:%d/cosmos/wrapper/", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()
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

func HealthCheck() {
	url := fmt.Sprintf("%s:%d/health", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetOsmosisPool() {
	url := fmt.Sprintf("%s:%d/osmosis/getPools", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetHoppersBalance() {
	url := fmt.Sprintf("%s:%d/cosmos/hoppers", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
