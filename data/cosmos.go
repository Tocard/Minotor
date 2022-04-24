package data

type CosmosTokens []struct {
	Price           float64 `json:"price"`
	Liquidity       float64 `json:"liquidity"`
	Volume24H       float64 `json:"volume_24h"`
	Volume24HChange float64 `json:"volume_24h_change"`
	Price24HChange  float64 `json:"price_24h_change"`
	Name            string  `json:"name"`
	Denom           string  `json:"denom"`
	Symbol          string  `json:"symbol"`
	Timestamp       string  `json:"@timestamp"`
	Main            bool    `json:"main"`
	Exponent        int     `json:"exponent"`
}
