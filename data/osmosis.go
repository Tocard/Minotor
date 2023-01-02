package data

type PoolTokensStruct struct {
	Name            string  `json:"name"`
	Denom           string  `json:"denom"`
	Price           float64 `json:"price"`
	Amount          float64 `json:"amount"`
	Symbol          string  `json:"symbol"`
	Display         string  `json:"display"`
	Percent         int     `json:"percent"`
	Exponent        int     `json:"exponent"`
	CoingeckoId     string  `json:"coingecko_id"`
	Price24HChange  float64 `json:"price_24h_change"`
	WeightOrScaling int64   `json:"weight_or_scaling"`
}

type PoolsStruct struct {
	Timestamp   string             `json:"@timestamp"`
	Main        bool               `json:"main"`
	Type        string             `json:"type"`
	Name        string             `json:"name"`
	PoolId      int                `json:"pool_id"`
	ExitFees    float64            `json:"exit_fees"`
	Liquidity   float64            `json:"liquidity"`
	SwapFees    float64            `json:"swap_fees"`
	Volume7D    float64            `json:"volume_7d"`
	Volume24H   float64            `json:"volume_24h"`
	PoolTokens  []PoolTokensStruct `json:"pool_tokens,omitempty"`
	PoolTokenA  PoolTokensStruct   `json:"pool_tokens_A,omitempty"`
	PoolTokenB  PoolTokensStruct   `json:"pool_tokens_B,omitempty"`
	TotalShares struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"total_shares"`
	Volume24HChange      float64 `json:"volume_24h_change"`
	FuturePoolGovernor   string  `json:"future_pool_governor"`
	Liquidity24HChange   float64 `json:"liquidity_24h_change"`
	TotalWeightOrScaling int64   `json:"total_weight_or_scaling"`
}

type AllPools struct {
	Pagination struct {
		NextOffset interface{} `json:"next_offset"`
		TotalPools int         `json:"total_pools"`
	} `json:"pagination"`
	Pools []PoolsStruct `json:"pools"`
}
