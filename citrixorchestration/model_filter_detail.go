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

// checks if the FilterDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterDetail{}

// FilterDetail Detailed information about filter.
type FilterDetail struct {
	FilterType *FilterType2 `json:"FilterType,omitempty"`
	// Filter value, if available.
	FilterValue NullableString `json:"FilterValue,omitempty"`
	// Indicate if the filter is a match.
	IsMatched *bool `json:"IsMatched,omitempty"`
	// The filter data. This may need to get translated to a user readable string.
	FilterData NullableString `json:"FilterData,omitempty"`
	// Valid only for the branch repeater filter and access control filter. Indicate the mode or condition.
	IsWithout NullableBool `json:"IsWithout,omitempty"`
	// Valid only for the access control filter. The access gateway.
	Gateway NullableString `json:"Gateway,omitempty"`
}

// NewFilterDetail instantiates a new FilterDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterDetail() *FilterDetail {
	this := FilterDetail{}
	return &this
}

// NewFilterDetailWithDefaults instantiates a new FilterDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterDetailWithDefaults() *FilterDetail {
	this := FilterDetail{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *FilterDetail) GetFilterType() FilterType2 {
	if o == nil || IsNil(o.FilterType) {
		var ret FilterType2
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDetail) GetFilterTypeOk() (*FilterType2, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *FilterDetail) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given FilterType2 and assigns it to the FilterType field.
func (o *FilterDetail) SetFilterType(v FilterType2) {
	o.FilterType = &v
}

// GetFilterValue returns the FilterValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDetail) GetFilterValue() string {
	if o == nil || IsNil(o.FilterValue.Get()) {
		var ret string
		return ret
	}
	return *o.FilterValue.Get()
}

// GetFilterValueOk returns a tuple with the FilterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDetail) GetFilterValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterValue.Get(), o.FilterValue.IsSet()
}

// HasFilterValue returns a boolean if a field has been set.
func (o *FilterDetail) HasFilterValue() bool {
	if o != nil && o.FilterValue.IsSet() {
		return true
	}

	return false
}

// SetFilterValue gets a reference to the given NullableString and assigns it to the FilterValue field.
func (o *FilterDetail) SetFilterValue(v string) {
	o.FilterValue.Set(&v)
}
// SetFilterValueNil sets the value for FilterValue to be an explicit nil
func (o *FilterDetail) SetFilterValueNil() {
	o.FilterValue.Set(nil)
}

// UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil
func (o *FilterDetail) UnsetFilterValue() {
	o.FilterValue.Unset()
}

// GetIsMatched returns the IsMatched field value if set, zero value otherwise.
func (o *FilterDetail) GetIsMatched() bool {
	if o == nil || IsNil(o.IsMatched) {
		var ret bool
		return ret
	}
	return *o.IsMatched
}

// GetIsMatchedOk returns a tuple with the IsMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDetail) GetIsMatchedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMatched) {
		return nil, false
	}
	return o.IsMatched, true
}

// HasIsMatched returns a boolean if a field has been set.
func (o *FilterDetail) HasIsMatched() bool {
	if o != nil && !IsNil(o.IsMatched) {
		return true
	}

	return false
}

// SetIsMatched gets a reference to the given bool and assigns it to the IsMatched field.
func (o *FilterDetail) SetIsMatched(v bool) {
	o.IsMatched = &v
}

// GetFilterData returns the FilterData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDetail) GetFilterData() string {
	if o == nil || IsNil(o.FilterData.Get()) {
		var ret string
		return ret
	}
	return *o.FilterData.Get()
}

// GetFilterDataOk returns a tuple with the FilterData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDetail) GetFilterDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterData.Get(), o.FilterData.IsSet()
}

// HasFilterData returns a boolean if a field has been set.
func (o *FilterDetail) HasFilterData() bool {
	if o != nil && o.FilterData.IsSet() {
		return true
	}

	return false
}

// SetFilterData gets a reference to the given NullableString and assigns it to the FilterData field.
func (o *FilterDetail) SetFilterData(v string) {
	o.FilterData.Set(&v)
}
// SetFilterDataNil sets the value for FilterData to be an explicit nil
func (o *FilterDetail) SetFilterDataNil() {
	o.FilterData.Set(nil)
}

// UnsetFilterData ensures that no value is present for FilterData, not even an explicit nil
func (o *FilterDetail) UnsetFilterData() {
	o.FilterData.Unset()
}

// GetIsWithout returns the IsWithout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDetail) GetIsWithout() bool {
	if o == nil || IsNil(o.IsWithout.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWithout.Get()
}

// GetIsWithoutOk returns a tuple with the IsWithout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDetail) GetIsWithoutOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWithout.Get(), o.IsWithout.IsSet()
}

// HasIsWithout returns a boolean if a field has been set.
func (o *FilterDetail) HasIsWithout() bool {
	if o != nil && o.IsWithout.IsSet() {
		return true
	}

	return false
}

// SetIsWithout gets a reference to the given NullableBool and assigns it to the IsWithout field.
func (o *FilterDetail) SetIsWithout(v bool) {
	o.IsWithout.Set(&v)
}
// SetIsWithoutNil sets the value for IsWithout to be an explicit nil
func (o *FilterDetail) SetIsWithoutNil() {
	o.IsWithout.Set(nil)
}

// UnsetIsWithout ensures that no value is present for IsWithout, not even an explicit nil
func (o *FilterDetail) UnsetIsWithout() {
	o.IsWithout.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterDetail) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterDetail) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *FilterDetail) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *FilterDetail) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *FilterDetail) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *FilterDetail) UnsetGateway() {
	o.Gateway.Unset()
}

func (o FilterDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterType) {
		toSerialize["FilterType"] = o.FilterType
	}
	if o.FilterValue.IsSet() {
		toSerialize["FilterValue"] = o.FilterValue.Get()
	}
	if !IsNil(o.IsMatched) {
		toSerialize["IsMatched"] = o.IsMatched
	}
	if o.FilterData.IsSet() {
		toSerialize["FilterData"] = o.FilterData.Get()
	}
	if o.IsWithout.IsSet() {
		toSerialize["IsWithout"] = o.IsWithout.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["Gateway"] = o.Gateway.Get()
	}
	return toSerialize, nil
}

type NullableFilterDetail struct {
	value *FilterDetail
	isSet bool
}

func (v NullableFilterDetail) Get() *FilterDetail {
	return v.value
}

func (v *NullableFilterDetail) Set(val *FilterDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterDetail(val *FilterDetail) *NullableFilterDetail {
	return &NullableFilterDetail{value: val, isSet: true}
}

func (v NullableFilterDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


