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

// checks if the VDAComponentsSelectionValidationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDAComponentsSelectionValidationResponseModel{}

// VDAComponentsSelectionValidationResponseModel Response model of VDA upgrade components and features related to the Machine Catalog.
type VDAComponentsSelectionValidationResponseModel struct {
	// Uid of the catalog.
	Uid *int32 `json:"Uid,omitempty"`
	// Id of the machine catalog.
	Id string `json:"Id"`
	// Name of the catalog.
	Name string `json:"Name"`
	// Validation result of VDA components selection of the catalog.
	IsValid *bool `json:"IsValid,omitempty"`
	// Validation message of VDA components selection of the catalog.
	Message NullableString `json:"Message,omitempty"`
}

// NewVDAComponentsSelectionValidationResponseModel instantiates a new VDAComponentsSelectionValidationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDAComponentsSelectionValidationResponseModel(id string, name string) *VDAComponentsSelectionValidationResponseModel {
	this := VDAComponentsSelectionValidationResponseModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewVDAComponentsSelectionValidationResponseModelWithDefaults instantiates a new VDAComponentsSelectionValidationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDAComponentsSelectionValidationResponseModelWithDefaults() *VDAComponentsSelectionValidationResponseModel {
	this := VDAComponentsSelectionValidationResponseModel{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *VDAComponentsSelectionValidationResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDAComponentsSelectionValidationResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *VDAComponentsSelectionValidationResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *VDAComponentsSelectionValidationResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetId returns the Id field value
func (o *VDAComponentsSelectionValidationResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VDAComponentsSelectionValidationResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VDAComponentsSelectionValidationResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VDAComponentsSelectionValidationResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VDAComponentsSelectionValidationResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VDAComponentsSelectionValidationResponseModel) SetName(v string) {
	o.Name = v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *VDAComponentsSelectionValidationResponseModel) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDAComponentsSelectionValidationResponseModel) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *VDAComponentsSelectionValidationResponseModel) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *VDAComponentsSelectionValidationResponseModel) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDAComponentsSelectionValidationResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDAComponentsSelectionValidationResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *VDAComponentsSelectionValidationResponseModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *VDAComponentsSelectionValidationResponseModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *VDAComponentsSelectionValidationResponseModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *VDAComponentsSelectionValidationResponseModel) UnsetMessage() {
	o.Message.Unset()
}

func (o VDAComponentsSelectionValidationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDAComponentsSelectionValidationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	if !IsNil(o.IsValid) {
		toSerialize["IsValid"] = o.IsValid
	}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableVDAComponentsSelectionValidationResponseModel struct {
	value *VDAComponentsSelectionValidationResponseModel
	isSet bool
}

func (v NullableVDAComponentsSelectionValidationResponseModel) Get() *VDAComponentsSelectionValidationResponseModel {
	return v.value
}

func (v *NullableVDAComponentsSelectionValidationResponseModel) Set(val *VDAComponentsSelectionValidationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVDAComponentsSelectionValidationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVDAComponentsSelectionValidationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDAComponentsSelectionValidationResponseModel(val *VDAComponentsSelectionValidationResponseModel) *NullableVDAComponentsSelectionValidationResponseModel {
	return &NullableVDAComponentsSelectionValidationResponseModel{value: val, isSet: true}
}

func (v NullableVDAComponentsSelectionValidationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDAComponentsSelectionValidationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


