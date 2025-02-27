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
	"minotor/db"
	"minotor/es"
	"minotor/utils"
	"time"
)

// AutonomysHarvestWalletSS58 retrieves balance using an SS58 address
/*func AutonomysHarvestWalletSS58(c *gin.Context) {
	ss58Address := c.Param("wallet")

	pubKey, prefix, err := utils.DecodeSS58(ss58Address)
	if err != nil {
		log.Println("Error:", err)
		c.String(400, "Invalid SS58 address")
		return
	}

	log.Printf("Custom Prefix: %d\nPublic Key: %x\n", prefix, pubKey)
	c.String(200, fmt.Sprintf("Custom Prefix: %d\nPublic Key: %x", prefix, pubKey))
}*/

func AutonomysHarvestWallet(c *gin.Context) {
	var Result = [][]byte{}
	var TimeMarker = time.Now().Format(time.RFC3339)
	meta, err := autonomys.Node.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("Failed to fetch metadata: %v\n", err)
	}
	Wallets, _ := db.GetAllWallets()
	var addresses []string
	for _, Wallet := range Wallets {
		addresses = append(addresses, Wallet.Address)
	}
	for _, ss58Address := range addresses {
		pubKey, _, err := utils.DecodeSS58(ss58Address)
		if err != nil {
			log.Println("Error:", err)
			c.String(400, "Invalid SS58 address")
			return
		}
		key, err := types.CreateStorageKey(meta, "System", "Account", pubKey)
		if err != nil {
			log.Printf("Failed to create storage key for public key 0x%s: %v\n", hex.EncodeToString(pubKey), err)
			continue
		}

		var accountInfo types.AccountInfo
		ok, err := autonomys.Node.RPC.State.GetStorageLatest(key, &accountInfo)
		if err != nil || !ok {
			log.Printf("Failed to get storage for public key 0x%s: %v\n", hex.EncodeToString(pubKey), err)
			continue
		}

		balance := accountInfo.Data.Free

		u128 := new(big.Int)
		u128.SetString(balance.String(), 10)

		float128 := data.U128ToFloat128(u128)
		divisor := big.NewFloat(1000000000000000000)
		result := new(big.Float).Quo(float128, divisor)

		var _Wallet = data.Wallet{Address: ss58Address, Amount: result, Timestamp: TimeMarker}
		log.Printf("Address %s ss58Address and/or Public Key 0x%s has a balance of %d\n", ss58Address, hex.EncodeToString(pubKey), balance)
		rawJson, _ := _Wallet.MarshalJSON()

		Result = append(Result, rawJson)
	}
	es.BulkData("minotor-autonomys-wallet", Result)
	c.String(200, "AutonomysHarvestWallet Done")
}

func RegisterWallet(c *gin.Context) {
	ss58Address := c.Param("wallet")

	pubKey, prefix, err := utils.DecodeSS58(ss58Address)
	if err != nil {
		log.Println("Error:", err)
		c.String(400, "Invalid SS58 address")
		return
	}

	log.Printf("Custom Prefix: %d\nPublic Key: %x\n for address %s\n", prefix, pubKey, ss58Address)

	WalletExist, err := db.WalletExists(ss58Address)
	if err != nil {
		resp := fmt.Sprintf("Unable to get wallets on RegisterWallet %s", err.Error())
		c.String(503, resp)
		return
	}

	if WalletExist {
		resp := fmt.Sprintf("wallet %s already registered", ss58Address)
		c.String(204, resp)
		return
	}

	Wallet := db.NewWallet(ss58Address)
	err = Wallet.Save()
	if err != nil {
		resp := fmt.Sprintf("something went wrong while registered %s. %s", ss58Address, err.Error())
		c.String(503, resp)
		return

	}
	c.String(201, fmt.Sprintf("Custom Prefix: %d\nPublic Key: %x for adress %s", prefix, pubKey, ss58Address))

}

func UnRegisterWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	WalletExist, err := db.WalletExists(wallet)
	if err != nil {
		resp := fmt.Sprintf("Unable to get wallets on RegisterWallet %s", err.Error())
		c.String(503, resp)
		return
	}
	if data.IsValidAutonomysAddress(wallet) == false {
		resp := fmt.Sprintf("wallet %s is not a correct pulic key", wallet)
		c.String(400, resp)
		return
	}
	if WalletExist {
		Wallet, err := db.GetWalletByAdresses(wallet)
		if err != nil {
			resp := fmt.Sprintf("Error while unregistered %s: %s", wallet, err.Error())
			c.String(404, resp)
		} else {
			err = Wallet.Delete()
			if err != nil {
				resp := fmt.Sprintf("Unable to delete wallet %s, contact admin", wallet)
				c.String(503, resp)
			} else {
				resp := fmt.Sprintf("wallet %s succefully removed", wallet)
				c.String(200, resp)
			}
		}
	} else {
		resp := fmt.Sprintf("wallet %s not registered", wallet)
		c.String(200, resp)
	}
}

func ListWallet(c *gin.Context) {
	Wallets, err := db.GetAllWallets()
	if err != nil {
		resp := fmt.Sprintf("something went wrong while getting wallet %s", err.Error())
		c.String(503, resp)

	} else {
		WalletsJson, _ := json.Marshal(Wallets)

		c.String(201, string(WalletsJson))
	}
}
