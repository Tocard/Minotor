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
	FluxServer := server.Group("/flux")
	{
		FluxServer.GET("/calcul_nodes_rentability", routes.CalculNodesRentability)
	}
	CosmosServer := server.Group("/cosmos")
	{
		CosmosServer.GET("/wrapper/", routes.WrapAllCosmosEndpoint)
		CosmosServer.GET("/GetBalance/", routes.GetCosmosWallet)
		CosmosServer.GET("/GetDelegation/", routes.GetCosmosBounding)
		CosmosServer.GET("/GetUnDelegation/", routes.GetCosmosUnBounding)
		CosmosServer.GET("/Register/:wallet", routes.RegisterWallet)
		CosmosServer.GET("/UnRegister/:wallet", routes.UnRegisterWallet)
	}
	WalletOverview := server.Group("/wallets")
	{
		WalletOverview.POST("/add", routes.PostWalletsOverview)
	}
	Streamr := server.Group("/streamr")
	{
		Streamr.GET("/status/:addr", routes.GetStreamrStatus)
	}
	Nibiru := server.Group("/nibiru")
	{
		Nibiru.GET("/validators", routes.GetNibiruValidatorsStatus)
	}
	Chia := server.Group("/chia")
	{
		Chia.POST("/plot_check_summary", routes.PostChiaPlotSummary)
		ChiaPool := Chia.Group("/pool")
		ChiaPool.GET("/blocks_win", routes.ChiaPoolBlockWins)
		ChiaPool.GET("/farmer", routes.ChiaPoolFarmers)
		ChiaPool.GET("/farmer_netspace", routes.ChiaPoolFarmerNetspace)
		ChiaPool.GET("/pool_netspace", routes.ChiaPoolPoolNetspace)
		ChiaPool.GET("/partials", routes.ChiaPoolPartial)
		ChiaPool.GET("/payments", routes.ChiaPoolFarmerPayment)
		ChiaPool.GET("/uptime", routes.ChiaPoolFarmerUptime)
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
