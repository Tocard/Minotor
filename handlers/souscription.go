package handlers

import (
	"2miner-monitoring/config"
	"2miner-monitoring/utils"
	"github.com/gin-gonic/gin"
)

func SuscribeWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	if !utils.IsValidAdresse(wallet) {
		c.String(400, "Not a valid adress")
		return
	}
	isregister := false
	for adressKey, _ := range config.Wtw.Adress {
		if config.Wtw.Adress[adressKey] == wallet {
			isregister = true
		}
	}
	if isregister == false {
		config.Wtw.Adress = append(config.Wtw.Adress, wallet)
		if utils.WriteYaml() {
			c.String(201, "Updated")
		} else {
			c.String(501, "Failed to Save it persistently, but still register in runtime")
		}
	}
	c.String(204, "Already register")
}

func UnSuscribeWallet(c *gin.Context) {
	wallet := c.Param("wallet")
	if !utils.IsValidAdresse(wallet) {
		c.String(400, "Not a valid adress")
		return
	}
	for key, val := range config.Wtw.Adress {
		if val == wallet {
			config.Wtw.Adress = append(config.Wtw.Adress[:key], config.Wtw.Adress[key+1:]...)
			if utils.WriteYaml() {
				c.String(200, "Deleted adresse")
				return
			} else {
				c.String(206, "Failed to Delete it persistently, but still unregister in runtime")
				return
			}
		}
	}
	c.String(404, "wallet not suscribed")
}
