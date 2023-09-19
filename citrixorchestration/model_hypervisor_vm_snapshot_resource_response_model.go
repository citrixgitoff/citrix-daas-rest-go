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

// checks if the HypervisorVmSnapshotResourceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorVmSnapshotResourceResponseModel{}

// HypervisorVmSnapshotResourceResponseModel struct for HypervisorVmSnapshotResourceResponseModel
type HypervisorVmSnapshotResourceResponseModel struct {
	// Id of the resource.
	Id NullableString `json:"Id,omitempty"`
	// Name of the resource.
	Name NullableString `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath NullableString `json:"XDPath,omitempty"`
	// Path to the object, relative to the resource pool from which it was queried. `{{vm name}}.vm/{{snapshot name}}.snapshot`
	RelativePath NullableString `json:"RelativePath,omitempty"`
	// Full path to the resource, including the hypervisor, relative to the root of the API. Example: `Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources/{{RelativePath}}`
	FullRelativePath NullableString `json:"FullRelativePath,omitempty"`
	// Type of resource.
	ResourceType string `json:"ResourceType"`
	// The type name of the hypervisor resource object.
	ObjectTypeName NullableString `json:"ObjectTypeName,omitempty"`
	ResourceContainer *HypervisorResourceRefResponseModel `json:"ResourceContainer,omitempty"`
	// Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}` This value 
	ResourcePoolXDPath NullableString `json:"ResourcePoolXDPath,omitempty"`
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
	ResourcePool HypervisorResourcePoolRefResponseModel `json:"ResourcePool"`
	// Indicates whether the object is a valid symbolic link.
	IsSymLink bool `json:"IsSymLink"`
	// Additional data about the object in the form of key-value pairs.
	AdditionalData []NameValueStringPairModel `json:"AdditionalData,omitempty"`
	// Number of CPUs, if known.
	CpuCount NullableInt32 `json:"CpuCount,omitempty"`
	// Memory in megabytes, if known.
	MemoryMB NullableInt32 `json:"MemoryMB,omitempty"`
	// Hard disk size in gigabytes, if known.
	HardDiskSizeGB NullableInt32 `json:"HardDiskSizeGB,omitempty"`
	// Minimum memory required to run this VM or snapshot, in megabytes, if known.
	MinMemoryMB NullableInt32 `json:"MinMemoryMB,omitempty"`
	// Network mappings associated with the VM, if known.
	NetworkMappings []NetworkMapResponseModel `json:"NetworkMappings,omitempty"`
	// Disks attached to the VM, if known.
	AttachedDisks []AttachedDiskResponseModel `json:"AttachedDisks,omitempty"`
}

// NewHypervisorVmSnapshotResourceResponseModel instantiates a new HypervisorVmSnapshotResourceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorVmSnapshotResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourcePoolRefResponseModel, isSymLink bool) *HypervisorVmSnapshotResourceResponseModel {
	this := HypervisorVmSnapshotResourceResponseModel{}
	this.ResourceType = resourceType
	this.FullName = fullName
	this.IsContainer = isContainer
	this.IsMachine = isMachine
	this.IsSnapshotable = isSnapshotable
	this.AllResourcesRelativePath = allResourcesRelativePath
	this.ResourcePool = resourcePool
	this.IsSymLink = isSymLink
	return &this
}

// NewHypervisorVmSnapshotResourceResponseModelWithDefaults instantiates a new HypervisorVmSnapshotResourceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorVmSnapshotResourceResponseModelWithDefaults() *HypervisorVmSnapshotResourceResponseModel {
	this := HypervisorVmSnapshotResourceResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetXDPath returns the XDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath.Get()) {
		var ret string
		return ret
	}
	return *o.XDPath.Get()
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XDPath.Get(), o.XDPath.IsSet()
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasXDPath() bool {
	if o != nil && o.XDPath.IsSet() {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given NullableString and assigns it to the XDPath field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetXDPath(v string) {
	o.XDPath.Set(&v)
}
// SetXDPathNil sets the value for XDPath to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetXDPathNil() {
	o.XDPath.Set(nil)
}

// UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetXDPath() {
	o.XDPath.Unset()
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetFullRelativePath returns the FullRelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetFullRelativePath() string {
	if o == nil || IsNil(o.FullRelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.FullRelativePath.Get()
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetFullRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullRelativePath.Get(), o.FullRelativePath.IsSet()
}

// HasFullRelativePath returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasFullRelativePath() bool {
	if o != nil && o.FullRelativePath.IsSet() {
		return true
	}

	return false
}

// SetFullRelativePath gets a reference to the given NullableString and assigns it to the FullRelativePath field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetFullRelativePath(v string) {
	o.FullRelativePath.Set(&v)
}
// SetFullRelativePathNil sets the value for FullRelativePath to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetFullRelativePathNil() {
	o.FullRelativePath.Set(nil)
}

