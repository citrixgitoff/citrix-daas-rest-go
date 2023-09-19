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

// checks if the ScopeReferenceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeReferenceModel{}

// ScopeReferenceModel struct for ScopeReferenceModel
type ScopeReferenceModel struct {
	ScopeId *string `json:"ScopeId,omitempty"`
	ScopeName NullableString `json:"ScopeName,omitempty"`
}

// NewScopeReferenceModel instantiates a new ScopeReferenceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeReferenceModel() *ScopeReferenceModel {
	this := ScopeReferenceModel{}
	return &this
}

// NewScopeReferenceModelWithDefaults instantiates a new ScopeReferenceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeReferenceModelWithDefaults() *ScopeReferenceModel {
	this := ScopeReferenceModel{}
	return &this
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise.
func (o *ScopeReferenceModel) GetScopeId() string {
	if o == nil || IsNil(o.ScopeId) {
		var ret string
		return ret
	}
	return *o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeReferenceModel) GetScopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeId) {
		return nil, false
	}
	return o.ScopeId, true
}

// HasScopeId returns a boolean if a field has been set.
func (o *ScopeReferenceModel) HasScopeId() bool {
	if o != nil && !IsNil(o.ScopeId) {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given string and assigns it to the ScopeId field.
func (o *ScopeReferenceModel) SetScopeId(v string) {
	o.ScopeId = &v
}

// GetScopeName returns the ScopeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScopeReferenceModel) GetScopeName() string {
	if o == nil || IsNil(o.ScopeName.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeName.Get()
}

// GetScopeNameOk returns a tuple with the ScopeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScopeReferenceModel) GetScopeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeName.Get(), o.ScopeName.IsSet()
}

// HasScopeName returns a boolean if a field has been set.
func (o *ScopeReferenceModel) HasScopeName() bool {
	if o != nil && o.ScopeName.IsSet() {
		return true
	}

	return false
}

// SetScopeName gets a reference to the given NullableString and assigns it to the ScopeName field.
func (o *ScopeReferenceModel) SetScopeName(v string) {
	o.ScopeName.Set(&v)
}
// SetScopeNameNil sets the value for ScopeName to be an explicit nil
func (o *ScopeReferenceModel) SetScopeNameNil() {
	o.ScopeName.Set(nil)
}

// UnsetScopeName ensures that no value is present for ScopeName, not even an explicit nil
func (o *ScopeReferenceModel) UnsetScopeName() {
	o.ScopeName.Unset()
}

func (o ScopeReferenceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeReferenceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScopeId) {
		toSerialize["ScopeId"] = o.ScopeId
	}
	if o.ScopeName.IsSet() {
		toSerialize["ScopeName"] = o.ScopeName.Get()
	}
	return toSerialize, nil
}

type NullableScopeReferenceModel struct {
	value *ScopeReferenceModel
	isSet bool
}

func (v NullableScopeReferenceModel) Get() *ScopeReferenceModel {
	return v.value
}

func (v *NullableScopeReferenceModel) Set(val *ScopeReferenceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeReferenceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeReferenceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeReferenceModel(val *ScopeReferenceModel) *NullableScopeReferenceModel {
	return &NullableScopeReferenceModel{value: val, isSet: true}
}

func (v NullableScopeReferenceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeReferenceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


