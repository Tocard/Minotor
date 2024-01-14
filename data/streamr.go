package data

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

type AllOperator struct {
	Data struct {
		Operator []Operator `json:"operators"`
	} `json:"data"`
}

type Operator struct {
	Stakes                      []Stakes         `json:"stakes,omitempty"`
	Delegations                 []Delegations    `json:"delegations,omitempty"`
	SlashingEvents              []SlashingEvents `json:"slashingEvents,omitempty"`
	Nodes                       []string         `json:"nodes"`
	QueueEntries                []interface{}    `json:"queueEntries"`
	DelegatorCount              int              `json:"delegatorCount"`
	ValueWithoutEarnings        string           `json:"valueWithoutEarnings"`
	Timestamp                   string           `json:"@timestamp"`
	Id                          string           `json:"id"`
	TotalStakeInSponsorshipsWei string           `json:"totalStakeInSponsorshipsWei"`
	DataTokenBalanceWei         string           `json:"dataTokenBalanceWei"`
	OperatorTokenTotalSupplyWei string           `json:"operatorTokenTotalSupplyWei"`
	MetadataJsonString          string           `json:"metadataJsonString,omitempty"`
	Metadata                    Metadata         `json:"metadataJsonStringSplited"`
	Owner                       string           `json:"owner"`
	CumulativeProfitsWei        string           `json:"cumulativeProfitsWei"`
	CumulativeOperatorsCutWei   string           `json:"cumulativeOperatorsCutWei"`
	OperatorsCutFraction        string           `json:"operatorsCutFraction"`
	Typename                    string           `json:"__typename"`
	TotalStakeInSponsorships    float64          `json:"totalStakeInSponsorships"`
	DataTokenBalance            float64          `json:"dataTokenBalance"`
	OperatorTokenTotalSupply    float64          `json:"operatorTokenTotalSupply"`
	CumulativeProfits           float64          `json:"cumulativeProfits²"`
	CumulativeOperatorsCut      float64          `json:"cumulativeOperatorsCut"`
	ValueWithoutEarningsReal    float64          `json:"valueWithoutEaarningsReal"`
	OperatorsCutFractionFloat   float64          `json:"operatorsCutFractionFloat"`
}

type Stakes struct {
	Operator struct {
		Id                 string   `json:"id"`
		MetadataJsonString string   `json:"metadataJsonString,omitempty"`
		Metadata           Metadata `json:"metadataJsonStringStakes"`
		Typename           string   `json:"__typename"`
	} `json:"operator"`
	Sponsorship struct {
		Id     string `json:"id"`
		Stream struct {
			Id       string `json:"id"`
			Metadata string `json:"metadata"`
			Typename string `json:"__typename"`
		} `json:"stream"`
		Metadata             string  `json:"metadata"`
		TotalPayoutWeiPerSec string  `json:"totalPayoutWeiPerSec"`
		IsRunning            bool    `json:"isRunning"`
		TotalPayoutPerSec    float64 `json:"totalPayoutPerSec"`
		Stakes               []struct {
			Operator struct {
				Id                 string `json:"id"`
				MetadataJsonString string `json:"metadataJsonString"`
				Typename           string `json:"__typename"`
			} `json:"operator"`
			AmountWei     string  `json:"amountWei"`
			EarningsWei   string  `json:"earningsWei"`
			LockedWei     string  `json:"lockedWei"`
			Typename      string  `json:"__typename"`
			JoinTimestamp int     `json:"joinTimestamp"`
			Amount        float64 `json:"amount"`
			Earnings      float64 `json:"earnings"`
			Locked        float64 `json:"locked"`
		} `json:"stakes"`
		OperatorCount               int         `json:"operatorCount"`
		MaxOperators                interface{} `json:"maxOperators"`
		TotalStakedWei              string      `json:"totalStakedWei"`
		RemainingWei                string      `json:"remainingWei"`
		ProjectedInsolvency         string      `json:"projectedInsolvency"`
		CumulativeSponsoring        string      `json:"cumulativeSponsoring"`
		MinimumStakingPeriodSeconds string      `json:"minimumStakingPeriodSeconds"`
		Creator                     string      `json:"creator"`
		SpotAPY                     string      `json:"spotAPY"`
		Typename                    string      `json:"__typename"`
		TotalStaked                 float64     `json:"totalStaked"`
		Remaining                   float64     `json:"remaining"`
	} `json:"sponsorship"`
	Timestamp      string  `json:"@timestamp"`
	OperatorSource string  `json:"OperatorSource"`
	AmountWei      string  `json:"amountWei"`
	EarningsWei    string  `json:"earningsWei"`
	LockedWei      string  `json:"lockedWei"`
	Typename       string  `json:"__typename"`
	JoinTimestamp  int     `json:"joinTimestamp"`
	Amount         float64 `json:"amount"`
	Earnings       float64 `json:"earnings"`
	Locked         float64 `json:"locked"`
}

