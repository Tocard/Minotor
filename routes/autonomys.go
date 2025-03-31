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
	"minotor/config"
	"minotor/data"
	"minotor/db"
	"minotor/es"
	"minotor/utils"
	"net/http"
	"time"
)

// @Summary Harvest wallet
// @Description Harvests rewards from the autonomous wallet
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Success 200 {string} string "AutonomysHarvestWallet Done"
// @Failure 400 {string} string "Invalid SS58 address"
// @Router /autonomys/wallet/harvest [get]
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

// @Summary Register a wallet
// @Description Registers a new wallet to the autonomys system
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Param wallet path string true "Wallet Address"
// @Success 201 {string} string "Wallet successfully registered"
// @Failure 400 {string} string "Invalid SS58 address"
// @Failure 503 {string} string "Service unavailable"
// @Router /autonomys/wallet/register/{wallet} [get]
func RegisterWallet(c *gin.Context) {
	ss58Address := c.Param("wallet")
	pubKey, prefix, err := utils.DecodeSS58(ss58Address)
	if err != nil {
		log.Println("Error:", err)
		c.String(400, "Invalid SS58 address")
		return
	}
	WalletExist, err := db.WalletExists(ss58Address)
	if err != nil {
		c.String(503, fmt.Sprintf("Unable to get wallets: %s", err.Error()))
		return
	}
	if WalletExist {
		c.String(204, fmt.Sprintf("Wallet %s already registered", ss58Address))
		return
	}
	Wallet := db.NewWallet(ss58Address)
	err = Wallet.Save()
	if err != nil {
		c.String(503, fmt.Sprintf("Error registering wallet %s: %s", ss58Address, err.Error()))
		return
	}
	c.String(201, fmt.Sprintf("Custom Prefix: %d\nPublic Key: %x for address %s", prefix, pubKey, ss58Address))
}

// @Summary Register a wallet
// @Description Registers a new wallet to the autonomys system
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Param wallet body data.Wallet true "Wallet Address"
// @Success 201 {string} string "Wallet successfully registered"
// @Failure 400 {string} string "Invalid SS58 address"
// @Failure 503 {string} string "Service unavailable"
// @Router /autonomys/wallet/register [post]
func RegisterWalletPayload(c *gin.Context) {
	var Payload struct {
		Wallet string `json:"wallet"`
	}
	if err := c.ShouldBindJSON(&Payload); err != nil {
		c.String(400, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	log.Println(Payload)
	pubKey, prefix, err := utils.DecodeSS58(Payload.Wallet)
	if err != nil {
		log.Println("Error:", err)
		c.String(400, "Invalid SS58 address")
		return
	}

	WalletExist, err := db.WalletExists(Payload.Wallet)
	if err != nil {
		c.String(503, fmt.Sprintf("Unable to get wallets: %s", err.Error()))
		return
	}
	if WalletExist {
		c.String(204, fmt.Sprintf("Wallet %s already registered", Payload.Wallet))
		return
	}

	Wallet := db.NewWallet(Payload.Wallet)
	err = Wallet.Save()
	if err != nil {
		c.String(503, fmt.Sprintf("Error registering wallet %s: %s", Payload.Wallet, err.Error()))
		return
	}

	c.String(201, fmt.Sprintf("Custom Prefix: %d\nPublic Key: %x for address %s", prefix, pubKey, Payload.Wallet))
}

// @Summary Unregister a wallet
// @Description Removes a wallet from the autonomys system
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Param wallet path string true "Wallet Address"
// @Success 200 {string} string "Wallet successfully removed"
// @Failure 400 {string} string "Invalid wallet address"
// @Failure 503 {string} string "Service unavailable"
// @Router /autonomys/wallet/unregister/{wallet} [get]
func UnRegisterWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	WalletExist, err := db.WalletExists(wallet)
	if err != nil {
		c.String(503, fmt.Sprintf("Unable to get wallets: %s", err.Error()))
		return
	}
	if !data.IsValidAutonomysAddress(wallet) {
		c.String(400, fmt.Sprintf("Wallet %s is not a correct public key", wallet))
		return
	}
	if WalletExist {
		Wallet, err := db.GetWalletByAdresses(wallet)
		if err != nil {
			c.String(404, fmt.Sprintf("Error unregistering %s: %s", wallet, err.Error()))
		} else {
			err = Wallet.Delete()
			if err != nil {
				c.String(503, fmt.Sprintf("Unable to delete wallet %s, contact admin", wallet))
			} else {
				c.String(200, fmt.Sprintf("Wallet %s successfully removed", wallet))
			}
		}
	} else {
		c.String(200, fmt.Sprintf("Wallet %s not registered", wallet))
	}
}

// @Summary Unregister a wallet
// @Description Removes a wallet from the autonomys system
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Param wallet body data.Wallet true "Wallet Address"
// @Success 200 {string} string "Wallet successfully removed"
// @Failure 400 {string} string "Invalid wallet address"
// @Failure 503 {string} string "Service unavailable"
// @Router /autonomys/wallet/unregister [post]
func UnRegisterWalletPayload(c *gin.Context) {
	var Payload struct {
		Wallet string `json:"wallet"`
	}
	if err := c.ShouldBindJSON(&Payload); err != nil {
		c.String(400, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}

	WalletExist, err := db.WalletExists(Payload.Wallet)
	if err != nil {
		c.String(503, fmt.Sprintf("Unable to get wallets: %s", err.Error()))
		return
	}
	if WalletExist {
		Wallet, err := db.GetWalletByAdresses(Payload.Wallet)
		if err != nil {
			c.String(404, fmt.Sprintf("Error unregistering %s: %s", Payload.Wallet, err.Error()))
		} else {
			err = Wallet.Delete()
			if err != nil {
				c.String(503, fmt.Sprintf("Unable to delete wallet %s, contact admin", Payload.Wallet))
			} else {
				c.String(200, fmt.Sprintf("Wallet %s successfully removed", Payload.Wallet))
			}
		}
	} else {
		c.String(200, fmt.Sprintf("Wallet %s not registered", Payload.Wallet))
	}
}

// @Summary List wallets
// @Description Lists all registered wallets
// @Tags autonomys_wallet
// @Accept json
// @Produce json
// @Success 200 {object} []data.Wallet "List of wallets in JSON format"
// @Failure 503 {string} string "Service unavailable"
// @Router /autonomys/wallet/list [get]
func ListWallet(c *gin.Context) {
	Wallets, err := db.GetAllWallets()
	if err != nil {
		c.String(503, fmt.Sprintf("Something went wrong while getting wallets: %s", err.Error()))
		return
	}
	WalletsJson, _ := json.Marshal(Wallets)
	c.String(200, string(WalletsJson))
}

func ServeWalletPage(c *gin.Context) {
	log.Println(config.Cfg.APIAdress, config.Cfg.APIPort)
	c.HTML(http.StatusOK, "wallet.html", gin.H{
		"apiAddress": config.Cfg.APIAdress,
		"apiPort":    config.Cfg.APIPort,
	})
}
