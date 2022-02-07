package main

import (
	"2miner-monitoring/cli"
	"2miner-monitoring/config"
	"2miner-monitoring/server"
	"2miner-monitoring/utils"
	"fmt"
	"net/http"
)

func task() {
	resp, err := http.Get("http://localhost:8080/cosmos/proc_winner")
	utils.HandleHttpError(err)
	fmt.Println(resp)
}

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	//log2miner.InitLogger("2miner.log2miner")
	//redis.InitRedis()
	server.GoGinServer()
}
