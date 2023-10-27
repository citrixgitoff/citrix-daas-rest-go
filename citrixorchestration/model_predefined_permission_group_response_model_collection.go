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

// checks if the PredefinedPermissionGroupResponseModelCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPermissionGroupResponseModelCollection{}

// PredefinedPermissionGroupResponseModelCollection struct for PredefinedPermissionGroupResponseModelCollection
type PredefinedPermissionGroupResponseModelCollection struct {
	// List of items.
	Items []PredefinedPermissionGroupResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewPredefinedPermissionGroupResponseModelCollection instantiates a new PredefinedPermissionGroupResponseModelCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPermissionGroupResponseModelCollection(items []PredefinedPermissionGroupResponseModel) *PredefinedPermissionGroupResponseModelCollection {
	this := PredefinedPermissionGroupResponseModelCollection{}
	this.Items = items
	return &this
}

// NewPredefinedPermissionGroupResponseModelCollectionWithDefaults instantiates a new PredefinedPermissionGroupResponseModelCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPermissionGroupResponseModelCollectionWithDefaults() *PredefinedPermissionGroupResponseModelCollection {
	this := PredefinedPermissionGroupResponseModelCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *PredefinedPermissionGroupResponseModelCollection) GetItems() []PredefinedPermissionGroupResponseModel {
	if o == nil {
		var ret []PredefinedPermissionGroupResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionGroupResponseModelCollection) GetItemsOk() ([]PredefinedPermissionGroupResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PredefinedPermissionGroupResponseModelCollection) SetItems(v []PredefinedPermissionGroupResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredefinedPermissionGroupResponseModelCollection) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredefinedPermissionGroupResponseModelCollection) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *PredefinedPermissionGroupResponseModelCollection) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *PredefinedPermissionGroupResponseModelCollection) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}
// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *PredefinedPermissionGroupResponseModelCollection) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *PredefinedPermissionGroupResponseModelCollection) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredefinedPermissionGroupResponseModelCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredefinedPermissionGroupResponseModelCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PredefinedPermissionGroupResponseModelCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *PredefinedPermissionGroupResponseModelCollection) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *PredefinedPermissionGroupResponseModelCollection) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *PredefinedPermissionGroupResponseModelCollection) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o PredefinedPermissionGroupResponseModelCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPermissionGroupResponseModelCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Items"] = o.Items
	if o.ContinuationToken.IsSet() {
		toSerialize["ContinuationToken"] = o.ContinuationToken.Get()
	}
	if o.TotalItems.IsSet() {
		toSerialize["TotalItems"] = o.TotalItems.Get()
	}
	return toSerialize, nil
}

type NullablePredefinedPermissionGroupResponseModelCollection struct {
	value *PredefinedPermissionGroupResponseModelCollection
	isSet bool
}

func (v NullablePredefinedPermissionGroupResponseModelCollection) Get() *PredefinedPermissionGroupResponseModelCollection {
	return v.value
}

func (v *NullablePredefinedPermissionGroupResponseModelCollection) Set(val *PredefinedPermissionGroupResponseModelCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPermissionGroupResponseModelCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPermissionGroupResponseModelCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPermissionGroupResponseModelCollection(val *PredefinedPermissionGroupResponseModelCollection) *NullablePredefinedPermissionGroupResponseModelCollection {
	return &NullablePredefinedPermissionGroupResponseModelCollection{value: val, isSet: true}
}

func (v NullablePredefinedPermissionGroupResponseModelCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPermissionGroupResponseModelCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


