package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/data"
	"minotor/es"
	"minotor/utils"
	"time"
)

func ChiaPoolBlockWins(c *gin.Context) {
	var BlockWinsJson [][]byte

	BlockWins, dbErr := data.GetAllBlockWins()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, BlockWin := range BlockWins {
		BlockWin.Timestamp, _ = utils.StringTimestampToRFC3339(BlockWin.Timestamp)
		BlockWin.FetchedAt = time.Now().Format(time.RFC3339)
		BlockWin.FetchedBy = "minotor"
		BlockWinJson, _ := json.Marshal(BlockWin)
		BlockWinsJson = append(BlockWinsJson, BlockWinJson)
	}
	//if err := rows.Err(); err != nil {
	//	c.String(500, fmt.Sprintf("%s error on ChiaPoolBlockWins:30", err))
	//}
	es.BulkData("minotor-chia-pool-block-win", BlockWinsJson)
	c.String(200, fmt.Sprintf("%s", BlockWinsJson))
}
