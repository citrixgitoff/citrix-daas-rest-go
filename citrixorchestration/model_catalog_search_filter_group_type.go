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

// CatalogSearchFilterGroupType The search filter group type
type CatalogSearchFilterGroupType string

// List of CatalogSearchFilterGroupType
const (
	CATALOGSEARCHFILTERGROUPTYPE_OR CatalogSearchFilterGroupType = "Or"
	CATALOGSEARCHFILTERGROUPTYPE_AND CatalogSearchFilterGroupType = "And"
)

// All allowed values of CatalogSearchFilterGroupType enum
var AllowedCatalogSearchFilterGroupTypeEnumValues = []CatalogSearchFilterGroupType{
	"Or",
	"And",
}

func (v *CatalogSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogSearchFilterGroupType(value)
	for _, existing := range AllowedCatalogSearchFilterGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogSearchFilterGroupType", value)
}

// NewCatalogSearchFilterGroupTypeFromValue returns a pointer to a valid CatalogSearchFilterGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogSearchFilterGroupTypeFromValue(v string) (*CatalogSearchFilterGroupType, error) {
	ev := CatalogSearchFilterGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogSearchFilterGroupType: valid values are %v", v, AllowedCatalogSearchFilterGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogSearchFilterGroupType) IsValid() bool {
	for _, existing := range AllowedCatalogSearchFilterGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogSearchFilterGroupType value
func (v CatalogSearchFilterGroupType) Ptr() *CatalogSearchFilterGroupType {
	return &v
}

type NullableCatalogSearchFilterGroupType struct {
	value *CatalogSearchFilterGroupType
	isSet bool
}

func (v NullableCatalogSearchFilterGroupType) Get() *CatalogSearchFilterGroupType {
	return v.value
}

func (v *NullableCatalogSearchFilterGroupType) Set(val *CatalogSearchFilterGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogSearchFilterGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogSearchFilterGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogSearchFilterGroupType(val *CatalogSearchFilterGroupType) *NullableCatalogSearchFilterGroupType {
	return &NullableCatalogSearchFilterGroupType{value: val, isSet: true}
}

func (v NullableCatalogSearchFilterGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

