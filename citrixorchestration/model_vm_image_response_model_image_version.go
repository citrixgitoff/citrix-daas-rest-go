/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the VMImageResponseModelImageVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMImageResponseModelImageVersion{}

// VMImageResponseModelImageVersion Reference to image version on the hypervisor.
type VMImageResponseModelImageVersion struct {
	// The image definition name.
	ImageDefinitionName *string `json:"ImageDefinitionName,omitempty"`
	// The image version number.
	ImageVersionNumber *string `json:"ImageVersionNumber,omitempty"`
	// The image version uid.
	ImageVersionUid *string `json:"ImageVersionUid,omitempty"`
}

// NewVMImageResponseModelImageVersion instantiates a new VMImageResponseModelImageVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMImageResponseModelImageVersion() *VMImageResponseModelImageVersion {
	this := VMImageResponseModelImageVersion{}
	return &this
}

// NewVMImageResponseModelImageVersionWithDefaults instantiates a new VMImageResponseModelImageVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMImageResponseModelImageVersionWithDefaults() *VMImageResponseModelImageVersion {
	this := VMImageResponseModelImageVersion{}
	return &this
}

// GetImageDefinitionName returns the ImageDefinitionName field value if set, zero value otherwise.
func (o *VMImageResponseModelImageVersion) GetImageDefinitionName() string {
	if o == nil || IsNil(o.ImageDefinitionName) {
		var ret string
		return ret
	}
	return *o.ImageDefinitionName
}

// GetImageDefinitionNameOk returns a tuple with the ImageDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMImageResponseModelImageVersion) GetImageDefinitionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDefinitionName) {
		return nil, false
	}
	return o.ImageDefinitionName, true
}

// HasImageDefinitionName returns a boolean if a field has been set.
func (o *VMImageResponseModelImageVersion) HasImageDefinitionName() bool {
	if o != nil && !IsNil(o.ImageDefinitionName) {
		return true
	}

	return false
}

// SetImageDefinitionName gets a reference to the given string and assigns it to the ImageDefinitionName field.
func (o *VMImageResponseModelImageVersion) SetImageDefinitionName(v string) {
	o.ImageDefinitionName = &v
}

// GetImageVersionNumber returns the ImageVersionNumber field value if set, zero value otherwise.
func (o *VMImageResponseModelImageVersion) GetImageVersionNumber() string {
	if o == nil || IsNil(o.ImageVersionNumber) {
		var ret string
		return ret
	}
	return *o.ImageVersionNumber
}

// GetImageVersionNumberOk returns a tuple with the ImageVersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMImageResponseModelImageVersion) GetImageVersionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ImageVersionNumber) {
		return nil, false
	}
	return o.ImageVersionNumber, true
}

// HasImageVersionNumber returns a boolean if a field has been set.
func (o *VMImageResponseModelImageVersion) HasImageVersionNumber() bool {
	if o != nil && !IsNil(o.ImageVersionNumber) {
		return true
	}

	return false
}

// SetImageVersionNumber gets a reference to the given string and assigns it to the ImageVersionNumber field.
func (o *VMImageResponseModelImageVersion) SetImageVersionNumber(v string) {
	o.ImageVersionNumber = &v
}

// GetImageVersionUid returns the ImageVersionUid field value if set, zero value otherwise.
func (o *VMImageResponseModelImageVersion) GetImageVersionUid() string {
	if o == nil || IsNil(o.ImageVersionUid) {
		var ret string
		return ret
	}
	return *o.ImageVersionUid
}

// GetImageVersionUidOk returns a tuple with the ImageVersionUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMImageResponseModelImageVersion) GetImageVersionUidOk() (*string, bool) {
	if o == nil || IsNil(o.ImageVersionUid) {
		return nil, false
	}
	return o.ImageVersionUid, true
}

// HasImageVersionUid returns a boolean if a field has been set.
func (o *VMImageResponseModelImageVersion) HasImageVersionUid() bool {
	if o != nil && !IsNil(o.ImageVersionUid) {
		return true
	}

	return false
}

// SetImageVersionUid gets a reference to the given string and assigns it to the ImageVersionUid field.
func (o *VMImageResponseModelImageVersion) SetImageVersionUid(v string) {
	o.ImageVersionUid = &v
}

func (o VMImageResponseModelImageVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMImageResponseModelImageVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageDefinitionName) {
		toSerialize["ImageDefinitionName"] = o.ImageDefinitionName
	}
	if !IsNil(o.ImageVersionNumber) {
		toSerialize["ImageVersionNumber"] = o.ImageVersionNumber
	}
	if !IsNil(o.ImageVersionUid) {
		toSerialize["ImageVersionUid"] = o.ImageVersionUid
	}
	return toSerialize, nil
}

type NullableVMImageResponseModelImageVersion struct {
	value *VMImageResponseModelImageVersion
	isSet bool
}

func (v NullableVMImageResponseModelImageVersion) Get() *VMImageResponseModelImageVersion {
	return v.value
}

func (v *NullableVMImageResponseModelImageVersion) Set(val *VMImageResponseModelImageVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableVMImageResponseModelImageVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVMImageResponseModelImageVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMImageResponseModelImageVersion(val *VMImageResponseModelImageVersion) *NullableVMImageResponseModelImageVersion {
	return &NullableVMImageResponseModelImageVersion{value: val, isSet: true}
}

func (v NullableVMImageResponseModelImageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMImageResponseModelImageVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


