package data

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
		TagIds             []interface{} `json:"tag_ids"`
		AutoTags           bool          `json:"auto_tags"`
		WorkersCount       int           `json:"workers_count"`
		RigsCount          int           `json:"rigs_count"`
		AsicsCount         int           `json:"asics_count"`
		DisabledRigsCount  int           `json:"disabled_rigs_count"`
		DisabledAsicsCount int           `json:"disabled_asics_count"`
		Owner              struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			Name  string `json:"name"`
			Me    bool   `json:"me"`
		} `json:"owner"`
		Money struct {
			IsPaid      bool    `json:"is_paid"`
			IsFree      bool    `json:"is_free"`
			PaidCause   string  `json:"paid_cause"`
			Balance     float64 `json:"balance"`
			Discount    int     `json:"discount"`
			DailyCost   float64 `json:"daily_cost"`
			MonthlyCost float64 `json:"monthly_cost"`
			DaysLeft    int     `json:"days_left"`
			Overdraft   bool    `json:"overdraft"`
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
			Algo     string `json:"algo"`
			Hashrate int    `json:"hashrate"`
		} `json:"hashrates"`
		HashratesByCoin []struct {
			Coin     string `json:"coin"`
			Algo     string `json:"algo"`
			Hashrate int    `json:"hashrate"`
		} `json:"hashrates_by_coin"`
		ChargeOnPool bool `json:"charge_on_pool"`
	} `json:"data"`
	Tags []interface{} `json:"tags"`
}
