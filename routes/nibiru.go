package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"time"
)

func GetNibiruValidatorsStatus(c *gin.Context) {
	Validators := data.Validators{}
	var ValidatorsJson [][]byte

	Code, Json := thirdapp.GetNibiruValidators()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err := json.Unmarshal(Json, &Validators)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on CalculNodesRentability", err))
		return
	}
	timer := time.Now().Format(time.RFC3339)
	for _, Raw := range Validators.Validator {
		Raw.Timestamp = timer
		rawJson, _ := json.Marshal(Raw)
		ValidatorsJson = append(ValidatorsJson, rawJson)
	}
	es.BulkData("minotor-nibiru-validators", ValidatorsJson)
}
