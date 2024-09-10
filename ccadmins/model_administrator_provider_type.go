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

// AdministratorProviderType the model 'AdministratorProviderType'
type AdministratorProviderType string

// List of AdministratorProviderType
const (
	ADMINISTRATORPROVIDERTYPE_AD AdministratorProviderType = "Ad"
	ADMINISTRATORPROVIDERTYPE_AZURE_AD AdministratorProviderType = "AzureAd"
	ADMINISTRATORPROVIDERTYPE_CITRIX_STS AdministratorProviderType = "CitrixSts"
	ADMINISTRATORPROVIDERTYPE_GOOGLE AdministratorProviderType = "Google"
)

// All allowed values of AdministratorProviderType enum
var AllowedAdministratorProviderTypeEnumValues = []AdministratorProviderType{
	"Ad",
	"AzureAd",
	"CitrixSts",
	"Google",
}

func (v *AdministratorProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdministratorProviderType(value)
	for _, existing := range AllowedAdministratorProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdministratorProviderType", value)
}

// NewAdministratorProviderTypeFromValue returns a pointer to a valid AdministratorProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdministratorProviderTypeFromValue(v string) (*AdministratorProviderType, error) {
	ev := AdministratorProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdministratorProviderType: valid values are %v", v, AllowedAdministratorProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdministratorProviderType) IsValid() bool {
	for _, existing := range AllowedAdministratorProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdministratorProviderType value
func (v AdministratorProviderType) Ptr() *AdministratorProviderType {
	return &v
}

type NullableAdministratorProviderType struct {
	value *AdministratorProviderType
	isSet bool
}

func (v NullableAdministratorProviderType) Get() *AdministratorProviderType {
	return v.value
}

func (v *NullableAdministratorProviderType) Set(val *AdministratorProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorProviderType(val *AdministratorProviderType) *NullableAdministratorProviderType {
	return &NullableAdministratorProviderType{value: val, isSet: true}
}

func (v NullableAdministratorProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

