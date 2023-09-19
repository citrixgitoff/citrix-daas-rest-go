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

// HomeZoneMode Modes by which the system may choose to use a user's or application's home zone.
type HomeZoneMode string

// List of HomeZoneMode
const (
	HOMEZONEMODE_UNKNOWN HomeZoneMode = "Unknown"
	HOMEZONEMODE_PREFER HomeZoneMode = "Prefer"
	HOMEZONEMODE_IGNORE HomeZoneMode = "Ignore"
	HOMEZONEMODE_ONLY HomeZoneMode = "Only"
	HOMEZONEMODE_USER HomeZoneMode = "User"
)

// All allowed values of HomeZoneMode enum
var AllowedHomeZoneModeEnumValues = []HomeZoneMode{
	"Unknown",
	"Prefer",
	"Ignore",
	"Only",
	"User",
}

func (v *HomeZoneMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HomeZoneMode(value)
	for _, existing := range AllowedHomeZoneModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HomeZoneMode", value)
}

// NewHomeZoneModeFromValue returns a pointer to a valid HomeZoneMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHomeZoneModeFromValue(v string) (*HomeZoneMode, error) {
	ev := HomeZoneMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HomeZoneMode: valid values are %v", v, AllowedHomeZoneModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HomeZoneMode) IsValid() bool {
	for _, existing := range AllowedHomeZoneModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HomeZoneMode value
func (v HomeZoneMode) Ptr() *HomeZoneMode {
	return &v
}

type NullableHomeZoneMode struct {
	value *HomeZoneMode
	isSet bool
}

func (v NullableHomeZoneMode) Get() *HomeZoneMode {
	return v.value
}

func (v *NullableHomeZoneMode) Set(val *HomeZoneMode) {
	v.value = val
	v.isSet = true
}

func (v NullableHomeZoneMode) IsSet() bool {
	return v.isSet
}

func (v *NullableHomeZoneMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHomeZoneMode(val *HomeZoneMode) *NullableHomeZoneMode {
	return &NullableHomeZoneMode{value: val, isSet: true}
}

func (v NullableHomeZoneMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHomeZoneMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

