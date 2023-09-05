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

// checks if the EditHypervisorWOLConnectionRequestModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHypervisorWOLConnectionRequestModelAllOf{}

// EditHypervisorWOLConnectionRequestModelAllOf struct for EditHypervisorWOLConnectionRequestModelAllOf
type EditHypervisorWOLConnectionRequestModelAllOf struct {
	// Specifies whether to use Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy for power management.  Optional.  If not specified, will not be changed.
	SccmWakeUpProxy *bool `json:"SccmWakeUpProxy,omitempty"`
	WakeOnLanPackets *WakeOnLanTransmission `json:"WakeOnLanPackets,omitempty"`
}

// NewEditHypervisorWOLConnectionRequestModelAllOf instantiates a new EditHypervisorWOLConnectionRequestModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHypervisorWOLConnectionRequestModelAllOf() *EditHypervisorWOLConnectionRequestModelAllOf {
	this := EditHypervisorWOLConnectionRequestModelAllOf{}
	return &this
}

// NewEditHypervisorWOLConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorWOLConnectionRequestModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHypervisorWOLConnectionRequestModelAllOfWithDefaults() *EditHypervisorWOLConnectionRequestModelAllOf {
	this := EditHypervisorWOLConnectionRequestModelAllOf{}
	return &this
}

// GetSccmWakeUpProxy returns the SccmWakeUpProxy field value if set, zero value otherwise.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetSccmWakeUpProxy() bool {
	if o == nil || IsNil(o.SccmWakeUpProxy) {
		var ret bool
		return ret
	}
	return *o.SccmWakeUpProxy
}

// GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetSccmWakeUpProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.SccmWakeUpProxy) {
		return nil, false
	}
	return o.SccmWakeUpProxy, true
}

// HasSccmWakeUpProxy returns a boolean if a field has been set.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) HasSccmWakeUpProxy() bool {
	if o != nil && !IsNil(o.SccmWakeUpProxy) {
		return true
	}

	return false
}

// SetSccmWakeUpProxy gets a reference to the given bool and assigns it to the SccmWakeUpProxy field.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) SetSccmWakeUpProxy(v bool) {
	o.SccmWakeUpProxy = &v
}

// GetWakeOnLanPackets returns the WakeOnLanPackets field value if set, zero value otherwise.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetWakeOnLanPackets() WakeOnLanTransmission {
	if o == nil || IsNil(o.WakeOnLanPackets) {
		var ret WakeOnLanTransmission
		return ret
	}
	return *o.WakeOnLanPackets
}

// GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool) {
	if o == nil || IsNil(o.WakeOnLanPackets) {
		return nil, false
	}
	return o.WakeOnLanPackets, true
}

// HasWakeOnLanPackets returns a boolean if a field has been set.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) HasWakeOnLanPackets() bool {
	if o != nil && !IsNil(o.WakeOnLanPackets) {
		return true
	}

	return false
}

// SetWakeOnLanPackets gets a reference to the given WakeOnLanTransmission and assigns it to the WakeOnLanPackets field.
func (o *EditHypervisorWOLConnectionRequestModelAllOf) SetWakeOnLanPackets(v WakeOnLanTransmission) {
	o.WakeOnLanPackets = &v
}

func (o EditHypervisorWOLConnectionRequestModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHypervisorWOLConnectionRequestModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SccmWakeUpProxy) {
		toSerialize["SccmWakeUpProxy"] = o.SccmWakeUpProxy
	}
	if !IsNil(o.WakeOnLanPackets) {
		toSerialize["WakeOnLanPackets"] = o.WakeOnLanPackets
	}
	return toSerialize, nil
}

type NullableEditHypervisorWOLConnectionRequestModelAllOf struct {
	value *EditHypervisorWOLConnectionRequestModelAllOf
	isSet bool
}

func (v NullableEditHypervisorWOLConnectionRequestModelAllOf) Get() *EditHypervisorWOLConnectionRequestModelAllOf {
	return v.value
}

func (v *NullableEditHypervisorWOLConnectionRequestModelAllOf) Set(val *EditHypervisorWOLConnectionRequestModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHypervisorWOLConnectionRequestModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHypervisorWOLConnectionRequestModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHypervisorWOLConnectionRequestModelAllOf(val *EditHypervisorWOLConnectionRequestModelAllOf) *NullableEditHypervisorWOLConnectionRequestModelAllOf {
	return &NullableEditHypervisorWOLConnectionRequestModelAllOf{value: val, isSet: true}
}

func (v NullableEditHypervisorWOLConnectionRequestModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHypervisorWOLConnectionRequestModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


