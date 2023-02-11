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

		s.Every(30).Second().Do(engine.GetHoppersBalance)

		s.StartAsync()
	}()
	server.GoGinServer()
}
