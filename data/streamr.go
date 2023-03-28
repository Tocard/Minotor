package data

import "time"

type StreamR struct {
	Timestamp string `json:"@timestamp"`
	Status    string `json:"status"`
	Data      struct {
		Node struct {
			Address      string  `json:"address"`
			Status       bool    `json:"status"`
			Staked       int     `json:"staked"`
			ToBeReceived float64 `json:"toBeReceived"`
			Sent         float64 `json:"sent"`
			Rewards      float64 `json:"rewards"`
			FirstClaim   struct {
				Id        string    `json:"id"`
				ClaimTime time.Time `json:"claimTime"`
			} `json:"firstClaim"`
			LastClaim struct {
				Id        string    `json:"id"`
				ClaimTime time.Time `json:"claimTime"`
			} `json:"lastClaim"`
			ClaimCount      int     `json:"claimCount"`
			ClaimPercentage float64 `json:"claimPercentage"`
			Payouts         []struct {
				Value     string `json:"value"`
				Timestamp string `json:"timestamp"`
			} `json:"payouts"`
			ClaimedRewardCodes []struct {
				Id        string    `json:"id"`
				ClaimTime time.Time `json:"claimTime"`
			} `json:"claimedRewardCodes"`
			PolygonScanURL string `json:"polygonScanURL"`
			IdenticonURL   string `json:"identiconURL"`
		} `json:"node"`
	} `json:"data"`
}
