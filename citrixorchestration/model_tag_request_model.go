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

// checks if the TagRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagRequestModel{}

// TagRequestModel Request object for tags.
type TagRequestModel struct {
	// Name of the tag.
	Name *string `json:"Name,omitempty"`
	// Description of the tag.
	Description *string `json:"Description,omitempty"`
	// Unique Id of the tag.
	Uuid *string `json:"Uuid,omitempty"`
	// Name of the Scopes.
	Scopes []string `json:"Scopes,omitempty"`
}

// NewTagRequestModel instantiates a new TagRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRequestModel() *TagRequestModel {
	this := TagRequestModel{}
	return &this
}

// NewTagRequestModelWithDefaults instantiates a new TagRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagRequestModelWithDefaults() *TagRequestModel {
	this := TagRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagRequestModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRequestModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagRequestModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagRequestModel) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagRequestModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagRequestModel) SetDescription(v string) {
	o.Description = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TagRequestModel) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRequestModel) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TagRequestModel) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TagRequestModel) SetUuid(v string) {
	o.Uuid = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TagRequestModel) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRequestModel) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TagRequestModel) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TagRequestModel) SetScopes(v []string) {
	o.Scopes = v
}

func (o TagRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.Scopes) {
		toSerialize["Scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableTagRequestModel struct {
	value *TagRequestModel
	isSet bool
}

func (v NullableTagRequestModel) Get() *TagRequestModel {
	return v.value
}

func (v *NullableTagRequestModel) Set(val *TagRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRequestModel(val *TagRequestModel) *NullableTagRequestModel {
	return &NullableTagRequestModel{value: val, isSet: true}
}

func (v NullableTagRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


