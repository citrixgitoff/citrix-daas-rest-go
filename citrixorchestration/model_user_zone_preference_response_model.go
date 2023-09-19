/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the UserZonePreferenceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserZonePreferenceResponseModel{}

// UserZonePreferenceResponseModel Data model for user zone preference.
type UserZonePreferenceResponseModel struct {
	// The full name of the user or user group.
	FullName NullableString `json:"FullName,omitempty"`
	// Name of the user or user group.
	Name NullableString `json:"Name,omitempty"`
	// Sid of the user or user group.
	Sid NullableString `json:"Sid,omitempty"`
	// UPN of the user or user group.
	Upn NullableString `json:"Upn,omitempty"`
	Zone *RefResponseModel `json:"Zone,omitempty"`
}

// NewUserZonePreferenceResponseModel instantiates a new UserZonePreferenceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserZonePreferenceResponseModel() *UserZonePreferenceResponseModel {
	this := UserZonePreferenceResponseModel{}
	return &this
}

// NewUserZonePreferenceResponseModelWithDefaults instantiates a new UserZonePreferenceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserZonePreferenceResponseModelWithDefaults() *UserZonePreferenceResponseModel {
	this := UserZonePreferenceResponseModel{}
	return &this
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserZonePreferenceResponseModel) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserZonePreferenceResponseModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *UserZonePreferenceResponseModel) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *UserZonePreferenceResponseModel) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *UserZonePreferenceResponseModel) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *UserZonePreferenceResponseModel) UnsetFullName() {
	o.FullName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserZonePreferenceResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserZonePreferenceResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserZonePreferenceResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserZonePreferenceResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserZonePreferenceResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserZonePreferenceResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserZonePreferenceResponseModel) GetSid() string {
	if o == nil || IsNil(o.Sid.Get()) {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserZonePreferenceResponseModel) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *UserZonePreferenceResponseModel) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *UserZonePreferenceResponseModel) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *UserZonePreferenceResponseModel) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *UserZonePreferenceResponseModel) UnsetSid() {
	o.Sid.Unset()
}

// GetUpn returns the Upn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserZonePreferenceResponseModel) GetUpn() string {
	if o == nil || IsNil(o.Upn.Get()) {
		var ret string
		return ret
	}
	return *o.Upn.Get()
}

// GetUpnOk returns a tuple with the Upn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserZonePreferenceResponseModel) GetUpnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Upn.Get(), o.Upn.IsSet()
}

// HasUpn returns a boolean if a field has been set.
func (o *UserZonePreferenceResponseModel) HasUpn() bool {
	if o != nil && o.Upn.IsSet() {
		return true
	}

	return false
}

// SetUpn gets a reference to the given NullableString and assigns it to the Upn field.
func (o *UserZonePreferenceResponseModel) SetUpn(v string) {
	o.Upn.Set(&v)
}
// SetUpnNil sets the value for Upn to be an explicit nil
func (o *UserZonePreferenceResponseModel) SetUpnNil() {
	o.Upn.Set(nil)
}

// UnsetUpn ensures that no value is present for Upn, not even an explicit nil
func (o *UserZonePreferenceResponseModel) UnsetUpn() {
	o.Upn.Unset()
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *UserZonePreferenceResponseModel) GetZone() RefResponseModel {
	if o == nil || IsNil(o.Zone) {
		var ret RefResponseModel
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserZonePreferenceResponseModel) GetZoneOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *UserZonePreferenceResponseModel) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given RefResponseModel and assigns it to the Zone field.
func (o *UserZonePreferenceResponseModel) SetZone(v RefResponseModel) {
	o.Zone = &v
}

func (o UserZonePreferenceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserZonePreferenceResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FullName.IsSet() {
		toSerialize["FullName"] = o.FullName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["Sid"] = o.Sid.Get()
	}
	if o.Upn.IsSet() {
		toSerialize["Upn"] = o.Upn.Get()
	}
	if !IsNil(o.Zone) {
		toSerialize["Zone"] = o.Zone
	}
	return toSerialize, nil
}

type NullableUserZonePreferenceResponseModel struct {
	value *UserZonePreferenceResponseModel
	isSet bool
}

func (v NullableUserZonePreferenceResponseModel) Get() *UserZonePreferenceResponseModel {
	return v.value
}

func (v *NullableUserZonePreferenceResponseModel) Set(val *UserZonePreferenceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserZonePreferenceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserZonePreferenceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserZonePreferenceResponseModel(val *UserZonePreferenceResponseModel) *NullableUserZonePreferenceResponseModel {
	return &NullableUserZonePreferenceResponseModel{value: val, isSet: true}
}

func (v NullableUserZonePreferenceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserZonePreferenceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