// UnsetFullRelativePath ensures that no value is present for FullRelativePath, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetFullRelativePath() {
	o.FullRelativePath.Unset()
}

// GetResourceType returns the ResourceType field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetResourceType(v string) {
	o.ResourceType = v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetObjectTypeName() string {
	if o == nil || IsNil(o.ObjectTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectTypeName.Get()
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetObjectTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypeName.Get(), o.ObjectTypeName.IsSet()
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasObjectTypeName() bool {
	if o != nil && o.ObjectTypeName.IsSet() {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given NullableString and assigns it to the ObjectTypeName field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetObjectTypeName(v string) {
	o.ObjectTypeName.Set(&v)
}
// SetObjectTypeNameNil sets the value for ObjectTypeName to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetObjectTypeNameNil() {
	o.ObjectTypeName.Set(nil)
}

// UnsetObjectTypeName ensures that no value is present for ObjectTypeName, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetObjectTypeName() {
	o.ObjectTypeName.Unset()
}

// GetResourceContainer returns the ResourceContainer field value if set, zero value otherwise.
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModel {
	if o == nil || IsNil(o.ResourceContainer) {
		var ret HypervisorResourceRefResponseModel
		return ret
	}
	return *o.ResourceContainer
}

// GetResourceContainerOk returns a tuple with the ResourceContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil || IsNil(o.ResourceContainer) {
		return nil, false
	}
	return o.ResourceContainer, true
}

// HasResourceContainer returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasResourceContainer() bool {
	if o != nil && !IsNil(o.ResourceContainer) {
		return true
	}

	return false
}

// SetResourceContainer gets a reference to the given HypervisorResourceRefResponseModel and assigns it to the ResourceContainer field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModel) {
	o.ResourceContainer = &v
}

// GetResourcePoolXDPath returns the ResourcePoolXDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourcePoolXDPath() string {
	if o == nil || IsNil(o.ResourcePoolXDPath.Get()) {
		var ret string
		return ret
	}
	return *o.ResourcePoolXDPath.Get()
}

// GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourcePoolXDPath.Get(), o.ResourcePoolXDPath.IsSet()
}

// HasResourcePoolXDPath returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasResourcePoolXDPath() bool {
	if o != nil && o.ResourcePoolXDPath.IsSet() {
		return true
	}

	return false
}

// SetResourcePoolXDPath gets a reference to the given NullableString and assigns it to the ResourcePoolXDPath field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetResourcePoolXDPath(v string) {
	o.ResourcePoolXDPath.Set(&v)
}
// SetResourcePoolXDPathNil sets the value for ResourcePoolXDPath to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetResourcePoolXDPathNil() {
	o.ResourcePoolXDPath.Set(nil)
}

// UnsetResourcePoolXDPath ensures that no value is present for ResourcePoolXDPath, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetResourcePoolXDPath() {
	o.ResourcePoolXDPath.Unset()
}

// GetFullName returns the FullName field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetFullName(v string) {
	o.FullName = v
}

// GetIsContainer returns the IsContainer field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsContainer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsContainerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsContainer, true
}

// SetIsContainer sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetIsContainer(v bool) {
	o.IsContainer = v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetChildren() []HypervisorResourceResponseModel {
	if o == nil {
		var ret []HypervisorResourceResponseModel
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetChildrenOk() ([]HypervisorResourceResponseModel, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasChildren() bool {
	if o != nil && IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []HypervisorResourceResponseModel and assigns it to the Children field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel) {
	o.Children = v
}

// GetIsMachine returns the IsMachine field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsMachine() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMachine
}

// GetIsMachineOk returns a tuple with the IsMachine field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsMachineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMachine, true
}

// SetIsMachine sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetIsMachine(v bool) {
	o.IsMachine = v
}

// GetIsSnapshotable returns the IsSnapshotable field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsSnapshotable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSnapshotable
}

// GetIsSnapshotableOk returns a tuple with the IsSnapshotable field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsSnapshotableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSnapshotable, true
}

// SetIsSnapshotable sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetIsSnapshotable(v bool) {
	o.IsSnapshotable = v
}

// GetAllResourcesRelativePath returns the AllResourcesRelativePath field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetAllResourcesRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AllResourcesRelativePath
}

// GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllResourcesRelativePath, true
}

// SetAllResourcesRelativePath sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetAllResourcesRelativePath(v string) {
	o.AllResourcesRelativePath = v
}

