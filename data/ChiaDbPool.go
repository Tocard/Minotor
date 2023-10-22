package data

import (
	"minotor/ChiaDbPool"
)

type Model struct {
	FetchedAt string `json:"fetched_at"`
	FetchedBy string `json:"fetched_by"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (BlocksWin) TableName() string {
	return "blocks_win"
}

type BlocksWin struct {
	Model
	Timestamp      string `sql:"timestamp" json:"@timestamp"`
	BlockHeight    int    `json:"block_height"`
	LauncherId     string `json:"launcher_id"`
	Amount         int64  `json:"amount"`
	PuzzleHash     string `json:"puzzle_hash"`
	CoinHash       string `json:"coin_hash"`
	Status         string `json:"status"`
	Distributed    bool   `json:"distributed"`
	Announced      bool   `json:"announced"`
	ParentCoinInfo string `json:"parent_coin_info"`
}

// NewBlockWin returns a BlockWin pointer.
func NewBlockWin() *BlocksWin {
	return &BlocksWin{}
}

// GetAllBlockWins return all BlockWin in db
func GetAllBlockWins() ([]BlocksWin, error) {
	var _BlocksWin []BlocksWin
	res := ChiaDbPool.ChiaDbPool.Model(BlocksWin{}).Find(&_BlocksWin)
	if res.Error != nil {
		return _BlocksWin, res.Error
	}
	return _BlocksWin, nil
}
