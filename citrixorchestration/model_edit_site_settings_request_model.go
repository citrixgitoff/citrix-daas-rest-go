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

// checks if the EditSiteSettingsRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditSiteSettingsRequestModel{}

// EditSiteSettingsRequestModel struct for EditSiteSettingsRequestModel
type EditSiteSettingsRequestModel struct {
	// Determines whether to use vertical scaling when considering RDS machines for launches. Vertical scaling would saturate machines in the current pool rather than send sessions to the least loaded machines. This would be a trade in performance vs. cost, where vertical scaling would be less costly.
	UseVerticalScalingForRdsLaunches NullableBool `json:"UseVerticalScalingForRdsLaunches,omitempty"`
	// Changes whether ICA files returned by a broker service to a user device contain the numeric IP address or the DNS name of the desktop machine to which a session should be established.
	DnsResolutionEnabled NullableBool `json:"DnsResolutionEnabled,omitempty"`
	// Changes whether the XML Service (as used by Storefront) implicitly trusts the originator of requests it receives, or whether it fully authenticates them.
	TrustRequestsSentToTheXmlServicePortEnabled NullableBool `json:"TrustRequestsSentToTheXmlServicePortEnabled,omitempty"`
	// Determine whether the policy set in web UI is enabled or not
	WebUiPolicySetEnabled NullableBool `json:"WebUiPolicySetEnabled,omitempty"`
	// The max number of minutes that console can be inactive.
	ConsoleInactivityTimeoutMinutes NullableInt32 `json:"ConsoleInactivityTimeoutMinutes,omitempty"`
	SupportedAuthenticators *Authenticator `json:"SupportedAuthenticators,omitempty"`
	// Applicable only for On-Premise. List of origins allowed to make cross-origin requests for Integrated Windows Authentication.
	AllowedCorsOriginsForIwa []string `json:"AllowedCorsOriginsForIwa,omitempty"`
	// Applicable only for On-Premise. Multiple sites configuration.
	MultiSites []MultiSiteModel `json:"MultiSites,omitempty"`
	// The default domain to be used in the login page.
	DefaultDomain NullableString `json:"DefaultDomain,omitempty"`
}

// NewEditSiteSettingsRequestModel instantiates a new EditSiteSettingsRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSiteSettingsRequestModel() *EditSiteSettingsRequestModel {
	this := EditSiteSettingsRequestModel{}
	return &this
}

// NewEditSiteSettingsRequestModelWithDefaults instantiates a new EditSiteSettingsRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSiteSettingsRequestModelWithDefaults() *EditSiteSettingsRequestModel {
	this := EditSiteSettingsRequestModel{}
	return &this
}

// GetUseVerticalScalingForRdsLaunches returns the UseVerticalScalingForRdsLaunches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetUseVerticalScalingForRdsLaunches() bool {
	if o == nil || IsNil(o.UseVerticalScalingForRdsLaunches.Get()) {
		var ret bool
		return ret
	}
	return *o.UseVerticalScalingForRdsLaunches.Get()
}

// GetUseVerticalScalingForRdsLaunchesOk returns a tuple with the UseVerticalScalingForRdsLaunches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetUseVerticalScalingForRdsLaunchesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseVerticalScalingForRdsLaunches.Get(), o.UseVerticalScalingForRdsLaunches.IsSet()
}

// HasUseVerticalScalingForRdsLaunches returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasUseVerticalScalingForRdsLaunches() bool {
	if o != nil && o.UseVerticalScalingForRdsLaunches.IsSet() {
		return true
	}

	return false
}

// SetUseVerticalScalingForRdsLaunches gets a reference to the given NullableBool and assigns it to the UseVerticalScalingForRdsLaunches field.
func (o *EditSiteSettingsRequestModel) SetUseVerticalScalingForRdsLaunches(v bool) {
	o.UseVerticalScalingForRdsLaunches.Set(&v)
}
// SetUseVerticalScalingForRdsLaunchesNil sets the value for UseVerticalScalingForRdsLaunches to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetUseVerticalScalingForRdsLaunchesNil() {
	o.UseVerticalScalingForRdsLaunches.Set(nil)
}

// UnsetUseVerticalScalingForRdsLaunches ensures that no value is present for UseVerticalScalingForRdsLaunches, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetUseVerticalScalingForRdsLaunches() {
	o.UseVerticalScalingForRdsLaunches.Unset()
}

