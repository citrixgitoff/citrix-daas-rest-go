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

// checks if the PolicySetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySetResponse{}

// PolicySetResponse GPO policy set data.
type PolicySetResponse struct {
	// Guid of the policy set.
	PolicySetGuid *string `json:"policySetGuid,omitempty"`
	PolicySetType *SdkGpoPolicySetType `json:"policySetType,omitempty"`
	// Name of the policy set.
	Name NullableString `json:"name,omitempty"`
	// Policy set description.
	Description NullableString `json:"description,omitempty"`
	// Number of policies in the policy set.
	PolicyCount *int32 `json:"policyCount,omitempty"`
	// Policies in the policy set.
	Policies []PolicyResponse `json:"policies,omitempty"`
	// Delegated admin scopes.
	Scopes []string `json:"scopes,omitempty"`
}

// NewPolicySetResponse instantiates a new PolicySetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySetResponse() *PolicySetResponse {
	this := PolicySetResponse{}
	return &this
}

// NewPolicySetResponseWithDefaults instantiates a new PolicySetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySetResponseWithDefaults() *PolicySetResponse {
	this := PolicySetResponse{}
	return &this
}

// GetPolicySetGuid returns the PolicySetGuid field value if set, zero value otherwise.
func (o *PolicySetResponse) GetPolicySetGuid() string {
	if o == nil || IsNil(o.PolicySetGuid) {
		var ret string
		return ret
	}
	return *o.PolicySetGuid
}

// GetPolicySetGuidOk returns a tuple with the PolicySetGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySetResponse) GetPolicySetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.PolicySetGuid) {
		return nil, false
	}
	return o.PolicySetGuid, true
}

// HasPolicySetGuid returns a boolean if a field has been set.
func (o *PolicySetResponse) HasPolicySetGuid() bool {
	if o != nil && !IsNil(o.PolicySetGuid) {
		return true
	}

	return false
}

// SetPolicySetGuid gets a reference to the given string and assigns it to the PolicySetGuid field.
func (o *PolicySetResponse) SetPolicySetGuid(v string) {
	o.PolicySetGuid = &v
}

// GetPolicySetType returns the PolicySetType field value if set, zero value otherwise.
func (o *PolicySetResponse) GetPolicySetType() SdkGpoPolicySetType {
	if o == nil || IsNil(o.PolicySetType) {
		var ret SdkGpoPolicySetType
		return ret
	}
	return *o.PolicySetType
}

// GetPolicySetTypeOk returns a tuple with the PolicySetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySetResponse) GetPolicySetTypeOk() (*SdkGpoPolicySetType, bool) {
	if o == nil || IsNil(o.PolicySetType) {
		return nil, false
	}
	return o.PolicySetType, true
}

// HasPolicySetType returns a boolean if a field has been set.
func (o *PolicySetResponse) HasPolicySetType() bool {
	if o != nil && !IsNil(o.PolicySetType) {
		return true
	}

	return false
}

// SetPolicySetType gets a reference to the given SdkGpoPolicySetType and assigns it to the PolicySetType field.
func (o *PolicySetResponse) SetPolicySetType(v SdkGpoPolicySetType) {
	o.PolicySetType = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PolicySetResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PolicySetResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PolicySetResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PolicySetResponse) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicySetResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PolicySetResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PolicySetResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PolicySetResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetPolicyCount returns the PolicyCount field value if set, zero value otherwise.
func (o *PolicySetResponse) GetPolicyCount() int32 {
	if o == nil || IsNil(o.PolicyCount) {
		var ret int32
		return ret
	}
	return *o.PolicyCount
}

// GetPolicyCountOk returns a tuple with the PolicyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySetResponse) GetPolicyCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyCount) {
		return nil, false
	}
	return o.PolicyCount, true
}

// HasPolicyCount returns a boolean if a field has been set.
func (o *PolicySetResponse) HasPolicyCount() bool {
	if o != nil && !IsNil(o.PolicyCount) {
		return true
	}

	return false
}

// SetPolicyCount gets a reference to the given int32 and assigns it to the PolicyCount field.
func (o *PolicySetResponse) SetPolicyCount(v int32) {
	o.PolicyCount = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetResponse) GetPolicies() []PolicyResponse {
	if o == nil {
		var ret []PolicyResponse
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetResponse) GetPoliciesOk() ([]PolicyResponse, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *PolicySetResponse) HasPolicies() bool {
	if o != nil && IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []PolicyResponse and assigns it to the Policies field.
func (o *PolicySetResponse) SetPolicies(v []PolicyResponse) {
	o.Policies = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PolicySetResponse) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *PolicySetResponse) SetScopes(v []string) {
	o.Scopes = v
}

func (o PolicySetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicySetGuid) {
		toSerialize["policySetGuid"] = o.PolicySetGuid
	}
	if !IsNil(o.PolicySetType) {
		toSerialize["policySetType"] = o.PolicySetType
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.PolicyCount) {
		toSerialize["policyCount"] = o.PolicyCount
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullablePolicySetResponse struct {
	value *PolicySetResponse
	isSet bool
}

func (v NullablePolicySetResponse) Get() *PolicySetResponse {
	return v.value
}

func (v *NullablePolicySetResponse) Set(val *PolicySetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySetResponse(val *PolicySetResponse) *NullablePolicySetResponse {
	return &NullablePolicySetResponse{value: val, isSet: true}
}

func (v NullablePolicySetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


