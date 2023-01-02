package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"time"
)

func GetOsmosisPool(c *gin.Context) {
	var OsmosisPool [][]byte

	_, pools := thirdapp.GetAllPool()
	allpools := data.AllPools{}
	err := json.Unmarshal(pools, &allpools)
	if err != nil {
		log.Println(err.Error())
	}
	for _, pool := range allpools.Pools {
		pool.Timestamp = time.Now().Format(time.RFC3339)
		for i, token := range pool.PoolTokens {
			if i == 0 {
				pool.PoolTokenA = token
				pool.Name = token.Name
			} else {
				pool.PoolTokenB = token
				pool.Name = fmt.Sprintf("%s / %s", pool.Name, token.Name)
			}
			pool.PoolTokens = nil
		}
		poolJson, _ := json.Marshal(pool)
		OsmosisPool = append(OsmosisPool, poolJson)
	}

	es.BulkData("osmosis-pool", OsmosisPool)
	c.String(200, "Harvest all pool")
}
