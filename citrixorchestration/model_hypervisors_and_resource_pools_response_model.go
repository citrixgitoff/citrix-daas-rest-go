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

// checks if the HypervisorsAndResourcePoolsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorsAndResourcePoolsResponseModel{}

// HypervisorsAndResourcePoolsResponseModel struct for HypervisorsAndResourcePoolsResponseModel
type HypervisorsAndResourcePoolsResponseModel struct {
	// Id of the resource.
	Id NullableString `json:"Id,omitempty"`
	// Name of the resource.
	Name NullableString `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath NullableString `json:"XDPath,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// Addresses that can be used to contact the required hypervisor. All the addresses are considered equivalent, that is, all of the addresses provide access to the same virtual machines, snapshots, network, and storage.
	Addresses []string `json:"Addresses"`
	// Indicates whether the hypervisor is in maintenance mode, which disables all communication between XenApp & XenDesktop and the Hypervisor.
	InMaintenanceMode bool `json:"InMaintenanceMode"`
	// Indicates whether is UnEntitled.
	Unentitled *bool `json:"Unentitled,omitempty"`
	// The class name for the Citrix Managed Machine library that is used to access the hypervisor.
	PluginId string `json:"PluginId"`
	// The list of administrative scopes that the connection is a part of. The scopes control which administrators are able to work with the connection.
	Scopes []ScopeResponseModel `json:"Scopes"`
	// The tenant(s) that the hypervisor is assigned to.  If `null`, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants.
	Tenants []RefResponseModel `json:"Tenants,omitempty"`
	// Indicates whether the hypervisor uses cloud infrastructure.
	UsesCloudInfrastructure bool `json:"UsesCloudInfrastructure"`
	Zone RefResponseModel `json:"Zone"`
	Fault *HypervisorFaultResponseModel `json:"Fault,omitempty"`
	// CustomProperties of hypervisor connection
	CustomProperties NullableString `json:"CustomProperties,omitempty"`
	// The broker id.
	Uid NullableInt32 `json:"Uid,omitempty"`
	// If this connection is virtual.
	IsVirtual *bool `json:"IsVirtual,omitempty"`
	// Hypervisor connection connection type display name.
	DisplayConnectionType string `json:"DisplayConnectionType"`
	// Indicates whether the hypervisor connection is broken.
	IsBroken bool `json:"IsBroken"`
	// Hypervisor connection broken errors.
	Errors []string `json:"Errors"`
	// Current hypervisor connection error status.
	ErrorState string `json:"ErrorState"`
	// The metadata of hypervisor connections.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	// Hypervisor resource pools ref response list.
	ResourcePools []HypervisorBaseResponseModel `json:"ResourcePools,omitempty"`
}

// NewHypervisorsAndResourcePoolsResponseModel instantiates a new HypervisorsAndResourcePoolsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorsAndResourcePoolsResponseModel(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel, displayConnectionType string, isBroken bool, errors []string, errorState string) *HypervisorsAndResourcePoolsResponseModel {
	this := HypervisorsAndResourcePoolsResponseModel{}
	this.ConnectionType = connectionType
	this.Addresses = addresses
	this.InMaintenanceMode = inMaintenanceMode
	this.PluginId = pluginId
	this.Scopes = scopes
	this.UsesCloudInfrastructure = usesCloudInfrastructure
	this.Zone = zone
	this.DisplayConnectionType = displayConnectionType
	this.IsBroken = isBroken
	this.Errors = errors
	this.ErrorState = errorState
	return &this
}

// NewHypervisorsAndResourcePoolsResponseModelWithDefaults instantiates a new HypervisorsAndResourcePoolsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorsAndResourcePoolsResponseModelWithDefaults() *HypervisorsAndResourcePoolsResponseModel {
	this := HypervisorsAndResourcePoolsResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetXDPath returns the XDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath.Get()) {
		var ret string
		return ret
	}
	return *o.XDPath.Get()
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XDPath.Get(), o.XDPath.IsSet()
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasXDPath() bool {
	if o != nil && o.XDPath.IsSet() {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given NullableString and assigns it to the XDPath field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetXDPath(v string) {
	o.XDPath.Set(&v)
}
// SetXDPathNil sets the value for XDPath to be an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) SetXDPathNil() {
	o.XDPath.Set(nil)
}

// UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) UnsetXDPath() {
	o.XDPath.Unset()
}

