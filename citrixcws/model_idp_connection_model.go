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

// checks if the IdpConnectionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpConnectionModel{}

// IdpConnectionModel struct for IdpConnectionModel
type IdpConnectionModel struct {
	IdpType NullableString `json:"idpType,omitempty"`
	ConnectedApp NullableString `json:"connectedApp,omitempty"`
	ConnectionStatus *bool `json:"connectionStatus,omitempty"`
	IdpProperties map[string]string `json:"idpProperties,omitempty"`
}

// NewIdpConnectionModel instantiates a new IdpConnectionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpConnectionModel() *IdpConnectionModel {
	this := IdpConnectionModel{}
	return &this
}

// NewIdpConnectionModelWithDefaults instantiates a new IdpConnectionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpConnectionModelWithDefaults() *IdpConnectionModel {
	this := IdpConnectionModel{}
	return &this
}

// GetIdpType returns the IdpType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpConnectionModel) GetIdpType() string {
	if o == nil || IsNil(o.IdpType.Get()) {
		var ret string
		return ret
	}
	return *o.IdpType.Get()
}

// GetIdpTypeOk returns a tuple with the IdpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpConnectionModel) GetIdpTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdpType.Get(), o.IdpType.IsSet()
}

// HasIdpType returns a boolean if a field has been set.
func (o *IdpConnectionModel) HasIdpType() bool {
	if o != nil && o.IdpType.IsSet() {
		return true
	}

	return false
}

// SetIdpType gets a reference to the given NullableString and assigns it to the IdpType field.
func (o *IdpConnectionModel) SetIdpType(v string) {
	o.IdpType.Set(&v)
}
// SetIdpTypeNil sets the value for IdpType to be an explicit nil
func (o *IdpConnectionModel) SetIdpTypeNil() {
	o.IdpType.Set(nil)
}

// UnsetIdpType ensures that no value is present for IdpType, not even an explicit nil
func (o *IdpConnectionModel) UnsetIdpType() {
	o.IdpType.Unset()
}

// GetConnectedApp returns the ConnectedApp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpConnectionModel) GetConnectedApp() string {
	if o == nil || IsNil(o.ConnectedApp.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectedApp.Get()
}

// GetConnectedAppOk returns a tuple with the ConnectedApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpConnectionModel) GetConnectedAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedApp.Get(), o.ConnectedApp.IsSet()
}

// HasConnectedApp returns a boolean if a field has been set.
func (o *IdpConnectionModel) HasConnectedApp() bool {
	if o != nil && o.ConnectedApp.IsSet() {
		return true
	}

	return false
}

// SetConnectedApp gets a reference to the given NullableString and assigns it to the ConnectedApp field.
func (o *IdpConnectionModel) SetConnectedApp(v string) {
	o.ConnectedApp.Set(&v)
}
// SetConnectedAppNil sets the value for ConnectedApp to be an explicit nil
func (o *IdpConnectionModel) SetConnectedAppNil() {
	o.ConnectedApp.Set(nil)
}

// UnsetConnectedApp ensures that no value is present for ConnectedApp, not even an explicit nil
func (o *IdpConnectionModel) UnsetConnectedApp() {
	o.ConnectedApp.Unset()
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *IdpConnectionModel) GetConnectionStatus() bool {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret bool
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnectionModel) GetConnectionStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *IdpConnectionModel) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given bool and assigns it to the ConnectionStatus field.
func (o *IdpConnectionModel) SetConnectionStatus(v bool) {
	o.ConnectionStatus = &v
}

// GetIdpProperties returns the IdpProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpConnectionModel) GetIdpProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.IdpProperties
}

// GetIdpPropertiesOk returns a tuple with the IdpProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpConnectionModel) GetIdpPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.IdpProperties) {
		return nil, false
	}
	return &o.IdpProperties, true
}

// HasIdpProperties returns a boolean if a field has been set.
func (o *IdpConnectionModel) HasIdpProperties() bool {
	if o != nil && IsNil(o.IdpProperties) {
		return true
	}

	return false
}

// SetIdpProperties gets a reference to the given map[string]string and assigns it to the IdpProperties field.
func (o *IdpConnectionModel) SetIdpProperties(v map[string]string) {
	o.IdpProperties = v
}

func (o IdpConnectionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpConnectionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IdpType.IsSet() {
		toSerialize["idpType"] = o.IdpType.Get()
	}
	if o.ConnectedApp.IsSet() {
		toSerialize["connectedApp"] = o.ConnectedApp.Get()
	}
	if !IsNil(o.ConnectionStatus) {
		toSerialize["connectionStatus"] = o.ConnectionStatus
	}
	if o.IdpProperties != nil {
		toSerialize["idpProperties"] = o.IdpProperties
	}
	return toSerialize, nil
}

type NullableIdpConnectionModel struct {
	value *IdpConnectionModel
	isSet bool
}

func (v NullableIdpConnectionModel) Get() *IdpConnectionModel {
	return v.value
}

func (v *NullableIdpConnectionModel) Set(val *IdpConnectionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpConnectionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpConnectionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpConnectionModel(val *IdpConnectionModel) *NullableIdpConnectionModel {
	return &NullableIdpConnectionModel{value: val, isSet: true}
}

func (v NullableIdpConnectionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpConnectionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


