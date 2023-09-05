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

// checks if the MachineAndSessionSearchFilterGroupRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineAndSessionSearchFilterGroupRequestModel{}

// MachineAndSessionSearchFilterGroupRequestModel Advanced search filter group for machines and sessions
type MachineAndSessionSearchFilterGroupRequestModel struct {
	SearchFilterGroupType *MachineAndSessionSearchFilterGroupType `json:"SearchFilterGroupType,omitempty"`
	// The search filters in search filter group
	SearchFilters []MachineAndSessionSearchFilterRequestModel `json:"SearchFilters,omitempty"`
	SearchFilterGroupsType *MachineAndSessionSearchFilterGroupsType `json:"SearchFilterGroupsType,omitempty"`
	// The search filter group in search filter groups
	SearchFilterGroups []MachineAndSessionSearchFilterGroupRequestModel `json:"SearchFilterGroups,omitempty"`
}

// NewMachineAndSessionSearchFilterGroupRequestModel instantiates a new MachineAndSessionSearchFilterGroupRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineAndSessionSearchFilterGroupRequestModel() *MachineAndSessionSearchFilterGroupRequestModel {
	this := MachineAndSessionSearchFilterGroupRequestModel{}
	return &this
}

// NewMachineAndSessionSearchFilterGroupRequestModelWithDefaults instantiates a new MachineAndSessionSearchFilterGroupRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineAndSessionSearchFilterGroupRequestModelWithDefaults() *MachineAndSessionSearchFilterGroupRequestModel {
	this := MachineAndSessionSearchFilterGroupRequestModel{}
	return &this
}

// GetSearchFilterGroupType returns the SearchFilterGroupType field value if set, zero value otherwise.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroupType() MachineAndSessionSearchFilterGroupType {
	if o == nil || IsNil(o.SearchFilterGroupType) {
		var ret MachineAndSessionSearchFilterGroupType
		return ret
	}
	return *o.SearchFilterGroupType
}

// GetSearchFilterGroupTypeOk returns a tuple with the SearchFilterGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroupTypeOk() (*MachineAndSessionSearchFilterGroupType, bool) {
	if o == nil || IsNil(o.SearchFilterGroupType) {
		return nil, false
	}
	return o.SearchFilterGroupType, true
}

// HasSearchFilterGroupType returns a boolean if a field has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) HasSearchFilterGroupType() bool {
	if o != nil && !IsNil(o.SearchFilterGroupType) {
		return true
	}

	return false
}

// SetSearchFilterGroupType gets a reference to the given MachineAndSessionSearchFilterGroupType and assigns it to the SearchFilterGroupType field.
func (o *MachineAndSessionSearchFilterGroupRequestModel) SetSearchFilterGroupType(v MachineAndSessionSearchFilterGroupType) {
	o.SearchFilterGroupType = &v
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilters() []MachineAndSessionSearchFilterRequestModel {
	if o == nil || IsNil(o.SearchFilters) {
		var ret []MachineAndSessionSearchFilterRequestModel
		return ret
	}
	return o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFiltersOk() ([]MachineAndSessionSearchFilterRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) HasSearchFilters() bool {
	if o != nil && !IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given []MachineAndSessionSearchFilterRequestModel and assigns it to the SearchFilters field.
func (o *MachineAndSessionSearchFilterGroupRequestModel) SetSearchFilters(v []MachineAndSessionSearchFilterRequestModel) {
	o.SearchFilters = v
}

// GetSearchFilterGroupsType returns the SearchFilterGroupsType field value if set, zero value otherwise.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroupsType() MachineAndSessionSearchFilterGroupsType {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		var ret MachineAndSessionSearchFilterGroupsType
		return ret
	}
	return *o.SearchFilterGroupsType
}

// GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroupsTypeOk() (*MachineAndSessionSearchFilterGroupsType, bool) {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		return nil, false
	}
	return o.SearchFilterGroupsType, true
}

// HasSearchFilterGroupsType returns a boolean if a field has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) HasSearchFilterGroupsType() bool {
	if o != nil && !IsNil(o.SearchFilterGroupsType) {
		return true
	}

	return false
}

// SetSearchFilterGroupsType gets a reference to the given MachineAndSessionSearchFilterGroupsType and assigns it to the SearchFilterGroupsType field.
func (o *MachineAndSessionSearchFilterGroupRequestModel) SetSearchFilterGroupsType(v MachineAndSessionSearchFilterGroupsType) {
	o.SearchFilterGroupsType = &v
}

// GetSearchFilterGroups returns the SearchFilterGroups field value if set, zero value otherwise.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroups() []MachineAndSessionSearchFilterGroupRequestModel {
	if o == nil || IsNil(o.SearchFilterGroups) {
		var ret []MachineAndSessionSearchFilterGroupRequestModel
		return ret
	}
	return o.SearchFilterGroups
}

// GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) GetSearchFilterGroupsOk() ([]MachineAndSessionSearchFilterGroupRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilterGroups) {
		return nil, false
	}
	return o.SearchFilterGroups, true
}

// HasSearchFilterGroups returns a boolean if a field has been set.
func (o *MachineAndSessionSearchFilterGroupRequestModel) HasSearchFilterGroups() bool {
	if o != nil && !IsNil(o.SearchFilterGroups) {
		return true
	}

	return false
}

// SetSearchFilterGroups gets a reference to the given []MachineAndSessionSearchFilterGroupRequestModel and assigns it to the SearchFilterGroups field.
func (o *MachineAndSessionSearchFilterGroupRequestModel) SetSearchFilterGroups(v []MachineAndSessionSearchFilterGroupRequestModel) {
	o.SearchFilterGroups = v
}

func (o MachineAndSessionSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineAndSessionSearchFilterGroupRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchFilterGroupType) {
		toSerialize["SearchFilterGroupType"] = o.SearchFilterGroupType
	}
	if !IsNil(o.SearchFilters) {
		toSerialize["SearchFilters"] = o.SearchFilters
	}
	if !IsNil(o.SearchFilterGroupsType) {
		toSerialize["SearchFilterGroupsType"] = o.SearchFilterGroupsType
	}
	if !IsNil(o.SearchFilterGroups) {
		toSerialize["SearchFilterGroups"] = o.SearchFilterGroups
	}
	return toSerialize, nil
}

type NullableMachineAndSessionSearchFilterGroupRequestModel struct {
	value *MachineAndSessionSearchFilterGroupRequestModel
	isSet bool
}

func (v NullableMachineAndSessionSearchFilterGroupRequestModel) Get() *MachineAndSessionSearchFilterGroupRequestModel {
	return v.value
}

func (v *NullableMachineAndSessionSearchFilterGroupRequestModel) Set(val *MachineAndSessionSearchFilterGroupRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAndSessionSearchFilterGroupRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAndSessionSearchFilterGroupRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAndSessionSearchFilterGroupRequestModel(val *MachineAndSessionSearchFilterGroupRequestModel) *NullableMachineAndSessionSearchFilterGroupRequestModel {
	return &NullableMachineAndSessionSearchFilterGroupRequestModel{value: val, isSet: true}
}

func (v NullableMachineAndSessionSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAndSessionSearchFilterGroupRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


