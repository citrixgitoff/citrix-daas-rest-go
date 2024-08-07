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

// LogoffOption Logoff options for VDA upgrade.             
type LogoffOption string

// List of LogoffOption
const (
	LOGOFFOPTION_DO_NOT_LOGOFF LogoffOption = "DoNotLogoff"
	LOGOFFOPTION_ACTIVE_SESSIONS_ONLY LogoffOption = "ActiveSessionsOnly"
	LOGOFFOPTION_DISCONNECTED_SESSIONS_ONLY LogoffOption = "DisconnectedSessionsOnly"
	LOGOFFOPTION_ACTIVE_AND_DISCONNECTED_SESSIONS LogoffOption = "ActiveAndDisconnectedSessions"
)

// All allowed values of LogoffOption enum
var AllowedLogoffOptionEnumValues = []LogoffOption{
	"DoNotLogoff",
	"ActiveSessionsOnly",
	"DisconnectedSessionsOnly",
	"ActiveAndDisconnectedSessions",
}

func (v *LogoffOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogoffOption(value)
	for _, existing := range AllowedLogoffOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogoffOption", value)
}

// NewLogoffOptionFromValue returns a pointer to a valid LogoffOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogoffOptionFromValue(v string) (*LogoffOption, error) {
	ev := LogoffOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogoffOption: valid values are %v", v, AllowedLogoffOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogoffOption) IsValid() bool {
	for _, existing := range AllowedLogoffOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogoffOption value
func (v LogoffOption) Ptr() *LogoffOption {
	return &v
}

type NullableLogoffOption struct {
	value *LogoffOption
	isSet bool
}

func (v NullableLogoffOption) Get() *LogoffOption {
	return v.value
}

func (v *NullableLogoffOption) Set(val *LogoffOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoffOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoffOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoffOption(val *LogoffOption) *NullableLogoffOption {
	return &NullableLogoffOption{value: val, isSet: true}
}

func (v NullableLogoffOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoffOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

