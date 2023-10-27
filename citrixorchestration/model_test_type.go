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

// TestType Defines test type.
type TestType string

// List of TestType
const (
	TESTTYPE_SITE TestType = "Site"
	TESTTYPE_DELIVERY_GROUP TestType = "DeliveryGroup"
	TESTTYPE_MACHINE_CATALOG TestType = "MachineCatalog"
	TESTTYPE_RESOURCE_POOL TestType = "ResourcePool"
	TESTTYPE_HYPERVISOR TestType = "Hypervisor"
	TESTTYPE_MACHINE TestType = "Machine"
	TESTTYPE_UNKNOWN TestType = "Unknown"
)

// All allowed values of TestType enum
var AllowedTestTypeEnumValues = []TestType{
	"Site",
	"DeliveryGroup",
	"MachineCatalog",
	"ResourcePool",
	"Hypervisor",
	"Machine",
	"Unknown",
}

func (v *TestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestType(value)
	for _, existing := range AllowedTestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestType", value)
}

// NewTestTypeFromValue returns a pointer to a valid TestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestTypeFromValue(v string) (*TestType, error) {
	ev := TestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestType: valid values are %v", v, AllowedTestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestType) IsValid() bool {
	for _, existing := range AllowedTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestType value
func (v TestType) Ptr() *TestType {
	return &v
}

type NullableTestType struct {
	value *TestType
	isSet bool
}

func (v NullableTestType) Get() *TestType {
	return v.value
}

func (v *NullableTestType) Set(val *TestType) {
	v.value = val
	v.isSet = true
}

func (v NullableTestType) IsSet() bool {
	return v.isSet
}

func (v *NullableTestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestType(val *TestType) *NullableTestType {
	return &NullableTestType{value: val, isSet: true}
}

func (v NullableTestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

