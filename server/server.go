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
	Autonomys := server.Group("/autonomys")
	{
		AutonomysWallet := Autonomys.Group("/wallet")
		AutonomysWallet.GET("/harvest", routes.AutonomysHarvestWallet)
		AutonomysWallet.GET("/register/:wallet", routes.RegisterWallet)
		AutonomysWallet.GET("/unregister/:wallet", routes.UnRegisterWallet)
		AutonomysWallet.GET("/list", routes.ListWallet)
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
