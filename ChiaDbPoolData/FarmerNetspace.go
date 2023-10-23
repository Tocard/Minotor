package ChiaDbPoolData

import "minotor/ChiaDbPool"

type FarmerNetspace struct {
	Model
	LauncherID           string `json:"launcher_id" gorm:"column:launcher_id"`
	Timestamp            int64  `json:"timestamp" gorm:"column:timestamp"`
	EstimatedNetspaceGib int64  `json:"estimated_netspace_gib" gorm:"column:estimated_netspace_gib"`
}

func (m *FarmerNetspace) TableName() string {
	return "farmer_netspace"
}

// GetAllFarmerNetspace return all BlockWin in db
func GetAllFarmerNetspace() ([]FarmerNetspace, error) {
	var _FarmerNetspace []FarmerNetspace
	res := ChiaDbPool.Client.Model(Farmer{}).Find(&_FarmerNetspace)
	if res.Error != nil {
		return _FarmerNetspace, res.Error
	}
	return _FarmerNetspace, nil
}

// GetAllFarmerNetspaceFromTimestamp returns all BlockWin records with Timestamp greater than a specified value.
func GetAllFarmerNetspaceFromTimestamp(Timestamp int64) ([]FarmerNetspace, error) {
	var _FarmerNetspace []FarmerNetspace
	res := ChiaDbPool.Client.Model(FarmerNetspace{}).Where("timestamp > ?", Timestamp).Find(&_FarmerNetspace)
	if res.Error != nil {
		return _FarmerNetspace, res.Error
	}
	return _FarmerNetspace, nil
}
