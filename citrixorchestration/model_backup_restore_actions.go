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

// BackupRestoreActions 
type BackupRestoreActions string

// List of BackupRestoreActions
const (
	BACKUPRESTOREACTIONS_BACKUP BackupRestoreActions = "Backup"
	BACKUPRESTOREACTIONS_RESTORE BackupRestoreActions = "Restore"
	BACKUPRESTOREACTIONS_PIN_BACKUP BackupRestoreActions = "PinBackup"
	BACKUPRESTOREACTIONS_UNPIN_BACKUP BackupRestoreActions = "UnpinBackup"
	BACKUPRESTOREACTIONS_DOWNLOAD_BACKUP BackupRestoreActions = "DownloadBackup"
	BACKUPRESTOREACTIONS_UPLOAD_BACKUP BackupRestoreActions = "UploadBackup"
	BACKUPRESTOREACTIONS_DELETE_BACKUP BackupRestoreActions = "DeleteBackup"
	BACKUPRESTOREACTIONS_DELETE_SCHEDULE BackupRestoreActions = "DeleteSchedule"
	BACKUPRESTOREACTIONS_CREATE_SCHEDULE BackupRestoreActions = "CreateSchedule"
	BACKUPRESTOREACTIONS_EDIT_SCHEDULE BackupRestoreActions = "EditSchedule"
	BACKUPRESTOREACTIONS_GET_MEMBERS_IN_FILE BackupRestoreActions = "GetMembersInFile"
	BACKUPRESTOREACTIONS_UNDEFINED BackupRestoreActions = "Undefined"
)

// All allowed values of BackupRestoreActions enum
var AllowedBackupRestoreActionsEnumValues = []BackupRestoreActions{
	"Backup",
	"Restore",
	"PinBackup",
	"UnpinBackup",
	"DownloadBackup",
	"UploadBackup",
	"DeleteBackup",
	"DeleteSchedule",
	"CreateSchedule",
	"EditSchedule",
	"GetMembersInFile",
	"Undefined",
}

func (v *BackupRestoreActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupRestoreActions(value)
	for _, existing := range AllowedBackupRestoreActionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupRestoreActions", value)
}

// NewBackupRestoreActionsFromValue returns a pointer to a valid BackupRestoreActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupRestoreActionsFromValue(v string) (*BackupRestoreActions, error) {
	ev := BackupRestoreActions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupRestoreActions: valid values are %v", v, AllowedBackupRestoreActionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupRestoreActions) IsValid() bool {
	for _, existing := range AllowedBackupRestoreActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRestoreActions value
func (v BackupRestoreActions) Ptr() *BackupRestoreActions {
	return &v
}

type NullableBackupRestoreActions struct {
	value *BackupRestoreActions
	isSet bool
}

func (v NullableBackupRestoreActions) Get() *BackupRestoreActions {
	return v.value
}

func (v *NullableBackupRestoreActions) Set(val *BackupRestoreActions) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreActions) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreActions(val *BackupRestoreActions) *NullableBackupRestoreActions {
	return &NullableBackupRestoreActions{value: val, isSet: true}
}

func (v NullableBackupRestoreActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

