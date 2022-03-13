package data

type HiveOsWorkerMinimal struct {
	Timestamp  string `json:"@timestamp,omitempty"`
	HiveOwner  string `json:"hive_owner,omitempty"`
	WorkerName string `json:"worker_name,omitempty"`
}

type HiveosCardLinker struct {
	busID int
}

type FlightSheet struct {
	ID     int    `json:"id"`
	FarmID int    `json:"farm_id"`
	Name   string `json:"name"`
	Items  []struct {
		Coin     string `json:"coin"`
		Pool     string `json:"pool"`
		Miner    string `json:"miner"`
		MinerAlt string `json:"miner_alt,omitempty"`
		WalID    int    `json:"wal_id"`
	} `json:"items"`
}

type EsFlightSheet struct {
	EsMinimal
	FarmID   int    `json:"farm_id"`
	Name     string `json:"name"`
	Coin     string `json:"coin"`
	Pool     string `json:"pool"`
	Miner    string `json:"miner"`
	MinerAlt string `json:"miner_alt,omitempty"`
	WalID    int    `json:"wal_id"`
}

type EsHiveOsWorkerGpu struct {
	HiveOsWorkerMinimal
	BusID     string  `json:"bus_id,omitempty"`
	BusNumber int     `json:"bus_number,omitempty"`
	BusNum    int     `json:"bus_num,omitempty"`
	Temp      int     `json:"temp,omitempty"`
	Fan       int     `json:"fan,omitempty"`
	Power     int     `json:"power,omitempty"`
	Memtemp   int     `json:"memtemp,omitempty"`
	Hash      float64 `json:"hash,omitempty"`
	Index     int     `json:"index,omitempty"`
	Brand     string  `json:"brand,omitempty"`
	Model     string  `json:"model,omitempty"`
	ShortName string  `json:"short_name,omitempty"`
	Details   struct {
		Mem       string `json:"mem,omitempty"`
		MemGb     int    `json:"mem_gb,omitempty"`
		MemType   string `json:"mem_type,omitempty"`
		MemOem    string `json:"mem_oem,omitempty"`
		Vbios     string `json:"vbios,omitempty"`
		Subvendor string `json:"subvendor,omitempty"`
		Oem       string `json:"oem,omitempty"`
	} `json:"details,omitempty"`
}

type GpuStats []struct {
	BusID     string  `json:"bus_id,omitempty"`
	BusNumber int     `json:"bus_number,omitempty"`
	BusNum    int     `json:"bus_num,omitempty"`
	Temp      int     `json:"temp,omitempty"`
	Fan       int     `json:"fan,omitempty"`
	Power     int     `json:"power,omitempty"`
	Memtemp   int     `json:"memtemp,omitempty"`
	Hash      float64 `json:"hash,omitempty"`
}

type GpuInfo []struct {
	BusID     string `json:"bus_id,omitempty"`
	BusNumber int    `json:"bus_number,omitempty"`
	Index     int    `json:"index,omitempty"`
	Brand     string `json:"brand,omitempty"`
	Model     string `json:"model,omitempty"`
	ShortName string `json:"short_name,omitempty"`
	Details   struct {
		Mem       string `json:"mem,omitempty"`
		MemGb     int    `json:"mem_gb,omitempty"`
		MemType   string `json:"mem_type,omitempty"`
		MemOem    string `json:"mem_oem,omitempty"`
		Vbios     string `json:"vbios,omitempty"`
		Subvendor string `json:"subvendor,omitempty"`
		Oem       string `json:"oem,omitempty"`
	} `json:"details,omitempty"`
}