// GetResourcePool returns the ResourcePool field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel {
	if o == nil {
		var ret HypervisorResourcePoolRefResponseModel
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel) {
	o.ResourcePool = v
}

// GetIsSymLink returns the IsSymLink field value
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsSymLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSymLink
}

// GetIsSymLinkOk returns a tuple with the IsSymLink field value
// and a boolean to check if the value has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) GetIsSymLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSymLink, true
}

// SetIsSymLink sets field value
func (o *HypervisorVmSnapshotResourceResponseModel) SetIsSymLink(v bool) {
	o.IsSymLink = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetAdditionalData() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetAdditionalDataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasAdditionalData() bool {
	if o != nil && IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given []NameValueStringPairModel and assigns it to the AdditionalData field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel) {
	o.AdditionalData = v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetCpuCount() int32 {
	if o == nil || IsNil(o.CpuCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CpuCount.Get()
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetCpuCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuCount.Get(), o.CpuCount.IsSet()
}

// HasCpuCount returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasCpuCount() bool {
	if o != nil && o.CpuCount.IsSet() {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given NullableInt32 and assigns it to the CpuCount field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetCpuCount(v int32) {
	o.CpuCount.Set(&v)
}
// SetCpuCountNil sets the value for CpuCount to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetCpuCountNil() {
	o.CpuCount.Set(nil)
}

// UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetCpuCount() {
	o.CpuCount.Unset()
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetMemoryMB() int32 {
	if o == nil || IsNil(o.MemoryMB.Get()) {
		var ret int32
		return ret
	}
	return *o.MemoryMB.Get()
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetMemoryMBOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryMB.Get(), o.MemoryMB.IsSet()
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasMemoryMB() bool {
	if o != nil && o.MemoryMB.IsSet() {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given NullableInt32 and assigns it to the MemoryMB field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetMemoryMB(v int32) {
	o.MemoryMB.Set(&v)
}
// SetMemoryMBNil sets the value for MemoryMB to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetMemoryMBNil() {
	o.MemoryMB.Set(nil)
}

// UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetMemoryMB() {
	o.MemoryMB.Unset()
}

// GetHardDiskSizeGB returns the HardDiskSizeGB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetHardDiskSizeGB() int32 {
	if o == nil || IsNil(o.HardDiskSizeGB.Get()) {
		var ret int32
		return ret
	}
	return *o.HardDiskSizeGB.Get()
}

// GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardDiskSizeGB.Get(), o.HardDiskSizeGB.IsSet()
}

// HasHardDiskSizeGB returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasHardDiskSizeGB() bool {
	if o != nil && o.HardDiskSizeGB.IsSet() {
		return true
	}

	return false
}

// SetHardDiskSizeGB gets a reference to the given NullableInt32 and assigns it to the HardDiskSizeGB field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetHardDiskSizeGB(v int32) {
	o.HardDiskSizeGB.Set(&v)
}
// SetHardDiskSizeGBNil sets the value for HardDiskSizeGB to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetHardDiskSizeGBNil() {
	o.HardDiskSizeGB.Set(nil)
}

// UnsetHardDiskSizeGB ensures that no value is present for HardDiskSizeGB, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetHardDiskSizeGB() {
	o.HardDiskSizeGB.Unset()
}

// GetMinMemoryMB returns the MinMemoryMB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetMinMemoryMB() int32 {
	if o == nil || IsNil(o.MinMemoryMB.Get()) {
		var ret int32
		return ret
	}
	return *o.MinMemoryMB.Get()
}

// GetMinMemoryMBOk returns a tuple with the MinMemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetMinMemoryMBOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinMemoryMB.Get(), o.MinMemoryMB.IsSet()
}

// HasMinMemoryMB returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasMinMemoryMB() bool {
	if o != nil && o.MinMemoryMB.IsSet() {
		return true
	}

	return false
}

// SetMinMemoryMB gets a reference to the given NullableInt32 and assigns it to the MinMemoryMB field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetMinMemoryMB(v int32) {
	o.MinMemoryMB.Set(&v)
}
// SetMinMemoryMBNil sets the value for MinMemoryMB to be an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) SetMinMemoryMBNil() {
	o.MinMemoryMB.Set(nil)
}

// UnsetMinMemoryMB ensures that no value is present for MinMemoryMB, not even an explicit nil
func (o *HypervisorVmSnapshotResourceResponseModel) UnsetMinMemoryMB() {
	o.MinMemoryMB.Unset()
}

// GetNetworkMappings returns the NetworkMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel {
	if o == nil {
		var ret []NetworkMapResponseModel
		return ret
	}
	return o.NetworkMappings
}

