//  Copyright(C) 2022. Huawei Technologies Co.,Ltd.  All rights reserved.

// Package devmanager this for device driver manager
package devmanager

import (
	"errors"
	"fmt"

	"huawei.com/npu-exporter/devmanager/common"
	"huawei.com/npu-exporter/devmanager/dcmi"
	"huawei.com/npu-exporter/hwlog"
)

// DeviceInterface for common device interface
type DeviceInterface interface {
	Init() error
	ShutDown() error
	GetDeviceCount() (int32, error)
	GetCardList() (int32, []int32, error)
	GetDeviceList() (int32, []int32, error)
	GetDeviceHealth(logicID int32) (uint32, error)
	GetDeviceNetWorkHealth(logicID int32) (uint32, error)
	GetDeviceUtilizationRate(logicID int32, deviceType common.DeviceType) (uint32, error)
	GetDeviceTemperature(logicID int32) (int32, error)
	GetDeviceVoltage(logicID int32) (float32, error)
	GetDevicePowerInfo(logicID int32) (float32, error)
	GetDeviceFrequency(logicID int32, deviceType common.DeviceType) (int32, error)
	GetDeviceMemoryInfo(logicID int32) (*common.MemoryInfo, error)
	GetDeviceHbmInfo(logicID int32) (*common.HbmInfo, error)
	GetDeviceErrorCode(logicID int32) (int32, int64, error)
	GetChipInfo(logicID int32) (*common.ChipInfo, error)
	GetPhysicIDFromLogicID(logicID int32) (int32, error)
	GetLogicIDFromPhysicID(physicID int32) (int32, error)
	GetDeviceLogicID(cardID, deviceID int32) (int32, error)
	GetDeviceIPAddress(logicID int32) (string, error)
	CreateVirtualDevice(logicID int32, aiCore uint32) (uint32, error)
	GetVirtualDeviceInfo(logicID int32) (common.VirtualDevInfo, error)
	DestroyVirtualDevice(logicID int32, vDevID uint32) error
}

// DeviceManager common device manager for Ascend910/310P/310
type DeviceManager struct {
	// DcMgr for common dev manager
	DcMgr dcmi.DcDriverInterface
	// DevType the value is the same as the device type corresponding to the DcMgr variable.
	// Options: common.Ascend310,common.Ascend310P,common.Ascend910
	DevType string
}

// AutoInit auto detect npu chip type and return the corresponding processing object
func AutoInit(dType string) (DeviceManager, error) {
	chipInfo, err := getChipInfoForInit()
	if err != nil {
		return DeviceManager{}, fmt.Errorf("auto init failed, err: %s", err)
	}
	devManager := DeviceManager{}
	devType := common.GetDeviceTypeByChipName(chipInfo.Name)
	switch devType {
	case common.Ascend910:
		devManager.DcMgr = &A910Manager{}
	case common.Ascend310P:
		devManager.DcMgr = &A310PManager{}
	case common.Ascend310:
		devManager.DcMgr = &A310Manager{}
	default:
		return DeviceManager{}, fmt.Errorf("unsupport device type (%s)", devType)
	}
	if dType != "" && devType != dType {
		return DeviceManager{}, fmt.Errorf("the value of dType(%s) is inconsistent with the actual chip type(%s)",
			dType, devType)
	}
	devManager.DevType = devType
	if err = devManager.Init(); err != nil {
		return DeviceManager{}, fmt.Errorf("deviceManager init failed, err: %v", err)
	}
	return devManager, nil
}

func getChipInfoForInit() (common.ChipInfo, error) {
	dcMgr := dcmi.DcManager{}
	if err := dcMgr.DcInit(); err != nil {
		return common.ChipInfo{}, fmt.Errorf("dc init failed, err: %v", err)
	}
	defer func() {
		if err := dcMgr.DcShutDown(); err != nil {
			hwlog.RunLog.Error(err)
		}
	}()
	// get card list
	carNum, cardList, err := dcMgr.DcGetCardList()
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.ChipInfo{}, fmt.Errorf("get card list failed for init")
	}
	if carNum == 0 {
		return common.ChipInfo{}, fmt.Errorf("get chip info failed, no card found")
	}
	// get device in card, then get chip info by cardID and deviceID
	for _, cardID := range cardList {
		devNum, err := dcMgr.DcGetDeviceNumInCard(cardID)
		if err != nil || devNum == 0 {
			hwlog.RunLog.Debugf("get device num by cardID(%d) failed, error: %v", cardID, err)
			continue
		}
		for devID := int32(0); devID < devNum; devID++ {
			chipInfo, err := dcMgr.DcGetChipInfo(cardID, devID)
			if err != nil {
				hwlog.RunLog.Debugf("get chip info failed by cardID(%d), deviceID(%d), error: %v", cardID, devID,
					err)
				continue
			}
			if !common.IsValidChipInfo(chipInfo) {
				hwlog.RunLog.Debugf("invalid chip info by cardID(%d), deviceID(%d), error: %v", cardID, devID,
					err)
				continue
			}
			return *chipInfo, nil
		}
	}

	return common.ChipInfo{}, errors.New("cannot get valid chip info")
}

