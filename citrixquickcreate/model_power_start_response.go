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

// checks if the PowerStartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerStartResponse{}

// PowerStartResponse Response of power operation
type PowerStartResponse struct {
	Code *PowerResponseCode `json:"code,omitempty"`
	// Description about the error
	ErrorDescription NullableString `json:"errorDescription,omitempty"`
}

// NewPowerStartResponse instantiates a new PowerStartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerStartResponse() *PowerStartResponse {
	this := PowerStartResponse{}
	return &this
}

// NewPowerStartResponseWithDefaults instantiates a new PowerStartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerStartResponseWithDefaults() *PowerStartResponse {
	this := PowerStartResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PowerStartResponse) GetCode() PowerResponseCode {
	if o == nil || IsNil(o.Code) {
		var ret PowerResponseCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerStartResponse) GetCodeOk() (*PowerResponseCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PowerStartResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given PowerResponseCode and assigns it to the Code field.
func (o *PowerStartResponse) SetCode(v PowerResponseCode) {
	o.Code = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerStartResponse) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDescription.Get()
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerStartResponse) GetErrorDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDescription.Get(), o.ErrorDescription.IsSet()
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *PowerStartResponse) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription.IsSet() {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given NullableString and assigns it to the ErrorDescription field.
func (o *PowerStartResponse) SetErrorDescription(v string) {
	o.ErrorDescription.Set(&v)
}
// SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil
func (o *PowerStartResponse) SetErrorDescriptionNil() {
	o.ErrorDescription.Set(nil)
}

// UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
func (o *PowerStartResponse) UnsetErrorDescription() {
	o.ErrorDescription.Unset()
}

func (o PowerStartResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerStartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.ErrorDescription.IsSet() {
		toSerialize["errorDescription"] = o.ErrorDescription.Get()
	}
	return toSerialize, nil
}

type NullablePowerStartResponse struct {
	value *PowerStartResponse
	isSet bool
}

func (v NullablePowerStartResponse) Get() *PowerStartResponse {
	return v.value
}

func (v *NullablePowerStartResponse) Set(val *PowerStartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerStartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerStartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerStartResponse(val *PowerStartResponse) *NullablePowerStartResponse {
	return &NullablePowerStartResponse{value: val, isSet: true}
}

func (v NullablePowerStartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerStartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


