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

// checks if the HypervisorResourcePoolDetailResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolDetailResponseModel{}

// HypervisorResourcePoolDetailResponseModel struct for HypervisorResourcePoolDetailResponseModel
type HypervisorResourcePoolDetailResponseModel struct {
	VirtualPrivateCloud HypervisorResourceRefResponseModel `json:"VirtualPrivateCloud"`
	AvailabilityZone HypervisorResourceRefResponseModel `json:"AvailabilityZone"`
	// List of networks that may be used within the resource pool.
	Networks []HypervisorResourceRefResponseModel `json:"Networks"`
	Region HypervisorResourceRefResponseModel `json:"Region"`
	VirtualNetwork HypervisorResourceRefResponseModel `json:"VirtualNetwork"`
	// List of subnets in the VirtualNetwork that may be used within the resource pool.
	Subnets []HypervisorResourceRefResponseModel `json:"Subnets"`
	Project HypervisorResourceRefResponseModel `json:"Project"`
	RootPath *HypervisorResourceRefResponseModel `json:"RootPath,omitempty"`
	// List of hypervisor-connected storage in the resource pool that is used for OS disks of virtual machines.
	Storage []HypervisorStorageResourceResponseModel `json:"Storage"`
	// List of hypervisor-connected storage in the resource pool that is used for temporary data storage for virtual machines.
	TemporaryStorage []HypervisorStorageResourceResponseModel `json:"TemporaryStorage"`
	// Indicates whether virtual machines created within this resource pool will use local storage caching for their disk images.
	UseLocalStorageCaching *bool `json:"UseLocalStorageCaching,omitempty"`
	// Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom.
	CustomProperties *string `json:"CustomProperties,omitempty"`
	// List of hypervisor-connected storage in the resource pool that is used for personal v disk data storage for virtual machines.
	PersonalvDiskStorage []HypervisorStorageResourceResponseModel `json:"PersonalvDiskStorage,omitempty"`
	// Id of the resource.
	Id *string `json:"Id,omitempty"`
	// Name of the resource.
	Name *string `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath *string `json:"XDPath,omitempty"`
	HypervisorConnection HypervisorResourcePoolResponseModelAllOfHypervisorConnection `json:"HypervisorConnection"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	DefaultNetwork HypervisorResourcePoolResponseModelAllOfDefaultNetwork `json:"DefaultNetwork"`
	// Indicates whether new virtual machines are tagged with metadata from the hypervisor.
	VMTaggingEnabled bool `json:"VMTaggingEnabled"`
	// Hypervisor resourcePool RootPath.
	ResourcePoolRootPath *string `json:"ResourcePoolRootPath,omitempty"`
	// Hypervisor resourcePool RootId.
	ResourcePoolRootId *string `json:"ResourcePoolRootId,omitempty"`
	// GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types.
	GpuTypes []HypervisorResourceRefResponseModel `json:"GpuTypes,omitempty"`
	// If the hypervisor resource pool use ExplicitStorage.
	UsesExplicitStorage *bool `json:"UsesExplicitStorage,omitempty"`
	// Metadata for hypervisor resource pool. 
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
}

// NewHypervisorResourcePoolDetailResponseModel instantiates a new HypervisorResourcePoolDetailResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolDetailResponseModel(virtualPrivateCloud HypervisorResourceRefResponseModel, availabilityZone HypervisorResourceRefResponseModel, networks []HypervisorResourceRefResponseModel, region HypervisorResourceRefResponseModel, virtualNetwork HypervisorResourceRefResponseModel, subnets []HypervisorResourceRefResponseModel, project HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, hypervisorConnection HypervisorResourcePoolResponseModelAllOfHypervisorConnection, connectionType HypervisorConnectionType, defaultNetwork HypervisorResourcePoolResponseModelAllOfDefaultNetwork, vMTaggingEnabled bool) *HypervisorResourcePoolDetailResponseModel {
	this := HypervisorResourcePoolDetailResponseModel{}
	this.HypervisorConnection = hypervisorConnection
	this.ConnectionType = connectionType
	this.DefaultNetwork = defaultNetwork
	this.VMTaggingEnabled = vMTaggingEnabled
	return &this
}

