package data

type Nodes struct {
	Status string `json:"status"`
	Data   struct {
		Total          int `json:"total"`
		Stable         int `json:"stable"`
		BasicEnabled   int `json:"basic-enabled"`
		SuperEnabled   int `json:"super-enabled"`
		BamfEnabled    int `json:"bamf-enabled"`
		CumulusEnabled int `json:"cumulus-enabled"`
		NimbusEnabled  int `json:"nimbus-enabled"`
		StratusEnabled int `json:"stratus-enabled"`
		Ipv4           int `json:"ipv4"`
		Ipv6           int `json:"ipv6"`
		Onion          int `json:"onion"`
	} `json:"data"`
}

type FluxBlocsStats struct {
	NBlocksMined         int     `json:"n_blocks_mined"`
	TimeBetweenBlocks    float64 `json:"time_between_blocks"`
	MinedCurrencyAmount  int64   `json:"mined_currency_amount"`
	TransactionFees      int     `json:"transaction_fees"`
	NumberOfTransactions int     `json:"number_of_transactions"`
	OutputsVolume        int64   `json:"outputs_volume"`
	Difficulty           string  `json:"difficulty"`
	NetworkHashPs        float64 `json:"network_hash_ps"`
	BlocksByPool         []struct {
		Address      string `json:"address"`
		PoolName     string `json:"poolName"`
		URL          string `json:"url"`
		BlocksFound  int    `json:"blocks_found"`
		PercentTotal string `json:"percent_total"`
	} `json:"blocks_by_pool"`
}

type FluxNodeOverview struct {
	Timestamp           string `json:"@timestamp"`
	Collateral          string `json:"collateral"`
	Txhash              string `json:"txhash"`
	Outidx              string `json:"outidx"`
	IP                  string `json:"ip"`
	Tier                string `json:"tier"`
	PaymentAddress      string `json:"payment_address"`
	Pubkey              string `json:"pubkey"`
	Activesince         string `json:"activesince,omitempty"`
	Lastpaid            string `json:"lastpaid,omitempty"`
	Network             string `json:"network"`
	Amount              string `json:"amount"`
	Status              string `json:"status"`
	EstimatedTimeToWin  string `json:"estimated_date_to_win,omitempty"`
	Rank                int    `json:"rank"`
	AddedHeight         int64  `json:"added_height"`
	ConfirmedHeight     int64  `json:"confirmed_height"`
	LastConfirmedHeight int64  `json:"last_confirmed_height"`
	LastPaidHeight      int64  `json:"last_paid_height"`
}

type FluxNodesOverview struct {
	Status   string             `json:"status"`
	FluxNode []FluxNodeOverview `json:"data"`
}

type GetZelNodeStatus struct {
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}
