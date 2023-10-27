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

// checks if the EnrollMachineRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollMachineRequestModel{}

// EnrollMachineRequestModel Model for Enrolling the websocket machine
type EnrollMachineRequestModel struct {
	// The id for the vda's key pair registered with the MFA trust service.
	VdaInstanceId NullableString `json:"VdaInstanceId,omitempty"`
	// The instance name for the key. it is also used as the machine's name for non-domain joined vda.
	VdaInstanceName NullableString `json:"VdaInstanceName,omitempty"`
	// The vda's public service key to be registered with FMA Trust Service.
	ServicePublicKey NullableString `json:"ServicePublicKey,omitempty"`
	// Real sid of AD machine for domain joined; or virtual sid if non-domain joined this parameter will be used to create machine
	MachineSid NullableString `json:"MachineSid,omitempty"`
	MachineMetadata *EnrollMachineMetaDataModel `json:"MachineMetadata,omitempty"`
}

// NewEnrollMachineRequestModel instantiates a new EnrollMachineRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollMachineRequestModel() *EnrollMachineRequestModel {
	this := EnrollMachineRequestModel{}
	return &this
}

// NewEnrollMachineRequestModelWithDefaults instantiates a new EnrollMachineRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollMachineRequestModelWithDefaults() *EnrollMachineRequestModel {
	this := EnrollMachineRequestModel{}
	return &this
}

// GetVdaInstanceId returns the VdaInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnrollMachineRequestModel) GetVdaInstanceId() string {
	if o == nil || IsNil(o.VdaInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.VdaInstanceId.Get()
}

// GetVdaInstanceIdOk returns a tuple with the VdaInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnrollMachineRequestModel) GetVdaInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VdaInstanceId.Get(), o.VdaInstanceId.IsSet()
}

// HasVdaInstanceId returns a boolean if a field has been set.
func (o *EnrollMachineRequestModel) HasVdaInstanceId() bool {
	if o != nil && o.VdaInstanceId.IsSet() {
		return true
	}

	return false
}

// SetVdaInstanceId gets a reference to the given NullableString and assigns it to the VdaInstanceId field.
func (o *EnrollMachineRequestModel) SetVdaInstanceId(v string) {
	o.VdaInstanceId.Set(&v)
}
// SetVdaInstanceIdNil sets the value for VdaInstanceId to be an explicit nil
func (o *EnrollMachineRequestModel) SetVdaInstanceIdNil() {
	o.VdaInstanceId.Set(nil)
}

// UnsetVdaInstanceId ensures that no value is present for VdaInstanceId, not even an explicit nil
func (o *EnrollMachineRequestModel) UnsetVdaInstanceId() {
	o.VdaInstanceId.Unset()
}

// GetVdaInstanceName returns the VdaInstanceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnrollMachineRequestModel) GetVdaInstanceName() string {
	if o == nil || IsNil(o.VdaInstanceName.Get()) {
		var ret string
		return ret
	}
	return *o.VdaInstanceName.Get()
}

// GetVdaInstanceNameOk returns a tuple with the VdaInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnrollMachineRequestModel) GetVdaInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VdaInstanceName.Get(), o.VdaInstanceName.IsSet()
}

// HasVdaInstanceName returns a boolean if a field has been set.
func (o *EnrollMachineRequestModel) HasVdaInstanceName() bool {
	if o != nil && o.VdaInstanceName.IsSet() {
		return true
	}

	return false
}

// SetVdaInstanceName gets a reference to the given NullableString and assigns it to the VdaInstanceName field.
func (o *EnrollMachineRequestModel) SetVdaInstanceName(v string) {
	o.VdaInstanceName.Set(&v)
}
// SetVdaInstanceNameNil sets the value for VdaInstanceName to be an explicit nil
func (o *EnrollMachineRequestModel) SetVdaInstanceNameNil() {
	o.VdaInstanceName.Set(nil)
}

// UnsetVdaInstanceName ensures that no value is present for VdaInstanceName, not even an explicit nil
func (o *EnrollMachineRequestModel) UnsetVdaInstanceName() {
	o.VdaInstanceName.Unset()
}

// GetServicePublicKey returns the ServicePublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnrollMachineRequestModel) GetServicePublicKey() string {
	if o == nil || IsNil(o.ServicePublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.ServicePublicKey.Get()
}

// GetServicePublicKeyOk returns a tuple with the ServicePublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnrollMachineRequestModel) GetServicePublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServicePublicKey.Get(), o.ServicePublicKey.IsSet()
}

