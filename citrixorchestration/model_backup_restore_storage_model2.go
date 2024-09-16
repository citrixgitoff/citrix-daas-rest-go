/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the BackupRestoreStorageModel2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreStorageModel2{}

// BackupRestoreStorageModel2 struct for BackupRestoreStorageModel2
type BackupRestoreStorageModel2 struct {
	StorageInfo *BackupStorageInfo `json:"StorageInfo,omitempty"`
	// Maximum Backups
	MaximumBackups *int32 `json:"MaximumBackups,omitempty"`
	// Maximum Pinned Backups 
	MaximumPinnedBackups *int32 `json:"MaximumPinnedBackups,omitempty"`
	// Maximum Auto-Pinned Backups 
	MaximumAutoPinnedBackups *int32 `json:"MaximumAutoPinnedBackups,omitempty"`
}

// NewBackupRestoreStorageModel2 instantiates a new BackupRestoreStorageModel2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreStorageModel2() *BackupRestoreStorageModel2 {
	this := BackupRestoreStorageModel2{}
	return &this
}

// NewBackupRestoreStorageModel2WithDefaults instantiates a new BackupRestoreStorageModel2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreStorageModel2WithDefaults() *BackupRestoreStorageModel2 {
	this := BackupRestoreStorageModel2{}
	return &this
}

// GetStorageInfo returns the StorageInfo field value if set, zero value otherwise.
func (o *BackupRestoreStorageModel2) GetStorageInfo() BackupStorageInfo {
	if o == nil || IsNil(o.StorageInfo) {
		var ret BackupStorageInfo
		return ret
	}
	return *o.StorageInfo
}

// GetStorageInfoOk returns a tuple with the StorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageModel2) GetStorageInfoOk() (*BackupStorageInfo, bool) {
	if o == nil || IsNil(o.StorageInfo) {
		return nil, false
	}
	return o.StorageInfo, true
}

// HasStorageInfo returns a boolean if a field has been set.
func (o *BackupRestoreStorageModel2) HasStorageInfo() bool {
	if o != nil && !IsNil(o.StorageInfo) {
		return true
	}

	return false
}

// SetStorageInfo gets a reference to the given BackupStorageInfo and assigns it to the StorageInfo field.
func (o *BackupRestoreStorageModel2) SetStorageInfo(v BackupStorageInfo) {
	o.StorageInfo = &v
}

// GetMaximumBackups returns the MaximumBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageModel2) GetMaximumBackups() int32 {
	if o == nil || IsNil(o.MaximumBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumBackups
}

// GetMaximumBackupsOk returns a tuple with the MaximumBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageModel2) GetMaximumBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumBackups) {
		return nil, false
	}
	return o.MaximumBackups, true
}

// HasMaximumBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageModel2) HasMaximumBackups() bool {
	if o != nil && !IsNil(o.MaximumBackups) {
		return true
	}

	return false
}

// SetMaximumBackups gets a reference to the given int32 and assigns it to the MaximumBackups field.
func (o *BackupRestoreStorageModel2) SetMaximumBackups(v int32) {
	o.MaximumBackups = &v
}

// GetMaximumPinnedBackups returns the MaximumPinnedBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageModel2) GetMaximumPinnedBackups() int32 {
	if o == nil || IsNil(o.MaximumPinnedBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumPinnedBackups
}

// GetMaximumPinnedBackupsOk returns a tuple with the MaximumPinnedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageModel2) GetMaximumPinnedBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumPinnedBackups) {
		return nil, false
	}
	return o.MaximumPinnedBackups, true
}

// HasMaximumPinnedBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageModel2) HasMaximumPinnedBackups() bool {
	if o != nil && !IsNil(o.MaximumPinnedBackups) {
		return true
	}

	return false
}

// SetMaximumPinnedBackups gets a reference to the given int32 and assigns it to the MaximumPinnedBackups field.
func (o *BackupRestoreStorageModel2) SetMaximumPinnedBackups(v int32) {
	o.MaximumPinnedBackups = &v
}

// GetMaximumAutoPinnedBackups returns the MaximumAutoPinnedBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageModel2) GetMaximumAutoPinnedBackups() int32 {
	if o == nil || IsNil(o.MaximumAutoPinnedBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumAutoPinnedBackups
}

// GetMaximumAutoPinnedBackupsOk returns a tuple with the MaximumAutoPinnedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageModel2) GetMaximumAutoPinnedBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAutoPinnedBackups) {
		return nil, false
	}
	return o.MaximumAutoPinnedBackups, true
}

// HasMaximumAutoPinnedBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageModel2) HasMaximumAutoPinnedBackups() bool {
	if o != nil && !IsNil(o.MaximumAutoPinnedBackups) {
		return true
	}

	return false
}

// SetMaximumAutoPinnedBackups gets a reference to the given int32 and assigns it to the MaximumAutoPinnedBackups field.
func (o *BackupRestoreStorageModel2) SetMaximumAutoPinnedBackups(v int32) {
	o.MaximumAutoPinnedBackups = &v
}

func (o BackupRestoreStorageModel2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreStorageModel2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageInfo) {
		toSerialize["StorageInfo"] = o.StorageInfo
	}
	if !IsNil(o.MaximumBackups) {
		toSerialize["MaximumBackups"] = o.MaximumBackups
	}
	if !IsNil(o.MaximumPinnedBackups) {
		toSerialize["MaximumPinnedBackups"] = o.MaximumPinnedBackups
	}
	if !IsNil(o.MaximumAutoPinnedBackups) {
		toSerialize["MaximumAutoPinnedBackups"] = o.MaximumAutoPinnedBackups
	}
	return toSerialize, nil
}

type NullableBackupRestoreStorageModel2 struct {
	value *BackupRestoreStorageModel2
	isSet bool
}

func (v NullableBackupRestoreStorageModel2) Get() *BackupRestoreStorageModel2 {
	return v.value
}

func (v *NullableBackupRestoreStorageModel2) Set(val *BackupRestoreStorageModel2) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreStorageModel2) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreStorageModel2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreStorageModel2(val *BackupRestoreStorageModel2) *NullableBackupRestoreStorageModel2 {
	return &NullableBackupRestoreStorageModel2{value: val, isSet: true}
}

func (v NullableBackupRestoreStorageModel2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreStorageModel2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


