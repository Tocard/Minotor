package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"minotor/ChiaDbPoolData"
	"minotor/es"
	"minotor/utils"
	"time"
)

func ChiaPoolFarmers(c *gin.Context) {
	var FarmersJson [][]byte

	Farmers, dbErr := ChiaDbPoolData.GetAllFarmers()
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, Farmer := range Farmers {
		Farmer.FetchedAt = time.Now().Format(time.RFC3339)
		Farmer.FetchedBy = "minotor"
		Farmer.EsTimestamp = time.Now().Format(time.RFC3339)
		FarmerJson, _ := json.Marshal(Farmer)
		FarmersJson = append(FarmersJson, FarmerJson)
	}
	//if err := rows.Err(); err != nil {
	//	c.String(500, fmt.Sprintf("%s error on ChiaPoolFarmers:30", err))
	//}
	es.BulkData("minotor-chia-pool-farmer", FarmersJson)
	c.String(200, fmt.Sprintf("%s", FarmersJson))
}

func ChiaPoolBlockWins(c *gin.Context) {
	var BlockWinsJson [][]byte
	blockHeight, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-block-win", "block_height")
	if err != nil {
		log.Println(err)
	}
	BlockWins, dbErr := ChiaDbPoolData.GetAllBlockWinsFromHeight(blockHeight)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, BlockWin := range BlockWins {
		BlockWin.EsTimestamp, _ = utils.StringTimestampToRFC3339(BlockWin.Timestamp)
		BlockWin.FetchedAt = time.Now().Format(time.RFC3339)
		BlockWin.FetchedBy = "minotor"
		BlockWinJson, _ := json.Marshal(BlockWin)
		BlockWinsJson = append(BlockWinsJson, BlockWinJson)
	}
	es.BulkData("minotor-chia-pool-block-win", BlockWinsJson)
	c.String(200, "minotor-chia-pool-block-win updated")
}

func ChiaPoolFarmerNetspace(c *gin.Context) {
	var FarmerNetspaceJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-farmer-netspace", "timestamp")
	if err != nil {
		log.Println(err)
	}
	FarmerNetspace, dbErr := ChiaDbPoolData.GetAllFarmerNetspaceFromTimestamp(timestamp)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, _FarmerNetspace := range FarmerNetspace {
		_FarmerNetspace.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_FarmerNetspace.Timestamp)
		_FarmerNetspace.FetchedAt = time.Now().Format(time.RFC3339)
		_FarmerNetspace.FetchedBy = "minotor"
		_FarmerNetspaceJson, _ := json.Marshal(_FarmerNetspace)
		FarmerNetspaceJson = append(FarmerNetspaceJson, _FarmerNetspaceJson)
	}
	es.BulkData("minotor-chia-pool-farmer-netspace", FarmerNetspaceJson)
	c.String(200, "minotor-chia-pool-farmer-netspace")
}

func ChiaPoolPoolNetspace(c *gin.Context) {
	var PoolNetspaceJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-netspace", "timestamp")
	if err != nil {
		log.Println(err)
	}
	PoolNetspace, dbErr := ChiaDbPoolData.GetAllPoolNetspaceFromTimestamp(timestamp)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, _PoolNetspace := range PoolNetspace {
		_PoolNetspace.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_PoolNetspace.Timestamp)
		_PoolNetspace.FetchedAt = time.Now().Format(time.RFC3339)
		_PoolNetspace.FetchedBy = "minotor"
		_PoolNetspaceJson, _ := json.Marshal(_PoolNetspace)
		PoolNetspaceJson = append(PoolNetspaceJson, _PoolNetspaceJson)
	}
	es.BulkData("minotor-chia-pool-netspace", PoolNetspaceJson)
	c.String(200, "minotor-chia-pool-netspace")
}

