package main

import (
	"minotor/cli"
	"minotor/config"
	"minotor/engine"
	"minotor/es"
	"minotor/redis"
	"minotor/server"
	"minotor/utils"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	utils.CreateNodes()
	redis.InitRedis()
	es.Connection()
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(10).Minutes().Do(engine.FluxNodeRentability)
		s.Every(5).Minutes().Do(engine.GetLastEthTx)
		s.Every(5).Minutes().Do(engine.FluxNodesOverview)

		s.Every(1).Minutes().Do(engine.GetCosmosTokens)
		s.Every(1).Minutes().Do(engine.GetCosmosMarket)
		s.Every(1).Minutes().Do(engine.GetLastEthBlock)
		s.Every(1).Minutes().Do(engine.HarvestCoinPrice)

		s.StartAsync()
	}()
	server.GoGinServer()
}
