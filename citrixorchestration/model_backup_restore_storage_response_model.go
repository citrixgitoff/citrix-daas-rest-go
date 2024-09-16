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

// checks if the BackupRestoreStorageResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreStorageResponseModel{}

// BackupRestoreStorageResponseModel struct for BackupRestoreStorageResponseModel
type BackupRestoreStorageResponseModel struct {
	StorageInfo *BackupStorageInfo `json:"StorageInfo,omitempty"`
	// Maximum Backups
	MaximumBackups *int32 `json:"MaximumBackups,omitempty"`
	// Maximum Pinned Backups 
	MaximumPinnedBackups *int32 `json:"MaximumPinnedBackups,omitempty"`
	// Maximum Auto-Pinned Backups 
	MaximumAutoPinnedBackups *int32 `json:"MaximumAutoPinnedBackups,omitempty"`
}

// NewBackupRestoreStorageResponseModel instantiates a new BackupRestoreStorageResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreStorageResponseModel() *BackupRestoreStorageResponseModel {
	this := BackupRestoreStorageResponseModel{}
	return &this
}

// NewBackupRestoreStorageResponseModelWithDefaults instantiates a new BackupRestoreStorageResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreStorageResponseModelWithDefaults() *BackupRestoreStorageResponseModel {
	this := BackupRestoreStorageResponseModel{}
	return &this
}

// GetStorageInfo returns the StorageInfo field value if set, zero value otherwise.
func (o *BackupRestoreStorageResponseModel) GetStorageInfo() BackupStorageInfo {
	if o == nil || IsNil(o.StorageInfo) {
		var ret BackupStorageInfo
		return ret
	}
	return *o.StorageInfo
}

// GetStorageInfoOk returns a tuple with the StorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageResponseModel) GetStorageInfoOk() (*BackupStorageInfo, bool) {
	if o == nil || IsNil(o.StorageInfo) {
		return nil, false
	}
	return o.StorageInfo, true
}

// HasStorageInfo returns a boolean if a field has been set.
func (o *BackupRestoreStorageResponseModel) HasStorageInfo() bool {
	if o != nil && !IsNil(o.StorageInfo) {
		return true
	}

	return false
}

// SetStorageInfo gets a reference to the given BackupStorageInfo and assigns it to the StorageInfo field.
func (o *BackupRestoreStorageResponseModel) SetStorageInfo(v BackupStorageInfo) {
	o.StorageInfo = &v
}

// GetMaximumBackups returns the MaximumBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageResponseModel) GetMaximumBackups() int32 {
	if o == nil || IsNil(o.MaximumBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumBackups
}

// GetMaximumBackupsOk returns a tuple with the MaximumBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageResponseModel) GetMaximumBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumBackups) {
		return nil, false
	}
	return o.MaximumBackups, true
}

// HasMaximumBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageResponseModel) HasMaximumBackups() bool {
	if o != nil && !IsNil(o.MaximumBackups) {
		return true
	}

	return false
}

// SetMaximumBackups gets a reference to the given int32 and assigns it to the MaximumBackups field.
func (o *BackupRestoreStorageResponseModel) SetMaximumBackups(v int32) {
	o.MaximumBackups = &v
}

// GetMaximumPinnedBackups returns the MaximumPinnedBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageResponseModel) GetMaximumPinnedBackups() int32 {
	if o == nil || IsNil(o.MaximumPinnedBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumPinnedBackups
}

// GetMaximumPinnedBackupsOk returns a tuple with the MaximumPinnedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageResponseModel) GetMaximumPinnedBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumPinnedBackups) {
		return nil, false
	}
	return o.MaximumPinnedBackups, true
}

// HasMaximumPinnedBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageResponseModel) HasMaximumPinnedBackups() bool {
	if o != nil && !IsNil(o.MaximumPinnedBackups) {
		return true
	}

	return false
}

// SetMaximumPinnedBackups gets a reference to the given int32 and assigns it to the MaximumPinnedBackups field.
func (o *BackupRestoreStorageResponseModel) SetMaximumPinnedBackups(v int32) {
	o.MaximumPinnedBackups = &v
}

// GetMaximumAutoPinnedBackups returns the MaximumAutoPinnedBackups field value if set, zero value otherwise.
func (o *BackupRestoreStorageResponseModel) GetMaximumAutoPinnedBackups() int32 {
	if o == nil || IsNil(o.MaximumAutoPinnedBackups) {
		var ret int32
		return ret
	}
	return *o.MaximumAutoPinnedBackups
}

// GetMaximumAutoPinnedBackupsOk returns a tuple with the MaximumAutoPinnedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreStorageResponseModel) GetMaximumAutoPinnedBackupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAutoPinnedBackups) {
		return nil, false
	}
	return o.MaximumAutoPinnedBackups, true
}

// HasMaximumAutoPinnedBackups returns a boolean if a field has been set.
func (o *BackupRestoreStorageResponseModel) HasMaximumAutoPinnedBackups() bool {
	if o != nil && !IsNil(o.MaximumAutoPinnedBackups) {
		return true
	}

	return false
}

// SetMaximumAutoPinnedBackups gets a reference to the given int32 and assigns it to the MaximumAutoPinnedBackups field.
func (o *BackupRestoreStorageResponseModel) SetMaximumAutoPinnedBackups(v int32) {
	o.MaximumAutoPinnedBackups = &v
}

func (o BackupRestoreStorageResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreStorageResponseModel) ToMap() (map[string]interface{}, error) {
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

type NullableBackupRestoreStorageResponseModel struct {
	value *BackupRestoreStorageResponseModel
	isSet bool
}

func (v NullableBackupRestoreStorageResponseModel) Get() *BackupRestoreStorageResponseModel {
	return v.value
}

func (v *NullableBackupRestoreStorageResponseModel) Set(val *BackupRestoreStorageResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreStorageResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreStorageResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreStorageResponseModel(val *BackupRestoreStorageResponseModel) *NullableBackupRestoreStorageResponseModel {
	return &NullableBackupRestoreStorageResponseModel{value: val, isSet: true}
}

func (v NullableBackupRestoreStorageResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreStorageResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


