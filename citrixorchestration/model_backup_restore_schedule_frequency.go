/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// BackupRestoreScheduleFrequency Backup Restore Schedule Frequency
type BackupRestoreScheduleFrequency string

// List of BackupRestoreScheduleFrequency
const (
	BACKUPRESTORESCHEDULEFREQUENCY_DAILY BackupRestoreScheduleFrequency = "Daily"
	BACKUPRESTORESCHEDULEFREQUENCY_WEEKLY BackupRestoreScheduleFrequency = "Weekly"
	BACKUPRESTORESCHEDULEFREQUENCY_MONTHLY BackupRestoreScheduleFrequency = "Monthly"
	BACKUPRESTORESCHEDULEFREQUENCY_ONCE BackupRestoreScheduleFrequency = "Once"
	BACKUPRESTORESCHEDULEFREQUENCY_NOT_DEFINED BackupRestoreScheduleFrequency = "NotDefined"
)

// All allowed values of BackupRestoreScheduleFrequency enum
var AllowedBackupRestoreScheduleFrequencyEnumValues = []BackupRestoreScheduleFrequency{
	"Daily",
	"Weekly",
	"Monthly",
	"Once",
	"NotDefined",
}

func (v *BackupRestoreScheduleFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupRestoreScheduleFrequency(value)
	for _, existing := range AllowedBackupRestoreScheduleFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupRestoreScheduleFrequency", value)
}

// NewBackupRestoreScheduleFrequencyFromValue returns a pointer to a valid BackupRestoreScheduleFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupRestoreScheduleFrequencyFromValue(v string) (*BackupRestoreScheduleFrequency, error) {
	ev := BackupRestoreScheduleFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupRestoreScheduleFrequency: valid values are %v", v, AllowedBackupRestoreScheduleFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupRestoreScheduleFrequency) IsValid() bool {
	for _, existing := range AllowedBackupRestoreScheduleFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRestoreScheduleFrequency value
func (v BackupRestoreScheduleFrequency) Ptr() *BackupRestoreScheduleFrequency {
	return &v
}

type NullableBackupRestoreScheduleFrequency struct {
	value *BackupRestoreScheduleFrequency
	isSet bool
}

func (v NullableBackupRestoreScheduleFrequency) Get() *BackupRestoreScheduleFrequency {
	return v.value
}

func (v *NullableBackupRestoreScheduleFrequency) Set(val *BackupRestoreScheduleFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreScheduleFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreScheduleFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreScheduleFrequency(val *BackupRestoreScheduleFrequency) *NullableBackupRestoreScheduleFrequency {
	return &NullableBackupRestoreScheduleFrequency{value: val, isSet: true}
}

func (v NullableBackupRestoreScheduleFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreScheduleFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

