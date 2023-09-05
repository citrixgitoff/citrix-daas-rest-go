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

// checks if the RebootScheduleResponseModelDeliveryGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RebootScheduleResponseModelDeliveryGroup{}

// RebootScheduleResponseModelDeliveryGroup The delivery group to which the reboot schedule is associated.
type RebootScheduleResponseModelDeliveryGroup struct {
	// Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility
	Id *string `json:"Id,omitempty"`
	// DEPRECATED. Use Id.
	Uid *int32 `json:"Uid,omitempty"`
	// Name of the object.
	Name *string `json:"Name,omitempty"`
}

// NewRebootScheduleResponseModelDeliveryGroup instantiates a new RebootScheduleResponseModelDeliveryGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootScheduleResponseModelDeliveryGroup() *RebootScheduleResponseModelDeliveryGroup {
	this := RebootScheduleResponseModelDeliveryGroup{}
	return &this
}

// NewRebootScheduleResponseModelDeliveryGroupWithDefaults instantiates a new RebootScheduleResponseModelDeliveryGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootScheduleResponseModelDeliveryGroupWithDefaults() *RebootScheduleResponseModelDeliveryGroup {
	this := RebootScheduleResponseModelDeliveryGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RebootScheduleResponseModelDeliveryGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RebootScheduleResponseModelDeliveryGroup) SetId(v string) {
	o.Id = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *RebootScheduleResponseModelDeliveryGroup) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *RebootScheduleResponseModelDeliveryGroup) SetUid(v int32) {
	o.Uid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RebootScheduleResponseModelDeliveryGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RebootScheduleResponseModelDeliveryGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RebootScheduleResponseModelDeliveryGroup) SetName(v string) {
	o.Name = &v
}

func (o RebootScheduleResponseModelDeliveryGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RebootScheduleResponseModelDeliveryGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRebootScheduleResponseModelDeliveryGroup struct {
	value *RebootScheduleResponseModelDeliveryGroup
	isSet bool
}

func (v NullableRebootScheduleResponseModelDeliveryGroup) Get() *RebootScheduleResponseModelDeliveryGroup {
	return v.value
}

func (v *NullableRebootScheduleResponseModelDeliveryGroup) Set(val *RebootScheduleResponseModelDeliveryGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootScheduleResponseModelDeliveryGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootScheduleResponseModelDeliveryGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootScheduleResponseModelDeliveryGroup(val *RebootScheduleResponseModelDeliveryGroup) *NullableRebootScheduleResponseModelDeliveryGroup {
	return &NullableRebootScheduleResponseModelDeliveryGroup{value: val, isSet: true}
}

func (v NullableRebootScheduleResponseModelDeliveryGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootScheduleResponseModelDeliveryGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


