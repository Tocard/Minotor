package ChiaDbPoolData

import "minotor/ChiaDbPool"

// GetAllFarmerPayment return all FarmerPayment in db
func GetAllFarmerPayment() ([]FarmerPayment, error) {
	var _FarmerPayment []FarmerPayment
	res := ChiaDbPool.Client.Model(FarmerPayment{}).Find(&_FarmerPayment)
	if res.Error != nil {
		return _FarmerPayment, res.Error
	}
	return _FarmerPayment, nil
}

type FarmerPayment struct {
	Model
	LauncherID              string `json:"launcher_id" gorm:"column:launcher_id"`
	Timestamp               int64  `json:"timestamp" gorm:"column:timestamp"`
	AmountToDistribute      int64  `json:"amount_to_distribute" gorm:"column:amount_to_distribute"`
	MojoPerPoint            int64  `json:"mojo_per_point" gorm:"column:mojo_per_point"`
	Points                  int64  `json:"points" gorm:"column:points"`
	PayoutInstructions      string `json:"payout_instructions" gorm:"column:payout_instructions"`
	Type                    string `json:"type" gorm:"column:type"`
	Status                  string `json:"status" gorm:"column:status"`
	Comment                 string `json:"comment" gorm:"column:comment"`
	TotalPoints             int64  `json:"total_points" gorm:"column:total_points"`
	AmountReallyDistributed int64  `json:"amount_really_distributed" gorm:"column:amount_really_distributed"`
	Amount                  int64  `json:"amount" gorm:"column:amount"`
	HeightBlockTransaction  int64  `json:"height_block_transaction" gorm:"column:height_block_transaction"`
	TxnHash                 string `json:"txn_hash" gorm:"column:txn_hash"`
	PayoutAddressXch        string `json:"payout_address_xch" gorm:"column:payout_address_xch"`
}

func (m *FarmerPayment) TableName() string {
	return "farmer_payment"
}

// GetAllFarmerPaymentFromTimestamp returns all BlockWin records with Timestamp greater than a specified value.
func GetAllFarmerPaymentFromTimestamp(Timestamp int64) ([]FarmerPayment, error) {
	var _FarmerPayment []FarmerPayment
	res := ChiaDbPool.Client.Model(FarmerPayment{}).Where("timestamp > ?", Timestamp).Find(&_FarmerPayment)
	if res.Error != nil {
		return _FarmerPayment, res.Error
	}
	return _FarmerPayment, nil
}
