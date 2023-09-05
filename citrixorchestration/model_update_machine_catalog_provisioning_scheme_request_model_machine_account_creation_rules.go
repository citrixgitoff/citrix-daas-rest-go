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

// checks if the UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules{}

// UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules `DEPRECATED.` If specified, the machine account creation rules will be updated. If not specified, the machine account creation rules will be left unchanged.
type UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules struct {
	// Defines the template name for AD accounts created in the identity pool.   If this is not specified, it will remain unchanged.   If the provisioning scheme is configured with a NamingScheme already, and this value is set to an empty string, the provisioning scheme will be reconfigured so that it will no longer automatically create machine accounts.
	NamingScheme *string `json:"NamingScheme,omitempty"`
	NamingSchemeType *NamingSchemeType `json:"NamingSchemeType,omitempty"`
	// The OU that computer accounts will be created into.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, then use the default account container specified by AD. This is the `Computers` container for out-of-the-box installations of AD.
	OU *string `json:"OU,omitempty"`
	// The AD domain name for the pool. Specify this in FQDN format; for example, MyDomain.com.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, this property is required.
	Domain *string `json:"Domain,omitempty"`
	// Defines the next value that will be used if creating new AD accounts.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, the default is a sequence of `0`s or `A`s, depending on the .
	NextValue *string `json:"NextValue,omitempty"`
}

// NewUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules() *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules {
	this := UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules{}
	return &this
}

// NewUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRulesWithDefaults instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRulesWithDefaults() *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules {
	this := UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules{}
	return &this
}

// GetNamingScheme returns the NamingScheme field value if set, zero value otherwise.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNamingScheme() string {
	if o == nil || IsNil(o.NamingScheme) {
		var ret string
		return ret
	}
	return *o.NamingScheme
}

// GetNamingSchemeOk returns a tuple with the NamingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNamingSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.NamingScheme) {
		return nil, false
	}
	return o.NamingScheme, true
}

// HasNamingScheme returns a boolean if a field has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) HasNamingScheme() bool {
	if o != nil && !IsNil(o.NamingScheme) {
		return true
	}

	return false
}

// SetNamingScheme gets a reference to the given string and assigns it to the NamingScheme field.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) SetNamingScheme(v string) {
	o.NamingScheme = &v
}

// GetNamingSchemeType returns the NamingSchemeType field value if set, zero value otherwise.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNamingSchemeType() NamingSchemeType {
	if o == nil || IsNil(o.NamingSchemeType) {
		var ret NamingSchemeType
		return ret
	}
	return *o.NamingSchemeType
}

// GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNamingSchemeTypeOk() (*NamingSchemeType, bool) {
	if o == nil || IsNil(o.NamingSchemeType) {
		return nil, false
	}
	return o.NamingSchemeType, true
}

// HasNamingSchemeType returns a boolean if a field has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) HasNamingSchemeType() bool {
	if o != nil && !IsNil(o.NamingSchemeType) {
		return true
	}

	return false
}

// SetNamingSchemeType gets a reference to the given NamingSchemeType and assigns it to the NamingSchemeType field.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) SetNamingSchemeType(v NamingSchemeType) {
	o.NamingSchemeType = &v
}

// GetOU returns the OU field value if set, zero value otherwise.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetOU() string {
	if o == nil || IsNil(o.OU) {
		var ret string
		return ret
	}
	return *o.OU
}

// GetOUOk returns a tuple with the OU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetOUOk() (*string, bool) {
	if o == nil || IsNil(o.OU) {
		return nil, false
	}
	return o.OU, true
}

// HasOU returns a boolean if a field has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) HasOU() bool {
	if o != nil && !IsNil(o.OU) {
		return true
	}

	return false
}

// SetOU gets a reference to the given string and assigns it to the OU field.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) SetOU(v string) {
	o.OU = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) SetDomain(v string) {
	o.Domain = &v
}

// GetNextValue returns the NextValue field value if set, zero value otherwise.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNextValue() string {
	if o == nil || IsNil(o.NextValue) {
		var ret string
		return ret
	}
	return *o.NextValue
}

// GetNextValueOk returns a tuple with the NextValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) GetNextValueOk() (*string, bool) {
	if o == nil || IsNil(o.NextValue) {
		return nil, false
	}
	return o.NextValue, true
}

// HasNextValue returns a boolean if a field has been set.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) HasNextValue() bool {
	if o != nil && !IsNil(o.NextValue) {
		return true
	}

	return false
}

// SetNextValue gets a reference to the given string and assigns it to the NextValue field.
func (o *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) SetNextValue(v string) {
	o.NextValue = &v
}

func (o UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NamingScheme) {
		toSerialize["NamingScheme"] = o.NamingScheme
	}
	if !IsNil(o.NamingSchemeType) {
		toSerialize["NamingSchemeType"] = o.NamingSchemeType
	}
	if !IsNil(o.OU) {
		toSerialize["OU"] = o.OU
	}
	if !IsNil(o.Domain) {
		toSerialize["Domain"] = o.Domain
	}
	if !IsNil(o.NextValue) {
		toSerialize["NextValue"] = o.NextValue
	}
	return toSerialize, nil
}

type NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules struct {
	value *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules
	isSet bool
}

func (v NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) Get() *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules {
	return v.value
}

func (v *NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) Set(val *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules(val *UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) *NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules {
	return &NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules{value: val, isSet: true}
}

func (v NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


