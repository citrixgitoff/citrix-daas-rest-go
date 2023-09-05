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

// checks if the HypervisorsAndResourcePoolsResponseModelCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorsAndResourcePoolsResponseModelCollection{}

// HypervisorsAndResourcePoolsResponseModelCollection struct for HypervisorsAndResourcePoolsResponseModelCollection
type HypervisorsAndResourcePoolsResponseModelCollection struct {
	// List of items.
	Items []HypervisorsAndResourcePoolsResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken *string `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems *int32 `json:"TotalItems,omitempty"`
}

// NewHypervisorsAndResourcePoolsResponseModelCollection instantiates a new HypervisorsAndResourcePoolsResponseModelCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorsAndResourcePoolsResponseModelCollection(items []HypervisorsAndResourcePoolsResponseModel) *HypervisorsAndResourcePoolsResponseModelCollection {
	this := HypervisorsAndResourcePoolsResponseModelCollection{}
	this.Items = items
	return &this
}

// NewHypervisorsAndResourcePoolsResponseModelCollectionWithDefaults instantiates a new HypervisorsAndResourcePoolsResponseModelCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorsAndResourcePoolsResponseModelCollectionWithDefaults() *HypervisorsAndResourcePoolsResponseModelCollection {
	this := HypervisorsAndResourcePoolsResponseModelCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetItems() []HypervisorsAndResourcePoolsResponseModel {
	if o == nil {
		var ret []HypervisorsAndResourcePoolsResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetItemsOk() ([]HypervisorsAndResourcePoolsResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetItems(v []HypervisorsAndResourcePoolsResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o HypervisorsAndResourcePoolsResponseModelCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorsAndResourcePoolsResponseModelCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Items"] = o.Items
	if !IsNil(o.ContinuationToken) {
		toSerialize["ContinuationToken"] = o.ContinuationToken
	}
	if !IsNil(o.TotalItems) {
		toSerialize["TotalItems"] = o.TotalItems
	}
	return toSerialize, nil
}

type NullableHypervisorsAndResourcePoolsResponseModelCollection struct {
	value *HypervisorsAndResourcePoolsResponseModelCollection
	isSet bool
}

func (v NullableHypervisorsAndResourcePoolsResponseModelCollection) Get() *HypervisorsAndResourcePoolsResponseModelCollection {
	return v.value
}

func (v *NullableHypervisorsAndResourcePoolsResponseModelCollection) Set(val *HypervisorsAndResourcePoolsResponseModelCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorsAndResourcePoolsResponseModelCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorsAndResourcePoolsResponseModelCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorsAndResourcePoolsResponseModelCollection(val *HypervisorsAndResourcePoolsResponseModelCollection) *NullableHypervisorsAndResourcePoolsResponseModelCollection {
	return &NullableHypervisorsAndResourcePoolsResponseModelCollection{value: val, isSet: true}
}

func (v NullableHypervisorsAndResourcePoolsResponseModelCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorsAndResourcePoolsResponseModelCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


