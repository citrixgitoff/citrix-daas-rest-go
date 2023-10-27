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

// checks if the EditOperationRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditOperationRequestModel{}

// EditOperationRequestModel Model for editing a high level log operation.
type EditOperationRequestModel struct {
	// The new labels of the operation.
	Labels []string `json:"Labels,omitempty"`
}

// NewEditOperationRequestModel instantiates a new EditOperationRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditOperationRequestModel() *EditOperationRequestModel {
	this := EditOperationRequestModel{}
	return &this
}

// NewEditOperationRequestModelWithDefaults instantiates a new EditOperationRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditOperationRequestModelWithDefaults() *EditOperationRequestModel {
	this := EditOperationRequestModel{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditOperationRequestModel) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditOperationRequestModel) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EditOperationRequestModel) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *EditOperationRequestModel) SetLabels(v []string) {
	o.Labels = v
}

func (o EditOperationRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditOperationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableEditOperationRequestModel struct {
	value *EditOperationRequestModel
	isSet bool
}

func (v NullableEditOperationRequestModel) Get() *EditOperationRequestModel {
	return v.value
}

func (v *NullableEditOperationRequestModel) Set(val *EditOperationRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditOperationRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditOperationRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditOperationRequestModel(val *EditOperationRequestModel) *NullableEditOperationRequestModel {
	return &NullableEditOperationRequestModel{value: val, isSet: true}
}

func (v NullableEditOperationRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditOperationRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


