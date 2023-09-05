/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the ActionErrorInnerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionErrorInnerError{}

// ActionErrorInnerError The inner error.
type ActionErrorInnerError struct {
	// The error message.
	Message *string `json:"Message,omitempty"`
	// The stack trace.
	StackTrace *string `json:"StackTrace,omitempty"`
	// The error data.
	ErrorData []NameValueStringPairModel `json:"ErrorData,omitempty"`
	InnerError *ActionErrorInnerError `json:"InnerError,omitempty"`
}

// NewActionErrorInnerError instantiates a new ActionErrorInnerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionErrorInnerError() *ActionErrorInnerError {
	this := ActionErrorInnerError{}
	return &this
}

// NewActionErrorInnerErrorWithDefaults instantiates a new ActionErrorInnerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionErrorInnerErrorWithDefaults() *ActionErrorInnerError {
	this := ActionErrorInnerError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ActionErrorInnerError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionErrorInnerError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ActionErrorInnerError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ActionErrorInnerError) SetMessage(v string) {
	o.Message = &v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *ActionErrorInnerError) GetStackTrace() string {
	if o == nil || IsNil(o.StackTrace) {
		var ret string
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionErrorInnerError) GetStackTraceOk() (*string, bool) {
	if o == nil || IsNil(o.StackTrace) {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *ActionErrorInnerError) HasStackTrace() bool {
	if o != nil && !IsNil(o.StackTrace) {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given string and assigns it to the StackTrace field.
func (o *ActionErrorInnerError) SetStackTrace(v string) {
	o.StackTrace = &v
}

// GetErrorData returns the ErrorData field value if set, zero value otherwise.
func (o *ActionErrorInnerError) GetErrorData() []NameValueStringPairModel {
	if o == nil || IsNil(o.ErrorData) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.ErrorData
}

// GetErrorDataOk returns a tuple with the ErrorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionErrorInnerError) GetErrorDataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.ErrorData) {
		return nil, false
	}
	return o.ErrorData, true
}

// HasErrorData returns a boolean if a field has been set.
func (o *ActionErrorInnerError) HasErrorData() bool {
	if o != nil && !IsNil(o.ErrorData) {
		return true
	}

	return false
}

// SetErrorData gets a reference to the given []NameValueStringPairModel and assigns it to the ErrorData field.
func (o *ActionErrorInnerError) SetErrorData(v []NameValueStringPairModel) {
	o.ErrorData = v
}

// GetInnerError returns the InnerError field value if set, zero value otherwise.
func (o *ActionErrorInnerError) GetInnerError() ActionErrorInnerError {
	if o == nil || IsNil(o.InnerError) {
		var ret ActionErrorInnerError
		return ret
	}
	return *o.InnerError
}

// GetInnerErrorOk returns a tuple with the InnerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionErrorInnerError) GetInnerErrorOk() (*ActionErrorInnerError, bool) {
	if o == nil || IsNil(o.InnerError) {
		return nil, false
	}
	return o.InnerError, true
}

// HasInnerError returns a boolean if a field has been set.
func (o *ActionErrorInnerError) HasInnerError() bool {
	if o != nil && !IsNil(o.InnerError) {
		return true
	}

	return false
}

// SetInnerError gets a reference to the given ActionErrorInnerError and assigns it to the InnerError field.
func (o *ActionErrorInnerError) SetInnerError(v ActionErrorInnerError) {
	o.InnerError = &v
}

func (o ActionErrorInnerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionErrorInnerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.StackTrace) {
		toSerialize["StackTrace"] = o.StackTrace
	}
	if !IsNil(o.ErrorData) {
		toSerialize["ErrorData"] = o.ErrorData
	}
	if !IsNil(o.InnerError) {
		toSerialize["InnerError"] = o.InnerError
	}
	return toSerialize, nil
}

type NullableActionErrorInnerError struct {
	value *ActionErrorInnerError
	isSet bool
}

func (v NullableActionErrorInnerError) Get() *ActionErrorInnerError {
	return v.value
}

func (v *NullableActionErrorInnerError) Set(val *ActionErrorInnerError) {
	v.value = val
	v.isSet = true
}

func (v NullableActionErrorInnerError) IsSet() bool {
	return v.isSet
}

func (v *NullableActionErrorInnerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionErrorInnerError(val *ActionErrorInnerError) *NullableActionErrorInnerError {
	return &NullableActionErrorInnerError{value: val, isSet: true}
}

func (v NullableActionErrorInnerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionErrorInnerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


