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

// checks if the CatalogNameCheckRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogNameCheckRequestModel{}

// CatalogNameCheckRequestModel Request Model for Checking Catalog Name with Admin Folder
type CatalogNameCheckRequestModel struct {
	// Catalog Name with Admin Folder
	Name *string `json:"Name,omitempty"`
}

// NewCatalogNameCheckRequestModel instantiates a new CatalogNameCheckRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogNameCheckRequestModel() *CatalogNameCheckRequestModel {
	this := CatalogNameCheckRequestModel{}
	return &this
}

// NewCatalogNameCheckRequestModelWithDefaults instantiates a new CatalogNameCheckRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogNameCheckRequestModelWithDefaults() *CatalogNameCheckRequestModel {
	this := CatalogNameCheckRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogNameCheckRequestModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogNameCheckRequestModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogNameCheckRequestModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogNameCheckRequestModel) SetName(v string) {
	o.Name = &v
}

func (o CatalogNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogNameCheckRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCatalogNameCheckRequestModel struct {
	value *CatalogNameCheckRequestModel
	isSet bool
}

func (v NullableCatalogNameCheckRequestModel) Get() *CatalogNameCheckRequestModel {
	return v.value
}

func (v *NullableCatalogNameCheckRequestModel) Set(val *CatalogNameCheckRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogNameCheckRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogNameCheckRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogNameCheckRequestModel(val *CatalogNameCheckRequestModel) *NullableCatalogNameCheckRequestModel {
	return &NullableCatalogNameCheckRequestModel{value: val, isSet: true}
}

func (v NullableCatalogNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogNameCheckRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


