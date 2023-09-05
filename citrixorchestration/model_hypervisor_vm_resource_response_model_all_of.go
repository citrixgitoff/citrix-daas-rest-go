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

// checks if the HypervisorVmResourceResponseModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorVmResourceResponseModelAllOf{}

// HypervisorVmResourceResponseModelAllOf struct for HypervisorVmResourceResponseModelAllOf
type HypervisorVmResourceResponseModelAllOf struct {
	// Id of the VM, as defined by the hypervisor.
	VMId string `json:"VMId"`
	// MAC address of the VM.
	MacAddress *string `json:"MacAddress,omitempty"`
}

// NewHypervisorVmResourceResponseModelAllOf instantiates a new HypervisorVmResourceResponseModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorVmResourceResponseModelAllOf(vMId string) *HypervisorVmResourceResponseModelAllOf {
	this := HypervisorVmResourceResponseModelAllOf{}
	this.VMId = vMId
	return &this
}

// NewHypervisorVmResourceResponseModelAllOfWithDefaults instantiates a new HypervisorVmResourceResponseModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorVmResourceResponseModelAllOfWithDefaults() *HypervisorVmResourceResponseModelAllOf {
	this := HypervisorVmResourceResponseModelAllOf{}
	return &this
}

// GetVMId returns the VMId field value
func (o *HypervisorVmResourceResponseModelAllOf) GetVMId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VMId
}

// GetVMIdOk returns a tuple with the VMId field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmResourceResponseModelAllOf) GetVMIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VMId, true
}

// SetVMId sets field value
func (o *HypervisorVmResourceResponseModelAllOf) SetVMId(v string) {
	o.VMId = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *HypervisorVmResourceResponseModelAllOf) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorVmResourceResponseModelAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *HypervisorVmResourceResponseModelAllOf) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *HypervisorVmResourceResponseModelAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o HypervisorVmResourceResponseModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorVmResourceResponseModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["VMId"] = o.VMId
	if !IsNil(o.MacAddress) {
		toSerialize["MacAddress"] = o.MacAddress
	}
	return toSerialize, nil
}

type NullableHypervisorVmResourceResponseModelAllOf struct {
	value *HypervisorVmResourceResponseModelAllOf
	isSet bool
}

func (v NullableHypervisorVmResourceResponseModelAllOf) Get() *HypervisorVmResourceResponseModelAllOf {
	return v.value
}

func (v *NullableHypervisorVmResourceResponseModelAllOf) Set(val *HypervisorVmResourceResponseModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorVmResourceResponseModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorVmResourceResponseModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorVmResourceResponseModelAllOf(val *HypervisorVmResourceResponseModelAllOf) *NullableHypervisorVmResourceResponseModelAllOf {
	return &NullableHypervisorVmResourceResponseModelAllOf{value: val, isSet: true}
}

func (v NullableHypervisorVmResourceResponseModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorVmResourceResponseModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


