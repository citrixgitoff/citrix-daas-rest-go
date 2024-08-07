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

// checks if the ResourcePriceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcePriceResponseModel{}

// ResourcePriceResponseModel The response object indicating the price of a resource.
type ResourcePriceResponseModel struct {
	// Gets or sets the location.
	Location NullableString `json:"Location,omitempty"`
	EffectivePrice *UnitPriceResponseModel `json:"EffectivePrice,omitempty"`
	RetailPrice *UnitPriceResponseModel `json:"RetailPrice,omitempty"`
	// Gets or sets the savings plans.
	SavingPlans []SavingPlanResponseModel `json:"SavingPlans,omitempty"`
	// Gets or sets the reservations.
	Reservations []ReservationResponseModel `json:"Reservations,omitempty"`
	// Gets or sets the additional data.
	AdditionalData map[string]string `json:"AdditionalData,omitempty"`
}

// NewResourcePriceResponseModel instantiates a new ResourcePriceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePriceResponseModel() *ResourcePriceResponseModel {
	this := ResourcePriceResponseModel{}
	return &this
}

// NewResourcePriceResponseModelWithDefaults instantiates a new ResourcePriceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePriceResponseModelWithDefaults() *ResourcePriceResponseModel {
	this := ResourcePriceResponseModel{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcePriceResponseModel) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcePriceResponseModel) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ResourcePriceResponseModel) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ResourcePriceResponseModel) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ResourcePriceResponseModel) UnsetLocation() {
	o.Location.Unset()
}

// GetEffectivePrice returns the EffectivePrice field value if set, zero value otherwise.
func (o *ResourcePriceResponseModel) GetEffectivePrice() UnitPriceResponseModel {
	if o == nil || IsNil(o.EffectivePrice) {
		var ret UnitPriceResponseModel
		return ret
	}
	return *o.EffectivePrice
}

// GetEffectivePriceOk returns a tuple with the EffectivePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePriceResponseModel) GetEffectivePriceOk() (*UnitPriceResponseModel, bool) {
	if o == nil || IsNil(o.EffectivePrice) {
		return nil, false
	}
	return o.EffectivePrice, true
}

// HasEffectivePrice returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasEffectivePrice() bool {
	if o != nil && !IsNil(o.EffectivePrice) {
		return true
	}

	return false
}

// SetEffectivePrice gets a reference to the given UnitPriceResponseModel and assigns it to the EffectivePrice field.
func (o *ResourcePriceResponseModel) SetEffectivePrice(v UnitPriceResponseModel) {
	o.EffectivePrice = &v
}

// GetRetailPrice returns the RetailPrice field value if set, zero value otherwise.
func (o *ResourcePriceResponseModel) GetRetailPrice() UnitPriceResponseModel {
	if o == nil || IsNil(o.RetailPrice) {
		var ret UnitPriceResponseModel
		return ret
	}
	return *o.RetailPrice
}

// GetRetailPriceOk returns a tuple with the RetailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePriceResponseModel) GetRetailPriceOk() (*UnitPriceResponseModel, bool) {
	if o == nil || IsNil(o.RetailPrice) {
		return nil, false
	}
	return o.RetailPrice, true
}

// HasRetailPrice returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasRetailPrice() bool {
	if o != nil && !IsNil(o.RetailPrice) {
		return true
	}

	return false
}

// SetRetailPrice gets a reference to the given UnitPriceResponseModel and assigns it to the RetailPrice field.
func (o *ResourcePriceResponseModel) SetRetailPrice(v UnitPriceResponseModel) {
	o.RetailPrice = &v
}

// GetSavingPlans returns the SavingPlans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcePriceResponseModel) GetSavingPlans() []SavingPlanResponseModel {
	if o == nil {
		var ret []SavingPlanResponseModel
		return ret
	}
	return o.SavingPlans
}

// GetSavingPlansOk returns a tuple with the SavingPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcePriceResponseModel) GetSavingPlansOk() ([]SavingPlanResponseModel, bool) {
	if o == nil || IsNil(o.SavingPlans) {
		return nil, false
	}
	return o.SavingPlans, true
}

// HasSavingPlans returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasSavingPlans() bool {
	if o != nil && IsNil(o.SavingPlans) {
		return true
	}

	return false
}

// SetSavingPlans gets a reference to the given []SavingPlanResponseModel and assigns it to the SavingPlans field.
func (o *ResourcePriceResponseModel) SetSavingPlans(v []SavingPlanResponseModel) {
	o.SavingPlans = v
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcePriceResponseModel) GetReservations() []ReservationResponseModel {
	if o == nil {
		var ret []ReservationResponseModel
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcePriceResponseModel) GetReservationsOk() ([]ReservationResponseModel, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasReservations() bool {
	if o != nil && IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []ReservationResponseModel and assigns it to the Reservations field.
func (o *ResourcePriceResponseModel) SetReservations(v []ReservationResponseModel) {
	o.Reservations = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcePriceResponseModel) GetAdditionalData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcePriceResponseModel) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return &o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ResourcePriceResponseModel) HasAdditionalData() bool {
	if o != nil && IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ResourcePriceResponseModel) SetAdditionalData(v map[string]string) {
	o.AdditionalData = v
}

func (o ResourcePriceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePriceResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["Location"] = o.Location.Get()
	}
	if !IsNil(o.EffectivePrice) {
		toSerialize["EffectivePrice"] = o.EffectivePrice
	}
	if !IsNil(o.RetailPrice) {
		toSerialize["RetailPrice"] = o.RetailPrice
	}
	if o.SavingPlans != nil {
		toSerialize["SavingPlans"] = o.SavingPlans
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}
	if o.AdditionalData != nil {
		toSerialize["AdditionalData"] = o.AdditionalData
	}
	return toSerialize, nil
}

type NullableResourcePriceResponseModel struct {
	value *ResourcePriceResponseModel
	isSet bool
}

func (v NullableResourcePriceResponseModel) Get() *ResourcePriceResponseModel {
	return v.value
}

func (v *NullableResourcePriceResponseModel) Set(val *ResourcePriceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePriceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePriceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePriceResponseModel(val *ResourcePriceResponseModel) *NullableResourcePriceResponseModel {
	return &NullableResourcePriceResponseModel{value: val, isSet: true}
}

func (v NullableResourcePriceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePriceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


