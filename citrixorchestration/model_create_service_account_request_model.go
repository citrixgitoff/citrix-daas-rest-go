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

// checks if the CreateServiceAccountRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceAccountRequestModel{}

// CreateServiceAccountRequestModel Request model for creating a service account.
type CreateServiceAccountRequestModel struct {
	// The identity provider type for the service account
	IdentityProviderType string `json:"IdentityProviderType"`
	// The identity provider id for the service account
	IdentityProviderIdentifier string `json:"IdentityProviderIdentifier"`
	// The identifier for the service account
	AccountId string `json:"AccountId"`
	// The secret expiration time for the service account
	SecretExpiryTime string `json:"SecretExpiryTime"`
	// The secret for the service account. E.g. Azure application (client) secret if 'IdentityProviderType' is AzureAD.
	AccountSecret string `json:"AccountSecret"`
	AccountSecretFormat *IdentityPasswordFormat `json:"AccountSecretFormat,omitempty"`
	// The capabilities for the service account
	Capabilities []string `json:"Capabilities,omitempty"`
	// The scopes for this object
	Scopes []string `json:"Scopes,omitempty"`
	// The tenant id
	TenantId NullableString `json:"TenantId,omitempty"`
	// Gets or sets the display name
	DisplayName NullableString `json:"DisplayName,omitempty"`
	// Gets or sets the description
	Description NullableString `json:"Description,omitempty"`
}

// NewCreateServiceAccountRequestModel instantiates a new CreateServiceAccountRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceAccountRequestModel(identityProviderType string, identityProviderIdentifier string, accountId string, secretExpiryTime string, accountSecret string) *CreateServiceAccountRequestModel {
	this := CreateServiceAccountRequestModel{}
	this.IdentityProviderType = identityProviderType
	this.IdentityProviderIdentifier = identityProviderIdentifier
	this.AccountId = accountId
	this.SecretExpiryTime = secretExpiryTime
	this.AccountSecret = accountSecret
	return &this
}

// NewCreateServiceAccountRequestModelWithDefaults instantiates a new CreateServiceAccountRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceAccountRequestModelWithDefaults() *CreateServiceAccountRequestModel {
	this := CreateServiceAccountRequestModel{}
	return &this
}

// GetIdentityProviderType returns the IdentityProviderType field value
func (o *CreateServiceAccountRequestModel) GetIdentityProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderType
}

// GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetIdentityProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderType, true
}

// SetIdentityProviderType sets field value
func (o *CreateServiceAccountRequestModel) SetIdentityProviderType(v string) {
	o.IdentityProviderType = v
}

// GetIdentityProviderIdentifier returns the IdentityProviderIdentifier field value
func (o *CreateServiceAccountRequestModel) GetIdentityProviderIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderIdentifier
}

// GetIdentityProviderIdentifierOk returns a tuple with the IdentityProviderIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetIdentityProviderIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderIdentifier, true
}

// SetIdentityProviderIdentifier sets field value
func (o *CreateServiceAccountRequestModel) SetIdentityProviderIdentifier(v string) {
	o.IdentityProviderIdentifier = v
}

// GetAccountId returns the AccountId field value
func (o *CreateServiceAccountRequestModel) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CreateServiceAccountRequestModel) SetAccountId(v string) {
	o.AccountId = v
}

// GetSecretExpiryTime returns the SecretExpiryTime field value
func (o *CreateServiceAccountRequestModel) GetSecretExpiryTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretExpiryTime
}

// GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetSecretExpiryTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretExpiryTime, true
}

// SetSecretExpiryTime sets field value
func (o *CreateServiceAccountRequestModel) SetSecretExpiryTime(v string) {
	o.SecretExpiryTime = v
}

// GetAccountSecret returns the AccountSecret field value
func (o *CreateServiceAccountRequestModel) GetAccountSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSecret
}

// GetAccountSecretOk returns a tuple with the AccountSecret field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetAccountSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSecret, true
}

// SetAccountSecret sets field value
func (o *CreateServiceAccountRequestModel) SetAccountSecret(v string) {
	o.AccountSecret = v
}

// GetAccountSecretFormat returns the AccountSecretFormat field value if set, zero value otherwise.
func (o *CreateServiceAccountRequestModel) GetAccountSecretFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.AccountSecretFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.AccountSecretFormat
}

// GetAccountSecretFormatOk returns a tuple with the AccountSecretFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountRequestModel) GetAccountSecretFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.AccountSecretFormat) {
		return nil, false
	}
	return o.AccountSecretFormat, true
}

// HasAccountSecretFormat returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasAccountSecretFormat() bool {
	if o != nil && !IsNil(o.AccountSecretFormat) {
		return true
	}

	return false
}

// SetAccountSecretFormat gets a reference to the given IdentityPasswordFormat and assigns it to the AccountSecretFormat field.
func (o *CreateServiceAccountRequestModel) SetAccountSecretFormat(v IdentityPasswordFormat) {
	o.AccountSecretFormat = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceAccountRequestModel) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceAccountRequestModel) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasCapabilities() bool {
	if o != nil && IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *CreateServiceAccountRequestModel) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceAccountRequestModel) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceAccountRequestModel) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CreateServiceAccountRequestModel) SetScopes(v []string) {
	o.Scopes = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceAccountRequestModel) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceAccountRequestModel) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *CreateServiceAccountRequestModel) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *CreateServiceAccountRequestModel) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *CreateServiceAccountRequestModel) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceAccountRequestModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceAccountRequestModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CreateServiceAccountRequestModel) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateServiceAccountRequestModel) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateServiceAccountRequestModel) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceAccountRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceAccountRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateServiceAccountRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateServiceAccountRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateServiceAccountRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateServiceAccountRequestModel) UnsetDescription() {
	o.Description.Unset()
}

func (o CreateServiceAccountRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceAccountRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IdentityProviderType"] = o.IdentityProviderType
	toSerialize["IdentityProviderIdentifier"] = o.IdentityProviderIdentifier
	toSerialize["AccountId"] = o.AccountId
	toSerialize["SecretExpiryTime"] = o.SecretExpiryTime
	toSerialize["AccountSecret"] = o.AccountSecret
	if !IsNil(o.AccountSecretFormat) {
		toSerialize["AccountSecretFormat"] = o.AccountSecretFormat
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if o.Scopes != nil {
		toSerialize["Scopes"] = o.Scopes
	}
	if o.TenantId.IsSet() {
		toSerialize["TenantId"] = o.TenantId.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableCreateServiceAccountRequestModel struct {
	value *CreateServiceAccountRequestModel
	isSet bool
}

func (v NullableCreateServiceAccountRequestModel) Get() *CreateServiceAccountRequestModel {
	return v.value
}

func (v *NullableCreateServiceAccountRequestModel) Set(val *CreateServiceAccountRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceAccountRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceAccountRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceAccountRequestModel(val *CreateServiceAccountRequestModel) *NullableCreateServiceAccountRequestModel {
	return &NullableCreateServiceAccountRequestModel{value: val, isSet: true}
}

func (v NullableCreateServiceAccountRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceAccountRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


