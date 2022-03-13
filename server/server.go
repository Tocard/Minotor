package server

import (
	"2miner-monitoring/config"
	"2miner-monitoring/es"
	"2miner-monitoring/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func engine() *gin.Engine {
	gin.ForceConsoleColor()
	server := gin.New()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	server.Use(gin.Recovery())
	server.GET("/miners", handlers.GetAllMiner)
	server.GET("/health", handlers.Health)

	serverMiner := server.Group("/harvest")
	{
		serverMiner.GET("/payments/:wallet", handlers.ExtractPaymentInfo)
		serverMiner.GET("/rewards/:wallet", handlers.ExtractRewardInfo)
		serverMiner.GET("/workers/:wallet", handlers.ExtractWorkerInfo)
		serverMiner.GET("/data/:wallet", handlers.ExtractSimpleField)
		serverMiner.GET("/stats/:wallet", handlers.ExtractStatInfo)
		serverMiner.GET("/sumrewards/:wallet", handlers.ExtractSumrewardsInfo)
	}
	server.GET("/balances", handlers.GetWalletsBalance)
	server.GET("/subscribe/:wallet", handlers.SuscribeWallet)
	server.GET("/unsubscribe/:wallet", handlers.UnSuscribeWallet)
	server.GET("/coins/price", handlers.GetCoinsPrice)
	server.GET("/stats", handlers.ExtractPoolStatInfo)
	server.GET("/transactions", handlers.GetLastTransaction)
	serverETH := server.Group("/ETH")
	{
		serverETH.GET("/lastblock", handlers.GetLastBlock)
	}
	hiveosServer := server.Group("/hiveos")
	{
		hiveosServer.GET("/refresh_auth", handlers.HiveosRefreshToken)
		hiveosServer.GET("/farms", handlers.GetHiveosFarm)
		hiveosServer.GET("/test", handlers.GetTest)
		hiveosServer.GET("/workers", handlers.GetHiveosWorkers)
		hiveosServer.GET("/worker/:worker", handlers.GetHiveosWorker)
	}
	server.GET("/hashrateNo", handlers.ScrapHashrateNo)
	return server
}

func GoGinServer() {
	es.Connection()
	server := engine()
	server.Use(gin.Logger())
	if err := engine().Run(":" + fmt.Sprint(config.Cfg.APIPort)); err != nil {
		log.Fatal("Unable to start:", err)
	}
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}
