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

// FilterMatchType 
type FilterMatchType string

// List of FilterMatchType
const (
	FILTERMATCHTYPE_UNKNOWN FilterMatchType = "Unknown"
	FILTERMATCHTYPE_MATCH_ANY FilterMatchType = "MatchAny"
	FILTERMATCHTYPE_MATCH_ALL FilterMatchType = "MatchAll"
)

// All allowed values of FilterMatchType enum
var AllowedFilterMatchTypeEnumValues = []FilterMatchType{
	"Unknown",
	"MatchAny",
	"MatchAll",
}

func (v *FilterMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterMatchType(value)
	for _, existing := range AllowedFilterMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterMatchType", value)
}

// NewFilterMatchTypeFromValue returns a pointer to a valid FilterMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterMatchTypeFromValue(v string) (*FilterMatchType, error) {
	ev := FilterMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterMatchType: valid values are %v", v, AllowedFilterMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterMatchType) IsValid() bool {
	for _, existing := range AllowedFilterMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterMatchType value
func (v FilterMatchType) Ptr() *FilterMatchType {
	return &v
}

type NullableFilterMatchType struct {
	value *FilterMatchType
	isSet bool
}

func (v NullableFilterMatchType) Get() *FilterMatchType {
	return v.value
}

func (v *NullableFilterMatchType) Set(val *FilterMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterMatchType(val *FilterMatchType) *NullableFilterMatchType {
	return &NullableFilterMatchType{value: val, isSet: true}
}

func (v NullableFilterMatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

