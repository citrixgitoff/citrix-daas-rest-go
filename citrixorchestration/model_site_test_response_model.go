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

// checks if the SiteTestResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteTestResponseModel{}

// SiteTestResponseModel Response object for a site test result.
type SiteTestResponseModel struct {
	Site RefResponseModel `json:"Site"`
	// The number of tests that passed.
	NumPassed int32 `json:"NumPassed"`
	// The number of warnings that were found.
	NumWarnings int32 `json:"NumWarnings"`
	// The number of tests that failed.
	NumFailures int32 `json:"NumFailures"`
}

// NewSiteTestResponseModel instantiates a new SiteTestResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteTestResponseModel(site RefResponseModel, numPassed int32, numWarnings int32, numFailures int32) *SiteTestResponseModel {
	this := SiteTestResponseModel{}
	this.Site = site
	this.NumPassed = numPassed
	this.NumWarnings = numWarnings
	this.NumFailures = numFailures
	return &this
}

// NewSiteTestResponseModelWithDefaults instantiates a new SiteTestResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteTestResponseModelWithDefaults() *SiteTestResponseModel {
	this := SiteTestResponseModel{}
	return &this
}

// GetSite returns the Site field value
func (o *SiteTestResponseModel) GetSite() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *SiteTestResponseModel) GetSiteOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *SiteTestResponseModel) SetSite(v RefResponseModel) {
	o.Site = v
}

// GetNumPassed returns the NumPassed field value
func (o *SiteTestResponseModel) GetNumPassed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPassed
}

// GetNumPassedOk returns a tuple with the NumPassed field value
// and a boolean to check if the value has been set.
func (o *SiteTestResponseModel) GetNumPassedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPassed, true
}

// SetNumPassed sets field value
func (o *SiteTestResponseModel) SetNumPassed(v int32) {
	o.NumPassed = v
}

// GetNumWarnings returns the NumWarnings field value
func (o *SiteTestResponseModel) GetNumWarnings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWarnings
}

// GetNumWarningsOk returns a tuple with the NumWarnings field value
// and a boolean to check if the value has been set.
func (o *SiteTestResponseModel) GetNumWarningsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWarnings, true
}

// SetNumWarnings sets field value
func (o *SiteTestResponseModel) SetNumWarnings(v int32) {
	o.NumWarnings = v
}

// GetNumFailures returns the NumFailures field value
func (o *SiteTestResponseModel) GetNumFailures() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumFailures
}

// GetNumFailuresOk returns a tuple with the NumFailures field value
// and a boolean to check if the value has been set.
func (o *SiteTestResponseModel) GetNumFailuresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumFailures, true
}

// SetNumFailures sets field value
func (o *SiteTestResponseModel) SetNumFailures(v int32) {
	o.NumFailures = v
}

func (o SiteTestResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteTestResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Site"] = o.Site
	toSerialize["NumPassed"] = o.NumPassed
	toSerialize["NumWarnings"] = o.NumWarnings
	toSerialize["NumFailures"] = o.NumFailures
	return toSerialize, nil
}

type NullableSiteTestResponseModel struct {
	value *SiteTestResponseModel
	isSet bool
}

func (v NullableSiteTestResponseModel) Get() *SiteTestResponseModel {
	return v.value
}

func (v *NullableSiteTestResponseModel) Set(val *SiteTestResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteTestResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteTestResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteTestResponseModel(val *SiteTestResponseModel) *NullableSiteTestResponseModel {
	return &NullableSiteTestResponseModel{value: val, isSet: true}
}

func (v NullableSiteTestResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteTestResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


