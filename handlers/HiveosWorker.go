package handlers

import (
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"2miner-monitoring/redis"
	"2miner-monitoring/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func setHiveosWorkerFlightsheet(FlightSheet data.FlightSheet, WorkerHarvestTime, Name, farmOwner string) {
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

func setHiveosWorkerGpusInfo(GpuStats data.GpuStats, GpuInfo data.GpuInfo, WorkerHarvestTime, name, farmOwner string, CardControlIndex *data.HiveosCardLinker) {
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
				//				esHiveosWorkerGpuJson, _ := json.Marshal(HiveosWorkerGpu)
				//				es.Bulk("2miners-hiveos-worker-gpu-info", string(esHiveosWorkerGpuJson))
				CardControlIndex.GPU = append(CardControlIndex.GPU, HiveosWorkerGpu)
				break
			}
		}
	}
}

func setHiveosWorkerGpus(Gpus data.Gpus, WorkerHarvestTime, workerName, farmOwner string) {
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

func setHiveosWorkerOverclock(Overclock data.Overclock, WorkerHarvestTime, workerName, farmOwner string, CardControlIndex *data.HiveosCardLinker) {
	esOverclock := data.HiveosOverclock{}
	//TODO: find another way to make it work with rig composed with AMD & NVIDIA
	// CardControlIndex & thoses loop beyond are broken is this case
	if Overclock.Nvidia.FanSpeed != "" {
		FanSpeed := strings.Split(Overclock.Nvidia.FanSpeed, " ")
		MemClock := strings.Split(Overclock.Nvidia.MemClock, " ")
		CoreClock := strings.Split(Overclock.Nvidia.CoreClock, " ")
		PowerLimit := strings.Split(Overclock.Nvidia.PowerLimit, " ")
		for i, _ := range FanSpeed { //Some value exepct Fan speed seems to be packed sometimes...
			k := i
			if len(CoreClock) < i {
				k = 0
			}
			esOverclock.Nvidia.MemClock = MemClock[k]
			esOverclock.Nvidia.CoreClock = CoreClock[k]
			esOverclock.Nvidia.PowerLimit = PowerLimit[k]
			esOverclock.Nvidia.FanSpeed = FanSpeed[i]
			CardControlIndex.GPU[i].HiveosOverclock = esOverclock
			esOverclockJson, _ := json.Marshal(esOverclock)
			es.Bulk("2miners-hiveos-gpu-total-info", string(esOverclockJson))
		}
	}
	if Overclock.Amd.FanSpeed != "" {
		MemMvdd := strings.Split(Overclock.Amd.MemMvdd, " ")
		CoreVddc := strings.Split(Overclock.Amd.CoreVddc, " ")
		FanSpeed := strings.Split(Overclock.Amd.FanSpeed, " ")
		MemClock := strings.Split(Overclock.Amd.MemClock, " ")
		MemVddci := strings.Split(Overclock.Amd.MemVddci, " ")
		CoreClock := strings.Split(Overclock.Amd.CoreClock, " ")
		CoreState := strings.Split(Overclock.Amd.CoreState, " ")
		esOverclock.Amd.Aggressive = Overclock.Amd.Aggressive
		for i, _ := range FanSpeed { //Some value exepct Fan speed seems to be packed sometimes...
			k := i
			if len(CoreClock) < i {
				k = 0
			}
			esOverclock.Amd.MemMvdd = MemMvdd[k]
			esOverclock.Amd.CoreVddc = CoreVddc[k]
			esOverclock.Amd.MemClock = MemClock[k]
			esOverclock.Amd.MemVddci = MemVddci[k]
			esOverclock.Amd.CoreClock = CoreClock[k]
			esOverclock.Amd.CoreState = CoreState[k]
			esOverclock.Amd.FanSpeed = FanSpeed[i]
			CardControlIndex.GPU[i].HiveosOverclock = esOverclock
			esOverclockJson, _ := json.Marshal(esOverclock)
			es.Bulk("2miners-hiveos-gpu-total-info", string(esOverclockJson))
		}
	}

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
			CarWorkerLinker := data.HiveosCardLinker{}
			worker.Timestamp = WorkerHarvestTime
			farmId := fmt.Sprintf("%d", worker.FarmID)
			farmOwner := redis.GetFromToRedis(0, farmId)
			worker.HiveOwner = farmOwner
			setHiveosWorkerFlightsheet(worker.FlightSheet, WorkerHarvestTime, worker.Name, farmOwner)
			worker.FlightSheet = data.FlightSheet{}
			setHiveosWorkerGpusInfo(worker.GpuStats, worker.GpuInfo, WorkerHarvestTime, worker.Name, farmOwner, &CarWorkerLinker)
			worker.GpuInfo = data.GpuInfo{}
			worker.GpuStats = data.GpuStats{}
			setHiveosWorkerGpus(worker.GpuSummary.Gpus, WorkerHarvestTime, worker.Name, farmOwner)
			worker.GpuSummary.Gpus = data.Gpus{}
			setHiveosWorkerOverclock(worker.Overclock, WorkerHarvestTime, worker.Name, farmOwner, &CarWorkerLinker)
			worker.Overclock = data.Overclock{}
			workerJson, _ := json.Marshal(worker)
			es.Bulk("2miners-hiveos-worker", string(workerJson))
		}
	}
	c.String(200, "Workers Harvested")
}
