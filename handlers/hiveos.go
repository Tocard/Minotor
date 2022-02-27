package handlers

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"log"
)

func HiveosRefreshToken(c *gin.Context) {
	code, res := thirdapp.HiveosRefreshToken()
	c.String(code, res)
}

func GetHiveosFarm(c *gin.Context) {
	code, res := thirdapp.HiveosGetFarms()
	var Farms = data.Farm{}
	FarmHarvestTime := time.Now().Format(time.RFC3339)
	err := json.Unmarshal(res, &Farms)
	if err != nil {
		c.String(500, err.Error())
	}
	for _, farm := range Farms.Data {
		farm.Timestamp = FarmHarvestTime
		farmId := fmt.Sprintf("%d", farm.ID)
		redis.WriteToRedis(0, farmId, farm.Owner.Name, "long")
		farm.HiveOwner = farm.Owner.Name
		for _, item := range farm.HashratesByCoin {
			var tmpHashratesByCoin = data.HashratesByCoin{}
			tmpHashratesByCoin.Coin = item.Coin
			tmpHashratesByCoin.Hashrate = item.Hashrate
			tmpHashratesByCoin.Timestamp = FarmHarvestTime
			tmpHashratesByCoin.HiveOwner = farm.Owner.Name
			tmpHashratesByCoin.Algo = item.Algo
			tmpHashratesByCoinJson, _ := json.Marshal(tmpHashratesByCoin)
			es.Bulk("2miners-hiveos-hashrate-coin", string(tmpHashratesByCoinJson))
		}
		for _, item := range farm.Hashrates {
			var tmpHashrates = data.Hashrates{}
			tmpHashrates.Hashrate = item.Hashrate
			tmpHashrates.Timestamp = FarmHarvestTime
			tmpHashrates.HiveOwner = farm.Owner.Name
			tmpHashrates.Algo = item.Algo
			tmpHashratesJson, _ := json.Marshal(tmpHashrates)
			es.Bulk("2miners-hiveos-hashrate", string(tmpHashratesJson))
		}
		farmJson, _ := json.Marshal(farm)
		es.Bulk("2miners-hiveos-farm", string(farmJson))
	}
	c.String(code, "Farm harvested")
}

func GetHiveosWorkers(c *gin.Context) {
	farmid := c.Param("farmid") //TODO: change harvest from redis or from this param
	farmId, _ := strconv.Atoi(farmid)
	code, res := thirdapp.HiveosGetWorkers(farmId)
	log.Printf("%s",res) 
	workers := data.Workers{}
	err := json.Unmarshal(res, &workers)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	WorkerHarvestTime := time.Now().Format(time.RFC3339)
	for _, worker := range workers.Data {
		worker.Timestamp = WorkerHarvestTime
		farmId := fmt.Sprintf("%d", worker.FarmID)
		worker.HiveOwner = redis.GetFromToRedis(0, farmId) //TODO: linked to first todo
		workerJson, _ := json.Marshal(worker)
		es.Bulk("2miners-hiveos-worker", string(workerJson))
	}
	c.String(code, "Workers Harvested")
}

func GetHiveosWorker(c *gin.Context) {
	workerid := c.Param("workerid")
	workerId, _ := strconv.Atoi(workerid)

	code, res := thirdapp.HiveosGetWorker(0, workerId)
	c.String(code, res)
}

func GetHiveosOc(c *gin.Context) {
	farmid := c.Param("farmid")
	farmId, _ := strconv.Atoi(farmid)
	code, res := thirdapp.HiveosGetOc(farmId)
	c.String(code, res)
}
