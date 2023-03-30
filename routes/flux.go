package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/thirdapp"
)

func CalculNodesRentability(c *gin.Context) {
	Nodes := data.Nodes{}
	Code, Json := thirdapp.GetNodesStats()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err := json.Unmarshal(Json, &Nodes)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on CalculNodesRentability", err))
		return
	}
	FluxBlocsStats := data.FluxBlocsStats{}
	Code, Json = thirdapp.GetBlocStats()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err = json.Unmarshal(Json, &FluxBlocsStats)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on HarvestBlocksInfo", err))
		return
	}
	code, State := data.CalCulRentability(Nodes, FluxBlocsStats)
	c.String(code, State)
}
