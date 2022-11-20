package data

import (
	"fmt"
	"minotor/config"
	"strings"
	"time"
)

func GetTokenUrl(wallet string) (string, string) {
	tokenType := ""
	for _, coin := range config.Cfg.TokenWatcher {
		if strings.Contains(wallet, coin) {
			for _, url := range config.Cfg.UrlTokenWatcher {
				if coin != "cosmos" {
					tokenType = fmt.Sprintf("%s-lcd", coin)
				} else {
					tokenType = "cosmos"
				}
				if strings.Contains(url, tokenType) {
					return url, coin
				}
			}
		}
	}
	return "Not yet supported", ""
}

func GetFactor(coin string) float64 {
	switch coin {
	case "cosmos":
		return 0.000001
	case "osmo":
		return 0.000001
	case "evmos":
		return 0.000000000000000001
	}
	return 0
}

type CosmosDelegation struct {
	Height string `json:"height"`
	Result []struct {
		Height     string  `json:"height"`
		Timestamp  string  `json:"@timestamp"`
		GovCoin    string  `json:"gov_coin"`
		Wallet     string  `json:"wallet"`
		Factor     float64 `json:"factor"`
		Denom      string  `json:"denom"`
		Delegation struct {
			DelegatorAddress string  `json:"delegator_address"`
			ValidatorAddress string  `json:"validator_address"`
			Shares           float64 `json:"shares,string"`
		} `json:"delegation"`
		Balance struct {
			Denom  string  `json:"denom"`
			Amount float64 `json:"amount,string"`
		} `json:"balance"`
	} `json:"result"`
}

type CosmosBalance struct {
	Height string `json:"height"`
	Result []struct {
		Timestamp string  `json:"@timestamp"`
		Wallet    string  `json:"wallet"`
		Height    string  `json:"height"`
		GovCoin   string  `json:"gov_coin"`
		Factor    float64 `json:"factor"`
		Denom     string  `json:"denom"`
		Amount    float64 `json:"amount,string"`
	} `json:"result"`
}

type CosmosUnDelegation struct {
	Height string `json:"height"`
	Result []struct {
		DelegatorAddress string  `json:"delegator_address"`
		ValidatorAddress string  `json:"validator_address"`
		Coin             string  `json:"coin"`
		Wallet           string  `json:"wallet"`
		Timestamp        string  `json:"@timestamp"`
		Height           string  `json:"height"`
		GovCoin          string  `json:"gov_coin"`
		Factor           float64 `json:"factor"`
		Entries          []struct {
			CreationHeight string    `json:"creation_height"`
			CompletionTime time.Time `json:"completion_time"`
			InitialBalance float64   `json:"initial_balance,string"`
			Balance        float64   `json:"balance,string"`
		} `json:"entries"`
	} `json:"result"`
}

type estimatedReward struct {
	Timestamp string `json:"@timestamp"`
	Height    string `json:"height"`
	Coin      string `json:"coin"`
	Wallet    string `json:"wallet"`
	Result    struct {
		Rewards []struct {
			ValidatorAddress string `json:"validator_address"`
			Reward           []struct {
				Denom  string  `json:"denom"`
				Amount float64 `json:"amount,string"`
			} `json:"reward"`
		} `json:"rewards"`
		Total []struct {
			Denom  string  `json:"denom"`
			Amount float64 `json:"amount,string"`
		} `json:"total"`
	} `json:"result"`
}

type Price struct {
	Timestamp string `json:"@timestamp"`
	Aioz      struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"aioz"`
	Akt struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"akt"`
	Atolo struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int64         `json:"max_supply"`
			CirculatingSupply float64       `json:"circulating_supply"`
			TotalSupply       float64       `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"atolo"`
	Atom struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"atom"`
	Band struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"band"`
	Bcna struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"bcna"`
	Boot struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int64         `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"boot"`
	Btsg struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"btsg"`
	Bze struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply int64     `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  int       `json:"percent_change_1h"`
					PercentChange24H int       `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"bze"`
	Cheq struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"cheq"`
	Cmdx struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"cmdx"`
	Crbrus struct {
		Price  float64     `json:"price"`
		Volume float64     `json:"volume"`
		Cmc    interface{} `json:"cmc"`
	} `json:"crbrus"`
	Cro struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply int64     `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"cro"`
	Ctk struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"ctk"`
	Cudos struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"cudos"`
	Darc struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"darc"`
	Dec struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"dec"`
	Dsm struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"dsm"`
	Dvpn struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int64     `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"dvpn"`
	Echelon struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"echelon"`
	Erowan struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"erowan"`
	Evmos struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"evmos"`
	Fct struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"fct"`
	Fet struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"fet"`
	Grav struct {
		Price  float64     `json:"price"`
		Volume float64     `json:"volume"`
		Cmc    interface{} `json:"cmc"`
	} `json:"grav"`
	Hash struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int64         `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int64         `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  int       `json:"percent_change_1h"`
					PercentChange24H int       `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"hash"`
	Huahua struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"huahua"`
	Inj struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"inj"`
	Iov struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"iov"`
	Iris struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"iris"`
	Ixo struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"ixo"`
	Juno struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"juno"`
	Kava struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"kava"`
	Kuji struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       float64       `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"kuji"`
	Lamb struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"lamb"`
	Like struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"like"`
	Lum struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"lum"`
	Luna struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply float64       `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"luna"`
	Lunc struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"lunc"`
	Med struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int64     `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"med"`
	Mntl struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int64         `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       float64       `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"mntl"`
	Ngm struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"ngm"`
	Okt struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"okt"`
	Orai struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"orai"`
	Osmo struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"osmo"`
	Point struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"point"`
	Regen struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"regen"`
	Scrt struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"scrt"`
	Somm struct {
		Price  float64     `json:"price"`
		Volume float64     `json:"volume"`
		Cmc    interface{} `json:"cmc"`
	} `json:"somm"`
	Stars struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"stars"`
	Swth struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int64     `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"swth"`
	Tick struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"tick"`
	Umee struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int64     `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"umee"`
	Vdl struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int           `json:"id"`
			Name              string        `json:"name"`
			Symbol            string        `json:"symbol"`
			Slug              string        `json:"slug"`
			NumMarketPairs    int           `json:"num_market_pairs"`
			DateAdded         time.Time     `json:"date_added"`
			Tags              []interface{} `json:"tags"`
			MaxSupply         int           `json:"max_supply"`
			CirculatingSupply int           `json:"circulating_supply"`
			TotalSupply       int           `json:"total_supply"`
			IsActive          int           `json:"is_active"`
			CmcRank           int           `json:"cmc_rank"`
			IsFiat            int           `json:"is_fiat"`
			LastUpdated       time.Time     `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"vdl"`
	Xki struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply int       `json:"circulating_supply"`
			TotalSupply       int       `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        int       `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"xki"`
	Xprt struct {
		Price  float64 `json:"price"`
		Volume float64 `json:"volume"`
		Cmc    struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			Tags              []string  `json:"tags"`
			MaxSupply         int       `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			IsActive          int       `json:"is_active"`
			CmcRank           int       `json:"cmc_rank"`
			IsFiat            int       `json:"is_fiat"`
			LastUpdated       time.Time `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"cmc"`
	} `json:"xprt"`
}
