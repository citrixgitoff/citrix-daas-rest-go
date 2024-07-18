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

// checks if the UpdateServiceAccountRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceAccountRequestModel{}

// UpdateServiceAccountRequestModel Request model for updating a service account.
type UpdateServiceAccountRequestModel struct {
	// The identifier for the service account
	AccountId NullableString `json:"AccountId,omitempty"`
	// The secret for the service account. E.g. Azure application (client) secret if 'IdentityProviderType' is AzureAD. The secret will be encrypted and stored in database.
	AccountSecret NullableString `json:"AccountSecret,omitempty"`
	AccountSecretFormat *IdentityPasswordFormat `json:"AccountSecretFormat,omitempty"`
	// The secret expiration time for the service account
	SecretExpiryTime NullableString `json:"SecretExpiryTime,omitempty"`
	// Gets or sets capabilities for the service account
	Capabilities []string `json:"Capabilities,omitempty"`
	// Gets or sets the display name
	DisplayName NullableString `json:"DisplayName,omitempty"`
	// Gets or sets the description
	Description NullableString `json:"Description,omitempty"`
	// Gets or sets the scopes for the service account.
	Scopes []string `json:"Scopes,omitempty"`
	// Tenants to associate with the service account.
	Tenants []string `json:"Tenants,omitempty"`
}

// NewUpdateServiceAccountRequestModel instantiates a new UpdateServiceAccountRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceAccountRequestModel() *UpdateServiceAccountRequestModel {
	this := UpdateServiceAccountRequestModel{}
	return &this
}

// NewUpdateServiceAccountRequestModelWithDefaults instantiates a new UpdateServiceAccountRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceAccountRequestModelWithDefaults() *UpdateServiceAccountRequestModel {
	this := UpdateServiceAccountRequestModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *UpdateServiceAccountRequestModel) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *UpdateServiceAccountRequestModel) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *UpdateServiceAccountRequestModel) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetAccountSecret returns the AccountSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetAccountSecret() string {
	if o == nil || IsNil(o.AccountSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AccountSecret.Get()
}

// GetAccountSecretOk returns a tuple with the AccountSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetAccountSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSecret.Get(), o.AccountSecret.IsSet()
}

// HasAccountSecret returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasAccountSecret() bool {
	if o != nil && o.AccountSecret.IsSet() {
		return true
	}

	return false
}

// SetAccountSecret gets a reference to the given NullableString and assigns it to the AccountSecret field.
func (o *UpdateServiceAccountRequestModel) SetAccountSecret(v string) {
	o.AccountSecret.Set(&v)
}
// SetAccountSecretNil sets the value for AccountSecret to be an explicit nil
func (o *UpdateServiceAccountRequestModel) SetAccountSecretNil() {
	o.AccountSecret.Set(nil)
}

// UnsetAccountSecret ensures that no value is present for AccountSecret, not even an explicit nil
func (o *UpdateServiceAccountRequestModel) UnsetAccountSecret() {
	o.AccountSecret.Unset()
}

// GetAccountSecretFormat returns the AccountSecretFormat field value if set, zero value otherwise.
func (o *UpdateServiceAccountRequestModel) GetAccountSecretFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.AccountSecretFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.AccountSecretFormat
}

// GetAccountSecretFormatOk returns a tuple with the AccountSecretFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceAccountRequestModel) GetAccountSecretFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.AccountSecretFormat) {
		return nil, false
	}
	return o.AccountSecretFormat, true
}

// HasAccountSecretFormat returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasAccountSecretFormat() bool {
	if o != nil && !IsNil(o.AccountSecretFormat) {
		return true
	}

	return false
}

// SetAccountSecretFormat gets a reference to the given IdentityPasswordFormat and assigns it to the AccountSecretFormat field.
func (o *UpdateServiceAccountRequestModel) SetAccountSecretFormat(v IdentityPasswordFormat) {
	o.AccountSecretFormat = &v
}

// GetSecretExpiryTime returns the SecretExpiryTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetSecretExpiryTime() string {
	if o == nil || IsNil(o.SecretExpiryTime.Get()) {
		var ret string
		return ret
	}
	return *o.SecretExpiryTime.Get()
}

// GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetSecretExpiryTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretExpiryTime.Get(), o.SecretExpiryTime.IsSet()
}

// HasSecretExpiryTime returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasSecretExpiryTime() bool {
	if o != nil && o.SecretExpiryTime.IsSet() {
		return true
	}

	return false
}

// SetSecretExpiryTime gets a reference to the given NullableString and assigns it to the SecretExpiryTime field.
func (o *UpdateServiceAccountRequestModel) SetSecretExpiryTime(v string) {
	o.SecretExpiryTime.Set(&v)
}
// SetSecretExpiryTimeNil sets the value for SecretExpiryTime to be an explicit nil
func (o *UpdateServiceAccountRequestModel) SetSecretExpiryTimeNil() {
	o.SecretExpiryTime.Set(nil)
}

// UnsetSecretExpiryTime ensures that no value is present for SecretExpiryTime, not even an explicit nil
func (o *UpdateServiceAccountRequestModel) UnsetSecretExpiryTime() {
	o.SecretExpiryTime.Unset()
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasCapabilities() bool {
	if o != nil && IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *UpdateServiceAccountRequestModel) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UpdateServiceAccountRequestModel) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UpdateServiceAccountRequestModel) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UpdateServiceAccountRequestModel) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateServiceAccountRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateServiceAccountRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateServiceAccountRequestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *UpdateServiceAccountRequestModel) SetScopes(v []string) {
	o.Scopes = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountRequestModel) GetTenants() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountRequestModel) GetTenantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *UpdateServiceAccountRequestModel) HasTenants() bool {
	if o != nil && IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []string and assigns it to the Tenants field.
func (o *UpdateServiceAccountRequestModel) SetTenants(v []string) {
	o.Tenants = v
}

func (o UpdateServiceAccountRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceAccountRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId.IsSet() {
		toSerialize["AccountId"] = o.AccountId.Get()
	}
	if o.AccountSecret.IsSet() {
		toSerialize["AccountSecret"] = o.AccountSecret.Get()
	}
	if !IsNil(o.AccountSecretFormat) {
		toSerialize["AccountSecretFormat"] = o.AccountSecretFormat
	}
	if o.SecretExpiryTime.IsSet() {
		toSerialize["SecretExpiryTime"] = o.SecretExpiryTime.Get()
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Scopes != nil {
		toSerialize["Scopes"] = o.Scopes
	}
	if o.Tenants != nil {
		toSerialize["Tenants"] = o.Tenants
	}
	return toSerialize, nil
}

type NullableUpdateServiceAccountRequestModel struct {
	value *UpdateServiceAccountRequestModel
	isSet bool
}

func (v NullableUpdateServiceAccountRequestModel) Get() *UpdateServiceAccountRequestModel {
	return v.value
}

func (v *NullableUpdateServiceAccountRequestModel) Set(val *UpdateServiceAccountRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceAccountRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceAccountRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceAccountRequestModel(val *UpdateServiceAccountRequestModel) *NullableUpdateServiceAccountRequestModel {
	return &NullableUpdateServiceAccountRequestModel{value: val, isSet: true}
}

func (v NullableUpdateServiceAccountRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceAccountRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


