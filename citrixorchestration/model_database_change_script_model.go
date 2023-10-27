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

// checks if the DatabaseChangeScriptModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseChangeScriptModel{}

// DatabaseChangeScriptModel Response object for DatabaseChangeScript.
type DatabaseChangeScriptModel struct {
	// The filename of the text file.
	FileName NullableString `json:"FileName,omitempty"`
	// The contents of the text file.
	Content NullableString `json:"Content,omitempty"`
}

// NewDatabaseChangeScriptModel instantiates a new DatabaseChangeScriptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseChangeScriptModel() *DatabaseChangeScriptModel {
	this := DatabaseChangeScriptModel{}
	return &this
}

// NewDatabaseChangeScriptModelWithDefaults instantiates a new DatabaseChangeScriptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseChangeScriptModelWithDefaults() *DatabaseChangeScriptModel {
	this := DatabaseChangeScriptModel{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatabaseChangeScriptModel) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseChangeScriptModel) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *DatabaseChangeScriptModel) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *DatabaseChangeScriptModel) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *DatabaseChangeScriptModel) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *DatabaseChangeScriptModel) UnsetFileName() {
	o.FileName.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatabaseChangeScriptModel) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseChangeScriptModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *DatabaseChangeScriptModel) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *DatabaseChangeScriptModel) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *DatabaseChangeScriptModel) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *DatabaseChangeScriptModel) UnsetContent() {
	o.Content.Unset()
}

func (o DatabaseChangeScriptModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseChangeScriptModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName.IsSet() {
		toSerialize["FileName"] = o.FileName.Get()
	}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	return toSerialize, nil
}

type NullableDatabaseChangeScriptModel struct {
	value *DatabaseChangeScriptModel
	isSet bool
}

func (v NullableDatabaseChangeScriptModel) Get() *DatabaseChangeScriptModel {
	return v.value
}

func (v *NullableDatabaseChangeScriptModel) Set(val *DatabaseChangeScriptModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseChangeScriptModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseChangeScriptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseChangeScriptModel(val *DatabaseChangeScriptModel) *NullableDatabaseChangeScriptModel {
	return &NullableDatabaseChangeScriptModel{value: val, isSet: true}
}

func (v NullableDatabaseChangeScriptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseChangeScriptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


