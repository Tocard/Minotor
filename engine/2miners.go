package engine

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func getAllMiner() []data.Miner {
	data.RefreshClock()
	log.Printf("Collecting all miner at %s\n", data.Clock)
	resp, err := http.Get(config.Cfg.TwoMinersURL + "/miners")
	utils.HandleHttpError(err)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	utils.HandleHttpError(err)

	var MinerArray []data.Miner
	miners := result["miners"].(map[string]interface{})
	for key, _ := range miners {
		if strings.HasPrefix(key, "0x") {
			miner := miners[key].(map[string]interface{})
			tmpMiner := data.Miner{}
			tmpMiner.Adress = key
			for minerKey, value := range miner {
				if minerKey == "hr" {
					tmpMiner.Hr = value.(float64)
				} else if minerKey == "offline" {
					tmpMiner.Offline = value.(bool)
				} else if minerKey == "currentLuck" {
					tmpMiner.CurrentLuck = value.(float64)
				} else if minerKey == "lastBeat" {
					tmpMiner.LastBeat = value.(float64)
				}
			}
			MinerArray = append(MinerArray, tmpMiner)
		}
	}
	duration := time.Since(data.Clock)
	data.RefreshClock()
	log.Printf("Collected all miner at %s. Duration :%f\n secondes", data.Clock, duration.Seconds())
	return MinerArray
}

func HarvestMinerStat(Urls map[string]string) []data.MinerStat {
	localTime := time.Now()
	log.Printf("HarvestMinerStat start %s\n", localTime)
	MinersInfo := AsyncGet(Urls)
	duration := time.Since(localTime)
	log.Printf("HarvestMinerStat stop %f seconde", duration.Seconds())
	return MinersInfo
}
