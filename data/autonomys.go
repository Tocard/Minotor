package data

import (
	"encoding/json"
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

func (w Wallet) MarshalJSON() ([]byte, error) {
	type Alias Wallet
	return json.Marshal(&struct {
		Amount float64 `json:"amount"` // Serialize as float64 directly
		*Alias
	}{
		Amount: func() float64 {
			if w.Amount != nil {
				f, _ := w.Amount.Float64()
				return f
			}
			return 0
		}(),
		Alias: (*Alias)(&w),
	})
}
