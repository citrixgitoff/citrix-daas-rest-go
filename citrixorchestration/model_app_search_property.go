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

// AppSearchProperty Application properties can be used for advanced application searches.             
type AppSearchProperty string

// List of AppSearchProperty
const (
	APPSEARCHPROPERTY_APPLICATION_CATEGORY AppSearchProperty = "ApplicationCategory"
	APPSEARCHPROPERTY_APPLICATION_GROUP AppSearchProperty = "ApplicationGroup"
	APPSEARCHPROPERTY_NAME AppSearchProperty = "Name"
	APPSEARCHPROPERTY_PATH AppSearchProperty = "Path"
	APPSEARCHPROPERTY_BROWSER_NAME AppSearchProperty = "BrowserName"
	APPSEARCHPROPERTY_COMMAND_LINE_ARGUMENTS AppSearchProperty = "CommandLineArguments"
	APPSEARCHPROPERTY_DELIVERY_GROUP AppSearchProperty = "DeliveryGroup"
	APPSEARCHPROPERTY_DESCRIPTION AppSearchProperty = "Description"
	APPSEARCHPROPERTY_PUBLISHED_NAME AppSearchProperty = "PublishedName"
	APPSEARCHPROPERTY_ENABLED AppSearchProperty = "Enabled"
	APPSEARCHPROPERTY_MAX_TOTAL_INSTANCES AppSearchProperty = "MaxTotalInstances"
	APPSEARCHPROPERTY_MAX_PER_USER_INSTANCES AppSearchProperty = "MaxPerUserInstances"
	APPSEARCHPROPERTY_SHORTCUT_ADDED_TO_DESKTOP AppSearchProperty = "ShortcutAddedToDesktop"
	APPSEARCHPROPERTY_SHORTCUT_ADDED_TO_START_MENU AppSearchProperty = "ShortcutAddedToStartMenu"
	APPSEARCHPROPERTY_START_MENU_FOLDER AppSearchProperty = "StartMenuFolder"
	APPSEARCHPROPERTY_TAGS AppSearchProperty = "Tags"
	APPSEARCHPROPERTY_APPLICATION_TYPE AppSearchProperty = "ApplicationType"
	APPSEARCHPROPERTY_VISIBLE AppSearchProperty = "Visible"
	APPSEARCHPROPERTY_WAIT_FOR_PRINTER_CREATION AppSearchProperty = "WaitForPrinterCreation"
	APPSEARCHPROPERTY_WORKING_DIRECTORY AppSearchProperty = "WorkingDirectory"
)

// All allowed values of AppSearchProperty enum
var AllowedAppSearchPropertyEnumValues = []AppSearchProperty{
	"ApplicationCategory",
	"ApplicationGroup",
	"Name",
	"Path",
	"BrowserName",
	"CommandLineArguments",
	"DeliveryGroup",
	"Description",
	"PublishedName",
	"Enabled",
	"MaxTotalInstances",
	"MaxPerUserInstances",
	"ShortcutAddedToDesktop",
	"ShortcutAddedToStartMenu",
	"StartMenuFolder",
	"Tags",
	"ApplicationType",
	"Visible",
	"WaitForPrinterCreation",
	"WorkingDirectory",
}

func (v *AppSearchProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppSearchProperty(value)
	for _, existing := range AllowedAppSearchPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppSearchProperty", value)
}

// NewAppSearchPropertyFromValue returns a pointer to a valid AppSearchProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppSearchPropertyFromValue(v string) (*AppSearchProperty, error) {
	ev := AppSearchProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppSearchProperty: valid values are %v", v, AllowedAppSearchPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppSearchProperty) IsValid() bool {
	for _, existing := range AllowedAppSearchPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppSearchProperty value
func (v AppSearchProperty) Ptr() *AppSearchProperty {
	return &v
}

type NullableAppSearchProperty struct {
	value *AppSearchProperty
	isSet bool
}

func (v NullableAppSearchProperty) Get() *AppSearchProperty {
	return v.value
}

func (v *NullableAppSearchProperty) Set(val *AppSearchProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSearchProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSearchProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSearchProperty(val *AppSearchProperty) *NullableAppSearchProperty {
	return &NullableAppSearchProperty{value: val, isSet: true}
}

func (v NullableAppSearchProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSearchProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

