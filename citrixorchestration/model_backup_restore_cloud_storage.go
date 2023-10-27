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

// checks if the BackupRestoreCloudStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreCloudStorage{}

// BackupRestoreCloudStorage Backup/Restore cloud storage definition
type BackupRestoreCloudStorage struct {
	BlobStorageType *BackupRestoreBlobStorage `json:"BlobStorageType,omitempty"`
	// Info 1 for cloud storage
	Info1 NullableString `json:"Info1,omitempty"`
	// Info 2 for cloud storage
	Info2 NullableString `json:"Info2,omitempty"`
	// Info 2 for cloud storage
	Info3 NullableString `json:"Info3,omitempty"`
	// Info 4 for cloud storage
	Info4 NullableString `json:"Info4,omitempty"`
}

// NewBackupRestoreCloudStorage instantiates a new BackupRestoreCloudStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreCloudStorage() *BackupRestoreCloudStorage {
	this := BackupRestoreCloudStorage{}
	return &this
}

// NewBackupRestoreCloudStorageWithDefaults instantiates a new BackupRestoreCloudStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreCloudStorageWithDefaults() *BackupRestoreCloudStorage {
	this := BackupRestoreCloudStorage{}
	return &this
}

// GetBlobStorageType returns the BlobStorageType field value if set, zero value otherwise.
func (o *BackupRestoreCloudStorage) GetBlobStorageType() BackupRestoreBlobStorage {
	if o == nil || IsNil(o.BlobStorageType) {
		var ret BackupRestoreBlobStorage
		return ret
	}
	return *o.BlobStorageType
}

// GetBlobStorageTypeOk returns a tuple with the BlobStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreCloudStorage) GetBlobStorageTypeOk() (*BackupRestoreBlobStorage, bool) {
	if o == nil || IsNil(o.BlobStorageType) {
		return nil, false
	}
	return o.BlobStorageType, true
}

// HasBlobStorageType returns a boolean if a field has been set.
func (o *BackupRestoreCloudStorage) HasBlobStorageType() bool {
	if o != nil && !IsNil(o.BlobStorageType) {
		return true
	}

	return false
}

// SetBlobStorageType gets a reference to the given BackupRestoreBlobStorage and assigns it to the BlobStorageType field.
func (o *BackupRestoreCloudStorage) SetBlobStorageType(v BackupRestoreBlobStorage) {
	o.BlobStorageType = &v
}

// GetInfo1 returns the Info1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreCloudStorage) GetInfo1() string {
	if o == nil || IsNil(o.Info1.Get()) {
		var ret string
		return ret
	}
	return *o.Info1.Get()
}

// GetInfo1Ok returns a tuple with the Info1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreCloudStorage) GetInfo1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info1.Get(), o.Info1.IsSet()
}

// HasInfo1 returns a boolean if a field has been set.
func (o *BackupRestoreCloudStorage) HasInfo1() bool {
	if o != nil && o.Info1.IsSet() {
		return true
	}

	return false
}

// SetInfo1 gets a reference to the given NullableString and assigns it to the Info1 field.
func (o *BackupRestoreCloudStorage) SetInfo1(v string) {
	o.Info1.Set(&v)
}
// SetInfo1Nil sets the value for Info1 to be an explicit nil
func (o *BackupRestoreCloudStorage) SetInfo1Nil() {
	o.Info1.Set(nil)
}

// UnsetInfo1 ensures that no value is present for Info1, not even an explicit nil
func (o *BackupRestoreCloudStorage) UnsetInfo1() {
	o.Info1.Unset()
}

// GetInfo2 returns the Info2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreCloudStorage) GetInfo2() string {
	if o == nil || IsNil(o.Info2.Get()) {
		var ret string
		return ret
	}
	return *o.Info2.Get()
}

// GetInfo2Ok returns a tuple with the Info2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreCloudStorage) GetInfo2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info2.Get(), o.Info2.IsSet()
}

// HasInfo2 returns a boolean if a field has been set.
func (o *BackupRestoreCloudStorage) HasInfo2() bool {
	if o != nil && o.Info2.IsSet() {
		return true
	}

	return false
}

