package data

import (
	"math/big"
)

type Wallet struct {
	Address   string     `json:"address"`
	Amount    *big.Float `json:"amount"`
	Timestamp string     `json:"@timestamp,omitempty"`
}

func U128ToFloat128(u128 *big.Int) *big.Float {
	// Convert big.Int to big.Float
	return new(big.Float).SetInt(u128)
}
