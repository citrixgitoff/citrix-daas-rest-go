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

// checks if the HypervisorResourcePoolRefResponseModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolRefResponseModelAllOf{}

// HypervisorResourcePoolRefResponseModelAllOf Reference to a hypervisor resource pool, AKA \"HostingUnit\". To dereference and access the details about the resource pool, use GetHypervisorResourcePool with Hypervisor.Id and Id path parameters.
type HypervisorResourcePoolRefResponseModelAllOf struct {
	// Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: `Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources`
	FullRelativePath string `json:"FullRelativePath"`
	Hypervisor HypervisorResourcePoolRefResponseModelAllOfHypervisor `json:"Hypervisor"`
}

// NewHypervisorResourcePoolRefResponseModelAllOf instantiates a new HypervisorResourcePoolRefResponseModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolRefResponseModelAllOf(fullRelativePath string, hypervisor HypervisorResourcePoolRefResponseModelAllOfHypervisor) *HypervisorResourcePoolRefResponseModelAllOf {
	this := HypervisorResourcePoolRefResponseModelAllOf{}
	this.FullRelativePath = fullRelativePath
	this.Hypervisor = hypervisor
	return &this
}

// NewHypervisorResourcePoolRefResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolRefResponseModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolRefResponseModelAllOfWithDefaults() *HypervisorResourcePoolRefResponseModelAllOf {
	this := HypervisorResourcePoolRefResponseModelAllOf{}
	return &this
}

// GetFullRelativePath returns the FullRelativePath field value
func (o *HypervisorResourcePoolRefResponseModelAllOf) GetFullRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolRefResponseModelAllOf) GetFullRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullRelativePath, true
}

// SetFullRelativePath sets field value
func (o *HypervisorResourcePoolRefResponseModelAllOf) SetFullRelativePath(v string) {
	o.FullRelativePath = v
}

// GetHypervisor returns the Hypervisor field value
func (o *HypervisorResourcePoolRefResponseModelAllOf) GetHypervisor() HypervisorResourcePoolRefResponseModelAllOfHypervisor {
	if o == nil {
		var ret HypervisorResourcePoolRefResponseModelAllOfHypervisor
		return ret
	}

	return o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolRefResponseModelAllOf) GetHypervisorOk() (*HypervisorResourcePoolRefResponseModelAllOfHypervisor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hypervisor, true
}

// SetHypervisor sets field value
func (o *HypervisorResourcePoolRefResponseModelAllOf) SetHypervisor(v HypervisorResourcePoolRefResponseModelAllOfHypervisor) {
	o.Hypervisor = v
}

func (o HypervisorResourcePoolRefResponseModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolRefResponseModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FullRelativePath"] = o.FullRelativePath
	toSerialize["Hypervisor"] = o.Hypervisor
	return toSerialize, nil
}

type NullableHypervisorResourcePoolRefResponseModelAllOf struct {
	value *HypervisorResourcePoolRefResponseModelAllOf
	isSet bool
}

func (v NullableHypervisorResourcePoolRefResponseModelAllOf) Get() *HypervisorResourcePoolRefResponseModelAllOf {
	return v.value
}

func (v *NullableHypervisorResourcePoolRefResponseModelAllOf) Set(val *HypervisorResourcePoolRefResponseModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolRefResponseModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolRefResponseModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolRefResponseModelAllOf(val *HypervisorResourcePoolRefResponseModelAllOf) *NullableHypervisorResourcePoolRefResponseModelAllOf {
	return &NullableHypervisorResourcePoolRefResponseModelAllOf{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolRefResponseModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolRefResponseModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


