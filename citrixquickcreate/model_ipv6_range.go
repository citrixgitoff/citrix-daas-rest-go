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

// checks if the Ipv6Range type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6Range{}

// Ipv6Range struct for Ipv6Range
type Ipv6Range struct {
	CidrIpv6 NullableString `json:"cidrIpv6,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewIpv6Range instantiates a new Ipv6Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Range() *Ipv6Range {
	this := Ipv6Range{}
	return &this
}

// NewIpv6RangeWithDefaults instantiates a new Ipv6Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6RangeWithDefaults() *Ipv6Range {
	this := Ipv6Range{}
	return &this
}

// GetCidrIpv6 returns the CidrIpv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ipv6Range) GetCidrIpv6() string {
	if o == nil || IsNil(o.CidrIpv6.Get()) {
		var ret string
		return ret
	}
	return *o.CidrIpv6.Get()
}

// GetCidrIpv6Ok returns a tuple with the CidrIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ipv6Range) GetCidrIpv6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CidrIpv6.Get(), o.CidrIpv6.IsSet()
}

// HasCidrIpv6 returns a boolean if a field has been set.
func (o *Ipv6Range) HasCidrIpv6() bool {
	if o != nil && o.CidrIpv6.IsSet() {
		return true
	}

	return false
}

// SetCidrIpv6 gets a reference to the given NullableString and assigns it to the CidrIpv6 field.
func (o *Ipv6Range) SetCidrIpv6(v string) {
	o.CidrIpv6.Set(&v)
}
// SetCidrIpv6Nil sets the value for CidrIpv6 to be an explicit nil
func (o *Ipv6Range) SetCidrIpv6Nil() {
	o.CidrIpv6.Set(nil)
}

// UnsetCidrIpv6 ensures that no value is present for CidrIpv6, not even an explicit nil
func (o *Ipv6Range) UnsetCidrIpv6() {
	o.CidrIpv6.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ipv6Range) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ipv6Range) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Ipv6Range) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Ipv6Range) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Ipv6Range) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Ipv6Range) UnsetDescription() {
	o.Description.Unset()
}

func (o Ipv6Range) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6Range) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CidrIpv6.IsSet() {
		toSerialize["cidrIpv6"] = o.CidrIpv6.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableIpv6Range struct {
	value *Ipv6Range
	isSet bool
}

func (v NullableIpv6Range) Get() *Ipv6Range {
	return v.value
}

func (v *NullableIpv6Range) Set(val *Ipv6Range) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Range) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Range) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Range(val *Ipv6Range) *NullableIpv6Range {
	return &NullableIpv6Range{value: val, isSet: true}
}

func (v NullableIpv6Range) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Range) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


