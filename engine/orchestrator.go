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

func HarvestMiners(wallet string) {
	url := fmt.Sprintf("http://127.0.0.1:%d/miners", config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}

func HarvestFactory(endpoint string) {
	for key := range config.Cfg.Adress {
		go func(wallet string) {
			url := fmt.Sprintf("http://127.0.0.1:%d/harvest/%s/%s", config.Cfg.APIPort, endpoint, wallet)
			resp, err := http.Get(url)
			utils.HandleHttpError(err)
			defer resp.Body.Close()
			fmt.Println(resp)
		}(config.Cfg.Adress[key])

	}
}
