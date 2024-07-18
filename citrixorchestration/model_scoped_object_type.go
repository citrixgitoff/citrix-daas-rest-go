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

// ScopedObjectType Type of an object within a delegated admin scope.
type ScopedObjectType string

// List of ScopedObjectType
const (
	SCOPEDOBJECTTYPE_UNKNOWN ScopedObjectType = "Unknown"
	SCOPEDOBJECTTYPE_HYPERVISOR_CONNECTION ScopedObjectType = "HypervisorConnection"
	SCOPEDOBJECTTYPE_MACHINE_CATALOG ScopedObjectType = "MachineCatalog"
	SCOPEDOBJECTTYPE_DELIVERY_GROUP ScopedObjectType = "DeliveryGroup"
	SCOPEDOBJECTTYPE_APPLICATION_GROUP ScopedObjectType = "ApplicationGroup"
	SCOPEDOBJECTTYPE_TAG ScopedObjectType = "Tag"
	SCOPEDOBJECTTYPE_POLICY_SET ScopedObjectType = "PolicySet"
	SCOPEDOBJECTTYPE_SERVICE_ACCOUNT ScopedObjectType = "ServiceAccount"
)

// All allowed values of ScopedObjectType enum
var AllowedScopedObjectTypeEnumValues = []ScopedObjectType{
	"Unknown",
	"HypervisorConnection",
	"MachineCatalog",
	"DeliveryGroup",
	"ApplicationGroup",
	"Tag",
	"PolicySet",
	"ServiceAccount",
}

func (v *ScopedObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopedObjectType(value)
	for _, existing := range AllowedScopedObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopedObjectType", value)
}

// NewScopedObjectTypeFromValue returns a pointer to a valid ScopedObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopedObjectTypeFromValue(v string) (*ScopedObjectType, error) {
	ev := ScopedObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopedObjectType: valid values are %v", v, AllowedScopedObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopedObjectType) IsValid() bool {
	for _, existing := range AllowedScopedObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScopedObjectType value
func (v ScopedObjectType) Ptr() *ScopedObjectType {
	return &v
}

type NullableScopedObjectType struct {
	value *ScopedObjectType
	isSet bool
}

func (v NullableScopedObjectType) Get() *ScopedObjectType {
	return v.value
}

func (v *NullableScopedObjectType) Set(val *ScopedObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedObjectType(val *ScopedObjectType) *NullableScopedObjectType {
	return &NullableScopedObjectType{value: val, isSet: true}
}

func (v NullableScopedObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