type HiveoOsGpus struct {
	HiveOsWorkerMinimal
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Gpus []struct {
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type HiveosOverclock struct {
	HiveOsWorkerMinimal
	Nvidia struct {
		LogoOff    bool   `json:"logo_off,omitempty"`
		FanSpeed   string `json:"fan_speed,omitempty"`
		MemClock   string `json:"mem_clock,omitempty"`
		CoreClock  string `json:"core_clock,omitempty"`
		PowerLimit string `json:"power_limit,omitempty"`
	} `json:"nvidia,omitempty"`
	Amd struct {
		MemMvdd    string `json:"mem_mvdd,omitempty"`
		CoreVddc   string `json:"core_vddc,omitempty"`
		FanSpeed   string `json:"fan_speed,omitempty"`
		MemClock   string `json:"mem_clock,omitempty"`
		MemVddci   string `json:"mem_vddci,omitempty"`
		Aggressive bool   `json:"aggressive,omitempty"`
		CoreClock  string `json:"core_clock,omitempty"`
		CoreState  string `json:"core_state,omitempty"`
	} `json:"amd,omitempty"`
	Tweakers struct {
		Amdmemtweak []struct {
			Params struct {
				Ref string `json:"ref,omitempty"`
			} `json:"params,omitempty"`
		} `json:"amdmemtweak,omitempty"`
	} `json:"tweakers,omitempty"`
}

type Overclock struct {
	Nvidia struct {
		LogoOff    bool   `json:"logo_off,omitempty"`
		FanSpeed   string `json:"fan_speed,omitempty"`
		MemClock   string `json:"mem_clock,omitempty"`
		CoreClock  string `json:"core_clock,omitempty"`
		PowerLimit string `json:"power_limit,omitempty"`
	} `json:"nvidia,omitempty"`
	Amd struct {
		MemMvdd    string `json:"mem_mvdd,omitempty"`
		CoreVddc   string `json:"core_vddc,omitempty"`
		FanSpeed   string `json:"fan_speed,omitempty"`
		MemClock   string `json:"mem_clock,omitempty"`
		MemVddci   string `json:"mem_vddci,omitempty"`
		Aggressive bool   `json:"aggressive,omitempty"`
		CoreClock  string `json:"core_clock,omitempty"`
		CoreState  string `json:"core_state,omitempty"`
	} `json:"amd,omitempty"`
	Tweakers struct { //TODO: Handle it
		Amdmemtweak []struct {
			Gpus   []int `json:"gpus,omitempty"`
			Params struct {
				Ref string `json:"ref,omitempty"`
			} `json:"params,omitempty"`
		} `json:"amdmemtweak,omitempty"`
	} `json:"tweakers,omitempty"`
}

type Workers struct {
	Data []struct {
		EsMinimal
		ID           int    `json:"id"`
		FarmID       int    `json:"farm_id"`
		Name         string `json:"name"`
		Active       bool   `json:"active"`
		SystemType   string `json:"system_type"`
		NeedsUpgrade bool   `json:"needs_upgrade"`
		Versions     struct {
			Hive         string `json:"hive"`
			Kernel       string `json:"kernel"`
			AmdDriver    string `json:"amd_driver"`
			NvidiaDriver string `json:"nvidia_driver"`
		} `json:"versions"`
		Stats struct {
			Online         bool     `json:"online"`
			BootTime       int      `json:"boot_time"`
			StatsTime      int      `json:"stats_time"`
			GpusOnline     int      `json:"gpus_online"`
			GpusOffline    int      `json:"gpus_offline"`
			GpusOverheated int      `json:"gpus_overheated"`
			CpusOnline     int      `json:"cpus_online"`
			MinerStartTime int      `json:"miner_start_time"`
			PowerDraw      int      `json:"power_draw"`
			Invalid        bool     `json:"invalid"`
			LowAsr         bool     `json:"low_asr"`
			Overloaded     bool     `json:"overloaded"`
			Overheated     bool     `json:"overheated"`
			Problems       []string `json:"problems"`
		} `json:"stats"`
		HardwareInfo struct {
			Motherboard struct {
				Manufacturer string `json:"manufacturer"`
				Model        string `json:"model"`
				Bios         string `json:"bios"`
			} `json:"motherboard"`
			CPU struct {
				ID    string `json:"id"`
				Model string `json:"model"`
				Cores int    `json:"cores"`
				Aes   bool   `json:"aes"`
			} `json:"cpu"`
			Disk struct {
				Model string `json:"model"`
			} `json:"disk"`
		} `json:"hardware_info"`
		HardwareStats struct { //TODO: Parse it to harvest cpu core temps
			Df      string    `json:"df"`
			Cpuavg  []float64 `json:"cpuavg"`
			Cputemp []int     `json:"cputemp"`
			Memory  struct {
				Total int `json:"total"`
				Free  int `json:"free"`
			} `json:"memory"`
			CPUCores int `json:"cpu_cores"`
		} `json:"hardware_stats"`
		UnitsCount    int  `json:"units_count"`
		RedTemp       int  `json:"red_temp"`
		RedFan        int  `json:"red_fan"`
		RedAsr        int  `json:"red_asr"`
		RedLa         int  `json:"red_la"`
		RedCPUTemp    int  `json:"red_cpu_temp"`
		RedMemTemp    int  `json:"red_mem_temp"`
		HasAmd        bool `json:"has_amd"`
		HasNvidia     bool `json:"has_nvidia"`
		FlightSheet   `json:"flight_sheet"`
		Overclock     `json:"overclock"`
		MinersSummary struct {
			Hashrates []struct { //TODO: Parse it to harvest miner info
				Miner string `json:"miner"`
				Ver   string `json:"ver"`
				Algo  string `json:"algo"`
				Coin  string `json:"coin"`
				//	Hash   float64 `json:"hash"`
				Shares struct {
					Accepted int     `json:"accepted"`
					Rejected int     `json:"rejected"`
					Invalid  int     `json:"invalid"`
					Total    int     `json:"total"`
					Ratio    float64 `json:"ratio"`
				} `json:"shares,omitempty"`
			} `json:"hashrates"`
		} `json:"miners_summary"`
		MinersStats struct { //TODO: Parse it to harvest miner info
			Hashrates []struct {
				Miner      string    `json:"miner"`
				Algo       string    `json:"algo"`
				Coin       string    `json:"coin"`
				Hashes     []float64 `json:"hashes"`
				Temps      []int     `json:"temps"`
				Fans       []int     `json:"fans,omitempty"`
				BusNumbers []int     `json:"bus_numbers,omitempty"`
			} `json:"hashrates"`
		} `json:"miners_stats"` //TODO: Parse it to harvest hashrate stats
		Watchdog struct {
			Enabled         bool   `json:"enabled"`
			RestartTimeout  int    `json:"restart_timeout"`
			RebootTimeout   int    `json:"reboot_timeout"`
			CheckPower      bool   `json:"check_power"`
			CheckConnection bool   `json:"check_connection"`
			PowerAction     string `json:"power_action"`
			CheckGpu        bool   `json:"check_gpu"`
			Type            string `json:"type"`
			Options         struct {
				ByMiner []interface{} `json:"by_miner"`
				ByAlgo  []struct {
					Algo    string  `json:"algo"`
					Minhash float64 `json:"minhash"`
					Units   string  `json:"units"`
				} `json:"by_algo"`
			} `json:"options"`
		} `json:"watchdog,omitempty"`
		GpuSummary struct {
			Gpus    `json:"gpus"`
			MaxTemp int `json:"max_temp"`
			MaxFan  int `json:"max_fan"`
		} `json:"gpu_summary"`
		GpuStats `json:"gpu_stats"`
		GpuInfo  `json:"gpu_info"`
		Options  struct {
			MaintenanceMode int `json:"maintenance_mode"`
		} `json:"options,omitempty"`
		Autofan struct {
			Enabled        bool `json:"enabled"`
			TargetTemp     int  `json:"target_temp"`
			TargetMemTemp  int  `json:"target_mem_temp"`
			MinFan         int  `json:"min_fan"`
			MaxFan         int  `json:"max_fan"`
			CriticalTemp   int  `json:"critical_temp"`
			NoAmd          bool `json:"no_amd"`
			RebootOnErrors bool `json:"reboot_on_errors"`
			SmartMode      bool `json:"smart_mode"`
		} `json:"autofan,omitempty"`
	} `json:"data"`
	Tags []struct {
		EsMinimal
		ID           int         `json:"id"`
		TypeID       int         `json:"type_id"`
		FarmID       int         `json:"farm_id"`
		UserID       interface{} `json:"user_id"`
		Name         string      `json:"name"`
		Color        int         `json:"color"`
		IsAuto       bool        `json:"is_auto"`
		WorkersCount int         `json:"workers_count"`
		FarmsCount   interface{} `json:"farms_count"`
	} `json:"tags"`
}
