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

// checks if the DatabaseChangeScriptPropertiesModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseChangeScriptPropertiesModel{}

// DatabaseChangeScriptPropertiesModel Database change script properties model.
type DatabaseChangeScriptPropertiesModel struct {
	// Name of database
	DatabaseName string `json:"DatabaseName"`
	// Server address of database
	DatabaseServerAddress string `json:"DatabaseServerAddress"`
}

// NewDatabaseChangeScriptPropertiesModel instantiates a new DatabaseChangeScriptPropertiesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseChangeScriptPropertiesModel(databaseName string, databaseServerAddress string) *DatabaseChangeScriptPropertiesModel {
	this := DatabaseChangeScriptPropertiesModel{}
	this.DatabaseName = databaseName
	this.DatabaseServerAddress = databaseServerAddress
	return &this
}

// NewDatabaseChangeScriptPropertiesModelWithDefaults instantiates a new DatabaseChangeScriptPropertiesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseChangeScriptPropertiesModelWithDefaults() *DatabaseChangeScriptPropertiesModel {
	this := DatabaseChangeScriptPropertiesModel{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value
func (o *DatabaseChangeScriptPropertiesModel) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DatabaseChangeScriptPropertiesModel) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *DatabaseChangeScriptPropertiesModel) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetDatabaseServerAddress returns the DatabaseServerAddress field value
func (o *DatabaseChangeScriptPropertiesModel) GetDatabaseServerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseServerAddress
}

// GetDatabaseServerAddressOk returns a tuple with the DatabaseServerAddress field value
// and a boolean to check if the value has been set.
func (o *DatabaseChangeScriptPropertiesModel) GetDatabaseServerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseServerAddress, true
}

// SetDatabaseServerAddress sets field value
func (o *DatabaseChangeScriptPropertiesModel) SetDatabaseServerAddress(v string) {
	o.DatabaseServerAddress = v
}

func (o DatabaseChangeScriptPropertiesModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseChangeScriptPropertiesModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DatabaseName"] = o.DatabaseName
	toSerialize["DatabaseServerAddress"] = o.DatabaseServerAddress
	return toSerialize, nil
}

type NullableDatabaseChangeScriptPropertiesModel struct {
	value *DatabaseChangeScriptPropertiesModel
	isSet bool
}

func (v NullableDatabaseChangeScriptPropertiesModel) Get() *DatabaseChangeScriptPropertiesModel {
	return v.value
}

func (v *NullableDatabaseChangeScriptPropertiesModel) Set(val *DatabaseChangeScriptPropertiesModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseChangeScriptPropertiesModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseChangeScriptPropertiesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseChangeScriptPropertiesModel(val *DatabaseChangeScriptPropertiesModel) *NullableDatabaseChangeScriptPropertiesModel {
	return &NullableDatabaseChangeScriptPropertiesModel{value: val, isSet: true}
}

func (v NullableDatabaseChangeScriptPropertiesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseChangeScriptPropertiesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


