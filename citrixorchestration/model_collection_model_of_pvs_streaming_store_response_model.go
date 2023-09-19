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

// checks if the CollectionModelOfPvsStreamingStoreResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionModelOfPvsStreamingStoreResponseModel{}

// CollectionModelOfPvsStreamingStoreResponseModel Response object for collections of items.
type CollectionModelOfPvsStreamingStoreResponseModel struct {
	// List of items.
	Items []PvsStreamingStoreResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewCollectionModelOfPvsStreamingStoreResponseModel instantiates a new CollectionModelOfPvsStreamingStoreResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionModelOfPvsStreamingStoreResponseModel(items []PvsStreamingStoreResponseModel) *CollectionModelOfPvsStreamingStoreResponseModel {
	this := CollectionModelOfPvsStreamingStoreResponseModel{}
	this.Items = items
	return &this
}

// NewCollectionModelOfPvsStreamingStoreResponseModelWithDefaults instantiates a new CollectionModelOfPvsStreamingStoreResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionModelOfPvsStreamingStoreResponseModelWithDefaults() *CollectionModelOfPvsStreamingStoreResponseModel {
	this := CollectionModelOfPvsStreamingStoreResponseModel{}
	return &this
}

// GetItems returns the Items field value
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetItems() []PvsStreamingStoreResponseModel {
	if o == nil {
		var ret []PvsStreamingStoreResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetItemsOk() ([]PvsStreamingStoreResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CollectionModelOfPvsStreamingStoreResponseModel) SetItems(v []PvsStreamingStoreResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionModelOfPvsStreamingStoreResponseModel) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *CollectionModelOfPvsStreamingStoreResponseModel) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}
// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *CollectionModelOfPvsStreamingStoreResponseModel) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *CollectionModelOfPvsStreamingStoreResponseModel) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfPvsStreamingStoreResponseModel) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionModelOfPvsStreamingStoreResponseModel) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *CollectionModelOfPvsStreamingStoreResponseModel) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *CollectionModelOfPvsStreamingStoreResponseModel) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *CollectionModelOfPvsStreamingStoreResponseModel) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o CollectionModelOfPvsStreamingStoreResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionModelOfPvsStreamingStoreResponseModel) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionModelOfPvsStreamingStoreResponseModel struct {
	value *CollectionModelOfPvsStreamingStoreResponseModel
	isSet bool
}

func (v NullableCollectionModelOfPvsStreamingStoreResponseModel) Get() *CollectionModelOfPvsStreamingStoreResponseModel {
	return v.value
}

func (v *NullableCollectionModelOfPvsStreamingStoreResponseModel) Set(val *CollectionModelOfPvsStreamingStoreResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionModelOfPvsStreamingStoreResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionModelOfPvsStreamingStoreResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionModelOfPvsStreamingStoreResponseModel(val *CollectionModelOfPvsStreamingStoreResponseModel) *NullableCollectionModelOfPvsStreamingStoreResponseModel {
	return &NullableCollectionModelOfPvsStreamingStoreResponseModel{value: val, isSet: true}
}

func (v NullableCollectionModelOfPvsStreamingStoreResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionModelOfPvsStreamingStoreResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


