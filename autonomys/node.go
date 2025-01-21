package autonomys

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"log"
	"minotor/config"
)

var (
	Node *gsrpc.SubstrateAPI
)

func ConnectNode() {
	var err error
	Node, err = gsrpc.NewSubstrateAPI(config.Cfg.AutonomysNodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to node: %v\n", err)
	}

	chain, err := Node.RPC.System.Chain()
	if err != nil {
		log.Fatalf("Failed to get chain info: %v\n", err)
	}
	nodeName, err := Node.RPC.System.Name()
	if err != nil {
		log.Fatalf("Failed to get node name: %v\n", err)
	}
	nodeVersion, err := Node.RPC.System.Version()
	if err != nil {
		log.Fatalf("Failed to get node version: %v\n", err)
	}

	log.Printf("Connected to chain %v using %v v%v\n", chain, nodeName, nodeVersion)
}
