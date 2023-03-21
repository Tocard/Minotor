package data

type WalletOverviewPure struct {
	Pseudo         string           `json:"pseudo"`
	WalletOverview []WalletOverview `json:"wallet_overview"`
}

type WalletOverview struct {
	Token     string  `json:"token"`
	Amount    float64 `json:"amount"`
	Value     float64 `json:"value"`
	Price     float64 `json:"price"`
	Pseudo    string  `json:"pseudo"`
	Timestamp string  `json:"@timestamp"`
}
