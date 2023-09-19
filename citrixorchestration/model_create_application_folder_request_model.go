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

// checks if the CreateApplicationFolderRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApplicationFolderRequestModel{}

// CreateApplicationFolderRequestModel Request object for creating an application folder.
type CreateApplicationFolderRequestModel struct {
	// Name of the folder to create. May be omitted if Path is used; otherwise required.  If both  and  are specified, a folder with the specified name will be created as a child of the specified path.
	Name NullableString `json:"Name,omitempty"`
	// Path where the folder should be created. Exclusive with ParentId.  One or the other property may be used but not both.  Path to where the folder should be created. Any folders in the path which don't exist will be created.  If not specified, default is the root folder path.
	Path NullableString `json:"Path,omitempty"`
	// Parent folder where the folder should be created. Exclusive with Path.  One or the other property may be used but not both.  If not specified, default is the root folder path.
	ParentId NullableString `json:"ParentId,omitempty"`
}

// NewCreateApplicationFolderRequestModel instantiates a new CreateApplicationFolderRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApplicationFolderRequestModel() *CreateApplicationFolderRequestModel {
	this := CreateApplicationFolderRequestModel{}
	return &this
}

// NewCreateApplicationFolderRequestModelWithDefaults instantiates a new CreateApplicationFolderRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApplicationFolderRequestModelWithDefaults() *CreateApplicationFolderRequestModel {
	this := CreateApplicationFolderRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateApplicationFolderRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateApplicationFolderRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateApplicationFolderRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateApplicationFolderRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateApplicationFolderRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateApplicationFolderRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateApplicationFolderRequestModel) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateApplicationFolderRequestModel) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *CreateApplicationFolderRequestModel) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *CreateApplicationFolderRequestModel) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *CreateApplicationFolderRequestModel) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *CreateApplicationFolderRequestModel) UnsetPath() {
	o.Path.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateApplicationFolderRequestModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateApplicationFolderRequestModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateApplicationFolderRequestModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *CreateApplicationFolderRequestModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *CreateApplicationFolderRequestModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *CreateApplicationFolderRequestModel) UnsetParentId() {
	o.ParentId.Unset()
}

func (o CreateApplicationFolderRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApplicationFolderRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["ParentId"] = o.ParentId.Get()
	}
	return toSerialize, nil
}

type NullableCreateApplicationFolderRequestModel struct {
	value *CreateApplicationFolderRequestModel
	isSet bool
}

func (v NullableCreateApplicationFolderRequestModel) Get() *CreateApplicationFolderRequestModel {
	return v.value
}

func (v *NullableCreateApplicationFolderRequestModel) Set(val *CreateApplicationFolderRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationFolderRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationFolderRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationFolderRequestModel(val *CreateApplicationFolderRequestModel) *NullableCreateApplicationFolderRequestModel {
	return &NullableCreateApplicationFolderRequestModel{value: val, isSet: true}
}

func (v NullableCreateApplicationFolderRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationFolderRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