// GetDnsResolutionEnabled returns the DnsResolutionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetDnsResolutionEnabled() bool {
	if o == nil || IsNil(o.DnsResolutionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.DnsResolutionEnabled.Get()
}

// GetDnsResolutionEnabledOk returns a tuple with the DnsResolutionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetDnsResolutionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsResolutionEnabled.Get(), o.DnsResolutionEnabled.IsSet()
}

// HasDnsResolutionEnabled returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasDnsResolutionEnabled() bool {
	if o != nil && o.DnsResolutionEnabled.IsSet() {
		return true
	}

	return false
}

// SetDnsResolutionEnabled gets a reference to the given NullableBool and assigns it to the DnsResolutionEnabled field.
func (o *EditSiteSettingsRequestModel) SetDnsResolutionEnabled(v bool) {
	o.DnsResolutionEnabled.Set(&v)
}
// SetDnsResolutionEnabledNil sets the value for DnsResolutionEnabled to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetDnsResolutionEnabledNil() {
	o.DnsResolutionEnabled.Set(nil)
}

// UnsetDnsResolutionEnabled ensures that no value is present for DnsResolutionEnabled, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetDnsResolutionEnabled() {
	o.DnsResolutionEnabled.Unset()
}

// GetTrustRequestsSentToTheXmlServicePortEnabled returns the TrustRequestsSentToTheXmlServicePortEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetTrustRequestsSentToTheXmlServicePortEnabled() bool {
	if o == nil || IsNil(o.TrustRequestsSentToTheXmlServicePortEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TrustRequestsSentToTheXmlServicePortEnabled.Get()
}

// GetTrustRequestsSentToTheXmlServicePortEnabledOk returns a tuple with the TrustRequestsSentToTheXmlServicePortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetTrustRequestsSentToTheXmlServicePortEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustRequestsSentToTheXmlServicePortEnabled.Get(), o.TrustRequestsSentToTheXmlServicePortEnabled.IsSet()
}

// HasTrustRequestsSentToTheXmlServicePortEnabled returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasTrustRequestsSentToTheXmlServicePortEnabled() bool {
	if o != nil && o.TrustRequestsSentToTheXmlServicePortEnabled.IsSet() {
		return true
	}

	return false
}

// SetTrustRequestsSentToTheXmlServicePortEnabled gets a reference to the given NullableBool and assigns it to the TrustRequestsSentToTheXmlServicePortEnabled field.
func (o *EditSiteSettingsRequestModel) SetTrustRequestsSentToTheXmlServicePortEnabled(v bool) {
	o.TrustRequestsSentToTheXmlServicePortEnabled.Set(&v)
}
// SetTrustRequestsSentToTheXmlServicePortEnabledNil sets the value for TrustRequestsSentToTheXmlServicePortEnabled to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetTrustRequestsSentToTheXmlServicePortEnabledNil() {
	o.TrustRequestsSentToTheXmlServicePortEnabled.Set(nil)
}

// UnsetTrustRequestsSentToTheXmlServicePortEnabled ensures that no value is present for TrustRequestsSentToTheXmlServicePortEnabled, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetTrustRequestsSentToTheXmlServicePortEnabled() {
	o.TrustRequestsSentToTheXmlServicePortEnabled.Unset()
}

// GetWebUiPolicySetEnabled returns the WebUiPolicySetEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetWebUiPolicySetEnabled() bool {
	if o == nil || IsNil(o.WebUiPolicySetEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.WebUiPolicySetEnabled.Get()
}

// GetWebUiPolicySetEnabledOk returns a tuple with the WebUiPolicySetEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetWebUiPolicySetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebUiPolicySetEnabled.Get(), o.WebUiPolicySetEnabled.IsSet()
}

// HasWebUiPolicySetEnabled returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasWebUiPolicySetEnabled() bool {
	if o != nil && o.WebUiPolicySetEnabled.IsSet() {
		return true
	}

	return false
}

// SetWebUiPolicySetEnabled gets a reference to the given NullableBool and assigns it to the WebUiPolicySetEnabled field.
func (o *EditSiteSettingsRequestModel) SetWebUiPolicySetEnabled(v bool) {
	o.WebUiPolicySetEnabled.Set(&v)
}
// SetWebUiPolicySetEnabledNil sets the value for WebUiPolicySetEnabled to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetWebUiPolicySetEnabledNil() {
	o.WebUiPolicySetEnabled.Set(nil)
}

