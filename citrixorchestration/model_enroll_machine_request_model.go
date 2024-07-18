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
	VdaInstanceId string `json:"VdaInstanceId"`
	// The instance name for the key. it is also used as the machine's name for non-domain joined vda.
	VdaInstanceName string `json:"VdaInstanceName"`
	// The vda's public service key to be registered with FMA Trust Service.
	ServicePublicKey string `json:"ServicePublicKey"`
	// Real sid of AD machine for domain joined; or virtual sid if non-domain joined this parameter will be used to create machine
	MachineSid string `json:"MachineSid"`
	MachineMetadata *EnrollMachineMetaDataModel `json:"MachineMetadata,omitempty"`
}

// NewEnrollMachineRequestModel instantiates a new EnrollMachineRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollMachineRequestModel(vdaInstanceId string, vdaInstanceName string, servicePublicKey string, machineSid string) *EnrollMachineRequestModel {
	this := EnrollMachineRequestModel{}
	this.VdaInstanceId = vdaInstanceId
	this.VdaInstanceName = vdaInstanceName
	this.ServicePublicKey = servicePublicKey
	this.MachineSid = machineSid
	return &this
}

// NewEnrollMachineRequestModelWithDefaults instantiates a new EnrollMachineRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollMachineRequestModelWithDefaults() *EnrollMachineRequestModel {
	this := EnrollMachineRequestModel{}
	return &this
}

// GetVdaInstanceId returns the VdaInstanceId field value
func (o *EnrollMachineRequestModel) GetVdaInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VdaInstanceId
}

// GetVdaInstanceIdOk returns a tuple with the VdaInstanceId field value
// and a boolean to check if the value has been set.
func (o *EnrollMachineRequestModel) GetVdaInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdaInstanceId, true
}

// SetVdaInstanceId sets field value
func (o *EnrollMachineRequestModel) SetVdaInstanceId(v string) {
	o.VdaInstanceId = v
}

// GetVdaInstanceName returns the VdaInstanceName field value
func (o *EnrollMachineRequestModel) GetVdaInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VdaInstanceName
}

// GetVdaInstanceNameOk returns a tuple with the VdaInstanceName field value
// and a boolean to check if the value has been set.
func (o *EnrollMachineRequestModel) GetVdaInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdaInstanceName, true
}

// SetVdaInstanceName sets field value
func (o *EnrollMachineRequestModel) SetVdaInstanceName(v string) {
	o.VdaInstanceName = v
}

// GetServicePublicKey returns the ServicePublicKey field value
func (o *EnrollMachineRequestModel) GetServicePublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePublicKey
}

// GetServicePublicKeyOk returns a tuple with the ServicePublicKey field value
// and a boolean to check if the value has been set.
func (o *EnrollMachineRequestModel) GetServicePublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServicePublicKey, true
}

// SetServicePublicKey sets field value
func (o *EnrollMachineRequestModel) SetServicePublicKey(v string) {
	o.ServicePublicKey = v
}

// GetMachineSid returns the MachineSid field value
func (o *EnrollMachineRequestModel) GetMachineSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MachineSid
}

// GetMachineSidOk returns a tuple with the MachineSid field value
// and a boolean to check if the value has been set.
func (o *EnrollMachineRequestModel) GetMachineSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MachineSid, true
}

// SetMachineSid sets field value
func (o *EnrollMachineRequestModel) SetMachineSid(v string) {
	o.MachineSid = v
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
	toSerialize["VdaInstanceId"] = o.VdaInstanceId
	toSerialize["VdaInstanceName"] = o.VdaInstanceName
	toSerialize["ServicePublicKey"] = o.ServicePublicKey
	toSerialize["MachineSid"] = o.MachineSid
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


