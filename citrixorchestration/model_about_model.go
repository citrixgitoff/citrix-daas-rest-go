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

// checks if the AboutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AboutModel{}

// AboutModel About model of current Orchestration instance.             
type AboutModel struct {
	// The commit hash when building.             
	Commit *string `json:"Commit,omitempty"`
	// Indicates if the  current environment is cloud.
	IsCloud *bool `json:"IsCloud,omitempty"`
}

// NewAboutModel instantiates a new AboutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAboutModel() *AboutModel {
	this := AboutModel{}
	return &this
}

// NewAboutModelWithDefaults instantiates a new AboutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAboutModelWithDefaults() *AboutModel {
	this := AboutModel{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *AboutModel) GetCommit() string {
	if o == nil || IsNil(o.Commit) {
		var ret string
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutModel) GetCommitOk() (*string, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *AboutModel) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given string and assigns it to the Commit field.
func (o *AboutModel) SetCommit(v string) {
	o.Commit = &v
}

// GetIsCloud returns the IsCloud field value if set, zero value otherwise.
func (o *AboutModel) GetIsCloud() bool {
	if o == nil || IsNil(o.IsCloud) {
		var ret bool
		return ret
	}
	return *o.IsCloud
}

// GetIsCloudOk returns a tuple with the IsCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutModel) GetIsCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloud) {
		return nil, false
	}
	return o.IsCloud, true
}

// HasIsCloud returns a boolean if a field has been set.
func (o *AboutModel) HasIsCloud() bool {
	if o != nil && !IsNil(o.IsCloud) {
		return true
	}

	return false
}

// SetIsCloud gets a reference to the given bool and assigns it to the IsCloud field.
func (o *AboutModel) SetIsCloud(v bool) {
	o.IsCloud = &v
}

func (o AboutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AboutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commit) {
		toSerialize["Commit"] = o.Commit
	}
	if !IsNil(o.IsCloud) {
		toSerialize["IsCloud"] = o.IsCloud
	}
	return toSerialize, nil
}

type NullableAboutModel struct {
	value *AboutModel
	isSet bool
}

func (v NullableAboutModel) Get() *AboutModel {
	return v.value
}

func (v *NullableAboutModel) Set(val *AboutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAboutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAboutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAboutModel(val *AboutModel) *NullableAboutModel {
	return &NullableAboutModel{value: val, isSet: true}
}

func (v NullableAboutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAboutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


