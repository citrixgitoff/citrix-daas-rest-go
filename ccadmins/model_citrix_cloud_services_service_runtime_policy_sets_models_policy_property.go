/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty{}

// CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty struct for CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty
type CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty struct {
	Value *bool `json:"value,omitempty"`
	CanChangeValue *bool `json:"canChangeValue,omitempty"`
}

// NewCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty() *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty {
	this := CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty{}
	return &this
}

// NewCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyPropertyWithDefaults instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyPropertyWithDefaults() *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty {
	this := CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) SetValue(v bool) {
	o.Value = &v
}

// GetCanChangeValue returns the CanChangeValue field value if set, zero value otherwise.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) GetCanChangeValue() bool {
	if o == nil || IsNil(o.CanChangeValue) {
		var ret bool
		return ret
	}
	return *o.CanChangeValue
}

// GetCanChangeValueOk returns a tuple with the CanChangeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) GetCanChangeValueOk() (*bool, bool) {
	if o == nil || IsNil(o.CanChangeValue) {
		return nil, false
	}
	return o.CanChangeValue, true
}

// HasCanChangeValue returns a boolean if a field has been set.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) HasCanChangeValue() bool {
	if o != nil && !IsNil(o.CanChangeValue) {
		return true
	}

	return false
}

// SetCanChangeValue gets a reference to the given bool and assigns it to the CanChangeValue field.
func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) SetCanChangeValue(v bool) {
	o.CanChangeValue = &v
}

func (o CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.CanChangeValue) {
		toSerialize["canChangeValue"] = o.CanChangeValue
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty struct {
	value *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty
	isSet bool
}

func (v NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) Get() *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty {
	return v.value
}

func (v *NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) Set(val *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty(val *CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) *NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty {
	return &NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


