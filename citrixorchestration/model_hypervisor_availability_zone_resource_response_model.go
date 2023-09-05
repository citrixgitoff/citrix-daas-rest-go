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

// checks if the HypervisorAvailabilityZoneResourceResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorAvailabilityZoneResourceResponseModel{}

// HypervisorAvailabilityZoneResourceResponseModel struct for HypervisorAvailabilityZoneResourceResponseModel
type HypervisorAvailabilityZoneResourceResponseModel struct {
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
	// Indicates whether this zone supports the use of security groups for isolation.
	SupportsSecurityGroups bool `json:"SupportsSecurityGroups"`
}

// NewHypervisorAvailabilityZoneResourceResponseModel instantiates a new HypervisorAvailabilityZoneResourceResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorAvailabilityZoneResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourceResponseModelAllOfResourcePool, isSymLink bool, supportsSecurityGroups bool) *HypervisorAvailabilityZoneResourceResponseModel {
	this := HypervisorAvailabilityZoneResourceResponseModel{}
	this.ResourceType = resourceType
	this.FullName = fullName
	this.IsContainer = isContainer
	this.IsMachine = isMachine
	this.IsSnapshotable = isSnapshotable
	this.AllResourcesRelativePath = allResourcesRelativePath
	this.ResourcePool = resourcePool
	this.IsSymLink = isSymLink
	this.SupportsSecurityGroups = supportsSecurityGroups
	return &this
}

// NewHypervisorAvailabilityZoneResourceResponseModelWithDefaults instantiates a new HypervisorAvailabilityZoneResourceResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorAvailabilityZoneResourceResponseModelWithDefaults() *HypervisorAvailabilityZoneResourceResponseModel {
	this := HypervisorAvailabilityZoneResourceResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetName(v string) {
	o.Name = &v
}

// GetXDPath returns the XDPath field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath) {
		var ret string
		return ret
	}
	return *o.XDPath
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.XDPath) {
		return nil, false
	}
	return o.XDPath, true
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasXDPath() bool {
	if o != nil && !IsNil(o.XDPath) {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given string and assigns it to the XDPath field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetXDPath(v string) {
	o.XDPath = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetRelativePath(v string) {
	o.RelativePath = &v
}

// GetFullRelativePath returns the FullRelativePath field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetFullRelativePath() string {
	if o == nil || IsNil(o.FullRelativePath) {
		var ret string
		return ret
	}
	return *o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetFullRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.FullRelativePath) {
		return nil, false
	}
	return o.FullRelativePath, true
}

// HasFullRelativePath returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasFullRelativePath() bool {
	if o != nil && !IsNil(o.FullRelativePath) {
		return true
	}

	return false
}

// SetFullRelativePath gets a reference to the given string and assigns it to the FullRelativePath field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetFullRelativePath(v string) {
	o.FullRelativePath = &v
}

// GetResourceType returns the ResourceType field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetResourceType(v string) {
	o.ResourceType = v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetObjectTypeName() string {
	if o == nil || IsNil(o.ObjectTypeName) {
		var ret string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetObjectTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectTypeName) {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasObjectTypeName() bool {
	if o != nil && !IsNil(o.ObjectTypeName) {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given string and assigns it to the ObjectTypeName field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetObjectTypeName(v string) {
	o.ObjectTypeName = &v
}

// GetResourceContainer returns the ResourceContainer field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer {
	if o == nil || IsNil(o.ResourceContainer) {
		var ret HypervisorResourceRefResponseModelAllOfResourceContainer
		return ret
	}
	return *o.ResourceContainer
}

// GetResourceContainerOk returns a tuple with the ResourceContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool) {
	if o == nil || IsNil(o.ResourceContainer) {
		return nil, false
	}
	return o.ResourceContainer, true
}

// HasResourceContainer returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasResourceContainer() bool {
	if o != nil && !IsNil(o.ResourceContainer) {
		return true
	}

	return false
}

// SetResourceContainer gets a reference to the given HypervisorResourceRefResponseModelAllOfResourceContainer and assigns it to the ResourceContainer field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer) {
	o.ResourceContainer = &v
}

// GetResourcePoolXDPath returns the ResourcePoolXDPath field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourcePoolXDPath() string {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		var ret string
		return ret
	}
	return *o.ResourcePoolXDPath
}

// GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		return nil, false
	}
	return o.ResourcePoolXDPath, true
}

