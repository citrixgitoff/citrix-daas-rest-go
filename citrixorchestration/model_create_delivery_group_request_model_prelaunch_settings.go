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

// checks if the CreateDeliveryGroupRequestModelPrelaunchSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeliveryGroupRequestModelPrelaunchSettings{}

// CreateDeliveryGroupRequestModelPrelaunchSettings Prelaunch settings for users with resources published from the delivery group.
type CreateDeliveryGroupRequestModelPrelaunchSettings struct {
	// Indicates if the pre-launch or lingering settings are enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// Specifies the average load threshold across the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest will be terminated to reduce load.
	MaxAverageLoadThreshold *int32 `json:"MaxAverageLoadThreshold,omitempty"`
	// Specifies the maximum load threshold per machine in the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest on each loaded machine will be terminated to reduce load. Was: public int? MaxPerDesktopLoadThreshold { get; set; }
	MaxLoadPerMachineThreshold *int32 `json:"MaxLoadPerMachineThreshold,omitempty"`
	// Specifies the time by which a pre-launched or lingering session will be disconnected.
	MaxTimeBeforeDisconnectMinutes *int32 `json:"MaxTimeBeforeDisconnectMinutes,omitempty"`
	// Specifies the time by which a pre-launched or lingering session will be terminated.
	MaxTimeBeforeTerminateMinutes *int32 `json:"MaxTimeBeforeTerminateMinutes,omitempty"`
	// Specifies whether the IncludedUsers filter is enabled or disabled.  When the user filter is enabled, pre-launch or lingering are enabled only to users who appear in the filter (either explicitly or by virtue of group membership). Was: UserFilterEnabled
	IncludedUserFilterEnabled *bool `json:"IncludedUserFilterEnabled,omitempty"`
	// Specifies the users and user groups to whom the pre-launch or lingering settings apply. Was: users
	IncludedUsers []string `json:"IncludedUsers,omitempty"`
}

// NewCreateDeliveryGroupRequestModelPrelaunchSettings instantiates a new CreateDeliveryGroupRequestModelPrelaunchSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeliveryGroupRequestModelPrelaunchSettings() *CreateDeliveryGroupRequestModelPrelaunchSettings {
	this := CreateDeliveryGroupRequestModelPrelaunchSettings{}
	var enabled bool = true
	this.Enabled = &enabled
	var maxAverageLoadThreshold int32 = 0
	this.MaxAverageLoadThreshold = &maxAverageLoadThreshold
	var maxLoadPerMachineThreshold int32 = 0
	this.MaxLoadPerMachineThreshold = &maxLoadPerMachineThreshold
	var maxTimeBeforeDisconnectMinutes int32 = 0
	this.MaxTimeBeforeDisconnectMinutes = &maxTimeBeforeDisconnectMinutes
	var maxTimeBeforeTerminateMinutes int32 = 0
	this.MaxTimeBeforeTerminateMinutes = &maxTimeBeforeTerminateMinutes
	var includedUserFilterEnabled bool = false
	this.IncludedUserFilterEnabled = &includedUserFilterEnabled
	return &this
}

// NewCreateDeliveryGroupRequestModelPrelaunchSettingsWithDefaults instantiates a new CreateDeliveryGroupRequestModelPrelaunchSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeliveryGroupRequestModelPrelaunchSettingsWithDefaults() *CreateDeliveryGroupRequestModelPrelaunchSettings {
	this := CreateDeliveryGroupRequestModelPrelaunchSettings{}
	var enabled bool = true
	this.Enabled = &enabled
	var maxAverageLoadThreshold int32 = 0
	this.MaxAverageLoadThreshold = &maxAverageLoadThreshold
	var maxLoadPerMachineThreshold int32 = 0
	this.MaxLoadPerMachineThreshold = &maxLoadPerMachineThreshold
	var maxTimeBeforeDisconnectMinutes int32 = 0
	this.MaxTimeBeforeDisconnectMinutes = &maxTimeBeforeDisconnectMinutes
	var maxTimeBeforeTerminateMinutes int32 = 0
	this.MaxTimeBeforeTerminateMinutes = &maxTimeBeforeTerminateMinutes
	var includedUserFilterEnabled bool = false
	this.IncludedUserFilterEnabled = &includedUserFilterEnabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxAverageLoadThreshold returns the MaxAverageLoadThreshold field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxAverageLoadThreshold() int32 {
	if o == nil || IsNil(o.MaxAverageLoadThreshold) {
		var ret int32
		return ret
	}
	return *o.MaxAverageLoadThreshold
}

