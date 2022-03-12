package data

var HiveOsController *HiveosControl

type HiveosControl struct {
	Id []int
}

type EsMinimal struct {
	Timestamp string `json:"@timestamp"`
	HiveOwner string `json:"hive_owner"`
}

type Hashrates struct {
	EsMinimal
	Algo     string  `json:"extract_algo"`
	Hashrate float64 `json:"extract_hashrate"`
}

type HashratesByCoin struct {
	EsMinimal
	Coin     string  `json:"extract_coin"`
	Algo     string  `json:"extract_coin_algo"`
	Hashrate float64 `json:"extract_coin_hashrate"`
}

type HiveosAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type HiveosToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Farm struct {
	Data []struct {
		EsMinimal
		Wallet             string        `json:"wallet_keyword"`
		ID                 int           `json:"id"`
		Name               string        `json:"name"`
		Timezone           string        `json:"timezone"`
		Nonfree            bool          `json:"nonfree"`
		Role               string        `json:"role"`
		TwofaRequired      bool          `json:"twofa_required"`
		Trusted            bool          `json:"trusted"`
		GpuRedTemp         int           `json:"gpu_red_temp"`
		AsicRedTemp        int           `json:"asic_red_temp"`
		GpuRedFan          int           `json:"gpu_red_fan"`
		AsicRedFan         int           `json:"asic_red_fan"`
		GpuRedAsr          int           `json:"gpu_red_asr"`
		AsicRedAsr         int           `json:"asic_red_asr"`
		GpuRedLa           float64       `json:"gpu_red_la"`
		AsicRedLa          float64       `json:"asic_red_la"`
		GpuRedCPUTemp      int           `json:"gpu_red_cpu_temp"`
		GpuRedMemTemp      int           `json:"gpu_red_mem_temp"`
		AsicRedBoardTemp   int           `json:"asic_red_board_temp"`
		AutocreateHash     string        `json:"autocreate_hash"`
		Locked             bool          `json:"locked"`
		PowerPrice         float64       `json:"power_price"`
		AutoTags           bool          `json:"auto_tags"`
		WorkersCount       int           `json:"workers_count"`
		RigsCount          int           `json:"rigs_count"`
		AsicsCount         int           `json:"asics_count"`
		DisabledRigsCount  int           `json:"disabled_rigs_count"`
		DisabledAsicsCount int           `json:"disabled_asics_count"`
		TagIds             []interface{} `json:"tag_ids"`
		Owner              struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			Name  string `json:"name"`
			Me    bool   `json:"me"`
		} `json:"owner"`
		Money struct {
			IsPaid      bool    `json:"is_paid"`
			IsFree      bool    `json:"is_free"`
			Overdraft   bool    `json:"overdraft"`
			PaidCause   string  `json:"paid_cause"`
			Discount    int     `json:"discount"`
			DaysLeft    int     `json:"days_left"`
			Balance     float64 `json:"balance"`
			DailyCost   float64 `json:"daily_cost"`
			MonthlyCost float64 `json:"monthly_cost"`
			CostDetails []struct {
				Type         int     `json:"type"`
				TypeName     string  `json:"type_name"`
				Amount       float64 `json:"amount"`
				MonthlyPrice float64 `json:"monthly_price"`
				MonthlyCost  float64 `json:"monthly_cost"`
				DailyCost    float64 `json:"daily_cost"`
			} `json:"cost_details"`
			DailyPrice    float64 `json:"daily_price"`
			MonthlyPrice  int     `json:"monthly_price"`
			DailyUseRigs  int     `json:"daily_use_rigs"`
			DailyUseAsics int     `json:"daily_use_asics"`
			PricePerRig   int     `json:"price_per_rig"`
			PricePerAsic  int     `json:"price_per_asic"`
		} `json:"money"`
		Stats struct {
			WorkersTotal      int     `json:"workers_total"`
			WorkersOnline     int     `json:"workers_online"`
			WorkersOffline    int     `json:"workers_offline"`
			WorkersOverheated int     `json:"workers_overheated"`
			WorkersNoTemp     int     `json:"workers_no_temp"`
			WorkersOverloaded int     `json:"workers_overloaded"`
			WorkersInvalid    int     `json:"workers_invalid"`
			WorkersLowAsr     int     `json:"workers_low_asr"`
			WorkersNoHashrate int     `json:"workers_no_hashrate"`
			RigsTotal         int     `json:"rigs_total"`
			RigsOnline        int     `json:"rigs_online"`
			RigsOffline       int     `json:"rigs_offline"`
			RigsPower         int     `json:"rigs_power"`
			GpusTotal         int     `json:"gpus_total"`
			GpusOnline        int     `json:"gpus_online"`
			GpusOffline       int     `json:"gpus_offline"`
			GpusOverheated    int     `json:"gpus_overheated"`
			GpusNoTemp        int     `json:"gpus_no_temp"`
			AsicsTotal        int     `json:"asics_total"`
			AsicsOnline       int     `json:"asics_online"`
			AsicsOffline      int     `json:"asics_offline"`
			AsicsPower        int     `json:"asics_power"`
			BoardsTotal       int     `json:"boards_total"`
			BoardsOnline      int     `json:"boards_online"`
			BoardsOffline     int     `json:"boards_offline"`
			BoardsOverheated  int     `json:"boards_overheated"`
			BoardsNoTemp      int     `json:"boards_no_temp"`
			CpusOnline        int     `json:"cpus_online"`
			DevicesTotal      int     `json:"devices_total"`
			DevicesOnline     int     `json:"devices_online"`
			DevicesOffline    int     `json:"devices_offline"`
			PowerDraw         int     `json:"power_draw"`
			PowerCost         float64 `json:"power_cost"`
			Asr               float64 `json:"asr"`
		} `json:"stats"`
		Hashrates []struct {
			Algo     string  `json:"algo"`
			Hashrate float64 `json:"hashrate"`
		} `json:"hashrates"`
		HashratesByCoin []struct {
			Coin     string  `json:"coin"`
			Algo     string  `json:"algo"`
			Hashrate float64 `json:"hashrate"`
		} `json:"hashrates_by_coin"`
		ChargeOnPool bool `json:"charge_on_pool"`
	} `json:"data"`
	Tags []interface{} `json:"tags"`
}

