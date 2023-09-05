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

// checks if the LicensingAlertModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensingAlertModel{}

// LicensingAlertModel Licensing alert.
type LicensingAlertModel struct {
	// The priority of the licensing alert.
	Priority *int32 `json:"Priority,omitempty"`
	LicensingAlertLevel *LicensingAlertLevel `json:"LicensingAlertLevel,omitempty"`
	// The type code of the licensing alert (internal code used by licensing).
	TypeCode *string `json:"TypeCode,omitempty"`
	// The title of the alert.
	Title *string `json:"Title,omitempty"`
	// More detailed information about the licensing alert.
	Detail *string `json:"Detail,omitempty"`
	// The action the user should take to resolve the alert.
	Action *string `json:"Action,omitempty"`
}

// NewLicensingAlertModel instantiates a new LicensingAlertModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensingAlertModel() *LicensingAlertModel {
	this := LicensingAlertModel{}
	return &this
}

// NewLicensingAlertModelWithDefaults instantiates a new LicensingAlertModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensingAlertModelWithDefaults() *LicensingAlertModel {
	this := LicensingAlertModel{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *LicensingAlertModel) SetPriority(v int32) {
	o.Priority = &v
}

// GetLicensingAlertLevel returns the LicensingAlertLevel field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetLicensingAlertLevel() LicensingAlertLevel {
	if o == nil || IsNil(o.LicensingAlertLevel) {
		var ret LicensingAlertLevel
		return ret
	}
	return *o.LicensingAlertLevel
}

// GetLicensingAlertLevelOk returns a tuple with the LicensingAlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetLicensingAlertLevelOk() (*LicensingAlertLevel, bool) {
	if o == nil || IsNil(o.LicensingAlertLevel) {
		return nil, false
	}
	return o.LicensingAlertLevel, true
}

// HasLicensingAlertLevel returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasLicensingAlertLevel() bool {
	if o != nil && !IsNil(o.LicensingAlertLevel) {
		return true
	}

	return false
}

// SetLicensingAlertLevel gets a reference to the given LicensingAlertLevel and assigns it to the LicensingAlertLevel field.
func (o *LicensingAlertModel) SetLicensingAlertLevel(v LicensingAlertLevel) {
	o.LicensingAlertLevel = &v
}

// GetTypeCode returns the TypeCode field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetTypeCode() string {
	if o == nil || IsNil(o.TypeCode) {
		var ret string
		return ret
	}
	return *o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TypeCode) {
		return nil, false
	}
	return o.TypeCode, true
}

// HasTypeCode returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasTypeCode() bool {
	if o != nil && !IsNil(o.TypeCode) {
		return true
	}

	return false
}

// SetTypeCode gets a reference to the given string and assigns it to the TypeCode field.
func (o *LicensingAlertModel) SetTypeCode(v string) {
	o.TypeCode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LicensingAlertModel) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *LicensingAlertModel) SetDetail(v string) {
	o.Detail = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *LicensingAlertModel) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensingAlertModel) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *LicensingAlertModel) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *LicensingAlertModel) SetAction(v string) {
	o.Action = &v
}

func (o LicensingAlertModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensingAlertModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.LicensingAlertLevel) {
		toSerialize["LicensingAlertLevel"] = o.LicensingAlertLevel
	}
	if !IsNil(o.TypeCode) {
		toSerialize["TypeCode"] = o.TypeCode
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.Detail) {
		toSerialize["Detail"] = o.Detail
	}
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	return toSerialize, nil
}

type NullableLicensingAlertModel struct {
	value *LicensingAlertModel
	isSet bool
}

func (v NullableLicensingAlertModel) Get() *LicensingAlertModel {
	return v.value
}

func (v *NullableLicensingAlertModel) Set(val *LicensingAlertModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingAlertModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingAlertModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingAlertModel(val *LicensingAlertModel) *NullableLicensingAlertModel {
	return &NullableLicensingAlertModel{value: val, isSet: true}
}

func (v NullableLicensingAlertModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingAlertModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


