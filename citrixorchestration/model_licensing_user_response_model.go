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

// checks if the LicensingUserResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensingUserResponseModel{}

// LicensingUserResponseModel Response object for the license administrator.
type LicensingUserResponseModel struct {
	// Gets or sets the account name
	Account *string `json:"Account,omitempty"`
	// Gets or sets the account SID 
	AccountSid *string `json:"AccountSid,omitempty"`
	PermissionLevel *LicensingPermissionLevel `json:"PermissionLevel,omitempty"`
	// Gets or sets a value indicating whether it is an individual account or a group
	IsGroup *bool `json:"IsGroup,omitempty"`
}

// NewLicensingUserResponseModel instantiates a new LicensingUserResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensingUserResponseModel() *LicensingUserResponseModel {
	this := LicensingUserResponseModel{}
	return &this
}

// NewLicensingUserResponseModelWithDefaults instantiates a new LicensingUserResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensingUserResponseModelWithDefaults() *LicensingUserResponseModel {
	this := LicensingUserResponseModel{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *LicensingUserResponseModel) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingUserResponseModel) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *LicensingUserResponseModel) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *LicensingUserResponseModel) SetAccount(v string) {
	o.Account = &v
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise.
func (o *LicensingUserResponseModel) GetAccountSid() string {
	if o == nil || IsNil(o.AccountSid) {
		var ret string
		return ret
	}
	return *o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingUserResponseModel) GetAccountSidOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSid) {
		return nil, false
	}
	return o.AccountSid, true
}

// HasAccountSid returns a boolean if a field has been set.
func (o *LicensingUserResponseModel) HasAccountSid() bool {
	if o != nil && !IsNil(o.AccountSid) {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given string and assigns it to the AccountSid field.
func (o *LicensingUserResponseModel) SetAccountSid(v string) {
	o.AccountSid = &v
}

// GetPermissionLevel returns the PermissionLevel field value if set, zero value otherwise.
func (o *LicensingUserResponseModel) GetPermissionLevel() LicensingPermissionLevel {
	if o == nil || IsNil(o.PermissionLevel) {
		var ret LicensingPermissionLevel
		return ret
	}
	return *o.PermissionLevel
}

// GetPermissionLevelOk returns a tuple with the PermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingUserResponseModel) GetPermissionLevelOk() (*LicensingPermissionLevel, bool) {
	if o == nil || IsNil(o.PermissionLevel) {
		return nil, false
	}
	return o.PermissionLevel, true
}

// HasPermissionLevel returns a boolean if a field has been set.
func (o *LicensingUserResponseModel) HasPermissionLevel() bool {
	if o != nil && !IsNil(o.PermissionLevel) {
		return true
	}

	return false
}

// SetPermissionLevel gets a reference to the given LicensingPermissionLevel and assigns it to the PermissionLevel field.
func (o *LicensingUserResponseModel) SetPermissionLevel(v LicensingPermissionLevel) {
	o.PermissionLevel = &v
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise.
func (o *LicensingUserResponseModel) GetIsGroup() bool {
	if o == nil || IsNil(o.IsGroup) {
		var ret bool
		return ret
	}
	return *o.IsGroup
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingUserResponseModel) GetIsGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGroup) {
		return nil, false
	}
	return o.IsGroup, true
}

// HasIsGroup returns a boolean if a field has been set.
func (o *LicensingUserResponseModel) HasIsGroup() bool {
	if o != nil && !IsNil(o.IsGroup) {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given bool and assigns it to the IsGroup field.
func (o *LicensingUserResponseModel) SetIsGroup(v bool) {
	o.IsGroup = &v
}

func (o LicensingUserResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensingUserResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["Account"] = o.Account
	}
	if !IsNil(o.AccountSid) {
		toSerialize["AccountSid"] = o.AccountSid
	}
	if !IsNil(o.PermissionLevel) {
		toSerialize["PermissionLevel"] = o.PermissionLevel
	}
	if !IsNil(o.IsGroup) {
		toSerialize["IsGroup"] = o.IsGroup
	}
	return toSerialize, nil
}

type NullableLicensingUserResponseModel struct {
	value *LicensingUserResponseModel
	isSet bool
}

func (v NullableLicensingUserResponseModel) Get() *LicensingUserResponseModel {
	return v.value
}

func (v *NullableLicensingUserResponseModel) Set(val *LicensingUserResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingUserResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingUserResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingUserResponseModel(val *LicensingUserResponseModel) *NullableLicensingUserResponseModel {
	return &NullableLicensingUserResponseModel{value: val, isSet: true}
}

func (v NullableLicensingUserResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingUserResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


