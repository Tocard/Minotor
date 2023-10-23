package main

import (
	"github.com/go-co-op/gocron"
	"log"
	"minotor/ChiaDbPool"
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
	es.Connection()
	db.Migrate()
	if err := ChiaDbPool.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(10).Minutes().Do(engine.GetStreamR)
		s.Every(1).Minutes().Do(engine.HealthCheck)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolFarmerPayment)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolFarmerUptime)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolBlockWins)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolPoolNetspace)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolFarmers)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolFarmerNetspace)
		s.Every(1).Minutes().Do(engine.EngineChiaPoolPartial)

		s.StartAsync()
	}()
	server.GoGinServer()
}
