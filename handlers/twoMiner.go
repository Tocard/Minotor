package handlers

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAllMiner(c *gin.Context) {
	RedisMiner := redis.GetFromToRedis(1, "Miners")
	if RedisMiner == "" {
		log.Printf("Collecting all miner at %s\n", data.Clock)
		resp, err := http.Get(config.Cfg.TwoMinersURL + "/miners")
		utils.HandleHttpError(err)
		bodyBytes, err := io.ReadAll(resp.Body)
		utils.HandleHttpError(err)
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
					} else if minerKey == "lastBeat" {
						tmpMiner.LastBeat = value.(float64)
					}
				}
				MinerArray = append(MinerArray, tmpMiner)
			}
		}
		save, _ := json.Marshal(MinerArray)
		redis.WriteToRedis(1, "Miners", string(save), "long")
		duration := time.Since(data.Clock)
		log.Printf("Collected all miner at %s. Duration :%f\n secondes", data.Clock, duration.Seconds())
	}
	c.String(200, "OK")
}

func ExtractWorkerInfo(c *gin.Context) {
	result, wallet := RequestStorage(c)
	for key, _ := range result["workers"].(map[string]interface{}) {
		miner := result["workers"].(map[string]interface{})[key].(map[string]interface{})
		tmpMiner := data.Worker{}
		tmpMiner.Name = key
		for minerKey, value := range miner {
			if minerKey == "hr" {
				tmpMiner.Hr = value.(float64)
			} else if minerKey == "offline" {
				tmpMiner.Offline = value.(bool)
			} else if minerKey == "Hr2" {
				tmpMiner.Hr2 = value.(float64)
			} else if minerKey == "lastBeat" {
				tmpMiner.LastBeat = time.Unix(int64(value.(float64)), 0).Format(time.RFC3339)
			} else if minerKey == "sharesValid" {
				tmpMiner.SharesValid = value.(float64)
			} else if minerKey == "sharesInvalid" {
				tmpMiner.SharesInvalid = value.(float64)
			} else if minerKey == "sharesStale" {
				tmpMiner.SharesStale = value.(float64)
			}
		}
		tmpMiner.Wallet = wallet
		tmpMiner.Timestamp = time.Now().Format(time.RFC3339)
		tmpMinerJson, _ := json.Marshal(tmpMiner)
		es.Bulk("2miners-worker", string(tmpMinerJson))
	}
	c.String(200, "OK")
}

func ExtractSimpleField(c *gin.Context) {
	result, wallet := RequestStorage(c)
	var esBulk data.MinerInfo
	esBulk.TwoMiners.Wallet = wallet
	esBulk.Timestamp = time.Now().Format(time.RFC3339)
	esBulk.Two4Hnumreward = result["24hnumreward"].(float64)
	esBulk.Two4Hreward = result["24hreward"].(float64)
	esBulk.APIVersion = result["apiVersion"].(float64)
	if result["allowedMaxPayout"] != nil {
		esBulk.AllowedMaxPayout = result["allowedMaxPayout"].(int64)
	}
	if result["allowedMinPayout"] != nil {
		esBulk.AllowedMinPayout = result["allowedMinPayout"].(int)
	}
	if result["defaultMinPayout"] != nil {
		esBulk.DefaultMinPayout = result["defaultMinPayout"].(int)
	}
	if result["ipHint"] != nil {
		esBulk.IPHint = result["ipHint"].(string)
	}
	if result["ipWorkerName"] != nil {
		esBulk.IPWorkerName = result["ipWorkerName"].(string)
	}
	if result["minPayout"] != nil {
		esBulk.MinPayout = result["minPayout"].(int)
	}
	esBulk.CurrentHashrate = result["currentHashrate"].(float64)
	esBulk.CurrentLuck, _ = strconv.ParseFloat(result["currentLuck"].(string), 64)
	esBulk.Hashrate = result["hashrate"].(float64)
	esBulk.PageSize = result["pageSize"].(float64)
	esBulk.UpdatedAt = result["updatedAt"].(float64)
	esBulk.WorkersOffline = result["workersOffline"].(float64)
	esBulk.WorkersOnline = result["workersOnline"].(float64)
	esBulk.WorkersTotal = result["workersTotal"].(float64)
	EsBulkJson, err := json.Marshal(esBulk)
	utils.HandleHttpError(err)
	es.Bulk("2miners-data", string(EsBulkJson))
	c.String(200, "OK")

}

func ExtractSumrewardsInfo(c *gin.Context) {
	result, wallet := RequestStorage(c)
	for key, _ := range result["sumrewards"].([]interface{}) {
		sumreward := result["sumrewards"].([]interface{})[key].(map[string]interface{})
		tmpSumrewards := data.Sumrewards{}
		for StatKey, value := range sumreward {
			if StatKey == "inverval" {
				tmpSumrewards.Inverval = value.(float64)
			} else if StatKey == "reward" {
				tmpSumrewards.Reward = value.(float64)
			} else if StatKey == "numreward" {
				tmpSumrewards.Numreward = value.(float64)
			} else if StatKey == "name" {
				tmpSumrewards.Name = value.(string)
			} else if StatKey == "offset" {
				tmpSumrewards.Offset = value.(float64)
			}
		}
		tmpSumrewards.Wallet = wallet
		tmpSumrewards.Timestamp = time.Now().Format(time.RFC3339)
		tmpsumRewardsJson, _ := json.Marshal(tmpSumrewards)
		es.Bulk("2miners-sumreward", string(tmpsumRewardsJson))
	}
	c.String(200, "OK")
}

