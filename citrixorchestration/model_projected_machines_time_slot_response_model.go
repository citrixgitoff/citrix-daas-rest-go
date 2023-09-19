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

// checks if the ProjectedMachinesTimeSlotResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectedMachinesTimeSlotResponseModel{}

// ProjectedMachinesTimeSlotResponseModel ProjectedMachinesTimeSlotResponseModel
type ProjectedMachinesTimeSlotResponseModel struct {
	// The start time (in 24-hour clock) for the 30-minute period represented that is based on the time zone of the delivery group.
	Time string `json:"Time"`
	// The projected machines as per the pool size configured for the current 30 minute period.
	PoolSizeCount int32 `json:"PoolSizeCount"`
	// The projected machines as per the buffer percentage configured for the current 30 minute period.
	BufferSizeCount int32 `json:"BufferSizeCount"`
}

// NewProjectedMachinesTimeSlotResponseModel instantiates a new ProjectedMachinesTimeSlotResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectedMachinesTimeSlotResponseModel(time string, poolSizeCount int32, bufferSizeCount int32) *ProjectedMachinesTimeSlotResponseModel {
	this := ProjectedMachinesTimeSlotResponseModel{}
	this.Time = time
	this.PoolSizeCount = poolSizeCount
	this.BufferSizeCount = bufferSizeCount
	return &this
}

// NewProjectedMachinesTimeSlotResponseModelWithDefaults instantiates a new ProjectedMachinesTimeSlotResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectedMachinesTimeSlotResponseModelWithDefaults() *ProjectedMachinesTimeSlotResponseModel {
	this := ProjectedMachinesTimeSlotResponseModel{}
	return &this
}

// GetTime returns the Time field value
func (o *ProjectedMachinesTimeSlotResponseModel) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ProjectedMachinesTimeSlotResponseModel) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ProjectedMachinesTimeSlotResponseModel) SetTime(v string) {
	o.Time = v
}

// GetPoolSizeCount returns the PoolSizeCount field value
func (o *ProjectedMachinesTimeSlotResponseModel) GetPoolSizeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PoolSizeCount
}

// GetPoolSizeCountOk returns a tuple with the PoolSizeCount field value
// and a boolean to check if the value has been set.
func (o *ProjectedMachinesTimeSlotResponseModel) GetPoolSizeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolSizeCount, true
}

// SetPoolSizeCount sets field value
func (o *ProjectedMachinesTimeSlotResponseModel) SetPoolSizeCount(v int32) {
	o.PoolSizeCount = v
}

// GetBufferSizeCount returns the BufferSizeCount field value
func (o *ProjectedMachinesTimeSlotResponseModel) GetBufferSizeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BufferSizeCount
}

// GetBufferSizeCountOk returns a tuple with the BufferSizeCount field value
// and a boolean to check if the value has been set.
func (o *ProjectedMachinesTimeSlotResponseModel) GetBufferSizeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BufferSizeCount, true
}

// SetBufferSizeCount sets field value
func (o *ProjectedMachinesTimeSlotResponseModel) SetBufferSizeCount(v int32) {
	o.BufferSizeCount = v
}

func (o ProjectedMachinesTimeSlotResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectedMachinesTimeSlotResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Time"] = o.Time
	toSerialize["PoolSizeCount"] = o.PoolSizeCount
	toSerialize["BufferSizeCount"] = o.BufferSizeCount
	return toSerialize, nil
}

type NullableProjectedMachinesTimeSlotResponseModel struct {
	value *ProjectedMachinesTimeSlotResponseModel
	isSet bool
}

func (v NullableProjectedMachinesTimeSlotResponseModel) Get() *ProjectedMachinesTimeSlotResponseModel {
	return v.value
}

func (v *NullableProjectedMachinesTimeSlotResponseModel) Set(val *ProjectedMachinesTimeSlotResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectedMachinesTimeSlotResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectedMachinesTimeSlotResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectedMachinesTimeSlotResponseModel(val *ProjectedMachinesTimeSlotResponseModel) *NullableProjectedMachinesTimeSlotResponseModel {
	return &NullableProjectedMachinesTimeSlotResponseModel{value: val, isSet: true}
}

func (v NullableProjectedMachinesTimeSlotResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectedMachinesTimeSlotResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


