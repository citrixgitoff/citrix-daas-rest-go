/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the LicenseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseUser{}

// LicenseUser struct for LicenseUser
type LicenseUser struct {
	HdxLicenseState *LicenseState `json:"hdxLicenseState,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Windows365Licenses []string `json:"windows365Licenses,omitempty"`
	// AAD User ID for which we want to assign/revoke license
	UserId NullableString `json:"userId,omitempty"`
	// AAD UPN for provided User
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewLicenseUser instantiates a new LicenseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseUser() *LicenseUser {
	this := LicenseUser{}
	return &this
}

// NewLicenseUserWithDefaults instantiates a new LicenseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseUserWithDefaults() *LicenseUser {
	this := LicenseUser{}
	return &this
}

// GetHdxLicenseState returns the HdxLicenseState field value if set, zero value otherwise.
func (o *LicenseUser) GetHdxLicenseState() LicenseState {
	if o == nil || IsNil(o.HdxLicenseState) {
		var ret LicenseState
		return ret
	}
	return *o.HdxLicenseState
}

// GetHdxLicenseStateOk returns a tuple with the HdxLicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseUser) GetHdxLicenseStateOk() (*LicenseState, bool) {
	if o == nil || IsNil(o.HdxLicenseState) {
		return nil, false
	}
	return o.HdxLicenseState, true
}

// HasHdxLicenseState returns a boolean if a field has been set.
func (o *LicenseUser) HasHdxLicenseState() bool {
	if o != nil && !IsNil(o.HdxLicenseState) {
		return true
	}

	return false
}

// SetHdxLicenseState gets a reference to the given LicenseState and assigns it to the HdxLicenseState field.
func (o *LicenseUser) SetHdxLicenseState(v LicenseState) {
	o.HdxLicenseState = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LicenseUser) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *LicenseUser) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *LicenseUser) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *LicenseUser) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetWindows365Licenses returns the Windows365Licenses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseUser) GetWindows365Licenses() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Windows365Licenses
}

// GetWindows365LicensesOk returns a tuple with the Windows365Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseUser) GetWindows365LicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.Windows365Licenses) {
		return nil, false
	}
	return o.Windows365Licenses, true
}

// HasWindows365Licenses returns a boolean if a field has been set.
func (o *LicenseUser) HasWindows365Licenses() bool {
	if o != nil && IsNil(o.Windows365Licenses) {
		return true
	}

	return false
}

// SetWindows365Licenses gets a reference to the given []string and assigns it to the Windows365Licenses field.
func (o *LicenseUser) SetWindows365Licenses(v []string) {
	o.Windows365Licenses = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseUser) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *LicenseUser) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *LicenseUser) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *LicenseUser) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *LicenseUser) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseUser) GetUserPrincipalName() string {
	if o == nil || IsNil(o.UserPrincipalName.Get()) {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseUser) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *LicenseUser) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *LicenseUser) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *LicenseUser) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *LicenseUser) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o LicenseUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HdxLicenseState) {
		toSerialize["hdxLicenseState"] = o.HdxLicenseState
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Windows365Licenses != nil {
		toSerialize["windows365Licenses"] = o.Windows365Licenses
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return toSerialize, nil
}

type NullableLicenseUser struct {
	value *LicenseUser
	isSet bool
}

func (v NullableLicenseUser) Get() *LicenseUser {
	return v.value
}

func (v *NullableLicenseUser) Set(val *LicenseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseUser(val *LicenseUser) *NullableLicenseUser {
	return &NullableLicenseUser{value: val, isSet: true}
}

func (v NullableLicenseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


