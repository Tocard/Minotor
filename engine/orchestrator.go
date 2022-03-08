package engine

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/utils"
	"fmt"
	"log"
	"net/http"
)

func Minertarget() []data.Miner {
	var MinerList []data.Miner
	if config.Cfg.MinerListing == "ALL" {
		return MinerList
	} else {
		for key := range config.Wtw.Adress {
			tmpMiner := data.Miner{}
			tmpMiner.Adress = config.Wtw.Adress[key]
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
	for key := range config.Wtw.Adress {
		go func(wallet string) {
			url := fmt.Sprintf("%s:%d/harvest/%s/%s", config.Cfg.APIAdress, config.Cfg.APIFrontPort, endpoint, wallet)
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error with wallet %s, unsubscribe\n", wallet)
				url := fmt.Sprintf("%s:%d/unsubscribe/%s", config.Cfg.APIAdress, config.Cfg.APIFrontPort, wallet)
				_, err := http.Get(url)
				if err != nil {
					log.Printf("Unable to unsubscribe %s, err : %s", wallet, err)
				}
			}
			defer resp.Body.Close()
			fmt.Println(resp)
		}(config.Wtw.Adress[key])

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

func GetLastEthBlock() {
	url := fmt.Sprintf("%s:%d/ETH/lastblock", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetLastEthTx() {
	url := fmt.Sprintf("%s:%d/transactions", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetHiveosFarm() {
	url := fmt.Sprintf("%s:%d/hiveos/farms", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func GetHiveosWorker() {
	url := fmt.Sprintf("%s:%d/hiveos/workers", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func ScrapHashrateNo() {
	url := fmt.Sprintf("%s:%d/hashrateNo", config.Cfg.APIAdress, config.Cfg.APIFrontPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
