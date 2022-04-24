package main

import (
	"2miner-monitoring/cli"
	"2miner-monitoring/config"
	"2miner-monitoring/data"
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
	config.LoadCardYamlConfig()
	data.InitHiveosControl()
	utils.CreateNodes()
	//log2miner.InitLogger("2miner.log2miner")
	redis.InitRedis()
	es.Connection()
	go func() {
		s := gocron.NewScheduler(time.Local)
		// cron expressions supported
		s.Every(1).Hours().Do(engine.HarvestMiners)
		s.Every(1).Hours().Do(engine.HarvestFactory, "rewards")

		s.Every(10).Minutes().Do(engine.FluxNodeRentability)
		s.Every(10).Minutes().Do(engine.HarvestBalance)
		s.Every(10).Minutes().Do(engine.HarvestPoolStat)
		s.Every(10).Minutes().Do(engine.HarvestFactory, "stats")
		s.Every(10).Minutes().Do(engine.HarvestFactory, "payments")
		s.Every(10).Minutes().Do(engine.HarvestFactory, "sumrewards")

		s.Every(5).Minutes().Do(engine.GetLastEthTx)
		s.Every(5).Minutes().Do(engine.FluxNodesOverview)

		s.Every(1).Minutes().Do(engine.GetCosmosTokens)
		s.Every(1).Minutes().Do(engine.GetHiveosFarm)
		s.Every(1).Minutes().Do(engine.GetHiveosWorkers)
		s.Every(1).Minutes().Do(engine.GetLastEthBlock)
		s.Every(1).Minutes().Do(engine.HarvestCoinPrice)
		s.Every(1).Minutes().Do(engine.ScrapHashrateNo)

		s.Every(10).Seconds().Do(engine.HarvestFactory, "data")

		s.Every(1).Second().Do(engine.HarvestFactory, "workers")
		// you can start running the scheduler in two different ways:
		// starts the scheduler asynchronously
		s.StartAsync()
	}()
	server.GoGinServer()
}
