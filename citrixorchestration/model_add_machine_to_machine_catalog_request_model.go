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

// checks if the AddMachineToMachineCatalogRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMachineToMachineCatalogRequestModel{}

// AddMachineToMachineCatalogRequestModel Add a machine to a machine catalog.
type AddMachineToMachineCatalogRequestModel struct {
	// Specify the name of the machine to create (in the form 'domain\\machine').  A SID can also be specified.
	MachineName NullableString `json:"MachineName,omitempty"`
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
	// Specifies whether the machine is initially in maintenance mode.  A machine in maintenance mode is not available for new sessions, and for managed machines all automatic power management is disabled. Optional; default is `false`.
	InMaintenanceMode NullableBool `json:"InMaintenanceMode,omitempty"`
}

// NewAddMachineToMachineCatalogRequestModel instantiates a new AddMachineToMachineCatalogRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMachineToMachineCatalogRequestModel() *AddMachineToMachineCatalogRequestModel {
	this := AddMachineToMachineCatalogRequestModel{}
	var inMaintenanceMode bool = false
	this.InMaintenanceMode = *NewNullableBool(&inMaintenanceMode)
	return &this
}

// NewAddMachineToMachineCatalogRequestModelWithDefaults instantiates a new AddMachineToMachineCatalogRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMachineToMachineCatalogRequestModelWithDefaults() *AddMachineToMachineCatalogRequestModel {
	this := AddMachineToMachineCatalogRequestModel{}
	var inMaintenanceMode bool = false
	this.InMaintenanceMode = *NewNullableBool(&inMaintenanceMode)
	return &this
}

// GetMachineName returns the MachineName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetMachineName() string {
	if o == nil || IsNil(o.MachineName.Get()) {
		var ret string
		return ret
	}
	return *o.MachineName.Get()
}

// GetMachineNameOk returns a tuple with the MachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetMachineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineName.Get(), o.MachineName.IsSet()
}

// HasMachineName returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasMachineName() bool {
	if o != nil && o.MachineName.IsSet() {
		return true
	}

	return false
}

// SetMachineName gets a reference to the given NullableString and assigns it to the MachineName field.
func (o *AddMachineToMachineCatalogRequestModel) SetMachineName(v string) {
	o.MachineName.Set(&v)
}
// SetMachineNameNil sets the value for MachineName to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetMachineNameNil() {
	o.MachineName.Set(nil)
}

// UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetMachineName() {
	o.MachineName.Unset()
}

// GetAssignedClientName returns the AssignedClientName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedClientName() string {
	if o == nil || IsNil(o.AssignedClientName.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedClientName.Get()
}

// GetAssignedClientNameOk returns a tuple with the AssignedClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedClientName.Get(), o.AssignedClientName.IsSet()
}

// HasAssignedClientName returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasAssignedClientName() bool {
	if o != nil && o.AssignedClientName.IsSet() {
		return true
	}

	return false
}

// SetAssignedClientName gets a reference to the given NullableString and assigns it to the AssignedClientName field.
func (o *AddMachineToMachineCatalogRequestModel) SetAssignedClientName(v string) {
	o.AssignedClientName.Set(&v)
}
// SetAssignedClientNameNil sets the value for AssignedClientName to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetAssignedClientNameNil() {
	o.AssignedClientName.Set(nil)
}

// UnsetAssignedClientName ensures that no value is present for AssignedClientName, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetAssignedClientName() {
	o.AssignedClientName.Unset()
}

// GetAssignedIPAddress returns the AssignedIPAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedIPAddress() string {
	if o == nil || IsNil(o.AssignedIPAddress.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedIPAddress.Get()
}

// GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedIPAddress.Get(), o.AssignedIPAddress.IsSet()
}

// HasAssignedIPAddress returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasAssignedIPAddress() bool {
	if o != nil && o.AssignedIPAddress.IsSet() {
		return true
	}

	return false
}

// SetAssignedIPAddress gets a reference to the given NullableString and assigns it to the AssignedIPAddress field.
func (o *AddMachineToMachineCatalogRequestModel) SetAssignedIPAddress(v string) {
	o.AssignedIPAddress.Set(&v)
}
// SetAssignedIPAddressNil sets the value for AssignedIPAddress to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetAssignedIPAddressNil() {
	o.AssignedIPAddress.Set(nil)
}

// UnsetAssignedIPAddress ensures that no value is present for AssignedIPAddress, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetAssignedIPAddress() {
	o.AssignedIPAddress.Unset()
}

// GetAssignedUsers returns the AssignedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssignedUsers
}

// GetAssignedUsersOk returns a tuple with the AssignedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetAssignedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUsers) {
		return nil, false
	}
	return o.AssignedUsers, true
}

