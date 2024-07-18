/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel{}

// CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel struct for CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel
type CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel struct {
	UserIds []string `json:"userIds,omitempty"`
	UcOids []string `json:"ucOids,omitempty"`
}

// NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel instantiates a new CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel() *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel {
	this := CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel{}
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel {
	this := CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) SetUserIds(v []string) {
	o.UserIds = v
}

// GetUcOids returns the UcOids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) GetUcOids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UcOids
}

// GetUcOidsOk returns a tuple with the UcOids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) GetUcOidsOk() ([]string, bool) {
	if o == nil || IsNil(o.UcOids) {
		return nil, false
	}
	return o.UcOids, true
}

// HasUcOids returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) HasUcOids() bool {
	if o != nil && !IsNil(o.UcOids) {
		return true
	}

	return false
}

// SetUcOids gets a reference to the given []string and assigns it to the UcOids field.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) SetUcOids(v []string) {
	o.UcOids = v
}

func (o CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if o.UcOids != nil {
		toSerialize["ucOids"] = o.UcOids
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel struct {
	value *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) Get() *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) Set(val *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel(val *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel {
	return &NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


