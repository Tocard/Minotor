package handlers

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/thirdapp"
	"encoding/json"
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
	for farm, _ := range Farms.Data {
		Farms.Data[farm].Timestamp = FarmHarvestTime
		farmJson, _ := json.Marshal(Farms.Data[farm])
		es.Bulk("2miners-hiveos-farm", string(farmJson))
	}
	c.String(code, "Farm harvested")
}

func GetHiveosWorkers(c *gin.Context) {
	farmid := c.Param("farmid")
	farmId, _ := strconv.Atoi(farmid)
	code, res := thirdapp.HiveosGetWorkers(farmId)
	c.String(code, res)
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
