package routes

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"minotor/autonomys"
	"minotor/data"
	"minotor/es"
	"minotor/utils"
	"time"
)

func AutonomysHarvestWallet(c *gin.Context) {
	var Result = [][]byte{}
	var TimeMarker = time.Now().Format(time.RFC3339)
	meta, err := autonomys.Node.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("Failed to fetch metadata: %v\n", err)
	}
	_addresses := []string{
		"0x62654ea0a3a04edf47ccd0fc36906cb27b9eb3fa3c3b7a6670c0fc167a7f2277",
	}
	addresses := utils.CleanAddressesArray(_addresses)
	for _, address := range addresses {
		key, err := types.CreateStorageKey(meta, "System", "Account", address)
		if err != nil {
			log.Printf("Failed to create storage key for public key 0x%s: %v\n", hex.EncodeToString(address), err)
			continue
		}

		var accountInfo types.AccountInfo
		ok, err := autonomys.Node.RPC.State.GetStorageLatest(key, &accountInfo)
		if err != nil || !ok {
			log.Printf("Failed to get storage for public key 0x%s: %v\n", hex.EncodeToString(address), err)
			continue
		}

		balance := accountInfo.Data.Free

		_addr := fmt.Sprintf("0x%s", hex.EncodeToString(address))
		u128 := new(big.Int)
		u128.SetString(balance.String(), 10)

		float128 := data.U128ToFloat128(u128)
		divisor := big.NewFloat(1000000000000000000)
		result := new(big.Float).Quo(float128, divisor)

		var _Wallet = data.Wallet{Address: _addr, Amount: result, Timestamp: TimeMarker}
		log.Printf("Public Key 0x%s has a balance of %d\n", hex.EncodeToString(address), balance)
		rawJson, _ := json.Marshal(_Wallet)

		Result = append(Result, rawJson)
	}
	es.BulkData("minotor-autonomys-wallet", Result)
	c.String(200, "AutonomysHarvestWallet Done")
}