type Overclock struct {
	Aggressive bool   `json:"aggressive"`
	LogoOff    bool   `json:"logo_off"`
	MemMvdd    string `json:"mem_mvdd"`
	CoreVddc   string `json:"core_vddc"`
	FanSpeed   string `json:"fan_speed"`
	MemClock   string `json:"mem_clock"`
	MemVddci   string `json:"mem_vddci"`
	CoreClock  string `json:"core_clock"`
	PowerLimit string `json:"power_limit"`
	CoreState  string `json:"core_state"`
	Tweakers   struct {
		Amdmemtweak []struct {
			Gpus   []int `json:"gpus"`
			Params struct {
				Ref string `json:"ref"`
			} `json:"params"`
		} `json:"amdmemtweak"`
	} `json:"tweakers"`
}

type Overclocks struct {
	Overclock `json:"amd,nvidia"`
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
		HardwareStats struct {
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
		Overclocks    `json:"overclock"`
		MinersSummary struct {
			Hashrates []struct {
				Miner  string  `json:"miner"`
				Ver    string  `json:"ver"`
				Algo   string  `json:"algo"`
				Coin   string  `json:"coin"`
				//	Hash   float64 `json:"hash"`
				Shares struct {
					Accepted int `json:"accepted"`
					Rejected int `json:"rejected"`
					Invalid  int `json:"invalid"`
					Total    int `json:"total"`
					Ratio    float64 `json:"ratio"`
				} `json:"shares,omitempty"`
			} `json:"hashrates"`
		} `json:"miners_summary"`
		MinersStats struct {
			Hashrates []struct {
				Miner      string    `json:"miner"`
				Algo       string    `json:"algo"`
				Coin       string    `json:"coin"`
				Hashes     []float64 `json:"hashes"`
				Temps      []int     `json:"temps"`
				Fans       []int     `json:"fans,omitempty"`
				BusNumbers []int     `json:"bus_numbers,omitempty"`
			} `json:"hashrates"`
		} `json:"miners_stats"`
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
			Gpus []struct {
				Name   string `json:"name"`
				Amount int    `json:"amount"`
			} `json:"gpus"`
			MaxTemp int `json:"max_temp"`
			MaxFan  int `json:"max_fan"`
		} `json:"gpu_summary"`
		GpuStats []struct {
			BusID     string  `json:"bus_id"`
			BusNumber int     `json:"bus_number"`
			BusNum    int     `json:"bus_num"`
			Temp      int     `json:"temp"`
			Fan       int     `json:"fan"`
			Power     int     `json:"power"`
			Memtemp   int     `json:"memtemp"`
			Hash      float64 `json:"hash"`
		} `json:"gpu_stats"`
		GpuInfo []struct {
			BusID     string `json:"bus_id"`
			BusNumber int    `json:"bus_number"`
			Index     int    `json:"index"`
			Brand     string `json:"brand"`
			Model     string `json:"model"`
			ShortName string `json:"short_name"`
			Details   struct {
				Mem       string `json:"mem"`
				MemGb     int    `json:"mem_gb"`
				MemType   string `json:"mem_type"`
				MemOem    string `json:"mem_oem"`
				Vbios     string `json:"vbios"`
				Subvendor string `json:"subvendor"`
				Oem       string `json:"oem"`
			} `json:"details"`
		} `json:"gpu_info"`
		Options struct {
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

func InitHiveosControl() {
	HiveosTmpControl := HiveosControl{}
	HiveOsController = &HiveosTmpControl
}
