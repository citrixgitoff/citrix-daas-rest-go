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

// checks if the IdentitySubnetResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySubnetResponseModel{}

// IdentitySubnetResponseModel The Subnet Identity object of a Site.
type IdentitySubnetResponseModel struct {
	// The name associated with the subnet object.
	Name NullableString `json:"Name,omitempty"`
	// The guid associated with the subnet object.
	Guid NullableString `json:"Guid,omitempty"`
}

// NewIdentitySubnetResponseModel instantiates a new IdentitySubnetResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySubnetResponseModel() *IdentitySubnetResponseModel {
	this := IdentitySubnetResponseModel{}
	return &this
}

// NewIdentitySubnetResponseModelWithDefaults instantiates a new IdentitySubnetResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySubnetResponseModelWithDefaults() *IdentitySubnetResponseModel {
	this := IdentitySubnetResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySubnetResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySubnetResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentitySubnetResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdentitySubnetResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentitySubnetResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentitySubnetResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySubnetResponseModel) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySubnetResponseModel) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *IdentitySubnetResponseModel) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *IdentitySubnetResponseModel) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *IdentitySubnetResponseModel) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *IdentitySubnetResponseModel) UnsetGuid() {
	o.Guid.Unset()
}

func (o IdentitySubnetResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySubnetResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["Guid"] = o.Guid.Get()
	}
	return toSerialize, nil
}

type NullableIdentitySubnetResponseModel struct {
	value *IdentitySubnetResponseModel
	isSet bool
}

func (v NullableIdentitySubnetResponseModel) Get() *IdentitySubnetResponseModel {
	return v.value
}

func (v *NullableIdentitySubnetResponseModel) Set(val *IdentitySubnetResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySubnetResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySubnetResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySubnetResponseModel(val *IdentitySubnetResponseModel) *NullableIdentitySubnetResponseModel {
	return &NullableIdentitySubnetResponseModel{value: val, isSet: true}
}

func (v NullableIdentitySubnetResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySubnetResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