func ExtractRewardInfo(c *gin.Context) {
	result, wallet := RequestStorage(c)
	for key, _ := range result["rewards"].([]interface{}) {
		reward := result["rewards"].([]interface{})[key].(map[string]interface{})
		tmpReward := data.Rewards{}
		for StatKey, value := range reward {
			if StatKey == "blockheight" {
				tmpReward.Blockheight = value.(float64)
			} else if StatKey == "timestamp" {
				tmpReward.RewardDate = time.Unix(int64(value.(float64)), 0).Format(time.RFC3339)
			} else if StatKey == "reward" {
				tmpReward.Reward = value.(float64)
			} else if StatKey == "percent" {
				tmpReward.Percent = value.(float64)
			} else if StatKey == "immature" {
				tmpReward.Immature = value.(bool)
			} else if StatKey == "orphan" {
				tmpReward.Orphan = value.(bool)
			} else if StatKey == "uncle" {
				tmpReward.Uncle = value.(bool)
			}
		}
		tmpReward.Wallet = wallet
		tmpReward.Timestamp = time.Now().Format(time.RFC3339)
		tmpRewardsJson, _ := json.Marshal(tmpReward)
		es.Bulk("2miners-reward", string(tmpRewardsJson))
	}
	c.String(200, "OK")

}

func ExtractPaymentInfo(c *gin.Context) {
	result, wallet := RequestStorage(c)
	for key, _ := range result["payments"].([]interface{}) {
		payment := result["payments"].([]interface{})[key].(map[string]interface{})
		tmpPayments := data.Payments{}
		for StatKey, value := range payment {
			if StatKey == "amount" {
				tmpPayments.Amount = value.(float64)
			} else if StatKey == "tx" {
				tmpPayments.Tx = value.(string)
			} else if StatKey == "txFee" {
				tmpPayments.TxFee = value.(float64)
			} else if StatKey == "timestamp" {
				tmpPayments.PaymentDate = time.Unix(int64(value.(float64)), 0).Format(time.RFC3339)
			}
		}
		tmpPayments.Wallet = wallet
		tmpPayments.Timestamp = time.Now().Format(time.RFC3339)
		tmpPaymentsJson, _ := json.Marshal(tmpPayments)
		es.Bulk("2miners-payment", string(tmpPaymentsJson))
	}
	c.String(200, "OK")
}

func ExtractStatInfo(c *gin.Context) {
	result, wallet := RequestStorage(c)
	tmpStat := data.MinerStats{}
	for StatKey, value := range result["stats"].(map[string]interface{}) {
		if StatKey == "balance" {
			tmpStat.Balance = value.(float64)
		} else if StatKey == "blocksFound" {
			tmpStat.BlocksFound = value.(float64)
		} else if StatKey == "gas" {
			tmpStat.Gas = value.(float64)
		} else if StatKey == "immature" {
			tmpStat.Immature = value.(float64)
		} else if StatKey == "lastShare" {
			tmpStat.LastShare = time.Unix(int64(value.(float64)), 0).Format(time.RFC3339)
		} else if StatKey == "paid" {
			tmpStat.Paid = value.(float64)
		} else if StatKey == "pending" {
			tmpStat.Pending = value.(float64)
		}
	}
	tmpStat.Wallet = wallet
	tmpStat.Timestamp = time.Now().Format(time.RFC3339)
	tmpStatJson, _ := json.Marshal(tmpStat)
	es.Bulk("2miners-stat", string(tmpStatJson))
	c.String(200, "OK")
}

func ExtractPoolStatInfo(c *gin.Context) {
	client := http.Client{
		Timeout: 180 * time.Second,
	}
	url := fmt.Sprintf("%s/stats", config.Cfg.TwoMinersURL)
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		c.String(500, "unable to fetch 2miners pool stats")
		return
	}
	stats, _ := ioutil.ReadAll(resp.Body)
	var statsMap map[string]interface{}
	err = json.Unmarshal(stats, &statsMap)
	utils.HandleHttpError(err)
	tmpStat := data.PoolStats{}
	tmpStat.Hashrate = statsMap["hashrate"].(float64)
	nodes := statsMap["nodes"].([]interface{})
	nodesRaw := nodes[0].(map[string]interface{})
	tmpStat.Difficulty, _ = strconv.ParseFloat(nodesRaw["difficulty"].(string), 64)
	tmpStat.Height, _ = strconv.ParseInt(nodesRaw["height"].(string), 10,  64)
	tmpStat.Timestamp = time.Now().Format(time.RFC3339)
	tmpStat.PoolName = "2miners-ETH"
	tmpStatJson, _ := json.Marshal(tmpStat)
	es.Bulk("2miners-poolstat", string(tmpStatJson))
	c.String(200, "Ok")
}
