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

// checks if the UpdateMachineRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMachineRequestModel{}

// UpdateMachineRequestModel Update a machine.
type UpdateMachineRequestModel struct {
	// The client name to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time.
	AssignedClientName NullableString `json:"AssignedClientName,omitempty"`
	// The client IP address to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time.
	AssignedIPAddress NullableString `json:"AssignedIPAddress,omitempty"`
	// The user(s) to whom this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time.
	AssignedUsers []string `json:"AssignedUsers,omitempty"`
	// The unique ID by which the hypervisor recognizes the machine. Omit this for machines that are not power-managed.
	HostedMachineId NullableString `json:"HostedMachineId,omitempty"`
	// Hypervisor connection to use for power management of the machine.
	HypervisorConnection NullableString `json:"HypervisorConnection,omitempty"`
	// Specifies whether the machine is initially in maintenance mode.  A machine in maintenance mode is not available for new sessions, and for managed machines all automatic power management is disabled. If `null`, will not be changed.
	InMaintenanceMode NullableBool `json:"InMaintenanceMode,omitempty"`
	// Customized name of the machine that is displayed in StoreFront, if the machine has been published. It can be set only for private desktops. If `null`, will not be changed. If empty string (`\"\"`), the machine will be unassigned from any published name.
	PublishedName NullableString `json:"PublishedName,omitempty"`
}

// NewUpdateMachineRequestModel instantiates a new UpdateMachineRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMachineRequestModel() *UpdateMachineRequestModel {
	this := UpdateMachineRequestModel{}
	return &this
}

// NewUpdateMachineRequestModelWithDefaults instantiates a new UpdateMachineRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMachineRequestModelWithDefaults() *UpdateMachineRequestModel {
	this := UpdateMachineRequestModel{}
	return &this
}

// GetAssignedClientName returns the AssignedClientName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetAssignedClientName() string {
	if o == nil || IsNil(o.AssignedClientName.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedClientName.Get()
}

// GetAssignedClientNameOk returns a tuple with the AssignedClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetAssignedClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedClientName.Get(), o.AssignedClientName.IsSet()
}

// HasAssignedClientName returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasAssignedClientName() bool {
	if o != nil && o.AssignedClientName.IsSet() {
		return true
	}

	return false
}

// SetAssignedClientName gets a reference to the given NullableString and assigns it to the AssignedClientName field.
func (o *UpdateMachineRequestModel) SetAssignedClientName(v string) {
	o.AssignedClientName.Set(&v)
}
// SetAssignedClientNameNil sets the value for AssignedClientName to be an explicit nil
func (o *UpdateMachineRequestModel) SetAssignedClientNameNil() {
	o.AssignedClientName.Set(nil)
}

// UnsetAssignedClientName ensures that no value is present for AssignedClientName, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetAssignedClientName() {
	o.AssignedClientName.Unset()
}

// GetAssignedIPAddress returns the AssignedIPAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetAssignedIPAddress() string {
	if o == nil || IsNil(o.AssignedIPAddress.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedIPAddress.Get()
}

// GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetAssignedIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedIPAddress.Get(), o.AssignedIPAddress.IsSet()
}

// HasAssignedIPAddress returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasAssignedIPAddress() bool {
	if o != nil && o.AssignedIPAddress.IsSet() {
		return true
	}

	return false
}

// SetAssignedIPAddress gets a reference to the given NullableString and assigns it to the AssignedIPAddress field.
func (o *UpdateMachineRequestModel) SetAssignedIPAddress(v string) {
	o.AssignedIPAddress.Set(&v)
}
// SetAssignedIPAddressNil sets the value for AssignedIPAddress to be an explicit nil
func (o *UpdateMachineRequestModel) SetAssignedIPAddressNil() {
	o.AssignedIPAddress.Set(nil)
}

// UnsetAssignedIPAddress ensures that no value is present for AssignedIPAddress, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetAssignedIPAddress() {
	o.AssignedIPAddress.Unset()
}

// GetAssignedUsers returns the AssignedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetAssignedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssignedUsers
}

// GetAssignedUsersOk returns a tuple with the AssignedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetAssignedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUsers) {
		return nil, false
	}
	return o.AssignedUsers, true
}

// HasAssignedUsers returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasAssignedUsers() bool {
	if o != nil && IsNil(o.AssignedUsers) {
		return true
	}

	return false
}

// SetAssignedUsers gets a reference to the given []string and assigns it to the AssignedUsers field.
func (o *UpdateMachineRequestModel) SetAssignedUsers(v []string) {
	o.AssignedUsers = v
}

// GetHostedMachineId returns the HostedMachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetHostedMachineId() string {
	if o == nil || IsNil(o.HostedMachineId.Get()) {
		var ret string
		return ret
	}
	return *o.HostedMachineId.Get()
}

// GetHostedMachineIdOk returns a tuple with the HostedMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetHostedMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostedMachineId.Get(), o.HostedMachineId.IsSet()
}

// HasHostedMachineId returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasHostedMachineId() bool {
	if o != nil && o.HostedMachineId.IsSet() {
		return true
	}

	return false
}

