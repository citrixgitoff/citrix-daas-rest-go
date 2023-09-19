/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the LicensePermissionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensePermissionResponseModel{}

// LicensePermissionResponseModel The permission to the license server
type LicensePermissionResponseModel struct {
	// The administrator  for the license server
	LicensingAdministrator NullableString `json:"LicensingAdministrator,omitempty"`
	LicensingPermissionLevel *LicensingPermissionLevel `json:"LicensingPermissionLevel,omitempty"`
}

// NewLicensePermissionResponseModel instantiates a new LicensePermissionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensePermissionResponseModel() *LicensePermissionResponseModel {
	this := LicensePermissionResponseModel{}
	return &this
}

// NewLicensePermissionResponseModelWithDefaults instantiates a new LicensePermissionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensePermissionResponseModelWithDefaults() *LicensePermissionResponseModel {
	this := LicensePermissionResponseModel{}
	return &this
}

// GetLicensingAdministrator returns the LicensingAdministrator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePermissionResponseModel) GetLicensingAdministrator() string {
	if o == nil || IsNil(o.LicensingAdministrator.Get()) {
		var ret string
		return ret
	}
	return *o.LicensingAdministrator.Get()
}

// GetLicensingAdministratorOk returns a tuple with the LicensingAdministrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePermissionResponseModel) GetLicensingAdministratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicensingAdministrator.Get(), o.LicensingAdministrator.IsSet()
}

// HasLicensingAdministrator returns a boolean if a field has been set.
func (o *LicensePermissionResponseModel) HasLicensingAdministrator() bool {
	if o != nil && o.LicensingAdministrator.IsSet() {
		return true
	}

	return false
}

// SetLicensingAdministrator gets a reference to the given NullableString and assigns it to the LicensingAdministrator field.
func (o *LicensePermissionResponseModel) SetLicensingAdministrator(v string) {
	o.LicensingAdministrator.Set(&v)
}
// SetLicensingAdministratorNil sets the value for LicensingAdministrator to be an explicit nil
func (o *LicensePermissionResponseModel) SetLicensingAdministratorNil() {
	o.LicensingAdministrator.Set(nil)
}

// UnsetLicensingAdministrator ensures that no value is present for LicensingAdministrator, not even an explicit nil
func (o *LicensePermissionResponseModel) UnsetLicensingAdministrator() {
	o.LicensingAdministrator.Unset()
}

// GetLicensingPermissionLevel returns the LicensingPermissionLevel field value if set, zero value otherwise.
func (o *LicensePermissionResponseModel) GetLicensingPermissionLevel() LicensingPermissionLevel {
	if o == nil || IsNil(o.LicensingPermissionLevel) {
		var ret LicensingPermissionLevel
		return ret
	}
	return *o.LicensingPermissionLevel
}

// GetLicensingPermissionLevelOk returns a tuple with the LicensingPermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensePermissionResponseModel) GetLicensingPermissionLevelOk() (*LicensingPermissionLevel, bool) {
	if o == nil || IsNil(o.LicensingPermissionLevel) {
		return nil, false
	}
	return o.LicensingPermissionLevel, true
}

// HasLicensingPermissionLevel returns a boolean if a field has been set.
func (o *LicensePermissionResponseModel) HasLicensingPermissionLevel() bool {
	if o != nil && !IsNil(o.LicensingPermissionLevel) {
		return true
	}

	return false
}

// SetLicensingPermissionLevel gets a reference to the given LicensingPermissionLevel and assigns it to the LicensingPermissionLevel field.
func (o *LicensePermissionResponseModel) SetLicensingPermissionLevel(v LicensingPermissionLevel) {
	o.LicensingPermissionLevel = &v
}

func (o LicensePermissionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensePermissionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LicensingAdministrator.IsSet() {
		toSerialize["LicensingAdministrator"] = o.LicensingAdministrator.Get()
	}
	if !IsNil(o.LicensingPermissionLevel) {
		toSerialize["LicensingPermissionLevel"] = o.LicensingPermissionLevel
	}
	return toSerialize, nil
}

type NullableLicensePermissionResponseModel struct {
	value *LicensePermissionResponseModel
	isSet bool
}

func (v NullableLicensePermissionResponseModel) Get() *LicensePermissionResponseModel {
	return v.value
}

func (v *NullableLicensePermissionResponseModel) Set(val *LicensePermissionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensePermissionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensePermissionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensePermissionResponseModel(val *LicensePermissionResponseModel) *NullableLicensePermissionResponseModel {
	return &NullableLicensePermissionResponseModel{value: val, isSet: true}
}

func (v NullableLicensePermissionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensePermissionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


