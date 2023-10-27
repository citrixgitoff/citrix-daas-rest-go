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

// checks if the SearchFilter3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchFilter3{}

// SearchFilter3 Search filter for settings.
type SearchFilter3 struct {
	Property SettingProperty `json:"Property"`
	// Value to match. For boolean properties this must be `\"true\"` or `\"false\"`. For string properties and enum values, this is case insensitive.
	Value string `json:"Value"`
	Operator SearchOperator `json:"Operator"`
}

// NewSearchFilter3 instantiates a new SearchFilter3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchFilter3(property SettingProperty, value string, operator SearchOperator) *SearchFilter3 {
	this := SearchFilter3{}
	this.Property = property
	this.Value = value
	this.Operator = operator
	return &this
}

// NewSearchFilter3WithDefaults instantiates a new SearchFilter3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchFilter3WithDefaults() *SearchFilter3 {
	this := SearchFilter3{}
	return &this
}

// GetProperty returns the Property field value
func (o *SearchFilter3) GetProperty() SettingProperty {
	if o == nil {
		var ret SettingProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *SearchFilter3) GetPropertyOk() (*SettingProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *SearchFilter3) SetProperty(v SettingProperty) {
	o.Property = v
}

// GetValue returns the Value field value
func (o *SearchFilter3) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SearchFilter3) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SearchFilter3) SetValue(v string) {
	o.Value = v
}

// GetOperator returns the Operator field value
func (o *SearchFilter3) GetOperator() SearchOperator {
	if o == nil {
		var ret SearchOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SearchFilter3) GetOperatorOk() (*SearchOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SearchFilter3) SetOperator(v SearchOperator) {
	o.Operator = v
}

func (o SearchFilter3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchFilter3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	toSerialize["Value"] = o.Value
	toSerialize["Operator"] = o.Operator
	return toSerialize, nil
}

type NullableSearchFilter3 struct {
	value *SearchFilter3
	isSet bool
}

func (v NullableSearchFilter3) Get() *SearchFilter3 {
	return v.value
}

func (v *NullableSearchFilter3) Set(val *SearchFilter3) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFilter3) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFilter3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFilter3(val *SearchFilter3) *NullableSearchFilter3 {
	return &NullableSearchFilter3{value: val, isSet: true}
}

func (v NullableSearchFilter3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFilter3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


