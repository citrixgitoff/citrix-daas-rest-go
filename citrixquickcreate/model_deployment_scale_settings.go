/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the DeploymentScaleSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentScaleSettings{}

// DeploymentScaleSettings Power management scale settings
type DeploymentScaleSettings struct {
	// Indicates whether the scale settings are enabled for the deployment
	AutoScaleEnabled NullableBool `json:"autoScaleEnabled,omitempty"`
	// The number of minutes before the configured action should be  performed after a user session disconnects outside peak hours.
	OffPeakDisconnectTimeoutMinutes NullableInt32 `json:"offPeakDisconnectTimeoutMinutes,omitempty"`
	// The number of minutes before the configured action should be  performed after a user session ends outside peak hours.
	OffPeakLogOffTimeoutMinutes NullableInt32 `json:"offPeakLogOffTimeoutMinutes,omitempty"`
	// Specifies the time in minutes after which an idle session  belonging to the delivery group is disconnected during off-peak time.
	SessionIdleTimeoutMinutes NullableInt32 `json:"sessionIdleTimeoutMinutes,omitempty"`
	// The percentage of machines in the delivery group that should be  kept available in an idle state outside peak hours.
	OffPeakBufferSizePercent NullableInt32 `json:"offPeakBufferSizePercent,omitempty"`
}

// NewDeploymentScaleSettings instantiates a new DeploymentScaleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentScaleSettings() *DeploymentScaleSettings {
	this := DeploymentScaleSettings{}
	return &this
}

// NewDeploymentScaleSettingsWithDefaults instantiates a new DeploymentScaleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentScaleSettingsWithDefaults() *DeploymentScaleSettings {
	this := DeploymentScaleSettings{}
	return &this
}

// GetAutoScaleEnabled returns the AutoScaleEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentScaleSettings) GetAutoScaleEnabled() bool {
	if o == nil || IsNil(o.AutoScaleEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoScaleEnabled.Get()
}

// GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentScaleSettings) GetAutoScaleEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScaleEnabled.Get(), o.AutoScaleEnabled.IsSet()
}

// HasAutoScaleEnabled returns a boolean if a field has been set.
func (o *DeploymentScaleSettings) HasAutoScaleEnabled() bool {
	if o != nil && o.AutoScaleEnabled.IsSet() {
		return true
	}

	return false
}

// SetAutoScaleEnabled gets a reference to the given NullableBool and assigns it to the AutoScaleEnabled field.
func (o *DeploymentScaleSettings) SetAutoScaleEnabled(v bool) {
	o.AutoScaleEnabled.Set(&v)
}
// SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil
func (o *DeploymentScaleSettings) SetAutoScaleEnabledNil() {
	o.AutoScaleEnabled.Set(nil)
}

// UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
func (o *DeploymentScaleSettings) UnsetAutoScaleEnabled() {
	o.AutoScaleEnabled.Unset()
}

// GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentScaleSettings) GetOffPeakDisconnectTimeoutMinutes() int32 {
	if o == nil || IsNil(o.OffPeakDisconnectTimeoutMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.OffPeakDisconnectTimeoutMinutes.Get()
}

// GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentScaleSettings) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OffPeakDisconnectTimeoutMinutes.Get(), o.OffPeakDisconnectTimeoutMinutes.IsSet()
}

// HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.
func (o *DeploymentScaleSettings) HasOffPeakDisconnectTimeoutMinutes() bool {
	if o != nil && o.OffPeakDisconnectTimeoutMinutes.IsSet() {
		return true
	}

	return false
}

// SetOffPeakDisconnectTimeoutMinutes gets a reference to the given NullableInt32 and assigns it to the OffPeakDisconnectTimeoutMinutes field.
func (o *DeploymentScaleSettings) SetOffPeakDisconnectTimeoutMinutes(v int32) {
	o.OffPeakDisconnectTimeoutMinutes.Set(&v)
}
// SetOffPeakDisconnectTimeoutMinutesNil sets the value for OffPeakDisconnectTimeoutMinutes to be an explicit nil
func (o *DeploymentScaleSettings) SetOffPeakDisconnectTimeoutMinutesNil() {
	o.OffPeakDisconnectTimeoutMinutes.Set(nil)
}

// UnsetOffPeakDisconnectTimeoutMinutes ensures that no value is present for OffPeakDisconnectTimeoutMinutes, not even an explicit nil
func (o *DeploymentScaleSettings) UnsetOffPeakDisconnectTimeoutMinutes() {
	o.OffPeakDisconnectTimeoutMinutes.Unset()
}

// GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentScaleSettings) GetOffPeakLogOffTimeoutMinutes() int32 {
	if o == nil || IsNil(o.OffPeakLogOffTimeoutMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.OffPeakLogOffTimeoutMinutes.Get()
}

// GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentScaleSettings) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OffPeakLogOffTimeoutMinutes.Get(), o.OffPeakLogOffTimeoutMinutes.IsSet()
}

// HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.
func (o *DeploymentScaleSettings) HasOffPeakLogOffTimeoutMinutes() bool {
	if o != nil && o.OffPeakLogOffTimeoutMinutes.IsSet() {
		return true
	}

	return false
}

