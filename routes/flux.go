package routes

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
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

func GetNodesOverwiew(c *gin.Context) {
	var FluxNodeOverviewByte [][]byte
	FluxBlocsStats := data.FluxBlocsStats{}
	Code, Json := thirdapp.GetBlocStats()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err := json.Unmarshal(Json, &FluxBlocsStats)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetNodesOverwiew", err))
		FluxBlocsStats.TimeBetweenBlocks = 120
	}
	Code, Json = thirdapp.GetNodesOverview()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	FluxNodesOverview := data.FluxNodesOverview{}
	err = json.Unmarshal(Json, &FluxNodesOverview)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetNodesOverwiew", err))
		return
	}
	HarvestTime := time.Now().Format(time.RFC3339)
	for _, FluxNodeOverview := range FluxNodesOverview.FluxNode {
		FluxNodeOverview.Timestamp = HarvestTime
		i, _ := strconv.ParseInt(FluxNodeOverview.Lastpaid, 10, 64)
		FluxNodeOverview.Lastpaid = time.Unix(i, 0).Format(time.RFC3339)
		i, _ = strconv.ParseInt(FluxNodeOverview.Activesince, 10, 64)
		FluxNodeOverview.Activesince = time.Unix(i, 0).Format(time.RFC3339)

		Code, Json = thirdapp.GetZelNodeStatus(FluxNodeOverview.IP)
		if Code != 200 {
			c.String(Code, string(Json))
		}
		GetZelNodeStatus := data.GetZelNodeStatus{} //TODO:fix it
		err = json.Unmarshal(Json, &GetZelNodeStatus)
		if err != nil {
			c.String(500, fmt.Sprintf("%s error on GetNodesOverwiew", err))
			return
		}
		FluxNodeOverview.Status = GetZelNodeStatus.Data.Status
		EstimatedTimeToWin := time.Now().Add(time.Duration(int(FluxBlocsStats.TimeBetweenBlocks) / 60 * FluxNodeOverview.Rank))
		FluxNodeOverview.EstimatedTimeToWin = EstimatedTimeToWin.Format(time.RFC3339)
		if FluxNodeOverview.PaymentAddress == "" {
			log.Printf("Adresse : %s\n", FluxNodeOverview.PaymentAddress)
			log.Println(FluxNodeOverview)
		}
		FluxNodeOverviewJson, _ := json.Marshal(FluxNodeOverview)
		FluxNodeOverviewByte = append(FluxNodeOverviewByte, FluxNodeOverviewJson)
	}
	es.BulkData("flux-node-overview", FluxNodeOverviewByte)
	c.String(Code, string(Json))
}
