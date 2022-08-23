//  Copyright(C) 2022. Huawei Technologies Co.,Ltd.  All rights reserved.

// Package common define common types
package common

// MemoryInfo memory information struct
type MemoryInfo struct {
	MemorySize      uint64 `json:"memory_size"`
	MemoryAvailable uint64 `json:"memory_available"`
	Frequency       uint32 `json:"memory_frequency"`
	Utilization     uint32 `json:"memory_utilization"`
}

// HbmInfo HBM info
type HbmInfo struct {
	MemorySize        uint64 `json:"memory_size"`        // HBM total size,KB
	Frequency         uint32 `json:"hbm_frequency"`      // HBM frequency MHz
	Usage             uint64 `json:"memory_usage"`       // HBM memory usage,KB
	Temp              int32  `json:"hbm_temperature"`    // HBM temperature
	BandWidthUtilRate uint32 `json:"hbm_bandwidth_util"` // HBM bandwidth utilization
}

// ChipInfo chip info
type ChipInfo struct {
	Type    string `json:"chip_type"`
	Name    string `json:"chip_name"`
	Version string `json:"chip_version"`
}

// CgoCreateVDevOut create virtual device output info
type CgoCreateVDevOut struct {
	VDevID     uint32
	PcieBus    uint32
	PcieDevice uint32
	PcieFunc   uint32
	VfgID      uint32
	Reserved   []uint8
}

// CgoCreateVDevRes create virtual device input info
type CgoCreateVDevRes struct {
	VDevID       uint32
	VfgID        uint32
	TemplateName string
	Reserved     []uint8
}

// CgoBaseResource base resource info
type CgoBaseResource struct {
	Token       uint64
	TokenMax    uint64
	TaskTimeout uint64
	VfgID       uint32
	VipMode     uint8
	Reserved    []uint8
}

// CgoComputingResource compute resource info
type CgoComputingResource struct {
	// accelator resource
	Aic     float32
	Aiv     float32
	Dsa     uint16
	Rtsq    uint16
	Acsq    uint16
	Cdqm    uint16
	CCore   uint16
	Ffts    uint16
	Sdma    uint16
	PcieDma uint16

	// memory resource, MB as unit
	MemorySize uint64

	// id resource
	EventID  uint32
	NotifyID uint32
	StreamID uint32
	ModelID  uint32

	// cpu resource
	TopicScheduleAicpu uint16
	HostCtrlCPU        uint16
	HostAicpu          uint16
	DeviceAicpu        uint16
	TopicCtrlCPUSlot   uint16

	Reserved []uint8
}

// CgoMediaResource media resource info
type CgoMediaResource struct {
	Jpegd    float32
	Jpege    float32
	Vpc      float32
	Vdec     float32
	Pngd     float32
	Venc     float32
	Reserved []uint8
}

// CgoVDevQueryInfo virtual resource special info
type CgoVDevQueryInfo struct {
	Name            string
	Status          uint32
	IsContainerUsed uint32
	Vfid            uint32
	VfgID           uint32
	ContainerID     uint64
	Base            CgoBaseResource
	Computing       CgoComputingResource
	Media           CgoMediaResource
}

// CgoVDevQueryStru virtual resource info
type CgoVDevQueryStru struct {
	VDevID    uint32
	QueryInfo CgoVDevQueryInfo
}

// CgoSocFreeResource soc free resource info
type CgoSocFreeResource struct {
	VfgNum    uint32
	VfgBitmap uint32
	Base      CgoBaseResource
	Computing CgoComputingResource
	Media     CgoMediaResource
}

// CgoSocTotalResource soc total resource info
type CgoSocTotalResource struct {
	VDevNum   uint32
	VDevID    []uint32
	VfgNum    uint32
	VfgBitmap uint32
	Base      CgoBaseResource
	Computing CgoComputingResource
	Media     CgoMediaResource
}

// VirtualDevInfo virtual device infos
type VirtualDevInfo struct {
	TotalResource CgoSocTotalResource
	FreeResource  CgoSocFreeResource
	VDevInfo      []CgoVDevQueryStru
}
