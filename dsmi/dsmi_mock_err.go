//  Copyright(C) 2020. Huawei Technologies Co.,Ltd. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dsmi interface
package dsmi

import "C"
import "fmt"

// DeviceManagerMockErr  struct definition
type DeviceManagerMockErr struct {
}

var errorMsg = "mock error"

// NewDeviceManagerMockErr new DeviceManagerMockErr instance
func NewDeviceManagerMockErr() *DeviceManagerMockErr {
	return &DeviceManagerMockErr{}
}

// GetDeviceCount get ascend910 device quantity
func (d *DeviceManagerMockErr) GetDeviceCount() (int32, error) {

	return 0, fmt.Errorf(errorMsg)
}

// GetDeviceList  get device list
func (d *DeviceManagerMockErr) GetDeviceList(devices *[HiAIMaxDeviceNum]uint32) (int32, error) {
	return 0, fmt.Errorf(errorMsg)
}

// GetDeviceHealth get device health by id
func (d *DeviceManagerMockErr) GetDeviceHealth(logicID int32) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)

}

// GetDeviceUtilizationRate get device utils rate by id
// DeviceType  Ascend910 1,2,3,4,5,6,10  Ascend310 1,2,3,4,5
func (d *DeviceManagerMockErr) GetDeviceUtilizationRate(logicID int32, deviceType DeviceType) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)
}

// GetDeviceTemperature get the device temperature
func (d *DeviceManagerMockErr) GetDeviceTemperature(logicID int32) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)
}

// GetDeviceVoltage get the device voltage
func (d *DeviceManagerMockErr) GetDeviceVoltage(logicID int32) (float32, error) {
	return 0.00025, fmt.Errorf(errorMsg)
}

// GetDevicePower get the power info of the device, the result like : 8.2w
func (d *DeviceManagerMockErr) GetDevicePower(logicID int32) (float32, error) {
	return 0.0007, fmt.Errorf(errorMsg)

}

// GetDeviceFrequency get device frequency, unit MHz
// Ascend910 1,6,7,9
// Ascend310 1,2,3,4,5
// subType enum:  Memory,6HBM,AI_Core_Current_Fre,AI_Core_Normal_Fre(1,6,7,9)    see DeviceType
func (d *DeviceManagerMockErr) GetDeviceFrequency(logicID int32, subType DeviceType) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)
}

// GetDeviceMemoryInfo get memory information
func (d *DeviceManagerMockErr) GetDeviceMemoryInfo(logicID int32) (*MemoryInfo, error) {

	return nil, fmt.Errorf(errorMsg)
}

// GetDeviceHbmInfo get HBM information , only for Ascend910
func (d *DeviceManagerMockErr) GetDeviceHbmInfo(logicID int32) (*HbmInfo, error) {
	return nil, fmt.Errorf(errorMsg)
}

// GetDeviceErrCode get the error count and errorcode of the device
func (d *DeviceManagerMockErr) GetDeviceErrCode(logicID int32) (int32, int32, error) {
	return int32(0), int32(0), fmt.Errorf(errorMsg)
}

// GetChipInfo get chip info
func (d *DeviceManagerMockErr) GetChipInfo(logicID int32) (*ChipInfo, error) {
	return nil, fmt.Errorf(errorMsg)
}

// GetPhyIDFromLogicID get physic id form logic id
func (d *DeviceManagerMockErr) GetPhyIDFromLogicID(logicID uint32) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)
}

// GetLogicIDFromPhyID get logic id form physic id
func (d *DeviceManagerMockErr) GetLogicIDFromPhyID(phyID uint32) (int32, error) {
	return int32(0), fmt.Errorf(errorMsg)
}

// GetNPUMajorID query the MajorID of NPU devices
func (d *DeviceManagerMockErr) GetNPUMajorID() (string, error) {
	return "", fmt.Errorf(errorMsg)
}
