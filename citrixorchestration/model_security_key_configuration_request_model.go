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

// checks if the SecurityKeyConfigurationRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityKeyConfigurationRequestModel{}

// SecurityKeyConfigurationRequestModel Request object for modification of the site security key configuration.
type SecurityKeyConfigurationRequestModel struct {
	// The first security key.
	Key1 *string `json:"Key1,omitempty"`
	// The second security key.
	Key2 *string `json:"Key2,omitempty"`
	// Indicates whether to require a key to authenticate communications over the STA port.
	RequireKeyForSta *bool `json:"RequireKeyForSta,omitempty"`
	// Indicates whether to require a key to authenticate communications over the XML port.
	RequireKeyForXml *bool `json:"RequireKeyForXml,omitempty"`
}

// NewSecurityKeyConfigurationRequestModel instantiates a new SecurityKeyConfigurationRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityKeyConfigurationRequestModel() *SecurityKeyConfigurationRequestModel {
	this := SecurityKeyConfigurationRequestModel{}
	return &this
}

// NewSecurityKeyConfigurationRequestModelWithDefaults instantiates a new SecurityKeyConfigurationRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityKeyConfigurationRequestModelWithDefaults() *SecurityKeyConfigurationRequestModel {
	this := SecurityKeyConfigurationRequestModel{}
	return &this
}

// GetKey1 returns the Key1 field value if set, zero value otherwise.
func (o *SecurityKeyConfigurationRequestModel) GetKey1() string {
	if o == nil || IsNil(o.Key1) {
		var ret string
		return ret
	}
	return *o.Key1
}

// GetKey1Ok returns a tuple with the Key1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationRequestModel) GetKey1Ok() (*string, bool) {
	if o == nil || IsNil(o.Key1) {
		return nil, false
	}
	return o.Key1, true
}

// HasKey1 returns a boolean if a field has been set.
func (o *SecurityKeyConfigurationRequestModel) HasKey1() bool {
	if o != nil && !IsNil(o.Key1) {
		return true
	}

	return false
}

// SetKey1 gets a reference to the given string and assigns it to the Key1 field.
func (o *SecurityKeyConfigurationRequestModel) SetKey1(v string) {
	o.Key1 = &v
}

// GetKey2 returns the Key2 field value if set, zero value otherwise.
func (o *SecurityKeyConfigurationRequestModel) GetKey2() string {
	if o == nil || IsNil(o.Key2) {
		var ret string
		return ret
	}
	return *o.Key2
}

// GetKey2Ok returns a tuple with the Key2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationRequestModel) GetKey2Ok() (*string, bool) {
	if o == nil || IsNil(o.Key2) {
		return nil, false
	}
	return o.Key2, true
}

// HasKey2 returns a boolean if a field has been set.
func (o *SecurityKeyConfigurationRequestModel) HasKey2() bool {
	if o != nil && !IsNil(o.Key2) {
		return true
	}

	return false
}

// SetKey2 gets a reference to the given string and assigns it to the Key2 field.
func (o *SecurityKeyConfigurationRequestModel) SetKey2(v string) {
	o.Key2 = &v
}

// GetRequireKeyForSta returns the RequireKeyForSta field value if set, zero value otherwise.
func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForSta() bool {
	if o == nil || IsNil(o.RequireKeyForSta) {
		var ret bool
		return ret
	}
	return *o.RequireKeyForSta
}

// GetRequireKeyForStaOk returns a tuple with the RequireKeyForSta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForStaOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireKeyForSta) {
		return nil, false
	}
	return o.RequireKeyForSta, true
}

// HasRequireKeyForSta returns a boolean if a field has been set.
func (o *SecurityKeyConfigurationRequestModel) HasRequireKeyForSta() bool {
	if o != nil && !IsNil(o.RequireKeyForSta) {
		return true
	}

	return false
}

// SetRequireKeyForSta gets a reference to the given bool and assigns it to the RequireKeyForSta field.
func (o *SecurityKeyConfigurationRequestModel) SetRequireKeyForSta(v bool) {
	o.RequireKeyForSta = &v
}

// GetRequireKeyForXml returns the RequireKeyForXml field value if set, zero value otherwise.
func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForXml() bool {
	if o == nil || IsNil(o.RequireKeyForXml) {
		var ret bool
		return ret
	}
	return *o.RequireKeyForXml
}

// GetRequireKeyForXmlOk returns a tuple with the RequireKeyForXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForXmlOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireKeyForXml) {
		return nil, false
	}
	return o.RequireKeyForXml, true
}

// HasRequireKeyForXml returns a boolean if a field has been set.
func (o *SecurityKeyConfigurationRequestModel) HasRequireKeyForXml() bool {
	if o != nil && !IsNil(o.RequireKeyForXml) {
		return true
	}

	return false
}

// SetRequireKeyForXml gets a reference to the given bool and assigns it to the RequireKeyForXml field.
func (o *SecurityKeyConfigurationRequestModel) SetRequireKeyForXml(v bool) {
	o.RequireKeyForXml = &v
}

func (o SecurityKeyConfigurationRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityKeyConfigurationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key1) {
		toSerialize["Key1"] = o.Key1
	}
	if !IsNil(o.Key2) {
		toSerialize["Key2"] = o.Key2
	}
	if !IsNil(o.RequireKeyForSta) {
		toSerialize["RequireKeyForSta"] = o.RequireKeyForSta
	}
	if !IsNil(o.RequireKeyForXml) {
		toSerialize["RequireKeyForXml"] = o.RequireKeyForXml
	}
	return toSerialize, nil
}

type NullableSecurityKeyConfigurationRequestModel struct {
	value *SecurityKeyConfigurationRequestModel
	isSet bool
}

func (v NullableSecurityKeyConfigurationRequestModel) Get() *SecurityKeyConfigurationRequestModel {
	return v.value
}

func (v *NullableSecurityKeyConfigurationRequestModel) Set(val *SecurityKeyConfigurationRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityKeyConfigurationRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityKeyConfigurationRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityKeyConfigurationRequestModel(val *SecurityKeyConfigurationRequestModel) *NullableSecurityKeyConfigurationRequestModel {
	return &NullableSecurityKeyConfigurationRequestModel{value: val, isSet: true}
}

func (v NullableSecurityKeyConfigurationRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityKeyConfigurationRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


