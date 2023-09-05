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

// checks if the DuplicateApplicationRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicateApplicationRequestModel{}

// DuplicateApplicationRequestModel The Duplicate Application Request Model.
type DuplicateApplicationRequestModel struct {
	// Optional name for the duplicated application.  If not specified, the system will choose a new name automatically, based on the original application name. If specified, must be unique within the site.
	NewName *string `json:"NewName,omitempty"`
	// Optional folder in which to create the duplicated application.  If not specified, the new application will be created in the same folder as the original. May be specified as either Path or Id.  If specified as a path, and the target does not exist, it will be automatically created.
	NewFolder *string `json:"NewFolder,omitempty"`
	// Optional. If not specified, or specified as `false`, the new application will have the same visibility and enabled state as the original application.  This may be undesirable; if the original application is visible, it means the new duplicate will also be immediately visible to end users. Setting this to `true` causes the new duplicate to start out invisible and disabled, allowing it to be further modified before making it visible to end users.
	CreateDisabled *bool `json:"CreateDisabled,omitempty"`
}

// NewDuplicateApplicationRequestModel instantiates a new DuplicateApplicationRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicateApplicationRequestModel() *DuplicateApplicationRequestModel {
	this := DuplicateApplicationRequestModel{}
	var createDisabled bool = false
	this.CreateDisabled = &createDisabled
	return &this
}

// NewDuplicateApplicationRequestModelWithDefaults instantiates a new DuplicateApplicationRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicateApplicationRequestModelWithDefaults() *DuplicateApplicationRequestModel {
	this := DuplicateApplicationRequestModel{}
	var createDisabled bool = false
	this.CreateDisabled = &createDisabled
	return &this
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DuplicateApplicationRequestModel) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicateApplicationRequestModel) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DuplicateApplicationRequestModel) SetNewName(v string) {
	o.NewName = &v
}

// GetNewFolder returns the NewFolder field value if set, zero value otherwise.
func (o *DuplicateApplicationRequestModel) GetNewFolder() string {
	if o == nil || IsNil(o.NewFolder) {
		var ret string
		return ret
	}
	return *o.NewFolder
}

// GetNewFolderOk returns a tuple with the NewFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicateApplicationRequestModel) GetNewFolderOk() (*string, bool) {
	if o == nil || IsNil(o.NewFolder) {
		return nil, false
	}
	return o.NewFolder, true
}

// HasNewFolder returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasNewFolder() bool {
	if o != nil && !IsNil(o.NewFolder) {
		return true
	}

	return false
}

// SetNewFolder gets a reference to the given string and assigns it to the NewFolder field.
func (o *DuplicateApplicationRequestModel) SetNewFolder(v string) {
	o.NewFolder = &v
}

// GetCreateDisabled returns the CreateDisabled field value if set, zero value otherwise.
func (o *DuplicateApplicationRequestModel) GetCreateDisabled() bool {
	if o == nil || IsNil(o.CreateDisabled) {
		var ret bool
		return ret
	}
	return *o.CreateDisabled
}

// GetCreateDisabledOk returns a tuple with the CreateDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicateApplicationRequestModel) GetCreateDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateDisabled) {
		return nil, false
	}
	return o.CreateDisabled, true
}

// HasCreateDisabled returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasCreateDisabled() bool {
	if o != nil && !IsNil(o.CreateDisabled) {
		return true
	}

	return false
}

// SetCreateDisabled gets a reference to the given bool and assigns it to the CreateDisabled field.
func (o *DuplicateApplicationRequestModel) SetCreateDisabled(v bool) {
	o.CreateDisabled = &v
}

func (o DuplicateApplicationRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicateApplicationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewName) {
		toSerialize["NewName"] = o.NewName
	}
	if !IsNil(o.NewFolder) {
		toSerialize["NewFolder"] = o.NewFolder
	}
	if !IsNil(o.CreateDisabled) {
		toSerialize["CreateDisabled"] = o.CreateDisabled
	}
	return toSerialize, nil
}

type NullableDuplicateApplicationRequestModel struct {
	value *DuplicateApplicationRequestModel
	isSet bool
}

func (v NullableDuplicateApplicationRequestModel) Get() *DuplicateApplicationRequestModel {
	return v.value
}

func (v *NullableDuplicateApplicationRequestModel) Set(val *DuplicateApplicationRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicateApplicationRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicateApplicationRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicateApplicationRequestModel(val *DuplicateApplicationRequestModel) *NullableDuplicateApplicationRequestModel {
	return &NullableDuplicateApplicationRequestModel{value: val, isSet: true}
}

func (v NullableDuplicateApplicationRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicateApplicationRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


