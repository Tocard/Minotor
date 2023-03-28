package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"time"
)

func GetStreamrStatus(c *gin.Context) {
	addr := c.Param("addr")
	code, Status := thirdapp.GetStreamStatsFromBruberScan(addr)
	if code != 200 {
		c.String(code, string(Status))
	}
	StreamRStatus := data.StreamR{}
	err := json.Unmarshal(Status, &StreamRStatus)
	if err != nil {
		c.String(500, err.Error())
	}
	StreamRStatus.Timestamp = time.Now().Format(time.RFC3339)
	StreamRJson, err := json.Marshal(StreamRStatus)
	if err != nil {
		c.String(500, err.Error())
	}
	es.Bulk("minotor-streamr", string(StreamRJson))
	c.String(201, string(StreamRJson))
}
