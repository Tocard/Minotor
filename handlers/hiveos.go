package handlers

import (
	"2miner-monitoring/thirdapp"
	"github.com/gin-gonic/gin"
)

func GetHiveosToken(c *gin.Context) {
	thirdapp.HiveosGetAuthToken()
	c.String(200, "OK")
}

func HiveosRefreshToken(c *gin.Context) {
	thirdapp.HiveosRefreshToken()
	c.String(200, "OK")
}
