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
	server.GET("/health", routes.Health)

	DefiExperience := server.Group("/DeFi")
	{
		DefiExperience.GET("/balances", routes.GetWalletsBalance)
		DefiExperience.GET("/subscribe/:wallet", routes.SuscribeWallet)
		DefiExperience.GET("/unsubscribe/:wallet", routes.UnSuscribeWallet)
		DefiExperience.GET("/coins/price", routes.GetCoinsPrice)
		DefiExperience.GET("/transactions", routes.GetLastTransaction)
	}
	serverETH := server.Group("/ETH")
	{
		serverETH.GET("/lastblock", routes.GetLastBlock)
	}
	FluxServer := server.Group("/flux")
	{
		FluxServer.GET("/calcul_nodes_rentability", routes.CalculNodesRentability)
		FluxServer.GET("/flux_nodes_overview", routes.GetNodesOverwiew)
	}
	CosmosServer := server.Group("/cosmos")
	{
		CosmosServer.GET("/get_tokens", routes.GetCosmosTokens)
		CosmosServer.GET("/get_market", routes.GetCosmosMarket)

	}
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
