package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"time"
)

// @Summary Health Check
// @Description Returns the health status of Minotor
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "In Minotor We Trust"
// @Router /health [get]
func Health(c *gin.Context) {
	Health := data.Health{}
	Status := string("In Minotor We Trust")
	Health.Timestamp = time.Now().Format(time.RFC3339)
	Health.Status = Status
	HealthJson, _ := json.Marshal(Health)
	es.Bulk("minotor-health", string(HealthJson))
	c.String(200, Status)
}
