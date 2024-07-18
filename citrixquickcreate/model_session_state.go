/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"fmt"
)

// SessionState Session states.
type SessionState string

// List of SessionState
const (
	SESSIONSTATE_UNKNOWN SessionState = "Unknown"
	SESSIONSTATE_CONNECTED SessionState = "Connected"
	SESSIONSTATE_ACTIVE SessionState = "Active"
	SESSIONSTATE_DISCONNECTED SessionState = "Disconnected"
)

// All allowed values of SessionState enum
var AllowedSessionStateEnumValues = []SessionState{
	"Unknown",
	"Connected",
	"Active",
	"Disconnected",
}

func (v *SessionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionState(value)
	for _, existing := range AllowedSessionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionState", value)
}

// NewSessionStateFromValue returns a pointer to a valid SessionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionStateFromValue(v string) (*SessionState, error) {
	ev := SessionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionState: valid values are %v", v, AllowedSessionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionState) IsValid() bool {
	for _, existing := range AllowedSessionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionState value
func (v SessionState) Ptr() *SessionState {
	return &v
}

type NullableSessionState struct {
	value *SessionState
	isSet bool
}

func (v NullableSessionState) Get() *SessionState {
	return v.value
}

func (v *NullableSessionState) Set(val *SessionState) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionState) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionState(val *SessionState) *NullableSessionState {
	return &NullableSessionState{value: val, isSet: true}
}

func (v NullableSessionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