// SetHostedMachineId gets a reference to the given NullableString and assigns it to the HostedMachineId field.
func (o *UpdateMachineRequestModel) SetHostedMachineId(v string) {
	o.HostedMachineId.Set(&v)
}
// SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil
func (o *UpdateMachineRequestModel) SetHostedMachineIdNil() {
	o.HostedMachineId.Set(nil)
}

// UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetHostedMachineId() {
	o.HostedMachineId.Unset()
}

// GetHypervisorConnection returns the HypervisorConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetHypervisorConnection() string {
	if o == nil || IsNil(o.HypervisorConnection.Get()) {
		var ret string
		return ret
	}
	return *o.HypervisorConnection.Get()
}

// GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetHypervisorConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HypervisorConnection.Get(), o.HypervisorConnection.IsSet()
}

// HasHypervisorConnection returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasHypervisorConnection() bool {
	if o != nil && o.HypervisorConnection.IsSet() {
		return true
	}

	return false
}

// SetHypervisorConnection gets a reference to the given NullableString and assigns it to the HypervisorConnection field.
func (o *UpdateMachineRequestModel) SetHypervisorConnection(v string) {
	o.HypervisorConnection.Set(&v)
}
// SetHypervisorConnectionNil sets the value for HypervisorConnection to be an explicit nil
func (o *UpdateMachineRequestModel) SetHypervisorConnectionNil() {
	o.HypervisorConnection.Set(nil)
}

// UnsetHypervisorConnection ensures that no value is present for HypervisorConnection, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetHypervisorConnection() {
	o.HypervisorConnection.Unset()
}

// GetInMaintenanceMode returns the InMaintenanceMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetInMaintenanceMode() bool {
	if o == nil || IsNil(o.InMaintenanceMode.Get()) {
		var ret bool
		return ret
	}
	return *o.InMaintenanceMode.Get()
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InMaintenanceMode.Get(), o.InMaintenanceMode.IsSet()
}

// HasInMaintenanceMode returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasInMaintenanceMode() bool {
	if o != nil && o.InMaintenanceMode.IsSet() {
		return true
	}

	return false
}

// SetInMaintenanceMode gets a reference to the given NullableBool and assigns it to the InMaintenanceMode field.
func (o *UpdateMachineRequestModel) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode.Set(&v)
}
// SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil
func (o *UpdateMachineRequestModel) SetInMaintenanceModeNil() {
	o.InMaintenanceMode.Set(nil)
}

// UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetInMaintenanceMode() {
	o.InMaintenanceMode.Unset()
}

// GetPublishedName returns the PublishedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMachineRequestModel) GetPublishedName() string {
	if o == nil || IsNil(o.PublishedName.Get()) {
		var ret string
		return ret
	}
	return *o.PublishedName.Get()
}

// GetPublishedNameOk returns a tuple with the PublishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMachineRequestModel) GetPublishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublishedName.Get(), o.PublishedName.IsSet()
}

// HasPublishedName returns a boolean if a field has been set.
func (o *UpdateMachineRequestModel) HasPublishedName() bool {
	if o != nil && o.PublishedName.IsSet() {
		return true
	}

	return false
}

// SetPublishedName gets a reference to the given NullableString and assigns it to the PublishedName field.
func (o *UpdateMachineRequestModel) SetPublishedName(v string) {
	o.PublishedName.Set(&v)
}
// SetPublishedNameNil sets the value for PublishedName to be an explicit nil
func (o *UpdateMachineRequestModel) SetPublishedNameNil() {
	o.PublishedName.Set(nil)
}

// UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
func (o *UpdateMachineRequestModel) UnsetPublishedName() {
	o.PublishedName.Unset()
}

func (o UpdateMachineRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMachineRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedClientName.IsSet() {
		toSerialize["AssignedClientName"] = o.AssignedClientName.Get()
	}
	if o.AssignedIPAddress.IsSet() {
		toSerialize["AssignedIPAddress"] = o.AssignedIPAddress.Get()
	}
	if o.AssignedUsers != nil {
		toSerialize["AssignedUsers"] = o.AssignedUsers
	}
	if o.HostedMachineId.IsSet() {
		toSerialize["HostedMachineId"] = o.HostedMachineId.Get()
	}
	if o.HypervisorConnection.IsSet() {
		toSerialize["HypervisorConnection"] = o.HypervisorConnection.Get()
	}
	if o.InMaintenanceMode.IsSet() {
		toSerialize["InMaintenanceMode"] = o.InMaintenanceMode.Get()
	}
	if o.PublishedName.IsSet() {
		toSerialize["PublishedName"] = o.PublishedName.Get()
	}
	return toSerialize, nil
}

type NullableUpdateMachineRequestModel struct {
	value *UpdateMachineRequestModel
	isSet bool
}

func (v NullableUpdateMachineRequestModel) Get() *UpdateMachineRequestModel {
	return v.value
}

func (v *NullableUpdateMachineRequestModel) Set(val *UpdateMachineRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMachineRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMachineRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMachineRequestModel(val *UpdateMachineRequestModel) *NullableUpdateMachineRequestModel {
	return &NullableUpdateMachineRequestModel{value: val, isSet: true}
}

func (v NullableUpdateMachineRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMachineRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


