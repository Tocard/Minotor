package ChiaDbPoolData

import "minotor/ChiaDbPool"

type Farmer struct {
	Model
	LauncherID              string `json:"launcher_id" gorm:"column:launcher_id"`
	P2SingletonPuzzleHash   string `json:"p2_singleton_puzzle_hash" gorm:"column:p2_singleton_puzzle_hash"`
	DelayTime               int64  `json:"delay_time" gorm:"column:delay_time"`
	DelayPuzzleHash         string `json:"delay_puzzle_hash" gorm:"column:delay_puzzle_hash"`
	AuthenticationPublicKey string `json:"authentication_public_key" gorm:"column:authentication_public_key"`
	SingletonTip            string `json:"singleton_tip" gorm:"column:singleton_tip"`
	SingletonTipState       string `json:"singleton_tip_state" gorm:"column:singleton_tip_state"`
	Points                  int64  `json:"points" gorm:"column:points"`
	Difficulty              int64  `json:"difficulty" gorm:"column:difficulty"`
	PayoutInstructions      string `json:"payout_instructions" gorm:"column:payout_instructions"`
	IsPoolMember            int16  `json:"is_pool_member" gorm:"column:is_pool_member"`
	EstimatedNetspaceGib    int64  `json:"estimated_netspace_gib" gorm:"column:estimated_netspace_gib"`
	PayoutAddressXch        string `json:"payout_address_xch" gorm:"column:payout_address_xch"`
	UptimeTable             string `json:"uptime_table" gorm:"column:uptime_table"`
	VirtualBalance          int64  `json:"virtual_balance" gorm:"column:virtual_balance"`
	TotalSend               int64  `json:"total_send" gorm:"column:total_send"`
	Alias                   string `json:"alias" gorm:"column:alias"`
	IamHereTable            string `json:"iam_here_table" gorm:"column:iam_here_table"`
	OverridePayoutAddress   string `json:"override_payout_address" gorm:"column:override_payout_address"`
}

func (m *Farmer) TableName() string {
	return "farmer"
}

// GetAllFarmers return all BlockWin in db
func GetAllFarmers() ([]Farmer, error) {
	var _Farmer []Farmer
	res := ChiaDbPool.Client.Model(Farmer{}).Find(&_Farmer)
	if res.Error != nil {
		return _Farmer, res.Error
	}
	return _Farmer, nil
}
