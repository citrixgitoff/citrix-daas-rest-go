/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the ProvisioningSchemeResponseModelImageRuntimeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSchemeResponseModelImageRuntimeInfo{}

// ProvisioningSchemeResponseModelImageRuntimeInfo Image runtime information.
type ProvisioningSchemeResponseModelImageRuntimeInfo struct {
	// The JSON schema version of image runtime info file.
	Version *string `json:"Version,omitempty"`
	OperatingSystemInfo *ImageRuntimeInfoResponseModelOperatingSystemInfo `json:"OperatingSystemInfo,omitempty"`
	// Installed VDA components information
	VdaComponents []VdaComponentResponseModel `json:"VdaComponents,omitempty"`
	SystemManagementInfo *ImageRuntimeInfoResponseModelSystemManagementInfo `json:"SystemManagementInfo,omitempty"`
}

// NewProvisioningSchemeResponseModelImageRuntimeInfo instantiates a new ProvisioningSchemeResponseModelImageRuntimeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSchemeResponseModelImageRuntimeInfo() *ProvisioningSchemeResponseModelImageRuntimeInfo {
	this := ProvisioningSchemeResponseModelImageRuntimeInfo{}
	return &this
}

// NewProvisioningSchemeResponseModelImageRuntimeInfoWithDefaults instantiates a new ProvisioningSchemeResponseModelImageRuntimeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSchemeResponseModelImageRuntimeInfoWithDefaults() *ProvisioningSchemeResponseModelImageRuntimeInfo {
	this := ProvisioningSchemeResponseModelImageRuntimeInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetVersion(v string) {
	o.Version = &v
}

// GetOperatingSystemInfo returns the OperatingSystemInfo field value if set, zero value otherwise.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetOperatingSystemInfo() ImageRuntimeInfoResponseModelOperatingSystemInfo {
	if o == nil || IsNil(o.OperatingSystemInfo) {
		var ret ImageRuntimeInfoResponseModelOperatingSystemInfo
		return ret
	}
	return *o.OperatingSystemInfo
}

// GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetOperatingSystemInfoOk() (*ImageRuntimeInfoResponseModelOperatingSystemInfo, bool) {
	if o == nil || IsNil(o.OperatingSystemInfo) {
		return nil, false
	}
	return o.OperatingSystemInfo, true
}

// HasOperatingSystemInfo returns a boolean if a field has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasOperatingSystemInfo() bool {
	if o != nil && !IsNil(o.OperatingSystemInfo) {
		return true
	}

	return false
}

// SetOperatingSystemInfo gets a reference to the given ImageRuntimeInfoResponseModelOperatingSystemInfo and assigns it to the OperatingSystemInfo field.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetOperatingSystemInfo(v ImageRuntimeInfoResponseModelOperatingSystemInfo) {
	o.OperatingSystemInfo = &v
}

// GetVdaComponents returns the VdaComponents field value if set, zero value otherwise.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVdaComponents() []VdaComponentResponseModel {
	if o == nil || IsNil(o.VdaComponents) {
		var ret []VdaComponentResponseModel
		return ret
	}
	return o.VdaComponents
}

// GetVdaComponentsOk returns a tuple with the VdaComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVdaComponentsOk() ([]VdaComponentResponseModel, bool) {
	if o == nil || IsNil(o.VdaComponents) {
		return nil, false
	}
	return o.VdaComponents, true
}

// HasVdaComponents returns a boolean if a field has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasVdaComponents() bool {
	if o != nil && !IsNil(o.VdaComponents) {
		return true
	}

	return false
}

// SetVdaComponents gets a reference to the given []VdaComponentResponseModel and assigns it to the VdaComponents field.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetVdaComponents(v []VdaComponentResponseModel) {
	o.VdaComponents = v
}

// GetSystemManagementInfo returns the SystemManagementInfo field value if set, zero value otherwise.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetSystemManagementInfo() ImageRuntimeInfoResponseModelSystemManagementInfo {
	if o == nil || IsNil(o.SystemManagementInfo) {
		var ret ImageRuntimeInfoResponseModelSystemManagementInfo
		return ret
	}
	return *o.SystemManagementInfo
}

// GetSystemManagementInfoOk returns a tuple with the SystemManagementInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetSystemManagementInfoOk() (*ImageRuntimeInfoResponseModelSystemManagementInfo, bool) {
	if o == nil || IsNil(o.SystemManagementInfo) {
		return nil, false
	}
	return o.SystemManagementInfo, true
}

// HasSystemManagementInfo returns a boolean if a field has been set.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasSystemManagementInfo() bool {
	if o != nil && !IsNil(o.SystemManagementInfo) {
		return true
	}

	return false
}

// SetSystemManagementInfo gets a reference to the given ImageRuntimeInfoResponseModelSystemManagementInfo and assigns it to the SystemManagementInfo field.
func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetSystemManagementInfo(v ImageRuntimeInfoResponseModelSystemManagementInfo) {
	o.SystemManagementInfo = &v
}

func (o ProvisioningSchemeResponseModelImageRuntimeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSchemeResponseModelImageRuntimeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.OperatingSystemInfo) {
		toSerialize["OperatingSystemInfo"] = o.OperatingSystemInfo
	}
	if !IsNil(o.VdaComponents) {
		toSerialize["VdaComponents"] = o.VdaComponents
	}
	if !IsNil(o.SystemManagementInfo) {
		toSerialize["SystemManagementInfo"] = o.SystemManagementInfo
	}
	return toSerialize, nil
}

type NullableProvisioningSchemeResponseModelImageRuntimeInfo struct {
	value *ProvisioningSchemeResponseModelImageRuntimeInfo
	isSet bool
}

func (v NullableProvisioningSchemeResponseModelImageRuntimeInfo) Get() *ProvisioningSchemeResponseModelImageRuntimeInfo {
	return v.value
}

func (v *NullableProvisioningSchemeResponseModelImageRuntimeInfo) Set(val *ProvisioningSchemeResponseModelImageRuntimeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSchemeResponseModelImageRuntimeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSchemeResponseModelImageRuntimeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSchemeResponseModelImageRuntimeInfo(val *ProvisioningSchemeResponseModelImageRuntimeInfo) *NullableProvisioningSchemeResponseModelImageRuntimeInfo {
	return &NullableProvisioningSchemeResponseModelImageRuntimeInfo{value: val, isSet: true}
}

func (v NullableProvisioningSchemeResponseModelImageRuntimeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSchemeResponseModelImageRuntimeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


