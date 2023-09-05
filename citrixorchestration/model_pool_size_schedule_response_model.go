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

// checks if the PoolSizeScheduleResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolSizeScheduleResponseModel{}

// PoolSizeScheduleResponseModel The pool size to power on during a time range.
type PoolSizeScheduleResponseModel struct {
	// Time range during which the pool size applies.
	TimeRange *string `json:"TimeRange,omitempty"`
	// The number of machines (either as an absolute number or a percentage of the machines in the delivery group, depending on the value of PoolUsingPercentage) that are to be maintained in a running state, whether they are in use or not.
	PoolSize *int32 `json:"PoolSize,omitempty"`
}

// NewPoolSizeScheduleResponseModel instantiates a new PoolSizeScheduleResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolSizeScheduleResponseModel() *PoolSizeScheduleResponseModel {
	this := PoolSizeScheduleResponseModel{}
	return &this
}

// NewPoolSizeScheduleResponseModelWithDefaults instantiates a new PoolSizeScheduleResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolSizeScheduleResponseModelWithDefaults() *PoolSizeScheduleResponseModel {
	this := PoolSizeScheduleResponseModel{}
	return &this
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *PoolSizeScheduleResponseModel) GetTimeRange() string {
	if o == nil || IsNil(o.TimeRange) {
		var ret string
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSizeScheduleResponseModel) GetTimeRangeOk() (*string, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *PoolSizeScheduleResponseModel) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given string and assigns it to the TimeRange field.
func (o *PoolSizeScheduleResponseModel) SetTimeRange(v string) {
	o.TimeRange = &v
}

// GetPoolSize returns the PoolSize field value if set, zero value otherwise.
func (o *PoolSizeScheduleResponseModel) GetPoolSize() int32 {
	if o == nil || IsNil(o.PoolSize) {
		var ret int32
		return ret
	}
	return *o.PoolSize
}

// GetPoolSizeOk returns a tuple with the PoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSizeScheduleResponseModel) GetPoolSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PoolSize) {
		return nil, false
	}
	return o.PoolSize, true
}

// HasPoolSize returns a boolean if a field has been set.
func (o *PoolSizeScheduleResponseModel) HasPoolSize() bool {
	if o != nil && !IsNil(o.PoolSize) {
		return true
	}

	return false
}

// SetPoolSize gets a reference to the given int32 and assigns it to the PoolSize field.
func (o *PoolSizeScheduleResponseModel) SetPoolSize(v int32) {
	o.PoolSize = &v
}

func (o PoolSizeScheduleResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolSizeScheduleResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeRange) {
		toSerialize["TimeRange"] = o.TimeRange
	}
	if !IsNil(o.PoolSize) {
		toSerialize["PoolSize"] = o.PoolSize
	}
	return toSerialize, nil
}

type NullablePoolSizeScheduleResponseModel struct {
	value *PoolSizeScheduleResponseModel
	isSet bool
}

func (v NullablePoolSizeScheduleResponseModel) Get() *PoolSizeScheduleResponseModel {
	return v.value
}

func (v *NullablePoolSizeScheduleResponseModel) Set(val *PoolSizeScheduleResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolSizeScheduleResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolSizeScheduleResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolSizeScheduleResponseModel(val *PoolSizeScheduleResponseModel) *NullablePoolSizeScheduleResponseModel {
	return &NullablePoolSizeScheduleResponseModel{value: val, isSet: true}
}

func (v NullablePoolSizeScheduleResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolSizeScheduleResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


