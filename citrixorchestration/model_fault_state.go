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

// FaultState Fault states for machines.
type FaultState string

// List of FaultState
const (
	FAULTSTATE_UNKNOWN FaultState = "Unknown"
	FAULTSTATE_NONE FaultState = "None"
	FAULTSTATE_FAILED_TO_START FaultState = "FailedToStart"
	FAULTSTATE_STUCK_ON_BOOT FaultState = "StuckOnBoot"
	FAULTSTATE_UNREGISTERED FaultState = "Unregistered"
	FAULTSTATE_MAX_CAPACITY FaultState = "MaxCapacity"
)

// All allowed values of FaultState enum
var AllowedFaultStateEnumValues = []FaultState{
	"Unknown",
	"None",
	"FailedToStart",
	"StuckOnBoot",
	"Unregistered",
	"MaxCapacity",
}

func (v *FaultState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FaultState(value)
	for _, existing := range AllowedFaultStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FaultState", value)
}

// NewFaultStateFromValue returns a pointer to a valid FaultState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFaultStateFromValue(v string) (*FaultState, error) {
	ev := FaultState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FaultState: valid values are %v", v, AllowedFaultStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaultState) IsValid() bool {
	for _, existing := range AllowedFaultStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FaultState value
func (v FaultState) Ptr() *FaultState {
	return &v
}

type NullableFaultState struct {
	value *FaultState
	isSet bool
}

func (v NullableFaultState) Get() *FaultState {
	return v.value
}

func (v *NullableFaultState) Set(val *FaultState) {
	v.value = val
	v.isSet = true
}

func (v NullableFaultState) IsSet() bool {
	return v.isSet
}

func (v *NullableFaultState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaultState(val *FaultState) *NullableFaultState {
	return &NullableFaultState{value: val, isSet: true}
}

func (v NullableFaultState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaultState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

