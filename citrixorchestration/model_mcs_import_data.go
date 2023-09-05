/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the MCSImportData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MCSImportData{}

// MCSImportData struct for MCSImportData
type MCSImportData struct {
	IdentityPool *string `json:"IdentityPool,omitempty"`
	ProvisioningScheme *string `json:"ProvisioningScheme,omitempty"`
}

// NewMCSImportData instantiates a new MCSImportData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMCSImportData() *MCSImportData {
	this := MCSImportData{}
	return &this
}

// NewMCSImportDataWithDefaults instantiates a new MCSImportData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMCSImportDataWithDefaults() *MCSImportData {
	this := MCSImportData{}
	return &this
}

// GetIdentityPool returns the IdentityPool field value if set, zero value otherwise.
func (o *MCSImportData) GetIdentityPool() string {
	if o == nil || IsNil(o.IdentityPool) {
		var ret string
		return ret
	}
	return *o.IdentityPool
}

// GetIdentityPoolOk returns a tuple with the IdentityPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MCSImportData) GetIdentityPoolOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityPool) {
		return nil, false
	}
	return o.IdentityPool, true
}

// HasIdentityPool returns a boolean if a field has been set.
func (o *MCSImportData) HasIdentityPool() bool {
	if o != nil && !IsNil(o.IdentityPool) {
		return true
	}

	return false
}

// SetIdentityPool gets a reference to the given string and assigns it to the IdentityPool field.
func (o *MCSImportData) SetIdentityPool(v string) {
	o.IdentityPool = &v
}

// GetProvisioningScheme returns the ProvisioningScheme field value if set, zero value otherwise.
func (o *MCSImportData) GetProvisioningScheme() string {
	if o == nil || IsNil(o.ProvisioningScheme) {
		var ret string
		return ret
	}
	return *o.ProvisioningScheme
}

// GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MCSImportData) GetProvisioningSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvisioningScheme) {
		return nil, false
	}
	return o.ProvisioningScheme, true
}

// HasProvisioningScheme returns a boolean if a field has been set.
func (o *MCSImportData) HasProvisioningScheme() bool {
	if o != nil && !IsNil(o.ProvisioningScheme) {
		return true
	}

	return false
}

// SetProvisioningScheme gets a reference to the given string and assigns it to the ProvisioningScheme field.
func (o *MCSImportData) SetProvisioningScheme(v string) {
	o.ProvisioningScheme = &v
}

func (o MCSImportData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MCSImportData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityPool) {
		toSerialize["IdentityPool"] = o.IdentityPool
	}
	if !IsNil(o.ProvisioningScheme) {
		toSerialize["ProvisioningScheme"] = o.ProvisioningScheme
	}
	return toSerialize, nil
}

type NullableMCSImportData struct {
	value *MCSImportData
	isSet bool
}

func (v NullableMCSImportData) Get() *MCSImportData {
	return v.value
}

func (v *NullableMCSImportData) Set(val *MCSImportData) {
	v.value = val
	v.isSet = true
}

func (v NullableMCSImportData) IsSet() bool {
	return v.isSet
}

func (v *NullableMCSImportData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMCSImportData(val *MCSImportData) *NullableMCSImportData {
	return &NullableMCSImportData{value: val, isSet: true}
}

func (v NullableMCSImportData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMCSImportData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


