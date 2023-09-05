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

// checks if the HypervisorTemplateResourceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorTemplateResourceResponseModel{}

// HypervisorTemplateResourceResponseModel struct for HypervisorTemplateResourceResponseModel
type HypervisorTemplateResourceResponseModel struct {
	// Id of the resource.
	Id *string `json:"Id,omitempty"`
	// Name of the resource.
	Name *string `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath *string `json:"XDPath,omitempty"`
	// Path to the object, relative to the resource pool from which it was queried. `{{vm name}}.vm/{{snapshot name}}.snapshot`
	RelativePath *string `json:"RelativePath,omitempty"`
	// Full path to the resource, including the hypervisor, relative to the root of the API. Example: `Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources/{{RelativePath}}`
	FullRelativePath *string `json:"FullRelativePath,omitempty"`
	// Type of resource.
	ResourceType string `json:"ResourceType"`
	// The type name of the hypervisor resource object.
	ObjectTypeName *string `json:"ObjectTypeName,omitempty"`
	ResourceContainer *HypervisorResourceRefResponseModelAllOfResourceContainer `json:"ResourceContainer,omitempty"`
	// Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}` This value 
	ResourcePoolXDPath *string `json:"ResourcePoolXDPath,omitempty"`
	// Name of the resource, with the type concatenated. i.e. \"name.type\".
	FullName string `json:"FullName"`
	// Indicates whether the resource is a container.
	IsContainer bool `json:"IsContainer"`
	// Child resources, when the resource is a container.
	Children []HypervisorResourceResponseModel `json:"Children,omitempty"`
	// Indicates whether the resource is a machine.
	IsMachine bool `json:"IsMachine"`
	// Indicates whether the resource can have a snapshot taken.
	IsSnapshotable bool `json:"IsSnapshotable"`
	// Path to the resource, relative to the special \"AllResources\" resource pool associated with the hypervisor.
	AllResourcesRelativePath string `json:"AllResourcesRelativePath"`
	ResourcePool HypervisorResourceResponseModelAllOfResourcePool `json:"ResourcePool"`
	// Indicates whether the object is a valid symbolic link.
	IsSymLink bool `json:"IsSymLink"`
	// Additional data about the object in the form of key-value pairs.
	AdditionalData []NameValueStringPairModel `json:"AdditionalData,omitempty"`
	// Number of CPUs, if known.
	CpuCount *int32 `json:"CpuCount,omitempty"`
	// Memory in megabytes, if known.
	MemoryMB *int32 `json:"MemoryMB,omitempty"`
	// Hard disk size in gigabytes, if known.
	HardDiskSizeGB *int32 `json:"HardDiskSizeGB,omitempty"`
	// Minimum memory required to run this VM or snapshot, in megabytes, if known.
	MinMemoryMB *int32 `json:"MinMemoryMB,omitempty"`
	// Network mappings associated with the VM, if known.
	NetworkMappings []NetworkMapResponseModel `json:"NetworkMappings,omitempty"`
	// Disks attached to the VM, if known.
	AttachedDisks []AttachedDiskResponseModel `json:"AttachedDisks,omitempty"`
	// The account ID for the owner of this template.
	Owner *string `json:"Owner,omitempty"`
	// Indicates whether this is a Windows OS template, if known.
	IsWindowsTemplate *bool `json:"IsWindowsTemplate,omitempty"`
	// Human-readable description of this template, as supplied when the offering was created in the cloud management API or console.
	Description *string `json:"Description,omitempty"`
	// Indicates whether this template has a persistent root volume (eg. is an EBS-backed image on AWS).
	HasPersistentRootVolume bool `json:"HasPersistentRootVolume"`
}

// NewHypervisorTemplateResourceResponseModel instantiates a new HypervisorTemplateResourceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorTemplateResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourceResponseModelAllOfResourcePool, isSymLink bool, hasPersistentRootVolume bool) *HypervisorTemplateResourceResponseModel {
	this := HypervisorTemplateResourceResponseModel{}
	this.ResourceType = resourceType
	this.FullName = fullName
	this.IsContainer = isContainer
	this.IsMachine = isMachine
	this.IsSnapshotable = isSnapshotable
	this.AllResourcesRelativePath = allResourcesRelativePath
	this.ResourcePool = resourcePool
	this.IsSymLink = isSymLink
	this.HasPersistentRootVolume = hasPersistentRootVolume
	return &this
}

