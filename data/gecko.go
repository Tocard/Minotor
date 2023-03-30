package data

import "time"

type GeckoAdvanceCoin struct {
	Timestamp                    string      `json:"@timestamp"`
	ID                           string      `json:"id"`
	Symbol                       string      `json:"symbol"`
	Name                         string      `json:"name"`
	Image                        string      `json:"image"`
	MarketCap                    int64       `json:"market_cap"`
	MarketCapRank                int         `json:"market_cap_rank"`
	TotalVolume                  int         `json:"total_volume"`
	MarketCapChange24H           int         `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H float64     `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64     `json:"circulating_supply"`
	TotalSupply                  float64     `json:"total_supply"`
	Ath                          float64     `json:"ath"`
	AthChangePercentage          float64     `json:"ath_change_percentage"`
	High24H                      float64     `json:"high_24h"`
	Low24H                       float64     `json:"low_24h"`
	PriceChange24H               float64     `json:"price_change_24h"`
	PriceChangePercentage24H     float64     `json:"price_change_percentage_24h"`
	CurrentPrice                 float64     `json:"current_price"`
	AthDate                      time.Time   `json:"ath_date"`
	Roi                          interface{} `json:"roi"`
	LastUpdated                  time.Time   `json:"last_updated"`
	SparklineIn7D                struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage1HInCurrency   float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24HInCurrency  float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7DInCurrency   float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14DInCurrency  float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30DInCurrency  float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200DInCurrency float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1YInCurrency   float64 `json:"price_change_percentage_1y_in_currency"`
}

type GeckoAdvanceCoins struct {
	GeckoAdvanceCoins []GeckoAdvanceCoin
}