// HasResourcePoolXDPath returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasResourcePoolXDPath() bool {
	if o != nil && !IsNil(o.ResourcePoolXDPath) {
		return true
	}

	return false
}

// SetResourcePoolXDPath gets a reference to the given string and assigns it to the ResourcePoolXDPath field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetResourcePoolXDPath(v string) {
	o.ResourcePoolXDPath = &v
}

// GetFullName returns the FullName field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetFullName(v string) {
	o.FullName = v
}

// GetIsContainer returns the IsContainer field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsContainer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsContainerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsContainer, true
}

// SetIsContainer sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetIsContainer(v bool) {
	o.IsContainer = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetChildren() []HypervisorResourceResponseModel {
	if o == nil || IsNil(o.Children) {
		var ret []HypervisorResourceResponseModel
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetChildrenOk() ([]HypervisorResourceResponseModel, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []HypervisorResourceResponseModel and assigns it to the Children field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel) {
	o.Children = v
}

// GetIsMachine returns the IsMachine field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsMachine() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMachine
}

// GetIsMachineOk returns a tuple with the IsMachine field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsMachineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMachine, true
}

// SetIsMachine sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetIsMachine(v bool) {
	o.IsMachine = v
}

// GetIsSnapshotable returns the IsSnapshotable field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsSnapshotable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSnapshotable
}

// GetIsSnapshotableOk returns a tuple with the IsSnapshotable field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsSnapshotableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSnapshotable, true
}

// SetIsSnapshotable sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetIsSnapshotable(v bool) {
	o.IsSnapshotable = v
}

// GetAllResourcesRelativePath returns the AllResourcesRelativePath field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetAllResourcesRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AllResourcesRelativePath
}

// GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllResourcesRelativePath, true
}

// SetAllResourcesRelativePath sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetAllResourcesRelativePath(v string) {
	o.AllResourcesRelativePath = v
}

// GetResourcePool returns the ResourcePool field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourcePool() HypervisorResourceResponseModelAllOfResourcePool {
	if o == nil {
		var ret HypervisorResourceResponseModelAllOfResourcePool
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetResourcePoolOk() (*HypervisorResourceResponseModelAllOfResourcePool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetResourcePool(v HypervisorResourceResponseModelAllOfResourcePool) {
	o.ResourcePool = v
}

// GetIsSymLink returns the IsSymLink field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsSymLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSymLink
}

// GetIsSymLinkOk returns a tuple with the IsSymLink field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetIsSymLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSymLink, true
}

// SetIsSymLink sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetIsSymLink(v bool) {
	o.IsSymLink = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetAdditionalData() []NameValueStringPairModel {
	if o == nil || IsNil(o.AdditionalData) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetAdditionalDataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given []NameValueStringPairModel and assigns it to the AdditionalData field.
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel) {
	o.AdditionalData = v
}

// GetSupportsSecurityGroups returns the SupportsSecurityGroups field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetSupportsSecurityGroups() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsSecurityGroups
}

// GetSupportsSecurityGroupsOk returns a tuple with the SupportsSecurityGroups field value
// and a boolean to check if the value has been set.
func (o *HypervisorAvailabilityZoneResourceResponseModel) GetSupportsSecurityGroupsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportsSecurityGroups, true
}

// SetSupportsSecurityGroups sets field value
func (o *HypervisorAvailabilityZoneResourceResponseModel) SetSupportsSecurityGroups(v bool) {
	o.SupportsSecurityGroups = v
}

func (o HypervisorAvailabilityZoneResourceResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorAvailabilityZoneResourceResponseModel) ToMap() (map[string]interface{}, error) {
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
	toSerialize["SupportsSecurityGroups"] = o.SupportsSecurityGroups
	return toSerialize, nil
}

type NullableHypervisorAvailabilityZoneResourceResponseModel struct {
	value *HypervisorAvailabilityZoneResourceResponseModel
	isSet bool
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModel) Get() *HypervisorAvailabilityZoneResourceResponseModel {
	return v.value
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModel) Set(val *HypervisorAvailabilityZoneResourceResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorAvailabilityZoneResourceResponseModel(val *HypervisorAvailabilityZoneResourceResponseModel) *NullableHypervisorAvailabilityZoneResourceResponseModel {
	return &NullableHypervisorAvailabilityZoneResourceResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorAvailabilityZoneResourceResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorAvailabilityZoneResourceResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


