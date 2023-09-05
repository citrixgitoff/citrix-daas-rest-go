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

// ProvSchemeWarningType 
type ProvSchemeWarningType string

// List of ProvSchemeWarningType
const (
	PROVSCHEMEWARNINGTYPE_UNKNOWN ProvSchemeWarningType = "Unknown"
	PROVSCHEMEWARNINGTYPE_GET_CUSTOM_PROPERTIES_FAILED ProvSchemeWarningType = "GetCustomPropertiesFailed"
	PROVSCHEMEWARNINGTYPE_GET_VM_METADATA_FAILED ProvSchemeWarningType = "GetVMMetadataFailed"
)

// All allowed values of ProvSchemeWarningType enum
var AllowedProvSchemeWarningTypeEnumValues = []ProvSchemeWarningType{
	"Unknown",
	"GetCustomPropertiesFailed",
	"GetVMMetadataFailed",
}

func (v *ProvSchemeWarningType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvSchemeWarningType(value)
	for _, existing := range AllowedProvSchemeWarningTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvSchemeWarningType", value)
}

// NewProvSchemeWarningTypeFromValue returns a pointer to a valid ProvSchemeWarningType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvSchemeWarningTypeFromValue(v string) (*ProvSchemeWarningType, error) {
	ev := ProvSchemeWarningType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvSchemeWarningType: valid values are %v", v, AllowedProvSchemeWarningTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvSchemeWarningType) IsValid() bool {
	for _, existing := range AllowedProvSchemeWarningTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvSchemeWarningType value
func (v ProvSchemeWarningType) Ptr() *ProvSchemeWarningType {
	return &v
}

type NullableProvSchemeWarningType struct {
	value *ProvSchemeWarningType
	isSet bool
}

func (v NullableProvSchemeWarningType) Get() *ProvSchemeWarningType {
	return v.value
}

func (v *NullableProvSchemeWarningType) Set(val *ProvSchemeWarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableProvSchemeWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableProvSchemeWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvSchemeWarningType(val *ProvSchemeWarningType) *NullableProvSchemeWarningType {
	return &NullableProvSchemeWarningType{value: val, isSet: true}
}

func (v NullableProvSchemeWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvSchemeWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

