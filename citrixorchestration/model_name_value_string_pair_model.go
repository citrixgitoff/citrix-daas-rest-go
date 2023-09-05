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

// checks if the NameValueStringPairModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameValueStringPairModel{}

// NameValueStringPairModel A name/value pair, with the value of (JSON) type String.
type NameValueStringPairModel struct {
	// Name.
	Name *string `json:"Name,omitempty"`
	// Value.
	Value *string `json:"Value,omitempty"`
}

// NewNameValueStringPairModel instantiates a new NameValueStringPairModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameValueStringPairModel() *NameValueStringPairModel {
	this := NameValueStringPairModel{}
	return &this
}

// NewNameValueStringPairModelWithDefaults instantiates a new NameValueStringPairModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameValueStringPairModelWithDefaults() *NameValueStringPairModel {
	this := NameValueStringPairModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NameValueStringPairModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueStringPairModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NameValueStringPairModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NameValueStringPairModel) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NameValueStringPairModel) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueStringPairModel) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NameValueStringPairModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NameValueStringPairModel) SetValue(v string) {
	o.Value = &v
}

func (o NameValueStringPairModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameValueStringPairModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableNameValueStringPairModel struct {
	value *NameValueStringPairModel
	isSet bool
}

func (v NullableNameValueStringPairModel) Get() *NameValueStringPairModel {
	return v.value
}

func (v *NullableNameValueStringPairModel) Set(val *NameValueStringPairModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValueStringPairModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValueStringPairModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValueStringPairModel(val *NameValueStringPairModel) *NullableNameValueStringPairModel {
	return &NullableNameValueStringPairModel{value: val, isSet: true}
}

func (v NullableNameValueStringPairModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValueStringPairModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