// NewHypervisorResourcePoolDetailResponseModelWithDefaults instantiates a new HypervisorResourcePoolDetailResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolDetailResponseModelWithDefaults() *HypervisorResourcePoolDetailResponseModel {
	this := HypervisorResourcePoolDetailResponseModel{}
	return &this
}

// GetVirtualPrivateCloud returns the VirtualPrivateCloud field value
func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualPrivateCloud() HypervisorResourceRefResponseModel {
	if o == nil {
		var ret HypervisorResourceRefResponseModel
		return ret
	}

	return o.VirtualPrivateCloud
}

// GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualPrivateCloudOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualPrivateCloud, true
}

// SetVirtualPrivateCloud sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetVirtualPrivateCloud(v HypervisorResourceRefResponseModel) {
	o.VirtualPrivateCloud = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *HypervisorResourcePoolDetailResponseModel) GetAvailabilityZone() HypervisorResourceRefResponseModel {
	if o == nil {
		var ret HypervisorResourceRefResponseModel
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetAvailabilityZoneOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetAvailabilityZone(v HypervisorResourceRefResponseModel) {
	o.AvailabilityZone = v
}

// GetNetworks returns the Networks field value
func (o *HypervisorResourcePoolDetailResponseModel) GetNetworks() []HypervisorResourceRefResponseModel {
	if o == nil {
		var ret []HypervisorResourceRefResponseModel
		return ret
	}

	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetNetworksOk() ([]HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Networks, true
}

// SetNetworks sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetNetworks(v []HypervisorResourceRefResponseModel) {
	o.Networks = v
}

// GetRegion returns the Region field value
func (o *HypervisorResourcePoolDetailResponseModel) GetRegion() HypervisorResourceRefResponseModel {
	if o == nil {
		var ret HypervisorResourceRefResponseModel
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetRegionOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetRegion(v HypervisorResourceRefResponseModel) {
	o.Region = v
}

// GetVirtualNetwork returns the VirtualNetwork field value
func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualNetwork() HypervisorResourceRefResponseModel {
	if o == nil {
		var ret HypervisorResourceRefResponseModel
		return ret
	}

	return o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualNetworkOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetwork, true
}

// SetVirtualNetwork sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetVirtualNetwork(v HypervisorResourceRefResponseModel) {
	o.VirtualNetwork = v
}

// GetSubnets returns the Subnets field value
func (o *HypervisorResourcePoolDetailResponseModel) GetSubnets() []HypervisorResourceRefResponseModel {
	if o == nil {
		var ret []HypervisorResourceRefResponseModel
		return ret
	}

	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetSubnetsOk() ([]HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnets, true
}

// SetSubnets sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetSubnets(v []HypervisorResourceRefResponseModel) {
	o.Subnets = v
}

// GetProject returns the Project field value
func (o *HypervisorResourcePoolDetailResponseModel) GetProject() HypervisorResourceRefResponseModel {
	if o == nil {
		var ret HypervisorResourceRefResponseModel
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetProjectOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetProject(v HypervisorResourceRefResponseModel) {
	o.Project = v
}

// GetRootPath returns the RootPath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetRootPath() HypervisorResourceRefResponseModel {
	if o == nil || IsNil(o.RootPath) {
		var ret HypervisorResourceRefResponseModel
		return ret
	}
	return *o.RootPath
}

// GetRootPathOk returns a tuple with the RootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetRootPathOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil || IsNil(o.RootPath) {
		return nil, false
	}
	return o.RootPath, true
}

// HasRootPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasRootPath() bool {
	if o != nil && !IsNil(o.RootPath) {
		return true
	}

	return false
}

// SetRootPath gets a reference to the given HypervisorResourceRefResponseModel and assigns it to the RootPath field.
func (o *HypervisorResourcePoolDetailResponseModel) SetRootPath(v HypervisorResourceRefResponseModel) {
	o.RootPath = &v
}

// GetStorage returns the Storage field value
func (o *HypervisorResourcePoolDetailResponseModel) GetStorage() []HypervisorStorageResourceResponseModel {
	if o == nil {
		var ret []HypervisorStorageResourceResponseModel
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetStorageOk() ([]HypervisorStorageResourceResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage, true
}

// SetStorage sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetStorage(v []HypervisorStorageResourceResponseModel) {
	o.Storage = v
}

// GetTemporaryStorage returns the TemporaryStorage field value
func (o *HypervisorResourcePoolDetailResponseModel) GetTemporaryStorage() []HypervisorStorageResourceResponseModel {
	if o == nil {
		var ret []HypervisorStorageResourceResponseModel
		return ret
	}

	return o.TemporaryStorage
}

// GetTemporaryStorageOk returns a tuple with the TemporaryStorage field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetTemporaryStorageOk() ([]HypervisorStorageResourceResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemporaryStorage, true
}

// SetTemporaryStorage sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel) {
	o.TemporaryStorage = v
}

