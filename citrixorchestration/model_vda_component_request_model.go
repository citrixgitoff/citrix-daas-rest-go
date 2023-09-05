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

// checks if the VDAComponentRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDAComponentRequestModel{}

// VDAComponentRequestModel struct for VDAComponentRequestModel
type VDAComponentRequestModel struct {
	// Id of the component.
	ComponentId string `json:"ComponentId"`
	Parameters []VDAComponentParameterRequestModel `json:"Parameters,omitempty"`
}

// NewVDAComponentRequestModel instantiates a new VDAComponentRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDAComponentRequestModel(componentId string) *VDAComponentRequestModel {
	this := VDAComponentRequestModel{}
	this.ComponentId = componentId
	return &this
}

// NewVDAComponentRequestModelWithDefaults instantiates a new VDAComponentRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDAComponentRequestModelWithDefaults() *VDAComponentRequestModel {
	this := VDAComponentRequestModel{}
	return &this
}

// GetComponentId returns the ComponentId field value
func (o *VDAComponentRequestModel) GetComponentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *VDAComponentRequestModel) GetComponentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *VDAComponentRequestModel) SetComponentId(v string) {
	o.ComponentId = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *VDAComponentRequestModel) GetParameters() []VDAComponentParameterRequestModel {
	if o == nil || IsNil(o.Parameters) {
		var ret []VDAComponentParameterRequestModel
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDAComponentRequestModel) GetParametersOk() ([]VDAComponentParameterRequestModel, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *VDAComponentRequestModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []VDAComponentParameterRequestModel and assigns it to the Parameters field.
func (o *VDAComponentRequestModel) SetParameters(v []VDAComponentParameterRequestModel) {
	o.Parameters = v
}

func (o VDAComponentRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDAComponentRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ComponentId"] = o.ComponentId
	if !IsNil(o.Parameters) {
		toSerialize["Parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableVDAComponentRequestModel struct {
	value *VDAComponentRequestModel
	isSet bool
}

func (v NullableVDAComponentRequestModel) Get() *VDAComponentRequestModel {
	return v.value
}

func (v *NullableVDAComponentRequestModel) Set(val *VDAComponentRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVDAComponentRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVDAComponentRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDAComponentRequestModel(val *VDAComponentRequestModel) *NullableVDAComponentRequestModel {
	return &NullableVDAComponentRequestModel{value: val, isSet: true}
}

func (v NullableVDAComponentRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDAComponentRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