// GetNetworkMappingsOk returns a tuple with the NetworkMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetNetworkMappingsOk() ([]NetworkMapResponseModel, bool) {
	if o == nil || IsNil(o.NetworkMappings) {
		return nil, false
	}
	return o.NetworkMappings, true
}

// HasNetworkMappings returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasNetworkMappings() bool {
	if o != nil && IsNil(o.NetworkMappings) {
		return true
	}

	return false
}

// SetNetworkMappings gets a reference to the given []NetworkMapResponseModel and assigns it to the NetworkMappings field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel) {
	o.NetworkMappings = v
}

// GetAttachedDisks returns the AttachedDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorVmSnapshotResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel {
	if o == nil {
		var ret []AttachedDiskResponseModel
		return ret
	}
	return o.AttachedDisks
}

// GetAttachedDisksOk returns a tuple with the AttachedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorVmSnapshotResourceResponseModel) GetAttachedDisksOk() ([]AttachedDiskResponseModel, bool) {
	if o == nil || IsNil(o.AttachedDisks) {
		return nil, false
	}
	return o.AttachedDisks, true
}

// HasAttachedDisks returns a boolean if a field has been set.
func (o *HypervisorVmSnapshotResourceResponseModel) HasAttachedDisks() bool {
	if o != nil && IsNil(o.AttachedDisks) {
		return true
	}

	return false
}

// SetAttachedDisks gets a reference to the given []AttachedDiskResponseModel and assigns it to the AttachedDisks field.
func (o *HypervisorVmSnapshotResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel) {
	o.AttachedDisks = v
}

func (o HypervisorVmSnapshotResourceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorVmSnapshotResourceResponseModel) ToMap() (map[string]interface{}, error) {
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
	if o.RelativePath.IsSet() {
		toSerialize["RelativePath"] = o.RelativePath.Get()
	}
	if o.FullRelativePath.IsSet() {
		toSerialize["FullRelativePath"] = o.FullRelativePath.Get()
	}
	toSerialize["ResourceType"] = o.ResourceType
	if o.ObjectTypeName.IsSet() {
		toSerialize["ObjectTypeName"] = o.ObjectTypeName.Get()
	}
	if !IsNil(o.ResourceContainer) {
		toSerialize["ResourceContainer"] = o.ResourceContainer
	}
	if o.ResourcePoolXDPath.IsSet() {
		toSerialize["ResourcePoolXDPath"] = o.ResourcePoolXDPath.Get()
	}
	toSerialize["FullName"] = o.FullName
	toSerialize["IsContainer"] = o.IsContainer
	if o.Children != nil {
		toSerialize["Children"] = o.Children
	}
	toSerialize["IsMachine"] = o.IsMachine
	toSerialize["IsSnapshotable"] = o.IsSnapshotable
	toSerialize["AllResourcesRelativePath"] = o.AllResourcesRelativePath
	toSerialize["ResourcePool"] = o.ResourcePool
	toSerialize["IsSymLink"] = o.IsSymLink
	if o.AdditionalData != nil {
		toSerialize["AdditionalData"] = o.AdditionalData
	}
	if o.CpuCount.IsSet() {
		toSerialize["CpuCount"] = o.CpuCount.Get()
	}
	if o.MemoryMB.IsSet() {
		toSerialize["MemoryMB"] = o.MemoryMB.Get()
	}
	if o.HardDiskSizeGB.IsSet() {
		toSerialize["HardDiskSizeGB"] = o.HardDiskSizeGB.Get()
	}
	if o.MinMemoryMB.IsSet() {
		toSerialize["MinMemoryMB"] = o.MinMemoryMB.Get()
	}
	if o.NetworkMappings != nil {
		toSerialize["NetworkMappings"] = o.NetworkMappings
	}
	if o.AttachedDisks != nil {
		toSerialize["AttachedDisks"] = o.AttachedDisks
	}
	return toSerialize, nil
}

type NullableHypervisorVmSnapshotResourceResponseModel struct {
	value *HypervisorVmSnapshotResourceResponseModel
	isSet bool
}

func (v NullableHypervisorVmSnapshotResourceResponseModel) Get() *HypervisorVmSnapshotResourceResponseModel {
	return v.value
}

func (v *NullableHypervisorVmSnapshotResourceResponseModel) Set(val *HypervisorVmSnapshotResourceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorVmSnapshotResourceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorVmSnapshotResourceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorVmSnapshotResourceResponseModel(val *HypervisorVmSnapshotResourceResponseModel) *NullableHypervisorVmSnapshotResourceResponseModel {
	return &NullableHypervisorVmSnapshotResourceResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorVmSnapshotResourceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorVmSnapshotResourceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


