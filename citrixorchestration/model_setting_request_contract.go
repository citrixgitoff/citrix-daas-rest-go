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

// checks if the SettingRequestContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingRequestContract{}

// SettingRequestContract Setting instance contract
type SettingRequestContract struct {
	// Setting type. Is globally unique.
	SettingType NullableString `json:"SettingType,omitempty"`
	// The setting value. * For all setting types, excluding booleans, if null, use the default value. * For all setting types, including booleans, otherwise, use the specified value. * For boolean types, the value can be \"true\"/\"false\", \"1\"/\"0\", null will be accepted as using default value.
	SettingValue NullableString `json:"SettingValue,omitempty"`
	// The value in native C# type. If this value is null, use SettingValue, otherwise use this value.
	TypedValue map[string]interface{} `json:"TypedValue,omitempty"`
}

// NewSettingRequestContract instantiates a new SettingRequestContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingRequestContract() *SettingRequestContract {
	this := SettingRequestContract{}
	return &this
}

// NewSettingRequestContractWithDefaults instantiates a new SettingRequestContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingRequestContractWithDefaults() *SettingRequestContract {
	this := SettingRequestContract{}
	return &this
}

// GetSettingType returns the SettingType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingRequestContract) GetSettingType() string {
	if o == nil || IsNil(o.SettingType.Get()) {
		var ret string
		return ret
	}
	return *o.SettingType.Get()
}

// GetSettingTypeOk returns a tuple with the SettingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingRequestContract) GetSettingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingType.Get(), o.SettingType.IsSet()
}

// HasSettingType returns a boolean if a field has been set.
func (o *SettingRequestContract) HasSettingType() bool {
	if o != nil && o.SettingType.IsSet() {
		return true
	}

	return false
}

// SetSettingType gets a reference to the given NullableString and assigns it to the SettingType field.
func (o *SettingRequestContract) SetSettingType(v string) {
	o.SettingType.Set(&v)
}
// SetSettingTypeNil sets the value for SettingType to be an explicit nil
func (o *SettingRequestContract) SetSettingTypeNil() {
	o.SettingType.Set(nil)
}

// UnsetSettingType ensures that no value is present for SettingType, not even an explicit nil
func (o *SettingRequestContract) UnsetSettingType() {
	o.SettingType.Unset()
}

// GetSettingValue returns the SettingValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingRequestContract) GetSettingValue() string {
	if o == nil || IsNil(o.SettingValue.Get()) {
		var ret string
		return ret
	}
	return *o.SettingValue.Get()
}

// GetSettingValueOk returns a tuple with the SettingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingRequestContract) GetSettingValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingValue.Get(), o.SettingValue.IsSet()
}

// HasSettingValue returns a boolean if a field has been set.
func (o *SettingRequestContract) HasSettingValue() bool {
	if o != nil && o.SettingValue.IsSet() {
		return true
	}

	return false
}

// SetSettingValue gets a reference to the given NullableString and assigns it to the SettingValue field.
func (o *SettingRequestContract) SetSettingValue(v string) {
	o.SettingValue.Set(&v)
}
// SetSettingValueNil sets the value for SettingValue to be an explicit nil
func (o *SettingRequestContract) SetSettingValueNil() {
	o.SettingValue.Set(nil)
}

// UnsetSettingValue ensures that no value is present for SettingValue, not even an explicit nil
func (o *SettingRequestContract) UnsetSettingValue() {
	o.SettingValue.Unset()
}

// GetTypedValue returns the TypedValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingRequestContract) GetTypedValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TypedValue
}

// GetTypedValueOk returns a tuple with the TypedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingRequestContract) GetTypedValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TypedValue) {
		return map[string]interface{}{}, false
	}
	return o.TypedValue, true
}

// HasTypedValue returns a boolean if a field has been set.
func (o *SettingRequestContract) HasTypedValue() bool {
	if o != nil && IsNil(o.TypedValue) {
		return true
	}

	return false
}

// SetTypedValue gets a reference to the given map[string]interface{} and assigns it to the TypedValue field.
func (o *SettingRequestContract) SetTypedValue(v map[string]interface{}) {
	o.TypedValue = v
}

func (o SettingRequestContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingRequestContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SettingType.IsSet() {
		toSerialize["SettingType"] = o.SettingType.Get()
	}
	if o.SettingValue.IsSet() {
		toSerialize["SettingValue"] = o.SettingValue.Get()
	}
	if o.TypedValue != nil {
		toSerialize["TypedValue"] = o.TypedValue
	}
	return toSerialize, nil
}

type NullableSettingRequestContract struct {
	value *SettingRequestContract
	isSet bool
}

func (v NullableSettingRequestContract) Get() *SettingRequestContract {
	return v.value
}

func (v *NullableSettingRequestContract) Set(val *SettingRequestContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingRequestContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingRequestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingRequestContract(val *SettingRequestContract) *NullableSettingRequestContract {
	return &NullableSettingRequestContract{value: val, isSet: true}
}

func (v NullableSettingRequestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingRequestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


