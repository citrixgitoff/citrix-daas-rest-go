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

// checks if the RefResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefResponseModel{}

// RefResponseModel Response object for an object reference.
type RefResponseModel struct {
	// Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility
	Id NullableString `json:"Id,omitempty"`
	// DEPRECATED. Use Id.
	Uid NullableInt32 `json:"Uid,omitempty"`
	// Name of the object.
	Name NullableString `json:"Name,omitempty"`
}

// NewRefResponseModel instantiates a new RefResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefResponseModel() *RefResponseModel {
	this := RefResponseModel{}
	return &this
}

// NewRefResponseModelWithDefaults instantiates a new RefResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefResponseModelWithDefaults() *RefResponseModel {
	this := RefResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RefResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *RefResponseModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RefResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RefResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefResponseModel) GetUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *RefResponseModel) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *RefResponseModel) SetUid(v int32) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *RefResponseModel) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *RefResponseModel) UnsetUid() {
	o.Uid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RefResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RefResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RefResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RefResponseModel) UnsetName() {
	o.Name.Unset()
}

func (o RefResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["Uid"] = o.Uid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableRefResponseModel struct {
	value *RefResponseModel
	isSet bool
}

func (v NullableRefResponseModel) Get() *RefResponseModel {
	return v.value
}

func (v *NullableRefResponseModel) Set(val *RefResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRefResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRefResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefResponseModel(val *RefResponseModel) *NullableRefResponseModel {
	return &NullableRefResponseModel{value: val, isSet: true}
}

func (v NullableRefResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