// GetUseLocalStorageCaching returns the UseLocalStorageCaching field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetUseLocalStorageCaching() bool {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		var ret bool
		return ret
	}
	return *o.UseLocalStorageCaching
}

// GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetUseLocalStorageCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		return nil, false
	}
	return o.UseLocalStorageCaching, true
}

// HasUseLocalStorageCaching returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasUseLocalStorageCaching() bool {
	if o != nil && !IsNil(o.UseLocalStorageCaching) {
		return true
	}

	return false
}

// SetUseLocalStorageCaching gets a reference to the given bool and assigns it to the UseLocalStorageCaching field.
func (o *HypervisorResourcePoolDetailResponseModel) SetUseLocalStorageCaching(v bool) {
	o.UseLocalStorageCaching = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties) {
		var ret string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given string and assigns it to the CustomProperties field.
func (o *HypervisorResourcePoolDetailResponseModel) SetCustomProperties(v string) {
	o.CustomProperties = &v
}

// GetPersonalvDiskStorage returns the PersonalvDiskStorage field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		var ret []HypervisorStorageResourceResponseModel
		return ret
	}
	return o.PersonalvDiskStorage
}

// GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetPersonalvDiskStorageOk() ([]HypervisorStorageResourceResponseModel, bool) {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		return nil, false
	}
	return o.PersonalvDiskStorage, true
}

// HasPersonalvDiskStorage returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasPersonalvDiskStorage() bool {
	if o != nil && !IsNil(o.PersonalvDiskStorage) {
		return true
	}

	return false
}

// SetPersonalvDiskStorage gets a reference to the given []HypervisorStorageResourceResponseModel and assigns it to the PersonalvDiskStorage field.
func (o *HypervisorResourcePoolDetailResponseModel) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel) {
	o.PersonalvDiskStorage = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HypervisorResourcePoolDetailResponseModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HypervisorResourcePoolDetailResponseModel) SetName(v string) {
	o.Name = &v
}

// GetXDPath returns the XDPath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath) {
		var ret string
		return ret
	}
	return *o.XDPath
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.XDPath) {
		return nil, false
	}
	return o.XDPath, true
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasXDPath() bool {
	if o != nil && !IsNil(o.XDPath) {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given string and assigns it to the XDPath field.
func (o *HypervisorResourcePoolDetailResponseModel) SetXDPath(v string) {
	o.XDPath = &v
}

// GetHypervisorConnection returns the HypervisorConnection field value
func (o *HypervisorResourcePoolDetailResponseModel) GetHypervisorConnection() HypervisorResourcePoolResponseModelAllOfHypervisorConnection {
	if o == nil {
		var ret HypervisorResourcePoolResponseModelAllOfHypervisorConnection
		return ret
	}

	return o.HypervisorConnection
}

// GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetHypervisorConnectionOk() (*HypervisorResourcePoolResponseModelAllOfHypervisorConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HypervisorConnection, true
}

// SetHypervisorConnection sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetHypervisorConnection(v HypervisorResourcePoolResponseModelAllOfHypervisorConnection) {
	o.HypervisorConnection = v
}

