package data

var Cumulus *NodesDefault
var Nimbus *NodesDefault
var Stratus *NodesDefault
var Titan *NodesDefault
var Cirrus *NodesDefault

var FluxRewardPerBlock = 75.0

var NodesList []NodesDefault

type NodesDefault struct {
	Name        string  `json:"name,omitempty"`
	DiskType    string  `json:"disk_type,omitempty"`
	Cpu         int     `json:"cpu,omitempty"`
	Threads     int     `json:"threads,omitempty"`
	GbRam       int     `json:"gb_ram,omitempty"`
	Collateral  int     `json:"collateral,omitempty"`
	GbDiskSize  int     `json:"gb_disk_size,omitempty"`
	MbDiskSpeed int     `json:"mb_disk_speed,omitempty"`
	MbBandwidth int     `json:"mb_bandwidth,omitempty"`
	MbEpsMin    int     `json:"mb_eps_min,omitempty"`
	LockTime    int     `json:"lock_time,omitempty"`
	Reward      float64 `json:"reward,omitempty"`
}
