/*
Administrators APIs

APIs for managing CC administrators.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
	"fmt"
)

// AdministratorAccessType the model 'AdministratorAccessType'
type AdministratorAccessType string

// List of AdministratorAccessType
const (
	ADMINISTRATORACCESSTYPE_FULL AdministratorAccessType = "Full"
	ADMINISTRATORACCESSTYPE_CUSTOM AdministratorAccessType = "Custom"
	ADMINISTRATORACCESSTYPE_REINVITE AdministratorAccessType = "Reinvite"
)

// All allowed values of AdministratorAccessType enum
var AllowedAdministratorAccessTypeEnumValues = []AdministratorAccessType{
	"Full",
	"Custom",
	"Reinvite",
}

func (v *AdministratorAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdministratorAccessType(value)
	for _, existing := range AllowedAdministratorAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdministratorAccessType", value)
}

// NewAdministratorAccessTypeFromValue returns a pointer to a valid AdministratorAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdministratorAccessTypeFromValue(v string) (*AdministratorAccessType, error) {
	ev := AdministratorAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdministratorAccessType: valid values are %v", v, AllowedAdministratorAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdministratorAccessType) IsValid() bool {
	for _, existing := range AllowedAdministratorAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdministratorAccessType value
func (v AdministratorAccessType) Ptr() *AdministratorAccessType {
	return &v
}

type NullableAdministratorAccessType struct {
	value *AdministratorAccessType
	isSet bool
}

func (v NullableAdministratorAccessType) Get() *AdministratorAccessType {
	return v.value
}

func (v *NullableAdministratorAccessType) Set(val *AdministratorAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorAccessType(val *AdministratorAccessType) *NullableAdministratorAccessType {
	return &NullableAdministratorAccessType{value: val, isSet: true}
}

func (v NullableAdministratorAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