// NewHypervisorTemplateResourceResponseModelWithDefaults instantiates a new HypervisorTemplateResourceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorTemplateResourceResponseModelWithDefaults() *HypervisorTemplateResourceResponseModel {
	this := HypervisorTemplateResourceResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HypervisorTemplateResourceResponseModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HypervisorTemplateResourceResponseModel) SetName(v string) {
	o.Name = &v
}

// GetXDPath returns the XDPath field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath) {
		var ret string
		return ret
	}
	return *o.XDPath
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.XDPath) {
		return nil, false
	}
	return o.XDPath, true
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasXDPath() bool {
	if o != nil && !IsNil(o.XDPath) {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given string and assigns it to the XDPath field.
func (o *HypervisorTemplateResourceResponseModel) SetXDPath(v string) {
	o.XDPath = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *HypervisorTemplateResourceResponseModel) SetRelativePath(v string) {
	o.RelativePath = &v
}

// GetFullRelativePath returns the FullRelativePath field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetFullRelativePath() string {
	if o == nil || IsNil(o.FullRelativePath) {
		var ret string
		return ret
	}
	return *o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetFullRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.FullRelativePath) {
		return nil, false
	}
	return o.FullRelativePath, true
}

// HasFullRelativePath returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasFullRelativePath() bool {
	if o != nil && !IsNil(o.FullRelativePath) {
		return true
	}

	return false
}

// SetFullRelativePath gets a reference to the given string and assigns it to the FullRelativePath field.
func (o *HypervisorTemplateResourceResponseModel) SetFullRelativePath(v string) {
	o.FullRelativePath = &v
}

// GetResourceType returns the ResourceType field value
func (o *HypervisorTemplateResourceResponseModel) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *HypervisorTemplateResourceResponseModel) SetResourceType(v string) {
	o.ResourceType = v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetObjectTypeName() string {
	if o == nil || IsNil(o.ObjectTypeName) {
		var ret string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetObjectTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectTypeName) {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasObjectTypeName() bool {
	if o != nil && !IsNil(o.ObjectTypeName) {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given string and assigns it to the ObjectTypeName field.
func (o *HypervisorTemplateResourceResponseModel) SetObjectTypeName(v string) {
	o.ObjectTypeName = &v
}

// GetResourceContainer returns the ResourceContainer field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer {
	if o == nil || IsNil(o.ResourceContainer) {
		var ret HypervisorResourceRefResponseModelAllOfResourceContainer
		return ret
	}
	return *o.ResourceContainer
}

// GetResourceContainerOk returns a tuple with the ResourceContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool) {
	if o == nil || IsNil(o.ResourceContainer) {
		return nil, false
	}
	return o.ResourceContainer, true
}

// HasResourceContainer returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasResourceContainer() bool {
	if o != nil && !IsNil(o.ResourceContainer) {
		return true
	}

	return false
}

// SetResourceContainer gets a reference to the given HypervisorResourceRefResponseModelAllOfResourceContainer and assigns it to the ResourceContainer field.
func (o *HypervisorTemplateResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer) {
	o.ResourceContainer = &v
}

// GetResourcePoolXDPath returns the ResourcePoolXDPath field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolXDPath() string {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		var ret string
		return ret
	}
	return *o.ResourcePoolXDPath
}

// GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		return nil, false
	}
	return o.ResourcePoolXDPath, true
}

// HasResourcePoolXDPath returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasResourcePoolXDPath() bool {
	if o != nil && !IsNil(o.ResourcePoolXDPath) {
		return true
	}

	return false
}

// SetResourcePoolXDPath gets a reference to the given string and assigns it to the ResourcePoolXDPath field.
func (o *HypervisorTemplateResourceResponseModel) SetResourcePoolXDPath(v string) {
	o.ResourcePoolXDPath = &v
}

// GetFullName returns the FullName field value
func (o *HypervisorTemplateResourceResponseModel) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *HypervisorTemplateResourceResponseModel) SetFullName(v string) {
	o.FullName = v
}

