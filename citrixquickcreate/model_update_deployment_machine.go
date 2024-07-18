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

// checks if the UpdateDeploymentMachine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeploymentMachine{}

// UpdateDeploymentMachine Adds machines to deployment
type UpdateDeploymentMachine struct {
	AccountType AccountType `json:"accountType"`
}

// NewUpdateDeploymentMachine instantiates a new UpdateDeploymentMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeploymentMachine(accountType AccountType) *UpdateDeploymentMachine {
	this := UpdateDeploymentMachine{}
	this.AccountType = accountType
	return &this
}

// NewUpdateDeploymentMachineWithDefaults instantiates a new UpdateDeploymentMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeploymentMachineWithDefaults() *UpdateDeploymentMachine {
	this := UpdateDeploymentMachine{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *UpdateDeploymentMachine) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *UpdateDeploymentMachine) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *UpdateDeploymentMachine) SetAccountType(v AccountType) {
	o.AccountType = v
}

func (o UpdateDeploymentMachine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeploymentMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	return toSerialize, nil
}

type NullableUpdateDeploymentMachine struct {
	value *UpdateDeploymentMachine
	isSet bool
}

func (v NullableUpdateDeploymentMachine) Get() *UpdateDeploymentMachine {
	return v.value
}

func (v *NullableUpdateDeploymentMachine) Set(val *UpdateDeploymentMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeploymentMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeploymentMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeploymentMachine(val *UpdateDeploymentMachine) *NullableUpdateDeploymentMachine {
	return &NullableUpdateDeploymentMachine{value: val, isSet: true}
}

func (v NullableUpdateDeploymentMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeploymentMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


