package routes

import (
	"minotor/data"
	"minotor/es"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func Health(c *gin.Context) {
	Health := data.Health{}
	Status := string("In Minotor We Trust")
	Health.Timestamp = time.Now().Format(time.RFC3339)
	Health.Status = Status
	HealthJson, _ := json.Marshal(Health)
	es.Bulk("minotor-health", string(HealthJson))
	c.String(200, Status)
}
