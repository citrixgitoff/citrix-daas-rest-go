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

// checks if the CollectionModelOfDeliveryControllerDeleteScriptModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionModelOfDeliveryControllerDeleteScriptModel{}

// CollectionModelOfDeliveryControllerDeleteScriptModel Response object for collections of items.
type CollectionModelOfDeliveryControllerDeleteScriptModel struct {
	// List of items.
	Items []DeliveryControllerDeleteScriptModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken *string `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems *int32 `json:"TotalItems,omitempty"`
}

// NewCollectionModelOfDeliveryControllerDeleteScriptModel instantiates a new CollectionModelOfDeliveryControllerDeleteScriptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionModelOfDeliveryControllerDeleteScriptModel(items []DeliveryControllerDeleteScriptModel) *CollectionModelOfDeliveryControllerDeleteScriptModel {
	this := CollectionModelOfDeliveryControllerDeleteScriptModel{}
	this.Items = items
	return &this
}

// NewCollectionModelOfDeliveryControllerDeleteScriptModelWithDefaults instantiates a new CollectionModelOfDeliveryControllerDeleteScriptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionModelOfDeliveryControllerDeleteScriptModelWithDefaults() *CollectionModelOfDeliveryControllerDeleteScriptModel {
	this := CollectionModelOfDeliveryControllerDeleteScriptModel{}
	return &this
}

// GetItems returns the Items field value
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetItems() []DeliveryControllerDeleteScriptModel {
	if o == nil {
		var ret []DeliveryControllerDeleteScriptModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetItemsOk() ([]DeliveryControllerDeleteScriptModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) SetItems(v []DeliveryControllerDeleteScriptModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionModelOfDeliveryControllerDeleteScriptModel) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o CollectionModelOfDeliveryControllerDeleteScriptModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionModelOfDeliveryControllerDeleteScriptModel) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionModelOfDeliveryControllerDeleteScriptModel struct {
	value *CollectionModelOfDeliveryControllerDeleteScriptModel
	isSet bool
}

func (v NullableCollectionModelOfDeliveryControllerDeleteScriptModel) Get() *CollectionModelOfDeliveryControllerDeleteScriptModel {
	return v.value
}

func (v *NullableCollectionModelOfDeliveryControllerDeleteScriptModel) Set(val *CollectionModelOfDeliveryControllerDeleteScriptModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionModelOfDeliveryControllerDeleteScriptModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionModelOfDeliveryControllerDeleteScriptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionModelOfDeliveryControllerDeleteScriptModel(val *CollectionModelOfDeliveryControllerDeleteScriptModel) *NullableCollectionModelOfDeliveryControllerDeleteScriptModel {
	return &NullableCollectionModelOfDeliveryControllerDeleteScriptModel{value: val, isSet: true}
}

func (v NullableCollectionModelOfDeliveryControllerDeleteScriptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionModelOfDeliveryControllerDeleteScriptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