func (s *Stakes) ConvertWeiFieldsToFloat(OperatorSource, clock string) {
	err := json.Unmarshal([]byte(s.Operator.MetadataJsonString), &s.Operator.Metadata)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Not possible to unmarhsal MetadaString for %s: error :%s"), s.Operator.MetadataJsonString, err.Error())
	}
	s.Operator.MetadataJsonString = ""
	s.Timestamp = clock
	s.OperatorSource = OperatorSource
	s.Amount = convertWeiToFloat(s.AmountWei)
	s.Earnings = convertWeiToFloat(s.EarningsWei)
	s.Locked = convertWeiToFloat(s.LockedWei)
	s.Sponsorship.TotalStaked = convertWeiToFloat(s.Sponsorship.TotalStakedWei)
	s.Sponsorship.TotalPayoutPerSec = convertWeiToFloat(s.Sponsorship.TotalPayoutWeiPerSec)
	for _, Stake := range s.Sponsorship.Stakes {
		Stake.Amount = convertWeiToFloat(Stake.AmountWei)
		Stake.Locked = convertWeiToFloat(Stake.LockedWei)
		Stake.Earnings = convertWeiToFloat(Stake.EarningsWei)

	}

}

type Delegations struct {
	Timestamp               string  `json:"@timestamp"`
	ValueDataWei            string  `json:"valueDataWei"`
	OperatorTokenBalanceWei string  `json:"operatorTokenBalanceWei"`
	Id                      string  `json:"id"`
	Typename                string  `json:"__typename"`
	OperatorSource          string  `json:"OperatorSource"`
	ValueData               float64 `json:"valueData"`
	OperatorTokenBalance    float64 `json:"operatorTokenBalance"`
	DelegatorPendingEarning float64 `json:"DelegatorPendingEarning"`
	Delegator               struct {
		Id       string `json:"id"`
		Typename string `json:"__typename"`
	} `json:"delegator"`
}

func (d *Delegations) ConvertWeiFieldsToFloat(OperatorName, clock string) {
	d.ValueData = convertWeiToFloat(d.ValueDataWei)
	d.Timestamp = clock
	d.OperatorTokenBalance = convertWeiToFloat(d.OperatorTokenBalanceWei)
	d.DelegatorPendingEarning = d.ValueData - d.OperatorTokenBalance
	d.OperatorSource = OperatorName
}

type SlashingEvents struct {
	OperatorSource string  `json:"OperatorSource"`
	Timestamp      string  `json:"@timestamp"`
	Date           string  `json:"date"`
	Amount         string  `json:"amount"`
	Typename       string  `json:"__typename"`
	AmountReal     float64 `json:"amountReal"`
	Sponsorship    struct {
		Id     string `json:"id"`
		Stream struct {
			Id       string `json:"id"`
			Typename string `json:"__typename"`
		} `json:"stream"`
		Typename string `json:"__typename"`
	} `json:"sponsorship"`
}

func (s *SlashingEvents) ConvertWeiFieldsToFloat(OperatorName, clock string) {
	s.Timestamp = clock
	s.AmountReal = convertWeiToFloat(s.Amount)
	s.OperatorSource = OperatorName
}

type Metadata struct {
	Name             string `json:"name"`
	RedundancyFactor int    `json:"redundancyFactor"`
	ImageIpfsCid     string `json:"imageIpfsCid"`
}

func (o *Operator) ConvertWeiFieldsToFloat(clock string) {
	o.Timestamp = clock
	o.ValueWithoutEarningsReal = convertWeiToFloat(o.ValueWithoutEarnings)
	o.TotalStakeInSponsorships = convertWeiToFloat(o.TotalStakeInSponsorshipsWei)
	o.DataTokenBalance = convertWeiToFloat(o.DataTokenBalanceWei)
	o.OperatorTokenTotalSupply = convertWeiToFloat(o.OperatorTokenTotalSupplyWei)
	o.CumulativeProfits = convertWeiToFloat(o.CumulativeProfitsWei)
	o.CumulativeOperatorsCut = convertWeiToFloat(o.CumulativeOperatorsCutWei)
	o.OperatorsCutFractionFloat = convertCutFraction(o.OperatorsCutFraction)

	err := json.Unmarshal([]byte(o.MetadataJsonString), &o.Metadata)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Not possible to unmarhsal MetadaString for %s: error :%s"), o.Owner, err.Error())
	}
	o.MetadataJsonString = ""
}

func (o *Operator) CleanupFields() {
	o.SlashingEvents = nil
	o.Stakes = nil
	o.Delegations = nil
}

func convertWeiToFloat(weiStr string) float64 {
	wei, ok := new(big.Int).SetString(weiStr, 10)
	if !ok {
		return 0.0
	}
	weiFloat := new(big.Float).SetInt(wei)
	weiFloat.Quo(weiFloat, big.NewFloat(1e18))
	result, _ := weiFloat.Float64()
	return result
}

func convertCutFraction(CutFraction string) float64 {
	wei, ok := new(big.Int).SetString(CutFraction, 10)
	if !ok {
		return 0.0
	}
	weiFloat := new(big.Float).SetInt(wei)
	weiFloat.Quo(weiFloat, big.NewFloat(1e16))
	result, _ := weiFloat.Float64()
	return result
}
