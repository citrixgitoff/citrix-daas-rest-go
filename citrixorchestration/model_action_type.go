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

// ActionType The action type.
type ActionType string

// List of ActionType
const (
	ACTIONTYPE_UNKNOWN ActionType = "Unknown"
	ACTIONTYPE_CREATE_ITEMS ActionType = "CreateItems"
	ACTIONTYPE_DELETE_ITEMS ActionType = "DeleteItems"
	ACTIONTYPE_CREATE_CATALOG ActionType = "CreateCatalog"
	ACTIONTYPE_UPDATE_IMAGE ActionType = "UpdateImage"
	ACTIONTYPE_ADD_MACHINES ActionType = "AddMachines"
	ACTIONTYPE_REMOVE_MACHINES ActionType = "RemoveMachines"
	ACTIONTYPE_UPDATE_SKY_WAY_IMAGE ActionType = "UpdateSkyWayImage"
	ACTIONTYPE_DELETE_MACHINES ActionType = "DeleteMachines"
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	"Unknown",
	"CreateItems",
	"DeleteItems",
	"CreateCatalog",
	"UpdateImage",
	"AddMachines",
	"RemoveMachines",
	"UpdateSkyWayImage",
	"DeleteMachines",
}

func (v *ActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionType(value)
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionType", value)
}

// NewActionTypeFromValue returns a pointer to a valid ActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionTypeFromValue(v string) (*ActionType, error) {
	ev := ActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionType: valid values are %v", v, AllowedActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionType value
func (v ActionType) Ptr() *ActionType {
	return &v
}

type NullableActionType struct {
	value *ActionType
	isSet bool
}

func (v NullableActionType) Get() *ActionType {
	return v.value
}

func (v *NullableActionType) Set(val *ActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionType(val *ActionType) *NullableActionType {
	return &NullableActionType{value: val, isSet: true}
}

func (v NullableActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

