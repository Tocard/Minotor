package utils

import "2miner-monitoring/data"

func CreateNodes() {
	data.Cumulus = &data.NodesDefault{
		Name:        "Cumulus",
		DiskType:    "SSD/NVME",
		Cpu:         2,
		Threads:     4,
		GbRam:       8,
		Collateral:  1000,
		GbDiskSize:  220,
		MbDiskSpeed: 180,
		MbBandwidth: 25,
		MbEpsMin:    250,
		Reward:      7.5,
		LockTime:    0,
	}
	data.NodesList = append(data.NodesList, *data.Cumulus)
	data.Nimbus = &data.NodesDefault{
		Name:        "Nimbus",
		DiskType:    "SSD/NVME",
		Cpu:         4,
		Threads:     8,
		GbRam:       32,
		Collateral:  12500,
		GbDiskSize:  440,
		MbDiskSpeed: 180,
		MbBandwidth: 50,
		MbEpsMin:    640,
		Reward:      12.5,
		LockTime:    0,
	}
	data.NodesList = append(data.NodesList, *data.Nimbus)
	data.Stratus = &data.NodesDefault{
		Name:        "Stratus",
		DiskType:    "SSD/NVME",
		Cpu:         8,
		Threads:     16,
		GbRam:       64,
		Collateral:  40000,
		GbDiskSize:  880,
		MbDiskSpeed: 400,
		MbBandwidth: 100,
		MbEpsMin:    1520,
		Reward:      30,
		LockTime:    0,
	}
	data.NodesList = append(data.NodesList, *data.Stratus)
	//data.Titan = &data.NodesDefault{
	//	Name: "Titan",
	//	DiskType:    "SSD/NVME",
	//	Cpu:          8,
	//	Threads:     16,
	//	GbRam:       64,
	//	Collateral:  250,
	//	GbDiskSize:  880,
	//	MbDiskSpeed: 400,
	//	MbBandwidth: 100,
	//	MbEpsMin:    1520,
	//	Reward:      0,
	//	LockTime: 6,
	//}
	//data.NodesList = append(data.NodesList, *data.Titan)
	//data.Cirrus = &data.NodesDefault{
	//	Name: "Cirrus",
	//	DiskType:    "SSD/NVME",
	//	Cpu:          2,
	//	Threads:     4,
	//	GbRam:       8,
	//	Collateral:  5,
	//	GbDiskSize:  220,
	//	MbDiskSpeed: 180,
	//	MbBandwidth: 25,
	//	MbEpsMin:    280,
	//	Reward:      0,
	//	LockTime: 6,
	//}
	//data.NodesList = append(data.NodesList, *data.Cirrus)
}
