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

// checks if the PvsCollectionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PvsCollectionResponseModel{}

// PvsCollectionResponseModel PVS collection object.
type PvsCollectionResponseModel struct {
	// Id of the PVS collection.
	Id string `json:"Id"`
	// Name of the PVS collection.
	Name string `json:"Name"`
	// Active directory domain of machines in the collection. internally: when querying a collection, get a single device from the collection and get its AD domain.
	Domain NullableString `json:"Domain,omitempty"`
	PvsSite RefResponseModel `json:"PvsSite"`
}

// NewPvsCollectionResponseModel instantiates a new PvsCollectionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPvsCollectionResponseModel(id string, name string, pvsSite RefResponseModel) *PvsCollectionResponseModel {
	this := PvsCollectionResponseModel{}
	this.Id = id
	this.Name = name
	this.PvsSite = pvsSite
	return &this
}

// NewPvsCollectionResponseModelWithDefaults instantiates a new PvsCollectionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvsCollectionResponseModelWithDefaults() *PvsCollectionResponseModel {
	this := PvsCollectionResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *PvsCollectionResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PvsCollectionResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PvsCollectionResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PvsCollectionResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PvsCollectionResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PvsCollectionResponseModel) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PvsCollectionResponseModel) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PvsCollectionResponseModel) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *PvsCollectionResponseModel) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *PvsCollectionResponseModel) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *PvsCollectionResponseModel) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *PvsCollectionResponseModel) UnsetDomain() {
	o.Domain.Unset()
}

// GetPvsSite returns the PvsSite field value
func (o *PvsCollectionResponseModel) GetPvsSite() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.PvsSite
}

// GetPvsSiteOk returns a tuple with the PvsSite field value
// and a boolean to check if the value has been set.
func (o *PvsCollectionResponseModel) GetPvsSiteOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PvsSite, true
}

// SetPvsSite sets field value
func (o *PvsCollectionResponseModel) SetPvsSite(v RefResponseModel) {
	o.PvsSite = v
}

func (o PvsCollectionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PvsCollectionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	if o.Domain.IsSet() {
		toSerialize["Domain"] = o.Domain.Get()
	}
	toSerialize["PvsSite"] = o.PvsSite
	return toSerialize, nil
}

type NullablePvsCollectionResponseModel struct {
	value *PvsCollectionResponseModel
	isSet bool
}

func (v NullablePvsCollectionResponseModel) Get() *PvsCollectionResponseModel {
	return v.value
}

func (v *NullablePvsCollectionResponseModel) Set(val *PvsCollectionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePvsCollectionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePvsCollectionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePvsCollectionResponseModel(val *PvsCollectionResponseModel) *NullablePvsCollectionResponseModel {
	return &NullablePvsCollectionResponseModel{value: val, isSet: true}
}

func (v NullablePvsCollectionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePvsCollectionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


