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

// checks if the DeliveryControllerDeleteScriptModelCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryControllerDeleteScriptModelCollection{}

// DeliveryControllerDeleteScriptModelCollection struct for DeliveryControllerDeleteScriptModelCollection
type DeliveryControllerDeleteScriptModelCollection struct {
	// List of items.
	Items []DeliveryControllerDeleteScriptModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewDeliveryControllerDeleteScriptModelCollection instantiates a new DeliveryControllerDeleteScriptModelCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryControllerDeleteScriptModelCollection(items []DeliveryControllerDeleteScriptModel) *DeliveryControllerDeleteScriptModelCollection {
	this := DeliveryControllerDeleteScriptModelCollection{}
	this.Items = items
	return &this
}

// NewDeliveryControllerDeleteScriptModelCollectionWithDefaults instantiates a new DeliveryControllerDeleteScriptModelCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryControllerDeleteScriptModelCollectionWithDefaults() *DeliveryControllerDeleteScriptModelCollection {
	this := DeliveryControllerDeleteScriptModelCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *DeliveryControllerDeleteScriptModelCollection) GetItems() []DeliveryControllerDeleteScriptModel {
	if o == nil {
		var ret []DeliveryControllerDeleteScriptModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DeliveryControllerDeleteScriptModelCollection) GetItemsOk() ([]DeliveryControllerDeleteScriptModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DeliveryControllerDeleteScriptModelCollection) SetItems(v []DeliveryControllerDeleteScriptModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryControllerDeleteScriptModelCollection) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryControllerDeleteScriptModelCollection) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *DeliveryControllerDeleteScriptModelCollection) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *DeliveryControllerDeleteScriptModelCollection) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}
// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *DeliveryControllerDeleteScriptModelCollection) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *DeliveryControllerDeleteScriptModelCollection) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryControllerDeleteScriptModelCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryControllerDeleteScriptModelCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *DeliveryControllerDeleteScriptModelCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *DeliveryControllerDeleteScriptModelCollection) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *DeliveryControllerDeleteScriptModelCollection) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *DeliveryControllerDeleteScriptModelCollection) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o DeliveryControllerDeleteScriptModelCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryControllerDeleteScriptModelCollection) ToMap() (map[string]interface{}, error) {
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

type NullableDeliveryControllerDeleteScriptModelCollection struct {
	value *DeliveryControllerDeleteScriptModelCollection
	isSet bool
}

func (v NullableDeliveryControllerDeleteScriptModelCollection) Get() *DeliveryControllerDeleteScriptModelCollection {
	return v.value
}

func (v *NullableDeliveryControllerDeleteScriptModelCollection) Set(val *DeliveryControllerDeleteScriptModelCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryControllerDeleteScriptModelCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryControllerDeleteScriptModelCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryControllerDeleteScriptModelCollection(val *DeliveryControllerDeleteScriptModelCollection) *NullableDeliveryControllerDeleteScriptModelCollection {
	return &NullableDeliveryControllerDeleteScriptModelCollection{value: val, isSet: true}
}

func (v NullableDeliveryControllerDeleteScriptModelCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryControllerDeleteScriptModelCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


