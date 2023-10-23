package ChiaDbPoolData

import "minotor/ChiaDbPool"

// GetAllPoolNetspace return all PoolNetspace in db
func GetAllPoolNetspace() ([]PoolNetspace, error) {
	var _PoolNetspace []PoolNetspace
	res := ChiaDbPool.Client.Model(PoolNetspace{}).Find(&_PoolNetspace)
	if res.Error != nil {
		return _PoolNetspace, res.Error
	}
	return _PoolNetspace, nil
}

type PoolNetspace struct {
	Model
	Timestamp            int64 `json:"timestamp" gorm:"column:timestamp"`
	EstimatedNetspaceGib int64 `json:"estimated_netspace_gib" gorm:"column:estimated_netspace_gib"`
}

func (m *PoolNetspace) TableName() string {
	return "pool_netspace"
}

// GetAllPoolNetspaceFromTimestamp returns all BlockWin records with Timestamp greater than a specified value.
func GetAllPoolNetspaceFromTimestamp(Timestamp int64) ([]PoolNetspace, error) {
	var _PoolNetspace []PoolNetspace
	res := ChiaDbPool.Client.Model(PoolNetspace{}).Where("timestamp > ?", Timestamp).Find(&_PoolNetspace)
	if res.Error != nil {
		return _PoolNetspace, res.Error
	}
	return _PoolNetspace, nil
}