// SetOffPeakLogOffTimeoutMinutes gets a reference to the given NullableInt32 and assigns it to the OffPeakLogOffTimeoutMinutes field.
func (o *DeploymentScaleSettings) SetOffPeakLogOffTimeoutMinutes(v int32) {
	o.OffPeakLogOffTimeoutMinutes.Set(&v)
}
// SetOffPeakLogOffTimeoutMinutesNil sets the value for OffPeakLogOffTimeoutMinutes to be an explicit nil
func (o *DeploymentScaleSettings) SetOffPeakLogOffTimeoutMinutesNil() {
	o.OffPeakLogOffTimeoutMinutes.Set(nil)
}

// UnsetOffPeakLogOffTimeoutMinutes ensures that no value is present for OffPeakLogOffTimeoutMinutes, not even an explicit nil
func (o *DeploymentScaleSettings) UnsetOffPeakLogOffTimeoutMinutes() {
	o.OffPeakLogOffTimeoutMinutes.Unset()
}

// GetSessionIdleTimeoutMinutes returns the SessionIdleTimeoutMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentScaleSettings) GetSessionIdleTimeoutMinutes() int32 {
	if o == nil || IsNil(o.SessionIdleTimeoutMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.SessionIdleTimeoutMinutes.Get()
}

// GetSessionIdleTimeoutMinutesOk returns a tuple with the SessionIdleTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentScaleSettings) GetSessionIdleTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionIdleTimeoutMinutes.Get(), o.SessionIdleTimeoutMinutes.IsSet()
}

// HasSessionIdleTimeoutMinutes returns a boolean if a field has been set.
func (o *DeploymentScaleSettings) HasSessionIdleTimeoutMinutes() bool {
	if o != nil && o.SessionIdleTimeoutMinutes.IsSet() {
		return true
	}

	return false
}

// SetSessionIdleTimeoutMinutes gets a reference to the given NullableInt32 and assigns it to the SessionIdleTimeoutMinutes field.
func (o *DeploymentScaleSettings) SetSessionIdleTimeoutMinutes(v int32) {
	o.SessionIdleTimeoutMinutes.Set(&v)
}
// SetSessionIdleTimeoutMinutesNil sets the value for SessionIdleTimeoutMinutes to be an explicit nil
func (o *DeploymentScaleSettings) SetSessionIdleTimeoutMinutesNil() {
	o.SessionIdleTimeoutMinutes.Set(nil)
}

// UnsetSessionIdleTimeoutMinutes ensures that no value is present for SessionIdleTimeoutMinutes, not even an explicit nil
func (o *DeploymentScaleSettings) UnsetSessionIdleTimeoutMinutes() {
	o.SessionIdleTimeoutMinutes.Unset()
}

// GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentScaleSettings) GetOffPeakBufferSizePercent() int32 {
	if o == nil || IsNil(o.OffPeakBufferSizePercent.Get()) {
		var ret int32
		return ret
	}
	return *o.OffPeakBufferSizePercent.Get()
}

// GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentScaleSettings) GetOffPeakBufferSizePercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OffPeakBufferSizePercent.Get(), o.OffPeakBufferSizePercent.IsSet()
}

// HasOffPeakBufferSizePercent returns a boolean if a field has been set.
func (o *DeploymentScaleSettings) HasOffPeakBufferSizePercent() bool {
	if o != nil && o.OffPeakBufferSizePercent.IsSet() {
		return true
	}

	return false
}

// SetOffPeakBufferSizePercent gets a reference to the given NullableInt32 and assigns it to the OffPeakBufferSizePercent field.
func (o *DeploymentScaleSettings) SetOffPeakBufferSizePercent(v int32) {
	o.OffPeakBufferSizePercent.Set(&v)
}
// SetOffPeakBufferSizePercentNil sets the value for OffPeakBufferSizePercent to be an explicit nil
func (o *DeploymentScaleSettings) SetOffPeakBufferSizePercentNil() {
	o.OffPeakBufferSizePercent.Set(nil)
}

// UnsetOffPeakBufferSizePercent ensures that no value is present for OffPeakBufferSizePercent, not even an explicit nil
func (o *DeploymentScaleSettings) UnsetOffPeakBufferSizePercent() {
	o.OffPeakBufferSizePercent.Unset()
}

func (o DeploymentScaleSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentScaleSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoScaleEnabled.IsSet() {
		toSerialize["autoScaleEnabled"] = o.AutoScaleEnabled.Get()
	}
	if o.OffPeakDisconnectTimeoutMinutes.IsSet() {
		toSerialize["offPeakDisconnectTimeoutMinutes"] = o.OffPeakDisconnectTimeoutMinutes.Get()
	}
	if o.OffPeakLogOffTimeoutMinutes.IsSet() {
		toSerialize["offPeakLogOffTimeoutMinutes"] = o.OffPeakLogOffTimeoutMinutes.Get()
	}
	if o.SessionIdleTimeoutMinutes.IsSet() {
		toSerialize["sessionIdleTimeoutMinutes"] = o.SessionIdleTimeoutMinutes.Get()
	}
	if o.OffPeakBufferSizePercent.IsSet() {
		toSerialize["offPeakBufferSizePercent"] = o.OffPeakBufferSizePercent.Get()
	}
	return toSerialize, nil
}

type NullableDeploymentScaleSettings struct {
	value *DeploymentScaleSettings
	isSet bool
}

func (v NullableDeploymentScaleSettings) Get() *DeploymentScaleSettings {
	return v.value
}

func (v *NullableDeploymentScaleSettings) Set(val *DeploymentScaleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentScaleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentScaleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentScaleSettings(val *DeploymentScaleSettings) *NullableDeploymentScaleSettings {
	return &NullableDeploymentScaleSettings{value: val, isSet: true}
}

func (v NullableDeploymentScaleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentScaleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


