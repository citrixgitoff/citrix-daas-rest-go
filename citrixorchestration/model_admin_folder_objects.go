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

// AdminFolderObjects Admin folder object name, indicating object in the folder.
type AdminFolderObjects string

// List of AdminFolderObjects
const (
	ADMINFOLDEROBJECTS_UNKNOWN AdminFolderObjects = "Unknown"
	ADMINFOLDEROBJECTS_APPLICATIONS AdminFolderObjects = "Applications"
	ADMINFOLDEROBJECTS_MACHINE_CATALOGS AdminFolderObjects = "MachineCatalogs"
	ADMINFOLDEROBJECTS_DELIVERY_GROUPS AdminFolderObjects = "DeliveryGroups"
	ADMINFOLDEROBJECTS_APPLICATION_GROUPS AdminFolderObjects = "ApplicationGroups"
)

// All allowed values of AdminFolderObjects enum
var AllowedAdminFolderObjectsEnumValues = []AdminFolderObjects{
	"Unknown",
	"Applications",
	"MachineCatalogs",
	"DeliveryGroups",
	"ApplicationGroups",
}

func (v *AdminFolderObjects) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdminFolderObjects(value)
	for _, existing := range AllowedAdminFolderObjectsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdminFolderObjects", value)
}

// NewAdminFolderObjectsFromValue returns a pointer to a valid AdminFolderObjects
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdminFolderObjectsFromValue(v string) (*AdminFolderObjects, error) {
	ev := AdminFolderObjects(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdminFolderObjects: valid values are %v", v, AllowedAdminFolderObjectsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdminFolderObjects) IsValid() bool {
	for _, existing := range AllowedAdminFolderObjectsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdminFolderObjects value
func (v AdminFolderObjects) Ptr() *AdminFolderObjects {
	return &v
}

type NullableAdminFolderObjects struct {
	value *AdminFolderObjects
	isSet bool
}

func (v NullableAdminFolderObjects) Get() *AdminFolderObjects {
	return v.value
}

func (v *NullableAdminFolderObjects) Set(val *AdminFolderObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminFolderObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminFolderObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminFolderObjects(val *AdminFolderObjects) *NullableAdminFolderObjects {
	return &NullableAdminFolderObjects{value: val, isSet: true}
}

func (v NullableAdminFolderObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminFolderObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

