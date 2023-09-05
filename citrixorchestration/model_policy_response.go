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

// checks if the PolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyResponse{}

// PolicyResponse Policy model
type PolicyResponse struct {
	// The blob that this policy belongs to.
	PolicySetGuid string `json:"policySetGuid"`
	// Policy GUID.
	PolicyGuid string `json:"policyGuid"`
	// Policy name, translated.
	PolicyName *string `json:"policyName,omitempty"`
	// Policy priority
	Priority int32 `json:"priority"`
	// Is policy enabled
	IsEnabled bool `json:"isEnabled"`
	// Policy description
	Description *string `json:"description,omitempty"`
	// Policy settings
	Settings []SettingResponse `json:"settings,omitempty"`
	// Policy filters
	Filters []FilterResponse `json:"filters,omitempty"`
}

// NewPolicyResponse instantiates a new PolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyResponse(policySetGuid string, policyGuid string, priority int32, isEnabled bool) *PolicyResponse {
	this := PolicyResponse{}
	this.PolicySetGuid = policySetGuid
	this.PolicyGuid = policyGuid
	this.Priority = priority
	this.IsEnabled = isEnabled
	return &this
}

// NewPolicyResponseWithDefaults instantiates a new PolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyResponseWithDefaults() *PolicyResponse {
	this := PolicyResponse{}
	return &this
}

// GetPolicySetGuid returns the PolicySetGuid field value
func (o *PolicyResponse) GetPolicySetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicySetGuid
}

// GetPolicySetGuidOk returns a tuple with the PolicySetGuid field value
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetPolicySetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicySetGuid, true
}

// SetPolicySetGuid sets field value
func (o *PolicyResponse) SetPolicySetGuid(v string) {
	o.PolicySetGuid = v
}

// GetPolicyGuid returns the PolicyGuid field value
func (o *PolicyResponse) GetPolicyGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyGuid
}

// GetPolicyGuidOk returns a tuple with the PolicyGuid field value
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetPolicyGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyGuid, true
}

// SetPolicyGuid sets field value
func (o *PolicyResponse) SetPolicyGuid(v string) {
	o.PolicyGuid = v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *PolicyResponse) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PolicyResponse) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *PolicyResponse) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPriority returns the Priority field value
func (o *PolicyResponse) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *PolicyResponse) SetPriority(v int32) {
	o.Priority = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *PolicyResponse) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *PolicyResponse) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *PolicyResponse) GetSettings() []SettingResponse {
	if o == nil || IsNil(o.Settings) {
		var ret []SettingResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetSettingsOk() ([]SettingResponse, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PolicyResponse) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []SettingResponse and assigns it to the Settings field.
func (o *PolicyResponse) SetSettings(v []SettingResponse) {
	o.Settings = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PolicyResponse) GetFilters() []FilterResponse {
	if o == nil || IsNil(o.Filters) {
		var ret []FilterResponse
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResponse) GetFiltersOk() ([]FilterResponse, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PolicyResponse) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []FilterResponse and assigns it to the Filters field.
func (o *PolicyResponse) SetFilters(v []FilterResponse) {
	o.Filters = v
}

func (o PolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policySetGuid"] = o.PolicySetGuid
	toSerialize["policyGuid"] = o.PolicyGuid
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	toSerialize["priority"] = o.Priority
	toSerialize["isEnabled"] = o.IsEnabled
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullablePolicyResponse struct {
	value *PolicyResponse
	isSet bool
}

func (v NullablePolicyResponse) Get() *PolicyResponse {
	return v.value
}

func (v *NullablePolicyResponse) Set(val *PolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyResponse(val *PolicyResponse) *NullablePolicyResponse {
	return &NullablePolicyResponse{value: val, isSet: true}
}

func (v NullablePolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


