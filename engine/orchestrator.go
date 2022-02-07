package engine

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"fmt"
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

func GetTargetUrl(MinerList []data.Miner) map[string]string {
	Urls := map[string]string{}
	for key, _ := range MinerList {
		url := fmt.Sprintf("%s/accounts/%s", config.Cfg.TwoMinersURL, MinerList[key].Adress)
		Urls[MinerList[key].Adress] = url
	}
	return Urls
}
