/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
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
	StartDateTimeUtc NullableTime `json:"StartDateTimeUtc,omitempty"`
	// Timeout duration in hours. Valid range is 1 to 24.
	DurationInHours int32 `json:"DurationInHours"`
	VDAComponentsAndFeaturesRequestModel *VDAComponentsSelectionValidationRequestModel `json:"VDAComponentsAndFeaturesRequestModel,omitempty"`
	// Custom location to download the VDA Workstation package from. Currently, only network shares (specified using a UNC path) are supported.
	VdaWorkstationPackageUri NullableString `json:"VdaWorkstationPackageUri,omitempty"`
	// Custom location to download the VDA Server package from. Currently, only network shares (specified using a UNC path) are supported.
	VdaServerPackageUri NullableString `json:"VdaServerPackageUri,omitempty"`
	// Limits the Concurrent upgrades that can happen at any time in Upgrade Window.
	ConcurrencyLevel NullableInt32 `json:"ConcurrencyLevel,omitempty"`
	// Limits the number of failures that can take place during a scheduled upgrade.
	FailureThreshold NullableInt32 `json:"FailureThreshold,omitempty"`
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

// GetStartDateTimeUtc returns the StartDateTimeUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtc() time.Time {
	if o == nil || IsNil(o.StartDateTimeUtc.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTimeUtc.Get()
}

// GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDateTimeUtc.Get(), o.StartDateTimeUtc.IsSet()
}

// HasStartDateTimeUtc returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasStartDateTimeUtc() bool {
	if o != nil && o.StartDateTimeUtc.IsSet() {
		return true
	}

	return false
}

// SetStartDateTimeUtc gets a reference to the given NullableTime and assigns it to the StartDateTimeUtc field.
func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtc(v time.Time) {
	o.StartDateTimeUtc.Set(&v)
}
// SetStartDateTimeUtcNil sets the value for StartDateTimeUtc to be an explicit nil
func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtcNil() {
	o.StartDateTimeUtc.Set(nil)
}

// UnsetStartDateTimeUtc ensures that no value is present for StartDateTimeUtc, not even an explicit nil
func (o *CreateUpgradeScheduleRequestModel) UnsetStartDateTimeUtc() {
	o.StartDateTimeUtc.Unset()
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
func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModel() VDAComponentsSelectionValidationRequestModel {
	if o == nil || IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		var ret VDAComponentsSelectionValidationRequestModel
		return ret
	}
	return *o.VDAComponentsAndFeaturesRequestModel
}

// GetVDAComponentsAndFeaturesRequestModelOk returns a tuple with the VDAComponentsAndFeaturesRequestModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModelOk() (*VDAComponentsSelectionValidationRequestModel, bool) {
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

// SetVDAComponentsAndFeaturesRequestModel gets a reference to the given VDAComponentsSelectionValidationRequestModel and assigns it to the VDAComponentsAndFeaturesRequestModel field.
func (o *CreateUpgradeScheduleRequestModel) SetVDAComponentsAndFeaturesRequestModel(v VDAComponentsSelectionValidationRequestModel) {
	o.VDAComponentsAndFeaturesRequestModel = &v
}

// GetVdaWorkstationPackageUri returns the VdaWorkstationPackageUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeScheduleRequestModel) GetVdaWorkstationPackageUri() string {
	if o == nil || IsNil(o.VdaWorkstationPackageUri.Get()) {
		var ret string
		return ret
	}
	return *o.VdaWorkstationPackageUri.Get()
}

// GetVdaWorkstationPackageUriOk returns a tuple with the VdaWorkstationPackageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeScheduleRequestModel) GetVdaWorkstationPackageUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VdaWorkstationPackageUri.Get(), o.VdaWorkstationPackageUri.IsSet()
}

// HasVdaWorkstationPackageUri returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasVdaWorkstationPackageUri() bool {
	if o != nil && o.VdaWorkstationPackageUri.IsSet() {
		return true
	}

	return false
}

// SetVdaWorkstationPackageUri gets a reference to the given NullableString and assigns it to the VdaWorkstationPackageUri field.
func (o *CreateUpgradeScheduleRequestModel) SetVdaWorkstationPackageUri(v string) {
	o.VdaWorkstationPackageUri.Set(&v)
}
// SetVdaWorkstationPackageUriNil sets the value for VdaWorkstationPackageUri to be an explicit nil
func (o *CreateUpgradeScheduleRequestModel) SetVdaWorkstationPackageUriNil() {
	o.VdaWorkstationPackageUri.Set(nil)
}

// UnsetVdaWorkstationPackageUri ensures that no value is present for VdaWorkstationPackageUri, not even an explicit nil
func (o *CreateUpgradeScheduleRequestModel) UnsetVdaWorkstationPackageUri() {
	o.VdaWorkstationPackageUri.Unset()
}