// GetIsContainer returns the IsContainer field value
func (o *HypervisorTemplateResourceResponseModel) GetIsContainer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIsContainerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsContainer, true
}

// SetIsContainer sets field value
func (o *HypervisorTemplateResourceResponseModel) SetIsContainer(v bool) {
	o.IsContainer = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetChildren() []HypervisorResourceResponseModel {
	if o == nil || IsNil(o.Children) {
		var ret []HypervisorResourceResponseModel
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetChildrenOk() ([]HypervisorResourceResponseModel, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []HypervisorResourceResponseModel and assigns it to the Children field.
func (o *HypervisorTemplateResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel) {
	o.Children = v
}

// GetIsMachine returns the IsMachine field value
func (o *HypervisorTemplateResourceResponseModel) GetIsMachine() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMachine
}

// GetIsMachineOk returns a tuple with the IsMachine field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIsMachineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMachine, true
}

// SetIsMachine sets field value
func (o *HypervisorTemplateResourceResponseModel) SetIsMachine(v bool) {
	o.IsMachine = v
}

// GetIsSnapshotable returns the IsSnapshotable field value
func (o *HypervisorTemplateResourceResponseModel) GetIsSnapshotable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSnapshotable
}

// GetIsSnapshotableOk returns a tuple with the IsSnapshotable field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIsSnapshotableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSnapshotable, true
}

// SetIsSnapshotable sets field value
func (o *HypervisorTemplateResourceResponseModel) SetIsSnapshotable(v bool) {
	o.IsSnapshotable = v
}

// GetAllResourcesRelativePath returns the AllResourcesRelativePath field value
func (o *HypervisorTemplateResourceResponseModel) GetAllResourcesRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AllResourcesRelativePath
}

// GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllResourcesRelativePath, true
}

// SetAllResourcesRelativePath sets field value
func (o *HypervisorTemplateResourceResponseModel) SetAllResourcesRelativePath(v string) {
	o.AllResourcesRelativePath = v
}

// GetResourcePool returns the ResourcePool field value
func (o *HypervisorTemplateResourceResponseModel) GetResourcePool() HypervisorResourceResponseModelAllOfResourcePool {
	if o == nil {
		var ret HypervisorResourceResponseModelAllOfResourcePool
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolOk() (*HypervisorResourceResponseModelAllOfResourcePool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *HypervisorTemplateResourceResponseModel) SetResourcePool(v HypervisorResourceResponseModelAllOfResourcePool) {
	o.ResourcePool = v
}

// GetIsSymLink returns the IsSymLink field value
func (o *HypervisorTemplateResourceResponseModel) GetIsSymLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSymLink
}

// GetIsSymLinkOk returns a tuple with the IsSymLink field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIsSymLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSymLink, true
}

