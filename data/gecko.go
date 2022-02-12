package data

type GeckoCoin struct {
	Timestamp string  `json:"@timestamp"`
	Coin      string  `json:"coin_keyword"`
	USD       float32 `json:"usd"`
	EUR       float32 `json:"eur"`
}
