package main

import (
	"2miner-monitoring/cli"
	"2miner-monitoring/config"
	"2miner-monitoring/engine"
	"2miner-monitoring/log2miner"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	log2miner.InitLogger("2miner.log2miner")
	engine.Life()
}