// GetConnectionType returns the ConnectionType field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetAddresses returns the Addresses field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetAddresses(v []string) {
	o.Addresses = v
}

// GetInMaintenanceMode returns the InMaintenanceMode field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetInMaintenanceMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InMaintenanceMode
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InMaintenanceMode, true
}

// SetInMaintenanceMode sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode = v
}

// GetUnentitled returns the Unentitled field value if set, zero value otherwise.
func (o *HypervisorsAndResourcePoolsResponseModel) GetUnentitled() bool {
	if o == nil || IsNil(o.Unentitled) {
		var ret bool
		return ret
	}
	return *o.Unentitled
}

// GetUnentitledOk returns a tuple with the Unentitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetUnentitledOk() (*bool, bool) {
	if o == nil || IsNil(o.Unentitled) {
		return nil, false
	}
	return o.Unentitled, true
}

// HasUnentitled returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasUnentitled() bool {
	if o != nil && !IsNil(o.Unentitled) {
		return true
	}

	return false
}

// SetUnentitled gets a reference to the given bool and assigns it to the Unentitled field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetUnentitled(v bool) {
	o.Unentitled = &v
}

// GetPluginId returns the PluginId field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetPluginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginId, true
}

// SetPluginId sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetPluginId(v string) {
	o.PluginId = v
}

// GetScopes returns the Scopes field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetScopes() []ScopeResponseModel {
	if o == nil {
		var ret []ScopeResponseModel
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetScopesOk() ([]ScopeResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetScopes(v []ScopeResponseModel) {
	o.Scopes = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetTenants() []RefResponseModel {
	if o == nil {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasTenants() bool {
	if o != nil && IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

// GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetUsesCloudInfrastructure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsesCloudInfrastructure
}

// GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsesCloudInfrastructure, true
}

// SetUsesCloudInfrastructure sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetUsesCloudInfrastructure(v bool) {
	o.UsesCloudInfrastructure = v
}

// GetZone returns the Zone field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetZone() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetZoneOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetZone(v RefResponseModel) {
	o.Zone = v
}

// GetFault returns the Fault field value if set, zero value otherwise.
func (o *HypervisorsAndResourcePoolsResponseModel) GetFault() HypervisorFaultResponseModel {
	if o == nil || IsNil(o.Fault) {
		var ret HypervisorFaultResponseModel
		return ret
	}
	return *o.Fault
}

// GetFaultOk returns a tuple with the Fault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool) {
	if o == nil || IsNil(o.Fault) {
		return nil, false
	}
	return o.Fault, true
}

// HasFault returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasFault() bool {
	if o != nil && !IsNil(o.Fault) {
		return true
	}

	return false
}

// SetFault gets a reference to the given HypervisorFaultResponseModel and assigns it to the Fault field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetFault(v HypervisorFaultResponseModel) {
	o.Fault = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties.Get()) {
		var ret string
		return ret
	}
	return *o.CustomProperties.Get()
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomProperties.Get(), o.CustomProperties.IsSet()
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasCustomProperties() bool {
	if o != nil && o.CustomProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given NullableString and assigns it to the CustomProperties field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetCustomProperties(v string) {
	o.CustomProperties.Set(&v)
}
// SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) SetCustomPropertiesNil() {
	o.CustomProperties.Set(nil)
}

// UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) UnsetCustomProperties() {
	o.CustomProperties.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetUid(v int32) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *HypervisorsAndResourcePoolsResponseModel) UnsetUid() {
	o.Uid.Unset()
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *HypervisorsAndResourcePoolsResponseModel) GetIsVirtual() bool {
	if o == nil || IsNil(o.IsVirtual) {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetIsVirtualOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVirtual) {
		return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasIsVirtual() bool {
	if o != nil && !IsNil(o.IsVirtual) {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

// GetDisplayConnectionType returns the DisplayConnectionType field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetDisplayConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayConnectionType
}

// GetDisplayConnectionTypeOk returns a tuple with the DisplayConnectionType field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetDisplayConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayConnectionType, true
}

// SetDisplayConnectionType sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetDisplayConnectionType(v string) {
	o.DisplayConnectionType = v
}

// GetIsBroken returns the IsBroken field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetIsBroken() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBroken
}

