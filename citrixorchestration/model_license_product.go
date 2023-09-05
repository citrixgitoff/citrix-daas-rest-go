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

// LicenseProduct License Product
type LicenseProduct string

// List of LicenseProduct
const (
	LICENSEPRODUCT_UNKNOWN LicenseProduct = "Unknown"
	LICENSEPRODUCT_XEN_DESKTOP_XDS LicenseProduct = "XenDesktopXds"
	LICENSEPRODUCT_XEN_DESKTOP_XDT LicenseProduct = "XenDesktopXdt"
	LICENSEPRODUCT_XEN_DESKTOP_XDT_TP LicenseProduct = "XenDesktopXdtTP"
	LICENSEPRODUCT_XEN_APP LicenseProduct = "XenApp"
	LICENSEPRODUCT_CVAP LicenseProduct = "CVAP"
)

// All allowed values of LicenseProduct enum
var AllowedLicenseProductEnumValues = []LicenseProduct{
	"Unknown",
	"XenDesktopXds",
	"XenDesktopXdt",
	"XenDesktopXdtTP",
	"XenApp",
	"CVAP",
}

func (v *LicenseProduct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseProduct(value)
	for _, existing := range AllowedLicenseProductEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseProduct", value)
}

// NewLicenseProductFromValue returns a pointer to a valid LicenseProduct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseProductFromValue(v string) (*LicenseProduct, error) {
	ev := LicenseProduct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseProduct: valid values are %v", v, AllowedLicenseProductEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseProduct) IsValid() bool {
	for _, existing := range AllowedLicenseProductEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseProduct value
func (v LicenseProduct) Ptr() *LicenseProduct {
	return &v
}

type NullableLicenseProduct struct {
	value *LicenseProduct
	isSet bool
}

func (v NullableLicenseProduct) Get() *LicenseProduct {
	return v.value
}

func (v *NullableLicenseProduct) Set(val *LicenseProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseProduct(val *LicenseProduct) *NullableLicenseProduct {
	return &NullableLicenseProduct{value: val, isSet: true}
}

func (v NullableLicenseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

