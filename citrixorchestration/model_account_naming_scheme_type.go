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

// AccountNamingSchemeType 
type AccountNamingSchemeType string

// List of AccountNamingSchemeType
const (
	ACCOUNTNAMINGSCHEMETYPE_UNKNOWN AccountNamingSchemeType = "Unknown"
	ACCOUNTNAMINGSCHEMETYPE_NONE AccountNamingSchemeType = "None"
	ACCOUNTNAMINGSCHEMETYPE_ALPHABETIC AccountNamingSchemeType = "Alphabetic"
	ACCOUNTNAMINGSCHEMETYPE_NUMERIC AccountNamingSchemeType = "Numeric"
)

// All allowed values of AccountNamingSchemeType enum
var AllowedAccountNamingSchemeTypeEnumValues = []AccountNamingSchemeType{
	"Unknown",
	"None",
	"Alphabetic",
	"Numeric",
}

func (v *AccountNamingSchemeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountNamingSchemeType(value)
	for _, existing := range AllowedAccountNamingSchemeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountNamingSchemeType", value)
}

// NewAccountNamingSchemeTypeFromValue returns a pointer to a valid AccountNamingSchemeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountNamingSchemeTypeFromValue(v string) (*AccountNamingSchemeType, error) {
	ev := AccountNamingSchemeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountNamingSchemeType: valid values are %v", v, AllowedAccountNamingSchemeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountNamingSchemeType) IsValid() bool {
	for _, existing := range AllowedAccountNamingSchemeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountNamingSchemeType value
func (v AccountNamingSchemeType) Ptr() *AccountNamingSchemeType {
	return &v
}

type NullableAccountNamingSchemeType struct {
	value *AccountNamingSchemeType
	isSet bool
}

func (v NullableAccountNamingSchemeType) Get() *AccountNamingSchemeType {
	return v.value
}

func (v *NullableAccountNamingSchemeType) Set(val *AccountNamingSchemeType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNamingSchemeType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNamingSchemeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNamingSchemeType(val *AccountNamingSchemeType) *NullableAccountNamingSchemeType {
	return &NullableAccountNamingSchemeType{value: val, isSet: true}
}

func (v NullableAccountNamingSchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNamingSchemeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

