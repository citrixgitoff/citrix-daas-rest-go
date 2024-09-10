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

// checks if the AdministratorNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministratorNotification{}

// AdministratorNotification struct for AdministratorNotification
type AdministratorNotification struct {
	Type *AdministratorNotificationType `json:"type,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewAdministratorNotification instantiates a new AdministratorNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministratorNotification() *AdministratorNotification {
	this := AdministratorNotification{}
	return &this
}

// NewAdministratorNotificationWithDefaults instantiates a new AdministratorNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministratorNotificationWithDefaults() *AdministratorNotification {
	this := AdministratorNotification{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdministratorNotification) GetType() AdministratorNotificationType {
	if o == nil || IsNil(o.Type) {
		var ret AdministratorNotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministratorNotification) GetTypeOk() (*AdministratorNotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdministratorNotification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AdministratorNotificationType and assigns it to the Type field.
func (o *AdministratorNotification) SetType(v AdministratorNotificationType) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdministratorNotification) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdministratorNotification) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *AdministratorNotification) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *AdministratorNotification) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *AdministratorNotification) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *AdministratorNotification) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o AdministratorNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministratorNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableAdministratorNotification struct {
	value *AdministratorNotification
	isSet bool
}

func (v NullableAdministratorNotification) Get() *AdministratorNotification {
	return v.value
}

func (v *NullableAdministratorNotification) Set(val *AdministratorNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorNotification(val *AdministratorNotification) *NullableAdministratorNotification {
	return &NullableAdministratorNotification{value: val, isSet: true}
}

func (v NullableAdministratorNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


