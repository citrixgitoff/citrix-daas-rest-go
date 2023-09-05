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

// checks if the EditHypervisorAzureConnectionRequestModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHypervisorAzureConnectionRequestModelAllOf{}

// EditHypervisorAzureConnectionRequestModelAllOf struct for EditHypervisorAzureConnectionRequestModelAllOf
type EditHypervisorAzureConnectionRequestModelAllOf struct {
	// Application ID of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, then ApplicationSecret must also be specified.
	ApplicationId *string `json:"ApplicationId,omitempty"`
	// The Application Secret of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, must in the format indicated by ApplicationSecretFormat.
	ApplicationSecret *string `json:"ApplicationSecret,omitempty"`
	ApplicationSecretFormat *IdentityPasswordFormat `json:"ApplicationSecretFormat,omitempty"`
	// The properties of host connection that are specific to the target hosting infrastructure.
	CustomProperties *string `json:"CustomProperties,omitempty"`
}

// NewEditHypervisorAzureConnectionRequestModelAllOf instantiates a new EditHypervisorAzureConnectionRequestModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHypervisorAzureConnectionRequestModelAllOf() *EditHypervisorAzureConnectionRequestModelAllOf {
	this := EditHypervisorAzureConnectionRequestModelAllOf{}
	return &this
}

// NewEditHypervisorAzureConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorAzureConnectionRequestModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHypervisorAzureConnectionRequestModelAllOfWithDefaults() *EditHypervisorAzureConnectionRequestModelAllOf {
	this := EditHypervisorAzureConnectionRequestModelAllOf{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationSecret returns the ApplicationSecret field value if set, zero value otherwise.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecret() string {
	if o == nil || IsNil(o.ApplicationSecret) {
		var ret string
		return ret
	}
	return *o.ApplicationSecret
}

// GetApplicationSecretOk returns a tuple with the ApplicationSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationSecret) {
		return nil, false
	}
	return o.ApplicationSecret, true
}

// HasApplicationSecret returns a boolean if a field has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationSecret() bool {
	if o != nil && !IsNil(o.ApplicationSecret) {
		return true
	}

	return false
}

// SetApplicationSecret gets a reference to the given string and assigns it to the ApplicationSecret field.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationSecret(v string) {
	o.ApplicationSecret = &v
}

// GetApplicationSecretFormat returns the ApplicationSecretFormat field value if set, zero value otherwise.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.ApplicationSecretFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.ApplicationSecretFormat
}

// GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.ApplicationSecretFormat) {
		return nil, false
	}
	return o.ApplicationSecretFormat, true
}

// HasApplicationSecretFormat returns a boolean if a field has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationSecretFormat() bool {
	if o != nil && !IsNil(o.ApplicationSecretFormat) {
		return true
	}

	return false
}

// SetApplicationSecretFormat gets a reference to the given IdentityPasswordFormat and assigns it to the ApplicationSecretFormat field.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationSecretFormat(v IdentityPasswordFormat) {
	o.ApplicationSecretFormat = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties) {
		var ret string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetCustomPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given string and assigns it to the CustomProperties field.
func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetCustomProperties(v string) {
	o.CustomProperties = &v
}

func (o EditHypervisorAzureConnectionRequestModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHypervisorAzureConnectionRequestModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["ApplicationId"] = o.ApplicationId
	}
	if !IsNil(o.ApplicationSecret) {
		toSerialize["ApplicationSecret"] = o.ApplicationSecret
	}
	if !IsNil(o.ApplicationSecretFormat) {
		toSerialize["ApplicationSecretFormat"] = o.ApplicationSecretFormat
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["CustomProperties"] = o.CustomProperties
	}
	return toSerialize, nil
}

type NullableEditHypervisorAzureConnectionRequestModelAllOf struct {
	value *EditHypervisorAzureConnectionRequestModelAllOf
	isSet bool
}

func (v NullableEditHypervisorAzureConnectionRequestModelAllOf) Get() *EditHypervisorAzureConnectionRequestModelAllOf {
	return v.value
}

func (v *NullableEditHypervisorAzureConnectionRequestModelAllOf) Set(val *EditHypervisorAzureConnectionRequestModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHypervisorAzureConnectionRequestModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHypervisorAzureConnectionRequestModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHypervisorAzureConnectionRequestModelAllOf(val *EditHypervisorAzureConnectionRequestModelAllOf) *NullableEditHypervisorAzureConnectionRequestModelAllOf {
	return &NullableEditHypervisorAzureConnectionRequestModelAllOf{value: val, isSet: true}
}

func (v NullableEditHypervisorAzureConnectionRequestModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHypervisorAzureConnectionRequestModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


