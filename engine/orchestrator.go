package engine

import (
	"fmt"
	"minotor/config"
	"minotor/utils"
	"net/http"
)

func EngineHarvestAutonomysWallet() {
	url := fmt.Sprintf("%s:%d/autonomys/wallet/harvest", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
