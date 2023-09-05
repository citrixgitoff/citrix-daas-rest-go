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

// DeliveryGroupSearchFilterGroupType The search filter group type
type DeliveryGroupSearchFilterGroupType string

// List of DeliveryGroupSearchFilterGroupType
const (
	DELIVERYGROUPSEARCHFILTERGROUPTYPE_OR DeliveryGroupSearchFilterGroupType = "Or"
	DELIVERYGROUPSEARCHFILTERGROUPTYPE_AND DeliveryGroupSearchFilterGroupType = "And"
)

// All allowed values of DeliveryGroupSearchFilterGroupType enum
var AllowedDeliveryGroupSearchFilterGroupTypeEnumValues = []DeliveryGroupSearchFilterGroupType{
	"Or",
	"And",
}

func (v *DeliveryGroupSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryGroupSearchFilterGroupType(value)
	for _, existing := range AllowedDeliveryGroupSearchFilterGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryGroupSearchFilterGroupType", value)
}

// NewDeliveryGroupSearchFilterGroupTypeFromValue returns a pointer to a valid DeliveryGroupSearchFilterGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryGroupSearchFilterGroupTypeFromValue(v string) (*DeliveryGroupSearchFilterGroupType, error) {
	ev := DeliveryGroupSearchFilterGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryGroupSearchFilterGroupType: valid values are %v", v, AllowedDeliveryGroupSearchFilterGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryGroupSearchFilterGroupType) IsValid() bool {
	for _, existing := range AllowedDeliveryGroupSearchFilterGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryGroupSearchFilterGroupType value
func (v DeliveryGroupSearchFilterGroupType) Ptr() *DeliveryGroupSearchFilterGroupType {
	return &v
}

type NullableDeliveryGroupSearchFilterGroupType struct {
	value *DeliveryGroupSearchFilterGroupType
	isSet bool
}

func (v NullableDeliveryGroupSearchFilterGroupType) Get() *DeliveryGroupSearchFilterGroupType {
	return v.value
}

func (v *NullableDeliveryGroupSearchFilterGroupType) Set(val *DeliveryGroupSearchFilterGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGroupSearchFilterGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGroupSearchFilterGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGroupSearchFilterGroupType(val *DeliveryGroupSearchFilterGroupType) *NullableDeliveryGroupSearchFilterGroupType {
	return &NullableDeliveryGroupSearchFilterGroupType{value: val, isSet: true}
}

func (v NullableDeliveryGroupSearchFilterGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGroupSearchFilterGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