// SetInfo2 gets a reference to the given NullableString and assigns it to the Info2 field.
func (o *BackupRestoreCloudStorage) SetInfo2(v string) {
	o.Info2.Set(&v)
}
// SetInfo2Nil sets the value for Info2 to be an explicit nil
func (o *BackupRestoreCloudStorage) SetInfo2Nil() {
	o.Info2.Set(nil)
}

// UnsetInfo2 ensures that no value is present for Info2, not even an explicit nil
func (o *BackupRestoreCloudStorage) UnsetInfo2() {
	o.Info2.Unset()
}

// GetInfo3 returns the Info3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreCloudStorage) GetInfo3() string {
	if o == nil || IsNil(o.Info3.Get()) {
		var ret string
		return ret
	}
	return *o.Info3.Get()
}

// GetInfo3Ok returns a tuple with the Info3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreCloudStorage) GetInfo3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info3.Get(), o.Info3.IsSet()
}

// HasInfo3 returns a boolean if a field has been set.
func (o *BackupRestoreCloudStorage) HasInfo3() bool {
	if o != nil && o.Info3.IsSet() {
		return true
	}

	return false
}

// SetInfo3 gets a reference to the given NullableString and assigns it to the Info3 field.
func (o *BackupRestoreCloudStorage) SetInfo3(v string) {
	o.Info3.Set(&v)
}
// SetInfo3Nil sets the value for Info3 to be an explicit nil
func (o *BackupRestoreCloudStorage) SetInfo3Nil() {
	o.Info3.Set(nil)
}

// UnsetInfo3 ensures that no value is present for Info3, not even an explicit nil
func (o *BackupRestoreCloudStorage) UnsetInfo3() {
	o.Info3.Unset()
}

// GetInfo4 returns the Info4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreCloudStorage) GetInfo4() string {
	if o == nil || IsNil(o.Info4.Get()) {
		var ret string
		return ret
	}
	return *o.Info4.Get()
}

// GetInfo4Ok returns a tuple with the Info4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreCloudStorage) GetInfo4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info4.Get(), o.Info4.IsSet()
}

// HasInfo4 returns a boolean if a field has been set.
func (o *BackupRestoreCloudStorage) HasInfo4() bool {
	if o != nil && o.Info4.IsSet() {
		return true
	}

	return false
}

// SetInfo4 gets a reference to the given NullableString and assigns it to the Info4 field.
func (o *BackupRestoreCloudStorage) SetInfo4(v string) {
	o.Info4.Set(&v)
}
// SetInfo4Nil sets the value for Info4 to be an explicit nil
func (o *BackupRestoreCloudStorage) SetInfo4Nil() {
	o.Info4.Set(nil)
}

// UnsetInfo4 ensures that no value is present for Info4, not even an explicit nil
func (o *BackupRestoreCloudStorage) UnsetInfo4() {
	o.Info4.Unset()
}

func (o BackupRestoreCloudStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreCloudStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobStorageType) {
		toSerialize["BlobStorageType"] = o.BlobStorageType
	}
	if o.Info1.IsSet() {
		toSerialize["Info1"] = o.Info1.Get()
	}
	if o.Info2.IsSet() {
		toSerialize["Info2"] = o.Info2.Get()
	}
	if o.Info3.IsSet() {
		toSerialize["Info3"] = o.Info3.Get()
	}
	if o.Info4.IsSet() {
		toSerialize["Info4"] = o.Info4.Get()
	}
	return toSerialize, nil
}

type NullableBackupRestoreCloudStorage struct {
	value *BackupRestoreCloudStorage
	isSet bool
}

func (v NullableBackupRestoreCloudStorage) Get() *BackupRestoreCloudStorage {
	return v.value
}

func (v *NullableBackupRestoreCloudStorage) Set(val *BackupRestoreCloudStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreCloudStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreCloudStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreCloudStorage(val *BackupRestoreCloudStorage) *NullableBackupRestoreCloudStorage {
	return &NullableBackupRestoreCloudStorage{value: val, isSet: true}
}

func (v NullableBackupRestoreCloudStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreCloudStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


