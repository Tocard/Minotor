package main

import (
	"github.com/go-co-op/gocron"
	"minotor/cli"
	"minotor/config"
	"minotor/engine"
	"minotor/es"
	"minotor/server"
	"time"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	//utils.CreateNodes()
	es.Connection()
	//db.Migrate()
	//if err := ChiaDbPool.ConnectToDB(); err != nil {
	//	log.Fatal(err)
	//}
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(1).Minutes().Do(engine.GetAllNodesStatus)

		s.StartAsync()
	}()
	server.GoGinServer()
}
