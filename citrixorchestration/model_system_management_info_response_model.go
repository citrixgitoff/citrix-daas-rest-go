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

// checks if the SystemManagementInfoResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemManagementInfoResponseModel{}

// SystemManagementInfoResponseModel System management information.
type SystemManagementInfoResponseModel struct {
	AzureADJoinStatus *SystemManagementInfoResponseModelAzureADJoinStatus `json:"AzureADJoinStatus,omitempty"`
	DeviceRegistrationStatus *SystemManagementInfoResponseModelDeviceRegistrationStatus `json:"DeviceRegistrationStatus,omitempty"`
	// System paging file settings.
	PagingFileSettings []PagingFileSettingResponseModel `json:"PagingFileSettings,omitempty"`
}

// NewSystemManagementInfoResponseModel instantiates a new SystemManagementInfoResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemManagementInfoResponseModel() *SystemManagementInfoResponseModel {
	this := SystemManagementInfoResponseModel{}
	return &this
}

// NewSystemManagementInfoResponseModelWithDefaults instantiates a new SystemManagementInfoResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemManagementInfoResponseModelWithDefaults() *SystemManagementInfoResponseModel {
	this := SystemManagementInfoResponseModel{}
	return &this
}

// GetAzureADJoinStatus returns the AzureADJoinStatus field value if set, zero value otherwise.
func (o *SystemManagementInfoResponseModel) GetAzureADJoinStatus() SystemManagementInfoResponseModelAzureADJoinStatus {
	if o == nil || IsNil(o.AzureADJoinStatus) {
		var ret SystemManagementInfoResponseModelAzureADJoinStatus
		return ret
	}
	return *o.AzureADJoinStatus
}

// GetAzureADJoinStatusOk returns a tuple with the AzureADJoinStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemManagementInfoResponseModel) GetAzureADJoinStatusOk() (*SystemManagementInfoResponseModelAzureADJoinStatus, bool) {
	if o == nil || IsNil(o.AzureADJoinStatus) {
		return nil, false
	}
	return o.AzureADJoinStatus, true
}

// HasAzureADJoinStatus returns a boolean if a field has been set.
func (o *SystemManagementInfoResponseModel) HasAzureADJoinStatus() bool {
	if o != nil && !IsNil(o.AzureADJoinStatus) {
		return true
	}

	return false
}

// SetAzureADJoinStatus gets a reference to the given SystemManagementInfoResponseModelAzureADJoinStatus and assigns it to the AzureADJoinStatus field.
func (o *SystemManagementInfoResponseModel) SetAzureADJoinStatus(v SystemManagementInfoResponseModelAzureADJoinStatus) {
	o.AzureADJoinStatus = &v
}

// GetDeviceRegistrationStatus returns the DeviceRegistrationStatus field value if set, zero value otherwise.
func (o *SystemManagementInfoResponseModel) GetDeviceRegistrationStatus() SystemManagementInfoResponseModelDeviceRegistrationStatus {
	if o == nil || IsNil(o.DeviceRegistrationStatus) {
		var ret SystemManagementInfoResponseModelDeviceRegistrationStatus
		return ret
	}
	return *o.DeviceRegistrationStatus
}

// GetDeviceRegistrationStatusOk returns a tuple with the DeviceRegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemManagementInfoResponseModel) GetDeviceRegistrationStatusOk() (*SystemManagementInfoResponseModelDeviceRegistrationStatus, bool) {
	if o == nil || IsNil(o.DeviceRegistrationStatus) {
		return nil, false
	}
	return o.DeviceRegistrationStatus, true
}

// HasDeviceRegistrationStatus returns a boolean if a field has been set.
func (o *SystemManagementInfoResponseModel) HasDeviceRegistrationStatus() bool {
	if o != nil && !IsNil(o.DeviceRegistrationStatus) {
		return true
	}

	return false
}

// SetDeviceRegistrationStatus gets a reference to the given SystemManagementInfoResponseModelDeviceRegistrationStatus and assigns it to the DeviceRegistrationStatus field.
func (o *SystemManagementInfoResponseModel) SetDeviceRegistrationStatus(v SystemManagementInfoResponseModelDeviceRegistrationStatus) {
	o.DeviceRegistrationStatus = &v
}

// GetPagingFileSettings returns the PagingFileSettings field value if set, zero value otherwise.
func (o *SystemManagementInfoResponseModel) GetPagingFileSettings() []PagingFileSettingResponseModel {
	if o == nil || IsNil(o.PagingFileSettings) {
		var ret []PagingFileSettingResponseModel
		return ret
	}
	return o.PagingFileSettings
}

// GetPagingFileSettingsOk returns a tuple with the PagingFileSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemManagementInfoResponseModel) GetPagingFileSettingsOk() ([]PagingFileSettingResponseModel, bool) {
	if o == nil || IsNil(o.PagingFileSettings) {
		return nil, false
	}
	return o.PagingFileSettings, true
}

// HasPagingFileSettings returns a boolean if a field has been set.
func (o *SystemManagementInfoResponseModel) HasPagingFileSettings() bool {
	if o != nil && !IsNil(o.PagingFileSettings) {
		return true
	}

	return false
}

// SetPagingFileSettings gets a reference to the given []PagingFileSettingResponseModel and assigns it to the PagingFileSettings field.
func (o *SystemManagementInfoResponseModel) SetPagingFileSettings(v []PagingFileSettingResponseModel) {
	o.PagingFileSettings = v
}

func (o SystemManagementInfoResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemManagementInfoResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AzureADJoinStatus) {
		toSerialize["AzureADJoinStatus"] = o.AzureADJoinStatus
	}
	if !IsNil(o.DeviceRegistrationStatus) {
		toSerialize["DeviceRegistrationStatus"] = o.DeviceRegistrationStatus
	}
	if !IsNil(o.PagingFileSettings) {
		toSerialize["PagingFileSettings"] = o.PagingFileSettings
	}
	return toSerialize, nil
}

type NullableSystemManagementInfoResponseModel struct {
	value *SystemManagementInfoResponseModel
	isSet bool
}

func (v NullableSystemManagementInfoResponseModel) Get() *SystemManagementInfoResponseModel {
	return v.value
}

func (v *NullableSystemManagementInfoResponseModel) Set(val *SystemManagementInfoResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemManagementInfoResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemManagementInfoResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemManagementInfoResponseModel(val *SystemManagementInfoResponseModel) *NullableSystemManagementInfoResponseModel {
	return &NullableSystemManagementInfoResponseModel{value: val, isSet: true}
}

func (v NullableSystemManagementInfoResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemManagementInfoResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


