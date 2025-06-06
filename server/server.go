package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"minotor/config"
	"minotor/docs"
	"minotor/routes"
	"time"
)

func engine() *gin.Engine {
	gin.ForceConsoleColor()
	server := gin.New()
	TemplatePath := fmt.Sprintf("%s/*", config.Cfg.APItemplatesPath)
	server.LoadHTMLGlob(TemplatePath)
	docs.SwaggerInfo.BasePath = ""
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to restrict allowed origins
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	Autonomys := server.Group("/autonomys")
	{
		Autonomys.GET("/wallet", routes.ServeWalletPage)
		AutonomysWallet := Autonomys.Group("/wallet")
		{
			AutonomysWallet.GET("/harvest", routes.AutonomysHarvestWallet)
			AutonomysWallet.GET("/register/:wallet", routes.RegisterWallet)
			AutonomysWallet.GET("/unregister/:wallet", routes.UnRegisterWallet)
			AutonomysWallet.POST("/register", routes.RegisterWalletPayload)
			AutonomysWallet.POST("/unregister", routes.UnRegisterWalletPayload)
			AutonomysWallet.GET("/list", routes.ListWallet)
		}
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
