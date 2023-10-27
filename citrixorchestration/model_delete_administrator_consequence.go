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

// DeleteAdministratorConsequence 
type DeleteAdministratorConsequence string

// List of DeleteAdministratorConsequence
const (
	DELETEADMINISTRATORCONSEQUENCE_UNKNOWN DeleteAdministratorConsequence = "Unknown"
	DELETEADMINISTRATORCONSEQUENCE_NO_RIGHTS_LOST DeleteAdministratorConsequence = "NoRightsLost"
	DELETEADMINISTRATORCONSEQUENCE_SOME_RIGHTS_LOST DeleteAdministratorConsequence = "SomeRightsLost"
	DELETEADMINISTRATORCONSEQUENCE_ALL_RIGHTS_LOST DeleteAdministratorConsequence = "AllRightsLost"
	DELETEADMINISTRATORCONSEQUENCE_CANNOT_DELETE_LAST_FULL_ADMIN DeleteAdministratorConsequence = "CannotDeleteLastFullAdmin"
)

// All allowed values of DeleteAdministratorConsequence enum
var AllowedDeleteAdministratorConsequenceEnumValues = []DeleteAdministratorConsequence{
	"Unknown",
	"NoRightsLost",
	"SomeRightsLost",
	"AllRightsLost",
	"CannotDeleteLastFullAdmin",
}

func (v *DeleteAdministratorConsequence) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeleteAdministratorConsequence(value)
	for _, existing := range AllowedDeleteAdministratorConsequenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeleteAdministratorConsequence", value)
}

// NewDeleteAdministratorConsequenceFromValue returns a pointer to a valid DeleteAdministratorConsequence
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeleteAdministratorConsequenceFromValue(v string) (*DeleteAdministratorConsequence, error) {
	ev := DeleteAdministratorConsequence(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeleteAdministratorConsequence: valid values are %v", v, AllowedDeleteAdministratorConsequenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeleteAdministratorConsequence) IsValid() bool {
	for _, existing := range AllowedDeleteAdministratorConsequenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteAdministratorConsequence value
func (v DeleteAdministratorConsequence) Ptr() *DeleteAdministratorConsequence {
	return &v
}

type NullableDeleteAdministratorConsequence struct {
	value *DeleteAdministratorConsequence
	isSet bool
}

func (v NullableDeleteAdministratorConsequence) Get() *DeleteAdministratorConsequence {
	return v.value
}

func (v *NullableDeleteAdministratorConsequence) Set(val *DeleteAdministratorConsequence) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAdministratorConsequence) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAdministratorConsequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAdministratorConsequence(val *DeleteAdministratorConsequence) *NullableDeleteAdministratorConsequence {
	return &NullableDeleteAdministratorConsequence{value: val, isSet: true}
}

func (v NullableDeleteAdministratorConsequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAdministratorConsequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

