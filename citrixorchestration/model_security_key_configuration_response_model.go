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

// checks if the SecurityKeyConfigurationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityKeyConfigurationResponseModel{}

// SecurityKeyConfigurationResponseModel Response object for the site security key configuration.
type SecurityKeyConfigurationResponseModel struct {
	// The first security key.
	Key1 string `json:"Key1"`
	// The second security key.
	Key2 string `json:"Key2"`
	// Indicates whether to require a key to authenticate communications over the STA port.
	RequireKeyForSta bool `json:"RequireKeyForSta"`
	// Indicates whether to require a key to authenticate communications over the XML port.
	RequireKeyForXml bool `json:"RequireKeyForXml"`
}

// NewSecurityKeyConfigurationResponseModel instantiates a new SecurityKeyConfigurationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityKeyConfigurationResponseModel(key1 string, key2 string, requireKeyForSta bool, requireKeyForXml bool) *SecurityKeyConfigurationResponseModel {
	this := SecurityKeyConfigurationResponseModel{}
	this.Key1 = key1
	this.Key2 = key2
	this.RequireKeyForSta = requireKeyForSta
	this.RequireKeyForXml = requireKeyForXml
	return &this
}

// NewSecurityKeyConfigurationResponseModelWithDefaults instantiates a new SecurityKeyConfigurationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityKeyConfigurationResponseModelWithDefaults() *SecurityKeyConfigurationResponseModel {
	this := SecurityKeyConfigurationResponseModel{}
	return &this
}

// GetKey1 returns the Key1 field value
func (o *SecurityKeyConfigurationResponseModel) GetKey1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key1
}

// GetKey1Ok returns a tuple with the Key1 field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationResponseModel) GetKey1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key1, true
}

// SetKey1 sets field value
func (o *SecurityKeyConfigurationResponseModel) SetKey1(v string) {
	o.Key1 = v
}

// GetKey2 returns the Key2 field value
func (o *SecurityKeyConfigurationResponseModel) GetKey2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key2
}

// GetKey2Ok returns a tuple with the Key2 field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationResponseModel) GetKey2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key2, true
}

// SetKey2 sets field value
func (o *SecurityKeyConfigurationResponseModel) SetKey2(v string) {
	o.Key2 = v
}

// GetRequireKeyForSta returns the RequireKeyForSta field value
func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForSta() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireKeyForSta
}

// GetRequireKeyForStaOk returns a tuple with the RequireKeyForSta field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForStaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireKeyForSta, true
}

// SetRequireKeyForSta sets field value
func (o *SecurityKeyConfigurationResponseModel) SetRequireKeyForSta(v bool) {
	o.RequireKeyForSta = v
}

// GetRequireKeyForXml returns the RequireKeyForXml field value
func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForXml() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireKeyForXml
}

// GetRequireKeyForXmlOk returns a tuple with the RequireKeyForXml field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForXmlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireKeyForXml, true
}

// SetRequireKeyForXml sets field value
func (o *SecurityKeyConfigurationResponseModel) SetRequireKeyForXml(v bool) {
	o.RequireKeyForXml = v
}

func (o SecurityKeyConfigurationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityKeyConfigurationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Key1"] = o.Key1
	toSerialize["Key2"] = o.Key2
	toSerialize["RequireKeyForSta"] = o.RequireKeyForSta
	toSerialize["RequireKeyForXml"] = o.RequireKeyForXml
	return toSerialize, nil
}

type NullableSecurityKeyConfigurationResponseModel struct {
	value *SecurityKeyConfigurationResponseModel
	isSet bool
}

func (v NullableSecurityKeyConfigurationResponseModel) Get() *SecurityKeyConfigurationResponseModel {
	return v.value
}

func (v *NullableSecurityKeyConfigurationResponseModel) Set(val *SecurityKeyConfigurationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityKeyConfigurationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityKeyConfigurationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityKeyConfigurationResponseModel(val *SecurityKeyConfigurationResponseModel) *NullableSecurityKeyConfigurationResponseModel {
	return &NullableSecurityKeyConfigurationResponseModel{value: val, isSet: true}
}

func (v NullableSecurityKeyConfigurationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityKeyConfigurationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


