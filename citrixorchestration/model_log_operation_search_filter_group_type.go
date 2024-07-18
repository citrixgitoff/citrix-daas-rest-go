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

// LogOperationSearchFilterGroupType The group type of advanced search filters for configuration logs.
type LogOperationSearchFilterGroupType string

// List of LogOperationSearchFilterGroupType
const (
	LOGOPERATIONSEARCHFILTERGROUPTYPE_OR LogOperationSearchFilterGroupType = "Or"
	LOGOPERATIONSEARCHFILTERGROUPTYPE_AND LogOperationSearchFilterGroupType = "And"
)

// All allowed values of LogOperationSearchFilterGroupType enum
var AllowedLogOperationSearchFilterGroupTypeEnumValues = []LogOperationSearchFilterGroupType{
	"Or",
	"And",
}

func (v *LogOperationSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogOperationSearchFilterGroupType(value)
	for _, existing := range AllowedLogOperationSearchFilterGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogOperationSearchFilterGroupType", value)
}

// NewLogOperationSearchFilterGroupTypeFromValue returns a pointer to a valid LogOperationSearchFilterGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogOperationSearchFilterGroupTypeFromValue(v string) (*LogOperationSearchFilterGroupType, error) {
	ev := LogOperationSearchFilterGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogOperationSearchFilterGroupType: valid values are %v", v, AllowedLogOperationSearchFilterGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogOperationSearchFilterGroupType) IsValid() bool {
	for _, existing := range AllowedLogOperationSearchFilterGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogOperationSearchFilterGroupType value
func (v LogOperationSearchFilterGroupType) Ptr() *LogOperationSearchFilterGroupType {
	return &v
}

type NullableLogOperationSearchFilterGroupType struct {
	value *LogOperationSearchFilterGroupType
	isSet bool
}

func (v NullableLogOperationSearchFilterGroupType) Get() *LogOperationSearchFilterGroupType {
	return v.value
}

func (v *NullableLogOperationSearchFilterGroupType) Set(val *LogOperationSearchFilterGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogOperationSearchFilterGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogOperationSearchFilterGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogOperationSearchFilterGroupType(val *LogOperationSearchFilterGroupType) *NullableLogOperationSearchFilterGroupType {
	return &NullableLogOperationSearchFilterGroupType{value: val, isSet: true}
}

func (v NullableLogOperationSearchFilterGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogOperationSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

