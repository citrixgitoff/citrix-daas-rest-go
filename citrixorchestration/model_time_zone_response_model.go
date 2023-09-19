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

// checks if the TimeZoneResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeZoneResponseModel{}

// TimeZoneResponseModel Time zone response details.
type TimeZoneResponseModel struct {
	// Human-readable name of the time zone, including the UTC offset of the standard (non-daylight-savings) time within the time zone.  For example, \"(UTC-10:00) Aleutian Islands\".   This is localized into the language requested by the caller.
	Name NullableString `json:"Name,omitempty"`
	// Human-readable name of the time zone, including the UTC offset of the standard (non-daylight-savings) time within the time zone.  For example, \"(UTC-10:00) Aleutian Islands\".   This is localized into the language requested by the caller.
	FullName NullableString `json:"FullName,omitempty"`
	// IANA identifier of the time zone within the primary territory covered by the time zone.   The  and  identifiers are usually, but not always, the same.  For example, the primary entry for \"India Standard Time\" is \"Asia/Calcutta\" while the canonical entry is \"Asia/Kolkata\".
	Primary NullableString `json:"Primary,omitempty"`
	// Canonical IANA identifier of the time zone.   The  and  identifiers are usually, but not always, the same.  For example, the primary entry for \"India Standard Time\" is \"Asia/Calcutta\" while the canonical entry is \"Asia/Kolkata\".
	Canonical NullableString `json:"Canonical,omitempty"`
	// Windows identifier of the time zone.   Note that although this may appear to be human-readable, it is **NOT** localized and therefore, should not be displayed in user interfaces.  Instead, use  or  properties for this purpose.
	WindowsId NullableString `json:"WindowsId,omitempty"`
	// UTC offset of the standard (non-daylight-savings) time within the time zone.
	UtcOffsetSeconds *int32 `json:"UtcOffsetSeconds,omitempty"`
	// String representation of UtcOffsetSeconds.  Always starts with `+` or `-`, followed by a two hours digits, `:`, and two minutes digits.  (e.g. `-05:00` / `+00:00`)
	UtcOffset NullableString `json:"UtcOffset,omitempty"`
	// List of IANA identifiers which can be used to represent this time zone.  May be non-exhaustive.
	TzdbIds []string `json:"TzdbIds,omitempty"`
}

// NewTimeZoneResponseModel instantiates a new TimeZoneResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeZoneResponseModel() *TimeZoneResponseModel {
	this := TimeZoneResponseModel{}
	return &this
}

// NewTimeZoneResponseModelWithDefaults instantiates a new TimeZoneResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeZoneResponseModelWithDefaults() *TimeZoneResponseModel {
	this := TimeZoneResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TimeZoneResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TimeZoneResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *TimeZoneResponseModel) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *TimeZoneResponseModel) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetFullName() {
	o.FullName.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetPrimary() string {
	if o == nil || IsNil(o.Primary.Get()) {
		var ret string
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasPrimary() bool {
	if o != nil && o.Primary.IsSet() {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given NullableString and assigns it to the Primary field.
func (o *TimeZoneResponseModel) SetPrimary(v string) {
	o.Primary.Set(&v)
}
// SetPrimaryNil sets the value for Primary to be an explicit nil
func (o *TimeZoneResponseModel) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetPrimary() {
	o.Primary.Unset()
}

// GetCanonical returns the Canonical field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetCanonical() string {
	if o == nil || IsNil(o.Canonical.Get()) {
		var ret string
		return ret
	}
	return *o.Canonical.Get()
}

// GetCanonicalOk returns a tuple with the Canonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Canonical.Get(), o.Canonical.IsSet()
}

// HasCanonical returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasCanonical() bool {
	if o != nil && o.Canonical.IsSet() {
		return true
	}

	return false
}

// SetCanonical gets a reference to the given NullableString and assigns it to the Canonical field.
func (o *TimeZoneResponseModel) SetCanonical(v string) {
	o.Canonical.Set(&v)
}
// SetCanonicalNil sets the value for Canonical to be an explicit nil
func (o *TimeZoneResponseModel) SetCanonicalNil() {
	o.Canonical.Set(nil)
}

// UnsetCanonical ensures that no value is present for Canonical, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetCanonical() {
	o.Canonical.Unset()
}

// GetWindowsId returns the WindowsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetWindowsId() string {
	if o == nil || IsNil(o.WindowsId.Get()) {
		var ret string
		return ret
	}
	return *o.WindowsId.Get()
}

// GetWindowsIdOk returns a tuple with the WindowsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetWindowsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsId.Get(), o.WindowsId.IsSet()
}

// HasWindowsId returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasWindowsId() bool {
	if o != nil && o.WindowsId.IsSet() {
		return true
	}

	return false
}

