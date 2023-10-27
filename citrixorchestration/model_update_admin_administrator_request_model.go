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

// checks if the UpdateAdminAdministratorRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAdminAdministratorRequestModel{}

// UpdateAdminAdministratorRequestModel Model for editing administrator.
type UpdateAdminAdministratorRequestModel struct {
	// Rights associated with the administrator.
	Rights []AdminRightRequestModel `json:"Rights,omitempty"`
	// Indicates whether the administrator is enabled.  Disabled administrators cannot administer the site unless they are a member of a different user group which is granted access by a different administrator record.
	Enabled NullableBool `json:"Enabled,omitempty"`
}

// NewUpdateAdminAdministratorRequestModel instantiates a new UpdateAdminAdministratorRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAdminAdministratorRequestModel() *UpdateAdminAdministratorRequestModel {
	this := UpdateAdminAdministratorRequestModel{}
	return &this
}

// NewUpdateAdminAdministratorRequestModelWithDefaults instantiates a new UpdateAdminAdministratorRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAdminAdministratorRequestModelWithDefaults() *UpdateAdminAdministratorRequestModel {
	this := UpdateAdminAdministratorRequestModel{}
	return &this
}

// GetRights returns the Rights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAdminAdministratorRequestModel) GetRights() []AdminRightRequestModel {
	if o == nil {
		var ret []AdminRightRequestModel
		return ret
	}
	return o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAdminAdministratorRequestModel) GetRightsOk() ([]AdminRightRequestModel, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *UpdateAdminAdministratorRequestModel) HasRights() bool {
	if o != nil && IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given []AdminRightRequestModel and assigns it to the Rights field.
func (o *UpdateAdminAdministratorRequestModel) SetRights(v []AdminRightRequestModel) {
	o.Rights = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAdminAdministratorRequestModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAdminAdministratorRequestModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateAdminAdministratorRequestModel) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateAdminAdministratorRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateAdminAdministratorRequestModel) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateAdminAdministratorRequestModel) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o UpdateAdminAdministratorRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAdminAdministratorRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rights != nil {
		toSerialize["Rights"] = o.Rights
	}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableUpdateAdminAdministratorRequestModel struct {
	value *UpdateAdminAdministratorRequestModel
	isSet bool
}

func (v NullableUpdateAdminAdministratorRequestModel) Get() *UpdateAdminAdministratorRequestModel {
	return v.value
}

func (v *NullableUpdateAdminAdministratorRequestModel) Set(val *UpdateAdminAdministratorRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAdminAdministratorRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAdminAdministratorRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAdminAdministratorRequestModel(val *UpdateAdminAdministratorRequestModel) *NullableUpdateAdminAdministratorRequestModel {
	return &NullableUpdateAdminAdministratorRequestModel{value: val, isSet: true}
}

func (v NullableUpdateAdminAdministratorRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAdminAdministratorRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


