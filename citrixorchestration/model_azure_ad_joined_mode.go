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

// AzureAdJoinedMode Azure AD joined type for machines
type AzureAdJoinedMode string

// List of AzureAdJoinedMode
const (
	AZUREADJOINEDMODE_UNKNOWN AzureAdJoinedMode = "Unknown"
	AZUREADJOINEDMODE_NOT_AAD_JOINED AzureAdJoinedMode = "NotAadJoined"
	AZUREADJOINEDMODE_HYBRID_AAD_JOINED AzureAdJoinedMode = "HybridAadJoined"
	AZUREADJOINEDMODE_PURE_AAD_JOINED AzureAdJoinedMode = "PureAadJoined"
)

// All allowed values of AzureAdJoinedMode enum
var AllowedAzureAdJoinedModeEnumValues = []AzureAdJoinedMode{
	"Unknown",
	"NotAadJoined",
	"HybridAadJoined",
	"PureAadJoined",
}

func (v *AzureAdJoinedMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureAdJoinedMode(value)
	for _, existing := range AllowedAzureAdJoinedModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureAdJoinedMode", value)
}

// NewAzureAdJoinedModeFromValue returns a pointer to a valid AzureAdJoinedMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureAdJoinedModeFromValue(v string) (*AzureAdJoinedMode, error) {
	ev := AzureAdJoinedMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureAdJoinedMode: valid values are %v", v, AllowedAzureAdJoinedModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureAdJoinedMode) IsValid() bool {
	for _, existing := range AllowedAzureAdJoinedModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureAdJoinedMode value
func (v AzureAdJoinedMode) Ptr() *AzureAdJoinedMode {
	return &v
}

type NullableAzureAdJoinedMode struct {
	value *AzureAdJoinedMode
	isSet bool
}

func (v NullableAzureAdJoinedMode) Get() *AzureAdJoinedMode {
	return v.value
}

func (v *NullableAzureAdJoinedMode) Set(val *AzureAdJoinedMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureAdJoinedMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureAdJoinedMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureAdJoinedMode(val *AzureAdJoinedMode) *NullableAzureAdJoinedMode {
	return &NullableAzureAdJoinedMode{value: val, isSet: true}
}

func (v NullableAzureAdJoinedMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureAdJoinedMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