// HasServicePublicKey returns a boolean if a field has been set.
func (o *EnrollMachineRequestModel) HasServicePublicKey() bool {
	if o != nil && o.ServicePublicKey.IsSet() {
		return true
	}

	return false
}

// SetServicePublicKey gets a reference to the given NullableString and assigns it to the ServicePublicKey field.
func (o *EnrollMachineRequestModel) SetServicePublicKey(v string) {
	o.ServicePublicKey.Set(&v)
}
// SetServicePublicKeyNil sets the value for ServicePublicKey to be an explicit nil
func (o *EnrollMachineRequestModel) SetServicePublicKeyNil() {
	o.ServicePublicKey.Set(nil)
}

// UnsetServicePublicKey ensures that no value is present for ServicePublicKey, not even an explicit nil
func (o *EnrollMachineRequestModel) UnsetServicePublicKey() {
	o.ServicePublicKey.Unset()
}

// GetMachineSid returns the MachineSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnrollMachineRequestModel) GetMachineSid() string {
	if o == nil || IsNil(o.MachineSid.Get()) {
		var ret string
		return ret
	}
	return *o.MachineSid.Get()
}

// GetMachineSidOk returns a tuple with the MachineSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnrollMachineRequestModel) GetMachineSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineSid.Get(), o.MachineSid.IsSet()
}

// HasMachineSid returns a boolean if a field has been set.
func (o *EnrollMachineRequestModel) HasMachineSid() bool {
	if o != nil && o.MachineSid.IsSet() {
		return true
	}

	return false
}

// SetMachineSid gets a reference to the given NullableString and assigns it to the MachineSid field.
func (o *EnrollMachineRequestModel) SetMachineSid(v string) {
	o.MachineSid.Set(&v)
}
// SetMachineSidNil sets the value for MachineSid to be an explicit nil
func (o *EnrollMachineRequestModel) SetMachineSidNil() {
	o.MachineSid.Set(nil)
}

// UnsetMachineSid ensures that no value is present for MachineSid, not even an explicit nil
func (o *EnrollMachineRequestModel) UnsetMachineSid() {
	o.MachineSid.Unset()
}

// GetMachineMetadata returns the MachineMetadata field value if set, zero value otherwise.
func (o *EnrollMachineRequestModel) GetMachineMetadata() EnrollMachineMetaDataModel {
	if o == nil || IsNil(o.MachineMetadata) {
		var ret EnrollMachineMetaDataModel
		return ret
	}
	return *o.MachineMetadata
}

// GetMachineMetadataOk returns a tuple with the MachineMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollMachineRequestModel) GetMachineMetadataOk() (*EnrollMachineMetaDataModel, bool) {
	if o == nil || IsNil(o.MachineMetadata) {
		return nil, false
	}
	return o.MachineMetadata, true
}

// HasMachineMetadata returns a boolean if a field has been set.
func (o *EnrollMachineRequestModel) HasMachineMetadata() bool {
	if o != nil && !IsNil(o.MachineMetadata) {
		return true
	}

	return false
}

// SetMachineMetadata gets a reference to the given EnrollMachineMetaDataModel and assigns it to the MachineMetadata field.
func (o *EnrollMachineRequestModel) SetMachineMetadata(v EnrollMachineMetaDataModel) {
	o.MachineMetadata = &v
}

func (o EnrollMachineRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollMachineRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VdaInstanceId.IsSet() {
		toSerialize["VdaInstanceId"] = o.VdaInstanceId.Get()
	}
	if o.VdaInstanceName.IsSet() {
		toSerialize["VdaInstanceName"] = o.VdaInstanceName.Get()
	}
	if o.ServicePublicKey.IsSet() {
		toSerialize["ServicePublicKey"] = o.ServicePublicKey.Get()
	}
	if o.MachineSid.IsSet() {
		toSerialize["MachineSid"] = o.MachineSid.Get()
	}
	if !IsNil(o.MachineMetadata) {
		toSerialize["MachineMetadata"] = o.MachineMetadata
	}
	return toSerialize, nil
}

type NullableEnrollMachineRequestModel struct {
	value *EnrollMachineRequestModel
	isSet bool
}

func (v NullableEnrollMachineRequestModel) Get() *EnrollMachineRequestModel {
	return v.value
}

func (v *NullableEnrollMachineRequestModel) Set(val *EnrollMachineRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollMachineRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollMachineRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollMachineRequestModel(val *EnrollMachineRequestModel) *NullableEnrollMachineRequestModel {
	return &NullableEnrollMachineRequestModel{value: val, isSet: true}
}

func (v NullableEnrollMachineRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollMachineRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


