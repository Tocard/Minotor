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
	url := fmt.Sprintf("%s:%d/cosmos/wrapper/%s", config.Cfg.APIAdress, config.Cfg.APIPort, "evmos18477p09j434edhcrvczraneu7rlrz6fsf7tgd9")
	resp, err := http.Get(url)
	utils.HandleHttpError(err)

	fmt.Println(resp)
	url = fmt.Sprintf("%s:%d/cosmos/wrapper/%s", config.Cfg.APIAdress, config.Cfg.APIPort, "cosmos10nhs553a5g2ve7mh72je7f7zeeg9az7qdmglsj")
	resp, err = http.Get(url)
	utils.HandleHttpError(err)

	fmt.Println(resp)
	url = fmt.Sprintf("%s:%d/cosmos/wrapper/%s", config.Cfg.APIAdress, config.Cfg.APIPort, "osmo10nhs553a5g2ve7mh72je7f7zeeg9az7q9qm0xq")
	resp, err = http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
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
