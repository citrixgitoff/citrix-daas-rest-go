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

// checks if the SendMessageRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendMessageRequestModel{}

// SendMessageRequestModel Request to send a message to a session.
type SendMessageRequestModel struct {
	Style *MessageStyle `json:"Style,omitempty"`
	// Text to display in the message box title bar.
	Title NullableString `json:"Title,omitempty"`
	// The message to display.
	Text NullableString `json:"Text,omitempty"`
}

// NewSendMessageRequestModel instantiates a new SendMessageRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessageRequestModel() *SendMessageRequestModel {
	this := SendMessageRequestModel{}
	return &this
}

// NewSendMessageRequestModelWithDefaults instantiates a new SendMessageRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessageRequestModelWithDefaults() *SendMessageRequestModel {
	this := SendMessageRequestModel{}
	return &this
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *SendMessageRequestModel) GetStyle() MessageStyle {
	if o == nil || IsNil(o.Style) {
		var ret MessageStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequestModel) GetStyleOk() (*MessageStyle, bool) {
	if o == nil || IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *SendMessageRequestModel) HasStyle() bool {
	if o != nil && !IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given MessageStyle and assigns it to the Style field.
func (o *SendMessageRequestModel) SetStyle(v MessageStyle) {
	o.Style = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendMessageRequestModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendMessageRequestModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SendMessageRequestModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SendMessageRequestModel) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SendMessageRequestModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SendMessageRequestModel) UnsetTitle() {
	o.Title.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendMessageRequestModel) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendMessageRequestModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *SendMessageRequestModel) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *SendMessageRequestModel) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *SendMessageRequestModel) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *SendMessageRequestModel) UnsetText() {
	o.Text.Unset()
}

func (o SendMessageRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendMessageRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Style) {
		toSerialize["Style"] = o.Style
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.Text.IsSet() {
		toSerialize["Text"] = o.Text.Get()
	}
	return toSerialize, nil
}

type NullableSendMessageRequestModel struct {
	value *SendMessageRequestModel
	isSet bool
}

func (v NullableSendMessageRequestModel) Get() *SendMessageRequestModel {
	return v.value
}

func (v *NullableSendMessageRequestModel) Set(val *SendMessageRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessageRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessageRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessageRequestModel(val *SendMessageRequestModel) *NullableSendMessageRequestModel {
	return &NullableSendMessageRequestModel{value: val, isSet: true}
}

func (v NullableSendMessageRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessageRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


