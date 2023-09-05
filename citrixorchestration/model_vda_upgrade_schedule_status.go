/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// VdaUpgradeScheduleStatus VDA package upgrade schedule status.
type VdaUpgradeScheduleStatus string

// List of VdaUpgradeScheduleStatus
const (
	VDAUPGRADESCHEDULESTATUS_NOT_SCHEDULED VdaUpgradeScheduleStatus = "NotScheduled"
	VDAUPGRADESCHEDULESTATUS_UPGRADE_SCHEDULED VdaUpgradeScheduleStatus = "UpgradeScheduled"
	VDAUPGRADESCHEDULESTATUS_UPGRADE_IN_PROGRESS VdaUpgradeScheduleStatus = "UpgradeInProgress"
	VDAUPGRADESCHEDULESTATUS_UPGRADE_CANCELLED VdaUpgradeScheduleStatus = "UpgradeCancelled"
	VDAUPGRADESCHEDULESTATUS_UPGRADE_FAILED VdaUpgradeScheduleStatus = "UpgradeFailed"
	VDAUPGRADESCHEDULESTATUS_UPGRADE_SUCCESSFUL VdaUpgradeScheduleStatus = "UpgradeSuccessful"
)

// All allowed values of VdaUpgradeScheduleStatus enum
var AllowedVdaUpgradeScheduleStatusEnumValues = []VdaUpgradeScheduleStatus{
	"NotScheduled",
	"UpgradeScheduled",
	"UpgradeInProgress",
	"UpgradeCancelled",
	"UpgradeFailed",
	"UpgradeSuccessful",
}

func (v *VdaUpgradeScheduleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VdaUpgradeScheduleStatus(value)
	for _, existing := range AllowedVdaUpgradeScheduleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VdaUpgradeScheduleStatus", value)
}

// NewVdaUpgradeScheduleStatusFromValue returns a pointer to a valid VdaUpgradeScheduleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVdaUpgradeScheduleStatusFromValue(v string) (*VdaUpgradeScheduleStatus, error) {
	ev := VdaUpgradeScheduleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VdaUpgradeScheduleStatus: valid values are %v", v, AllowedVdaUpgradeScheduleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VdaUpgradeScheduleStatus) IsValid() bool {
	for _, existing := range AllowedVdaUpgradeScheduleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VdaUpgradeScheduleStatus value
func (v VdaUpgradeScheduleStatus) Ptr() *VdaUpgradeScheduleStatus {
	return &v
}

type NullableVdaUpgradeScheduleStatus struct {
	value *VdaUpgradeScheduleStatus
	isSet bool
}

func (v NullableVdaUpgradeScheduleStatus) Get() *VdaUpgradeScheduleStatus {
	return v.value
}

func (v *NullableVdaUpgradeScheduleStatus) Set(val *VdaUpgradeScheduleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVdaUpgradeScheduleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVdaUpgradeScheduleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdaUpgradeScheduleStatus(val *VdaUpgradeScheduleStatus) *NullableVdaUpgradeScheduleStatus {
	return &NullableVdaUpgradeScheduleStatus{value: val, isSet: true}
}

func (v NullableVdaUpgradeScheduleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdaUpgradeScheduleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

