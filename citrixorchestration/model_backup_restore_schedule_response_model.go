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

// checks if the BackupRestoreScheduleResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreScheduleResponseModel{}

// BackupRestoreScheduleResponseModel struct for BackupRestoreScheduleResponseModel
type BackupRestoreScheduleResponseModel struct {
	// Name
	Name *string `json:"Name,omitempty"`
	// Id
	Uid *int32 `json:"Uid,omitempty"`
	Day *BackupRestoreScheduleDays `json:"Day,omitempty"`
	// Days in Week
	DaysInWeek []BackupRestoreScheduleDays `json:"DaysInWeek,omitempty"`
	DayInMonth *BackupRestoreScheduleDays `json:"DayInMonth,omitempty"`
	WeekInMonth *BackupRestoreScheduleWeeks `json:"WeekInMonth,omitempty"`
	// Start Date
	StartDate *string `json:"StartDate,omitempty"`
	// Frequency Factor
	FrequencyFactor *int32 `json:"FrequencyFactor,omitempty"`
	// Description
	Description *string `json:"Description,omitempty"`
	// Enabled             
	Enabled *bool `json:"Enabled,omitempty"`
	Frequency *BackupRestoreScheduleFrequency `json:"Frequency,omitempty"`
	// Start Time
	StartTime *string `json:"StartTime,omitempty"`
	// Time Zone             
	TimeZone *string `json:"TimeZone,omitempty"`
}

// NewBackupRestoreScheduleResponseModel instantiates a new BackupRestoreScheduleResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreScheduleResponseModel() *BackupRestoreScheduleResponseModel {
	this := BackupRestoreScheduleResponseModel{}
	return &this
}

// NewBackupRestoreScheduleResponseModelWithDefaults instantiates a new BackupRestoreScheduleResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreScheduleResponseModelWithDefaults() *BackupRestoreScheduleResponseModel {
	this := BackupRestoreScheduleResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupRestoreScheduleResponseModel) SetName(v string) {
	o.Name = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *BackupRestoreScheduleResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetDay() BackupRestoreScheduleDays {
	if o == nil || IsNil(o.Day) {
		var ret BackupRestoreScheduleDays
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetDayOk() (*BackupRestoreScheduleDays, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasDay() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given BackupRestoreScheduleDays and assigns it to the Day field.
func (o *BackupRestoreScheduleResponseModel) SetDay(v BackupRestoreScheduleDays) {
	o.Day = &v
}

// GetDaysInWeek returns the DaysInWeek field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetDaysInWeek() []BackupRestoreScheduleDays {
	if o == nil || IsNil(o.DaysInWeek) {
		var ret []BackupRestoreScheduleDays
		return ret
	}
	return o.DaysInWeek
}

// GetDaysInWeekOk returns a tuple with the DaysInWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetDaysInWeekOk() ([]BackupRestoreScheduleDays, bool) {
	if o == nil || IsNil(o.DaysInWeek) {
		return nil, false
	}
	return o.DaysInWeek, true
}

// HasDaysInWeek returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasDaysInWeek() bool {
	if o != nil && !IsNil(o.DaysInWeek) {
		return true
	}

	return false
}

// SetDaysInWeek gets a reference to the given []BackupRestoreScheduleDays and assigns it to the DaysInWeek field.
func (o *BackupRestoreScheduleResponseModel) SetDaysInWeek(v []BackupRestoreScheduleDays) {
	o.DaysInWeek = v
}

// GetDayInMonth returns the DayInMonth field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetDayInMonth() BackupRestoreScheduleDays {
	if o == nil || IsNil(o.DayInMonth) {
		var ret BackupRestoreScheduleDays
		return ret
	}
	return *o.DayInMonth
}

// GetDayInMonthOk returns a tuple with the DayInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetDayInMonthOk() (*BackupRestoreScheduleDays, bool) {
	if o == nil || IsNil(o.DayInMonth) {
		return nil, false
	}
	return o.DayInMonth, true
}

// HasDayInMonth returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasDayInMonth() bool {
	if o != nil && !IsNil(o.DayInMonth) {
		return true
	}

	return false
}

// SetDayInMonth gets a reference to the given BackupRestoreScheduleDays and assigns it to the DayInMonth field.
func (o *BackupRestoreScheduleResponseModel) SetDayInMonth(v BackupRestoreScheduleDays) {
	o.DayInMonth = &v
}

