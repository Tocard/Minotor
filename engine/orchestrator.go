package engine

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"fmt"
	"log"
	"time"
)

func Minertarget() []data.Miner {
	var MinerList []data.Miner
	if config.Cfg.MinerListing == "ALL" {
		MinerList = getAllMiner()
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

func ShipToEs(Miner []data.MinerStat) {
	for key, _ := range Miner {
		es.Bulk("2miners-data", es.PrepareDataForEs(Miner[key]))
	}
}

func Life() {
	es.Connection()
	MinerList := Minertarget()
	if MinerList == nil {
		log.Fatal("Minerlist is empty, abort")
	}
	Urls := GetTargetUrl(MinerList)
	for {
		log.Printf("Collecting for clock %s.", data.Clock)
		localTime := time.Now()
		MinerInfo := HarvestMinerStat(Urls)
		ShipToEs(MinerInfo)
		duration := time.Since(localTime)
		log.Printf("Collected & saved miner data. Duration :%f secondes", duration.Seconds())
		time.Sleep(2 * time.Second)
	}
}
