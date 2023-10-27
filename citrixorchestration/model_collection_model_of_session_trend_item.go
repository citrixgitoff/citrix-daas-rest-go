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

// checks if the CollectionModelOfSessionTrendItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionModelOfSessionTrendItem{}

// CollectionModelOfSessionTrendItem Response object for collections of items.
type CollectionModelOfSessionTrendItem struct {
	// List of items.
	Items []SessionTrendItem `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewCollectionModelOfSessionTrendItem instantiates a new CollectionModelOfSessionTrendItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionModelOfSessionTrendItem(items []SessionTrendItem) *CollectionModelOfSessionTrendItem {
	this := CollectionModelOfSessionTrendItem{}
	this.Items = items
	return &this
}

// NewCollectionModelOfSessionTrendItemWithDefaults instantiates a new CollectionModelOfSessionTrendItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionModelOfSessionTrendItemWithDefaults() *CollectionModelOfSessionTrendItem {
	this := CollectionModelOfSessionTrendItem{}
	return &this
}

// GetItems returns the Items field value
func (o *CollectionModelOfSessionTrendItem) GetItems() []SessionTrendItem {
	if o == nil {
		var ret []SessionTrendItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CollectionModelOfSessionTrendItem) GetItemsOk() ([]SessionTrendItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CollectionModelOfSessionTrendItem) SetItems(v []SessionTrendItem) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfSessionTrendItem) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfSessionTrendItem) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionModelOfSessionTrendItem) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *CollectionModelOfSessionTrendItem) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}
// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *CollectionModelOfSessionTrendItem) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *CollectionModelOfSessionTrendItem) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfSessionTrendItem) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfSessionTrendItem) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionModelOfSessionTrendItem) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *CollectionModelOfSessionTrendItem) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *CollectionModelOfSessionTrendItem) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *CollectionModelOfSessionTrendItem) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o CollectionModelOfSessionTrendItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionModelOfSessionTrendItem) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionModelOfSessionTrendItem struct {
	value *CollectionModelOfSessionTrendItem
	isSet bool
}

func (v NullableCollectionModelOfSessionTrendItem) Get() *CollectionModelOfSessionTrendItem {
	return v.value
}

func (v *NullableCollectionModelOfSessionTrendItem) Set(val *CollectionModelOfSessionTrendItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionModelOfSessionTrendItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionModelOfSessionTrendItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionModelOfSessionTrendItem(val *CollectionModelOfSessionTrendItem) *NullableCollectionModelOfSessionTrendItem {
	return &NullableCollectionModelOfSessionTrendItem{value: val, isSet: true}
}

func (v NullableCollectionModelOfSessionTrendItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionModelOfSessionTrendItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


