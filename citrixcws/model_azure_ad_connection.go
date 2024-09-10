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

// checks if the AzureAdConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureAdConnection{}

// AzureAdConnection struct for AzureAdConnection
type AzureAdConnection struct {
	AuthDomainName NullableString `json:"authDomainName,omitempty"`
	AuthDomainDisplayName NullableString `json:"authDomainDisplayName,omitempty"`
	HintApp NullableString `json:"hintApp,omitempty"`
	Resources NullableString `json:"resources,omitempty"`
}

// NewAzureAdConnection instantiates a new AzureAdConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureAdConnection() *AzureAdConnection {
	this := AzureAdConnection{}
	return &this
}

// NewAzureAdConnectionWithDefaults instantiates a new AzureAdConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureAdConnectionWithDefaults() *AzureAdConnection {
	this := AzureAdConnection{}
	return &this
}

// GetAuthDomainName returns the AuthDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureAdConnection) GetAuthDomainName() string {
	if o == nil || IsNil(o.AuthDomainName.Get()) {
		var ret string
		return ret
	}
	return *o.AuthDomainName.Get()
}

// GetAuthDomainNameOk returns a tuple with the AuthDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureAdConnection) GetAuthDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDomainName.Get(), o.AuthDomainName.IsSet()
}

// HasAuthDomainName returns a boolean if a field has been set.
func (o *AzureAdConnection) HasAuthDomainName() bool {
	if o != nil && o.AuthDomainName.IsSet() {
		return true
	}

	return false
}

// SetAuthDomainName gets a reference to the given NullableString and assigns it to the AuthDomainName field.
func (o *AzureAdConnection) SetAuthDomainName(v string) {
	o.AuthDomainName.Set(&v)
}
// SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil
func (o *AzureAdConnection) SetAuthDomainNameNil() {
	o.AuthDomainName.Set(nil)
}

// UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
func (o *AzureAdConnection) UnsetAuthDomainName() {
	o.AuthDomainName.Unset()
}

// GetAuthDomainDisplayName returns the AuthDomainDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureAdConnection) GetAuthDomainDisplayName() string {
	if o == nil || IsNil(o.AuthDomainDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.AuthDomainDisplayName.Get()
}

// GetAuthDomainDisplayNameOk returns a tuple with the AuthDomainDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureAdConnection) GetAuthDomainDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDomainDisplayName.Get(), o.AuthDomainDisplayName.IsSet()
}

// HasAuthDomainDisplayName returns a boolean if a field has been set.
func (o *AzureAdConnection) HasAuthDomainDisplayName() bool {
	if o != nil && o.AuthDomainDisplayName.IsSet() {
		return true
	}

	return false
}

// SetAuthDomainDisplayName gets a reference to the given NullableString and assigns it to the AuthDomainDisplayName field.
func (o *AzureAdConnection) SetAuthDomainDisplayName(v string) {
	o.AuthDomainDisplayName.Set(&v)
}
// SetAuthDomainDisplayNameNil sets the value for AuthDomainDisplayName to be an explicit nil
func (o *AzureAdConnection) SetAuthDomainDisplayNameNil() {
	o.AuthDomainDisplayName.Set(nil)
}

// UnsetAuthDomainDisplayName ensures that no value is present for AuthDomainDisplayName, not even an explicit nil
func (o *AzureAdConnection) UnsetAuthDomainDisplayName() {
	o.AuthDomainDisplayName.Unset()
}

// GetHintApp returns the HintApp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureAdConnection) GetHintApp() string {
	if o == nil || IsNil(o.HintApp.Get()) {
		var ret string
		return ret
	}
	return *o.HintApp.Get()
}

// GetHintAppOk returns a tuple with the HintApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureAdConnection) GetHintAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HintApp.Get(), o.HintApp.IsSet()
}

// HasHintApp returns a boolean if a field has been set.
func (o *AzureAdConnection) HasHintApp() bool {
	if o != nil && o.HintApp.IsSet() {
		return true
	}

	return false
}

// SetHintApp gets a reference to the given NullableString and assigns it to the HintApp field.
func (o *AzureAdConnection) SetHintApp(v string) {
	o.HintApp.Set(&v)
}
// SetHintAppNil sets the value for HintApp to be an explicit nil
func (o *AzureAdConnection) SetHintAppNil() {
	o.HintApp.Set(nil)
}

// UnsetHintApp ensures that no value is present for HintApp, not even an explicit nil
func (o *AzureAdConnection) UnsetHintApp() {
	o.HintApp.Unset()
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureAdConnection) GetResources() string {
	if o == nil || IsNil(o.Resources.Get()) {
		var ret string
		return ret
	}
	return *o.Resources.Get()
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureAdConnection) GetResourcesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources.Get(), o.Resources.IsSet()
}

// HasResources returns a boolean if a field has been set.
func (o *AzureAdConnection) HasResources() bool {
	if o != nil && o.Resources.IsSet() {
		return true
	}

	return false
}

// SetResources gets a reference to the given NullableString and assigns it to the Resources field.
func (o *AzureAdConnection) SetResources(v string) {
	o.Resources.Set(&v)
}
// SetResourcesNil sets the value for Resources to be an explicit nil
func (o *AzureAdConnection) SetResourcesNil() {
	o.Resources.Set(nil)
}

// UnsetResources ensures that no value is present for Resources, not even an explicit nil
func (o *AzureAdConnection) UnsetResources() {
	o.Resources.Unset()
}

func (o AzureAdConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureAdConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthDomainName.IsSet() {
		toSerialize["authDomainName"] = o.AuthDomainName.Get()
	}
	if o.AuthDomainDisplayName.IsSet() {
		toSerialize["authDomainDisplayName"] = o.AuthDomainDisplayName.Get()
	}
	if o.HintApp.IsSet() {
		toSerialize["hintApp"] = o.HintApp.Get()
	}
	if o.Resources.IsSet() {
		toSerialize["resources"] = o.Resources.Get()
	}
	return toSerialize, nil
}

type NullableAzureAdConnection struct {
	value *AzureAdConnection
	isSet bool
}

func (v NullableAzureAdConnection) Get() *AzureAdConnection {
	return v.value
}

func (v *NullableAzureAdConnection) Set(val *AzureAdConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureAdConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureAdConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureAdConnection(val *AzureAdConnection) *NullableAzureAdConnection {
	return &NullableAzureAdConnection{value: val, isSet: true}
}

func (v NullableAzureAdConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureAdConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


