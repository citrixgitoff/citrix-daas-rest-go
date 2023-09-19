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

// checks if the SiteServiceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteServiceResponseModel{}

// SiteServiceResponseModel Service capabilities.
type SiteServiceResponseModel struct {
	// Name of the service.
	ServiceName string `json:"ServiceName"`
	// Type of the service, which will not be changed across languages.
	ServiceType string `json:"ServiceType"`
	// Current schema version of the service. Will be `null` for XenApp & XenDesktop service.
	CurrentSchemaVersion NullableString `json:"CurrentSchemaVersion,omitempty"`
	// Desired schema version of the service. Will be `null` for XenApp & XenDesktop service.
	DesiredSchemaVersion NullableString `json:"DesiredSchemaVersion,omitempty"`
	// List of capabilities exposed by the service.
	Capabilities []string `json:"Capabilities"`
}

// NewSiteServiceResponseModel instantiates a new SiteServiceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteServiceResponseModel(serviceName string, serviceType string, capabilities []string) *SiteServiceResponseModel {
	this := SiteServiceResponseModel{}
	this.ServiceName = serviceName
	this.ServiceType = serviceType
	this.Capabilities = capabilities
	return &this
}

// NewSiteServiceResponseModelWithDefaults instantiates a new SiteServiceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteServiceResponseModelWithDefaults() *SiteServiceResponseModel {
	this := SiteServiceResponseModel{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *SiteServiceResponseModel) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *SiteServiceResponseModel) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *SiteServiceResponseModel) SetServiceName(v string) {
	o.ServiceName = v
}

// GetServiceType returns the ServiceType field value
func (o *SiteServiceResponseModel) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *SiteServiceResponseModel) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *SiteServiceResponseModel) SetServiceType(v string) {
	o.ServiceType = v
}

// GetCurrentSchemaVersion returns the CurrentSchemaVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteServiceResponseModel) GetCurrentSchemaVersion() string {
	if o == nil || IsNil(o.CurrentSchemaVersion.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentSchemaVersion.Get()
}

// GetCurrentSchemaVersionOk returns a tuple with the CurrentSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteServiceResponseModel) GetCurrentSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentSchemaVersion.Get(), o.CurrentSchemaVersion.IsSet()
}

// HasCurrentSchemaVersion returns a boolean if a field has been set.
func (o *SiteServiceResponseModel) HasCurrentSchemaVersion() bool {
	if o != nil && o.CurrentSchemaVersion.IsSet() {
		return true
	}

	return false
}

// SetCurrentSchemaVersion gets a reference to the given NullableString and assigns it to the CurrentSchemaVersion field.
func (o *SiteServiceResponseModel) SetCurrentSchemaVersion(v string) {
	o.CurrentSchemaVersion.Set(&v)
}
// SetCurrentSchemaVersionNil sets the value for CurrentSchemaVersion to be an explicit nil
func (o *SiteServiceResponseModel) SetCurrentSchemaVersionNil() {
	o.CurrentSchemaVersion.Set(nil)
}

// UnsetCurrentSchemaVersion ensures that no value is present for CurrentSchemaVersion, not even an explicit nil
func (o *SiteServiceResponseModel) UnsetCurrentSchemaVersion() {
	o.CurrentSchemaVersion.Unset()
}

// GetDesiredSchemaVersion returns the DesiredSchemaVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteServiceResponseModel) GetDesiredSchemaVersion() string {
	if o == nil || IsNil(o.DesiredSchemaVersion.Get()) {
		var ret string
		return ret
	}
	return *o.DesiredSchemaVersion.Get()
}

// GetDesiredSchemaVersionOk returns a tuple with the DesiredSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteServiceResponseModel) GetDesiredSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DesiredSchemaVersion.Get(), o.DesiredSchemaVersion.IsSet()
}

// HasDesiredSchemaVersion returns a boolean if a field has been set.
func (o *SiteServiceResponseModel) HasDesiredSchemaVersion() bool {
	if o != nil && o.DesiredSchemaVersion.IsSet() {
		return true
	}

	return false
}

// SetDesiredSchemaVersion gets a reference to the given NullableString and assigns it to the DesiredSchemaVersion field.
func (o *SiteServiceResponseModel) SetDesiredSchemaVersion(v string) {
	o.DesiredSchemaVersion.Set(&v)
}
// SetDesiredSchemaVersionNil sets the value for DesiredSchemaVersion to be an explicit nil
func (o *SiteServiceResponseModel) SetDesiredSchemaVersionNil() {
	o.DesiredSchemaVersion.Set(nil)
}

// UnsetDesiredSchemaVersion ensures that no value is present for DesiredSchemaVersion, not even an explicit nil
func (o *SiteServiceResponseModel) UnsetDesiredSchemaVersion() {
	o.DesiredSchemaVersion.Unset()
}

// GetCapabilities returns the Capabilities field value
func (o *SiteServiceResponseModel) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *SiteServiceResponseModel) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *SiteServiceResponseModel) SetCapabilities(v []string) {
	o.Capabilities = v
}

func (o SiteServiceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteServiceResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ServiceName"] = o.ServiceName
	toSerialize["ServiceType"] = o.ServiceType
	if o.CurrentSchemaVersion.IsSet() {
		toSerialize["CurrentSchemaVersion"] = o.CurrentSchemaVersion.Get()
	}
	if o.DesiredSchemaVersion.IsSet() {
		toSerialize["DesiredSchemaVersion"] = o.DesiredSchemaVersion.Get()
	}
	toSerialize["Capabilities"] = o.Capabilities
	return toSerialize, nil
}

type NullableSiteServiceResponseModel struct {
	value *SiteServiceResponseModel
	isSet bool
}

func (v NullableSiteServiceResponseModel) Get() *SiteServiceResponseModel {
	return v.value
}

func (v *NullableSiteServiceResponseModel) Set(val *SiteServiceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteServiceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteServiceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteServiceResponseModel(val *SiteServiceResponseModel) *NullableSiteServiceResponseModel {
	return &NullableSiteServiceResponseModel{value: val, isSet: true}
}

func (v NullableSiteServiceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteServiceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


