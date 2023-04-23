package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
)

func PostChiaPlotSummary(c *gin.Context) {
	ChiaPlotCheckSummary := data.ChiaPlotCheckSummary{}
	var ChiaPlots [][]byte

	err := c.BindJSON(&ChiaPlotCheckSummary)
	if err != nil {
		c.String(500, "Error on PostChiaPlotSummary | BindJson: %s ", err.Error())
	}
	for _, Plot := range ChiaPlotCheckSummary.Plots {
		Plot.Pseudo = ChiaPlotCheckSummary.Pseudo
		PlotJson, err := json.Marshal(Plot)
		if err != nil {
			c.String(500, "Error on PostChiaPlotSummary | Marshal Plot: %s ", err.Error())
		}
		ChiaPlots = append(ChiaPlots, PlotJson)
	}
	es.BulkData("minotor-chia_plot_check_summary", ChiaPlots)
	ChiaPlotCheckSummary.Plots = nil
	ChiaPlotCheckSummaryJson, err := json.Marshal(ChiaPlotCheckSummary)
	if err != nil {
		c.String(500, "Error on PostChiaPlotSummary | Marshal ChiaPlotCheckSummary : %s ", err.Error())
	}
	es.Bulk("minotor-chia_plot_check_summary", string(ChiaPlotCheckSummaryJson))
	c.String(201, "Data added to Minotor")
}
