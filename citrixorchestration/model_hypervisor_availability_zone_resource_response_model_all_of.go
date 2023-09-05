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

// checks if the HypervisorAvailabilityZoneResourceResponseModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorAvailabilityZoneResourceResponseModelAllOf{}

// HypervisorAvailabilityZoneResourceResponseModelAllOf struct for HypervisorAvailabilityZoneResourceResponseModelAllOf
type HypervisorAvailabilityZoneResourceResponseModelAllOf struct {
	// Indicates whether this zone supports the use of security groups for isolation.
	SupportsSecurityGroups bool `json:"SupportsSecurityGroups"`
}

// NewHypervisorAvailabilityZoneResourceResponseModelAllOf instantiates a new HypervisorAvailabilityZoneResourceResponseModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorAvailabilityZoneResourceResponseModelAllOf(supportsSecurityGroups bool) *HypervisorAvailabilityZoneResourceResponseModelAllOf {
	this := HypervisorAvailabilityZoneResourceResponseModelAllOf{}
	this.SupportsSecurityGroups = supportsSecurityGroups
	return &this
}

// NewHypervisorAvailabilityZoneResourceResponseModelAllOfWithDefaults instantiates a new HypervisorAvailabilityZoneResourceResponseModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorAvailabilityZoneResourceResponseModelAllOfWithDefaults() *HypervisorAvailabilityZoneResourceResponseModelAllOf {
	this := HypervisorAvailabilityZoneResourceResponseModelAllOf{}
	return &this
}

// GetSupportsSecurityGroups returns the SupportsSecurityGroups field value
func (o *HypervisorAvailabilityZoneResourceResponseModelAllOf) GetSupportsSecurityGroups() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsSecurityGroups
}

// GetSupportsSecurityGroupsOk returns a tuple with the SupportsSecurityGroups field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModelAllOf) GetSupportsSecurityGroupsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportsSecurityGroups, true
}

// SetSupportsSecurityGroups sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModelAllOf) SetSupportsSecurityGroups(v bool) {
	o.SupportsSecurityGroups = v
}

func (o HypervisorAvailabilityZoneResourceResponseModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorAvailabilityZoneResourceResponseModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SupportsSecurityGroups"] = o.SupportsSecurityGroups
	return toSerialize, nil
}

type NullableHypervisorAvailabilityZoneResourceResponseModelAllOf struct {
	value *HypervisorAvailabilityZoneResourceResponseModelAllOf
	isSet bool
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) Get() *HypervisorAvailabilityZoneResourceResponseModelAllOf {
	return v.value
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) Set(val *HypervisorAvailabilityZoneResourceResponseModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorAvailabilityZoneResourceResponseModelAllOf(val *HypervisorAvailabilityZoneResourceResponseModelAllOf) *NullableHypervisorAvailabilityZoneResourceResponseModelAllOf {
	return &NullableHypervisorAvailabilityZoneResourceResponseModelAllOf{value: val, isSet: true}
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


