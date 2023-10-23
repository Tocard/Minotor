package ChiaDbPoolData

import "minotor/ChiaDbPool"

type Partial struct {
	Model
	LauncherID  string `json:"launcher_id" gorm:"column:launcher_id"`
	Timestamp   int64  `json:"timestamp" gorm:"column:timestamp"`
	Difficulty  int64  `json:"difficulty" gorm:"column:difficulty"`
	HarvesterID string `json:"harvester_id" gorm:"column:harvester_id"`
	Error       int16  `json:"error" gorm:"column:error"`
}

func (m *Partial) TableName() string {
	return "partial"
}

// GetAllPartialFromTimestampPaginated returns a page of Partial records with Timestamp greater than a specified value.
func GetAllPartialFromTimestampPaginated(Timestamp int64, pageSize, pageNumber int) ([]Partial, error) {
	var _Partial []Partial
	offset := (pageNumber - 1) * pageSize
	res := ChiaDbPool.Client.Model(Partial{}).Where("timestamp > ?", Timestamp).Offset(offset).Limit(pageSize).Find(&_Partial)
	if res.Error != nil {
		return _Partial, res.Error
	}
	return _Partial, nil
}

// GetAllPartialFromTimestamp returns all Partial records with Timestamp greater than a specified value.
func GetAllPartialFromTimestamp(Timestamp int64) ([]Partial, error) {
	var _Partial []Partial
	res := ChiaDbPool.Client.Model(Partial{}).Where("timestamp > ?", Timestamp).Find(&_Partial)
	if res.Error != nil {
		return _Partial, res.Error
	}
	return _Partial, nil
}
