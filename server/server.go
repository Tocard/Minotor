package server

import (
	"2miner-monitoring/config"
	"2miner-monitoring/routes"
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
	server.GET("/miners", routes.GetAllMiner)
	server.GET("/health", routes.Health)

	serverMiner := server.Group("/harvest")
	{
		serverMiner.GET("/payments/:wallet", routes.ExtractPaymentInfo)
		serverMiner.GET("/rewards/:wallet", routes.ExtractRewardInfo)
		serverMiner.GET("/workers/:wallet", routes.ExtractWorkerInfo)
		serverMiner.GET("/data/:wallet", routes.ExtractSimpleField)
		serverMiner.GET("/stats/:wallet", routes.ExtractStatInfo)
		serverMiner.GET("/sumrewards/:wallet", routes.ExtractSumrewardsInfo)
	}
	server.GET("/balances", routes.GetWalletsBalance)
	server.GET("/subscribe/:wallet", routes.SuscribeWallet)
	server.GET("/unsubscribe/:wallet", routes.UnSuscribeWallet)
	server.GET("/coins/price", routes.GetCoinsPrice)
	server.GET("/stats", routes.ExtractPoolStatInfo)
	server.GET("/transactions", routes.GetLastTransaction)
	serverETH := server.Group("/ETH")
	{
		serverETH.GET("/lastblock", routes.GetLastBlock)
	}
	hiveosServer := server.Group("/hiveos")
	{
		hiveosServer.GET("/refresh_auth", routes.HiveosRefreshToken)
		hiveosServer.GET("/farms", routes.GetHiveosFarm)
		hiveosServer.GET("/test", routes.GetTest)
		hiveosServer.GET("/workers", routes.GetHiveosWorkers)
		hiveosServer.GET("/worker/:worker", routes.GetHiveosWorker)
	}
	FluxServer := server.Group("/flux")
	{
		FluxServer.GET("/calcul_nodes_rentability", routes.CalculNodesRentability)
		FluxServer.GET("/flux_nodes_overview", routes.GetNodesOverwiew)
	}
	CosmosServer := server.Group("/cosmos/get_tokens", routes.GetCosmosTokens)
	{
		CosmosServer.GET("/calcul_nodes_rentability", routes.CalculNodesRentability)
	}

	server.GET("/hashrateNo", routes.ScrapHashrateNo)
	return server
}

func GoGinServer() {
	server := engine()
	server.Use(gin.Logger())
	if err := engine().Run(":" + fmt.Sprint(config.Cfg.APIPort)); err != nil {
		log.Fatal("Unable to start:", err)
	}
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}
