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

// checks if the PagingFileSettingResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingFileSettingResponseModel{}

// PagingFileSettingResponseModel Paging file location like: ?:\\pagefile.sys, C:\\pagefile.sys
type PagingFileSettingResponseModel struct {
	// Paging file location like: ?:\\pagefile.sys, C:\\pagefile.sys
	Location NullableString `json:"Location,omitempty"`
	// Paging file minimum size in MB.
	MinSize *int32 `json:"MinSize,omitempty"`
	// Paging file maximum size in MB.
	MaxSize *int32 `json:"MaxSize,omitempty"`
}

// NewPagingFileSettingResponseModel instantiates a new PagingFileSettingResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingFileSettingResponseModel() *PagingFileSettingResponseModel {
	this := PagingFileSettingResponseModel{}
	return &this
}

// NewPagingFileSettingResponseModelWithDefaults instantiates a new PagingFileSettingResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingFileSettingResponseModelWithDefaults() *PagingFileSettingResponseModel {
	this := PagingFileSettingResponseModel{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PagingFileSettingResponseModel) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingFileSettingResponseModel) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PagingFileSettingResponseModel) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *PagingFileSettingResponseModel) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *PagingFileSettingResponseModel) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PagingFileSettingResponseModel) UnsetLocation() {
	o.Location.Unset()
}

// GetMinSize returns the MinSize field value if set, zero value otherwise.
func (o *PagingFileSettingResponseModel) GetMinSize() int32 {
	if o == nil || IsNil(o.MinSize) {
		var ret int32
		return ret
	}
	return *o.MinSize
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingFileSettingResponseModel) GetMinSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSize) {
		return nil, false
	}
	return o.MinSize, true
}

// HasMinSize returns a boolean if a field has been set.
func (o *PagingFileSettingResponseModel) HasMinSize() bool {
	if o != nil && !IsNil(o.MinSize) {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given int32 and assigns it to the MinSize field.
func (o *PagingFileSettingResponseModel) SetMinSize(v int32) {
	o.MinSize = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *PagingFileSettingResponseModel) GetMaxSize() int32 {
	if o == nil || IsNil(o.MaxSize) {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingFileSettingResponseModel) GetMaxSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSize) {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *PagingFileSettingResponseModel) HasMaxSize() bool {
	if o != nil && !IsNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *PagingFileSettingResponseModel) SetMaxSize(v int32) {
	o.MaxSize = &v
}

func (o PagingFileSettingResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingFileSettingResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["Location"] = o.Location.Get()
	}
	if !IsNil(o.MinSize) {
		toSerialize["MinSize"] = o.MinSize
	}
	if !IsNil(o.MaxSize) {
		toSerialize["MaxSize"] = o.MaxSize
	}
	return toSerialize, nil
}

type NullablePagingFileSettingResponseModel struct {
	value *PagingFileSettingResponseModel
	isSet bool
}

func (v NullablePagingFileSettingResponseModel) Get() *PagingFileSettingResponseModel {
	return v.value
}

func (v *NullablePagingFileSettingResponseModel) Set(val *PagingFileSettingResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingFileSettingResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingFileSettingResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingFileSettingResponseModel(val *PagingFileSettingResponseModel) *NullablePagingFileSettingResponseModel {
	return &NullablePagingFileSettingResponseModel{value: val, isSet: true}
}

func (v NullablePagingFileSettingResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingFileSettingResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


