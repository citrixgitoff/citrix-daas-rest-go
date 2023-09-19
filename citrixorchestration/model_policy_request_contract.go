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

// checks if the PolicyRequestContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRequestContract{}

// PolicyRequestContract Policy request model
type PolicyRequestContract struct {
	// Policy name as identifier, translated.
	PolicyName NullableString `json:"PolicyName,omitempty"`
	// Is policy enabled
	IsEnabled NullableBool `json:"IsEnabled,omitempty"`
	// Policy description
	Description NullableString `json:"Description,omitempty"`
	// New policy priority
	Priority NullableInt32 `json:"Priority,omitempty"`
	// Policy settings
	Settings []SettingRequestContract `json:"Settings,omitempty"`
	// Policy filters
	Filters []FilterRequestContract `json:"Filters,omitempty"`
}

// NewPolicyRequestContract instantiates a new PolicyRequestContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRequestContract() *PolicyRequestContract {
	this := PolicyRequestContract{}
	var isEnabled bool = false
	this.IsEnabled = *NewNullableBool(&isEnabled)
	return &this
}

// NewPolicyRequestContractWithDefaults instantiates a new PolicyRequestContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRequestContractWithDefaults() *PolicyRequestContract {
	this := PolicyRequestContract{}
	var isEnabled bool = false
	this.IsEnabled = *NewNullableBool(&isEnabled)
	return &this
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *PolicyRequestContract) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *PolicyRequestContract) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *PolicyRequestContract) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *PolicyRequestContract) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *PolicyRequestContract) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *PolicyRequestContract) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PolicyRequestContract) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PolicyRequestContract) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PolicyRequestContract) UnsetDescription() {
	o.Description.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *PolicyRequestContract) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *PolicyRequestContract) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *PolicyRequestContract) UnsetPriority() {
	o.Priority.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetSettings() []SettingRequestContract {
	if o == nil {
		var ret []SettingRequestContract
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetSettingsOk() ([]SettingRequestContract, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasSettings() bool {
	if o != nil && IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []SettingRequestContract and assigns it to the Settings field.
func (o *PolicyRequestContract) SetSettings(v []SettingRequestContract) {
	o.Settings = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequestContract) GetFilters() []FilterRequestContract {
	if o == nil {
		var ret []FilterRequestContract
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequestContract) GetFiltersOk() ([]FilterRequestContract, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PolicyRequestContract) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []FilterRequestContract and assigns it to the Filters field.
func (o *PolicyRequestContract) SetFilters(v []FilterRequestContract) {
	o.Filters = v
}

func (o PolicyRequestContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRequestContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicyName.IsSet() {
		toSerialize["PolicyName"] = o.PolicyName.Get()
	}
	if o.IsEnabled.IsSet() {
		toSerialize["IsEnabled"] = o.IsEnabled.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["Priority"] = o.Priority.Get()
	}
	if o.Settings != nil {
		toSerialize["Settings"] = o.Settings
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullablePolicyRequestContract struct {
	value *PolicyRequestContract
	isSet bool
}

func (v NullablePolicyRequestContract) Get() *PolicyRequestContract {
	return v.value
}

func (v *NullablePolicyRequestContract) Set(val *PolicyRequestContract) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRequestContract) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRequestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRequestContract(val *PolicyRequestContract) *NullablePolicyRequestContract {
	return &NullablePolicyRequestContract{value: val, isSet: true}
}

func (v NullablePolicyRequestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRequestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


