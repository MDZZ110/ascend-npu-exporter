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

// ChipType chip type enum
type ChipType string

const (
	// HiAIMaxDeviceNum the max device num
	HiAIMaxDeviceNum = 64
	// HIAIMaxCardNum the max card num
	HIAIMaxCardNum = 8
	// Ascend910 Enum
	Ascend910 ChipType = "Ascend910"
	// Ascend710 chip type enum
	Ascend710 ChipType = "Ascend710"
	// Ascend310 chip type enum
	Ascend310 ChipType = "Ascend310"
	// DefaultErrorValue default error value
	DefaultErrorValue = -1
)