// UnsetWebUiPolicySetEnabled ensures that no value is present for WebUiPolicySetEnabled, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetWebUiPolicySetEnabled() {
	o.WebUiPolicySetEnabled.Unset()
}

// GetConsoleInactivityTimeoutMinutes returns the ConsoleInactivityTimeoutMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetConsoleInactivityTimeoutMinutes() int32 {
	if o == nil || IsNil(o.ConsoleInactivityTimeoutMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.ConsoleInactivityTimeoutMinutes.Get()
}

// GetConsoleInactivityTimeoutMinutesOk returns a tuple with the ConsoleInactivityTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetConsoleInactivityTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsoleInactivityTimeoutMinutes.Get(), o.ConsoleInactivityTimeoutMinutes.IsSet()
}

// HasConsoleInactivityTimeoutMinutes returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasConsoleInactivityTimeoutMinutes() bool {
	if o != nil && o.ConsoleInactivityTimeoutMinutes.IsSet() {
		return true
	}

	return false
}

// SetConsoleInactivityTimeoutMinutes gets a reference to the given NullableInt32 and assigns it to the ConsoleInactivityTimeoutMinutes field.
func (o *EditSiteSettingsRequestModel) SetConsoleInactivityTimeoutMinutes(v int32) {
	o.ConsoleInactivityTimeoutMinutes.Set(&v)
}
// SetConsoleInactivityTimeoutMinutesNil sets the value for ConsoleInactivityTimeoutMinutes to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetConsoleInactivityTimeoutMinutesNil() {
	o.ConsoleInactivityTimeoutMinutes.Set(nil)
}

// UnsetConsoleInactivityTimeoutMinutes ensures that no value is present for ConsoleInactivityTimeoutMinutes, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetConsoleInactivityTimeoutMinutes() {
	o.ConsoleInactivityTimeoutMinutes.Unset()
}

// GetSupportedAuthenticators returns the SupportedAuthenticators field value if set, zero value otherwise.
func (o *EditSiteSettingsRequestModel) GetSupportedAuthenticators() Authenticator {
	if o == nil || IsNil(o.SupportedAuthenticators) {
		var ret Authenticator
		return ret
	}
	return *o.SupportedAuthenticators
}

// GetSupportedAuthenticatorsOk returns a tuple with the SupportedAuthenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSiteSettingsRequestModel) GetSupportedAuthenticatorsOk() (*Authenticator, bool) {
	if o == nil || IsNil(o.SupportedAuthenticators) {
		return nil, false
	}
	return o.SupportedAuthenticators, true
}

// HasSupportedAuthenticators returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasSupportedAuthenticators() bool {
	if o != nil && !IsNil(o.SupportedAuthenticators) {
		return true
	}

	return false
}

// SetSupportedAuthenticators gets a reference to the given Authenticator and assigns it to the SupportedAuthenticators field.
func (o *EditSiteSettingsRequestModel) SetSupportedAuthenticators(v Authenticator) {
	o.SupportedAuthenticators = &v
}

// GetAllowedCorsOriginsForIwa returns the AllowedCorsOriginsForIwa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetAllowedCorsOriginsForIwa() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedCorsOriginsForIwa
}

// GetAllowedCorsOriginsForIwaOk returns a tuple with the AllowedCorsOriginsForIwa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetAllowedCorsOriginsForIwaOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedCorsOriginsForIwa) {
		return nil, false
	}
	return o.AllowedCorsOriginsForIwa, true
}

// HasAllowedCorsOriginsForIwa returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasAllowedCorsOriginsForIwa() bool {
	if o != nil && IsNil(o.AllowedCorsOriginsForIwa) {
		return true
	}

	return false
}

// SetAllowedCorsOriginsForIwa gets a reference to the given []string and assigns it to the AllowedCorsOriginsForIwa field.
func (o *EditSiteSettingsRequestModel) SetAllowedCorsOriginsForIwa(v []string) {
	o.AllowedCorsOriginsForIwa = v
}

// GetMultiSites returns the MultiSites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetMultiSites() []MultiSiteModel {
	if o == nil {
		var ret []MultiSiteModel
		return ret
	}
	return o.MultiSites
}

