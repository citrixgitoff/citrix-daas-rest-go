/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the AllResponseContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllResponseContract{}

// AllResponseContract ALL: Policy, Template, Setting Definitions and Filter Definitions
type AllResponseContract struct {
	// Policies
	Policies []PolicyResponseContract `json:"Policies,omitempty"`
	// Templates
	Templates []TemplateResponseContract `json:"Templates,omitempty"`
	// Setting Definitions
	SettingDefinitions []SettingDefinitionContract `json:"SettingDefinitions,omitempty"`
	// Filter Definitions
	FilterDefinitions []FilterDefinitionContract `json:"FilterDefinitions,omitempty"`
}

// NewAllResponseContract instantiates a new AllResponseContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllResponseContract() *AllResponseContract {
	this := AllResponseContract{}
	return &this
}

// NewAllResponseContractWithDefaults instantiates a new AllResponseContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllResponseContractWithDefaults() *AllResponseContract {
	this := AllResponseContract{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllResponseContract) GetPolicies() []PolicyResponseContract {
	if o == nil {
		var ret []PolicyResponseContract
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllResponseContract) GetPoliciesOk() ([]PolicyResponseContract, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AllResponseContract) HasPolicies() bool {
	if o != nil && IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []PolicyResponseContract and assigns it to the Policies field.
func (o *AllResponseContract) SetPolicies(v []PolicyResponseContract) {
	o.Policies = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllResponseContract) GetTemplates() []TemplateResponseContract {
	if o == nil {
		var ret []TemplateResponseContract
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllResponseContract) GetTemplatesOk() ([]TemplateResponseContract, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *AllResponseContract) HasTemplates() bool {
	if o != nil && IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []TemplateResponseContract and assigns it to the Templates field.
func (o *AllResponseContract) SetTemplates(v []TemplateResponseContract) {
	o.Templates = v
}

// GetSettingDefinitions returns the SettingDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllResponseContract) GetSettingDefinitions() []SettingDefinitionContract {
	if o == nil {
		var ret []SettingDefinitionContract
		return ret
	}
	return o.SettingDefinitions
}

// GetSettingDefinitionsOk returns a tuple with the SettingDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllResponseContract) GetSettingDefinitionsOk() ([]SettingDefinitionContract, bool) {
	if o == nil || IsNil(o.SettingDefinitions) {
		return nil, false
	}
	return o.SettingDefinitions, true
}

// HasSettingDefinitions returns a boolean if a field has been set.
func (o *AllResponseContract) HasSettingDefinitions() bool {
	if o != nil && IsNil(o.SettingDefinitions) {
		return true
	}

	return false
}

// SetSettingDefinitions gets a reference to the given []SettingDefinitionContract and assigns it to the SettingDefinitions field.
func (o *AllResponseContract) SetSettingDefinitions(v []SettingDefinitionContract) {
	o.SettingDefinitions = v
}

// GetFilterDefinitions returns the FilterDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllResponseContract) GetFilterDefinitions() []FilterDefinitionContract {
	if o == nil {
		var ret []FilterDefinitionContract
		return ret
	}
	return o.FilterDefinitions
}

// GetFilterDefinitionsOk returns a tuple with the FilterDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllResponseContract) GetFilterDefinitionsOk() ([]FilterDefinitionContract, bool) {
	if o == nil || IsNil(o.FilterDefinitions) {
		return nil, false
	}
	return o.FilterDefinitions, true
}

// HasFilterDefinitions returns a boolean if a field has been set.
func (o *AllResponseContract) HasFilterDefinitions() bool {
	if o != nil && IsNil(o.FilterDefinitions) {
		return true
	}

	return false
}

// SetFilterDefinitions gets a reference to the given []FilterDefinitionContract and assigns it to the FilterDefinitions field.
func (o *AllResponseContract) SetFilterDefinitions(v []FilterDefinitionContract) {
	o.FilterDefinitions = v
}

func (o AllResponseContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllResponseContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Policies != nil {
		toSerialize["Policies"] = o.Policies
	}
	if o.Templates != nil {
		toSerialize["Templates"] = o.Templates
	}
	if o.SettingDefinitions != nil {
		toSerialize["SettingDefinitions"] = o.SettingDefinitions
	}
	if o.FilterDefinitions != nil {
		toSerialize["FilterDefinitions"] = o.FilterDefinitions
	}
	return toSerialize, nil
}

type NullableAllResponseContract struct {
	value *AllResponseContract
	isSet bool
}

func (v NullableAllResponseContract) Get() *AllResponseContract {
	return v.value
}

func (v *NullableAllResponseContract) Set(val *AllResponseContract) {
	v.value = val
	v.isSet = true
}

func (v NullableAllResponseContract) IsSet() bool {
	return v.isSet
}

func (v *NullableAllResponseContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllResponseContract(val *AllResponseContract) *NullableAllResponseContract {
	return &NullableAllResponseContract{value: val, isSet: true}
}

func (v NullableAllResponseContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllResponseContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


