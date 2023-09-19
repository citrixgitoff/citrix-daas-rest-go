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

// checks if the NameValueIntPairModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameValueIntPairModel{}

// NameValueIntPairModel A name/value pair, with the value of (JSON) type Number.
type NameValueIntPairModel struct {
	// Name.
	Name NullableString `json:"Name,omitempty"`
	// Value.
	Value *int32 `json:"Value,omitempty"`
}

// NewNameValueIntPairModel instantiates a new NameValueIntPairModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameValueIntPairModel() *NameValueIntPairModel {
	this := NameValueIntPairModel{}
	return &this
}

// NewNameValueIntPairModelWithDefaults instantiates a new NameValueIntPairModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameValueIntPairModelWithDefaults() *NameValueIntPairModel {
	this := NameValueIntPairModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameValueIntPairModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameValueIntPairModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NameValueIntPairModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NameValueIntPairModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NameValueIntPairModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NameValueIntPairModel) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NameValueIntPairModel) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueIntPairModel) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NameValueIntPairModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *NameValueIntPairModel) SetValue(v int32) {
	o.Value = &v
}

func (o NameValueIntPairModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameValueIntPairModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableNameValueIntPairModel struct {
	value *NameValueIntPairModel
	isSet bool
}

func (v NullableNameValueIntPairModel) Get() *NameValueIntPairModel {
	return v.value
}

func (v *NullableNameValueIntPairModel) Set(val *NameValueIntPairModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValueIntPairModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValueIntPairModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValueIntPairModel(val *NameValueIntPairModel) *NullableNameValueIntPairModel {
	return &NullableNameValueIntPairModel{value: val, isSet: true}
}

func (v NullableNameValueIntPairModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValueIntPairModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


