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

// LicensingPermissionLevel 
type LicensingPermissionLevel string

// List of LicensingPermissionLevel
const (
	LICENSINGPERMISSIONLEVEL_UNKNOWN LicensingPermissionLevel = "Unknown"
	LICENSINGPERMISSIONLEVEL_READ_ONLY LicensingPermissionLevel = "ReadOnly"
	LICENSINGPERMISSIONLEVEL_FULL LicensingPermissionLevel = "Full"
	LICENSINGPERMISSIONLEVEL_NONE LicensingPermissionLevel = "None"
)

// All allowed values of LicensingPermissionLevel enum
var AllowedLicensingPermissionLevelEnumValues = []LicensingPermissionLevel{
	"Unknown",
	"ReadOnly",
	"Full",
	"None",
}

func (v *LicensingPermissionLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicensingPermissionLevel(value)
	for _, existing := range AllowedLicensingPermissionLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicensingPermissionLevel", value)
}

// NewLicensingPermissionLevelFromValue returns a pointer to a valid LicensingPermissionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicensingPermissionLevelFromValue(v string) (*LicensingPermissionLevel, error) {
	ev := LicensingPermissionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicensingPermissionLevel: valid values are %v", v, AllowedLicensingPermissionLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicensingPermissionLevel) IsValid() bool {
	for _, existing := range AllowedLicensingPermissionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicensingPermissionLevel value
func (v LicensingPermissionLevel) Ptr() *LicensingPermissionLevel {
	return &v
}

type NullableLicensingPermissionLevel struct {
	value *LicensingPermissionLevel
	isSet bool
}

func (v NullableLicensingPermissionLevel) Get() *LicensingPermissionLevel {
	return v.value
}

func (v *NullableLicensingPermissionLevel) Set(val *LicensingPermissionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingPermissionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingPermissionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingPermissionLevel(val *LicensingPermissionLevel) *NullableLicensingPermissionLevel {
	return &NullableLicensingPermissionLevel{value: val, isSet: true}
}

func (v NullableLicensingPermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingPermissionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

