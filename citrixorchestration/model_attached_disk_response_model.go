/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the AttachedDiskResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachedDiskResponseModel{}

// AttachedDiskResponseModel Attached disk.
type AttachedDiskResponseModel struct {
	// Name of the attached disk.
	Name NullableString `json:"Name,omitempty"`
	// Indicates whether the disk is the boot disk of the VM.
	IsBootDisk *bool `json:"IsBootDisk,omitempty"`
	// Size of the disk, in whole GB.
	SizeGB *int32 `json:"SizeGB,omitempty"`
	// Opaque identifier used by the underlying infrastructure to identify the disk.
	InfrastructureUid NullableString `json:"InfrastructureUid,omitempty"`
}

// NewAttachedDiskResponseModel instantiates a new AttachedDiskResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachedDiskResponseModel() *AttachedDiskResponseModel {
	this := AttachedDiskResponseModel{}
	return &this
}

// NewAttachedDiskResponseModelWithDefaults instantiates a new AttachedDiskResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachedDiskResponseModelWithDefaults() *AttachedDiskResponseModel {
	this := AttachedDiskResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachedDiskResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachedDiskResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AttachedDiskResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AttachedDiskResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AttachedDiskResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AttachedDiskResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetIsBootDisk returns the IsBootDisk field value if set, zero value otherwise.
func (o *AttachedDiskResponseModel) GetIsBootDisk() bool {
	if o == nil || IsNil(o.IsBootDisk) {
		var ret bool
		return ret
	}
	return *o.IsBootDisk
}

// GetIsBootDiskOk returns a tuple with the IsBootDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachedDiskResponseModel) GetIsBootDiskOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBootDisk) {
		return nil, false
	}
	return o.IsBootDisk, true
}

// HasIsBootDisk returns a boolean if a field has been set.
func (o *AttachedDiskResponseModel) HasIsBootDisk() bool {
	if o != nil && !IsNil(o.IsBootDisk) {
		return true
	}

	return false
}

// SetIsBootDisk gets a reference to the given bool and assigns it to the IsBootDisk field.
func (o *AttachedDiskResponseModel) SetIsBootDisk(v bool) {
	o.IsBootDisk = &v
}

// GetSizeGB returns the SizeGB field value if set, zero value otherwise.
func (o *AttachedDiskResponseModel) GetSizeGB() int32 {
	if o == nil || IsNil(o.SizeGB) {
		var ret int32
		return ret
	}
	return *o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachedDiskResponseModel) GetSizeGBOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeGB) {
		return nil, false
	}
	return o.SizeGB, true
}

// HasSizeGB returns a boolean if a field has been set.
func (o *AttachedDiskResponseModel) HasSizeGB() bool {
	if o != nil && !IsNil(o.SizeGB) {
		return true
	}

	return false
}

// SetSizeGB gets a reference to the given int32 and assigns it to the SizeGB field.
func (o *AttachedDiskResponseModel) SetSizeGB(v int32) {
	o.SizeGB = &v
}

// GetInfrastructureUid returns the InfrastructureUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachedDiskResponseModel) GetInfrastructureUid() string {
	if o == nil || IsNil(o.InfrastructureUid.Get()) {
		var ret string
		return ret
	}
	return *o.InfrastructureUid.Get()
}

// GetInfrastructureUidOk returns a tuple with the InfrastructureUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachedDiskResponseModel) GetInfrastructureUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfrastructureUid.Get(), o.InfrastructureUid.IsSet()
}

// HasInfrastructureUid returns a boolean if a field has been set.
func (o *AttachedDiskResponseModel) HasInfrastructureUid() bool {
	if o != nil && o.InfrastructureUid.IsSet() {
		return true
	}

	return false
}

// SetInfrastructureUid gets a reference to the given NullableString and assigns it to the InfrastructureUid field.
func (o *AttachedDiskResponseModel) SetInfrastructureUid(v string) {
	o.InfrastructureUid.Set(&v)
}
// SetInfrastructureUidNil sets the value for InfrastructureUid to be an explicit nil
func (o *AttachedDiskResponseModel) SetInfrastructureUidNil() {
	o.InfrastructureUid.Set(nil)
}

// UnsetInfrastructureUid ensures that no value is present for InfrastructureUid, not even an explicit nil
func (o *AttachedDiskResponseModel) UnsetInfrastructureUid() {
	o.InfrastructureUid.Unset()
}

func (o AttachedDiskResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachedDiskResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.IsBootDisk) {
		toSerialize["IsBootDisk"] = o.IsBootDisk
	}
	if !IsNil(o.SizeGB) {
		toSerialize["SizeGB"] = o.SizeGB
	}
	if o.InfrastructureUid.IsSet() {
		toSerialize["InfrastructureUid"] = o.InfrastructureUid.Get()
	}
	return toSerialize, nil
}

type NullableAttachedDiskResponseModel struct {
	value *AttachedDiskResponseModel
	isSet bool
}

func (v NullableAttachedDiskResponseModel) Get() *AttachedDiskResponseModel {
	return v.value
}

func (v *NullableAttachedDiskResponseModel) Set(val *AttachedDiskResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachedDiskResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachedDiskResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachedDiskResponseModel(val *AttachedDiskResponseModel) *NullableAttachedDiskResponseModel {
	return &NullableAttachedDiskResponseModel{value: val, isSet: true}
}

func (v NullableAttachedDiskResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachedDiskResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


