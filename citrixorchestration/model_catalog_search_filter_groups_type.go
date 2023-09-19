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

// CatalogSearchFilterGroupsType The search filter groups type
type CatalogSearchFilterGroupsType string

// List of CatalogSearchFilterGroupsType
const (
	CATALOGSEARCHFILTERGROUPSTYPE_OR CatalogSearchFilterGroupsType = "Or"
	CATALOGSEARCHFILTERGROUPSTYPE_AND CatalogSearchFilterGroupsType = "And"
)

// All allowed values of CatalogSearchFilterGroupsType enum
var AllowedCatalogSearchFilterGroupsTypeEnumValues = []CatalogSearchFilterGroupsType{
	"Or",
	"And",
}

func (v *CatalogSearchFilterGroupsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogSearchFilterGroupsType(value)
	for _, existing := range AllowedCatalogSearchFilterGroupsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogSearchFilterGroupsType", value)
}

// NewCatalogSearchFilterGroupsTypeFromValue returns a pointer to a valid CatalogSearchFilterGroupsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogSearchFilterGroupsTypeFromValue(v string) (*CatalogSearchFilterGroupsType, error) {
	ev := CatalogSearchFilterGroupsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogSearchFilterGroupsType: valid values are %v", v, AllowedCatalogSearchFilterGroupsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogSearchFilterGroupsType) IsValid() bool {
	for _, existing := range AllowedCatalogSearchFilterGroupsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogSearchFilterGroupsType value
func (v CatalogSearchFilterGroupsType) Ptr() *CatalogSearchFilterGroupsType {
	return &v
}

type NullableCatalogSearchFilterGroupsType struct {
	value *CatalogSearchFilterGroupsType
	isSet bool
}

func (v NullableCatalogSearchFilterGroupsType) Get() *CatalogSearchFilterGroupsType {
	return v.value
}

func (v *NullableCatalogSearchFilterGroupsType) Set(val *CatalogSearchFilterGroupsType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogSearchFilterGroupsType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogSearchFilterGroupsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogSearchFilterGroupsType(val *CatalogSearchFilterGroupsType) *NullableCatalogSearchFilterGroupsType {
	return &NullableCatalogSearchFilterGroupsType{value: val, isSet: true}
}

func (v NullableCatalogSearchFilterGroupsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogSearchFilterGroupsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

