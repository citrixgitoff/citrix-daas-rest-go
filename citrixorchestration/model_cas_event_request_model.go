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

// checks if the CasEventRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CasEventRequestModel{}

// CasEventRequestModel Request model for sending Cas event.
type CasEventRequestModel struct {
	// Cas Event Type Header
	EventType string `json:"EventType"`
	// Information that need to send to cas, which will shown in cas as payloads.
	Payloads []NameValueStringPairModel `json:"Payloads,omitempty"`
}

// NewCasEventRequestModel instantiates a new CasEventRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCasEventRequestModel(eventType string) *CasEventRequestModel {
	this := CasEventRequestModel{}
	this.EventType = eventType
	return &this
}

// NewCasEventRequestModelWithDefaults instantiates a new CasEventRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCasEventRequestModelWithDefaults() *CasEventRequestModel {
	this := CasEventRequestModel{}
	return &this
}

// GetEventType returns the EventType field value
func (o *CasEventRequestModel) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *CasEventRequestModel) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *CasEventRequestModel) SetEventType(v string) {
	o.EventType = v
}

// GetPayloads returns the Payloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CasEventRequestModel) GetPayloads() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Payloads
}

// GetPayloadsOk returns a tuple with the Payloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CasEventRequestModel) GetPayloadsOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Payloads) {
		return nil, false
	}
	return o.Payloads, true
}

// HasPayloads returns a boolean if a field has been set.
func (o *CasEventRequestModel) HasPayloads() bool {
	if o != nil && IsNil(o.Payloads) {
		return true
	}

	return false
}

// SetPayloads gets a reference to the given []NameValueStringPairModel and assigns it to the Payloads field.
func (o *CasEventRequestModel) SetPayloads(v []NameValueStringPairModel) {
	o.Payloads = v
}

func (o CasEventRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CasEventRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["EventType"] = o.EventType
	if o.Payloads != nil {
		toSerialize["Payloads"] = o.Payloads
	}
	return toSerialize, nil
}

type NullableCasEventRequestModel struct {
	value *CasEventRequestModel
	isSet bool
}

func (v NullableCasEventRequestModel) Get() *CasEventRequestModel {
	return v.value
}

func (v *NullableCasEventRequestModel) Set(val *CasEventRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCasEventRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCasEventRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCasEventRequestModel(val *CasEventRequestModel) *NullableCasEventRequestModel {
	return &NullableCasEventRequestModel{value: val, isSet: true}
}

func (v NullableCasEventRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCasEventRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


