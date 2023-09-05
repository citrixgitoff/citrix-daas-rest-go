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

// UserManagementModel User management model to use for a delivery group.
type UserManagementModel string

// List of UserManagementModel
const (
	USERMANAGEMENTMODEL_UNKNOWN UserManagementModel = "Unknown"
	USERMANAGEMENTMODEL_CLOUD_LIBRARY UserManagementModel = "CloudLibrary"
	USERMANAGEMENTMODEL_STUDIO UserManagementModel = "Studio"
)

// All allowed values of UserManagementModel enum
var AllowedUserManagementModelEnumValues = []UserManagementModel{
	"Unknown",
	"CloudLibrary",
	"Studio",
}

func (v *UserManagementModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserManagementModel(value)
	for _, existing := range AllowedUserManagementModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserManagementModel", value)
}

// NewUserManagementModelFromValue returns a pointer to a valid UserManagementModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserManagementModelFromValue(v string) (*UserManagementModel, error) {
	ev := UserManagementModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserManagementModel: valid values are %v", v, AllowedUserManagementModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserManagementModel) IsValid() bool {
	for _, existing := range AllowedUserManagementModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserManagementModel value
func (v UserManagementModel) Ptr() *UserManagementModel {
	return &v
}

type NullableUserManagementModel struct {
	value *UserManagementModel
	isSet bool
}

func (v NullableUserManagementModel) Get() *UserManagementModel {
	return v.value
}

func (v *NullableUserManagementModel) Set(val *UserManagementModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserManagementModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserManagementModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserManagementModel(val *UserManagementModel) *NullableUserManagementModel {
	return &NullableUserManagementModel{value: val, isSet: true}
}

func (v NullableUserManagementModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserManagementModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

