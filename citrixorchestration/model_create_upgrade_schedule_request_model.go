/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"time"
)

// checks if the CreateUpgradeScheduleRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpgradeScheduleRequestModel{}

// CreateUpgradeScheduleRequestModel Request object for creation of VDA upgrade schedule on a machine catalog or machine object. There should not already be an upgrade schedule on the same object in UpgradeScheduled or UpgradeInProgress state. In such cases, you can only update or cancel the schedule.
type CreateUpgradeScheduleRequestModel struct {
	// UTC time to start the Vda upgrade. Must be a future time. If set to null, the upgrade should be started at once.
	StartDateTimeUtc *time.Time `json:"StartDateTimeUtc,omitempty"`
	// Timeout duration in hours. Valid range is 1 to 24.
	DurationInHours int32 `json:"DurationInHours"`
	VDAComponentsAndFeaturesRequestModel *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel `json:"VDAComponentsAndFeaturesRequestModel,omitempty"`
}

// NewCreateUpgradeScheduleRequestModel instantiates a new CreateUpgradeScheduleRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpgradeScheduleRequestModel(durationInHours int32) *CreateUpgradeScheduleRequestModel {
	this := CreateUpgradeScheduleRequestModel{}
	this.DurationInHours = durationInHours
	return &this
}

// NewCreateUpgradeScheduleRequestModelWithDefaults instantiates a new CreateUpgradeScheduleRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpgradeScheduleRequestModelWithDefaults() *CreateUpgradeScheduleRequestModel {
	this := CreateUpgradeScheduleRequestModel{}
	return &this
}

// GetStartDateTimeUtc returns the StartDateTimeUtc field value if set, zero value otherwise.
func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtc() time.Time {
	if o == nil || IsNil(o.StartDateTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTimeUtc
}

// GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTimeUtc) {
		return nil, false
	}
	return o.StartDateTimeUtc, true
}

// HasStartDateTimeUtc returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasStartDateTimeUtc() bool {
	if o != nil && !IsNil(o.StartDateTimeUtc) {
		return true
	}

	return false
}

// SetStartDateTimeUtc gets a reference to the given time.Time and assigns it to the StartDateTimeUtc field.
func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtc(v time.Time) {
	o.StartDateTimeUtc = &v
}

// GetDurationInHours returns the DurationInHours field value
func (o *CreateUpgradeScheduleRequestModel) GetDurationInHours() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationInHours
}

// GetDurationInHoursOk returns a tuple with the DurationInHours field value
// and a boolean to check if the value has been set.
func (o *CreateUpgradeScheduleRequestModel) GetDurationInHoursOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInHours, true
}

// SetDurationInHours sets field value
func (o *CreateUpgradeScheduleRequestModel) SetDurationInHours(v int32) {
	o.DurationInHours = v
}

// GetVDAComponentsAndFeaturesRequestModel returns the VDAComponentsAndFeaturesRequestModel field value if set, zero value otherwise.
func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModel() CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel {
	if o == nil || IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		var ret CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel
		return ret
	}
	return *o.VDAComponentsAndFeaturesRequestModel
}

// GetVDAComponentsAndFeaturesRequestModelOk returns a tuple with the VDAComponentsAndFeaturesRequestModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModelOk() (*CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel, bool) {
	if o == nil || IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		return nil, false
	}
	return o.VDAComponentsAndFeaturesRequestModel, true
}

// HasVDAComponentsAndFeaturesRequestModel returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasVDAComponentsAndFeaturesRequestModel() bool {
	if o != nil && !IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		return true
	}

	return false
}

// SetVDAComponentsAndFeaturesRequestModel gets a reference to the given CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel and assigns it to the VDAComponentsAndFeaturesRequestModel field.
func (o *CreateUpgradeScheduleRequestModel) SetVDAComponentsAndFeaturesRequestModel(v CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) {
	o.VDAComponentsAndFeaturesRequestModel = &v
}

func (o CreateUpgradeScheduleRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUpgradeScheduleRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDateTimeUtc) {
		toSerialize["StartDateTimeUtc"] = o.StartDateTimeUtc
	}
	toSerialize["DurationInHours"] = o.DurationInHours
	if !IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		toSerialize["VDAComponentsAndFeaturesRequestModel"] = o.VDAComponentsAndFeaturesRequestModel
	}
	return toSerialize, nil
}

type NullableCreateUpgradeScheduleRequestModel struct {
	value *CreateUpgradeScheduleRequestModel
	isSet bool
}

func (v NullableCreateUpgradeScheduleRequestModel) Get() *CreateUpgradeScheduleRequestModel {
	return v.value
}

func (v *NullableCreateUpgradeScheduleRequestModel) Set(val *CreateUpgradeScheduleRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpgradeScheduleRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpgradeScheduleRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpgradeScheduleRequestModel(val *CreateUpgradeScheduleRequestModel) *NullableCreateUpgradeScheduleRequestModel {
	return &NullableCreateUpgradeScheduleRequestModel{value: val, isSet: true}
}

func (v NullableCreateUpgradeScheduleRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpgradeScheduleRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


