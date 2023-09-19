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

// IdentityUserType The type of User identity object to retrieve.
type IdentityUserType string

// List of IdentityUserType
const (
	IDENTITYUSERTYPE_UNKNOWN IdentityUserType = "Unknown"
	IDENTITYUSERTYPE_USER IdentityUserType = "User"
	IDENTITYUSERTYPE_GROUP IdentityUserType = "Group"
	IDENTITYUSERTYPE_ALL IdentityUserType = "All"
)

// All allowed values of IdentityUserType enum
var AllowedIdentityUserTypeEnumValues = []IdentityUserType{
	"Unknown",
	"User",
	"Group",
	"All",
}

func (v *IdentityUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityUserType(value)
	for _, existing := range AllowedIdentityUserTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityUserType", value)
}

// NewIdentityUserTypeFromValue returns a pointer to a valid IdentityUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityUserTypeFromValue(v string) (*IdentityUserType, error) {
	ev := IdentityUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityUserType: valid values are %v", v, AllowedIdentityUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityUserType) IsValid() bool {
	for _, existing := range AllowedIdentityUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityUserType value
func (v IdentityUserType) Ptr() *IdentityUserType {
	return &v
}

type NullableIdentityUserType struct {
	value *IdentityUserType
	isSet bool
}

func (v NullableIdentityUserType) Get() *IdentityUserType {
	return v.value
}

func (v *NullableIdentityUserType) Set(val *IdentityUserType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserType(val *IdentityUserType) *NullableIdentityUserType {
	return &NullableIdentityUserType{value: val, isSet: true}
}

func (v NullableIdentityUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