// GetConnectionType returns the ConnectionType field value
func (o *HypervisorResourcePoolDetailResponseModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetDefaultNetwork returns the DefaultNetwork field value
func (o *HypervisorResourcePoolDetailResponseModel) GetDefaultNetwork() HypervisorResourcePoolResponseModelAllOfDefaultNetwork {
	if o == nil {
		var ret HypervisorResourcePoolResponseModelAllOfDefaultNetwork
		return ret
	}

	return o.DefaultNetwork
}

// GetDefaultNetworkOk returns a tuple with the DefaultNetwork field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetDefaultNetworkOk() (*HypervisorResourcePoolResponseModelAllOfDefaultNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultNetwork, true
}

// SetDefaultNetwork sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetDefaultNetwork(v HypervisorResourcePoolResponseModelAllOfDefaultNetwork) {
	o.DefaultNetwork = v
}

// GetVMTaggingEnabled returns the VMTaggingEnabled field value
func (o *HypervisorResourcePoolDetailResponseModel) GetVMTaggingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VMTaggingEnabled
}

// GetVMTaggingEnabledOk returns a tuple with the VMTaggingEnabled field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetVMTaggingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VMTaggingEnabled, true
}

// SetVMTaggingEnabled sets field value
func (o *HypervisorResourcePoolDetailResponseModel) SetVMTaggingEnabled(v bool) {
	o.VMTaggingEnabled = v
}

// GetResourcePoolRootPath returns the ResourcePoolRootPath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootPath() string {
	if o == nil || IsNil(o.ResourcePoolRootPath) {
		var ret string
		return ret
	}
	return *o.ResourcePoolRootPath
}

// GetResourcePoolRootPathOk returns a tuple with the ResourcePoolRootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootPathOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolRootPath) {
		return nil, false
	}
	return o.ResourcePoolRootPath, true
}

// HasResourcePoolRootPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasResourcePoolRootPath() bool {
	if o != nil && !IsNil(o.ResourcePoolRootPath) {
		return true
	}

	return false
}

// SetResourcePoolRootPath gets a reference to the given string and assigns it to the ResourcePoolRootPath field.
func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootPath(v string) {
	o.ResourcePoolRootPath = &v
}

// GetResourcePoolRootId returns the ResourcePoolRootId field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootId() string {
	if o == nil || IsNil(o.ResourcePoolRootId) {
		var ret string
		return ret
	}
	return *o.ResourcePoolRootId
}

// GetResourcePoolRootIdOk returns a tuple with the ResourcePoolRootId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolRootId) {
		return nil, false
	}
	return o.ResourcePoolRootId, true
}

// HasResourcePoolRootId returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasResourcePoolRootId() bool {
	if o != nil && !IsNil(o.ResourcePoolRootId) {
		return true
	}

	return false
}

// SetResourcePoolRootId gets a reference to the given string and assigns it to the ResourcePoolRootId field.
func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootId(v string) {
	o.ResourcePoolRootId = &v
}

// GetGpuTypes returns the GpuTypes field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetGpuTypes() []HypervisorResourceRefResponseModel {
	if o == nil || IsNil(o.GpuTypes) {
		var ret []HypervisorResourceRefResponseModel
		return ret
	}
	return o.GpuTypes
}

// GetGpuTypesOk returns a tuple with the GpuTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetGpuTypesOk() ([]HypervisorResourceRefResponseModel, bool) {
	if o == nil || IsNil(o.GpuTypes) {
		return nil, false
	}
	return o.GpuTypes, true
}

// HasGpuTypes returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasGpuTypes() bool {
	if o != nil && !IsNil(o.GpuTypes) {
		return true
	}

	return false
}

