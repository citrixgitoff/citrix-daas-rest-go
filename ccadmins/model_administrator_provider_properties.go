/*
Administrators APIs

APIs for managing CC administrators.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the AdministratorProviderProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministratorProviderProperties{}

// AdministratorProviderProperties struct for AdministratorProviderProperties
type AdministratorProviderProperties struct {
	DisplayName NullableString `json:"displayName,omitempty"`
	Tid NullableString `json:"tid,omitempty"`
}

// NewAdministratorProviderProperties instantiates a new AdministratorProviderProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministratorProviderProperties() *AdministratorProviderProperties {
	this := AdministratorProviderProperties{}
	return &this
}

// NewAdministratorProviderPropertiesWithDefaults instantiates a new AdministratorProviderProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministratorProviderPropertiesWithDefaults() *AdministratorProviderProperties {
	this := AdministratorProviderProperties{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdministratorProviderProperties) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdministratorProviderProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AdministratorProviderProperties) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AdministratorProviderProperties) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AdministratorProviderProperties) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AdministratorProviderProperties) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTid returns the Tid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdministratorProviderProperties) GetTid() string {
	if o == nil || IsNil(o.Tid.Get()) {
		var ret string
		return ret
	}
	return *o.Tid.Get()
}

// GetTidOk returns a tuple with the Tid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdministratorProviderProperties) GetTidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tid.Get(), o.Tid.IsSet()
}

// HasTid returns a boolean if a field has been set.
func (o *AdministratorProviderProperties) HasTid() bool {
	if o != nil && o.Tid.IsSet() {
		return true
	}

	return false
}

// SetTid gets a reference to the given NullableString and assigns it to the Tid field.
func (o *AdministratorProviderProperties) SetTid(v string) {
	o.Tid.Set(&v)
}
// SetTidNil sets the value for Tid to be an explicit nil
func (o *AdministratorProviderProperties) SetTidNil() {
	o.Tid.Set(nil)
}

// UnsetTid ensures that no value is present for Tid, not even an explicit nil
func (o *AdministratorProviderProperties) UnsetTid() {
	o.Tid.Unset()
}

func (o AdministratorProviderProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministratorProviderProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Tid.IsSet() {
		toSerialize["tid"] = o.Tid.Get()
	}
	return toSerialize, nil
}

type NullableAdministratorProviderProperties struct {
	value *AdministratorProviderProperties
	isSet bool
}

func (v NullableAdministratorProviderProperties) Get() *AdministratorProviderProperties {
	return v.value
}

func (v *NullableAdministratorProviderProperties) Set(val *AdministratorProviderProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorProviderProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorProviderProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorProviderProperties(val *AdministratorProviderProperties) *NullableAdministratorProviderProperties {
	return &NullableAdministratorProviderProperties{value: val, isSet: true}
}

func (v NullableAdministratorProviderProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorProviderProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


