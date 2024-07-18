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

// checks if the ProvisioningOperationEventSearchRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningOperationEventSearchRequestModel{}

// ProvisioningOperationEventSearchRequestModel Request model for defining provisioning operation event.
type ProvisioningOperationEventSearchRequestModel struct {
	// Basic search string. Specify a string which will match if contained within some string property of the operation event.
	BasicSearchString NullableString `json:"BasicSearchString,omitempty"`
	// List of advanced search filters.
	SearchFilters []ProvisioningOperationEventSearchFilterRequestModel `json:"SearchFilters,omitempty"`
	// List of advanced search filter groups.
	SearchFilterGroups []ProvisioningOperationEventSearchFilterGroupRequestModel `json:"SearchFilterGroups,omitempty"`
	SearchFilterGroupsType *ProvisioningOperationEventSearchFilterGroupsType `json:"SearchFilterGroupsType,omitempty"`
	// Sort criteria for the results, multiple sorting criteria can be specified here.
	SortCriteriaItems []ProvisioningOperationEventSortCriteriaRequestModel `json:"SortCriteriaItems,omitempty"`
}

// NewProvisioningOperationEventSearchRequestModel instantiates a new ProvisioningOperationEventSearchRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningOperationEventSearchRequestModel() *ProvisioningOperationEventSearchRequestModel {
	this := ProvisioningOperationEventSearchRequestModel{}
	return &this
}

// NewProvisioningOperationEventSearchRequestModelWithDefaults instantiates a new ProvisioningOperationEventSearchRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningOperationEventSearchRequestModelWithDefaults() *ProvisioningOperationEventSearchRequestModel {
	this := ProvisioningOperationEventSearchRequestModel{}
	return &this
}

// GetBasicSearchString returns the BasicSearchString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningOperationEventSearchRequestModel) GetBasicSearchString() string {
	if o == nil || IsNil(o.BasicSearchString.Get()) {
		var ret string
		return ret
	}
	return *o.BasicSearchString.Get()
}

// GetBasicSearchStringOk returns a tuple with the BasicSearchString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningOperationEventSearchRequestModel) GetBasicSearchStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BasicSearchString.Get(), o.BasicSearchString.IsSet()
}

// HasBasicSearchString returns a boolean if a field has been set.
func (o *ProvisioningOperationEventSearchRequestModel) HasBasicSearchString() bool {
	if o != nil && o.BasicSearchString.IsSet() {
		return true
	}

	return false
}

// SetBasicSearchString gets a reference to the given NullableString and assigns it to the BasicSearchString field.
func (o *ProvisioningOperationEventSearchRequestModel) SetBasicSearchString(v string) {
	o.BasicSearchString.Set(&v)
}
// SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil
func (o *ProvisioningOperationEventSearchRequestModel) SetBasicSearchStringNil() {
	o.BasicSearchString.Set(nil)
}

// UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
func (o *ProvisioningOperationEventSearchRequestModel) UnsetBasicSearchString() {
	o.BasicSearchString.Unset()
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilters() []ProvisioningOperationEventSearchFilterRequestModel {
	if o == nil {
		var ret []ProvisioningOperationEventSearchFilterRequestModel
		return ret
	}
	return o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFiltersOk() ([]ProvisioningOperationEventSearchFilterRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilters() bool {
	if o != nil && IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given []ProvisioningOperationEventSearchFilterRequestModel and assigns it to the SearchFilters field.
func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilters(v []ProvisioningOperationEventSearchFilterRequestModel) {
	o.SearchFilters = v
}

// GetSearchFilterGroups returns the SearchFilterGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroups() []ProvisioningOperationEventSearchFilterGroupRequestModel {
	if o == nil {
		var ret []ProvisioningOperationEventSearchFilterGroupRequestModel
		return ret
	}
	return o.SearchFilterGroups
}

// GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsOk() ([]ProvisioningOperationEventSearchFilterGroupRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilterGroups) {
		return nil, false
	}
	return o.SearchFilterGroups, true
}

// HasSearchFilterGroups returns a boolean if a field has been set.
func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilterGroups() bool {
	if o != nil && IsNil(o.SearchFilterGroups) {
		return true
	}

	return false
}

// SetSearchFilterGroups gets a reference to the given []ProvisioningOperationEventSearchFilterGroupRequestModel and assigns it to the SearchFilterGroups field.
func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilterGroups(v []ProvisioningOperationEventSearchFilterGroupRequestModel) {
	o.SearchFilterGroups = v
}

// GetSearchFilterGroupsType returns the SearchFilterGroupsType field value if set, zero value otherwise.
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsType() ProvisioningOperationEventSearchFilterGroupsType {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		var ret ProvisioningOperationEventSearchFilterGroupsType
		return ret
	}
	return *o.SearchFilterGroupsType
}

// GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsTypeOk() (*ProvisioningOperationEventSearchFilterGroupsType, bool) {
	if o == nil || IsNil(o.SearchFilterGroupsType) {
		return nil, false
	}
	return o.SearchFilterGroupsType, true
}

// HasSearchFilterGroupsType returns a boolean if a field has been set.
func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilterGroupsType() bool {
	if o != nil && !IsNil(o.SearchFilterGroupsType) {
		return true
	}

	return false
}

// SetSearchFilterGroupsType gets a reference to the given ProvisioningOperationEventSearchFilterGroupsType and assigns it to the SearchFilterGroupsType field.
func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilterGroupsType(v ProvisioningOperationEventSearchFilterGroupsType) {
	o.SearchFilterGroupsType = &v
}

// GetSortCriteriaItems returns the SortCriteriaItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningOperationEventSearchRequestModel) GetSortCriteriaItems() []ProvisioningOperationEventSortCriteriaRequestModel {
	if o == nil {
		var ret []ProvisioningOperationEventSortCriteriaRequestModel
		return ret
	}
	return o.SortCriteriaItems
}

// GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningOperationEventSearchRequestModel) GetSortCriteriaItemsOk() ([]ProvisioningOperationEventSortCriteriaRequestModel, bool) {
	if o == nil || IsNil(o.SortCriteriaItems) {
		return nil, false
	}
	return o.SortCriteriaItems, true
}

// HasSortCriteriaItems returns a boolean if a field has been set.
func (o *ProvisioningOperationEventSearchRequestModel) HasSortCriteriaItems() bool {
	if o != nil && IsNil(o.SortCriteriaItems) {
		return true
	}

	return false
}

// SetSortCriteriaItems gets a reference to the given []ProvisioningOperationEventSortCriteriaRequestModel and assigns it to the SortCriteriaItems field.
func (o *ProvisioningOperationEventSearchRequestModel) SetSortCriteriaItems(v []ProvisioningOperationEventSortCriteriaRequestModel) {
	o.SortCriteriaItems = v
}

func (o ProvisioningOperationEventSearchRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningOperationEventSearchRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BasicSearchString.IsSet() {
		toSerialize["BasicSearchString"] = o.BasicSearchString.Get()
	}
	if o.SearchFilters != nil {
		toSerialize["SearchFilters"] = o.SearchFilters
	}
	if o.SearchFilterGroups != nil {
		toSerialize["SearchFilterGroups"] = o.SearchFilterGroups
	}
	if !IsNil(o.SearchFilterGroupsType) {
		toSerialize["SearchFilterGroupsType"] = o.SearchFilterGroupsType
	}
	if o.SortCriteriaItems != nil {
		toSerialize["SortCriteriaItems"] = o.SortCriteriaItems
	}
	return toSerialize, nil
}

type NullableProvisioningOperationEventSearchRequestModel struct {
	value *ProvisioningOperationEventSearchRequestModel
	isSet bool
}

func (v NullableProvisioningOperationEventSearchRequestModel) Get() *ProvisioningOperationEventSearchRequestModel {
	return v.value
}

func (v *NullableProvisioningOperationEventSearchRequestModel) Set(val *ProvisioningOperationEventSearchRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningOperationEventSearchRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningOperationEventSearchRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningOperationEventSearchRequestModel(val *ProvisioningOperationEventSearchRequestModel) *NullableProvisioningOperationEventSearchRequestModel {
	return &NullableProvisioningOperationEventSearchRequestModel{value: val, isSet: true}
}

func (v NullableProvisioningOperationEventSearchRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningOperationEventSearchRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


