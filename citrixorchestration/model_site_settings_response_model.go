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

// checks if the SiteSettingsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingsResponseModel{}

// SiteSettingsResponseModel struct for SiteSettingsResponseModel
type SiteSettingsResponseModel struct {
	UseVerticalScalingForRdsLaunches *bool `json:"UseVerticalScalingForRdsLaunches,omitempty"`
	DnsResolutionEnabled *bool `json:"DnsResolutionEnabled,omitempty"`
	TrustRequestsSentToTheXmlServicePortEnabled *bool `json:"TrustRequestsSentToTheXmlServicePortEnabled,omitempty"`
	// The policy set in web UI is enabled or not
	WebUiPolicySetEnabled *bool `json:"WebUiPolicySetEnabled,omitempty"`
}

// NewSiteSettingsResponseModel instantiates a new SiteSettingsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingsResponseModel() *SiteSettingsResponseModel {
	this := SiteSettingsResponseModel{}
	return &this
}

// NewSiteSettingsResponseModelWithDefaults instantiates a new SiteSettingsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingsResponseModelWithDefaults() *SiteSettingsResponseModel {
	this := SiteSettingsResponseModel{}
	return &this
}

// GetUseVerticalScalingForRdsLaunches returns the UseVerticalScalingForRdsLaunches field value if set, zero value otherwise.
func (o *SiteSettingsResponseModel) GetUseVerticalScalingForRdsLaunches() bool {
	if o == nil || IsNil(o.UseVerticalScalingForRdsLaunches) {
		var ret bool
		return ret
	}
	return *o.UseVerticalScalingForRdsLaunches
}

// GetUseVerticalScalingForRdsLaunchesOk returns a tuple with the UseVerticalScalingForRdsLaunches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingsResponseModel) GetUseVerticalScalingForRdsLaunchesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseVerticalScalingForRdsLaunches) {
		return nil, false
	}
	return o.UseVerticalScalingForRdsLaunches, true
}

// HasUseVerticalScalingForRdsLaunches returns a boolean if a field has been set.
func (o *SiteSettingsResponseModel) HasUseVerticalScalingForRdsLaunches() bool {
	if o != nil && !IsNil(o.UseVerticalScalingForRdsLaunches) {
		return true
	}

	return false
}

// SetUseVerticalScalingForRdsLaunches gets a reference to the given bool and assigns it to the UseVerticalScalingForRdsLaunches field.
func (o *SiteSettingsResponseModel) SetUseVerticalScalingForRdsLaunches(v bool) {
	o.UseVerticalScalingForRdsLaunches = &v
}

// GetDnsResolutionEnabled returns the DnsResolutionEnabled field value if set, zero value otherwise.
func (o *SiteSettingsResponseModel) GetDnsResolutionEnabled() bool {
	if o == nil || IsNil(o.DnsResolutionEnabled) {
		var ret bool
		return ret
	}
	return *o.DnsResolutionEnabled
}

// GetDnsResolutionEnabledOk returns a tuple with the DnsResolutionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingsResponseModel) GetDnsResolutionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsResolutionEnabled) {
		return nil, false
	}
	return o.DnsResolutionEnabled, true
}

// HasDnsResolutionEnabled returns a boolean if a field has been set.
func (o *SiteSettingsResponseModel) HasDnsResolutionEnabled() bool {
	if o != nil && !IsNil(o.DnsResolutionEnabled) {
		return true
	}

	return false
}

// SetDnsResolutionEnabled gets a reference to the given bool and assigns it to the DnsResolutionEnabled field.
func (o *SiteSettingsResponseModel) SetDnsResolutionEnabled(v bool) {
	o.DnsResolutionEnabled = &v
}

