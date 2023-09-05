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

// IdentityProviderType The type of identity provider.
type IdentityProviderType string

// List of IdentityProviderType
const (
	IDENTITYPROVIDERTYPE_UNKNOWN IdentityProviderType = "Unknown"
	IDENTITYPROVIDERTYPE_AD IdentityProviderType = "AD"
	IDENTITYPROVIDERTYPE_AZURE_AD IdentityProviderType = "AzureAD"
	IDENTITYPROVIDERTYPE_OKTA IdentityProviderType = "Okta"
	IDENTITYPROVIDERTYPE_GOOGLE IdentityProviderType = "Google"
	IDENTITYPROVIDERTYPE_ALL IdentityProviderType = "All"
)

// All allowed values of IdentityProviderType enum
var AllowedIdentityProviderTypeEnumValues = []IdentityProviderType{
	"Unknown",
	"AD",
	"AzureAD",
	"Okta",
	"Google",
	"All",
}

func (v *IdentityProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityProviderType(value)
	for _, existing := range AllowedIdentityProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityProviderType", value)
}

// NewIdentityProviderTypeFromValue returns a pointer to a valid IdentityProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityProviderTypeFromValue(v string) (*IdentityProviderType, error) {
	ev := IdentityProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityProviderType: valid values are %v", v, AllowedIdentityProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityProviderType) IsValid() bool {
	for _, existing := range AllowedIdentityProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityProviderType value
func (v IdentityProviderType) Ptr() *IdentityProviderType {
	return &v
}

type NullableIdentityProviderType struct {
	value *IdentityProviderType
	isSet bool
}

func (v NullableIdentityProviderType) Get() *IdentityProviderType {
	return v.value
}

func (v *NullableIdentityProviderType) Set(val *IdentityProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderType(val *IdentityProviderType) *NullableIdentityProviderType {
	return &NullableIdentityProviderType{value: val, isSet: true}
}

func (v NullableIdentityProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