// Init load symbol and initialize dcmi
func (d *DeviceManager) Init() error {
	return d.DcMgr.DcInit()
}

// ShutDown clean the dynamically loaded resource
func (d *DeviceManager) ShutDown() error {
	return d.DcMgr.DcShutDown()
}

// GetDeviceCount get npu device count
func (d *DeviceManager) GetDeviceCount() (int32, error) {
	return d.DcMgr.DcGetDeviceCount()
}

// GetCardList  get all card list
func (d *DeviceManager) GetCardList() (int32, []int32, error) {
	return d.DcMgr.DcGetCardList()
}

// GetDeviceList get all device logicID list
func (d *DeviceManager) GetDeviceList() (int32, []int32, error) {
	return d.DcMgr.DcGetLogicIDList()
}

// GetDeviceHealth query npu device health status
func (d *DeviceManager) GetDeviceHealth(logicID int32) (uint32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get health code by logicID(%d)", logicID)
	}
	healthCode, err := d.DcMgr.DcGetDeviceHealth(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get health code by logicID(%d)", logicID)
	}

	return uint32(healthCode), nil
}

// GetDeviceNetWorkHealth query npu device network health status
func (d *DeviceManager) GetDeviceNetWorkHealth(logicID int32) (uint32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get network health code by logicID(%d)", logicID)
	}
	healthCode, err := d.DcMgr.DcGetDeviceNetWorkHealth(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get network health code by logicID(%d)", logicID)
	}

	return healthCode, nil
}

// GetDeviceUtilizationRate get npu device utilization
func (d *DeviceManager) GetDeviceUtilizationRate(logicID int32, deviceType common.DeviceType) (uint32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get utilization by logicID(%d)", logicID)
	}
	rate, err := d.DcMgr.DcGetDeviceUtilizationRate(cardID, deviceID, deviceType)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get utilization by logicID(%d)", logicID)
	}

	return uint32(rate), nil
}

// GetDeviceTemperature get npu device temperature
func (d *DeviceManager) GetDeviceTemperature(logicID int32) (int32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, fmt.Errorf("failed to get temperature by logicID(%d)", logicID)
	}
	temp, err := d.DcMgr.DcGetDeviceTemperature(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, fmt.Errorf("failed to get temperature by logicID(%d)", logicID)
	}

	return temp, nil
}

// GetDeviceVoltage get npu device voltage
func (d *DeviceManager) GetDeviceVoltage(logicID int32) (float32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get voltage by logicID(%d)", logicID)
	}
	voltage, err := d.DcMgr.DcGetDeviceVoltage(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get voltage by logicID(%d)", logicID)
	}

	return voltage, nil
}

// GetDevicePowerInfo get npu device power info
func (d *DeviceManager) GetDevicePowerInfo(logicID int32) (float32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get power by logicID(%d)", logicID)
	}
	power, err := d.DcMgr.DcGetDevicePowerInfo(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.UnRetError, fmt.Errorf("failed to get power by logicID(%d)", logicID)
	}

	return power, nil
}

// GetDeviceFrequency get npu device work frequency
func (d *DeviceManager) GetDeviceFrequency(logicID int32, deviceType common.DeviceType) (int32, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, fmt.Errorf("failed to get frequency by logicID(%d)", logicID)
	}
	frequency, err := d.DcMgr.DcGetDeviceFrequency(cardID, deviceID, deviceType)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, fmt.Errorf("failed to get frequency by logicID(%d)", logicID)
	}

	return frequency, nil
}

// GetDeviceMemoryInfo get npu memory information
func (d *DeviceManager) GetDeviceMemoryInfo(logicID int32) (*common.MemoryInfo, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get memory info by logicID(%d)", logicID)
	}
	memInfo, err := d.DcMgr.DcGetMemoryInfo(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get memory info by logicID(%d)", logicID)
	}

	return memInfo, nil
}

// GetDeviceHbmInfo get npu HBM module memory and frequency information
func (d *DeviceManager) GetDeviceHbmInfo(logicID int32) (*common.HbmInfo, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get hbm info by logicID(%d)", logicID)
	}
	hbmInfo, err := d.DcMgr.DcGetHbmInfo(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get hbm info by logicID(%d)", logicID)
	}

	return hbmInfo, nil
}

