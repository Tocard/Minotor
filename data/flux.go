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
