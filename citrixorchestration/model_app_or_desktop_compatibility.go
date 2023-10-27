/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// AppOrDesktopCompatibility Compatibility set for Apps or Desktops.             
type AppOrDesktopCompatibility string

// List of AppOrDesktopCompatibility
const (
	APPORDESKTOPCOMPATIBILITY_UNKNOWN AppOrDesktopCompatibility = "Unknown"
	APPORDESKTOPCOMPATIBILITY_COMPATIBLE AppOrDesktopCompatibility = "Compatible"
	APPORDESKTOPCOMPATIBILITY_INCOMPATIBLE_DELIVERY_TYPE AppOrDesktopCompatibility = "IncompatibleDeliveryType"
)

// All allowed values of AppOrDesktopCompatibility enum
var AllowedAppOrDesktopCompatibilityEnumValues = []AppOrDesktopCompatibility{
	"Unknown",
	"Compatible",
	"IncompatibleDeliveryType",
}

func (v *AppOrDesktopCompatibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppOrDesktopCompatibility(value)
	for _, existing := range AllowedAppOrDesktopCompatibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppOrDesktopCompatibility", value)
}

// NewAppOrDesktopCompatibilityFromValue returns a pointer to a valid AppOrDesktopCompatibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppOrDesktopCompatibilityFromValue(v string) (*AppOrDesktopCompatibility, error) {
	ev := AppOrDesktopCompatibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppOrDesktopCompatibility: valid values are %v", v, AllowedAppOrDesktopCompatibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppOrDesktopCompatibility) IsValid() bool {
	for _, existing := range AllowedAppOrDesktopCompatibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppOrDesktopCompatibility value
func (v AppOrDesktopCompatibility) Ptr() *AppOrDesktopCompatibility {
	return &v
}

type NullableAppOrDesktopCompatibility struct {
	value *AppOrDesktopCompatibility
	isSet bool
}

func (v NullableAppOrDesktopCompatibility) Get() *AppOrDesktopCompatibility {
	return v.value
}

func (v *NullableAppOrDesktopCompatibility) Set(val *AppOrDesktopCompatibility) {
	v.value = val
	v.isSet = true
}

func (v NullableAppOrDesktopCompatibility) IsSet() bool {
	return v.isSet
}

func (v *NullableAppOrDesktopCompatibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppOrDesktopCompatibility(val *AppOrDesktopCompatibility) *NullableAppOrDesktopCompatibility {
	return &NullableAppOrDesktopCompatibility{value: val, isSet: true}
}

func (v NullableAppOrDesktopCompatibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppOrDesktopCompatibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

