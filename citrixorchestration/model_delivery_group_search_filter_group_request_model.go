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

// checks if the DeliveryGroupSearchFilterGroupRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryGroupSearchFilterGroupRequestModel{}

// DeliveryGroupSearchFilterGroupRequestModel Advanced search filter group for applications
type DeliveryGroupSearchFilterGroupRequestModel struct {
	SearchFilterGroupType *DeliveryGroupSearchFilterGroupType `json:"SearchFilterGroupType,omitempty"`
	// The search filters in search filter group
	SearchFilters []DeliveryGroupSearchFilterRequestModel `json:"SearchFilters,omitempty"`
	SearchFilterGroupsType *DeliveryGroupSearchFilterGroupsType `json:"SearchFilterGroupsType,omitempty"`
	// The search filter group in search filter groups
	SearchFilterGroups []DeliveryGroupSearchFilterGroupRequestModel `json:"SearchFilterGroups,omitempty"`
}

// NewDeliveryGroupSearchFilterGroupRequestModel instantiates a new DeliveryGroupSearchFilterGroupRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryGroupSearchFilterGroupRequestModel() *DeliveryGroupSearchFilterGroupRequestModel {
	this := DeliveryGroupSearchFilterGroupRequestModel{}
	return &this
}

// NewDeliveryGroupSearchFilterGroupRequestModelWithDefaults instantiates a new DeliveryGroupSearchFilterGroupRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryGroupSearchFilterGroupRequestModelWithDefaults() *DeliveryGroupSearchFilterGroupRequestModel {
	this := DeliveryGroupSearchFilterGroupRequestModel{}
	return &this
}

// GetSearchFilterGroupType returns the SearchFilterGroupType field value if set, zero value otherwise.
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroupType() DeliveryGroupSearchFilterGroupType {
	if o == nil || IsNil(o.SearchFilterGroupType) {
		var ret DeliveryGroupSearchFilterGroupType
		return ret
	}
	return *o.SearchFilterGroupType
}

// GetSearchFilterGroupTypeOk returns a tuple with the SearchFilterGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroupTypeOk() (*DeliveryGroupSearchFilterGroupType, bool) {
	if o == nil || IsNil(o.SearchFilterGroupType) {
		return nil, false
	}
	return o.SearchFilterGroupType, true
}

// HasSearchFilterGroupType returns a boolean if a field has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) HasSearchFilterGroupType() bool {
	if o != nil && !IsNil(o.SearchFilterGroupType) {
		return true
	}

	return false
}

// SetSearchFilterGroupType gets a reference to the given DeliveryGroupSearchFilterGroupType and assigns it to the SearchFilterGroupType field.
func (o *DeliveryGroupSearchFilterGroupRequestModel) SetSearchFilterGroupType(v DeliveryGroupSearchFilterGroupType) {
	o.SearchFilterGroupType = &v
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilters() []DeliveryGroupSearchFilterRequestModel {
	if o == nil {
		var ret []DeliveryGroupSearchFilterRequestModel
		return ret
	}
	return o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFiltersOk() ([]DeliveryGroupSearchFilterRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) HasSearchFilters() bool {
	if o != nil && IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given []DeliveryGroupSearchFilterRequestModel and assigns it to the SearchFilters field.
func (o *DeliveryGroupSearchFilterGroupRequestModel) SetSearchFilters(v []DeliveryGroupSearchFilterRequestModel) {
	o.SearchFilters = v
}

// GetSearchFilterGroupsType returns the SearchFilterGroupsType field value if set, zero value otherwise.
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroupsType() DeliveryGroupSearchFilterGroupsType {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		var ret DeliveryGroupSearchFilterGroupsType
		return ret
	}
	return *o.SearchFilterGroupsType
}

// GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroupsTypeOk() (*DeliveryGroupSearchFilterGroupsType, bool) {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		return nil, false
	}
	return o.SearchFilterGroupsType, true
}

// HasSearchFilterGroupsType returns a boolean if a field has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) HasSearchFilterGroupsType() bool {
	if o != nil && !IsNil(o.SearchFilterGroupsType) {
		return true
	}

	return false
}

// SetSearchFilterGroupsType gets a reference to the given DeliveryGroupSearchFilterGroupsType and assigns it to the SearchFilterGroupsType field.
func (o *DeliveryGroupSearchFilterGroupRequestModel) SetSearchFilterGroupsType(v DeliveryGroupSearchFilterGroupsType) {
	o.SearchFilterGroupsType = &v
}

// GetSearchFilterGroups returns the SearchFilterGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroups() []DeliveryGroupSearchFilterGroupRequestModel {
	if o == nil {
		var ret []DeliveryGroupSearchFilterGroupRequestModel
		return ret
	}
	return o.SearchFilterGroups
}

// GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryGroupSearchFilterGroupRequestModel) GetSearchFilterGroupsOk() ([]DeliveryGroupSearchFilterGroupRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilterGroups) {
		return nil, false
	}
	return o.SearchFilterGroups, true
}

// HasSearchFilterGroups returns a boolean if a field has been set.
func (o *DeliveryGroupSearchFilterGroupRequestModel) HasSearchFilterGroups() bool {
	if o != nil && IsNil(o.SearchFilterGroups) {
		return true
	}

	return false
}

// SetSearchFilterGroups gets a reference to the given []DeliveryGroupSearchFilterGroupRequestModel and assigns it to the SearchFilterGroups field.
func (o *DeliveryGroupSearchFilterGroupRequestModel) SetSearchFilterGroups(v []DeliveryGroupSearchFilterGroupRequestModel) {
	o.SearchFilterGroups = v
}

func (o DeliveryGroupSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryGroupSearchFilterGroupRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchFilterGroupType) {
		toSerialize["SearchFilterGroupType"] = o.SearchFilterGroupType
	}
	if o.SearchFilters != nil {
		toSerialize["SearchFilters"] = o.SearchFilters
	}
	if !IsNil(o.SearchFilterGroupsType) {
		toSerialize["SearchFilterGroupsType"] = o.SearchFilterGroupsType
	}
	if o.SearchFilterGroups != nil {
		toSerialize["SearchFilterGroups"] = o.SearchFilterGroups
	}
	return toSerialize, nil
}

type NullableDeliveryGroupSearchFilterGroupRequestModel struct {
	value *DeliveryGroupSearchFilterGroupRequestModel
	isSet bool
}

func (v NullableDeliveryGroupSearchFilterGroupRequestModel) Get() *DeliveryGroupSearchFilterGroupRequestModel {
	return v.value
}

func (v *NullableDeliveryGroupSearchFilterGroupRequestModel) Set(val *DeliveryGroupSearchFilterGroupRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGroupSearchFilterGroupRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGroupSearchFilterGroupRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGroupSearchFilterGroupRequestModel(val *DeliveryGroupSearchFilterGroupRequestModel) *NullableDeliveryGroupSearchFilterGroupRequestModel {
	return &NullableDeliveryGroupSearchFilterGroupRequestModel{value: val, isSet: true}
}

func (v NullableDeliveryGroupSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGroupSearchFilterGroupRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


