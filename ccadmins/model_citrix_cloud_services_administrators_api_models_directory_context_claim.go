/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim{}

// CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim struct for CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim
type CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim struct {
	Forest NullableString `json:"forest,omitempty"`
	Domain NullableString `json:"domain,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	IdentityProvider NullableString `json:"identityProvider,omitempty"`
}

// NewCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim instantiates a new CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim() *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim {
	this := CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim{}
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaimWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaimWithDefaults() *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim {
	this := CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim{}
	return &this
}

// GetForest returns the Forest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetForest() string {
	if o == nil || IsNil(o.Forest.Get()) {
		var ret string
		return ret
	}
	return *o.Forest.Get()
}

// GetForestOk returns a tuple with the Forest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetForestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forest.Get(), o.Forest.IsSet()
}

// HasForest returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) HasForest() bool {
	if o != nil && o.Forest.IsSet() {
		return true
	}

	return false
}

// SetForest gets a reference to the given NullableString and assigns it to the Forest field.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetForest(v string) {
	o.Forest.Set(&v)
}
// SetForestNil sets the value for Forest to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetForestNil() {
	o.Forest.Set(nil)
}

// UnsetForest ensures that no value is present for Forest, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) UnsetForest() {
	o.Forest.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) UnsetDomain() {
	o.Domain.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetIdentityProvider() string {
	if o == nil || IsNil(o.IdentityProvider.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityProvider.Get()
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) GetIdentityProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProvider.Get(), o.IdentityProvider.IsSet()
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider.IsSet() {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given NullableString and assigns it to the IdentityProvider field.
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetIdentityProvider(v string) {
	o.IdentityProvider.Set(&v)
}
// SetIdentityProviderNil sets the value for IdentityProvider to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) SetIdentityProviderNil() {
	o.IdentityProvider.Set(nil)
}

// UnsetIdentityProvider ensures that no value is present for IdentityProvider, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) UnsetIdentityProvider() {
	o.IdentityProvider.Unset()
}

func (o CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Forest.IsSet() {
		toSerialize["forest"] = o.Forest.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.IdentityProvider.IsSet() {
		toSerialize["identityProvider"] = o.IdentityProvider.Get()
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim struct {
	value *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) Get() *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) Set(val *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim(val *CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) *NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim {
	return &NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


