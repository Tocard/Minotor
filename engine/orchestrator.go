package engine

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/utils"
	"fmt"
	"net/http"
)

func Minertarget() []data.Miner {
	var MinerList []data.Miner
	if config.Cfg.MinerListing == "ALL" {
		return MinerList
	} else {
		for key := range config.Cfg.Adress {
			tmpMiner := data.Miner{}
			tmpMiner.Adress = config.Cfg.Adress[key]
			MinerList = append(MinerList, tmpMiner)
		}
		return MinerList
	}
}

func HarvestMiners() {
	url := fmt.Sprintf("%s:%d/miners", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HarvestFactory(endpoint string) {
	for key := range config.Cfg.Adress {
		go func(wallet string) {
			url := fmt.Sprintf("%s:%d/harvest/%s/%s", config.Cfg.APIAdress, config.Cfg.APIFrontPort, endpoint, wallet)
			resp, err := http.Get(url)
			utils.HandleHttpError(err)
			defer resp.Body.Close()
			fmt.Println(resp)
		}(config.Cfg.Adress[key])

	}
}

func HarvestBalance() {
	url := fmt.Sprintf("%s:%d/balances", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HarvestPoolStat() {
	url := fmt.Sprintf("%s:%d/stats", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HarvestCoinPrice() {
	url := fmt.Sprintf("%s:%d/coins/price", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
