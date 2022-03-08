package handlers

import (
	"2miner-monitoring/thirdapp"
	"github.com/gin-gonic/gin"
)

func ScrapHashrateNo(c *gin.Context) {
	thirdapp.RunCrawler()
	c.String(200, "OK")

}
