package data

type PaymentsModel struct {
	amount      int64  `json:"amount,omitempty"`
	timestamp   int64  `json:"timestamp,omitempty"`
	totalPayees int64  `json:"totalPayees,omitempty"`
	tx          string `json:"tx,omitempty"`
}

type PaymentReturnModel struct {
	payments      []PaymentsModel `json:"payments,omitempty"`
	paymentsTotal int64           `json:"payments_total,omitempty"`
}

type RewardsModel struct {
	blockheight int64   `json:"blockheight,omitempty"`
	timestamp   int64   `json:"timestamp,omitempty"`
	blockhash   string  `json:"blockhash,omitempty"`
	reward      int64   `json:"reward,omitempty"`
	percent     float64 `json:"percent,omitempty"`
	immature    bool    `json:"immature,omitempty"`
	currentLuck float64 `json:"currentLuck,omitempty"`
	uncle       bool    `json:"uncle,omitempty"`
}

type StatsModel struct {
	balance     int64 `json:"balance,omitempty"`
	blocksFound int64 `json:"blocksFound,omitempty"`
	immature    int64 `json:"immature,omitempty"`
	lastShare   int64 `json:"lastShare,omitempty"`
	paid        int64 `json:"paid,omitempty"`
	pending     bool  `json:"pending,omitempty"`
}

type SumrewardsModel struct {
	inverval  int64  `json:"inverval,omitempty"`
	reward    int64  `json:"reward,omitempty"`
	numreward int64  `json:"numreward,omitempty"`
	name      string `json:"name,omitempty"`
	offset    int64  `json:"offset,omitempty"`
}

type WorkerGroupModel struct {
	lastBeat string  `json:"lastBeat,omitempty"`
	hr       float64 `json:"hr,omitempty"`
	offline  bool    `json:"offline,omitempty"`
	hr2      float64 `json:"hr2,omitempty"`
}

type WorkerModel struct {
	workerGroup WorkerGroupModel `json:"workerGroup"`
}

type AccountReturnModel struct {
	currentHashrate float64           `json:"currentHashrate,omitempty"`
	currentLuck     string            `json:"currentLuck,omitempty"`
	hashrate        float64           `json:"hashrate,omitempty"`
	pageSize        int64             `json:"pageSize,omitempty"`
	payments        []PaymentsModel   `json:"payments,omitempty"`
	paymentsTotal   int64             `json:"paymentsTotal,omitempty"`
	rewards         []RewardsModel    `json:"rewards,omitempty"`
	roundShares     int64             `json:"roundShares,omitempty"`
	shares          []string          `json:"shares,omitempty"`
	stats           StatsModel        `json:"stats"`
	sumrewards      []SumrewardsModel `json:"sumrewards,omitempty"`
	workers         WorkerModel       `json:"workers"`
	workersOffline  int64             `json:"workersOffline,omitempty"`
	workersOnline   int64             `json:"workersOnline,omitempty"`
	workersTotal    int64             `json:"workersTotal,omitempty"`
	hreward         int64             `json:"24hreward,omitempty"`
	hnumreward      int64             `json:"24hnumreward,omitempty"`
}

type MinerUidModel struct {
	lastBeat int64 `json:"lastBeat,omitempty"`
	height   int64 `json:"height,omitempty"`
	offline  bool  `json:"offline,omitempty"`
}

type MinerModel struct {
	MinerUid MinerUidModel `json:"minerUid"`
}

type MinerReturnModel struct {
	Hashrate    float64      `json:"hashrate,omitempty"`
	Miners      []MinerModel `json:"miners,omitempty"`
	MinersTotal int64        `json:"minersTotal,omitempty"`
	Now         int64        `json:"now,omitempty"`
}

type Miner struct {
	Adress      string
	LastBeat    float64
	Hr          float64
	Offline     bool
	CurrentLuck string
}

type MinerInfo struct {
	Wallet         string `json:"wallet"`
	Timestamp      string `json:"@timestamp"`
	Two4Hnumreward int    `json:"24hnumreward"`
	Two4Hreward    int    `json:"24hreward"`
	APIVersion     int    `json:"apiVersion"`
	Config         struct {
		AllowedMaxPayout int64  `json:"allowedMaxPayout"`
		AllowedMinPayout int    `json:"allowedMinPayout"`
		DefaultMinPayout int    `json:"defaultMinPayout"`
		IPHint           string `json:"ipHint"`
		IPWorkerName     string `json:"ipWorkerName"`
		MinPayout        int    `json:"minPayout"`
	} `json:"config"`
	CurrentHashrate int    `json:"currentHashrate"`
	CurrentLuck     string `json:"currentLuck"`
	Hashrate        int    `json:"hashrate"`
	PageSize        int    `json:"pageSize"`
	Payments        []struct {
		Amount    int    `json:"amount"`
		Timestamp int    `json:"timestamp"`
		Tx        string `json:"tx"`
		TxFee     int    `json:"txFee"`
	} `json:"payments"`
	PaymentsTotal int `json:"paymentsTotal"`
	Rewards       []struct {
		Blockheight int     `json:"blockheight"`
		Timestamp   int     `json:"timestamp"`
		Reward      int     `json:"reward"`
		Percent     float64 `json:"percent"`
		Immature    bool    `json:"immature"`
		Orphan      bool    `json:"orphan"`
		Uncle       bool    `json:"uncle"`
	} `json:"rewards"`
	RoundShares   int `json:"roundShares"`
	SharesInvalid int `json:"sharesInvalid"`
	SharesStale   int `json:"sharesStale"`
	SharesValid   int `json:"sharesValid"`
	Stats         struct {
		Balance     int `json:"balance"`
		BlocksFound int `json:"blocksFound"`
		Gas         int `json:"gas"`
		Immature    int `json:"immature"`
		LastShare   int `json:"lastShare"`
		Paid        int `json:"paid"`
		Pending     int `json:"pending"`
	} `json:"stats"`
	Sumrewards []struct {
		Inverval  int    `json:"inverval"`
		Reward    int    `json:"reward"`
		Numreward int    `json:"numreward"`
		Name      string `json:"name"`
		Offset    int    `json:"offset"`
	} `json:"sumrewards"`
	UpdatedAt int64 `json:"updatedAt"`
	Workers   struct {
	} `json:"workers"`
	WorkersOffline int `json:"workersOffline"`
	WorkersOnline  int `json:"workersOnline"`
	WorkersTotal   int `json:"workersTotal"`
}
