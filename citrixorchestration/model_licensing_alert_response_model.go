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

// checks if the LicensingAlertResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensingAlertResponseModel{}

// LicensingAlertResponseModel Licensing alert.
type LicensingAlertResponseModel struct {
	// The priority of the licensing alert.
	Priority *int32 `json:"Priority,omitempty"`
	AlertLevel *LicensingAlertLevel2 `json:"AlertLevel,omitempty"`
	// The type code of the licensing alert (internal code used by licensing).
	TypeCode NullableString `json:"TypeCode,omitempty"`
	// The title of the alert.
	Title NullableString `json:"Title,omitempty"`
	// More detailed information about the licensing alert.
	Detail NullableString `json:"Detail,omitempty"`
	// The action the user should take to resolve the alert.
	Action NullableString `json:"Action,omitempty"`
}

// NewLicensingAlertResponseModel instantiates a new LicensingAlertResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensingAlertResponseModel() *LicensingAlertResponseModel {
	this := LicensingAlertResponseModel{}
	return &this
}

// NewLicensingAlertResponseModelWithDefaults instantiates a new LicensingAlertResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensingAlertResponseModelWithDefaults() *LicensingAlertResponseModel {
	this := LicensingAlertResponseModel{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LicensingAlertResponseModel) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertResponseModel) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *LicensingAlertResponseModel) SetPriority(v int32) {
	o.Priority = &v
}

// GetAlertLevel returns the AlertLevel field value if set, zero value otherwise.
func (o *LicensingAlertResponseModel) GetAlertLevel() LicensingAlertLevel2 {
	if o == nil || IsNil(o.AlertLevel) {
		var ret LicensingAlertLevel2
		return ret
	}
	return *o.AlertLevel
}

// GetAlertLevelOk returns a tuple with the AlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertResponseModel) GetAlertLevelOk() (*LicensingAlertLevel2, bool) {
	if o == nil || IsNil(o.AlertLevel) {
		return nil, false
	}
	return o.AlertLevel, true
}

// HasAlertLevel returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasAlertLevel() bool {
	if o != nil && !IsNil(o.AlertLevel) {
		return true
	}

	return false
}

// SetAlertLevel gets a reference to the given LicensingAlertLevel2 and assigns it to the AlertLevel field.
func (o *LicensingAlertResponseModel) SetAlertLevel(v LicensingAlertLevel2) {
	o.AlertLevel = &v
}

// GetTypeCode returns the TypeCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensingAlertResponseModel) GetTypeCode() string {
	if o == nil || IsNil(o.TypeCode.Get()) {
		var ret string
		return ret
	}
	return *o.TypeCode.Get()
}

// GetTypeCodeOk returns a tuple with the TypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensingAlertResponseModel) GetTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeCode.Get(), o.TypeCode.IsSet()
}

// HasTypeCode returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasTypeCode() bool {
	if o != nil && o.TypeCode.IsSet() {
		return true
	}

	return false
}

// SetTypeCode gets a reference to the given NullableString and assigns it to the TypeCode field.
func (o *LicensingAlertResponseModel) SetTypeCode(v string) {
	o.TypeCode.Set(&v)
}
// SetTypeCodeNil sets the value for TypeCode to be an explicit nil
func (o *LicensingAlertResponseModel) SetTypeCodeNil() {
	o.TypeCode.Set(nil)
}

// UnsetTypeCode ensures that no value is present for TypeCode, not even an explicit nil
func (o *LicensingAlertResponseModel) UnsetTypeCode() {
	o.TypeCode.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensingAlertResponseModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensingAlertResponseModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LicensingAlertResponseModel) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LicensingAlertResponseModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LicensingAlertResponseModel) UnsetTitle() {
	o.Title.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensingAlertResponseModel) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensingAlertResponseModel) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *LicensingAlertResponseModel) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *LicensingAlertResponseModel) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *LicensingAlertResponseModel) UnsetDetail() {
	o.Detail.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensingAlertResponseModel) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensingAlertResponseModel) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *LicensingAlertResponseModel) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *LicensingAlertResponseModel) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *LicensingAlertResponseModel) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *LicensingAlertResponseModel) UnsetAction() {
	o.Action.Unset()
}

func (o LicensingAlertResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensingAlertResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.AlertLevel) {
		toSerialize["AlertLevel"] = o.AlertLevel
	}
	if o.TypeCode.IsSet() {
		toSerialize["TypeCode"] = o.TypeCode.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["Detail"] = o.Detail.Get()
	}
	if o.Action.IsSet() {
		toSerialize["Action"] = o.Action.Get()
	}
	return toSerialize, nil
}

type NullableLicensingAlertResponseModel struct {
	value *LicensingAlertResponseModel
	isSet bool
}

func (v NullableLicensingAlertResponseModel) Get() *LicensingAlertResponseModel {
	return v.value
}

func (v *NullableLicensingAlertResponseModel) Set(val *LicensingAlertResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingAlertResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingAlertResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingAlertResponseModel(val *LicensingAlertResponseModel) *NullableLicensingAlertResponseModel {
	return &NullableLicensingAlertResponseModel{value: val, isSet: true}
}

func (v NullableLicensingAlertResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingAlertResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


