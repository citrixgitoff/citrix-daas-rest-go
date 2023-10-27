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

// PrintTrueTypeOption True type option.
type PrintTrueTypeOption string

// List of PrintTrueTypeOption
const (
	PRINTTRUETYPEOPTION_BITMAP PrintTrueTypeOption = "Bitmap"
	PRINTTRUETYPEOPTION_DOWNLOAD PrintTrueTypeOption = "Download"
	PRINTTRUETYPEOPTION_SUBSTITUTE PrintTrueTypeOption = "Substitute"
	PRINTTRUETYPEOPTION_OUTLINE PrintTrueTypeOption = "Outline"
)

// All allowed values of PrintTrueTypeOption enum
var AllowedPrintTrueTypeOptionEnumValues = []PrintTrueTypeOption{
	"Bitmap",
	"Download",
	"Substitute",
	"Outline",
}

func (v *PrintTrueTypeOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrintTrueTypeOption(value)
	for _, existing := range AllowedPrintTrueTypeOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrintTrueTypeOption", value)
}

// NewPrintTrueTypeOptionFromValue returns a pointer to a valid PrintTrueTypeOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrintTrueTypeOptionFromValue(v string) (*PrintTrueTypeOption, error) {
	ev := PrintTrueTypeOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrintTrueTypeOption: valid values are %v", v, AllowedPrintTrueTypeOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrintTrueTypeOption) IsValid() bool {
	for _, existing := range AllowedPrintTrueTypeOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrintTrueTypeOption value
func (v PrintTrueTypeOption) Ptr() *PrintTrueTypeOption {
	return &v
}

type NullablePrintTrueTypeOption struct {
	value *PrintTrueTypeOption
	isSet bool
}

func (v NullablePrintTrueTypeOption) Get() *PrintTrueTypeOption {
	return v.value
}

func (v *NullablePrintTrueTypeOption) Set(val *PrintTrueTypeOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintTrueTypeOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintTrueTypeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintTrueTypeOption(val *PrintTrueTypeOption) *NullablePrintTrueTypeOption {
	return &NullablePrintTrueTypeOption{value: val, isSet: true}
}

func (v NullablePrintTrueTypeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintTrueTypeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

