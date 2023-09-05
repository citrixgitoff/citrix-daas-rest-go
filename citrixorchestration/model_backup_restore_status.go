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

// BackupRestoreStatus Backup / Restore Status
type BackupRestoreStatus string

// List of BackupRestoreStatus
const (
	BACKUPRESTORESTATUS_UNKNOWN BackupRestoreStatus = "Unknown"
	BACKUPRESTORESTATUS_PENDING BackupRestoreStatus = "Pending"
	BACKUPRESTORESTATUS_IN_PROGRESS BackupRestoreStatus = "InProgress"
	BACKUPRESTORESTATUS_COMPLETE BackupRestoreStatus = "Complete"
	BACKUPRESTORESTATUS_COMPLETE_WITH_WARNING BackupRestoreStatus = "CompleteWithWarning"
	BACKUPRESTORESTATUS_FAILED BackupRestoreStatus = "Failed"
	BACKUPRESTORESTATUS_CANCELED BackupRestoreStatus = "Canceled"
	BACKUPRESTORESTATUS_NON_TERMINATING_ERROR BackupRestoreStatus = "NonTerminatingError"
	BACKUPRESTORESTATUS_STALLED BackupRestoreStatus = "Stalled"
)

// All allowed values of BackupRestoreStatus enum
var AllowedBackupRestoreStatusEnumValues = []BackupRestoreStatus{
	"Unknown",
	"Pending",
	"InProgress",
	"Complete",
	"CompleteWithWarning",
	"Failed",
	"Canceled",
	"NonTerminatingError",
	"Stalled",
}

func (v *BackupRestoreStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupRestoreStatus(value)
	for _, existing := range AllowedBackupRestoreStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupRestoreStatus", value)
}

// NewBackupRestoreStatusFromValue returns a pointer to a valid BackupRestoreStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupRestoreStatusFromValue(v string) (*BackupRestoreStatus, error) {
	ev := BackupRestoreStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupRestoreStatus: valid values are %v", v, AllowedBackupRestoreStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupRestoreStatus) IsValid() bool {
	for _, existing := range AllowedBackupRestoreStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRestoreStatus value
func (v BackupRestoreStatus) Ptr() *BackupRestoreStatus {
	return &v
}

type NullableBackupRestoreStatus struct {
	value *BackupRestoreStatus
	isSet bool
}

func (v NullableBackupRestoreStatus) Get() *BackupRestoreStatus {
	return v.value
}

func (v *NullableBackupRestoreStatus) Set(val *BackupRestoreStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreStatus(val *BackupRestoreStatus) *NullableBackupRestoreStatus {
	return &NullableBackupRestoreStatus{value: val, isSet: true}
}

func (v NullableBackupRestoreStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

