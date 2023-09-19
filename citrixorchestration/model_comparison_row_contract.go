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

// checks if the ComparisonRowContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComparisonRowContract{}

// ComparisonRowContract A row in a comparison report.
type ComparisonRowContract struct {
	// Can be either a category name or setting name.
	SettingName NullableString `json:"SettingName,omitempty"`
	// Indicate if SettingName is a setting name or category name.
	IsSetting *bool `json:"IsSetting,omitempty"`
	// Valid only when IsSetting is false. Used to indicate if there is any difference among the settings in the category.
	IsDifferent *bool `json:"IsDifferent,omitempty"`
	// There are Targets.Length number of items in the list, stored in the same order as the items in Targets. Value is true if the setting is configured in the target. Value is true if SettingName is a category name and at least one of the settings under the category has IsInUse set to true.
	IsInUse []bool `json:"IsInUse,omitempty"`
	// There are Targets.Length number of items in the list, stored in the same order as the items in Targets. Value should not be used if the corresponding IsInUse value is false. If the IsInUse value is false, a - should be displayed if IsSetting is true. Don't use the values here if IsSetting is false.
	Values []string `json:"Values,omitempty"`
	// For internal use to keep track of setting category.
	Category NullableString `json:"Category,omitempty"`
	// The default value as specified in the GPFX file for the setting.
	DefaultValue NullableString `json:"DefaultValue,omitempty"`
}

// NewComparisonRowContract instantiates a new ComparisonRowContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComparisonRowContract() *ComparisonRowContract {
	this := ComparisonRowContract{}
	return &this
}

// NewComparisonRowContractWithDefaults instantiates a new ComparisonRowContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComparisonRowContractWithDefaults() *ComparisonRowContract {
	this := ComparisonRowContract{}
	return &this
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComparisonRowContract) GetSettingName() string {
	if o == nil || IsNil(o.SettingName.Get()) {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComparisonRowContract) GetSettingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *ComparisonRowContract) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *ComparisonRowContract) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *ComparisonRowContract) UnsetSettingName() {
	o.SettingName.Unset()
}

// GetIsSetting returns the IsSetting field value if set, zero value otherwise.
func (o *ComparisonRowContract) GetIsSetting() bool {
	if o == nil || IsNil(o.IsSetting) {
		var ret bool
		return ret
	}
	return *o.IsSetting
}

// GetIsSettingOk returns a tuple with the IsSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonRowContract) GetIsSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSetting) {
		return nil, false
	}
	return o.IsSetting, true
}

// HasIsSetting returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasIsSetting() bool {
	if o != nil && !IsNil(o.IsSetting) {
		return true
	}

	return false
}

// SetIsSetting gets a reference to the given bool and assigns it to the IsSetting field.
func (o *ComparisonRowContract) SetIsSetting(v bool) {
	o.IsSetting = &v
}

// GetIsDifferent returns the IsDifferent field value if set, zero value otherwise.
func (o *ComparisonRowContract) GetIsDifferent() bool {
	if o == nil || IsNil(o.IsDifferent) {
		var ret bool
		return ret
	}
	return *o.IsDifferent
}

// GetIsDifferentOk returns a tuple with the IsDifferent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonRowContract) GetIsDifferentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDifferent) {
		return nil, false
	}
	return o.IsDifferent, true
}

// HasIsDifferent returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasIsDifferent() bool {
	if o != nil && !IsNil(o.IsDifferent) {
		return true
	}

	return false
}

// SetIsDifferent gets a reference to the given bool and assigns it to the IsDifferent field.
func (o *ComparisonRowContract) SetIsDifferent(v bool) {
	o.IsDifferent = &v
}

// GetIsInUse returns the IsInUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComparisonRowContract) GetIsInUse() []bool {
	if o == nil {
		var ret []bool
		return ret
	}
	return o.IsInUse
}

// GetIsInUseOk returns a tuple with the IsInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComparisonRowContract) GetIsInUseOk() ([]bool, bool) {
	if o == nil || IsNil(o.IsInUse) {
		return nil, false
	}
	return o.IsInUse, true
}

// HasIsInUse returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasIsInUse() bool {
	if o != nil && IsNil(o.IsInUse) {
		return true
	}

	return false
}

// SetIsInUse gets a reference to the given []bool and assigns it to the IsInUse field.
func (o *ComparisonRowContract) SetIsInUse(v []bool) {
	o.IsInUse = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComparisonRowContract) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComparisonRowContract) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasValues() bool {
	if o != nil && IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ComparisonRowContract) SetValues(v []string) {
	o.Values = v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComparisonRowContract) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComparisonRowContract) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ComparisonRowContract) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ComparisonRowContract) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ComparisonRowContract) UnsetCategory() {
	o.Category.Unset()
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComparisonRowContract) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComparisonRowContract) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ComparisonRowContract) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *ComparisonRowContract) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *ComparisonRowContract) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *ComparisonRowContract) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

func (o ComparisonRowContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComparisonRowContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SettingName.IsSet() {
		toSerialize["SettingName"] = o.SettingName.Get()
	}
	if !IsNil(o.IsSetting) {
		toSerialize["IsSetting"] = o.IsSetting
	}
	if !IsNil(o.IsDifferent) {
		toSerialize["IsDifferent"] = o.IsDifferent
	}
	if o.IsInUse != nil {
		toSerialize["IsInUse"] = o.IsInUse
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}
	if o.Category.IsSet() {
		toSerialize["Category"] = o.Category.Get()
	}
	if o.DefaultValue.IsSet() {
		toSerialize["DefaultValue"] = o.DefaultValue.Get()
	}
	return toSerialize, nil
}

type NullableComparisonRowContract struct {
	value *ComparisonRowContract
	isSet bool
}

func (v NullableComparisonRowContract) Get() *ComparisonRowContract {
	return v.value
}

func (v *NullableComparisonRowContract) Set(val *ComparisonRowContract) {
	v.value = val
	v.isSet = true
}

func (v NullableComparisonRowContract) IsSet() bool {
	return v.isSet
}

func (v *NullableComparisonRowContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparisonRowContract(val *ComparisonRowContract) *NullableComparisonRowContract {
	return &NullableComparisonRowContract{value: val, isSet: true}
}

func (v NullableComparisonRowContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComparisonRowContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


