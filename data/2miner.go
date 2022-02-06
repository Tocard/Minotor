package data

import "io"

type TwoMiners struct {
	Timestamp string `json:"@timestamp"`
	Wallet    string `json:"wallet_keyword"`
}

type Miner struct {
	Adress      string
	LastBeat    float64
	Hr          float64
	Offline     bool
	CurrentLuck string
}

type Worker struct {
	TwoMiners
	Name          string  `json:"name_keyword"`
	LastBeat      float64 `json:"lastBeat"`
	Hr            float64 `json:"hr"`
	Offline       bool    `json:"offline"`
	Hr2           float64 `json:"hr2"`
	Rhr           float64 `json:"rhr"`
	SharesValid   float64 `json:"sharesValid"`
	SharesInvalid float64 `json:"sharesInvalid"`
	SharesStale   float64 `json:"sharesStale"`
}

type MinerInfo struct {
	TwoMiners
	Two4Hnumreward   float64 `json:"24hnumreward"`
	Two4Hreward      float64 `json:"24hreward"`
	APIVersion       float64 `json:"apiVersion"`
	AllowedMaxPayout int64   `json:"allowedMaxPayout"`
	AllowedMinPayout int     `json:"allowedMinPayout"`
	DefaultMinPayout int     `json:"defaultMinPayout"`
	IPHint           string  `json:"ipHint_keyword"`
	IPWorkerName     string  `json:"ipWorkerName_keyword"`
	MinPayout        int     `json:"minPayout"`
	CurrentHashrate  float64 `json:"currentHashrate"`
	CurrentLuck      float64 `json:"currentLuck"`
	Hashrate         float64 `json:"hashrate"`
	PageSize         float64 `json:"pageSize"`
	UpdatedAt        float64 `json:"updatedAt"`
	WorkersOffline   float64 `json:"workersOffline"`
	WorkersOnline    float64 `json:"workersOnline"`
	WorkersTotal     float64 `json:"workersTotal"`
	PaymentsTotal    int     `json:"paymentsTotal"`
	RoundShares      int     `json:"roundShares"`
	SharesInvalid    int     `json:"sharesInvalid"`
	SharesStale      int     `json:"sharesStale"`
	SharesValid      int     `json:"sharesValid"`
	Stats            struct {
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
	Payments []struct {
		Amount    int    `json:"amount"`
		Timestamp int    `json:"timestamp"`
		Tx        string `json:"tx"`
		TxFee     int    `json:"txFee"`
	} `json:"payments"`
	Rewards []struct {
		Blockheight int     `json:"blockheight"`
		Timestamp   int     `json:"timestamp"`
		Reward      int     `json:"reward"`
		Percent     float64 `json:"percent"`
		Immature    bool    `json:"immature"`
		Orphan      bool    `json:"orphan"`
		Uncle       bool    `json:"uncle"`
	} `json:"rewards"`
}

type MinerStat struct {
	Json   io.ReadCloser
	Wallet string
}
