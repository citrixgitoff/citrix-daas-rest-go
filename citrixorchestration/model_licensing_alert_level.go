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

// LicensingAlertLevel 
type LicensingAlertLevel string

// List of LicensingAlertLevel
const (
	LICENSINGALERTLEVEL_UNKNOWN LicensingAlertLevel = "Unknown"
	LICENSINGALERTLEVEL_INFO LicensingAlertLevel = "Info"
	LICENSINGALERTLEVEL_ALERT LicensingAlertLevel = "Alert"
	LICENSINGALERTLEVEL_ALARM LicensingAlertLevel = "Alarm"
)

// All allowed values of LicensingAlertLevel enum
var AllowedLicensingAlertLevelEnumValues = []LicensingAlertLevel{
	"Unknown",
	"Info",
	"Alert",
	"Alarm",
}

func (v *LicensingAlertLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicensingAlertLevel(value)
	for _, existing := range AllowedLicensingAlertLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicensingAlertLevel", value)
}

// NewLicensingAlertLevelFromValue returns a pointer to a valid LicensingAlertLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicensingAlertLevelFromValue(v string) (*LicensingAlertLevel, error) {
	ev := LicensingAlertLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicensingAlertLevel: valid values are %v", v, AllowedLicensingAlertLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicensingAlertLevel) IsValid() bool {
	for _, existing := range AllowedLicensingAlertLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicensingAlertLevel value
func (v LicensingAlertLevel) Ptr() *LicensingAlertLevel {
	return &v
}

type NullableLicensingAlertLevel struct {
	value *LicensingAlertLevel
	isSet bool
}

func (v NullableLicensingAlertLevel) Get() *LicensingAlertLevel {
	return v.value
}

func (v *NullableLicensingAlertLevel) Set(val *LicensingAlertLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingAlertLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingAlertLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingAlertLevel(val *LicensingAlertLevel) *NullableLicensingAlertLevel {
	return &NullableLicensingAlertLevel{value: val, isSet: true}
}

func (v NullableLicensingAlertLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingAlertLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

