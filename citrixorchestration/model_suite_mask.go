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

// SuiteMask Windows product suit list.
type SuiteMask string

// List of SuiteMask
const (
	SUITEMASK_SMALL_BUSINESS SuiteMask = "SmallBusiness"
	SUITEMASK_ENTERPRISE SuiteMask = "Enterprise"
	SUITEMASK_BACK_OFFICE SuiteMask = "BackOffice"
	SUITEMASK_TERMINAL SuiteMask = "Terminal"
	SUITEMASK_SMALL_BUSINESS_RESTRICTED SuiteMask = "SmallBusinessRestricted"
	SUITEMASK_EMBEDDED_NT SuiteMask = "EmbeddedNT"
	SUITEMASK_DATACENTER SuiteMask = "Datacenter"
	SUITEMASK_SINGLE_USER_TS SuiteMask = "SingleUserTS"
	SUITEMASK_PERSONAL SuiteMask = "Personal"
	SUITEMASK_BLADE SuiteMask = "Blade"
	SUITEMASK_STORAGE_SERVER SuiteMask = "StorageServer"
	SUITEMASK_COMPUTE_SERVER SuiteMask = "ComputeServer"
	SUITEMASK_WINDOWS_HOME_SERVER SuiteMask = "WindowsHomeServer"
)

// All allowed values of SuiteMask enum
var AllowedSuiteMaskEnumValues = []SuiteMask{
	"SmallBusiness",
	"Enterprise",
	"BackOffice",
	"Terminal",
	"SmallBusinessRestricted",
	"EmbeddedNT",
	"Datacenter",
	"SingleUserTS",
	"Personal",
	"Blade",
	"StorageServer",
	"ComputeServer",
	"WindowsHomeServer",
}

func (v *SuiteMask) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SuiteMask(value)
	for _, existing := range AllowedSuiteMaskEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SuiteMask", value)
}

// NewSuiteMaskFromValue returns a pointer to a valid SuiteMask
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSuiteMaskFromValue(v string) (*SuiteMask, error) {
	ev := SuiteMask(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SuiteMask: valid values are %v", v, AllowedSuiteMaskEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SuiteMask) IsValid() bool {
	for _, existing := range AllowedSuiteMaskEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SuiteMask value
func (v SuiteMask) Ptr() *SuiteMask {
	return &v
}

type NullableSuiteMask struct {
	value *SuiteMask
	isSet bool
}

func (v NullableSuiteMask) Get() *SuiteMask {
	return v.value
}

func (v *NullableSuiteMask) Set(val *SuiteMask) {
	v.value = val
	v.isSet = true
}

func (v NullableSuiteMask) IsSet() bool {
	return v.isSet
}

func (v *NullableSuiteMask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuiteMask(val *SuiteMask) *NullableSuiteMask {
	return &NullableSuiteMask{value: val, isSet: true}
}

func (v NullableSuiteMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuiteMask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

