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

// checks if the DeliveryGroupTestResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryGroupTestResponseModel{}

// DeliveryGroupTestResponseModel Response object for a delivery group test result.
type DeliveryGroupTestResponseModel struct {
	DeliveryGroup RefResponseModel `json:"DeliveryGroup"`
	// The number of tests that passed.
	NumPassed int32 `json:"NumPassed"`
	// The number of warnings that were found.
	NumWarnings int32 `json:"NumWarnings"`
	// The number of tests that failed.
	NumFailures int32 `json:"NumFailures"`
}

// NewDeliveryGroupTestResponseModel instantiates a new DeliveryGroupTestResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryGroupTestResponseModel(deliveryGroup RefResponseModel, numPassed int32, numWarnings int32, numFailures int32) *DeliveryGroupTestResponseModel {
	this := DeliveryGroupTestResponseModel{}
	this.DeliveryGroup = deliveryGroup
	this.NumPassed = numPassed
	this.NumWarnings = numWarnings
	this.NumFailures = numFailures
	return &this
}

// NewDeliveryGroupTestResponseModelWithDefaults instantiates a new DeliveryGroupTestResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryGroupTestResponseModelWithDefaults() *DeliveryGroupTestResponseModel {
	this := DeliveryGroupTestResponseModel{}
	return &this
}

// GetDeliveryGroup returns the DeliveryGroup field value
func (o *DeliveryGroupTestResponseModel) GetDeliveryGroup() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.DeliveryGroup
}

// GetDeliveryGroupOk returns a tuple with the DeliveryGroup field value
// and a boolean to check if the value has been set.
func (o *DeliveryGroupTestResponseModel) GetDeliveryGroupOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryGroup, true
}

// SetDeliveryGroup sets field value
func (o *DeliveryGroupTestResponseModel) SetDeliveryGroup(v RefResponseModel) {
	o.DeliveryGroup = v
}

// GetNumPassed returns the NumPassed field value
func (o *DeliveryGroupTestResponseModel) GetNumPassed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPassed
}

// GetNumPassedOk returns a tuple with the NumPassed field value
// and a boolean to check if the value has been set.
func (o *DeliveryGroupTestResponseModel) GetNumPassedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPassed, true
}

// SetNumPassed sets field value
func (o *DeliveryGroupTestResponseModel) SetNumPassed(v int32) {
	o.NumPassed = v
}

// GetNumWarnings returns the NumWarnings field value
func (o *DeliveryGroupTestResponseModel) GetNumWarnings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWarnings
}

// GetNumWarningsOk returns a tuple with the NumWarnings field value
// and a boolean to check if the value has been set.
func (o *DeliveryGroupTestResponseModel) GetNumWarningsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWarnings, true
}

// SetNumWarnings sets field value
func (o *DeliveryGroupTestResponseModel) SetNumWarnings(v int32) {
	o.NumWarnings = v
}

// GetNumFailures returns the NumFailures field value
func (o *DeliveryGroupTestResponseModel) GetNumFailures() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumFailures
}

// GetNumFailuresOk returns a tuple with the NumFailures field value
// and a boolean to check if the value has been set.
func (o *DeliveryGroupTestResponseModel) GetNumFailuresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumFailures, true
}

// SetNumFailures sets field value
func (o *DeliveryGroupTestResponseModel) SetNumFailures(v int32) {
	o.NumFailures = v
}

func (o DeliveryGroupTestResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryGroupTestResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DeliveryGroup"] = o.DeliveryGroup
	toSerialize["NumPassed"] = o.NumPassed
	toSerialize["NumWarnings"] = o.NumWarnings
	toSerialize["NumFailures"] = o.NumFailures
	return toSerialize, nil
}

type NullableDeliveryGroupTestResponseModel struct {
	value *DeliveryGroupTestResponseModel
	isSet bool
}

func (v NullableDeliveryGroupTestResponseModel) Get() *DeliveryGroupTestResponseModel {
	return v.value
}

func (v *NullableDeliveryGroupTestResponseModel) Set(val *DeliveryGroupTestResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGroupTestResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGroupTestResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGroupTestResponseModel(val *DeliveryGroupTestResponseModel) *NullableDeliveryGroupTestResponseModel {
	return &NullableDeliveryGroupTestResponseModel{value: val, isSet: true}
}

func (v NullableDeliveryGroupTestResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGroupTestResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


