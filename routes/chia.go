package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
)

func PostChiaPlotSummary(c *gin.Context) {
	ChiaPlotCheckSummary := data.ChiaPlotCheckSummary{}

	err := c.BindJSON(&ChiaPlotCheckSummary)
	if err != nil {
		c.String(500, "Error on PostChiaPlotSummary | BindJson: %s ", err.Error())
	}
	ChiaPlotCheckSummaryJson, err := json.Marshal(ChiaPlotCheckSummary)
	if err != nil {
		c.String(500, "Error on PostChiaPlotSummary | Marshal: %s ", err.Error())
	}
	es.Bulk("minotor-chia_plot_check_summary", string(ChiaPlotCheckSummaryJson))
	c.String(201, "Data added to Minotor")
}
