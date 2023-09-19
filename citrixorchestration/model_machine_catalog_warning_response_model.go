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

// checks if the MachineCatalogWarningResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineCatalogWarningResponseModel{}

// MachineCatalogWarningResponseModel Describes a warning on a machine catalog.
type MachineCatalogWarningResponseModel struct {
	Type MachineCatalogWarningType `json:"Type"`
	Subtype *MachineCatalogWarningSubtype `json:"Subtype,omitempty"`
	// Message associated with warning
	Message NullableString `json:"Message,omitempty"`
}

// NewMachineCatalogWarningResponseModel instantiates a new MachineCatalogWarningResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineCatalogWarningResponseModel(type_ MachineCatalogWarningType) *MachineCatalogWarningResponseModel {
	this := MachineCatalogWarningResponseModel{}
	this.Type = type_
	return &this
}

// NewMachineCatalogWarningResponseModelWithDefaults instantiates a new MachineCatalogWarningResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineCatalogWarningResponseModelWithDefaults() *MachineCatalogWarningResponseModel {
	this := MachineCatalogWarningResponseModel{}
	return &this
}

// GetType returns the Type field value
func (o *MachineCatalogWarningResponseModel) GetType() MachineCatalogWarningType {
	if o == nil {
		var ret MachineCatalogWarningType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MachineCatalogWarningResponseModel) GetTypeOk() (*MachineCatalogWarningType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MachineCatalogWarningResponseModel) SetType(v MachineCatalogWarningType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *MachineCatalogWarningResponseModel) GetSubtype() MachineCatalogWarningSubtype {
	if o == nil || IsNil(o.Subtype) {
		var ret MachineCatalogWarningSubtype
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineCatalogWarningResponseModel) GetSubtypeOk() (*MachineCatalogWarningSubtype, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *MachineCatalogWarningResponseModel) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given MachineCatalogWarningSubtype and assigns it to the Subtype field.
func (o *MachineCatalogWarningResponseModel) SetSubtype(v MachineCatalogWarningSubtype) {
	o.Subtype = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineCatalogWarningResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineCatalogWarningResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MachineCatalogWarningResponseModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MachineCatalogWarningResponseModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MachineCatalogWarningResponseModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MachineCatalogWarningResponseModel) UnsetMessage() {
	o.Message.Unset()
}

func (o MachineCatalogWarningResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineCatalogWarningResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Type"] = o.Type
	if !IsNil(o.Subtype) {
		toSerialize["Subtype"] = o.Subtype
	}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableMachineCatalogWarningResponseModel struct {
	value *MachineCatalogWarningResponseModel
	isSet bool
}

func (v NullableMachineCatalogWarningResponseModel) Get() *MachineCatalogWarningResponseModel {
	return v.value
}

func (v *NullableMachineCatalogWarningResponseModel) Set(val *MachineCatalogWarningResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineCatalogWarningResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineCatalogWarningResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineCatalogWarningResponseModel(val *MachineCatalogWarningResponseModel) *NullableMachineCatalogWarningResponseModel {
	return &NullableMachineCatalogWarningResponseModel{value: val, isSet: true}
}

func (v NullableMachineCatalogWarningResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineCatalogWarningResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


