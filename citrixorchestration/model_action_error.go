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

// checks if the ActionError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionError{}

// ActionError Action Error class.
type ActionError struct {
	// The error message.
	Message NullableString `json:"Message,omitempty"`
	// The stack trace.
	StackTrace NullableString `json:"StackTrace,omitempty"`
	// The error data.
	ErrorData []NameValueStringPairModel `json:"ErrorData,omitempty"`
	InnerError *ActionError `json:"InnerError,omitempty"`
}

// NewActionError instantiates a new ActionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionError() *ActionError {
	this := ActionError{}
	return &this
}

// NewActionErrorWithDefaults instantiates a new ActionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionErrorWithDefaults() *ActionError {
	this := ActionError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionError) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ActionError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ActionError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ActionError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ActionError) UnsetMessage() {
	o.Message.Unset()
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionError) GetStackTrace() string {
	if o == nil || IsNil(o.StackTrace.Get()) {
		var ret string
		return ret
	}
	return *o.StackTrace.Get()
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionError) GetStackTraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StackTrace.Get(), o.StackTrace.IsSet()
}

// HasStackTrace returns a boolean if a field has been set.
func (o *ActionError) HasStackTrace() bool {
	if o != nil && o.StackTrace.IsSet() {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given NullableString and assigns it to the StackTrace field.
func (o *ActionError) SetStackTrace(v string) {
	o.StackTrace.Set(&v)
}
// SetStackTraceNil sets the value for StackTrace to be an explicit nil
func (o *ActionError) SetStackTraceNil() {
	o.StackTrace.Set(nil)
}

// UnsetStackTrace ensures that no value is present for StackTrace, not even an explicit nil
func (o *ActionError) UnsetStackTrace() {
	o.StackTrace.Unset()
}

// GetErrorData returns the ErrorData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionError) GetErrorData() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.ErrorData
}

// GetErrorDataOk returns a tuple with the ErrorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionError) GetErrorDataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.ErrorData) {
		return nil, false
	}
	return o.ErrorData, true
}

// HasErrorData returns a boolean if a field has been set.
func (o *ActionError) HasErrorData() bool {
	if o != nil && IsNil(o.ErrorData) {
		return true
	}

	return false
}

// SetErrorData gets a reference to the given []NameValueStringPairModel and assigns it to the ErrorData field.
func (o *ActionError) SetErrorData(v []NameValueStringPairModel) {
	o.ErrorData = v
}

// GetInnerError returns the InnerError field value if set, zero value otherwise.
func (o *ActionError) GetInnerError() ActionError {
	if o == nil || IsNil(o.InnerError) {
		var ret ActionError
		return ret
	}
	return *o.InnerError
}

// GetInnerErrorOk returns a tuple with the InnerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionError) GetInnerErrorOk() (*ActionError, bool) {
	if o == nil || IsNil(o.InnerError) {
		return nil, false
	}
	return o.InnerError, true
}

// HasInnerError returns a boolean if a field has been set.
func (o *ActionError) HasInnerError() bool {
	if o != nil && !IsNil(o.InnerError) {
		return true
	}

	return false
}

// SetInnerError gets a reference to the given ActionError and assigns it to the InnerError field.
func (o *ActionError) SetInnerError(v ActionError) {
	o.InnerError = &v
}

func (o ActionError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	if o.StackTrace.IsSet() {
		toSerialize["StackTrace"] = o.StackTrace.Get()
	}
	if o.ErrorData != nil {
		toSerialize["ErrorData"] = o.ErrorData
	}
	if !IsNil(o.InnerError) {
		toSerialize["InnerError"] = o.InnerError
	}
	return toSerialize, nil
}

type NullableActionError struct {
	value *ActionError
	isSet bool
}

func (v NullableActionError) Get() *ActionError {
	return v.value
}

func (v *NullableActionError) Set(val *ActionError) {
	v.value = val
	v.isSet = true
}

func (v NullableActionError) IsSet() bool {
	return v.isSet
}

func (v *NullableActionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionError(val *ActionError) *NullableActionError {
	return &NullableActionError{value: val, isSet: true}
}

func (v NullableActionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