// SetWindowsId gets a reference to the given NullableString and assigns it to the WindowsId field.
func (o *TimeZoneResponseModel) SetWindowsId(v string) {
	o.WindowsId.Set(&v)
}
// SetWindowsIdNil sets the value for WindowsId to be an explicit nil
func (o *TimeZoneResponseModel) SetWindowsIdNil() {
	o.WindowsId.Set(nil)
}

// UnsetWindowsId ensures that no value is present for WindowsId, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetWindowsId() {
	o.WindowsId.Unset()
}

// GetUtcOffsetSeconds returns the UtcOffsetSeconds field value if set, zero value otherwise.
func (o *TimeZoneResponseModel) GetUtcOffsetSeconds() int32 {
	if o == nil || IsNil(o.UtcOffsetSeconds) {
		var ret int32
		return ret
	}
	return *o.UtcOffsetSeconds
}

// GetUtcOffsetSecondsOk returns a tuple with the UtcOffsetSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZoneResponseModel) GetUtcOffsetSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.UtcOffsetSeconds) {
		return nil, false
	}
	return o.UtcOffsetSeconds, true
}

// HasUtcOffsetSeconds returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasUtcOffsetSeconds() bool {
	if o != nil && !IsNil(o.UtcOffsetSeconds) {
		return true
	}

	return false
}

// SetUtcOffsetSeconds gets a reference to the given int32 and assigns it to the UtcOffsetSeconds field.
func (o *TimeZoneResponseModel) SetUtcOffsetSeconds(v int32) {
	o.UtcOffsetSeconds = &v
}

// GetUtcOffset returns the UtcOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetUtcOffset() string {
	if o == nil || IsNil(o.UtcOffset.Get()) {
		var ret string
		return ret
	}
	return *o.UtcOffset.Get()
}

// GetUtcOffsetOk returns a tuple with the UtcOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetUtcOffsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UtcOffset.Get(), o.UtcOffset.IsSet()
}

// HasUtcOffset returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasUtcOffset() bool {
	if o != nil && o.UtcOffset.IsSet() {
		return true
	}

	return false
}

// SetUtcOffset gets a reference to the given NullableString and assigns it to the UtcOffset field.
func (o *TimeZoneResponseModel) SetUtcOffset(v string) {
	o.UtcOffset.Set(&v)
}
// SetUtcOffsetNil sets the value for UtcOffset to be an explicit nil
func (o *TimeZoneResponseModel) SetUtcOffsetNil() {
	o.UtcOffset.Set(nil)
}

// UnsetUtcOffset ensures that no value is present for UtcOffset, not even an explicit nil
func (o *TimeZoneResponseModel) UnsetUtcOffset() {
	o.UtcOffset.Unset()
}

// GetTzdbIds returns the TzdbIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeZoneResponseModel) GetTzdbIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TzdbIds
}

// GetTzdbIdsOk returns a tuple with the TzdbIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeZoneResponseModel) GetTzdbIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TzdbIds) {
		return nil, false
	}
	return o.TzdbIds, true
}

// HasTzdbIds returns a boolean if a field has been set.
func (o *TimeZoneResponseModel) HasTzdbIds() bool {
	if o != nil && IsNil(o.TzdbIds) {
		return true
	}

	return false
}

// SetTzdbIds gets a reference to the given []string and assigns it to the TzdbIds field.
func (o *TimeZoneResponseModel) SetTzdbIds(v []string) {
	o.TzdbIds = v
}

func (o TimeZoneResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeZoneResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["FullName"] = o.FullName.Get()
	}
	if o.Primary.IsSet() {
		toSerialize["Primary"] = o.Primary.Get()
	}
	if o.Canonical.IsSet() {
		toSerialize["Canonical"] = o.Canonical.Get()
	}
	if o.WindowsId.IsSet() {
		toSerialize["WindowsId"] = o.WindowsId.Get()
	}
	if !IsNil(o.UtcOffsetSeconds) {
		toSerialize["UtcOffsetSeconds"] = o.UtcOffsetSeconds
	}
	if o.UtcOffset.IsSet() {
		toSerialize["UtcOffset"] = o.UtcOffset.Get()
	}
	if o.TzdbIds != nil {
		toSerialize["TzdbIds"] = o.TzdbIds
	}
	return toSerialize, nil
}

type NullableTimeZoneResponseModel struct {
	value *TimeZoneResponseModel
	isSet bool
}

func (v NullableTimeZoneResponseModel) Get() *TimeZoneResponseModel {
	return v.value
}

func (v *NullableTimeZoneResponseModel) Set(val *TimeZoneResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeZoneResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeZoneResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeZoneResponseModel(val *TimeZoneResponseModel) *NullableTimeZoneResponseModel {
	return &NullableTimeZoneResponseModel{value: val, isSet: true}
}

func (v NullableTimeZoneResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeZoneResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


