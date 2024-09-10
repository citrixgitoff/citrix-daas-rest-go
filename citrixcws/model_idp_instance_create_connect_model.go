/*
Citrix.CloudServices.Cws.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixcws

import (
	"encoding/json"
)

// checks if the IdpInstanceCreateConnectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpInstanceCreateConnectModel{}

// IdpInstanceCreateConnectModel struct for IdpInstanceCreateConnectModel
type IdpInstanceCreateConnectModel struct {
	AdaptiveAuthenticationConnectionDetails *AdaptiveAuthenticationConnectionModel `json:"adaptiveAuthenticationConnectionDetails,omitempty"`
	AuthDomainName NullableString `json:"authDomainName,omitempty"`
	AzureAd *AzureAdConnectionSettings `json:"azureAd,omitempty"`
	GatewayConnectionDetails *GatewayConnectionModel `json:"gatewayConnectionDetails,omitempty"`
	GoogleConnectionModel *GoogleConnectionModel `json:"googleConnectionModel,omitempty"`
	IdentityProviderType string `json:"identityProviderType"`
	OktaConnectionModel *OktaConnectionModel `json:"oktaConnectionModel,omitempty"`
	SamlConnectionModel *SamlConnectionModel `json:"samlConnectionModel,omitempty"`
	IdentityProviderNickname string `json:"identityProviderNickname"`
}

// NewIdpInstanceCreateConnectModel instantiates a new IdpInstanceCreateConnectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpInstanceCreateConnectModel(identityProviderType string, identityProviderNickname string) *IdpInstanceCreateConnectModel {
	this := IdpInstanceCreateConnectModel{}
	this.IdentityProviderType = identityProviderType
	this.IdentityProviderNickname = identityProviderNickname
	return &this
}

// NewIdpInstanceCreateConnectModelWithDefaults instantiates a new IdpInstanceCreateConnectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpInstanceCreateConnectModelWithDefaults() *IdpInstanceCreateConnectModel {
	this := IdpInstanceCreateConnectModel{}
	return &this
}

// GetAdaptiveAuthenticationConnectionDetails returns the AdaptiveAuthenticationConnectionDetails field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetAdaptiveAuthenticationConnectionDetails() AdaptiveAuthenticationConnectionModel {
	if o == nil || IsNil(o.AdaptiveAuthenticationConnectionDetails) {
		var ret AdaptiveAuthenticationConnectionModel
		return ret
	}
	return *o.AdaptiveAuthenticationConnectionDetails
}

// GetAdaptiveAuthenticationConnectionDetailsOk returns a tuple with the AdaptiveAuthenticationConnectionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetAdaptiveAuthenticationConnectionDetailsOk() (*AdaptiveAuthenticationConnectionModel, bool) {
	if o == nil || IsNil(o.AdaptiveAuthenticationConnectionDetails) {
		return nil, false
	}
	return o.AdaptiveAuthenticationConnectionDetails, true
}

// HasAdaptiveAuthenticationConnectionDetails returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasAdaptiveAuthenticationConnectionDetails() bool {
	if o != nil && !IsNil(o.AdaptiveAuthenticationConnectionDetails) {
		return true
	}

	return false
}

// SetAdaptiveAuthenticationConnectionDetails gets a reference to the given AdaptiveAuthenticationConnectionModel and assigns it to the AdaptiveAuthenticationConnectionDetails field.
func (o *IdpInstanceCreateConnectModel) SetAdaptiveAuthenticationConnectionDetails(v AdaptiveAuthenticationConnectionModel) {
	o.AdaptiveAuthenticationConnectionDetails = &v
}

// GetAuthDomainName returns the AuthDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpInstanceCreateConnectModel) GetAuthDomainName() string {
	if o == nil || IsNil(o.AuthDomainName.Get()) {
		var ret string
		return ret
	}
	return *o.AuthDomainName.Get()
}

// GetAuthDomainNameOk returns a tuple with the AuthDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpInstanceCreateConnectModel) GetAuthDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDomainName.Get(), o.AuthDomainName.IsSet()
}

// HasAuthDomainName returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasAuthDomainName() bool {
	if o != nil && o.AuthDomainName.IsSet() {
		return true
	}

	return false
}

// SetAuthDomainName gets a reference to the given NullableString and assigns it to the AuthDomainName field.
func (o *IdpInstanceCreateConnectModel) SetAuthDomainName(v string) {
	o.AuthDomainName.Set(&v)
}
// SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil
func (o *IdpInstanceCreateConnectModel) SetAuthDomainNameNil() {
	o.AuthDomainName.Set(nil)
}

// UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
func (o *IdpInstanceCreateConnectModel) UnsetAuthDomainName() {
	o.AuthDomainName.Unset()
}

// GetAzureAd returns the AzureAd field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetAzureAd() AzureAdConnectionSettings {
	if o == nil || IsNil(o.AzureAd) {
		var ret AzureAdConnectionSettings
		return ret
	}
	return *o.AzureAd
}

// GetAzureAdOk returns a tuple with the AzureAd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetAzureAdOk() (*AzureAdConnectionSettings, bool) {
	if o == nil || IsNil(o.AzureAd) {
		return nil, false
	}
	return o.AzureAd, true
}

// HasAzureAd returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasAzureAd() bool {
	if o != nil && !IsNil(o.AzureAd) {
		return true
	}

	return false
}

// SetAzureAd gets a reference to the given AzureAdConnectionSettings and assigns it to the AzureAd field.
func (o *IdpInstanceCreateConnectModel) SetAzureAd(v AzureAdConnectionSettings) {
	o.AzureAd = &v
}

// GetGatewayConnectionDetails returns the GatewayConnectionDetails field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetGatewayConnectionDetails() GatewayConnectionModel {
	if o == nil || IsNil(o.GatewayConnectionDetails) {
		var ret GatewayConnectionModel
		return ret
	}
	return *o.GatewayConnectionDetails
}

// GetGatewayConnectionDetailsOk returns a tuple with the GatewayConnectionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetGatewayConnectionDetailsOk() (*GatewayConnectionModel, bool) {
	if o == nil || IsNil(o.GatewayConnectionDetails) {
		return nil, false
	}
	return o.GatewayConnectionDetails, true
}

// HasGatewayConnectionDetails returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasGatewayConnectionDetails() bool {
	if o != nil && !IsNil(o.GatewayConnectionDetails) {
		return true
	}

	return false
}

// SetGatewayConnectionDetails gets a reference to the given GatewayConnectionModel and assigns it to the GatewayConnectionDetails field.
func (o *IdpInstanceCreateConnectModel) SetGatewayConnectionDetails(v GatewayConnectionModel) {
	o.GatewayConnectionDetails = &v
}

// GetGoogleConnectionModel returns the GoogleConnectionModel field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetGoogleConnectionModel() GoogleConnectionModel {
	if o == nil || IsNil(o.GoogleConnectionModel) {
		var ret GoogleConnectionModel
		return ret
	}
	return *o.GoogleConnectionModel
}

// GetGoogleConnectionModelOk returns a tuple with the GoogleConnectionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetGoogleConnectionModelOk() (*GoogleConnectionModel, bool) {
	if o == nil || IsNil(o.GoogleConnectionModel) {
		return nil, false
	}
	return o.GoogleConnectionModel, true
}

// HasGoogleConnectionModel returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasGoogleConnectionModel() bool {
	if o != nil && !IsNil(o.GoogleConnectionModel) {
		return true
	}

	return false
}

// SetGoogleConnectionModel gets a reference to the given GoogleConnectionModel and assigns it to the GoogleConnectionModel field.
func (o *IdpInstanceCreateConnectModel) SetGoogleConnectionModel(v GoogleConnectionModel) {
	o.GoogleConnectionModel = &v
}

// GetIdentityProviderType returns the IdentityProviderType field value
func (o *IdpInstanceCreateConnectModel) GetIdentityProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderType
}

// GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field value
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetIdentityProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderType, true
}

// SetIdentityProviderType sets field value
func (o *IdpInstanceCreateConnectModel) SetIdentityProviderType(v string) {
	o.IdentityProviderType = v
}

// GetOktaConnectionModel returns the OktaConnectionModel field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetOktaConnectionModel() OktaConnectionModel {
	if o == nil || IsNil(o.OktaConnectionModel) {
		var ret OktaConnectionModel
		return ret
	}
	return *o.OktaConnectionModel
}

// GetOktaConnectionModelOk returns a tuple with the OktaConnectionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetOktaConnectionModelOk() (*OktaConnectionModel, bool) {
	if o == nil || IsNil(o.OktaConnectionModel) {
		return nil, false
	}
	return o.OktaConnectionModel, true
}

// HasOktaConnectionModel returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasOktaConnectionModel() bool {
	if o != nil && !IsNil(o.OktaConnectionModel) {
		return true
	}

	return false
}

// SetOktaConnectionModel gets a reference to the given OktaConnectionModel and assigns it to the OktaConnectionModel field.
func (o *IdpInstanceCreateConnectModel) SetOktaConnectionModel(v OktaConnectionModel) {
	o.OktaConnectionModel = &v
}

// GetSamlConnectionModel returns the SamlConnectionModel field value if set, zero value otherwise.
func (o *IdpInstanceCreateConnectModel) GetSamlConnectionModel() SamlConnectionModel {
	if o == nil || IsNil(o.SamlConnectionModel) {
		var ret SamlConnectionModel
		return ret
	}
	return *o.SamlConnectionModel
}

// GetSamlConnectionModelOk returns a tuple with the SamlConnectionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetSamlConnectionModelOk() (*SamlConnectionModel, bool) {
	if o == nil || IsNil(o.SamlConnectionModel) {
		return nil, false
	}
	return o.SamlConnectionModel, true
}

// HasSamlConnectionModel returns a boolean if a field has been set.
func (o *IdpInstanceCreateConnectModel) HasSamlConnectionModel() bool {
	if o != nil && !IsNil(o.SamlConnectionModel) {
		return true
	}

	return false
}

// SetSamlConnectionModel gets a reference to the given SamlConnectionModel and assigns it to the SamlConnectionModel field.
func (o *IdpInstanceCreateConnectModel) SetSamlConnectionModel(v SamlConnectionModel) {
	o.SamlConnectionModel = &v
}

// GetIdentityProviderNickname returns the IdentityProviderNickname field value
func (o *IdpInstanceCreateConnectModel) GetIdentityProviderNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderNickname
}

// GetIdentityProviderNicknameOk returns a tuple with the IdentityProviderNickname field value
// and a boolean to check if the value has been set.
func (o *IdpInstanceCreateConnectModel) GetIdentityProviderNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderNickname, true
}

// SetIdentityProviderNickname sets field value
func (o *IdpInstanceCreateConnectModel) SetIdentityProviderNickname(v string) {
	o.IdentityProviderNickname = v
}

func (o IdpInstanceCreateConnectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpInstanceCreateConnectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdaptiveAuthenticationConnectionDetails) {
		toSerialize["adaptiveAuthenticationConnectionDetails"] = o.AdaptiveAuthenticationConnectionDetails
	}
	if o.AuthDomainName.IsSet() {
		toSerialize["authDomainName"] = o.AuthDomainName.Get()
	}
	if !IsNil(o.AzureAd) {
		toSerialize["azureAd"] = o.AzureAd
	}
	if !IsNil(o.GatewayConnectionDetails) {
		toSerialize["gatewayConnectionDetails"] = o.GatewayConnectionDetails
	}
	if !IsNil(o.GoogleConnectionModel) {
		toSerialize["googleConnectionModel"] = o.GoogleConnectionModel
	}
	toSerialize["identityProviderType"] = o.IdentityProviderType
	if !IsNil(o.OktaConnectionModel) {
		toSerialize["oktaConnectionModel"] = o.OktaConnectionModel
	}
	if !IsNil(o.SamlConnectionModel) {
		toSerialize["samlConnectionModel"] = o.SamlConnectionModel
	}
	toSerialize["identityProviderNickname"] = o.IdentityProviderNickname
	return toSerialize, nil
}

type NullableIdpInstanceCreateConnectModel struct {
	value *IdpInstanceCreateConnectModel
	isSet bool
}

func (v NullableIdpInstanceCreateConnectModel) Get() *IdpInstanceCreateConnectModel {
	return v.value
}

func (v *NullableIdpInstanceCreateConnectModel) Set(val *IdpInstanceCreateConnectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpInstanceCreateConnectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpInstanceCreateConnectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpInstanceCreateConnectModel(val *IdpInstanceCreateConnectModel) *NullableIdpInstanceCreateConnectModel {
	return &NullableIdpInstanceCreateConnectModel{value: val, isSet: true}
}

func (v NullableIdpInstanceCreateConnectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpInstanceCreateConnectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