// GetDeviceErrorCode get npu device error code
func (d *DeviceManager) GetDeviceErrorCode(logicID int32) (int32, int64, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, common.RetError, fmt.Errorf("failed to get device error code by logicID(%d)",
			logicID)
	}
	errCount, errCode, err := d.DcMgr.DcGetDeviceErrorCode(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.RetError, common.RetError, fmt.Errorf("failed to get device error code by logicID(%d)",
			logicID)
	}

	return errCount, errCode, nil
}

// GetChipInfo get npu device error code
func (d *DeviceManager) GetChipInfo(logicID int32) (*common.ChipInfo, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get chip info code by logicID(%d)", logicID)
	}
	chipInfo, err := d.DcMgr.DcGetChipInfo(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return nil, fmt.Errorf("failed to get chip info code by logicID(%d)", logicID)
	}

	return chipInfo, nil
}

// GetPhysicIDFromLogicID get device physic id from logic id
func (d *DeviceManager) GetPhysicIDFromLogicID(logicID int32) (int32, error) {
	physicID, err := d.DcMgr.DcGetPhysicIDFromLogicID(logicID)
	if err != nil {
		return common.RetError, fmt.Errorf("failed to get physicID by logicID(%d)", logicID)
	}

	return physicID, nil
}

// GetLogicIDFromPhysicID get device logic id from physic id
func (d *DeviceManager) GetLogicIDFromPhysicID(physicID int32) (int32, error) {
	logicID, err := d.DcMgr.DcGetLogicIDFromPhysicID(physicID)
	if err != nil {
		return common.RetError, fmt.Errorf("failed to get logicID by physicID(%d)", physicID)
	}

	return logicID, nil
}

// GetDeviceLogicID get device logic id from card id and device id
func (d *DeviceManager) GetDeviceLogicID(cardID, deviceID int32) (int32, error) {
	return d.DcMgr.DcGetDeviceLogicID(cardID, deviceID)
}

// GetDeviceIPAddress get device ip address
func (d *DeviceManager) GetDeviceIPAddress(logicID int32) (string, error) {
	cardID, deviceID, err := d.DcMgr.DcGetCardIDDeviceID(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return "", fmt.Errorf("failed to get ip address by logicID(%d)", logicID)
	}
	ipAddr, err := d.DcMgr.DcGetDeviceIPAddress(cardID, deviceID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return "", fmt.Errorf("failed to get ip address by logicID(%d)", logicID)
	}

	return ipAddr, nil
}

// CreateVirtualDevice create virtual device
func (d *DeviceManager) CreateVirtualDevice(logicID int32, aiCore uint32) (uint32, error) {
	return d.DcMgr.DcCreateVDevice(logicID, aiCore)
}

// GetVirtualDeviceInfo get virtual device info
func (d *DeviceManager) GetVirtualDeviceInfo(logicID int32) (common.VirtualDevInfo, error) {
	dcmiVDevInfo, err := d.DcMgr.DcGetVDeviceInfo(logicID)
	if err != nil {
		hwlog.RunLog.Error(err)
		return common.VirtualDevInfo{}, fmt.Errorf("get virtual device info failed, error is: %v "+
			"and vdev num is: %d", err, int32(dcmiVDevInfo.VDevNum))
	}
	cgoVDevInfos := common.VirtualDevInfo{
		VDevNum:       dcmiVDevInfo.VDevNum,
		CoreNumUnused: uint32(dcmiVDevInfo.CoreNumUnused),
	}
	usedCoreCount := uint32(0)
	for i := uint32(0); i < cgoVDevInfos.VDevNum; i++ {
		usedCoreCount += uint32(dcmiVDevInfo.CoreNum[i])
		cgoVDevInfos.CgoDsmiSubVDevInfos = append(cgoVDevInfos.CgoDsmiSubVDevInfos, common.CgoDsmiSubVDevInfo{
			Status: dcmiVDevInfo.Status[i],
			VDevID: dcmiVDevInfo.VDevID[i],
			VfID:   dcmiVDevInfo.VfID[i],
			CID:    dcmiVDevInfo.CID[i],
			Spec: common.CgoDsmiVdevSpecInfo{
				CoreNum: fmt.Sprintf("%v", dcmiVDevInfo.CoreNum[i]),
			},
		})
	}
	cgoVDevInfos.CoreCount = cgoVDevInfos.CoreNumUnused + usedCoreCount
	return cgoVDevInfos, nil
}

// DestroyVirtualDevice destroy virtual device
func (d *DeviceManager) DestroyVirtualDevice(logicID int32, vDevID uint32) error {
	return d.DcMgr.DcDestroyVDevice(logicID, vDevID)
}