// GetTrustRequestsSentToTheXmlServicePortEnabled returns the TrustRequestsSentToTheXmlServicePortEnabled field value if set, zero value otherwise.
func (o *SiteSettingsResponseModel) GetTrustRequestsSentToTheXmlServicePortEnabled() bool {
	if o == nil || IsNil(o.TrustRequestsSentToTheXmlServicePortEnabled) {
		var ret bool
		return ret
	}
	return *o.TrustRequestsSentToTheXmlServicePortEnabled
}

// GetTrustRequestsSentToTheXmlServicePortEnabledOk returns a tuple with the TrustRequestsSentToTheXmlServicePortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingsResponseModel) GetTrustRequestsSentToTheXmlServicePortEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TrustRequestsSentToTheXmlServicePortEnabled) {
		return nil, false
	}
	return o.TrustRequestsSentToTheXmlServicePortEnabled, true
}

// HasTrustRequestsSentToTheXmlServicePortEnabled returns a boolean if a field has been set.
func (o *SiteSettingsResponseModel) HasTrustRequestsSentToTheXmlServicePortEnabled() bool {
	if o != nil && !IsNil(o.TrustRequestsSentToTheXmlServicePortEnabled) {
		return true
	}

	return false
}

// SetTrustRequestsSentToTheXmlServicePortEnabled gets a reference to the given bool and assigns it to the TrustRequestsSentToTheXmlServicePortEnabled field.
func (o *SiteSettingsResponseModel) SetTrustRequestsSentToTheXmlServicePortEnabled(v bool) {
	o.TrustRequestsSentToTheXmlServicePortEnabled = &v
}

// GetWebUiPolicySetEnabled returns the WebUiPolicySetEnabled field value if set, zero value otherwise.
func (o *SiteSettingsResponseModel) GetWebUiPolicySetEnabled() bool {
	if o == nil || IsNil(o.WebUiPolicySetEnabled) {
		var ret bool
		return ret
	}
	return *o.WebUiPolicySetEnabled
}

// GetWebUiPolicySetEnabledOk returns a tuple with the WebUiPolicySetEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingsResponseModel) GetWebUiPolicySetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WebUiPolicySetEnabled) {
		return nil, false
	}
	return o.WebUiPolicySetEnabled, true
}

// HasWebUiPolicySetEnabled returns a boolean if a field has been set.
func (o *SiteSettingsResponseModel) HasWebUiPolicySetEnabled() bool {
	if o != nil && !IsNil(o.WebUiPolicySetEnabled) {
		return true
	}

	return false
}

// SetWebUiPolicySetEnabled gets a reference to the given bool and assigns it to the WebUiPolicySetEnabled field.
func (o *SiteSettingsResponseModel) SetWebUiPolicySetEnabled(v bool) {
	o.WebUiPolicySetEnabled = &v
}

func (o SiteSettingsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseVerticalScalingForRdsLaunches) {
		toSerialize["UseVerticalScalingForRdsLaunches"] = o.UseVerticalScalingForRdsLaunches
	}
	if !IsNil(o.DnsResolutionEnabled) {
		toSerialize["DnsResolutionEnabled"] = o.DnsResolutionEnabled
	}
	if !IsNil(o.TrustRequestsSentToTheXmlServicePortEnabled) {
		toSerialize["TrustRequestsSentToTheXmlServicePortEnabled"] = o.TrustRequestsSentToTheXmlServicePortEnabled
	}
	if !IsNil(o.WebUiPolicySetEnabled) {
		toSerialize["WebUiPolicySetEnabled"] = o.WebUiPolicySetEnabled
	}
	return toSerialize, nil
}

type NullableSiteSettingsResponseModel struct {
	value *SiteSettingsResponseModel
	isSet bool
}

func (v NullableSiteSettingsResponseModel) Get() *SiteSettingsResponseModel {
	return v.value
}

func (v *NullableSiteSettingsResponseModel) Set(val *SiteSettingsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingsResponseModel(val *SiteSettingsResponseModel) *NullableSiteSettingsResponseModel {
	return &NullableSiteSettingsResponseModel{value: val, isSet: true}
}

func (v NullableSiteSettingsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


