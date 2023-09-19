/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"time"
)

// checks if the MachineCatalogUpgradeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineCatalogUpgradeDetail{}

// MachineCatalogUpgradeDetail Status detail of last VDA upgrade schedule for a machine catalog.
type MachineCatalogUpgradeDetail struct {
	ScheduleStatus *VdaUpgradeScheduleStatus `json:"ScheduleStatus,omitempty"`
	// UTC time when this upgrade status object changed status for the last time.
	LastStateChangeTimeUtc *time.Time `json:"LastStateChangeTimeUtc,omitempty"`
	// Count of machines with an upgrade schedule in this catalog. This does not always equal to the machine count of the catalog, as there might be machines which joined the catalog after a catalog level upgrade is schdeuled, so do not have an upgrade status.
	TotalMachines *int32 `json:"TotalMachines,omitempty"`
	// Count of machines whose last upgrade is in success state.
	SuccessCount *int32 `json:"SuccessCount,omitempty"`
	// Count of machines whose last upgrade failed during package validation.
	ValidationFailureCount *int32 `json:"ValidationFailureCount,omitempty"`
	// Count of machines whose upgrade is in progress.
	InProgressCount *int32 `json:"InProgressCount,omitempty"`
	// Count of machines whose last upgrade failed during package installtion.
	UpgradeFailureCount *int32 `json:"UpgradeFailureCount,omitempty"`
	// UTC time when this VDA upgrade was scheduled to start.
	ScheduledTimeUtc *time.Time `json:"ScheduledTimeUtc,omitempty"`
	// Timeout duration in hours, of the current VDA upgrade schdeule.
	DurationInHours *int32 `json:"DurationInHours,omitempty"`
	// Target package version of the current VDA upgrade schdeule.
	TargetPackageVersion NullableString `json:"TargetPackageVersion,omitempty"`
	// Count of machines whose last upgrade canceled during package installtion.
	CancelledUpgradeCount *int32 `json:"CancelledUpgradeCount,omitempty"`
	// Count of machines who is waiting to upgrade.
	WaitingToUpgradeCount *int32 `json:"WaitingToUpgradeCount,omitempty"`
}

// NewMachineCatalogUpgradeDetail instantiates a new MachineCatalogUpgradeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineCatalogUpgradeDetail() *MachineCatalogUpgradeDetail {
	this := MachineCatalogUpgradeDetail{}
	return &this
}

// NewMachineCatalogUpgradeDetailWithDefaults instantiates a new MachineCatalogUpgradeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineCatalogUpgradeDetailWithDefaults() *MachineCatalogUpgradeDetail {
	this := MachineCatalogUpgradeDetail{}
	return &this
}

// GetScheduleStatus returns the ScheduleStatus field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetScheduleStatus() VdaUpgradeScheduleStatus {
	if o == nil || IsNil(o.ScheduleStatus) {
		var ret VdaUpgradeScheduleStatus
		return ret
	}
	return *o.ScheduleStatus
}

// GetScheduleStatusOk returns a tuple with the ScheduleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeScheduleStatus, bool) {
	if o == nil || IsNil(o.ScheduleStatus) {
		return nil, false
	}
	return o.ScheduleStatus, true
}

// HasScheduleStatus returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasScheduleStatus() bool {
	if o != nil && !IsNil(o.ScheduleStatus) {
		return true
	}

	return false
}

// SetScheduleStatus gets a reference to the given VdaUpgradeScheduleStatus and assigns it to the ScheduleStatus field.
func (o *MachineCatalogUpgradeDetail) SetScheduleStatus(v VdaUpgradeScheduleStatus) {
	o.ScheduleStatus = &v
}

// GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetLastStateChangeTimeUtc() time.Time {
	if o == nil || IsNil(o.LastStateChangeTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.LastStateChangeTimeUtc
}

// GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetLastStateChangeTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastStateChangeTimeUtc) {
		return nil, false
	}
	return o.LastStateChangeTimeUtc, true
}

// HasLastStateChangeTimeUtc returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasLastStateChangeTimeUtc() bool {
	if o != nil && !IsNil(o.LastStateChangeTimeUtc) {
		return true
	}

	return false
}

// SetLastStateChangeTimeUtc gets a reference to the given time.Time and assigns it to the LastStateChangeTimeUtc field.
func (o *MachineCatalogUpgradeDetail) SetLastStateChangeTimeUtc(v time.Time) {
	o.LastStateChangeTimeUtc = &v
}

// GetTotalMachines returns the TotalMachines field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetTotalMachines() int32 {
	if o == nil || IsNil(o.TotalMachines) {
		var ret int32
		return ret
	}
	return *o.TotalMachines
}

// GetTotalMachinesOk returns a tuple with the TotalMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetTotalMachinesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalMachines) {
		return nil, false
	}
	return o.TotalMachines, true
}

// HasTotalMachines returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasTotalMachines() bool {
	if o != nil && !IsNil(o.TotalMachines) {
		return true
	}

	return false
}

// SetTotalMachines gets a reference to the given int32 and assigns it to the TotalMachines field.
func (o *MachineCatalogUpgradeDetail) SetTotalMachines(v int32) {
	o.TotalMachines = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetSuccessCount() int32 {
	if o == nil || IsNil(o.SuccessCount) {
		var ret int32
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetSuccessCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SuccessCount) {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasSuccessCount() bool {
	if o != nil && !IsNil(o.SuccessCount) {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.
func (o *MachineCatalogUpgradeDetail) SetSuccessCount(v int32) {
	o.SuccessCount = &v
}

// GetValidationFailureCount returns the ValidationFailureCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetValidationFailureCount() int32 {
	if o == nil || IsNil(o.ValidationFailureCount) {
		var ret int32
		return ret
	}
	return *o.ValidationFailureCount
}

// GetValidationFailureCountOk returns a tuple with the ValidationFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetValidationFailureCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidationFailureCount) {
		return nil, false
	}
	return o.ValidationFailureCount, true
}

// HasValidationFailureCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasValidationFailureCount() bool {
	if o != nil && !IsNil(o.ValidationFailureCount) {
		return true
	}

	return false
}

// SetValidationFailureCount gets a reference to the given int32 and assigns it to the ValidationFailureCount field.
func (o *MachineCatalogUpgradeDetail) SetValidationFailureCount(v int32) {
	o.ValidationFailureCount = &v
}

// GetInProgressCount returns the InProgressCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetInProgressCount() int32 {
	if o == nil || IsNil(o.InProgressCount) {
		var ret int32
		return ret
	}
	return *o.InProgressCount
}

// GetInProgressCountOk returns a tuple with the InProgressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetInProgressCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InProgressCount) {
		return nil, false
	}
	return o.InProgressCount, true
}

// HasInProgressCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasInProgressCount() bool {
	if o != nil && !IsNil(o.InProgressCount) {
		return true
	}

	return false
}

// SetInProgressCount gets a reference to the given int32 and assigns it to the InProgressCount field.
func (o *MachineCatalogUpgradeDetail) SetInProgressCount(v int32) {
	o.InProgressCount = &v
}

// GetUpgradeFailureCount returns the UpgradeFailureCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetUpgradeFailureCount() int32 {
	if o == nil || IsNil(o.UpgradeFailureCount) {
		var ret int32
		return ret
	}
	return *o.UpgradeFailureCount
}

// GetUpgradeFailureCountOk returns a tuple with the UpgradeFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetUpgradeFailureCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UpgradeFailureCount) {
		return nil, false
	}
	return o.UpgradeFailureCount, true
}

// HasUpgradeFailureCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasUpgradeFailureCount() bool {
	if o != nil && !IsNil(o.UpgradeFailureCount) {
		return true
	}

	return false
}

