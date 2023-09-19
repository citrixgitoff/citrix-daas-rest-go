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

// ApplicationDiscoveryItemType Type to indicate application discovery item type
type ApplicationDiscoveryItemType string

// List of ApplicationDiscoveryItemType
const (
	APPLICATIONDISCOVERYITEMTYPE_UNKNOWN ApplicationDiscoveryItemType = "Unknown"
	APPLICATIONDISCOVERYITEMTYPE_MACHINE_SHARES ApplicationDiscoveryItemType = "MachineShares"
	APPLICATIONDISCOVERYITEMTYPE_DIRECTORY_OR_FILES ApplicationDiscoveryItemType = "DirectoryOrFiles"
	APPLICATIONDISCOVERYITEMTYPE_APPLICATION_PROPERTY ApplicationDiscoveryItemType = "ApplicationProperty"
)

// All allowed values of ApplicationDiscoveryItemType enum
var AllowedApplicationDiscoveryItemTypeEnumValues = []ApplicationDiscoveryItemType{
	"Unknown",
	"MachineShares",
	"DirectoryOrFiles",
	"ApplicationProperty",
}

func (v *ApplicationDiscoveryItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationDiscoveryItemType(value)
	for _, existing := range AllowedApplicationDiscoveryItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationDiscoveryItemType", value)
}

// NewApplicationDiscoveryItemTypeFromValue returns a pointer to a valid ApplicationDiscoveryItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationDiscoveryItemTypeFromValue(v string) (*ApplicationDiscoveryItemType, error) {
	ev := ApplicationDiscoveryItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationDiscoveryItemType: valid values are %v", v, AllowedApplicationDiscoveryItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationDiscoveryItemType) IsValid() bool {
	for _, existing := range AllowedApplicationDiscoveryItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationDiscoveryItemType value
func (v ApplicationDiscoveryItemType) Ptr() *ApplicationDiscoveryItemType {
	return &v
}

type NullableApplicationDiscoveryItemType struct {
	value *ApplicationDiscoveryItemType
	isSet bool
}

func (v NullableApplicationDiscoveryItemType) Get() *ApplicationDiscoveryItemType {
	return v.value
}

func (v *NullableApplicationDiscoveryItemType) Set(val *ApplicationDiscoveryItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDiscoveryItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDiscoveryItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDiscoveryItemType(val *ApplicationDiscoveryItemType) *NullableApplicationDiscoveryItemType {
	return &NullableApplicationDiscoveryItemType{value: val, isSet: true}
}

func (v NullableApplicationDiscoveryItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDiscoveryItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