// SetGpuTypes gets a reference to the given []HypervisorResourceRefResponseModel and assigns it to the GpuTypes field.
func (o *HypervisorResourcePoolDetailResponseModel) SetGpuTypes(v []HypervisorResourceRefResponseModel) {
	o.GpuTypes = v
}

// GetUsesExplicitStorage returns the UsesExplicitStorage field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetUsesExplicitStorage() bool {
	if o == nil || IsNil(o.UsesExplicitStorage) {
		var ret bool
		return ret
	}
	return *o.UsesExplicitStorage
}

// GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetUsesExplicitStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesExplicitStorage) {
		return nil, false
	}
	return o.UsesExplicitStorage, true
}

// HasUsesExplicitStorage returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasUsesExplicitStorage() bool {
	if o != nil && !IsNil(o.UsesExplicitStorage) {
		return true
	}

	return false
}

// SetUsesExplicitStorage gets a reference to the given bool and assigns it to the UsesExplicitStorage field.
func (o *HypervisorResourcePoolDetailResponseModel) SetUsesExplicitStorage(v bool) {
	o.UsesExplicitStorage = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModel) GetMetadata() []NameValueStringPairModel {
	if o == nil || IsNil(o.Metadata) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *HypervisorResourcePoolDetailResponseModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

func (o HypervisorResourcePoolDetailResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolDetailResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["VirtualPrivateCloud"] = o.VirtualPrivateCloud
	toSerialize["AvailabilityZone"] = o.AvailabilityZone
	toSerialize["Networks"] = o.Networks
	toSerialize["Region"] = o.Region
	toSerialize["VirtualNetwork"] = o.VirtualNetwork
	toSerialize["Subnets"] = o.Subnets
	toSerialize["Project"] = o.Project
	if !IsNil(o.RootPath) {
		toSerialize["RootPath"] = o.RootPath
	}
	toSerialize["Storage"] = o.Storage
	toSerialize["TemporaryStorage"] = o.TemporaryStorage
	if !IsNil(o.UseLocalStorageCaching) {
		toSerialize["UseLocalStorageCaching"] = o.UseLocalStorageCaching
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["CustomProperties"] = o.CustomProperties
	}
	if !IsNil(o.PersonalvDiskStorage) {
		toSerialize["PersonalvDiskStorage"] = o.PersonalvDiskStorage
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.XDPath) {
		toSerialize["XDPath"] = o.XDPath
	}
	toSerialize["HypervisorConnection"] = o.HypervisorConnection
	toSerialize["ConnectionType"] = o.ConnectionType
	toSerialize["DefaultNetwork"] = o.DefaultNetwork
	toSerialize["VMTaggingEnabled"] = o.VMTaggingEnabled
	if !IsNil(o.ResourcePoolRootPath) {
		toSerialize["ResourcePoolRootPath"] = o.ResourcePoolRootPath
	}
	if !IsNil(o.ResourcePoolRootId) {
		toSerialize["ResourcePoolRootId"] = o.ResourcePoolRootId
	}
	if !IsNil(o.GpuTypes) {
		toSerialize["GpuTypes"] = o.GpuTypes
	}
	if !IsNil(o.UsesExplicitStorage) {
		toSerialize["UsesExplicitStorage"] = o.UsesExplicitStorage
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableHypervisorResourcePoolDetailResponseModel struct {
	value *HypervisorResourcePoolDetailResponseModel
	isSet bool
}

func (v NullableHypervisorResourcePoolDetailResponseModel) Get() *HypervisorResourcePoolDetailResponseModel {
	return v.value
}

func (v *NullableHypervisorResourcePoolDetailResponseModel) Set(val *HypervisorResourcePoolDetailResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolDetailResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolDetailResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolDetailResponseModel(val *HypervisorResourcePoolDetailResponseModel) *NullableHypervisorResourcePoolDetailResponseModel {
	return &NullableHypervisorResourcePoolDetailResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolDetailResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolDetailResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


