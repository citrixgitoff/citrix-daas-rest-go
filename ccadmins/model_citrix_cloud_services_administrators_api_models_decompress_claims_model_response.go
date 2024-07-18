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

// checks if the CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse{}

// CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse struct for CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse
type CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse struct {
	ClaimsModel *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel `json:"claimsModel,omitempty"`
}

// NewCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse instantiates a new CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse() *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse {
	this := CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse{}
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponseWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponseWithDefaults() *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse {
	this := CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse{}
	return &this
}

// GetClaimsModel returns the ClaimsModel field value if set, zero value otherwise.
func (o *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) GetClaimsModel() CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel {
	if o == nil || IsNil(o.ClaimsModel) {
		var ret CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel
		return ret
	}
	return *o.ClaimsModel
}

// GetClaimsModelOk returns a tuple with the ClaimsModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) GetClaimsModelOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel, bool) {
	if o == nil || IsNil(o.ClaimsModel) {
		return nil, false
	}
	return o.ClaimsModel, true
}

// HasClaimsModel returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) HasClaimsModel() bool {
	if o != nil && !IsNil(o.ClaimsModel) {
		return true
	}

	return false
}

// SetClaimsModel gets a reference to the given CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel and assigns it to the ClaimsModel field.
func (o *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) SetClaimsModel(v CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) {
	o.ClaimsModel = &v
}

func (o CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClaimsModel) {
		toSerialize["claimsModel"] = o.ClaimsModel
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse struct {
	value *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) Get() *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) Set(val *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse(val *CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) *NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse {
	return &NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


