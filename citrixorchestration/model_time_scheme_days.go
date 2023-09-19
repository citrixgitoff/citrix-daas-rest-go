/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// TimeSchemeDays Time scheme days.
type TimeSchemeDays string

// List of TimeSchemeDays
const (
	TIMESCHEMEDAYS_UNKNOWN TimeSchemeDays = "Unknown"
	TIMESCHEMEDAYS_SUNDAY TimeSchemeDays = "Sunday"
	TIMESCHEMEDAYS_MONDAY TimeSchemeDays = "Monday"
	TIMESCHEMEDAYS_TUESDAY TimeSchemeDays = "Tuesday"
	TIMESCHEMEDAYS_WEDNESDAY TimeSchemeDays = "Wednesday"
	TIMESCHEMEDAYS_THURSDAY TimeSchemeDays = "Thursday"
	TIMESCHEMEDAYS_FRIDAY TimeSchemeDays = "Friday"
	TIMESCHEMEDAYS_SATURDAY TimeSchemeDays = "Saturday"
	TIMESCHEMEDAYS_WEEKDAYS TimeSchemeDays = "Weekdays"
	TIMESCHEMEDAYS_WEEKEND TimeSchemeDays = "Weekend"
)

// All allowed values of TimeSchemeDays enum
var AllowedTimeSchemeDaysEnumValues = []TimeSchemeDays{
	"Unknown",
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Weekdays",
	"Weekend",
}

func (v *TimeSchemeDays) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeSchemeDays(value)
	for _, existing := range AllowedTimeSchemeDaysEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeSchemeDays", value)
}

// NewTimeSchemeDaysFromValue returns a pointer to a valid TimeSchemeDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeSchemeDaysFromValue(v string) (*TimeSchemeDays, error) {
	ev := TimeSchemeDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeSchemeDays: valid values are %v", v, AllowedTimeSchemeDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeSchemeDays) IsValid() bool {
	for _, existing := range AllowedTimeSchemeDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeSchemeDays value
func (v TimeSchemeDays) Ptr() *TimeSchemeDays {
	return &v
}

type NullableTimeSchemeDays struct {
	value *TimeSchemeDays
	isSet bool
}

func (v NullableTimeSchemeDays) Get() *TimeSchemeDays {
	return v.value
}

func (v *NullableTimeSchemeDays) Set(val *TimeSchemeDays) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSchemeDays) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSchemeDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSchemeDays(val *TimeSchemeDays) *NullableTimeSchemeDays {
	return &NullableTimeSchemeDays{value: val, isSet: true}
}

func (v NullableTimeSchemeDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSchemeDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

