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

// checks if the CreateHypervisorResourcePoolAzureRequestModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHypervisorResourcePoolAzureRequestModelAllOf{}

// CreateHypervisorResourcePoolAzureRequestModelAllOf struct for CreateHypervisorResourcePoolAzureRequestModelAllOf
type CreateHypervisorResourcePoolAzureRequestModelAllOf struct {
	// Azure region which the resource pool is connected to.  Required.
	Region string `json:"Region"`
	// Azure virtual network which the resource pool is connected to. Required.
	VirtualNetwork string `json:"VirtualNetwork"`
	// Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Subnets []string `json:"Subnets"`
}

// NewCreateHypervisorResourcePoolAzureRequestModelAllOf instantiates a new CreateHypervisorResourcePoolAzureRequestModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHypervisorResourcePoolAzureRequestModelAllOf(region string, virtualNetwork string, subnets []string) *CreateHypervisorResourcePoolAzureRequestModelAllOf {
	this := CreateHypervisorResourcePoolAzureRequestModelAllOf{}
	this.Region = region
	this.VirtualNetwork = virtualNetwork
	this.Subnets = subnets
	return &this
}

// NewCreateHypervisorResourcePoolAzureRequestModelAllOfWithDefaults instantiates a new CreateHypervisorResourcePoolAzureRequestModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHypervisorResourcePoolAzureRequestModelAllOfWithDefaults() *CreateHypervisorResourcePoolAzureRequestModelAllOf {
	this := CreateHypervisorResourcePoolAzureRequestModelAllOf{}
	return &this
}

// GetRegion returns the Region field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetRegion(v string) {
	o.Region = v
}

// GetVirtualNetwork returns the VirtualNetwork field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetVirtualNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetVirtualNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetwork, true
}

// SetVirtualNetwork sets field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetVirtualNetwork(v string) {
	o.VirtualNetwork = v
}

// GetSubnets returns the Subnets field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetSubnetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnets, true
}

// SetSubnets sets field value
func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetSubnets(v []string) {
	o.Subnets = v
}

func (o CreateHypervisorResourcePoolAzureRequestModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHypervisorResourcePoolAzureRequestModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Region"] = o.Region
	toSerialize["VirtualNetwork"] = o.VirtualNetwork
	toSerialize["Subnets"] = o.Subnets
	return toSerialize, nil
}

type NullableCreateHypervisorResourcePoolAzureRequestModelAllOf struct {
	value *CreateHypervisorResourcePoolAzureRequestModelAllOf
	isSet bool
}

func (v NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) Get() *CreateHypervisorResourcePoolAzureRequestModelAllOf {
	return v.value
}

func (v *NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) Set(val *CreateHypervisorResourcePoolAzureRequestModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHypervisorResourcePoolAzureRequestModelAllOf(val *CreateHypervisorResourcePoolAzureRequestModelAllOf) *NullableCreateHypervisorResourcePoolAzureRequestModelAllOf {
	return &NullableCreateHypervisorResourcePoolAzureRequestModelAllOf{value: val, isSet: true}
}

func (v NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHypervisorResourcePoolAzureRequestModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