// HasAssignedUsers returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasAssignedUsers() bool {
	if o != nil && IsNil(o.AssignedUsers) {
		return true
	}

	return false
}

// SetAssignedUsers gets a reference to the given []string and assigns it to the AssignedUsers field.
func (o *AddMachineToMachineCatalogRequestModel) SetAssignedUsers(v []string) {
	o.AssignedUsers = v
}

// GetHostedMachineId returns the HostedMachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetHostedMachineId() string {
	if o == nil || IsNil(o.HostedMachineId.Get()) {
		var ret string
		return ret
	}
	return *o.HostedMachineId.Get()
}

// GetHostedMachineIdOk returns a tuple with the HostedMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetHostedMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostedMachineId.Get(), o.HostedMachineId.IsSet()
}

// HasHostedMachineId returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasHostedMachineId() bool {
	if o != nil && o.HostedMachineId.IsSet() {
		return true
	}

	return false
}

// SetHostedMachineId gets a reference to the given NullableString and assigns it to the HostedMachineId field.
func (o *AddMachineToMachineCatalogRequestModel) SetHostedMachineId(v string) {
	o.HostedMachineId.Set(&v)
}
// SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetHostedMachineIdNil() {
	o.HostedMachineId.Set(nil)
}

// UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetHostedMachineId() {
	o.HostedMachineId.Unset()
}

// GetHypervisorConnection returns the HypervisorConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetHypervisorConnection() string {
	if o == nil || IsNil(o.HypervisorConnection.Get()) {
		var ret string
		return ret
	}
	return *o.HypervisorConnection.Get()
}

// GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetHypervisorConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HypervisorConnection.Get(), o.HypervisorConnection.IsSet()
}

// HasHypervisorConnection returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasHypervisorConnection() bool {
	if o != nil && o.HypervisorConnection.IsSet() {
		return true
	}

	return false
}

// SetHypervisorConnection gets a reference to the given NullableString and assigns it to the HypervisorConnection field.
func (o *AddMachineToMachineCatalogRequestModel) SetHypervisorConnection(v string) {
	o.HypervisorConnection.Set(&v)
}
// SetHypervisorConnectionNil sets the value for HypervisorConnection to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetHypervisorConnectionNil() {
	o.HypervisorConnection.Set(nil)
}

// UnsetHypervisorConnection ensures that no value is present for HypervisorConnection, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetHypervisorConnection() {
	o.HypervisorConnection.Unset()
}

// GetInMaintenanceMode returns the InMaintenanceMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMachineToMachineCatalogRequestModel) GetInMaintenanceMode() bool {
	if o == nil || IsNil(o.InMaintenanceMode.Get()) {
		var ret bool
		return ret
	}
	return *o.InMaintenanceMode.Get()
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMachineToMachineCatalogRequestModel) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InMaintenanceMode.Get(), o.InMaintenanceMode.IsSet()
}

// HasInMaintenanceMode returns a boolean if a field has been set.
func (o *AddMachineToMachineCatalogRequestModel) HasInMaintenanceMode() bool {
	if o != nil && o.InMaintenanceMode.IsSet() {
		return true
	}

	return false
}

// SetInMaintenanceMode gets a reference to the given NullableBool and assigns it to the InMaintenanceMode field.
func (o *AddMachineToMachineCatalogRequestModel) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode.Set(&v)
}
// SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) SetInMaintenanceModeNil() {
	o.InMaintenanceMode.Set(nil)
}

// UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
func (o *AddMachineToMachineCatalogRequestModel) UnsetInMaintenanceMode() {
	o.InMaintenanceMode.Unset()
}

func (o AddMachineToMachineCatalogRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddMachineToMachineCatalogRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MachineName.IsSet() {
		toSerialize["MachineName"] = o.MachineName.Get()
	}
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
	return toSerialize, nil
}

type NullableAddMachineToMachineCatalogRequestModel struct {
	value *AddMachineToMachineCatalogRequestModel
	isSet bool
}

func (v NullableAddMachineToMachineCatalogRequestModel) Get() *AddMachineToMachineCatalogRequestModel {
	return v.value
}

func (v *NullableAddMachineToMachineCatalogRequestModel) Set(val *AddMachineToMachineCatalogRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMachineToMachineCatalogRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMachineToMachineCatalogRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMachineToMachineCatalogRequestModel(val *AddMachineToMachineCatalogRequestModel) *NullableAddMachineToMachineCatalogRequestModel {
	return &NullableAddMachineToMachineCatalogRequestModel{value: val, isSet: true}
}

func (v NullableAddMachineToMachineCatalogRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMachineToMachineCatalogRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