// GetWeekInMonth returns the WeekInMonth field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetWeekInMonth() BackupRestoreScheduleWeeks {
	if o == nil || IsNil(o.WeekInMonth) {
		var ret BackupRestoreScheduleWeeks
		return ret
	}
	return *o.WeekInMonth
}

// GetWeekInMonthOk returns a tuple with the WeekInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetWeekInMonthOk() (*BackupRestoreScheduleWeeks, bool) {
	if o == nil || IsNil(o.WeekInMonth) {
		return nil, false
	}
	return o.WeekInMonth, true
}

// HasWeekInMonth returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasWeekInMonth() bool {
	if o != nil && !IsNil(o.WeekInMonth) {
		return true
	}

	return false
}

// SetWeekInMonth gets a reference to the given BackupRestoreScheduleWeeks and assigns it to the WeekInMonth field.
func (o *BackupRestoreScheduleResponseModel) SetWeekInMonth(v BackupRestoreScheduleWeeks) {
	o.WeekInMonth = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *BackupRestoreScheduleResponseModel) SetStartDate(v string) {
	o.StartDate = &v
}

// GetFrequencyFactor returns the FrequencyFactor field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetFrequencyFactor() int32 {
	if o == nil || IsNil(o.FrequencyFactor) {
		var ret int32
		return ret
	}
	return *o.FrequencyFactor
}

// GetFrequencyFactorOk returns a tuple with the FrequencyFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetFrequencyFactorOk() (*int32, bool) {
	if o == nil || IsNil(o.FrequencyFactor) {
		return nil, false
	}
	return o.FrequencyFactor, true
}

// HasFrequencyFactor returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasFrequencyFactor() bool {
	if o != nil && !IsNil(o.FrequencyFactor) {
		return true
	}

	return false
}

// SetFrequencyFactor gets a reference to the given int32 and assigns it to the FrequencyFactor field.
func (o *BackupRestoreScheduleResponseModel) SetFrequencyFactor(v int32) {
	o.FrequencyFactor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BackupRestoreScheduleResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BackupRestoreScheduleResponseModel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetFrequency() BackupRestoreScheduleFrequency {
	if o == nil || IsNil(o.Frequency) {
		var ret BackupRestoreScheduleFrequency
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetFrequencyOk() (*BackupRestoreScheduleFrequency, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given BackupRestoreScheduleFrequency and assigns it to the Frequency field.
func (o *BackupRestoreScheduleResponseModel) SetFrequency(v BackupRestoreScheduleFrequency) {
	o.Frequency = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *BackupRestoreScheduleResponseModel) SetStartTime(v string) {
	o.StartTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *BackupRestoreScheduleResponseModel) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreScheduleResponseModel) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *BackupRestoreScheduleResponseModel) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *BackupRestoreScheduleResponseModel) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o BackupRestoreScheduleResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreScheduleResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	if !IsNil(o.Day) {
		toSerialize["Day"] = o.Day
	}
	if !IsNil(o.DaysInWeek) {
		toSerialize["DaysInWeek"] = o.DaysInWeek
	}
	if !IsNil(o.DayInMonth) {
		toSerialize["DayInMonth"] = o.DayInMonth
	}
	if !IsNil(o.WeekInMonth) {
		toSerialize["WeekInMonth"] = o.WeekInMonth
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.FrequencyFactor) {
		toSerialize["FrequencyFactor"] = o.FrequencyFactor
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Frequency) {
		toSerialize["Frequency"] = o.Frequency
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["TimeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableBackupRestoreScheduleResponseModel struct {
	value *BackupRestoreScheduleResponseModel
	isSet bool
}

func (v NullableBackupRestoreScheduleResponseModel) Get() *BackupRestoreScheduleResponseModel {
	return v.value
}

func (v *NullableBackupRestoreScheduleResponseModel) Set(val *BackupRestoreScheduleResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreScheduleResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreScheduleResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreScheduleResponseModel(val *BackupRestoreScheduleResponseModel) *NullableBackupRestoreScheduleResponseModel {
	return &NullableBackupRestoreScheduleResponseModel{value: val, isSet: true}
}

func (v NullableBackupRestoreScheduleResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreScheduleResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


