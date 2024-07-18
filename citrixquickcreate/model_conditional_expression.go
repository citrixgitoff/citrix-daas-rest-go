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

// checks if the ConditionalExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionalExpression{}

// ConditionalExpression struct for ConditionalExpression
type ConditionalExpression struct {
	StringEquals NullableExpression `json:"stringEquals,omitempty"`
}

// NewConditionalExpression instantiates a new ConditionalExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalExpression() *ConditionalExpression {
	this := ConditionalExpression{}
	return &this
}

// NewConditionalExpressionWithDefaults instantiates a new ConditionalExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalExpressionWithDefaults() *ConditionalExpression {
	this := ConditionalExpression{}
	return &this
}

// GetStringEquals returns the StringEquals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConditionalExpression) GetStringEquals() Expression {
	if o == nil || IsNil(o.StringEquals.Get()) {
		var ret Expression
		return ret
	}
	return *o.StringEquals.Get()
}

// GetStringEqualsOk returns a tuple with the StringEquals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConditionalExpression) GetStringEqualsOk() (*Expression, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringEquals.Get(), o.StringEquals.IsSet()
}

// HasStringEquals returns a boolean if a field has been set.
func (o *ConditionalExpression) HasStringEquals() bool {
	if o != nil && o.StringEquals.IsSet() {
		return true
	}

	return false
}

// SetStringEquals gets a reference to the given NullableExpression and assigns it to the StringEquals field.
func (o *ConditionalExpression) SetStringEquals(v Expression) {
	o.StringEquals.Set(&v)
}
// SetStringEqualsNil sets the value for StringEquals to be an explicit nil
func (o *ConditionalExpression) SetStringEqualsNil() {
	o.StringEquals.Set(nil)
}

// UnsetStringEquals ensures that no value is present for StringEquals, not even an explicit nil
func (o *ConditionalExpression) UnsetStringEquals() {
	o.StringEquals.Unset()
}

func (o ConditionalExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StringEquals.IsSet() {
		toSerialize["stringEquals"] = o.StringEquals.Get()
	}
	return toSerialize, nil
}

type NullableConditionalExpression struct {
	value *ConditionalExpression
	isSet bool
}

func (v NullableConditionalExpression) Get() *ConditionalExpression {
	return v.value
}

func (v *NullableConditionalExpression) Set(val *ConditionalExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalExpression(val *ConditionalExpression) *NullableConditionalExpression {
	return &NullableConditionalExpression{value: val, isSet: true}
}

func (v NullableConditionalExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


