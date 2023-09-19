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

// checks if the StoreFrontServerRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreFrontServerRequestModel{}

// StoreFrontServerRequestModel Request object for specifying StoreFront URLs to configure within Receivers when they are hosted on a VDA.
type StoreFrontServerRequestModel struct {
	// ID of an existing StoreFront server.  If specified, this must be the only property specified in the request model.
	Id NullableString `json:"Id,omitempty"`
	// Name of the StoreFront server.
	Name NullableString `json:"Name,omitempty"`
	// Description of the StoreFront server.
	Description NullableString `json:"Description,omitempty"`
	// Url of the StoreFront server.
	Url NullableString `json:"Url,omitempty"`
	// Whether the StoreFront server is enabled.  Disabled StoreFront servers will not have thier URLs added to hosted receiver.
	Enabled NullableBool `json:"Enabled,omitempty"`
}

// NewStoreFrontServerRequestModel instantiates a new StoreFrontServerRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreFrontServerRequestModel() *StoreFrontServerRequestModel {
	this := StoreFrontServerRequestModel{}
	var enabled bool = true
	this.Enabled = *NewNullableBool(&enabled)
	return &this
}

// NewStoreFrontServerRequestModelWithDefaults instantiates a new StoreFrontServerRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreFrontServerRequestModelWithDefaults() *StoreFrontServerRequestModel {
	this := StoreFrontServerRequestModel{}
	var enabled bool = true
	this.Enabled = *NewNullableBool(&enabled)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoreFrontServerRequestModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreFrontServerRequestModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *StoreFrontServerRequestModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *StoreFrontServerRequestModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *StoreFrontServerRequestModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *StoreFrontServerRequestModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoreFrontServerRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreFrontServerRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StoreFrontServerRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StoreFrontServerRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *StoreFrontServerRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StoreFrontServerRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoreFrontServerRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreFrontServerRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StoreFrontServerRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StoreFrontServerRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StoreFrontServerRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StoreFrontServerRequestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoreFrontServerRequestModel) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreFrontServerRequestModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *StoreFrontServerRequestModel) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *StoreFrontServerRequestModel) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *StoreFrontServerRequestModel) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *StoreFrontServerRequestModel) UnsetUrl() {
	o.Url.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoreFrontServerRequestModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoreFrontServerRequestModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *StoreFrontServerRequestModel) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *StoreFrontServerRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *StoreFrontServerRequestModel) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *StoreFrontServerRequestModel) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o StoreFrontServerRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreFrontServerRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Url.IsSet() {
		toSerialize["Url"] = o.Url.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableStoreFrontServerRequestModel struct {
	value *StoreFrontServerRequestModel
	isSet bool
}

func (v NullableStoreFrontServerRequestModel) Get() *StoreFrontServerRequestModel {
	return v.value
}

func (v *NullableStoreFrontServerRequestModel) Set(val *StoreFrontServerRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreFrontServerRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreFrontServerRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreFrontServerRequestModel(val *StoreFrontServerRequestModel) *NullableStoreFrontServerRequestModel {
	return &NullableStoreFrontServerRequestModel{value: val, isSet: true}
}

func (v NullableStoreFrontServerRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreFrontServerRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


