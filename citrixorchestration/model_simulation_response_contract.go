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

// checks if the SimulationResponseContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimulationResponseContract{}

// SimulationResponseContract Result of a simulation run.
type SimulationResponseContract struct {
	// The report is for machine or user.
	IsUserRsop *bool `json:"IsUserRsop,omitempty"`
	// Filter evidence used for the simulation.
	FilterEvidence map[string]string `json:"FilterEvidence,omitempty"`
	// Applied settings.
	AppliedSettings []AppliedSetting `json:"AppliedSettings,omitempty"`
	// Applied policies.
	AppliedPolicies []AppliedPolicy `json:"AppliedPolicies,omitempty"`
	// Settings that did not get applied.
	LosingSettings []LosingSetting `json:"LosingSettings,omitempty"`
	// Policies that did not get applied. Some of them should be applied in theory but not applied in practice.
	LosingPolicies []LosingPolicy `json:"LosingPolicies,omitempty"`
}

// NewSimulationResponseContract instantiates a new SimulationResponseContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulationResponseContract() *SimulationResponseContract {
	this := SimulationResponseContract{}
	return &this
}

// NewSimulationResponseContractWithDefaults instantiates a new SimulationResponseContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulationResponseContractWithDefaults() *SimulationResponseContract {
	this := SimulationResponseContract{}
	return &this
}

// GetIsUserRsop returns the IsUserRsop field value if set, zero value otherwise.
func (o *SimulationResponseContract) GetIsUserRsop() bool {
	if o == nil || IsNil(o.IsUserRsop) {
		var ret bool
		return ret
	}
	return *o.IsUserRsop
}

// GetIsUserRsopOk returns a tuple with the IsUserRsop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulationResponseContract) GetIsUserRsopOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserRsop) {
		return nil, false
	}
	return o.IsUserRsop, true
}

// HasIsUserRsop returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasIsUserRsop() bool {
	if o != nil && !IsNil(o.IsUserRsop) {
		return true
	}

	return false
}

// SetIsUserRsop gets a reference to the given bool and assigns it to the IsUserRsop field.
func (o *SimulationResponseContract) SetIsUserRsop(v bool) {
	o.IsUserRsop = &v
}

// GetFilterEvidence returns the FilterEvidence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulationResponseContract) GetFilterEvidence() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.FilterEvidence
}

// GetFilterEvidenceOk returns a tuple with the FilterEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulationResponseContract) GetFilterEvidenceOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.FilterEvidence) {
		return nil, false
	}
	return &o.FilterEvidence, true
}

// HasFilterEvidence returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasFilterEvidence() bool {
	if o != nil && IsNil(o.FilterEvidence) {
		return true
	}

	return false
}

// SetFilterEvidence gets a reference to the given map[string]string and assigns it to the FilterEvidence field.
func (o *SimulationResponseContract) SetFilterEvidence(v map[string]string) {
	o.FilterEvidence = v
}

// GetAppliedSettings returns the AppliedSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulationResponseContract) GetAppliedSettings() []AppliedSetting {
	if o == nil {
		var ret []AppliedSetting
		return ret
	}
	return o.AppliedSettings
}

// GetAppliedSettingsOk returns a tuple with the AppliedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulationResponseContract) GetAppliedSettingsOk() ([]AppliedSetting, bool) {
	if o == nil || IsNil(o.AppliedSettings) {
		return nil, false
	}
	return o.AppliedSettings, true
}

// HasAppliedSettings returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasAppliedSettings() bool {
	if o != nil && IsNil(o.AppliedSettings) {
		return true
	}

	return false
}

// SetAppliedSettings gets a reference to the given []AppliedSetting and assigns it to the AppliedSettings field.
func (o *SimulationResponseContract) SetAppliedSettings(v []AppliedSetting) {
	o.AppliedSettings = v
}

