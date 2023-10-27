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

// checks if the EditAdminScopeRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAdminScopeRequestModel{}

// EditAdminScopeRequestModel Request object used when editing an admin scope.
type EditAdminScopeRequestModel struct {
	// The name of the scope. Name is globally unique.
	Name NullableString `json:"Name,omitempty"`
	// The description of the admin scope.
	Description NullableString `json:"Description,omitempty"`
	// Scoped objects associated with the scope object.
	ScopedObjects []ScopedObjectRequestModel `json:"ScopedObjects,omitempty"`
}

// NewEditAdminScopeRequestModel instantiates a new EditAdminScopeRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAdminScopeRequestModel() *EditAdminScopeRequestModel {
	this := EditAdminScopeRequestModel{}
	return &this
}

// NewEditAdminScopeRequestModelWithDefaults instantiates a new EditAdminScopeRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAdminScopeRequestModelWithDefaults() *EditAdminScopeRequestModel {
	this := EditAdminScopeRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAdminScopeRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAdminScopeRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditAdminScopeRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditAdminScopeRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EditAdminScopeRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditAdminScopeRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAdminScopeRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAdminScopeRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EditAdminScopeRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EditAdminScopeRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EditAdminScopeRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EditAdminScopeRequestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetScopedObjects returns the ScopedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAdminScopeRequestModel) GetScopedObjects() []ScopedObjectRequestModel {
	if o == nil {
		var ret []ScopedObjectRequestModel
		return ret
	}
	return o.ScopedObjects
}

// GetScopedObjectsOk returns a tuple with the ScopedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAdminScopeRequestModel) GetScopedObjectsOk() ([]ScopedObjectRequestModel, bool) {
	if o == nil || IsNil(o.ScopedObjects) {
		return nil, false
	}
	return o.ScopedObjects, true
}

// HasScopedObjects returns a boolean if a field has been set.
func (o *EditAdminScopeRequestModel) HasScopedObjects() bool {
	if o != nil && IsNil(o.ScopedObjects) {
		return true
	}

	return false
}

// SetScopedObjects gets a reference to the given []ScopedObjectRequestModel and assigns it to the ScopedObjects field.
func (o *EditAdminScopeRequestModel) SetScopedObjects(v []ScopedObjectRequestModel) {
	o.ScopedObjects = v
}

func (o EditAdminScopeRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAdminScopeRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.ScopedObjects != nil {
		toSerialize["ScopedObjects"] = o.ScopedObjects
	}
	return toSerialize, nil
}

type NullableEditAdminScopeRequestModel struct {
	value *EditAdminScopeRequestModel
	isSet bool
}

func (v NullableEditAdminScopeRequestModel) Get() *EditAdminScopeRequestModel {
	return v.value
}

func (v *NullableEditAdminScopeRequestModel) Set(val *EditAdminScopeRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAdminScopeRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAdminScopeRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAdminScopeRequestModel(val *EditAdminScopeRequestModel) *NullableEditAdminScopeRequestModel {
	return &NullableEditAdminScopeRequestModel{value: val, isSet: true}
}

func (v NullableEditAdminScopeRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAdminScopeRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


