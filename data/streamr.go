package data

type AllOperator struct {
	Data struct {
		Operator []Operator `json:"operators"`
	} `json:"data"`
}

type Stakes struct {
	Operator struct {
		Id                 string `json:"id"`
		MetadataJsonString string `json:"metadataJsonString"`
		Typename           string `json:"__typename"`
	} `json:"operator"`
	AmountWei     float64 `json:"amountWei"`
	EarningsWei   float64 `json:"earningsWei"`
	LockedWei     float64 `json:"lockedWei"`
	JoinTimestamp int     `json:"joinTimestamp"`
	Typename      string  `json:"__typename"`
	Sponsorship   struct {
		Id     string `json:"id"`
		Stream struct {
			Id       string `json:"id"`
			Metadata string `json:"metadata"`
			Typename string `json:"__typename"`
		} `json:"stream"`
		Metadata             string  `json:"metadata"`
		IsRunning            bool    `json:"isRunning"`
		TotalPayoutWeiPerSec float64 `json:"totalPayoutWeiPerSec"`
		Stakes               []struct {
			Operator struct {
				Id                 string `json:"id"`
				MetadataJsonString string `json:"metadataJsonString"`
				Typename           string `json:"__typename"`
			} `json:"operator"`
			AmountWei     float64 `json:"amountWei"`
			EarningsWei   float64 `json:"earningsWei"`
			LockedWei     float64 `json:"lockedWei"`
			JoinTimestamp int     `json:"joinTimestamp"`
			Typename      string  `json:"__typename"`
		} `json:"stakes"`
		OperatorCount               int         `json:"operatorCount"`
		MaxOperators                interface{} `json:"maxOperators"`
		TotalStakedWei              float64     `json:"totalStakedWei"`
		RemainingWei                float64     `json:"remainingWei"`
		ProjectedInsolvency         string      `json:"projectedInsolvency"`
		CumulativeSponsoring        string      `json:"cumulativeSponsoring"`
		MinimumStakingPeriodSeconds string      `json:"minimumStakingPeriodSeconds"`
		Creator                     string      `json:"creator"`
		SpotAPY                     string      `json:"spotAPY"`
		Typename                    string      `json:"__typename"`
	} `json:"sponsorship"`
}

type Delegations struct {
	Delegator struct {
		Id       string `json:"id"`
		Typename string `json:"__typename"`
	} `json:"delegator"`
	ValueDataWei            float64 `json:"valueDataWei"`
	OperatorTokenBalanceWei float64 `json:"operatorTokenBalanceWei"`
	Id                      string  `json:"id"`
	Typename                string  `json:"__typename"`
}

type SlashingEvents struct {
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	Sponsorship struct {
		Id     string `json:"id"`
		Stream struct {
			Id       string `json:"id"`
			Typename string `json:"__typename"`
		} `json:"stream"`
		Typename string `json:"__typename"`
	} `json:"sponsorship"`
	Typename string `json:"__typename"`
}

type Operator struct {
	Timestamp                   string           `json:"@timestamp"`
	Id                          string           `json:"id"`
	Stakes                      []Stakes         `json:"stakes"`
	Delegations                 []Delegations    `json:"delegations"`
	SlashingEvents              []SlashingEvents `json:"slashingEvents"`
	QueueEntries                []interface{}    `json:"queueEntries"`
	DelegatorCount              int              `json:"delegatorCount"`
	ValueWithoutEarnings        float64          `json:"valueWithoutEarnings"`
	TotalStakeInSponsorshipsWei float64          `json:"totalStakeInSponsorshipsWei"`
	DataTokenBalanceWei         float64          `json:"dataTokenBalanceWei"`
	OperatorTokenTotalSupplyWei float64          `json:"operatorTokenTotalSupplyWei"`
	MetadataJsonString          string           `json:"metadataJsonString"`
	Owner                       string           `json:"owner"`
	Nodes                       []string         `json:"nodes"`
	CumulativeProfitsWei        float64          `json:"cumulativeProfitsWei"`
	CumulativeOperatorsCutWei   float64          `json:"cumulativeOperatorsCutWei"`
	OperatorsCutFraction        string           `json:"operatorsCutFraction"`
	Typename                    string           `json:"__typename"`
}