// GetMultiSitesOk returns a tuple with the MultiSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetMultiSitesOk() ([]MultiSiteModel, bool) {
	if o == nil || IsNil(o.MultiSites) {
		return nil, false
	}
	return o.MultiSites, true
}

// HasMultiSites returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasMultiSites() bool {
	if o != nil && IsNil(o.MultiSites) {
		return true
	}

	return false
}

// SetMultiSites gets a reference to the given []MultiSiteModel and assigns it to the MultiSites field.
func (o *EditSiteSettingsRequestModel) SetMultiSites(v []MultiSiteModel) {
	o.MultiSites = v
}

// GetDefaultDomain returns the DefaultDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditSiteSettingsRequestModel) GetDefaultDomain() string {
	if o == nil || IsNil(o.DefaultDomain.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultDomain.Get()
}

// GetDefaultDomainOk returns a tuple with the DefaultDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditSiteSettingsRequestModel) GetDefaultDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDomain.Get(), o.DefaultDomain.IsSet()
}

// HasDefaultDomain returns a boolean if a field has been set.
func (o *EditSiteSettingsRequestModel) HasDefaultDomain() bool {
	if o != nil && o.DefaultDomain.IsSet() {
		return true
	}

	return false
}

// SetDefaultDomain gets a reference to the given NullableString and assigns it to the DefaultDomain field.
func (o *EditSiteSettingsRequestModel) SetDefaultDomain(v string) {
	o.DefaultDomain.Set(&v)
}
// SetDefaultDomainNil sets the value for DefaultDomain to be an explicit nil
func (o *EditSiteSettingsRequestModel) SetDefaultDomainNil() {
	o.DefaultDomain.Set(nil)
}

// UnsetDefaultDomain ensures that no value is present for DefaultDomain, not even an explicit nil
func (o *EditSiteSettingsRequestModel) UnsetDefaultDomain() {
	o.DefaultDomain.Unset()
}

func (o EditSiteSettingsRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditSiteSettingsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UseVerticalScalingForRdsLaunches.IsSet() {
		toSerialize["UseVerticalScalingForRdsLaunches"] = o.UseVerticalScalingForRdsLaunches.Get()
	}
	if o.DnsResolutionEnabled.IsSet() {
		toSerialize["DnsResolutionEnabled"] = o.DnsResolutionEnabled.Get()
	}
	if o.TrustRequestsSentToTheXmlServicePortEnabled.IsSet() {
		toSerialize["TrustRequestsSentToTheXmlServicePortEnabled"] = o.TrustRequestsSentToTheXmlServicePortEnabled.Get()
	}
	if o.WebUiPolicySetEnabled.IsSet() {
		toSerialize["WebUiPolicySetEnabled"] = o.WebUiPolicySetEnabled.Get()
	}
	if o.ConsoleInactivityTimeoutMinutes.IsSet() {
		toSerialize["ConsoleInactivityTimeoutMinutes"] = o.ConsoleInactivityTimeoutMinutes.Get()
	}
	if !IsNil(o.SupportedAuthenticators) {
		toSerialize["SupportedAuthenticators"] = o.SupportedAuthenticators
	}
	if o.AllowedCorsOriginsForIwa != nil {
		toSerialize["AllowedCorsOriginsForIwa"] = o.AllowedCorsOriginsForIwa
	}
	if o.MultiSites != nil {
		toSerialize["MultiSites"] = o.MultiSites
	}
	if o.DefaultDomain.IsSet() {
		toSerialize["DefaultDomain"] = o.DefaultDomain.Get()
	}
	return toSerialize, nil
}

type NullableEditSiteSettingsRequestModel struct {
	value *EditSiteSettingsRequestModel
	isSet bool
}

func (v NullableEditSiteSettingsRequestModel) Get() *EditSiteSettingsRequestModel {
	return v.value
}

func (v *NullableEditSiteSettingsRequestModel) Set(val *EditSiteSettingsRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSiteSettingsRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSiteSettingsRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSiteSettingsRequestModel(val *EditSiteSettingsRequestModel) *NullableEditSiteSettingsRequestModel {
	return &NullableEditSiteSettingsRequestModel{value: val, isSet: true}
}

func (v NullableEditSiteSettingsRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSiteSettingsRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


