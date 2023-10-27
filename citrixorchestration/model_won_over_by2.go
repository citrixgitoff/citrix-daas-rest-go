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

// checks if the WonOverBy2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WonOverBy2{}

// WonOverBy2 The reason why a setting lost. Just the name of one of the winning policies that have this setting.
type WonOverBy2 struct {
	// Name of the setting.
	SettingName NullableString `json:"SettingName,omitempty"`
	// The winning policy that has this setting configured.
	WinningPolicy NullableString `json:"WinningPolicy,omitempty"`
}

// NewWonOverBy2 instantiates a new WonOverBy2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWonOverBy2() *WonOverBy2 {
	this := WonOverBy2{}
	return &this
}

// NewWonOverBy2WithDefaults instantiates a new WonOverBy2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWonOverBy2WithDefaults() *WonOverBy2 {
	this := WonOverBy2{}
	return &this
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WonOverBy2) GetSettingName() string {
	if o == nil || IsNil(o.SettingName.Get()) {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WonOverBy2) GetSettingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *WonOverBy2) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *WonOverBy2) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *WonOverBy2) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *WonOverBy2) UnsetSettingName() {
	o.SettingName.Unset()
}

// GetWinningPolicy returns the WinningPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WonOverBy2) GetWinningPolicy() string {
	if o == nil || IsNil(o.WinningPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.WinningPolicy.Get()
}

// GetWinningPolicyOk returns a tuple with the WinningPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WonOverBy2) GetWinningPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WinningPolicy.Get(), o.WinningPolicy.IsSet()
}

// HasWinningPolicy returns a boolean if a field has been set.
func (o *WonOverBy2) HasWinningPolicy() bool {
	if o != nil && o.WinningPolicy.IsSet() {
		return true
	}

	return false
}

// SetWinningPolicy gets a reference to the given NullableString and assigns it to the WinningPolicy field.
func (o *WonOverBy2) SetWinningPolicy(v string) {
	o.WinningPolicy.Set(&v)
}
// SetWinningPolicyNil sets the value for WinningPolicy to be an explicit nil
func (o *WonOverBy2) SetWinningPolicyNil() {
	o.WinningPolicy.Set(nil)
}

// UnsetWinningPolicy ensures that no value is present for WinningPolicy, not even an explicit nil
func (o *WonOverBy2) UnsetWinningPolicy() {
	o.WinningPolicy.Unset()
}

func (o WonOverBy2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WonOverBy2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SettingName.IsSet() {
		toSerialize["SettingName"] = o.SettingName.Get()
	}
	if o.WinningPolicy.IsSet() {
		toSerialize["WinningPolicy"] = o.WinningPolicy.Get()
	}
	return toSerialize, nil
}

type NullableWonOverBy2 struct {
	value *WonOverBy2
	isSet bool
}

func (v NullableWonOverBy2) Get() *WonOverBy2 {
	return v.value
}

func (v *NullableWonOverBy2) Set(val *WonOverBy2) {
	v.value = val
	v.isSet = true
}

func (v NullableWonOverBy2) IsSet() bool {
	return v.isSet
}

func (v *NullableWonOverBy2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWonOverBy2(val *WonOverBy2) *NullableWonOverBy2 {
	return &NullableWonOverBy2{value: val, isSet: true}
}

func (v NullableWonOverBy2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWonOverBy2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