// SetUpgradeFailureCount gets a reference to the given int32 and assigns it to the UpgradeFailureCount field.
func (o *MachineCatalogUpgradeDetail) SetUpgradeFailureCount(v int32) {
	o.UpgradeFailureCount = &v
}

// GetScheduledTimeUtc returns the ScheduledTimeUtc field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetScheduledTimeUtc() time.Time {
	if o == nil || IsNil(o.ScheduledTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledTimeUtc
}

// GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetScheduledTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledTimeUtc) {
		return nil, false
	}
	return o.ScheduledTimeUtc, true
}

// HasScheduledTimeUtc returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasScheduledTimeUtc() bool {
	if o != nil && !IsNil(o.ScheduledTimeUtc) {
		return true
	}

	return false
}

// SetScheduledTimeUtc gets a reference to the given time.Time and assigns it to the ScheduledTimeUtc field.
func (o *MachineCatalogUpgradeDetail) SetScheduledTimeUtc(v time.Time) {
	o.ScheduledTimeUtc = &v
}

// GetDurationInHours returns the DurationInHours field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetDurationInHours() int32 {
	if o == nil || IsNil(o.DurationInHours) {
		var ret int32
		return ret
	}
	return *o.DurationInHours
}

// GetDurationInHoursOk returns a tuple with the DurationInHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetDurationInHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInHours) {
		return nil, false
	}
	return o.DurationInHours, true
}

// HasDurationInHours returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasDurationInHours() bool {
	if o != nil && !IsNil(o.DurationInHours) {
		return true
	}

	return false
}

// SetDurationInHours gets a reference to the given int32 and assigns it to the DurationInHours field.
func (o *MachineCatalogUpgradeDetail) SetDurationInHours(v int32) {
	o.DurationInHours = &v
}

// GetTargetPackageVersion returns the TargetPackageVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineCatalogUpgradeDetail) GetTargetPackageVersion() string {
	if o == nil || IsNil(o.TargetPackageVersion.Get()) {
		var ret string
		return ret
	}
	return *o.TargetPackageVersion.Get()
}

// GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineCatalogUpgradeDetail) GetTargetPackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetPackageVersion.Get(), o.TargetPackageVersion.IsSet()
}

// HasTargetPackageVersion returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasTargetPackageVersion() bool {
	if o != nil && o.TargetPackageVersion.IsSet() {
		return true
	}

	return false
}

// SetTargetPackageVersion gets a reference to the given NullableString and assigns it to the TargetPackageVersion field.
func (o *MachineCatalogUpgradeDetail) SetTargetPackageVersion(v string) {
	o.TargetPackageVersion.Set(&v)
}
// SetTargetPackageVersionNil sets the value for TargetPackageVersion to be an explicit nil
func (o *MachineCatalogUpgradeDetail) SetTargetPackageVersionNil() {
	o.TargetPackageVersion.Set(nil)
}

// UnsetTargetPackageVersion ensures that no value is present for TargetPackageVersion, not even an explicit nil
func (o *MachineCatalogUpgradeDetail) UnsetTargetPackageVersion() {
	o.TargetPackageVersion.Unset()
}

// GetCancelledUpgradeCount returns the CancelledUpgradeCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetCancelledUpgradeCount() int32 {
	if o == nil || IsNil(o.CancelledUpgradeCount) {
		var ret int32
		return ret
	}
	return *o.CancelledUpgradeCount
}

// GetCancelledUpgradeCountOk returns a tuple with the CancelledUpgradeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetCancelledUpgradeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelledUpgradeCount) {
		return nil, false
	}
	return o.CancelledUpgradeCount, true
}

// HasCancelledUpgradeCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasCancelledUpgradeCount() bool {
	if o != nil && !IsNil(o.CancelledUpgradeCount) {
		return true
	}

	return false
}

