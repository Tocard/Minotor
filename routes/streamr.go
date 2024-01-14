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

func GetNodeStatus(c *gin.Context) {
}

func GetAllNodesStatus(c *gin.Context) {
	var OperatorsToBulk [][]byte
	var DelegatorsToBulk [][]byte
	var SlashingsToBulk [][]byte
	var StakesToBulk [][]byte

	Operators, code := thirdapp.HarvestAllOperatorsInfo()
	if code != 200 {
		c.String(code, string(Operators))
	}
	AllOperators := data.AllOperator{}
	err := json.Unmarshal(Operators, &AllOperators)
	if err != nil {
		c.String(500, err.Error())
	}
	clock := time.Now().Format(time.RFC3339)
	for _, Operator := range AllOperators.Data.Operator {
		Operator.ConvertWeiFieldsToFloat(clock)
		for _, Delegator := range Operator.Delegations {
			Delegator.ConvertWeiFieldsToFloat(Operator.Metadata.Name, clock)
			DelegatorJson, _ := json.Marshal(Delegator)
			DelegatorsToBulk = append(DelegatorsToBulk, DelegatorJson)
		}
		for _, Slashing := range Operator.SlashingEvents {
			Slashing.ConvertWeiFieldsToFloat(Operator.Metadata.Name, clock)
			SlashingJson, _ := json.Marshal(Slashing)
			SlashingsToBulk = append(SlashingsToBulk, SlashingJson)
		}
		for _, Stakes := range Operator.Stakes {
			Stakes.ConvertWeiFieldsToFloat(Operator.Metadata.Name, clock)
			StakesJson, _ := json.Marshal(Stakes)
			StakesToBulk = append(StakesToBulk, StakesJson)
		}
		Operator.CleanupFields()
		OperatorJson, _ := json.Marshal(Operator)
		OperatorsToBulk = append(OperatorsToBulk, OperatorJson)
	}
	es.BulkData("minotor-streamr-stake", StakesToBulk)
	es.BulkData("minotor-streamr-slashing", SlashingsToBulk)
	es.BulkData("minotor-streamr-delegator", DelegatorsToBulk)
	es.BulkData("minotor-streamr-operator", OperatorsToBulk)
	c.String(201, fmt.Sprintf("%s", OperatorsToBulk))
}

func GetSlashingHistory(c *gin.Context) {
}
