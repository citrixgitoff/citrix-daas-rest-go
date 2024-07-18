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

// checks if the LosingPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LosingPolicy{}

// LosingPolicy Policy that should be applied but not applied due to various reasons.
type LosingPolicy struct {
	// Name of the policy.
	PolicyName NullableString `json:"PolicyName,omitempty"`
	// Name of the GPO that contains the policy that uses this setting.
	GpoName NullableString `json:"GpoName,omitempty"`
	// Reasons why the policy is not applied according to filtering results.
	Reasons map[string][]ReasonDetail `json:"Reasons,omitempty"`
	// Policy priority.
	Priority *int32 `json:"Priority,omitempty"`
	// The settings in the policy and the reason of losing for each setting.
	WinningSettings []WonOverBy `json:"WinningSettings,omitempty"`
}

// NewLosingPolicy instantiates a new LosingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLosingPolicy() *LosingPolicy {
	this := LosingPolicy{}
	return &this
}

// NewLosingPolicyWithDefaults instantiates a new LosingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLosingPolicyWithDefaults() *LosingPolicy {
	this := LosingPolicy{}
	return &this
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingPolicy) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingPolicy) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *LosingPolicy) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *LosingPolicy) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *LosingPolicy) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *LosingPolicy) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetGpoName returns the GpoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingPolicy) GetGpoName() string {
	if o == nil || IsNil(o.GpoName.Get()) {
		var ret string
		return ret
	}
	return *o.GpoName.Get()
}

// GetGpoNameOk returns a tuple with the GpoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingPolicy) GetGpoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpoName.Get(), o.GpoName.IsSet()
}

// HasGpoName returns a boolean if a field has been set.
func (o *LosingPolicy) HasGpoName() bool {
	if o != nil && o.GpoName.IsSet() {
		return true
	}

	return false
}

// SetGpoName gets a reference to the given NullableString and assigns it to the GpoName field.
func (o *LosingPolicy) SetGpoName(v string) {
	o.GpoName.Set(&v)
}
// SetGpoNameNil sets the value for GpoName to be an explicit nil
func (o *LosingPolicy) SetGpoNameNil() {
	o.GpoName.Set(nil)
}

// UnsetGpoName ensures that no value is present for GpoName, not even an explicit nil
func (o *LosingPolicy) UnsetGpoName() {
	o.GpoName.Unset()
}

// GetReasons returns the Reasons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingPolicy) GetReasons() map[string][]ReasonDetail {
	if o == nil {
		var ret map[string][]ReasonDetail
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingPolicy) GetReasonsOk() (*map[string][]ReasonDetail, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *LosingPolicy) HasReasons() bool {
	if o != nil && IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given map[string][]ReasonDetail and assigns it to the Reasons field.
func (o *LosingPolicy) SetReasons(v map[string][]ReasonDetail) {
	o.Reasons = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LosingPolicy) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LosingPolicy) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LosingPolicy) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *LosingPolicy) SetPriority(v int32) {
	o.Priority = &v
}

// GetWinningSettings returns the WinningSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingPolicy) GetWinningSettings() []WonOverBy {
	if o == nil {
		var ret []WonOverBy
		return ret
	}
	return o.WinningSettings
}

// GetWinningSettingsOk returns a tuple with the WinningSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingPolicy) GetWinningSettingsOk() ([]WonOverBy, bool) {
	if o == nil || IsNil(o.WinningSettings) {
		return nil, false
	}
	return o.WinningSettings, true
}

// HasWinningSettings returns a boolean if a field has been set.
func (o *LosingPolicy) HasWinningSettings() bool {
	if o != nil && IsNil(o.WinningSettings) {
		return true
	}

	return false
}

// SetWinningSettings gets a reference to the given []WonOverBy and assigns it to the WinningSettings field.
func (o *LosingPolicy) SetWinningSettings(v []WonOverBy) {
	o.WinningSettings = v
}

func (o LosingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LosingPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicyName.IsSet() {
		toSerialize["PolicyName"] = o.PolicyName.Get()
	}
	if o.GpoName.IsSet() {
		toSerialize["GpoName"] = o.GpoName.Get()
	}
	if o.Reasons != nil {
		toSerialize["Reasons"] = o.Reasons
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if o.WinningSettings != nil {
		toSerialize["WinningSettings"] = o.WinningSettings
	}
	return toSerialize, nil
}

type NullableLosingPolicy struct {
	value *LosingPolicy
	isSet bool
}

func (v NullableLosingPolicy) Get() *LosingPolicy {
	return v.value
}

func (v *NullableLosingPolicy) Set(val *LosingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableLosingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableLosingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLosingPolicy(val *LosingPolicy) *NullableLosingPolicy {
	return &NullableLosingPolicy{value: val, isSet: true}
}

func (v NullableLosingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLosingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


