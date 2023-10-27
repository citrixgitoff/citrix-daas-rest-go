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

// checks if the VersionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionDetail{}

// VersionDetail Detailed VDA version information.
type VersionDetail struct {
	// XA or XD, no other values.
	Product NullableString `json:"product,omitempty"`
	// One of the values in XaVersions or XdVersions.
	Version NullableString `json:"version,omitempty"`
	// MS or SS, no other values.
	Edition NullableString `json:"edition,omitempty"`
	// Display name.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewVersionDetail instantiates a new VersionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionDetail() *VersionDetail {
	this := VersionDetail{}
	return &this
}

// NewVersionDetailWithDefaults instantiates a new VersionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionDetailWithDefaults() *VersionDetail {
	this := VersionDetail{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionDetail) GetProduct() string {
	if o == nil || IsNil(o.Product.Get()) {
		var ret string
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionDetail) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *VersionDetail) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableString and assigns it to the Product field.
func (o *VersionDetail) SetProduct(v string) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *VersionDetail) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *VersionDetail) UnsetProduct() {
	o.Product.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionDetail) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionDetail) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionDetail) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *VersionDetail) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *VersionDetail) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *VersionDetail) UnsetVersion() {
	o.Version.Unset()
}

// GetEdition returns the Edition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionDetail) GetEdition() string {
	if o == nil || IsNil(o.Edition.Get()) {
		var ret string
		return ret
	}
	return *o.Edition.Get()
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionDetail) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Edition.Get(), o.Edition.IsSet()
}

// HasEdition returns a boolean if a field has been set.
func (o *VersionDetail) HasEdition() bool {
	if o != nil && o.Edition.IsSet() {
		return true
	}

	return false
}

// SetEdition gets a reference to the given NullableString and assigns it to the Edition field.
func (o *VersionDetail) SetEdition(v string) {
	o.Edition.Set(&v)
}
// SetEditionNil sets the value for Edition to be an explicit nil
func (o *VersionDetail) SetEditionNil() {
	o.Edition.Set(nil)
}

// UnsetEdition ensures that no value is present for Edition, not even an explicit nil
func (o *VersionDetail) UnsetEdition() {
	o.Edition.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionDetail) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionDetail) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *VersionDetail) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *VersionDetail) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *VersionDetail) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *VersionDetail) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o VersionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Edition.IsSet() {
		toSerialize["edition"] = o.Edition.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return toSerialize, nil
}

type NullableVersionDetail struct {
	value *VersionDetail
	isSet bool
}

func (v NullableVersionDetail) Get() *VersionDetail {
	return v.value
}

func (v *NullableVersionDetail) Set(val *VersionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionDetail(val *VersionDetail) *NullableVersionDetail {
	return &NullableVersionDetail{value: val, isSet: true}
}

func (v NullableVersionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