// SetIsSymLink sets field value
func (o *HypervisorTemplateResourceResponseModel) SetIsSymLink(v bool) {
	o.IsSymLink = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetAdditionalData() []NameValueStringPairModel {
	if o == nil || IsNil(o.AdditionalData) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetAdditionalDataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given []NameValueStringPairModel and assigns it to the AdditionalData field.
func (o *HypervisorTemplateResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel) {
	o.AdditionalData = v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetCpuCount() int32 {
	if o == nil || IsNil(o.CpuCount) {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetCpuCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuCount) {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasCpuCount() bool {
	if o != nil && !IsNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *HypervisorTemplateResourceResponseModel) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetMemoryMB() int32 {
	if o == nil || IsNil(o.MemoryMB) {
		var ret int32
		return ret
	}
	return *o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetMemoryMBOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryMB) {
		return nil, false
	}
	return o.MemoryMB, true
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasMemoryMB() bool {
	if o != nil && !IsNil(o.MemoryMB) {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given int32 and assigns it to the MemoryMB field.
func (o *HypervisorTemplateResourceResponseModel) SetMemoryMB(v int32) {
	o.MemoryMB = &v
}

// GetHardDiskSizeGB returns the HardDiskSizeGB field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetHardDiskSizeGB() int32 {
	if o == nil || IsNil(o.HardDiskSizeGB) {
		var ret int32
		return ret
	}
	return *o.HardDiskSizeGB
}

// GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool) {
	if o == nil || IsNil(o.HardDiskSizeGB) {
		return nil, false
	}
	return o.HardDiskSizeGB, true
}

// HasHardDiskSizeGB returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasHardDiskSizeGB() bool {
	if o != nil && !IsNil(o.HardDiskSizeGB) {
		return true
	}

	return false
}

// SetHardDiskSizeGB gets a reference to the given int32 and assigns it to the HardDiskSizeGB field.
func (o *HypervisorTemplateResourceResponseModel) SetHardDiskSizeGB(v int32) {
	o.HardDiskSizeGB = &v
}

// GetMinMemoryMB returns the MinMemoryMB field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetMinMemoryMB() int32 {
	if o == nil || IsNil(o.MinMemoryMB) {
		var ret int32
		return ret
	}
	return *o.MinMemoryMB
}

// GetMinMemoryMBOk returns a tuple with the MinMemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetMinMemoryMBOk() (*int32, bool) {
	if o == nil || IsNil(o.MinMemoryMB) {
		return nil, false
	}
	return o.MinMemoryMB, true
}

// HasMinMemoryMB returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasMinMemoryMB() bool {
	if o != nil && !IsNil(o.MinMemoryMB) {
		return true
	}

	return false
}

// SetMinMemoryMB gets a reference to the given int32 and assigns it to the MinMemoryMB field.
func (o *HypervisorTemplateResourceResponseModel) SetMinMemoryMB(v int32) {
	o.MinMemoryMB = &v
}

// GetNetworkMappings returns the NetworkMappings field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel {
	if o == nil || IsNil(o.NetworkMappings) {
		var ret []NetworkMapResponseModel
		return ret
	}
	return o.NetworkMappings
}

// GetNetworkMappingsOk returns a tuple with the NetworkMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetNetworkMappingsOk() ([]NetworkMapResponseModel, bool) {
	if o == nil || IsNil(o.NetworkMappings) {
		return nil, false
	}
	return o.NetworkMappings, true
}

// HasNetworkMappings returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasNetworkMappings() bool {
	if o != nil && !IsNil(o.NetworkMappings) {
		return true
	}

	return false
}

// SetNetworkMappings gets a reference to the given []NetworkMapResponseModel and assigns it to the NetworkMappings field.
func (o *HypervisorTemplateResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel) {
	o.NetworkMappings = v
}

// GetAttachedDisks returns the AttachedDisks field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel {
	if o == nil || IsNil(o.AttachedDisks) {
		var ret []AttachedDiskResponseModel
		return ret
	}
	return o.AttachedDisks
}

// GetAttachedDisksOk returns a tuple with the AttachedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetAttachedDisksOk() ([]AttachedDiskResponseModel, bool) {
	if o == nil || IsNil(o.AttachedDisks) {
		return nil, false
	}
	return o.AttachedDisks, true
}

// HasAttachedDisks returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasAttachedDisks() bool {
	if o != nil && !IsNil(o.AttachedDisks) {
		return true
	}

	return false
}

// SetAttachedDisks gets a reference to the given []AttachedDiskResponseModel and assigns it to the AttachedDisks field.
func (o *HypervisorTemplateResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel) {
	o.AttachedDisks = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *HypervisorTemplateResourceResponseModel) SetOwner(v string) {
	o.Owner = &v
}

// GetIsWindowsTemplate returns the IsWindowsTemplate field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetIsWindowsTemplate() bool {
	if o == nil || IsNil(o.IsWindowsTemplate) {
		var ret bool
		return ret
	}
	return *o.IsWindowsTemplate
}

// GetIsWindowsTemplateOk returns a tuple with the IsWindowsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetIsWindowsTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWindowsTemplate) {
		return nil, false
	}
	return o.IsWindowsTemplate, true
}

// HasIsWindowsTemplate returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasIsWindowsTemplate() bool {
	if o != nil && !IsNil(o.IsWindowsTemplate) {
		return true
	}

	return false
}

