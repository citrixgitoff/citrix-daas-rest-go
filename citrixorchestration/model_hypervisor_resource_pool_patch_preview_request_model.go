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

// checks if the HypervisorResourcePoolPatchPreviewRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolPatchPreviewRequestModel{}

// HypervisorResourcePoolPatchPreviewRequestModel struct for HypervisorResourcePoolPatchPreviewRequestModel
type HypervisorResourcePoolPatchPreviewRequestModel struct {
	// Path to the networks/subnets that are available for provisioning operations in this resource pool. At least one is required.
	Networks []string `json:"Networks,omitempty"`
}

// NewHypervisorResourcePoolPatchPreviewRequestModel instantiates a new HypervisorResourcePoolPatchPreviewRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolPatchPreviewRequestModel() *HypervisorResourcePoolPatchPreviewRequestModel {
	this := HypervisorResourcePoolPatchPreviewRequestModel{}
	return &this
}

// NewHypervisorResourcePoolPatchPreviewRequestModelWithDefaults instantiates a new HypervisorResourcePoolPatchPreviewRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolPatchPreviewRequestModelWithDefaults() *HypervisorResourcePoolPatchPreviewRequestModel {
	this := HypervisorResourcePoolPatchPreviewRequestModel{}
	return &this
}

// GetNetworks returns the Networks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *HypervisorResourcePoolPatchPreviewRequestModel) HasNetworks() bool {
	if o != nil && IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *HypervisorResourcePoolPatchPreviewRequestModel) SetNetworks(v []string) {
	o.Networks = v
}

func (o HypervisorResourcePoolPatchPreviewRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolPatchPreviewRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Networks != nil {
		toSerialize["Networks"] = o.Networks
	}
	return toSerialize, nil
}

type NullableHypervisorResourcePoolPatchPreviewRequestModel struct {
	value *HypervisorResourcePoolPatchPreviewRequestModel
	isSet bool
}

func (v NullableHypervisorResourcePoolPatchPreviewRequestModel) Get() *HypervisorResourcePoolPatchPreviewRequestModel {
	return v.value
}

func (v *NullableHypervisorResourcePoolPatchPreviewRequestModel) Set(val *HypervisorResourcePoolPatchPreviewRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolPatchPreviewRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolPatchPreviewRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolPatchPreviewRequestModel(val *HypervisorResourcePoolPatchPreviewRequestModel) *NullableHypervisorResourcePoolPatchPreviewRequestModel {
	return &NullableHypervisorResourcePoolPatchPreviewRequestModel{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolPatchPreviewRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolPatchPreviewRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


