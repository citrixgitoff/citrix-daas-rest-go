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

// checks if the ProjectedAutoscaleMachinesResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectedAutoscaleMachinesResponseModel{}

// ProjectedAutoscaleMachinesResponseModel Projected number of machines that Autoscale will keep powered on over the specified period
type ProjectedAutoscaleMachinesResponseModel struct {
	// The Id of the desired desktop group.
	DesktopGroupId string `json:"DesktopGroupId"`
	// The number of machines in the desktop group that would be managed by Autoscale. This number shall include all the power managed machines in the desktop group, except for machines in maintenance mode or, not tagged for Autoscale if tag restriction is being used.
	ManagedMachineCount int32 `json:"ManagedMachineCount"`
	ProjectedAutoscaleMachines ProjectedMachineResponseModel `json:"ProjectedAutoscaleMachines"`
	// The time zone in which this delivery group's machines reside.
	TimeZone string `json:"TimeZone"`
}

// NewProjectedAutoscaleMachinesResponseModel instantiates a new ProjectedAutoscaleMachinesResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectedAutoscaleMachinesResponseModel(desktopGroupId string, managedMachineCount int32, projectedAutoscaleMachines ProjectedMachineResponseModel, timeZone string) *ProjectedAutoscaleMachinesResponseModel {
	this := ProjectedAutoscaleMachinesResponseModel{}
	this.DesktopGroupId = desktopGroupId
	this.ManagedMachineCount = managedMachineCount
	this.ProjectedAutoscaleMachines = projectedAutoscaleMachines
	this.TimeZone = timeZone
	return &this
}

// NewProjectedAutoscaleMachinesResponseModelWithDefaults instantiates a new ProjectedAutoscaleMachinesResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectedAutoscaleMachinesResponseModelWithDefaults() *ProjectedAutoscaleMachinesResponseModel {
	this := ProjectedAutoscaleMachinesResponseModel{}
	return &this
}

// GetDesktopGroupId returns the DesktopGroupId field value
func (o *ProjectedAutoscaleMachinesResponseModel) GetDesktopGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesktopGroupId
}

// GetDesktopGroupIdOk returns a tuple with the DesktopGroupId field value
// and a boolean to check if the value has been set.
func (o *ProjectedAutoscaleMachinesResponseModel) GetDesktopGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesktopGroupId, true
}

// SetDesktopGroupId sets field value
func (o *ProjectedAutoscaleMachinesResponseModel) SetDesktopGroupId(v string) {
	o.DesktopGroupId = v
}

// GetManagedMachineCount returns the ManagedMachineCount field value
func (o *ProjectedAutoscaleMachinesResponseModel) GetManagedMachineCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ManagedMachineCount
}

// GetManagedMachineCountOk returns a tuple with the ManagedMachineCount field value
// and a boolean to check if the value has been set.
func (o *ProjectedAutoscaleMachinesResponseModel) GetManagedMachineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagedMachineCount, true
}

// SetManagedMachineCount sets field value
func (o *ProjectedAutoscaleMachinesResponseModel) SetManagedMachineCount(v int32) {
	o.ManagedMachineCount = v
}

// GetProjectedAutoscaleMachines returns the ProjectedAutoscaleMachines field value
func (o *ProjectedAutoscaleMachinesResponseModel) GetProjectedAutoscaleMachines() ProjectedMachineResponseModel {
	if o == nil {
		var ret ProjectedMachineResponseModel
		return ret
	}

	return o.ProjectedAutoscaleMachines
}

// GetProjectedAutoscaleMachinesOk returns a tuple with the ProjectedAutoscaleMachines field value
// and a boolean to check if the value has been set.
func (o *ProjectedAutoscaleMachinesResponseModel) GetProjectedAutoscaleMachinesOk() (*ProjectedMachineResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectedAutoscaleMachines, true
}

// SetProjectedAutoscaleMachines sets field value
func (o *ProjectedAutoscaleMachinesResponseModel) SetProjectedAutoscaleMachines(v ProjectedMachineResponseModel) {
	o.ProjectedAutoscaleMachines = v
}

// GetTimeZone returns the TimeZone field value
func (o *ProjectedAutoscaleMachinesResponseModel) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ProjectedAutoscaleMachinesResponseModel) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *ProjectedAutoscaleMachinesResponseModel) SetTimeZone(v string) {
	o.TimeZone = v
}

func (o ProjectedAutoscaleMachinesResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectedAutoscaleMachinesResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DesktopGroupId"] = o.DesktopGroupId
	toSerialize["ManagedMachineCount"] = o.ManagedMachineCount
	toSerialize["ProjectedAutoscaleMachines"] = o.ProjectedAutoscaleMachines
	toSerialize["TimeZone"] = o.TimeZone
	return toSerialize, nil
}

type NullableProjectedAutoscaleMachinesResponseModel struct {
	value *ProjectedAutoscaleMachinesResponseModel
	isSet bool
}

func (v NullableProjectedAutoscaleMachinesResponseModel) Get() *ProjectedAutoscaleMachinesResponseModel {
	return v.value
}

func (v *NullableProjectedAutoscaleMachinesResponseModel) Set(val *ProjectedAutoscaleMachinesResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectedAutoscaleMachinesResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectedAutoscaleMachinesResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectedAutoscaleMachinesResponseModel(val *ProjectedAutoscaleMachinesResponseModel) *NullableProjectedAutoscaleMachinesResponseModel {
	return &NullableProjectedAutoscaleMachinesResponseModel{value: val, isSet: true}
}

func (v NullableProjectedAutoscaleMachinesResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectedAutoscaleMachinesResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


