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

// MachineAndSessionSearchFilterGroupType The search filter group type
type MachineAndSessionSearchFilterGroupType string

// List of MachineAndSessionSearchFilterGroupType
const (
	MACHINEANDSESSIONSEARCHFILTERGROUPTYPE_OR MachineAndSessionSearchFilterGroupType = "Or"
	MACHINEANDSESSIONSEARCHFILTERGROUPTYPE_AND MachineAndSessionSearchFilterGroupType = "And"
)

// All allowed values of MachineAndSessionSearchFilterGroupType enum
var AllowedMachineAndSessionSearchFilterGroupTypeEnumValues = []MachineAndSessionSearchFilterGroupType{
	"Or",
	"And",
}

func (v *MachineAndSessionSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MachineAndSessionSearchFilterGroupType(value)
	for _, existing := range AllowedMachineAndSessionSearchFilterGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MachineAndSessionSearchFilterGroupType", value)
}

// NewMachineAndSessionSearchFilterGroupTypeFromValue returns a pointer to a valid MachineAndSessionSearchFilterGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineAndSessionSearchFilterGroupTypeFromValue(v string) (*MachineAndSessionSearchFilterGroupType, error) {
	ev := MachineAndSessionSearchFilterGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineAndSessionSearchFilterGroupType: valid values are %v", v, AllowedMachineAndSessionSearchFilterGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineAndSessionSearchFilterGroupType) IsValid() bool {
	for _, existing := range AllowedMachineAndSessionSearchFilterGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineAndSessionSearchFilterGroupType value
func (v MachineAndSessionSearchFilterGroupType) Ptr() *MachineAndSessionSearchFilterGroupType {
	return &v
}

type NullableMachineAndSessionSearchFilterGroupType struct {
	value *MachineAndSessionSearchFilterGroupType
	isSet bool
}

func (v NullableMachineAndSessionSearchFilterGroupType) Get() *MachineAndSessionSearchFilterGroupType {
	return v.value
}

func (v *NullableMachineAndSessionSearchFilterGroupType) Set(val *MachineAndSessionSearchFilterGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAndSessionSearchFilterGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAndSessionSearchFilterGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAndSessionSearchFilterGroupType(val *MachineAndSessionSearchFilterGroupType) *NullableMachineAndSessionSearchFilterGroupType {
	return &NullableMachineAndSessionSearchFilterGroupType{value: val, isSet: true}
}

func (v NullableMachineAndSessionSearchFilterGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAndSessionSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

