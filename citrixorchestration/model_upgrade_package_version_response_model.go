/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the UpgradePackageVersionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradePackageVersionResponseModel{}

// UpgradePackageVersionResponseModel Response object for a VDA upgrade package version.
type UpgradePackageVersionResponseModel struct {
	UpgradeType VdaUpgradeType `json:"UpgradeType"`
	// Latest released package version for the upgrade type.
	UpgradePackageVersion string `json:"UpgradePackageVersion"`
}

// NewUpgradePackageVersionResponseModel instantiates a new UpgradePackageVersionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradePackageVersionResponseModel(upgradeType VdaUpgradeType, upgradePackageVersion string) *UpgradePackageVersionResponseModel {
	this := UpgradePackageVersionResponseModel{}
	this.UpgradeType = upgradeType
	this.UpgradePackageVersion = upgradePackageVersion
	return &this
}

// NewUpgradePackageVersionResponseModelWithDefaults instantiates a new UpgradePackageVersionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradePackageVersionResponseModelWithDefaults() *UpgradePackageVersionResponseModel {
	this := UpgradePackageVersionResponseModel{}
	return &this
}

// GetUpgradeType returns the UpgradeType field value
func (o *UpgradePackageVersionResponseModel) GetUpgradeType() VdaUpgradeType {
	if o == nil {
		var ret VdaUpgradeType
		return ret
	}

	return o.UpgradeType
}

// GetUpgradeTypeOk returns a tuple with the UpgradeType field value
// and a boolean to check if the value has been set.
func (o *UpgradePackageVersionResponseModel) GetUpgradeTypeOk() (*VdaUpgradeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradeType, true
}

// SetUpgradeType sets field value
func (o *UpgradePackageVersionResponseModel) SetUpgradeType(v VdaUpgradeType) {
	o.UpgradeType = v
}

// GetUpgradePackageVersion returns the UpgradePackageVersion field value
func (o *UpgradePackageVersionResponseModel) GetUpgradePackageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpgradePackageVersion
}

// GetUpgradePackageVersionOk returns a tuple with the UpgradePackageVersion field value
// and a boolean to check if the value has been set.
func (o *UpgradePackageVersionResponseModel) GetUpgradePackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradePackageVersion, true
}

// SetUpgradePackageVersion sets field value
func (o *UpgradePackageVersionResponseModel) SetUpgradePackageVersion(v string) {
	o.UpgradePackageVersion = v
}

func (o UpgradePackageVersionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradePackageVersionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["UpgradeType"] = o.UpgradeType
	toSerialize["UpgradePackageVersion"] = o.UpgradePackageVersion
	return toSerialize, nil
}

type NullableUpgradePackageVersionResponseModel struct {
	value *UpgradePackageVersionResponseModel
	isSet bool
}

func (v NullableUpgradePackageVersionResponseModel) Get() *UpgradePackageVersionResponseModel {
	return v.value
}

func (v *NullableUpgradePackageVersionResponseModel) Set(val *UpgradePackageVersionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradePackageVersionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradePackageVersionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradePackageVersionResponseModel(val *UpgradePackageVersionResponseModel) *NullableUpgradePackageVersionResponseModel {
	return &NullableUpgradePackageVersionResponseModel{value: val, isSet: true}
}

func (v NullableUpgradePackageVersionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradePackageVersionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


