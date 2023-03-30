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
	es.Connection()
	db.Migrate()
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(10).Minutes().Do(engine.FluxNodeRentability)

		s.Every(10).Minutes().Do(engine.GetStreamR)
		s.Every(1).Minutes().Do(engine.HealthCheck)
		s.StartAsync()
	}()
	server.GoGinServer()
}
