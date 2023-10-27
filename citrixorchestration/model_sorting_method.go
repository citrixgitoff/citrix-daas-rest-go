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

// checks if the SortingMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortingMethod{}

// SortingMethod Policy set search result sort criteria.
type SortingMethod struct {
	Property SearchProperty `json:"Property"`
	SortDirection ListSortDirection `json:"SortDirection"`
}

// NewSortingMethod instantiates a new SortingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortingMethod(property SearchProperty, sortDirection ListSortDirection) *SortingMethod {
	this := SortingMethod{}
	this.Property = property
	this.SortDirection = sortDirection
	return &this
}

// NewSortingMethodWithDefaults instantiates a new SortingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortingMethodWithDefaults() *SortingMethod {
	this := SortingMethod{}
	return &this
}

// GetProperty returns the Property field value
func (o *SortingMethod) GetProperty() SearchProperty {
	if o == nil {
		var ret SearchProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *SortingMethod) GetPropertyOk() (*SearchProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *SortingMethod) SetProperty(v SearchProperty) {
	o.Property = v
}

// GetSortDirection returns the SortDirection field value
func (o *SortingMethod) GetSortDirection() ListSortDirection {
	if o == nil {
		var ret ListSortDirection
		return ret
	}

	return o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value
// and a boolean to check if the value has been set.
func (o *SortingMethod) GetSortDirectionOk() (*ListSortDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortDirection, true
}

// SetSortDirection sets field value
func (o *SortingMethod) SetSortDirection(v ListSortDirection) {
	o.SortDirection = v
}

func (o SortingMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortingMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	toSerialize["SortDirection"] = o.SortDirection
	return toSerialize, nil
}

type NullableSortingMethod struct {
	value *SortingMethod
	isSet bool
}

func (v NullableSortingMethod) Get() *SortingMethod {
	return v.value
}

func (v *NullableSortingMethod) Set(val *SortingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSortingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSortingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortingMethod(val *SortingMethod) *NullableSortingMethod {
	return &NullableSortingMethod{value: val, isSet: true}
}

func (v NullableSortingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


