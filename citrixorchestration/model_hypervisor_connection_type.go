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

// HypervisorConnectionType Types of hypervisor connections.
type HypervisorConnectionType string

// List of HypervisorConnectionType
const (
	HYPERVISORCONNECTIONTYPE_UNKNOWN HypervisorConnectionType = "Unknown"
	HYPERVISORCONNECTIONTYPE_XEN_SERVER HypervisorConnectionType = "XenServer"
	HYPERVISORCONNECTIONTYPE_SCVMM HypervisorConnectionType = "SCVMM"
	HYPERVISORCONNECTIONTYPE_V_CENTER HypervisorConnectionType = "VCenter"
	HYPERVISORCONNECTIONTYPE_CUSTOM HypervisorConnectionType = "Custom"
	HYPERVISORCONNECTIONTYPE_AWS HypervisorConnectionType = "AWS"
	HYPERVISORCONNECTIONTYPE_WAKE_ON_LAN HypervisorConnectionType = "WakeOnLAN"
	HYPERVISORCONNECTIONTYPE_AZURE_RM HypervisorConnectionType = "AzureRM"
	HYPERVISORCONNECTIONTYPE_GOOGLE_CLOUD_PLATFORM HypervisorConnectionType = "GoogleCloudPlatform"
	HYPERVISORCONNECTIONTYPE_CLOUD_PLATFORM HypervisorConnectionType = "CloudPlatform"
	HYPERVISORCONNECTIONTYPE_ORACLE_CLOUD_INFRASTRUCTURE HypervisorConnectionType = "OracleCloudInfrastructure"
	HYPERVISORCONNECTIONTYPE_AZURE_ARC HypervisorConnectionType = "AzureArc"
)

// All allowed values of HypervisorConnectionType enum
var AllowedHypervisorConnectionTypeEnumValues = []HypervisorConnectionType{
	"Unknown",
	"XenServer",
	"SCVMM",
	"VCenter",
	"Custom",
	"AWS",
	"WakeOnLAN",
	"AzureRM",
	"GoogleCloudPlatform",
	"CloudPlatform",
	"OracleCloudInfrastructure",
	"AzureArc",
}

func (v *HypervisorConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HypervisorConnectionType(value)
	for _, existing := range AllowedHypervisorConnectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HypervisorConnectionType", value)
}

// NewHypervisorConnectionTypeFromValue returns a pointer to a valid HypervisorConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHypervisorConnectionTypeFromValue(v string) (*HypervisorConnectionType, error) {
	ev := HypervisorConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HypervisorConnectionType: valid values are %v", v, AllowedHypervisorConnectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HypervisorConnectionType) IsValid() bool {
	for _, existing := range AllowedHypervisorConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HypervisorConnectionType value
func (v HypervisorConnectionType) Ptr() *HypervisorConnectionType {
	return &v
}

type NullableHypervisorConnectionType struct {
	value *HypervisorConnectionType
	isSet bool
}

func (v NullableHypervisorConnectionType) Get() *HypervisorConnectionType {
	return v.value
}

func (v *NullableHypervisorConnectionType) Set(val *HypervisorConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorConnectionType(val *HypervisorConnectionType) *NullableHypervisorConnectionType {
	return &NullableHypervisorConnectionType{value: val, isSet: true}
}

func (v NullableHypervisorConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

