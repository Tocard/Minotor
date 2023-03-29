package engine

import (
	"fmt"
	"io"
	"log"
	"minotor/config"
	"minotor/utils"
	"net/http"
)

func HarvestCoinPrice() {
	url := fmt.Sprintf("%s:%d/coins/price", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func HarvestComsosWallet() {

	url := fmt.Sprintf("%s:%d/cosmos/wrapper/", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)
}

func FluxNodeRentability() {
	url := fmt.Sprintf("%s:%d/flux/calcul_nodes_rentability", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func FluxNodesOverview() {
	url := fmt.Sprintf("%s:%d/flux/flux_nodes_overview", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func GetCosmosTokens() {
	url := fmt.Sprintf("%s:%d/cosmos/get_tokens", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func GetCosmosMarket() {
	url := fmt.Sprintf("%s:%d/cosmos/get_market", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func HealthCheck() {
	url := fmt.Sprintf("%s:%d/health", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func GetOsmosisPool() {
	url := fmt.Sprintf("%s:%d/osmosis/getPools", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}

func GetStreamR() {
	url := fmt.Sprintf("%s:%d/streamr/status/%s", config.Cfg.APIAdress, config.Cfg.APIPort, config.Cfg.StreamRAddr)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(resp.Body)

	fmt.Println(resp)
}
