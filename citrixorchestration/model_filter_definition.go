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

// checks if the FilterDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterDefinition{}

// FilterDefinition Filter definition.
type FilterDefinition struct {
	FilterType *FilterType `json:"filterType,omitempty"`
	// Localized filter name
	FilterName NullableString `json:"filterName,omitempty"`
	// Localized explanation
	Explanation NullableString `json:"explanation,omitempty"`
	// True = user filter, False = machine filter
	IsUserFilter *bool `json:"isUserFilter,omitempty"`
	// Is filter a singleton, only the NetScaler SD-Wan filter is singleton.
	IsSingleton *bool `json:"isSingleton,omitempty"`
}

// NewFilterDefinition instantiates a new FilterDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterDefinition() *FilterDefinition {
	this := FilterDefinition{}
	return &this
}

// NewFilterDefinitionWithDefaults instantiates a new FilterDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterDefinitionWithDefaults() *FilterDefinition {
	this := FilterDefinition{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *FilterDefinition) GetFilterType() FilterType {
	if o == nil || IsNil(o.FilterType) {
		var ret FilterType
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDefinition) GetFilterTypeOk() (*FilterType, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *FilterDefinition) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given FilterType and assigns it to the FilterType field.
func (o *FilterDefinition) SetFilterType(v FilterType) {
	o.FilterType = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDefinition) GetFilterName() string {
	if o == nil || IsNil(o.FilterName.Get()) {
		var ret string
		return ret
	}
	return *o.FilterName.Get()
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDefinition) GetFilterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterName.Get(), o.FilterName.IsSet()
}

// HasFilterName returns a boolean if a field has been set.
func (o *FilterDefinition) HasFilterName() bool {
	if o != nil && o.FilterName.IsSet() {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given NullableString and assigns it to the FilterName field.
func (o *FilterDefinition) SetFilterName(v string) {
	o.FilterName.Set(&v)
}
// SetFilterNameNil sets the value for FilterName to be an explicit nil
func (o *FilterDefinition) SetFilterNameNil() {
	o.FilterName.Set(nil)
}

// UnsetFilterName ensures that no value is present for FilterName, not even an explicit nil
func (o *FilterDefinition) UnsetFilterName() {
	o.FilterName.Unset()
}

// GetExplanation returns the Explanation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDefinition) GetExplanation() string {
	if o == nil || IsNil(o.Explanation.Get()) {
		var ret string
		return ret
	}
	return *o.Explanation.Get()
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDefinition) GetExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Explanation.Get(), o.Explanation.IsSet()
}

// HasExplanation returns a boolean if a field has been set.
func (o *FilterDefinition) HasExplanation() bool {
	if o != nil && o.Explanation.IsSet() {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given NullableString and assigns it to the Explanation field.
func (o *FilterDefinition) SetExplanation(v string) {
	o.Explanation.Set(&v)
}
// SetExplanationNil sets the value for Explanation to be an explicit nil
func (o *FilterDefinition) SetExplanationNil() {
	o.Explanation.Set(nil)
}

// UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
func (o *FilterDefinition) UnsetExplanation() {
	o.Explanation.Unset()
}

// GetIsUserFilter returns the IsUserFilter field value if set, zero value otherwise.
func (o *FilterDefinition) GetIsUserFilter() bool {
	if o == nil || IsNil(o.IsUserFilter) {
		var ret bool
		return ret
	}
	return *o.IsUserFilter
}

// GetIsUserFilterOk returns a tuple with the IsUserFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDefinition) GetIsUserFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserFilter) {
		return nil, false
	}
	return o.IsUserFilter, true
}

// HasIsUserFilter returns a boolean if a field has been set.
func (o *FilterDefinition) HasIsUserFilter() bool {
	if o != nil && !IsNil(o.IsUserFilter) {
		return true
	}

	return false
}

// SetIsUserFilter gets a reference to the given bool and assigns it to the IsUserFilter field.
func (o *FilterDefinition) SetIsUserFilter(v bool) {
	o.IsUserFilter = &v
}

// GetIsSingleton returns the IsSingleton field value if set, zero value otherwise.
func (o *FilterDefinition) GetIsSingleton() bool {
	if o == nil || IsNil(o.IsSingleton) {
		var ret bool
		return ret
	}
	return *o.IsSingleton
}

// GetIsSingletonOk returns a tuple with the IsSingleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDefinition) GetIsSingletonOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSingleton) {
		return nil, false
	}
	return o.IsSingleton, true
}

// HasIsSingleton returns a boolean if a field has been set.
func (o *FilterDefinition) HasIsSingleton() bool {
	if o != nil && !IsNil(o.IsSingleton) {
		return true
	}

	return false
}

// SetIsSingleton gets a reference to the given bool and assigns it to the IsSingleton field.
func (o *FilterDefinition) SetIsSingleton(v bool) {
	o.IsSingleton = &v
}

func (o FilterDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if o.FilterName.IsSet() {
		toSerialize["filterName"] = o.FilterName.Get()
	}
	if o.Explanation.IsSet() {
		toSerialize["explanation"] = o.Explanation.Get()
	}
	if !IsNil(o.IsUserFilter) {
		toSerialize["isUserFilter"] = o.IsUserFilter
	}
	if !IsNil(o.IsSingleton) {
		toSerialize["isSingleton"] = o.IsSingleton
	}
	return toSerialize, nil
}

type NullableFilterDefinition struct {
	value *FilterDefinition
	isSet bool
}

func (v NullableFilterDefinition) Get() *FilterDefinition {
	return v.value
}

func (v *NullableFilterDefinition) Set(val *FilterDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterDefinition(val *FilterDefinition) *NullableFilterDefinition {
	return &NullableFilterDefinition{value: val, isSet: true}
}

func (v NullableFilterDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


