package handlers

import (
	"2miner-monitoring/thirdapp"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetHiveosToken(c *gin.Context) {
	thirdapp.HiveosGetAuthToken()
	c.String(200, "OK")
}

func HiveosRefreshToken(c *gin.Context) {
	thirdapp.HiveosRefreshToken()
	c.String(200, "OK")
}

func GetHiveosFarm(c *gin.Context) {
	res := thirdapp.HiveosGetFarms()
	c.String(200, res)
}

func GetHiveosWorkers(c *gin.Context) {
	farmid := c.Param("farmid")
	farmId, _ := strconv.Atoi(farmid)
	res := thirdapp.HiveosGetWorkers(farmId)
	c.String(200, res)
}

func GetHiveosWorker(c *gin.Context) {
	workerid := c.Param("workerid")
	workerId, _ := strconv.Atoi(workerid)

	res := thirdapp.HiveosGetWorker(0, workerId)
	c.String(200, res)
}

func GetHiveosOc(c *gin.Context) {
	farmid := c.Param("farmid")
	farmId, _ := strconv.Atoi(farmid)
	res := thirdapp.HiveosGetWorkers(farmId)
	c.String(200, res)
}