// GetMaxAverageLoadThresholdOk returns a tuple with the MaxAverageLoadThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxAverageLoadThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAverageLoadThreshold) {
		return nil, false
	}
	return o.MaxAverageLoadThreshold, true
}

// HasMaxAverageLoadThreshold returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxAverageLoadThreshold() bool {
	if o != nil && !IsNil(o.MaxAverageLoadThreshold) {
		return true
	}

	return false
}

// SetMaxAverageLoadThreshold gets a reference to the given int32 and assigns it to the MaxAverageLoadThreshold field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxAverageLoadThreshold(v int32) {
	o.MaxAverageLoadThreshold = &v
}

// GetMaxLoadPerMachineThreshold returns the MaxLoadPerMachineThreshold field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxLoadPerMachineThreshold() int32 {
	if o == nil || IsNil(o.MaxLoadPerMachineThreshold) {
		var ret int32
		return ret
	}
	return *o.MaxLoadPerMachineThreshold
}

// GetMaxLoadPerMachineThresholdOk returns a tuple with the MaxLoadPerMachineThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxLoadPerMachineThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLoadPerMachineThreshold) {
		return nil, false
	}
	return o.MaxLoadPerMachineThreshold, true
}

// HasMaxLoadPerMachineThreshold returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxLoadPerMachineThreshold() bool {
	if o != nil && !IsNil(o.MaxLoadPerMachineThreshold) {
		return true
	}

	return false
}

// SetMaxLoadPerMachineThreshold gets a reference to the given int32 and assigns it to the MaxLoadPerMachineThreshold field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxLoadPerMachineThreshold(v int32) {
	o.MaxLoadPerMachineThreshold = &v
}

// GetMaxTimeBeforeDisconnectMinutes returns the MaxTimeBeforeDisconnectMinutes field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeDisconnectMinutes() int32 {
	if o == nil || IsNil(o.MaxTimeBeforeDisconnectMinutes) {
		var ret int32
		return ret
	}
	return *o.MaxTimeBeforeDisconnectMinutes
}

// GetMaxTimeBeforeDisconnectMinutesOk returns a tuple with the MaxTimeBeforeDisconnectMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeDisconnectMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTimeBeforeDisconnectMinutes) {
		return nil, false
	}
	return o.MaxTimeBeforeDisconnectMinutes, true
}

// HasMaxTimeBeforeDisconnectMinutes returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxTimeBeforeDisconnectMinutes() bool {
	if o != nil && !IsNil(o.MaxTimeBeforeDisconnectMinutes) {
		return true
	}

	return false
}

// SetMaxTimeBeforeDisconnectMinutes gets a reference to the given int32 and assigns it to the MaxTimeBeforeDisconnectMinutes field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxTimeBeforeDisconnectMinutes(v int32) {
	o.MaxTimeBeforeDisconnectMinutes = &v
}

// GetMaxTimeBeforeTerminateMinutes returns the MaxTimeBeforeTerminateMinutes field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeTerminateMinutes() int32 {
	if o == nil || IsNil(o.MaxTimeBeforeTerminateMinutes) {
		var ret int32
		return ret
	}
	return *o.MaxTimeBeforeTerminateMinutes
}

// GetMaxTimeBeforeTerminateMinutesOk returns a tuple with the MaxTimeBeforeTerminateMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeTerminateMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTimeBeforeTerminateMinutes) {
		return nil, false
	}
	return o.MaxTimeBeforeTerminateMinutes, true
}

// HasMaxTimeBeforeTerminateMinutes returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxTimeBeforeTerminateMinutes() bool {
	if o != nil && !IsNil(o.MaxTimeBeforeTerminateMinutes) {
		return true
	}

	return false
}

