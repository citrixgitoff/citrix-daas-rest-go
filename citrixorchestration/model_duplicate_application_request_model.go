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

// checks if the DuplicateApplicationRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicateApplicationRequestModel{}

// DuplicateApplicationRequestModel The Duplicate Application Request Model.
type DuplicateApplicationRequestModel struct {
	// Optional name for the duplicated application.  If not specified, the system will choose a new name automatically, based on the original application name. If specified, must be unique within the site.
	NewName NullableString `json:"NewName,omitempty"`
	// Optional folder in which to create the duplicated application.  If not specified, the new application will be created in the same folder as the original. May be specified as either Path or Id.  If specified as a path, and the target does not exist, it will be automatically created.
	NewFolder NullableString `json:"NewFolder,omitempty"`
	// Optional. If not specified, or specified as `false`, the new application will have the same visibility and enabled state as the original application.  This may be undesirable; if the original application is visible, it means the new duplicate will also be immediately visible to end users. Setting this to `true` causes the new duplicate to start out invisible and disabled, allowing it to be further modified before making it visible to end users.
	CreateDisabled NullableBool `json:"CreateDisabled,omitempty"`
}

// NewDuplicateApplicationRequestModel instantiates a new DuplicateApplicationRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicateApplicationRequestModel() *DuplicateApplicationRequestModel {
	this := DuplicateApplicationRequestModel{}
	var createDisabled bool = false
	this.CreateDisabled = *NewNullableBool(&createDisabled)
	return &this
}

// NewDuplicateApplicationRequestModelWithDefaults instantiates a new DuplicateApplicationRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicateApplicationRequestModelWithDefaults() *DuplicateApplicationRequestModel {
	this := DuplicateApplicationRequestModel{}
	var createDisabled bool = false
	this.CreateDisabled = *NewNullableBool(&createDisabled)
	return &this
}

// GetNewName returns the NewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateApplicationRequestModel) GetNewName() string {
	if o == nil || IsNil(o.NewName.Get()) {
		var ret string
		return ret
	}
	return *o.NewName.Get()
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateApplicationRequestModel) GetNewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewName.Get(), o.NewName.IsSet()
}

// HasNewName returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasNewName() bool {
	if o != nil && o.NewName.IsSet() {
		return true
	}

	return false
}

// SetNewName gets a reference to the given NullableString and assigns it to the NewName field.
func (o *DuplicateApplicationRequestModel) SetNewName(v string) {
	o.NewName.Set(&v)
}
// SetNewNameNil sets the value for NewName to be an explicit nil
func (o *DuplicateApplicationRequestModel) SetNewNameNil() {
	o.NewName.Set(nil)
}

// UnsetNewName ensures that no value is present for NewName, not even an explicit nil
func (o *DuplicateApplicationRequestModel) UnsetNewName() {
	o.NewName.Unset()
}

// GetNewFolder returns the NewFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateApplicationRequestModel) GetNewFolder() string {
	if o == nil || IsNil(o.NewFolder.Get()) {
		var ret string
		return ret
	}
	return *o.NewFolder.Get()
}

// GetNewFolderOk returns a tuple with the NewFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateApplicationRequestModel) GetNewFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewFolder.Get(), o.NewFolder.IsSet()
}

// HasNewFolder returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasNewFolder() bool {
	if o != nil && o.NewFolder.IsSet() {
		return true
	}

	return false
}

// SetNewFolder gets a reference to the given NullableString and assigns it to the NewFolder field.
func (o *DuplicateApplicationRequestModel) SetNewFolder(v string) {
	o.NewFolder.Set(&v)
}
// SetNewFolderNil sets the value for NewFolder to be an explicit nil
func (o *DuplicateApplicationRequestModel) SetNewFolderNil() {
	o.NewFolder.Set(nil)
}

// UnsetNewFolder ensures that no value is present for NewFolder, not even an explicit nil
func (o *DuplicateApplicationRequestModel) UnsetNewFolder() {
	o.NewFolder.Unset()
}

// GetCreateDisabled returns the CreateDisabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateApplicationRequestModel) GetCreateDisabled() bool {
	if o == nil || IsNil(o.CreateDisabled.Get()) {
		var ret bool
		return ret
	}
	return *o.CreateDisabled.Get()
}

// GetCreateDisabledOk returns a tuple with the CreateDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateApplicationRequestModel) GetCreateDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateDisabled.Get(), o.CreateDisabled.IsSet()
}

// HasCreateDisabled returns a boolean if a field has been set.
func (o *DuplicateApplicationRequestModel) HasCreateDisabled() bool {
	if o != nil && o.CreateDisabled.IsSet() {
		return true
	}

	return false
}

// SetCreateDisabled gets a reference to the given NullableBool and assigns it to the CreateDisabled field.
func (o *DuplicateApplicationRequestModel) SetCreateDisabled(v bool) {
	o.CreateDisabled.Set(&v)
}
// SetCreateDisabledNil sets the value for CreateDisabled to be an explicit nil
func (o *DuplicateApplicationRequestModel) SetCreateDisabledNil() {
	o.CreateDisabled.Set(nil)
}

// UnsetCreateDisabled ensures that no value is present for CreateDisabled, not even an explicit nil
func (o *DuplicateApplicationRequestModel) UnsetCreateDisabled() {
	o.CreateDisabled.Unset()
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
	if o.NewName.IsSet() {
		toSerialize["NewName"] = o.NewName.Get()
	}
	if o.NewFolder.IsSet() {
		toSerialize["NewFolder"] = o.NewFolder.Get()
	}
	if o.CreateDisabled.IsSet() {
		toSerialize["CreateDisabled"] = o.CreateDisabled.Get()
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


