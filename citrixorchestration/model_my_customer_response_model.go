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

// checks if the MyCustomerResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyCustomerResponseModel{}

// MyCustomerResponseModel Response object indicating a customer who the currently-logged-in admin has access to.
type MyCustomerResponseModel struct {
	// Id of the customer. internally: the CC customer id / short-name, or for on-prem this could be any generated guid.
	Id string `json:"Id"`
	// Human-readable name of the customer.  Not necessarily unique. internally: the CC customer full name, or for on-prem this can be something hardcoded.
	Name *string `json:"Name,omitempty"`
	// Sites which the admin has access to within the customer.
	Sites []MySiteResponseModel `json:"Sites"`
}

// NewMyCustomerResponseModel instantiates a new MyCustomerResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyCustomerResponseModel(id string, sites []MySiteResponseModel) *MyCustomerResponseModel {
	this := MyCustomerResponseModel{}
	this.Id = id
	this.Sites = sites
	return &this
}

// NewMyCustomerResponseModelWithDefaults instantiates a new MyCustomerResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyCustomerResponseModelWithDefaults() *MyCustomerResponseModel {
	this := MyCustomerResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *MyCustomerResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MyCustomerResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MyCustomerResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MyCustomerResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyCustomerResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MyCustomerResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MyCustomerResponseModel) SetName(v string) {
	o.Name = &v
}

// GetSites returns the Sites field value
func (o *MyCustomerResponseModel) GetSites() []MySiteResponseModel {
	if o == nil {
		var ret []MySiteResponseModel
		return ret
	}

	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value
// and a boolean to check if the value has been set.
func (o *MyCustomerResponseModel) GetSitesOk() ([]MySiteResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sites, true
}

// SetSites sets field value
func (o *MyCustomerResponseModel) SetSites(v []MySiteResponseModel) {
	o.Sites = v
}

func (o MyCustomerResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyCustomerResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	toSerialize["Sites"] = o.Sites
	return toSerialize, nil
}

type NullableMyCustomerResponseModel struct {
	value *MyCustomerResponseModel
	isSet bool
}

func (v NullableMyCustomerResponseModel) Get() *MyCustomerResponseModel {
	return v.value
}

func (v *NullableMyCustomerResponseModel) Set(val *MyCustomerResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMyCustomerResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMyCustomerResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyCustomerResponseModel(val *MyCustomerResponseModel) *NullableMyCustomerResponseModel {
	return &NullableMyCustomerResponseModel{value: val, isSet: true}
}

func (v NullableMyCustomerResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyCustomerResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


