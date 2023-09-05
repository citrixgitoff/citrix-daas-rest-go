/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// ScheduledReboot Scheduled reboot states.
type ScheduledReboot string

// List of ScheduledReboot
const (
	SCHEDULEDREBOOT_UNKNOWN ScheduledReboot = "Unknown"
	SCHEDULEDREBOOT_NONE ScheduledReboot = "None"
	SCHEDULEDREBOOT_PENDING ScheduledReboot = "Pending"
	SCHEDULEDREBOOT_DRAINING ScheduledReboot = "Draining"
	SCHEDULEDREBOOT_IN_PROGRESS ScheduledReboot = "InProgress"
	SCHEDULEDREBOOT_NATURAL ScheduledReboot = "Natural"
)

// All allowed values of ScheduledReboot enum
var AllowedScheduledRebootEnumValues = []ScheduledReboot{
	"Unknown",
	"None",
	"Pending",
	"Draining",
	"InProgress",
	"Natural",
}

func (v *ScheduledReboot) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduledReboot(value)
	for _, existing := range AllowedScheduledRebootEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduledReboot", value)
}

// NewScheduledRebootFromValue returns a pointer to a valid ScheduledReboot
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduledRebootFromValue(v string) (*ScheduledReboot, error) {
	ev := ScheduledReboot(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduledReboot: valid values are %v", v, AllowedScheduledRebootEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduledReboot) IsValid() bool {
	for _, existing := range AllowedScheduledRebootEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledReboot value
func (v ScheduledReboot) Ptr() *ScheduledReboot {
	return &v
}

type NullableScheduledReboot struct {
	value *ScheduledReboot
	isSet bool
}

func (v NullableScheduledReboot) Get() *ScheduledReboot {
	return v.value
}

func (v *NullableScheduledReboot) Set(val *ScheduledReboot) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledReboot) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledReboot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledReboot(val *ScheduledReboot) *NullableScheduledReboot {
	return &NullableScheduledReboot{value: val, isSet: true}
}

func (v NullableScheduledReboot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledReboot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

