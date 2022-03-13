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
	var Farmids []int
	for _, farm := range Farms.Data {
		farm.Timestamp = FarmHarvestTime
		Farmids = append(Farmids, farm.ID)
		farm.HiveOwner = farm.Owner.Name
		farmId := fmt.Sprintf("%d", farm.ID)
		redis.WriteToRedis(0, farmId, farm.Owner.Name, "long")
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
	data.HiveOsController.Id = Farmids
	jsonFarmID, _ := json.Marshal(Farmids)
	redis.WriteToRedis(0, "listFarmids", string(jsonFarmID), "long")
	c.String(code, "Farm harvested")
}

func GetHiveosWorker(c *gin.Context) {
	workerid := c.Param("workerid")
	workerId, _ := strconv.Atoi(workerid)

	code, res := thirdapp.HiveosGetWorker(0, workerId)
	c.String(code, res)
}

func GetTest(c *gin.Context) {
	test := es.EsSearch()

	c.String(200, test.String())
}
