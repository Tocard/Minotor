package engine

import (
	"fmt"
	"minotor/config"
	"minotor/utils"
	"net/http"
)

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

func HealthCheck() {
	url := fmt.Sprintf("%s:%d/health", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetAllNodesStatus() {
	url := fmt.Sprintf("%s:%d/streamr/all_operator", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolBlockWins() {
	url := fmt.Sprintf("%s:%d/chia/pool/blocks_win", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolFarmers() {
	url := fmt.Sprintf("%s:%d/chia/pool/farmer", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolFarmerNetspace() {
	url := fmt.Sprintf("%s:%d/chia/pool/farmer_netspace", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolPartial() {
	url := fmt.Sprintf("%s:%d/chia/pool/partials", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolFarmerPayment() {
	url := fmt.Sprintf("%s:%d/chia/pool/payments", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolFarmerUptime() {
	url := fmt.Sprintf("%s:%d/chia/pool/uptime", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func EngineChiaPoolPoolNetspace() {
	url := fmt.Sprintf("%s:%d/chia/pool/pool_netspace", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
