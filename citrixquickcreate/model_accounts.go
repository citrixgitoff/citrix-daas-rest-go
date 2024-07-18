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

// checks if the Accounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Accounts{}

// Accounts Enumerable of Account
type Accounts struct {
	// Enumerable of Account
	Items []AwsEdcAccount `json:"items,omitempty"`
}

// NewAccounts instantiates a new Accounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccounts() *Accounts {
	this := Accounts{}
	return &this
}

// NewAccountsWithDefaults instantiates a new Accounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsWithDefaults() *Accounts {
	this := Accounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Accounts) GetItems() []AwsEdcAccount {
	if o == nil {
		var ret []AwsEdcAccount
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Accounts) GetItemsOk() ([]AwsEdcAccount, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Accounts) HasItems() bool {
	if o != nil && IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AwsEdcAccount and assigns it to the Items field.
func (o *Accounts) SetItems(v []AwsEdcAccount) {
	o.Items = v
}

func (o Accounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Accounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAccounts struct {
	value *Accounts
	isSet bool
}

func (v NullableAccounts) Get() *Accounts {
	return v.value
}

func (v *NullableAccounts) Set(val *Accounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccounts(val *Accounts) *NullableAccounts {
	return &NullableAccounts{value: val, isSet: true}
}

func (v NullableAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