func ChiaPoolPartial(c *gin.Context) {
	var PartialJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-partial", "timestamp")
	if err != nil {
		log.Println(err)
	}
	if timestamp < 1696128382 {
		log.Println("Timestamp lower than 1696128382, so we will use pagination")
		pageSize := 50000 // Adjust the page size as needed.

		Partial, dbErr := ChiaDbPoolData.GetAllPartialFromTimestampPaginated(timestamp, pageSize, 1)
		if dbErr != nil {
			log.Println(dbErr.Error())
		}
		for _, _Partial := range Partial {
			_Partial.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_Partial.Timestamp)
			_Partial.FetchedAt = time.Now().Format(time.RFC3339)
			_Partial.FetchedBy = "minotor"
			_PartialJson, _ := json.Marshal(_Partial)
			PartialJson = append(PartialJson, _PartialJson)
		}
		es.BulkData("minotor-chia-pool-partial", PartialJson)
		c.String(200, "minotor-chia-pool-partial")
	} else {
		log.Println("Timestamp is not zero, so we will fetch delta")
		Partial, dbErr := ChiaDbPoolData.GetAllPartialFromTimestamp(timestamp)
		if dbErr != nil {
			log.Println(dbErr.Error())
		}
		for _, _Partial := range Partial {
			_Partial.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_Partial.Timestamp)
			_Partial.FetchedAt = time.Now().Format(time.RFC3339)
			_Partial.FetchedBy = "minotor"
			_PartialJson, _ := json.Marshal(_Partial)
			PartialJson = append(PartialJson, _PartialJson)
		}
		es.BulkData("minotor-chia-pool-partial", PartialJson)
		c.String(200, "minotor-chia-pool-partial")
	}
}

func ChiaPoolOldPartial(c *gin.Context) {
	var PartialJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-partial", "timestamp")
	if err != nil {
		log.Println(err)
	}
	Partial, dbErr := ChiaDbPoolData.GetAllPartialFromTimestamp(timestamp)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, _Partial := range Partial {
		_Partial.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_Partial.Timestamp)
		_Partial.FetchedAt = time.Now().Format(time.RFC3339)
		_Partial.FetchedBy = "minotor"
		_PartialJson, _ := json.Marshal(_Partial)
		PartialJson = append(PartialJson, _PartialJson)
	}
	es.BulkData("minotor-chia-pool-partial", PartialJson)
	c.String(200, "minotor-chia-pool-partial")
}

func ChiaPoolFarmerPayment(c *gin.Context) {
	var FarmerPaymentJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-farmer-payment", "timestamp")
	if err != nil {
		log.Println(err)
	}
	FarmerPayment, dbErr := ChiaDbPoolData.GetAllFarmerPaymentFromTimestamp(timestamp)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, _FarmerPayment := range FarmerPayment {
		_FarmerPayment.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_FarmerPayment.Timestamp)
		_FarmerPayment.FetchedAt = time.Now().Format(time.RFC3339)
		_FarmerPayment.FetchedBy = "minotor"
		_FarmerPaymentJson, _ := json.Marshal(_FarmerPayment)
		FarmerPaymentJson = append(FarmerPaymentJson, _FarmerPaymentJson)
	}
	es.BulkData("minotor-chia-pool-farmer-payment", FarmerPaymentJson)
	c.String(200, "minotor-chia-pool-farmer-payment")
}

func ChiaPoolFarmerUptime(c *gin.Context) {
	var FarmerUptimeJson [][]byte
	timestamp, err := es.GetMaxValueFromIndexForField("minotor-chia-pool-uptime", "timestamp")
	if err != nil {
		log.Println(err)
	}
	FarmerUptime, dbErr := ChiaDbPoolData.GetAllFarmerUptimeFromTimestamp(timestamp)
	if dbErr != nil {
		log.Println(dbErr.Error())
	}
	for _, _FarmerUptime := range FarmerUptime {
		_FarmerUptime.EsTimestamp, _ = utils.Int64TimestampToRFC3339(_FarmerUptime.Timestamp)
		_FarmerUptime.FetchedAt = time.Now().Format(time.RFC3339)
		_FarmerUptime.FetchedBy = "minotor"
		_FarmerUptimeJson, _ := json.Marshal(_FarmerUptime)
		FarmerUptimeJson = append(FarmerUptimeJson, _FarmerUptimeJson)
	}
	es.BulkData("minotor-chia-pool-uptime", FarmerUptimeJson)
	c.String(200, "minotor-chia-pool-uptime")
}
