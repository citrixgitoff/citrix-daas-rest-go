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

// WindowsConnectionSetting Windows connection setting.
type WindowsConnectionSetting string

// List of WindowsConnectionSetting
const (
	WINDOWSCONNECTIONSETTING_UNKNOWN WindowsConnectionSetting = "Unknown"
	WINDOWSCONNECTIONSETTING_LOGON_ENABLED WindowsConnectionSetting = "LogonEnabled"
	WINDOWSCONNECTIONSETTING_DRAINING WindowsConnectionSetting = "Draining"
	WINDOWSCONNECTIONSETTING_DRAINING_UNTIL_RESTART WindowsConnectionSetting = "DrainingUntilRestart"
	WINDOWSCONNECTIONSETTING_LOGON_DISABLED WindowsConnectionSetting = "LogonDisabled"
)

// All allowed values of WindowsConnectionSetting enum
var AllowedWindowsConnectionSettingEnumValues = []WindowsConnectionSetting{
	"Unknown",
	"LogonEnabled",
	"Draining",
	"DrainingUntilRestart",
	"LogonDisabled",
}

func (v *WindowsConnectionSetting) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WindowsConnectionSetting(value)
	for _, existing := range AllowedWindowsConnectionSettingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WindowsConnectionSetting", value)
}

// NewWindowsConnectionSettingFromValue returns a pointer to a valid WindowsConnectionSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWindowsConnectionSettingFromValue(v string) (*WindowsConnectionSetting, error) {
	ev := WindowsConnectionSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WindowsConnectionSetting: valid values are %v", v, AllowedWindowsConnectionSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WindowsConnectionSetting) IsValid() bool {
	for _, existing := range AllowedWindowsConnectionSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WindowsConnectionSetting value
func (v WindowsConnectionSetting) Ptr() *WindowsConnectionSetting {
	return &v
}

type NullableWindowsConnectionSetting struct {
	value *WindowsConnectionSetting
	isSet bool
}

func (v NullableWindowsConnectionSetting) Get() *WindowsConnectionSetting {
	return v.value
}

func (v *NullableWindowsConnectionSetting) Set(val *WindowsConnectionSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsConnectionSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsConnectionSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsConnectionSetting(val *WindowsConnectionSetting) *NullableWindowsConnectionSetting {
	return &NullableWindowsConnectionSetting{value: val, isSet: true}
}

func (v NullableWindowsConnectionSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsConnectionSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