// SetMaxTimeBeforeTerminateMinutes gets a reference to the given int32 and assigns it to the MaxTimeBeforeTerminateMinutes field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxTimeBeforeTerminateMinutes(v int32) {
	o.MaxTimeBeforeTerminateMinutes = &v
}

// GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUserFilterEnabled() bool {
	if o == nil || IsNil(o.IncludedUserFilterEnabled) {
		var ret bool
		return ret
	}
	return *o.IncludedUserFilterEnabled
}

// GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUserFilterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludedUserFilterEnabled) {
		return nil, false
	}
	return o.IncludedUserFilterEnabled, true
}

// HasIncludedUserFilterEnabled returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasIncludedUserFilterEnabled() bool {
	if o != nil && !IsNil(o.IncludedUserFilterEnabled) {
		return true
	}

	return false
}

// SetIncludedUserFilterEnabled gets a reference to the given bool and assigns it to the IncludedUserFilterEnabled field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetIncludedUserFilterEnabled(v bool) {
	o.IncludedUserFilterEnabled = &v
}

// GetIncludedUsers returns the IncludedUsers field value if set, zero value otherwise.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUsers() []string {
	if o == nil || IsNil(o.IncludedUsers) {
		var ret []string
		return ret
	}
	return o.IncludedUsers
}

// GetIncludedUsersOk returns a tuple with the IncludedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUsers) {
		return nil, false
	}
	return o.IncludedUsers, true
}

// HasIncludedUsers returns a boolean if a field has been set.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasIncludedUsers() bool {
	if o != nil && !IsNil(o.IncludedUsers) {
		return true
	}

	return false
}

// SetIncludedUsers gets a reference to the given []string and assigns it to the IncludedUsers field.
func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetIncludedUsers(v []string) {
	o.IncludedUsers = v
}

func (o CreateDeliveryGroupRequestModelPrelaunchSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeliveryGroupRequestModelPrelaunchSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.MaxAverageLoadThreshold) {
		toSerialize["MaxAverageLoadThreshold"] = o.MaxAverageLoadThreshold
	}
	if !IsNil(o.MaxLoadPerMachineThreshold) {
		toSerialize["MaxLoadPerMachineThreshold"] = o.MaxLoadPerMachineThreshold
	}
	if !IsNil(o.MaxTimeBeforeDisconnectMinutes) {
		toSerialize["MaxTimeBeforeDisconnectMinutes"] = o.MaxTimeBeforeDisconnectMinutes
	}
	if !IsNil(o.MaxTimeBeforeTerminateMinutes) {
		toSerialize["MaxTimeBeforeTerminateMinutes"] = o.MaxTimeBeforeTerminateMinutes
	}
	if !IsNil(o.IncludedUserFilterEnabled) {
		toSerialize["IncludedUserFilterEnabled"] = o.IncludedUserFilterEnabled
	}
	if !IsNil(o.IncludedUsers) {
		toSerialize["IncludedUsers"] = o.IncludedUsers
	}
	return toSerialize, nil
}

type NullableCreateDeliveryGroupRequestModelPrelaunchSettings struct {
	value *CreateDeliveryGroupRequestModelPrelaunchSettings
	isSet bool
}

func (v NullableCreateDeliveryGroupRequestModelPrelaunchSettings) Get() *CreateDeliveryGroupRequestModelPrelaunchSettings {
	return v.value
}

func (v *NullableCreateDeliveryGroupRequestModelPrelaunchSettings) Set(val *CreateDeliveryGroupRequestModelPrelaunchSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeliveryGroupRequestModelPrelaunchSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeliveryGroupRequestModelPrelaunchSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeliveryGroupRequestModelPrelaunchSettings(val *CreateDeliveryGroupRequestModelPrelaunchSettings) *NullableCreateDeliveryGroupRequestModelPrelaunchSettings {
	return &NullableCreateDeliveryGroupRequestModelPrelaunchSettings{value: val, isSet: true}
}

func (v NullableCreateDeliveryGroupRequestModelPrelaunchSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeliveryGroupRequestModelPrelaunchSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


