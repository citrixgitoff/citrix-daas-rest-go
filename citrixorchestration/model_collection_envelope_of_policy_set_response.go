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

// checks if the CollectionEnvelopeOfPolicySetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionEnvelopeOfPolicySetResponse{}

// CollectionEnvelopeOfPolicySetResponse Response object for collections of items.
type CollectionEnvelopeOfPolicySetResponse struct {
	// Items in the collection.
	Items []PolicySetResponse `json:"items,omitempty"`
	// If present, indicates to the caller that the query was not complete, and it should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"continuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of items returned, if there is a ContinuationToken. This is returned only in the response to $search APIs.
	TotalItems NullableInt32 `json:"totalItems,omitempty"`
}

// NewCollectionEnvelopeOfPolicySetResponse instantiates a new CollectionEnvelopeOfPolicySetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionEnvelopeOfPolicySetResponse() *CollectionEnvelopeOfPolicySetResponse {
	this := CollectionEnvelopeOfPolicySetResponse{}
	return &this
}

// NewCollectionEnvelopeOfPolicySetResponseWithDefaults instantiates a new CollectionEnvelopeOfPolicySetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionEnvelopeOfPolicySetResponseWithDefaults() *CollectionEnvelopeOfPolicySetResponse {
	this := CollectionEnvelopeOfPolicySetResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionEnvelopeOfPolicySetResponse) GetItems() []PolicySetResponse {
	if o == nil {
		var ret []PolicySetResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionEnvelopeOfPolicySetResponse) GetItemsOk() ([]PolicySetResponse, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CollectionEnvelopeOfPolicySetResponse) HasItems() bool {
	if o != nil && IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PolicySetResponse and assigns it to the Items field.
func (o *CollectionEnvelopeOfPolicySetResponse) SetItems(v []PolicySetResponse) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionEnvelopeOfPolicySetResponse) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionEnvelopeOfPolicySetResponse) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionEnvelopeOfPolicySetResponse) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *CollectionEnvelopeOfPolicySetResponse) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}
// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *CollectionEnvelopeOfPolicySetResponse) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *CollectionEnvelopeOfPolicySetResponse) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionEnvelopeOfPolicySetResponse) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionEnvelopeOfPolicySetResponse) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionEnvelopeOfPolicySetResponse) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *CollectionEnvelopeOfPolicySetResponse) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *CollectionEnvelopeOfPolicySetResponse) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *CollectionEnvelopeOfPolicySetResponse) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o CollectionEnvelopeOfPolicySetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionEnvelopeOfPolicySetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ContinuationToken.IsSet() {
		toSerialize["continuationToken"] = o.ContinuationToken.Get()
	}
	if o.TotalItems.IsSet() {
		toSerialize["totalItems"] = o.TotalItems.Get()
	}
	return toSerialize, nil
}

type NullableCollectionEnvelopeOfPolicySetResponse struct {
	value *CollectionEnvelopeOfPolicySetResponse
	isSet bool
}

func (v NullableCollectionEnvelopeOfPolicySetResponse) Get() *CollectionEnvelopeOfPolicySetResponse {
	return v.value
}

func (v *NullableCollectionEnvelopeOfPolicySetResponse) Set(val *CollectionEnvelopeOfPolicySetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionEnvelopeOfPolicySetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionEnvelopeOfPolicySetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionEnvelopeOfPolicySetResponse(val *CollectionEnvelopeOfPolicySetResponse) *NullableCollectionEnvelopeOfPolicySetResponse {
	return &NullableCollectionEnvelopeOfPolicySetResponse{value: val, isSet: true}
}

func (v NullableCollectionEnvelopeOfPolicySetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionEnvelopeOfPolicySetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


