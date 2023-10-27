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

// checks if the LosingSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LosingSetting{}

// LosingSetting Setting that is not applied.
type LosingSetting struct {
	PolicyIdentity *PolicyIdentity `json:"PolicyIdentity,omitempty"`
	// The setting that is not applied.
	SettingName NullableString `json:"SettingName,omitempty"`
}

// NewLosingSetting instantiates a new LosingSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLosingSetting() *LosingSetting {
	this := LosingSetting{}
	return &this
}

// NewLosingSettingWithDefaults instantiates a new LosingSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLosingSettingWithDefaults() *LosingSetting {
	this := LosingSetting{}
	return &this
}

// GetPolicyIdentity returns the PolicyIdentity field value if set, zero value otherwise.
func (o *LosingSetting) GetPolicyIdentity() PolicyIdentity {
	if o == nil || IsNil(o.PolicyIdentity) {
		var ret PolicyIdentity
		return ret
	}
	return *o.PolicyIdentity
}

// GetPolicyIdentityOk returns a tuple with the PolicyIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LosingSetting) GetPolicyIdentityOk() (*PolicyIdentity, bool) {
	if o == nil || IsNil(o.PolicyIdentity) {
		return nil, false
	}
	return o.PolicyIdentity, true
}

// HasPolicyIdentity returns a boolean if a field has been set.
func (o *LosingSetting) HasPolicyIdentity() bool {
	if o != nil && !IsNil(o.PolicyIdentity) {
		return true
	}

	return false
}

// SetPolicyIdentity gets a reference to the given PolicyIdentity and assigns it to the PolicyIdentity field.
func (o *LosingSetting) SetPolicyIdentity(v PolicyIdentity) {
	o.PolicyIdentity = &v
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingSetting) GetSettingName() string {
	if o == nil || IsNil(o.SettingName.Get()) {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingSetting) GetSettingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *LosingSetting) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *LosingSetting) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *LosingSetting) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *LosingSetting) UnsetSettingName() {
	o.SettingName.Unset()
}

func (o LosingSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LosingSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyIdentity) {
		toSerialize["PolicyIdentity"] = o.PolicyIdentity
	}
	if o.SettingName.IsSet() {
		toSerialize["SettingName"] = o.SettingName.Get()
	}
	return toSerialize, nil
}

type NullableLosingSetting struct {
	value *LosingSetting
	isSet bool
}

func (v NullableLosingSetting) Get() *LosingSetting {
	return v.value
}

func (v *NullableLosingSetting) Set(val *LosingSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableLosingSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableLosingSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLosingSetting(val *LosingSetting) *NullableLosingSetting {
	return &NullableLosingSetting{value: val, isSet: true}
}

func (v NullableLosingSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLosingSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


