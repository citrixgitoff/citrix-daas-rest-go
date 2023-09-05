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

// checks if the MachineAccountCreationRulesRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineAccountCreationRulesRequestModel{}

// MachineAccountCreationRulesRequestModel Rules for creation of machine accounts in Active Directory.
type MachineAccountCreationRulesRequestModel struct {
	// Defines the template name for AD accounts created in the identity pool.  Required.
	NamingScheme string `json:"NamingScheme"`
	NamingSchemeType *NamingSchemeType `json:"NamingSchemeType,omitempty"`
	// The OU that computer accounts will be created into.  Optional.
	OU *string `json:"OU,omitempty"`
	// The AD domain name for the pool. Specify this in FQDN format; for example, MyDomain.com. Required.
	Domain string `json:"Domain"`
	// Defines the next value that will be used if creating new AD accounts.  Optional.
	NextValue *string `json:"NextValue,omitempty"`
	// Existing identity pool id
	IdentityPoolId *string `json:"IdentityPoolId,omitempty"`
}

// NewMachineAccountCreationRulesRequestModel instantiates a new MachineAccountCreationRulesRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineAccountCreationRulesRequestModel(namingScheme string, domain string) *MachineAccountCreationRulesRequestModel {
	this := MachineAccountCreationRulesRequestModel{}
	this.NamingScheme = namingScheme
	this.Domain = domain
	return &this
}

// NewMachineAccountCreationRulesRequestModelWithDefaults instantiates a new MachineAccountCreationRulesRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineAccountCreationRulesRequestModelWithDefaults() *MachineAccountCreationRulesRequestModel {
	this := MachineAccountCreationRulesRequestModel{}
	return &this
}

// GetNamingScheme returns the NamingScheme field value
func (o *MachineAccountCreationRulesRequestModel) GetNamingScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamingScheme
}

// GetNamingSchemeOk returns a tuple with the NamingScheme field value
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamingScheme, true
}

// SetNamingScheme sets field value
func (o *MachineAccountCreationRulesRequestModel) SetNamingScheme(v string) {
	o.NamingScheme = v
}

// GetNamingSchemeType returns the NamingSchemeType field value if set, zero value otherwise.
func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeType() NamingSchemeType {
	if o == nil || IsNil(o.NamingSchemeType) {
		var ret NamingSchemeType
		return ret
	}
	return *o.NamingSchemeType
}

// GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeTypeOk() (*NamingSchemeType, bool) {
	if o == nil || IsNil(o.NamingSchemeType) {
		return nil, false
	}
	return o.NamingSchemeType, true
}

// HasNamingSchemeType returns a boolean if a field has been set.
func (o *MachineAccountCreationRulesRequestModel) HasNamingSchemeType() bool {
	if o != nil && !IsNil(o.NamingSchemeType) {
		return true
	}

	return false
}

// SetNamingSchemeType gets a reference to the given NamingSchemeType and assigns it to the NamingSchemeType field.
func (o *MachineAccountCreationRulesRequestModel) SetNamingSchemeType(v NamingSchemeType) {
	o.NamingSchemeType = &v
}

// GetOU returns the OU field value if set, zero value otherwise.
func (o *MachineAccountCreationRulesRequestModel) GetOU() string {
	if o == nil || IsNil(o.OU) {
		var ret string
		return ret
	}
	return *o.OU
}

// GetOUOk returns a tuple with the OU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetOUOk() (*string, bool) {
	if o == nil || IsNil(o.OU) {
		return nil, false
	}
	return o.OU, true
}

// HasOU returns a boolean if a field has been set.
func (o *MachineAccountCreationRulesRequestModel) HasOU() bool {
	if o != nil && !IsNil(o.OU) {
		return true
	}

	return false
}

// SetOU gets a reference to the given string and assigns it to the OU field.
func (o *MachineAccountCreationRulesRequestModel) SetOU(v string) {
	o.OU = &v
}

// GetDomain returns the Domain field value
func (o *MachineAccountCreationRulesRequestModel) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *MachineAccountCreationRulesRequestModel) SetDomain(v string) {
	o.Domain = v
}

// GetNextValue returns the NextValue field value if set, zero value otherwise.
func (o *MachineAccountCreationRulesRequestModel) GetNextValue() string {
	if o == nil || IsNil(o.NextValue) {
		var ret string
		return ret
	}
	return *o.NextValue
}

// GetNextValueOk returns a tuple with the NextValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetNextValueOk() (*string, bool) {
	if o == nil || IsNil(o.NextValue) {
		return nil, false
	}
	return o.NextValue, true
}

// HasNextValue returns a boolean if a field has been set.
func (o *MachineAccountCreationRulesRequestModel) HasNextValue() bool {
	if o != nil && !IsNil(o.NextValue) {
		return true
	}

	return false
}

// SetNextValue gets a reference to the given string and assigns it to the NextValue field.
func (o *MachineAccountCreationRulesRequestModel) SetNextValue(v string) {
	o.NextValue = &v
}

// GetIdentityPoolId returns the IdentityPoolId field value if set, zero value otherwise.
func (o *MachineAccountCreationRulesRequestModel) GetIdentityPoolId() string {
	if o == nil || IsNil(o.IdentityPoolId) {
		var ret string
		return ret
	}
	return *o.IdentityPoolId
}

// GetIdentityPoolIdOk returns a tuple with the IdentityPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAccountCreationRulesRequestModel) GetIdentityPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityPoolId) {
		return nil, false
	}
	return o.IdentityPoolId, true
}

// HasIdentityPoolId returns a boolean if a field has been set.
func (o *MachineAccountCreationRulesRequestModel) HasIdentityPoolId() bool {
	if o != nil && !IsNil(o.IdentityPoolId) {
		return true
	}

	return false
}

// SetIdentityPoolId gets a reference to the given string and assigns it to the IdentityPoolId field.
func (o *MachineAccountCreationRulesRequestModel) SetIdentityPoolId(v string) {
	o.IdentityPoolId = &v
}

func (o MachineAccountCreationRulesRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineAccountCreationRulesRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["NamingScheme"] = o.NamingScheme
	if !IsNil(o.NamingSchemeType) {
		toSerialize["NamingSchemeType"] = o.NamingSchemeType
	}
	if !IsNil(o.OU) {
		toSerialize["OU"] = o.OU
	}
	toSerialize["Domain"] = o.Domain
	if !IsNil(o.NextValue) {
		toSerialize["NextValue"] = o.NextValue
	}
	if !IsNil(o.IdentityPoolId) {
		toSerialize["IdentityPoolId"] = o.IdentityPoolId
	}
	return toSerialize, nil
}

type NullableMachineAccountCreationRulesRequestModel struct {
	value *MachineAccountCreationRulesRequestModel
	isSet bool
}

func (v NullableMachineAccountCreationRulesRequestModel) Get() *MachineAccountCreationRulesRequestModel {
	return v.value
}

func (v *NullableMachineAccountCreationRulesRequestModel) Set(val *MachineAccountCreationRulesRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAccountCreationRulesRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAccountCreationRulesRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAccountCreationRulesRequestModel(val *MachineAccountCreationRulesRequestModel) *NullableMachineAccountCreationRulesRequestModel {
	return &NullableMachineAccountCreationRulesRequestModel{value: val, isSet: true}
}

func (v NullableMachineAccountCreationRulesRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAccountCreationRulesRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


