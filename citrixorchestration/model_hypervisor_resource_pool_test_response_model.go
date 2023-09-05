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

// checks if the HypervisorResourcePoolTestResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolTestResponseModel{}

// HypervisorResourcePoolTestResponseModel Response object for a resource pool test result.
type HypervisorResourcePoolTestResponseModel struct {
	ResourcePool HypervisorResourcePoolTestResponseModelResourcePool `json:"ResourcePool"`
	// The number of tests that passed.
	NumPassed int32 `json:"NumPassed"`
	// The number of warnings that were found.
	NumWarnings int32 `json:"NumWarnings"`
	// The number of tests that failed.
	NumFailures int32 `json:"NumFailures"`
}

// NewHypervisorResourcePoolTestResponseModel instantiates a new HypervisorResourcePoolTestResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolTestResponseModel(resourcePool HypervisorResourcePoolTestResponseModelResourcePool, numPassed int32, numWarnings int32, numFailures int32) *HypervisorResourcePoolTestResponseModel {
	this := HypervisorResourcePoolTestResponseModel{}
	this.ResourcePool = resourcePool
	this.NumPassed = numPassed
	this.NumWarnings = numWarnings
	this.NumFailures = numFailures
	return &this
}

// NewHypervisorResourcePoolTestResponseModelWithDefaults instantiates a new HypervisorResourcePoolTestResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolTestResponseModelWithDefaults() *HypervisorResourcePoolTestResponseModel {
	this := HypervisorResourcePoolTestResponseModel{}
	return &this
}

// GetResourcePool returns the ResourcePool field value
func (o *HypervisorResourcePoolTestResponseModel) GetResourcePool() HypervisorResourcePoolTestResponseModelResourcePool {
	if o == nil {
		var ret HypervisorResourcePoolTestResponseModelResourcePool
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolTestResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolTestResponseModelResourcePool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *HypervisorResourcePoolTestResponseModel) SetResourcePool(v HypervisorResourcePoolTestResponseModelResourcePool) {
	o.ResourcePool = v
}

// GetNumPassed returns the NumPassed field value
func (o *HypervisorResourcePoolTestResponseModel) GetNumPassed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPassed
}

// GetNumPassedOk returns a tuple with the NumPassed field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolTestResponseModel) GetNumPassedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPassed, true
}

// SetNumPassed sets field value
func (o *HypervisorResourcePoolTestResponseModel) SetNumPassed(v int32) {
	o.NumPassed = v
}

// GetNumWarnings returns the NumWarnings field value
func (o *HypervisorResourcePoolTestResponseModel) GetNumWarnings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWarnings
}

// GetNumWarningsOk returns a tuple with the NumWarnings field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolTestResponseModel) GetNumWarningsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWarnings, true
}

// SetNumWarnings sets field value
func (o *HypervisorResourcePoolTestResponseModel) SetNumWarnings(v int32) {
	o.NumWarnings = v
}

// GetNumFailures returns the NumFailures field value
func (o *HypervisorResourcePoolTestResponseModel) GetNumFailures() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumFailures
}

// GetNumFailuresOk returns a tuple with the NumFailures field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolTestResponseModel) GetNumFailuresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumFailures, true
}

// SetNumFailures sets field value
func (o *HypervisorResourcePoolTestResponseModel) SetNumFailures(v int32) {
	o.NumFailures = v
}

func (o HypervisorResourcePoolTestResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolTestResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ResourcePool"] = o.ResourcePool
	toSerialize["NumPassed"] = o.NumPassed
	toSerialize["NumWarnings"] = o.NumWarnings
	toSerialize["NumFailures"] = o.NumFailures
	return toSerialize, nil
}

type NullableHypervisorResourcePoolTestResponseModel struct {
	value *HypervisorResourcePoolTestResponseModel
	isSet bool
}

func (v NullableHypervisorResourcePoolTestResponseModel) Get() *HypervisorResourcePoolTestResponseModel {
	return v.value
}

func (v *NullableHypervisorResourcePoolTestResponseModel) Set(val *HypervisorResourcePoolTestResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolTestResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolTestResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolTestResponseModel(val *HypervisorResourcePoolTestResponseModel) *NullableHypervisorResourcePoolTestResponseModel {
	return &NullableHypervisorResourcePoolTestResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolTestResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolTestResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


