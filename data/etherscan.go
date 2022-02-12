package data

import "github.com/nanmu42/etherscan-api"

type Balance struct {
	TwoMiners
	Balance *etherscan.BigInt `json:"wallet_balance"`
}
