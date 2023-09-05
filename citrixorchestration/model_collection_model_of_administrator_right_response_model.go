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

// checks if the CollectionModelOfAdministratorRightResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionModelOfAdministratorRightResponseModel{}

// CollectionModelOfAdministratorRightResponseModel Response object for collections of items.
type CollectionModelOfAdministratorRightResponseModel struct {
	// List of items.
	Items []AdministratorRightResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken *string `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems *int32 `json:"TotalItems,omitempty"`
}

// NewCollectionModelOfAdministratorRightResponseModel instantiates a new CollectionModelOfAdministratorRightResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionModelOfAdministratorRightResponseModel(items []AdministratorRightResponseModel) *CollectionModelOfAdministratorRightResponseModel {
	this := CollectionModelOfAdministratorRightResponseModel{}
	this.Items = items
	return &this
}

// NewCollectionModelOfAdministratorRightResponseModelWithDefaults instantiates a new CollectionModelOfAdministratorRightResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionModelOfAdministratorRightResponseModelWithDefaults() *CollectionModelOfAdministratorRightResponseModel {
	this := CollectionModelOfAdministratorRightResponseModel{}
	return &this
}

// GetItems returns the Items field value
func (o *CollectionModelOfAdministratorRightResponseModel) GetItems() []AdministratorRightResponseModel {
	if o == nil {
		var ret []AdministratorRightResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CollectionModelOfAdministratorRightResponseModel) GetItemsOk() ([]AdministratorRightResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CollectionModelOfAdministratorRightResponseModel) SetItems(v []AdministratorRightResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *CollectionModelOfAdministratorRightResponseModel) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionModelOfAdministratorRightResponseModel) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionModelOfAdministratorRightResponseModel) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *CollectionModelOfAdministratorRightResponseModel) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionModelOfAdministratorRightResponseModel) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionModelOfAdministratorRightResponseModel) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionModelOfAdministratorRightResponseModel) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionModelOfAdministratorRightResponseModel) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o CollectionModelOfAdministratorRightResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionModelOfAdministratorRightResponseModel) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionModelOfAdministratorRightResponseModel struct {
	value *CollectionModelOfAdministratorRightResponseModel
	isSet bool
}

func (v NullableCollectionModelOfAdministratorRightResponseModel) Get() *CollectionModelOfAdministratorRightResponseModel {
	return v.value
}

func (v *NullableCollectionModelOfAdministratorRightResponseModel) Set(val *CollectionModelOfAdministratorRightResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionModelOfAdministratorRightResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionModelOfAdministratorRightResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionModelOfAdministratorRightResponseModel(val *CollectionModelOfAdministratorRightResponseModel) *NullableCollectionModelOfAdministratorRightResponseModel {
	return &NullableCollectionModelOfAdministratorRightResponseModel{value: val, isSet: true}
}

func (v NullableCollectionModelOfAdministratorRightResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionModelOfAdministratorRightResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


