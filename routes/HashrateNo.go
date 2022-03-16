package routes

import (
	"2miner-monitoring/thirdapp"
	"github.com/gin-gonic/gin"
)

func ScrapHashrateNo(c *gin.Context) {
	code, status := thirdapp.RunCrawler()
	c.String(code, status)

}
