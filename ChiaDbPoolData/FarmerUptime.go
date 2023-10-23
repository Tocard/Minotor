package ChiaDbPoolData

import "minotor/ChiaDbPool"

// GetAllFarmerUptime return all FarmerUptime in db
func GetAllFarmerUptime() ([]FarmerUptime, error) {
	var _FarmerUptime []FarmerUptime
	res := ChiaDbPool.Client.Model(FarmerUptime{}).Find(&_FarmerUptime)
	if res.Error != nil {
		return _FarmerUptime, res.Error
	}
	return _FarmerUptime, nil
}

type FarmerUptime struct {
	Model
	LauncherID string `json:"launcher_id" gorm:"column:launcher_id"`
	Timestamp  int64  `json:"timestamp" gorm:"column:timestamp"`
}

func (m *FarmerUptime) TableName() string {
	return "farmer_uptime"
}

// GetAllFarmerUptimeFromTimestamp returns all BlockWin records with Timestamp greater than a specified value.
func GetAllFarmerUptimeFromTimestamp(Timestamp int64) ([]FarmerUptime, error) {
	var _FarmerUptime []FarmerUptime
	res := ChiaDbPool.Client.Model(FarmerUptime{}).Where("timestamp > ?", Timestamp).Find(&_FarmerUptime)
	if res.Error != nil {
		return _FarmerUptime, res.Error
	}
	return _FarmerUptime, nil
}
