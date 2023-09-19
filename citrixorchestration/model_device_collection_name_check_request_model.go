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

// checks if the DeviceCollectionNameCheckRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCollectionNameCheckRequestModel{}

// DeviceCollectionNameCheckRequestModel struct for DeviceCollectionNameCheckRequestModel
type DeviceCollectionNameCheckRequestModel struct {
	// Device Collection name
	Name NullableString `json:"Name,omitempty"`
	// PVS site id
	PvsSiteId NullableString `json:"PvsSiteId,omitempty"`
}

// NewDeviceCollectionNameCheckRequestModel instantiates a new DeviceCollectionNameCheckRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCollectionNameCheckRequestModel() *DeviceCollectionNameCheckRequestModel {
	this := DeviceCollectionNameCheckRequestModel{}
	return &this
}

// NewDeviceCollectionNameCheckRequestModelWithDefaults instantiates a new DeviceCollectionNameCheckRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCollectionNameCheckRequestModelWithDefaults() *DeviceCollectionNameCheckRequestModel {
	this := DeviceCollectionNameCheckRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCollectionNameCheckRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCollectionNameCheckRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceCollectionNameCheckRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceCollectionNameCheckRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceCollectionNameCheckRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceCollectionNameCheckRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetPvsSiteId returns the PvsSiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCollectionNameCheckRequestModel) GetPvsSiteId() string {
	if o == nil || IsNil(o.PvsSiteId.Get()) {
		var ret string
		return ret
	}
	return *o.PvsSiteId.Get()
}

// GetPvsSiteIdOk returns a tuple with the PvsSiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCollectionNameCheckRequestModel) GetPvsSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PvsSiteId.Get(), o.PvsSiteId.IsSet()
}

// HasPvsSiteId returns a boolean if a field has been set.
func (o *DeviceCollectionNameCheckRequestModel) HasPvsSiteId() bool {
	if o != nil && o.PvsSiteId.IsSet() {
		return true
	}

	return false
}

// SetPvsSiteId gets a reference to the given NullableString and assigns it to the PvsSiteId field.
func (o *DeviceCollectionNameCheckRequestModel) SetPvsSiteId(v string) {
	o.PvsSiteId.Set(&v)
}
// SetPvsSiteIdNil sets the value for PvsSiteId to be an explicit nil
func (o *DeviceCollectionNameCheckRequestModel) SetPvsSiteIdNil() {
	o.PvsSiteId.Set(nil)
}

// UnsetPvsSiteId ensures that no value is present for PvsSiteId, not even an explicit nil
func (o *DeviceCollectionNameCheckRequestModel) UnsetPvsSiteId() {
	o.PvsSiteId.Unset()
}

func (o DeviceCollectionNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCollectionNameCheckRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.PvsSiteId.IsSet() {
		toSerialize["PvsSiteId"] = o.PvsSiteId.Get()
	}
	return toSerialize, nil
}

type NullableDeviceCollectionNameCheckRequestModel struct {
	value *DeviceCollectionNameCheckRequestModel
	isSet bool
}

func (v NullableDeviceCollectionNameCheckRequestModel) Get() *DeviceCollectionNameCheckRequestModel {
	return v.value
}

func (v *NullableDeviceCollectionNameCheckRequestModel) Set(val *DeviceCollectionNameCheckRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCollectionNameCheckRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCollectionNameCheckRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCollectionNameCheckRequestModel(val *DeviceCollectionNameCheckRequestModel) *NullableDeviceCollectionNameCheckRequestModel {
	return &NullableDeviceCollectionNameCheckRequestModel{value: val, isSet: true}
}

func (v NullableDeviceCollectionNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCollectionNameCheckRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


