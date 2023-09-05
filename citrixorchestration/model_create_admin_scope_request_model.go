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

// checks if the CreateAdminScopeRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAdminScopeRequestModel{}

// CreateAdminScopeRequestModel struct for CreateAdminScopeRequestModel
type CreateAdminScopeRequestModel struct {
	// The name of the scope. Name is globally unique.
	Name string `json:"Name"`
	// The description of the admin scope.
	Description *string `json:"Description,omitempty"`
	// Indicates if it is tenant scope.
	IsTenantScope *bool `json:"IsTenantScope,omitempty"`
	// Scoped objects associated with the scope object.
	ScopedObjects []ScopedObjectRequestModel `json:"ScopedObjects,omitempty"`
}

// NewCreateAdminScopeRequestModel instantiates a new CreateAdminScopeRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAdminScopeRequestModel(name string) *CreateAdminScopeRequestModel {
	this := CreateAdminScopeRequestModel{}
	this.Name = name
	return &this
}

// NewCreateAdminScopeRequestModelWithDefaults instantiates a new CreateAdminScopeRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAdminScopeRequestModelWithDefaults() *CreateAdminScopeRequestModel {
	this := CreateAdminScopeRequestModel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAdminScopeRequestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAdminScopeRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAdminScopeRequestModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAdminScopeRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdminScopeRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAdminScopeRequestModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAdminScopeRequestModel) SetDescription(v string) {
	o.Description = &v
}

// GetIsTenantScope returns the IsTenantScope field value if set, zero value otherwise.
func (o *CreateAdminScopeRequestModel) GetIsTenantScope() bool {
	if o == nil || IsNil(o.IsTenantScope) {
		var ret bool
		return ret
	}
	return *o.IsTenantScope
}

// GetIsTenantScopeOk returns a tuple with the IsTenantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdminScopeRequestModel) GetIsTenantScopeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTenantScope) {
		return nil, false
	}
	return o.IsTenantScope, true
}

// HasIsTenantScope returns a boolean if a field has been set.
func (o *CreateAdminScopeRequestModel) HasIsTenantScope() bool {
	if o != nil && !IsNil(o.IsTenantScope) {
		return true
	}

	return false
}

// SetIsTenantScope gets a reference to the given bool and assigns it to the IsTenantScope field.
func (o *CreateAdminScopeRequestModel) SetIsTenantScope(v bool) {
	o.IsTenantScope = &v
}

// GetScopedObjects returns the ScopedObjects field value if set, zero value otherwise.
func (o *CreateAdminScopeRequestModel) GetScopedObjects() []ScopedObjectRequestModel {
	if o == nil || IsNil(o.ScopedObjects) {
		var ret []ScopedObjectRequestModel
		return ret
	}
	return o.ScopedObjects
}

// GetScopedObjectsOk returns a tuple with the ScopedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdminScopeRequestModel) GetScopedObjectsOk() ([]ScopedObjectRequestModel, bool) {
	if o == nil || IsNil(o.ScopedObjects) {
		return nil, false
	}
	return o.ScopedObjects, true
}

// HasScopedObjects returns a boolean if a field has been set.
func (o *CreateAdminScopeRequestModel) HasScopedObjects() bool {
	if o != nil && !IsNil(o.ScopedObjects) {
		return true
	}

	return false
}

// SetScopedObjects gets a reference to the given []ScopedObjectRequestModel and assigns it to the ScopedObjects field.
func (o *CreateAdminScopeRequestModel) SetScopedObjects(v []ScopedObjectRequestModel) {
	o.ScopedObjects = v
}

func (o CreateAdminScopeRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAdminScopeRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.IsTenantScope) {
		toSerialize["IsTenantScope"] = o.IsTenantScope
	}
	if !IsNil(o.ScopedObjects) {
		toSerialize["ScopedObjects"] = o.ScopedObjects
	}
	return toSerialize, nil
}

type NullableCreateAdminScopeRequestModel struct {
	value *CreateAdminScopeRequestModel
	isSet bool
}

func (v NullableCreateAdminScopeRequestModel) Get() *CreateAdminScopeRequestModel {
	return v.value
}

func (v *NullableCreateAdminScopeRequestModel) Set(val *CreateAdminScopeRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdminScopeRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdminScopeRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdminScopeRequestModel(val *CreateAdminScopeRequestModel) *NullableCreateAdminScopeRequestModel {
	return &NullableCreateAdminScopeRequestModel{value: val, isSet: true}
}

func (v NullableCreateAdminScopeRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAdminScopeRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