// GetAppliedPolicies returns the AppliedPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulationResponseContract) GetAppliedPolicies() []AppliedPolicy {
	if o == nil {
		var ret []AppliedPolicy
		return ret
	}
	return o.AppliedPolicies
}

// GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulationResponseContract) GetAppliedPoliciesOk() ([]AppliedPolicy, bool) {
	if o == nil || IsNil(o.AppliedPolicies) {
		return nil, false
	}
	return o.AppliedPolicies, true
}

// HasAppliedPolicies returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasAppliedPolicies() bool {
	if o != nil && IsNil(o.AppliedPolicies) {
		return true
	}

	return false
}

// SetAppliedPolicies gets a reference to the given []AppliedPolicy and assigns it to the AppliedPolicies field.
func (o *SimulationResponseContract) SetAppliedPolicies(v []AppliedPolicy) {
	o.AppliedPolicies = v
}

// GetLosingSettings returns the LosingSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulationResponseContract) GetLosingSettings() []LosingSetting {
	if o == nil {
		var ret []LosingSetting
		return ret
	}
	return o.LosingSettings
}

// GetLosingSettingsOk returns a tuple with the LosingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulationResponseContract) GetLosingSettingsOk() ([]LosingSetting, bool) {
	if o == nil || IsNil(o.LosingSettings) {
		return nil, false
	}
	return o.LosingSettings, true
}

// HasLosingSettings returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasLosingSettings() bool {
	if o != nil && IsNil(o.LosingSettings) {
		return true
	}

	return false
}

// SetLosingSettings gets a reference to the given []LosingSetting and assigns it to the LosingSettings field.
func (o *SimulationResponseContract) SetLosingSettings(v []LosingSetting) {
	o.LosingSettings = v
}

// GetLosingPolicies returns the LosingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulationResponseContract) GetLosingPolicies() []LosingPolicy {
	if o == nil {
		var ret []LosingPolicy
		return ret
	}
	return o.LosingPolicies
}

// GetLosingPoliciesOk returns a tuple with the LosingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulationResponseContract) GetLosingPoliciesOk() ([]LosingPolicy, bool) {
	if o == nil || IsNil(o.LosingPolicies) {
		return nil, false
	}
	return o.LosingPolicies, true
}

// HasLosingPolicies returns a boolean if a field has been set.
func (o *SimulationResponseContract) HasLosingPolicies() bool {
	if o != nil && IsNil(o.LosingPolicies) {
		return true
	}

	return false
}

// SetLosingPolicies gets a reference to the given []LosingPolicy and assigns it to the LosingPolicies field.
func (o *SimulationResponseContract) SetLosingPolicies(v []LosingPolicy) {
	o.LosingPolicies = v
}

func (o SimulationResponseContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimulationResponseContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsUserRsop) {
		toSerialize["IsUserRsop"] = o.IsUserRsop
	}
	if o.FilterEvidence != nil {
		toSerialize["FilterEvidence"] = o.FilterEvidence
	}
	if o.AppliedSettings != nil {
		toSerialize["AppliedSettings"] = o.AppliedSettings
	}
	if o.AppliedPolicies != nil {
		toSerialize["AppliedPolicies"] = o.AppliedPolicies
	}
	if o.LosingSettings != nil {
		toSerialize["LosingSettings"] = o.LosingSettings
	}
	if o.LosingPolicies != nil {
		toSerialize["LosingPolicies"] = o.LosingPolicies
	}
	return toSerialize, nil
}

type NullableSimulationResponseContract struct {
	value *SimulationResponseContract
	isSet bool
}

func (v NullableSimulationResponseContract) Get() *SimulationResponseContract {
	return v.value
}

func (v *NullableSimulationResponseContract) Set(val *SimulationResponseContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulationResponseContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulationResponseContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulationResponseContract(val *SimulationResponseContract) *NullableSimulationResponseContract {
	return &NullableSimulationResponseContract{value: val, isSet: true}
}

func (v NullableSimulationResponseContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulationResponseContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


