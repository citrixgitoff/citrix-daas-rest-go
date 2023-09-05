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

// checks if the SessionRecordingStatusResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionRecordingStatusResponseModel{}

// SessionRecordingStatusResponseModel Session Recording Status model of current Orchestration instance.             
type SessionRecordingStatusResponseModel struct {
	// Indicates if current session is being recorded.
	IsSessionBeingRecorded *bool `json:"IsSessionBeingRecorded,omitempty"`
}

// NewSessionRecordingStatusResponseModel instantiates a new SessionRecordingStatusResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRecordingStatusResponseModel() *SessionRecordingStatusResponseModel {
	this := SessionRecordingStatusResponseModel{}
	return &this
}

// NewSessionRecordingStatusResponseModelWithDefaults instantiates a new SessionRecordingStatusResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRecordingStatusResponseModelWithDefaults() *SessionRecordingStatusResponseModel {
	this := SessionRecordingStatusResponseModel{}
	return &this
}

// GetIsSessionBeingRecorded returns the IsSessionBeingRecorded field value if set, zero value otherwise.
func (o *SessionRecordingStatusResponseModel) GetIsSessionBeingRecorded() bool {
	if o == nil || IsNil(o.IsSessionBeingRecorded) {
		var ret bool
		return ret
	}
	return *o.IsSessionBeingRecorded
}

// GetIsSessionBeingRecordedOk returns a tuple with the IsSessionBeingRecorded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRecordingStatusResponseModel) GetIsSessionBeingRecordedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSessionBeingRecorded) {
		return nil, false
	}
	return o.IsSessionBeingRecorded, true
}

// HasIsSessionBeingRecorded returns a boolean if a field has been set.
func (o *SessionRecordingStatusResponseModel) HasIsSessionBeingRecorded() bool {
	if o != nil && !IsNil(o.IsSessionBeingRecorded) {
		return true
	}

	return false
}

// SetIsSessionBeingRecorded gets a reference to the given bool and assigns it to the IsSessionBeingRecorded field.
func (o *SessionRecordingStatusResponseModel) SetIsSessionBeingRecorded(v bool) {
	o.IsSessionBeingRecorded = &v
}

func (o SessionRecordingStatusResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionRecordingStatusResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSessionBeingRecorded) {
		toSerialize["IsSessionBeingRecorded"] = o.IsSessionBeingRecorded
	}
	return toSerialize, nil
}

type NullableSessionRecordingStatusResponseModel struct {
	value *SessionRecordingStatusResponseModel
	isSet bool
}

func (v NullableSessionRecordingStatusResponseModel) Get() *SessionRecordingStatusResponseModel {
	return v.value
}

func (v *NullableSessionRecordingStatusResponseModel) Set(val *SessionRecordingStatusResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRecordingStatusResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRecordingStatusResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRecordingStatusResponseModel(val *SessionRecordingStatusResponseModel) *NullableSessionRecordingStatusResponseModel {
	return &NullableSessionRecordingStatusResponseModel{value: val, isSet: true}
}

func (v NullableSessionRecordingStatusResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRecordingStatusResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