// SetIsWindowsTemplate gets a reference to the given bool and assigns it to the IsWindowsTemplate field.
func (o *HypervisorTemplateResourceResponseModel) SetIsWindowsTemplate(v bool) {
	o.IsWindowsTemplate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HypervisorTemplateResourceResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HypervisorTemplateResourceResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HypervisorTemplateResourceResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetHasPersistentRootVolume returns the HasPersistentRootVolume field value
func (o *HypervisorTemplateResourceResponseModel) GetHasPersistentRootVolume() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPersistentRootVolume
}

// GetHasPersistentRootVolumeOk returns a tuple with the HasPersistentRootVolume field value
// and a boolean to check if the value has been set.
func (o *HypervisorTemplateResourceResponseModel) GetHasPersistentRootVolumeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPersistentRootVolume, true
}

// SetHasPersistentRootVolume sets field value
func (o *HypervisorTemplateResourceResponseModel) SetHasPersistentRootVolume(v bool) {
	o.HasPersistentRootVolume = v
}

func (o HypervisorTemplateResourceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorTemplateResourceResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.XDPath) {
		toSerialize["XDPath"] = o.XDPath
	}
	if !IsNil(o.RelativePath) {
		toSerialize["RelativePath"] = o.RelativePath
	}
	if !IsNil(o.FullRelativePath) {
		toSerialize["FullRelativePath"] = o.FullRelativePath
	}
	toSerialize["ResourceType"] = o.ResourceType
	if !IsNil(o.ObjectTypeName) {
		toSerialize["ObjectTypeName"] = o.ObjectTypeName
	}
	if !IsNil(o.ResourceContainer) {
		toSerialize["ResourceContainer"] = o.ResourceContainer
	}
	if !IsNil(o.ResourcePoolXDPath) {
		toSerialize["ResourcePoolXDPath"] = o.ResourcePoolXDPath
	}
	toSerialize["FullName"] = o.FullName
	toSerialize["IsContainer"] = o.IsContainer
	if !IsNil(o.Children) {
		toSerialize["Children"] = o.Children
	}
	toSerialize["IsMachine"] = o.IsMachine
	toSerialize["IsSnapshotable"] = o.IsSnapshotable
	toSerialize["AllResourcesRelativePath"] = o.AllResourcesRelativePath
	toSerialize["ResourcePool"] = o.ResourcePool
	toSerialize["IsSymLink"] = o.IsSymLink
	if !IsNil(o.AdditionalData) {
		toSerialize["AdditionalData"] = o.AdditionalData
	}
	if !IsNil(o.CpuCount) {
		toSerialize["CpuCount"] = o.CpuCount
	}
	if !IsNil(o.MemoryMB) {
		toSerialize["MemoryMB"] = o.MemoryMB
	}
	if !IsNil(o.HardDiskSizeGB) {
		toSerialize["HardDiskSizeGB"] = o.HardDiskSizeGB
	}
	if !IsNil(o.MinMemoryMB) {
		toSerialize["MinMemoryMB"] = o.MinMemoryMB
	}
	if !IsNil(o.NetworkMappings) {
		toSerialize["NetworkMappings"] = o.NetworkMappings
	}
	if !IsNil(o.AttachedDisks) {
		toSerialize["AttachedDisks"] = o.AttachedDisks
	}
	if !IsNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	if !IsNil(o.IsWindowsTemplate) {
		toSerialize["IsWindowsTemplate"] = o.IsWindowsTemplate
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["HasPersistentRootVolume"] = o.HasPersistentRootVolume
	return toSerialize, nil
}

type NullableHypervisorTemplateResourceResponseModel struct {
	value *HypervisorTemplateResourceResponseModel
	isSet bool
}

func (v NullableHypervisorTemplateResourceResponseModel) Get() *HypervisorTemplateResourceResponseModel {
	return v.value
}

func (v *NullableHypervisorTemplateResourceResponseModel) Set(val *HypervisorTemplateResourceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorTemplateResourceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorTemplateResourceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorTemplateResourceResponseModel(val *HypervisorTemplateResourceResponseModel) *NullableHypervisorTemplateResourceResponseModel {
	return &NullableHypervisorTemplateResourceResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorTemplateResourceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorTemplateResourceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


