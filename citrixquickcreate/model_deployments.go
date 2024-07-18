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

// checks if the Deployments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deployments{}

// Deployments Enumerable of Deployment
type Deployments struct {
	// Enumerable of Deployment
	Items []AwsEdcDeployment `json:"items,omitempty"`
}

// NewDeployments instantiates a new Deployments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployments() *Deployments {
	this := Deployments{}
	return &this
}

// NewDeploymentsWithDefaults instantiates a new Deployments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsWithDefaults() *Deployments {
	this := Deployments{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployments) GetItems() []AwsEdcDeployment {
	if o == nil {
		var ret []AwsEdcDeployment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployments) GetItemsOk() ([]AwsEdcDeployment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Deployments) HasItems() bool {
	if o != nil && IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AwsEdcDeployment and assigns it to the Items field.
func (o *Deployments) SetItems(v []AwsEdcDeployment) {
	o.Items = v
}

func (o Deployments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deployments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableDeployments struct {
	value *Deployments
	isSet bool
}

func (v NullableDeployments) Get() *Deployments {
	return v.value
}

func (v *NullableDeployments) Set(val *Deployments) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployments) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployments(val *Deployments) *NullableDeployments {
	return &NullableDeployments{value: val, isSet: true}
}

func (v NullableDeployments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