// GetIsBrokenOk returns a tuple with the IsBroken field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetIsBrokenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBroken, true
}

// SetIsBroken sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetIsBroken(v bool) {
	o.IsBroken = v
}

// GetErrors returns the Errors field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetErrors(v []string) {
	o.Errors = v
}

// GetErrorState returns the ErrorState field value
func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value
// and a boolean to check if the value has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorState, true
}

// SetErrorState sets field value
func (o *HypervisorsAndResourcePoolsResponseModel) SetErrorState(v string) {
	o.ErrorState = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetMetadata() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetResourcePools returns the ResourcePools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorsAndResourcePoolsResponseModel) GetResourcePools() []HypervisorBaseResponseModel {
	if o == nil {
		var ret []HypervisorBaseResponseModel
		return ret
	}
	return o.ResourcePools
}

// GetResourcePoolsOk returns a tuple with the ResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorsAndResourcePoolsResponseModel) GetResourcePoolsOk() ([]HypervisorBaseResponseModel, bool) {
	if o == nil || IsNil(o.ResourcePools) {
		return nil, false
	}
	return o.ResourcePools, true
}

// HasResourcePools returns a boolean if a field has been set.
func (o *HypervisorsAndResourcePoolsResponseModel) HasResourcePools() bool {
	if o != nil && IsNil(o.ResourcePools) {
		return true
	}

	return false
}

// SetResourcePools gets a reference to the given []HypervisorBaseResponseModel and assigns it to the ResourcePools field.
func (o *HypervisorsAndResourcePoolsResponseModel) SetResourcePools(v []HypervisorBaseResponseModel) {
	o.ResourcePools = v
}

func (o HypervisorsAndResourcePoolsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorsAndResourcePoolsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.XDPath.IsSet() {
		toSerialize["XDPath"] = o.XDPath.Get()
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	toSerialize["Addresses"] = o.Addresses
	toSerialize["InMaintenanceMode"] = o.InMaintenanceMode
	if !IsNil(o.Unentitled) {
		toSerialize["Unentitled"] = o.Unentitled
	}
	toSerialize["PluginId"] = o.PluginId
	toSerialize["Scopes"] = o.Scopes
	if o.Tenants != nil {
		toSerialize["Tenants"] = o.Tenants
	}
	toSerialize["UsesCloudInfrastructure"] = o.UsesCloudInfrastructure
	toSerialize["Zone"] = o.Zone
	if !IsNil(o.Fault) {
		toSerialize["Fault"] = o.Fault
	}
	if o.CustomProperties.IsSet() {
		toSerialize["CustomProperties"] = o.CustomProperties.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["Uid"] = o.Uid.Get()
	}
	if !IsNil(o.IsVirtual) {
		toSerialize["IsVirtual"] = o.IsVirtual
	}
	toSerialize["DisplayConnectionType"] = o.DisplayConnectionType
	toSerialize["IsBroken"] = o.IsBroken
	toSerialize["Errors"] = o.Errors
	toSerialize["ErrorState"] = o.ErrorState
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if o.ResourcePools != nil {
		toSerialize["ResourcePools"] = o.ResourcePools
	}
	return toSerialize, nil
}

type NullableHypervisorsAndResourcePoolsResponseModel struct {
	value *HypervisorsAndResourcePoolsResponseModel
	isSet bool
}

func (v NullableHypervisorsAndResourcePoolsResponseModel) Get() *HypervisorsAndResourcePoolsResponseModel {
	return v.value
}

func (v *NullableHypervisorsAndResourcePoolsResponseModel) Set(val *HypervisorsAndResourcePoolsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorsAndResourcePoolsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorsAndResourcePoolsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorsAndResourcePoolsResponseModel(val *HypervisorsAndResourcePoolsResponseModel) *NullableHypervisorsAndResourcePoolsResponseModel {
	return &NullableHypervisorsAndResourcePoolsResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorsAndResourcePoolsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorsAndResourcePoolsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


