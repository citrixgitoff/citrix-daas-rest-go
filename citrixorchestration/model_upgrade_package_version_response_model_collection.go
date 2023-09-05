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

// checks if the UpgradePackageVersionResponseModelCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradePackageVersionResponseModelCollection{}

// UpgradePackageVersionResponseModelCollection struct for UpgradePackageVersionResponseModelCollection
type UpgradePackageVersionResponseModelCollection struct {
	// List of items.
	Items []UpgradePackageVersionResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken *string `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems *int32 `json:"TotalItems,omitempty"`
}

// NewUpgradePackageVersionResponseModelCollection instantiates a new UpgradePackageVersionResponseModelCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradePackageVersionResponseModelCollection(items []UpgradePackageVersionResponseModel) *UpgradePackageVersionResponseModelCollection {
	this := UpgradePackageVersionResponseModelCollection{}
	this.Items = items
	return &this
}

// NewUpgradePackageVersionResponseModelCollectionWithDefaults instantiates a new UpgradePackageVersionResponseModelCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradePackageVersionResponseModelCollectionWithDefaults() *UpgradePackageVersionResponseModelCollection {
	this := UpgradePackageVersionResponseModelCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *UpgradePackageVersionResponseModelCollection) GetItems() []UpgradePackageVersionResponseModel {
	if o == nil {
		var ret []UpgradePackageVersionResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UpgradePackageVersionResponseModelCollection) GetItemsOk() ([]UpgradePackageVersionResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *UpgradePackageVersionResponseModelCollection) SetItems(v []UpgradePackageVersionResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *UpgradePackageVersionResponseModelCollection) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePackageVersionResponseModelCollection) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *UpgradePackageVersionResponseModelCollection) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *UpgradePackageVersionResponseModelCollection) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *UpgradePackageVersionResponseModelCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePackageVersionResponseModelCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *UpgradePackageVersionResponseModelCollection) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *UpgradePackageVersionResponseModelCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o UpgradePackageVersionResponseModelCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradePackageVersionResponseModelCollection) ToMap() (map[string]interface{}, error) {
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

type NullableUpgradePackageVersionResponseModelCollection struct {
	value *UpgradePackageVersionResponseModelCollection
	isSet bool
}

func (v NullableUpgradePackageVersionResponseModelCollection) Get() *UpgradePackageVersionResponseModelCollection {
	return v.value
}

func (v *NullableUpgradePackageVersionResponseModelCollection) Set(val *UpgradePackageVersionResponseModelCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradePackageVersionResponseModelCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradePackageVersionResponseModelCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradePackageVersionResponseModelCollection(val *UpgradePackageVersionResponseModelCollection) *NullableUpgradePackageVersionResponseModelCollection {
	return &NullableUpgradePackageVersionResponseModelCollection{value: val, isSet: true}
}

func (v NullableUpgradePackageVersionResponseModelCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradePackageVersionResponseModelCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