// GetVdaServerPackageUri returns the VdaServerPackageUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeScheduleRequestModel) GetVdaServerPackageUri() string {
	if o == nil || IsNil(o.VdaServerPackageUri.Get()) {
		var ret string
		return ret
	}
	return *o.VdaServerPackageUri.Get()
}

// GetVdaServerPackageUriOk returns a tuple with the VdaServerPackageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeScheduleRequestModel) GetVdaServerPackageUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VdaServerPackageUri.Get(), o.VdaServerPackageUri.IsSet()
}

// HasVdaServerPackageUri returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasVdaServerPackageUri() bool {
	if o != nil && o.VdaServerPackageUri.IsSet() {
		return true
	}

	return false
}

// SetVdaServerPackageUri gets a reference to the given NullableString and assigns it to the VdaServerPackageUri field.
func (o *CreateUpgradeScheduleRequestModel) SetVdaServerPackageUri(v string) {
	o.VdaServerPackageUri.Set(&v)
}
// SetVdaServerPackageUriNil sets the value for VdaServerPackageUri to be an explicit nil
func (o *CreateUpgradeScheduleRequestModel) SetVdaServerPackageUriNil() {
	o.VdaServerPackageUri.Set(nil)
}

// UnsetVdaServerPackageUri ensures that no value is present for VdaServerPackageUri, not even an explicit nil
func (o *CreateUpgradeScheduleRequestModel) UnsetVdaServerPackageUri() {
	o.VdaServerPackageUri.Unset()
}

// GetConcurrencyLevel returns the ConcurrencyLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeScheduleRequestModel) GetConcurrencyLevel() int32 {
	if o == nil || IsNil(o.ConcurrencyLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.ConcurrencyLevel.Get()
}

// GetConcurrencyLevelOk returns a tuple with the ConcurrencyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeScheduleRequestModel) GetConcurrencyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConcurrencyLevel.Get(), o.ConcurrencyLevel.IsSet()
}

// HasConcurrencyLevel returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasConcurrencyLevel() bool {
	if o != nil && o.ConcurrencyLevel.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyLevel gets a reference to the given NullableInt32 and assigns it to the ConcurrencyLevel field.
func (o *CreateUpgradeScheduleRequestModel) SetConcurrencyLevel(v int32) {
	o.ConcurrencyLevel.Set(&v)
}
// SetConcurrencyLevelNil sets the value for ConcurrencyLevel to be an explicit nil
func (o *CreateUpgradeScheduleRequestModel) SetConcurrencyLevelNil() {
	o.ConcurrencyLevel.Set(nil)
}

// UnsetConcurrencyLevel ensures that no value is present for ConcurrencyLevel, not even an explicit nil
func (o *CreateUpgradeScheduleRequestModel) UnsetConcurrencyLevel() {
	o.ConcurrencyLevel.Unset()
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeScheduleRequestModel) GetFailureThreshold() int32 {
	if o == nil || IsNil(o.FailureThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.FailureThreshold.Get()
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeScheduleRequestModel) GetFailureThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureThreshold.Get(), o.FailureThreshold.IsSet()
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *CreateUpgradeScheduleRequestModel) HasFailureThreshold() bool {
	if o != nil && o.FailureThreshold.IsSet() {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given NullableInt32 and assigns it to the FailureThreshold field.
func (o *CreateUpgradeScheduleRequestModel) SetFailureThreshold(v int32) {
	o.FailureThreshold.Set(&v)
}
// SetFailureThresholdNil sets the value for FailureThreshold to be an explicit nil
func (o *CreateUpgradeScheduleRequestModel) SetFailureThresholdNil() {
	o.FailureThreshold.Set(nil)
}

// UnsetFailureThreshold ensures that no value is present for FailureThreshold, not even an explicit nil
func (o *CreateUpgradeScheduleRequestModel) UnsetFailureThreshold() {
	o.FailureThreshold.Unset()
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
	if o.StartDateTimeUtc.IsSet() {
		toSerialize["StartDateTimeUtc"] = o.StartDateTimeUtc.Get()
	}
	toSerialize["DurationInHours"] = o.DurationInHours
	if !IsNil(o.VDAComponentsAndFeaturesRequestModel) {
		toSerialize["VDAComponentsAndFeaturesRequestModel"] = o.VDAComponentsAndFeaturesRequestModel
	}
	if o.VdaWorkstationPackageUri.IsSet() {
		toSerialize["VdaWorkstationPackageUri"] = o.VdaWorkstationPackageUri.Get()
	}
	if o.VdaServerPackageUri.IsSet() {
		toSerialize["VdaServerPackageUri"] = o.VdaServerPackageUri.Get()
	}
	if o.ConcurrencyLevel.IsSet() {
		toSerialize["ConcurrencyLevel"] = o.ConcurrencyLevel.Get()
	}
	if o.FailureThreshold.IsSet() {
		toSerialize["FailureThreshold"] = o.FailureThreshold.Get()
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