// SetCancelledUpgradeCount gets a reference to the given int32 and assigns it to the CancelledUpgradeCount field.
func (o *MachineCatalogUpgradeDetail) SetCancelledUpgradeCount(v int32) {
	o.CancelledUpgradeCount = &v
}

// GetWaitingToUpgradeCount returns the WaitingToUpgradeCount field value if set, zero value otherwise.
func (o *MachineCatalogUpgradeDetail) GetWaitingToUpgradeCount() int32 {
	if o == nil || IsNil(o.WaitingToUpgradeCount) {
		var ret int32
		return ret
	}
	return *o.WaitingToUpgradeCount
}

// GetWaitingToUpgradeCountOk returns a tuple with the WaitingToUpgradeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogUpgradeDetail) GetWaitingToUpgradeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitingToUpgradeCount) {
		return nil, false
	}
	return o.WaitingToUpgradeCount, true
}

// HasWaitingToUpgradeCount returns a boolean if a field has been set.
func (o *MachineCatalogUpgradeDetail) HasWaitingToUpgradeCount() bool {
	if o != nil && !IsNil(o.WaitingToUpgradeCount) {
		return true
	}

	return false
}

// SetWaitingToUpgradeCount gets a reference to the given int32 and assigns it to the WaitingToUpgradeCount field.
func (o *MachineCatalogUpgradeDetail) SetWaitingToUpgradeCount(v int32) {
	o.WaitingToUpgradeCount = &v
}

func (o MachineCatalogUpgradeDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineCatalogUpgradeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduleStatus) {
		toSerialize["ScheduleStatus"] = o.ScheduleStatus
	}
	if !IsNil(o.LastStateChangeTimeUtc) {
		toSerialize["LastStateChangeTimeUtc"] = o.LastStateChangeTimeUtc
	}
	if !IsNil(o.TotalMachines) {
		toSerialize["TotalMachines"] = o.TotalMachines
	}
	if !IsNil(o.SuccessCount) {
		toSerialize["SuccessCount"] = o.SuccessCount
	}
	if !IsNil(o.ValidationFailureCount) {
		toSerialize["ValidationFailureCount"] = o.ValidationFailureCount
	}
	if !IsNil(o.InProgressCount) {
		toSerialize["InProgressCount"] = o.InProgressCount
	}
	if !IsNil(o.UpgradeFailureCount) {
		toSerialize["UpgradeFailureCount"] = o.UpgradeFailureCount
	}
	if !IsNil(o.ScheduledTimeUtc) {
		toSerialize["ScheduledTimeUtc"] = o.ScheduledTimeUtc
	}
	if !IsNil(o.DurationInHours) {
		toSerialize["DurationInHours"] = o.DurationInHours
	}
	if o.TargetPackageVersion.IsSet() {
		toSerialize["TargetPackageVersion"] = o.TargetPackageVersion.Get()
	}
	if !IsNil(o.CancelledUpgradeCount) {
		toSerialize["CancelledUpgradeCount"] = o.CancelledUpgradeCount
	}
	if !IsNil(o.WaitingToUpgradeCount) {
		toSerialize["WaitingToUpgradeCount"] = o.WaitingToUpgradeCount
	}
	return toSerialize, nil
}

type NullableMachineCatalogUpgradeDetail struct {
	value *MachineCatalogUpgradeDetail
	isSet bool
}

func (v NullableMachineCatalogUpgradeDetail) Get() *MachineCatalogUpgradeDetail {
	return v.value
}

func (v *NullableMachineCatalogUpgradeDetail) Set(val *MachineCatalogUpgradeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineCatalogUpgradeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineCatalogUpgradeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineCatalogUpgradeDetail(val *MachineCatalogUpgradeDetail) *NullableMachineCatalogUpgradeDetail {
	return &NullableMachineCatalogUpgradeDetail{value: val, isSet: true}
}

func (v NullableMachineCatalogUpgradeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineCatalogUpgradeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


