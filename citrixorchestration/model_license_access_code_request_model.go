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

// checks if the LicenseAccessCodeRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseAccessCodeRequestModel{}

// LicenseAccessCodeRequestModel The request model for License Access Code
type LicenseAccessCodeRequestModel struct {
	// The License Access Code provided by Citrix.
	LicenseAccessCode string `json:"LicenseAccessCode"`
}

// NewLicenseAccessCodeRequestModel instantiates a new LicenseAccessCodeRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAccessCodeRequestModel(licenseAccessCode string) *LicenseAccessCodeRequestModel {
	this := LicenseAccessCodeRequestModel{}
	this.LicenseAccessCode = licenseAccessCode
	return &this
}

// NewLicenseAccessCodeRequestModelWithDefaults instantiates a new LicenseAccessCodeRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAccessCodeRequestModelWithDefaults() *LicenseAccessCodeRequestModel {
	this := LicenseAccessCodeRequestModel{}
	return &this
}

// GetLicenseAccessCode returns the LicenseAccessCode field value
func (o *LicenseAccessCodeRequestModel) GetLicenseAccessCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseAccessCode
}

// GetLicenseAccessCodeOk returns a tuple with the LicenseAccessCode field value
// and a boolean to check if the value has been set.
func (o *LicenseAccessCodeRequestModel) GetLicenseAccessCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseAccessCode, true
}

// SetLicenseAccessCode sets field value
func (o *LicenseAccessCodeRequestModel) SetLicenseAccessCode(v string) {
	o.LicenseAccessCode = v
}

func (o LicenseAccessCodeRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseAccessCodeRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["LicenseAccessCode"] = o.LicenseAccessCode
	return toSerialize, nil
}

type NullableLicenseAccessCodeRequestModel struct {
	value *LicenseAccessCodeRequestModel
	isSet bool
}

func (v NullableLicenseAccessCodeRequestModel) Get() *LicenseAccessCodeRequestModel {
	return v.value
}

func (v *NullableLicenseAccessCodeRequestModel) Set(val *LicenseAccessCodeRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAccessCodeRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAccessCodeRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAccessCodeRequestModel(val *LicenseAccessCodeRequestModel) *NullableLicenseAccessCodeRequestModel {
	return &NullableLicenseAccessCodeRequestModel{value: val, isSet: true}
}

func (v NullableLicenseAccessCodeRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAccessCodeRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


