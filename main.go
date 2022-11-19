package main

import (
	"2miner-monitoring/cli"
	"2miner-monitoring/config"
	"2miner-monitoring/engine"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/server"
	"2miner-monitoring/utils"
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
