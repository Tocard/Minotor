package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/config"
	"minotor/routes"
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
		DefiExperience.GET("/coins/price", routes.GetCoinsPrice)
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
		CosmosServer.GET("/wrapper/", routes.WrapAllCosmosEndpoint)
		CosmosServer.GET("/GetBalance/", routes.GetCosmosWallet)
		CosmosServer.GET("/GetDelegation/", routes.GetCosmosBounding)
		CosmosServer.GET("/GetUnDelegation/", routes.GetCosmosUnBounding)
		CosmosServer.GET("/Register/:wallet", routes.RegisterWallet)
		CosmosServer.GET("/UnRegister/:wallet", routes.UnRegisterWallet)
		CosmosServer.GET("/hoppers", routes.GetHopersBalance)
	}
	OsmosisServer := server.Group("/osmosis")
	{
		OsmosisServer.GET("/getPools", routes.GetOsmosisPool)
	}
	GrafanaServer := server.Group("/grafana")
	{
		GrafanaServer.GET("/Register/:user", routes.CreateGrafanaUser)
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
