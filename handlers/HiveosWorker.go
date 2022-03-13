package handlers

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func setHiveosFlightsheet(FlightSheet data.FlightSheet, WorkerHarvestTime, Name, farmOwner string) {
	for _, flightsheet := range FlightSheet.Items {
		esflight := data.EsFlightSheet{}
		esflight.FarmID = FlightSheet.FarmID
		esflight.Timestamp = WorkerHarvestTime
		esflight.HiveOwner = farmOwner
		esflight.Name = Name
		esflight.Coin = flightsheet.Coin
		esflight.Miner = flightsheet.Miner
		esflight.MinerAlt = flightsheet.MinerAlt
		esflight.Pool = flightsheet.Pool
		esflight.WalID = flightsheet.WalID
		esflightJson, _ := json.Marshal(esflight)
		es.Bulk("2miners-hiveos-flightsheet", string(esflightJson))
	}
}

func setHiveosWorkerUnit(GpuStats data.GpuStats, GpuInfo data.GpuInfo, WorkerHarvestTime, name, farmOwner string) {
	for _, TmpGpuStats := range GpuStats {
		for _, TmpGpuInfo := range GpuInfo {
			if TmpGpuInfo.BusID == TmpGpuStats.BusID {
				HiveosWorkerGpu := data.EsHiveOsWorkerGpu{}
				HiveosWorkerGpu.HiveOsWorkerMinimal.Timestamp = WorkerHarvestTime
				HiveosWorkerGpu.HiveOsWorkerMinimal.HiveOwner = farmOwner
				HiveosWorkerGpu.HiveOsWorkerMinimal.WorkerName = name
				HiveosWorkerGpu.BusID = TmpGpuStats.BusID
				HiveosWorkerGpu.BusNumber = TmpGpuStats.BusNumber
				HiveosWorkerGpu.BusNum = TmpGpuStats.BusNum
				HiveosWorkerGpu.Temp = TmpGpuStats.Temp
				HiveosWorkerGpu.Fan = TmpGpuStats.Fan
				HiveosWorkerGpu.Power = TmpGpuStats.Power
				HiveosWorkerGpu.Memtemp = TmpGpuStats.Memtemp
				HiveosWorkerGpu.Hash = TmpGpuStats.Hash
				HiveosWorkerGpu.Index = TmpGpuInfo.Index
				HiveosWorkerGpu.Brand = TmpGpuInfo.Brand
				HiveosWorkerGpu.Model = TmpGpuInfo.Model
				HiveosWorkerGpu.ShortName = TmpGpuInfo.ShortName
				HiveosWorkerGpu.Details = TmpGpuInfo.Details
				esHiveosWorkerGpuJson, _ := json.Marshal(HiveosWorkerGpu)
				es.Bulk("2miners-hiveos-worker-gpu-info", string(esHiveosWorkerGpuJson))
				break
			}
		}
	}
}

func setHiveosGpus(Gpus data.Gpus, WorkerHarvestTime, workerName, farmOwner string) {
	for _, Gpu := range Gpus {
		esGpu := data.HiveoOsGpus{}
		esGpu.Timestamp = WorkerHarvestTime
		esGpu.HiveOwner = farmOwner
		esGpu.Name = Gpu.Name
		esGpu.Amount = Gpu.Amount
		esGpu.WorkerName = workerName
		esGpuJson, _ := json.Marshal(esGpu)
		es.Bulk("2miners-hiveos-gpu", string(esGpuJson))
	}
}

func setHiveosOverclock(Overclocks data.Overclocks, WorkerHarvestTime, workerName, farmOwner string) {
	log.Println(Overclocks)
	Overclocks.Timestamp = WorkerHarvestTime
	Overclocks.WorkerName = workerName
	Overclocks.HiveOwner = farmOwner
	esOverclockJson, _ := json.Marshal(Overclocks)
	es.Bulk("2miners-hiveos-gpu-overclock", string(esOverclockJson))
}

func GetHiveosWorkers(c *gin.Context) {
	tmpHiveOsController := data.HiveOsController.Id
	for _, farmid := range tmpHiveOsController {
		_, res := thirdapp.HiveosGetWorkers(farmid)
		workers := data.Workers{}
		err := json.Unmarshal(res, &workers)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		WorkerHarvestTime := time.Now().Format(time.RFC3339)
		for _, worker := range workers.Data {
			worker.Timestamp = WorkerHarvestTime
			farmId := fmt.Sprintf("%d", worker.FarmID)
			farmOwner := redis.GetFromToRedis(0, farmId)
			worker.HiveOwner = farmOwner
			setHiveosFlightsheet(worker.FlightSheet, WorkerHarvestTime, worker.Name, farmOwner)
			worker.FlightSheet = data.FlightSheet{}
			setHiveosWorkerUnit(worker.GpuStats, worker.GpuInfo, WorkerHarvestTime, worker.Name, farmOwner)
			worker.GpuInfo = data.GpuInfo{}
			worker.GpuStats = data.GpuStats{}
			setHiveosGpus(worker.GpuSummary.Gpus, WorkerHarvestTime, worker.Name, farmOwner)
			worker.GpuSummary.Gpus = data.Gpus{}
			setHiveosOverclock(worker.Overclocks, WorkerHarvestTime, worker.Name, farmOwner)
			worker.Overclocks = data.Overclocks{}
			//TODO: delete flighshett from original data to avoid double insert
			workerJson, _ := json.Marshal(worker)
			log.Printf(string(workerJson))
			es.Bulk("2miners-hiveos-worker", string(workerJson))
		}
	}
	c.String(200, "Workers Harvested")
}
