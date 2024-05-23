/*
Global App Config Admin

Describes API used by Global App Config Admin Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package globalappconfiguration

import (
	"encoding/json"
)

// checks if the CitrixErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixErrorModel{}

// CitrixErrorModel struct for CitrixErrorModel
type CitrixErrorModel struct {
	Detail *string `json:"detail,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCitrixErrorModel instantiates a new CitrixErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixErrorModel() *CitrixErrorModel {
	this := CitrixErrorModel{}
	return &this
}

// NewCitrixErrorModelWithDefaults instantiates a new CitrixErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixErrorModelWithDefaults() *CitrixErrorModel {
	this := CitrixErrorModel{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *CitrixErrorModel) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixErrorModel) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *CitrixErrorModel) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *CitrixErrorModel) SetDetail(v string) {
	o.Detail = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CitrixErrorModel) GetParameters() []Parameter {
	if o == nil || IsNil(o.Parameters) {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixErrorModel) GetParametersOk() ([]Parameter, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CitrixErrorModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *CitrixErrorModel) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CitrixErrorModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixErrorModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CitrixErrorModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CitrixErrorModel) SetType(v string) {
	o.Type = &v
}

func (o CitrixErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCitrixErrorModel struct {
	value *CitrixErrorModel
	isSet bool
}

func (v NullableCitrixErrorModel) Get() *CitrixErrorModel {
	return v.value
}

func (v *NullableCitrixErrorModel) Set(val *CitrixErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixErrorModel(val *CitrixErrorModel) *NullableCitrixErrorModel {
	return &NullableCitrixErrorModel{value: val, isSet: true}
}

func (v NullableCitrixErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


