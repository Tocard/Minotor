package main

import (
	"github.com/go-co-op/gocron"
	"minotor/cli"
	"minotor/config"
	"minotor/db"
	"minotor/engine"
	"minotor/es"
	"minotor/server"
	"minotor/utils"
	"time"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	utils.CreateNodes()
	//redis.InitRedis()
	es.Connection()
	db.Migrate()
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(10).Minutes().Do(engine.FluxNodeRentability)
		s.Every(5).Minutes().Do(engine.FluxNodesOverview)

		s.Every(1).Minutes().Do(engine.GetCosmosTokens)
		s.Every(1).Minutes().Do(engine.GetCosmosMarket)
		s.Every(1).Minutes().Do(engine.HarvestCoinPrice)
		s.Every(1).Minutes().Do(engine.HarvestComsosWallet)
		s.Every(1).Minutes().Do(engine.HealthCheck)
		s.StartAsync()
	}()
	server.GoGinServer()
}
