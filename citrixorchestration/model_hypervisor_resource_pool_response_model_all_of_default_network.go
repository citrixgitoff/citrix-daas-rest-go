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

// checks if the HypervisorResourcePoolResponseModelAllOfDefaultNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolResponseModelAllOfDefaultNetwork{}

// HypervisorResourcePoolResponseModelAllOfDefaultNetwork Default network to use for VMs provisioned using the resource pool.
type HypervisorResourcePoolResponseModelAllOfDefaultNetwork struct {
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
	ResourceType *string `json:"ResourceType,omitempty"`
	// The type name of the hypervisor resource object.
	ObjectTypeName *string `json:"ObjectTypeName,omitempty"`
	ResourceContainer *HypervisorResourceRefResponseModelAllOfResourceContainer `json:"ResourceContainer,omitempty"`
	// Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}` This value 
	ResourcePoolXDPath *string `json:"ResourcePoolXDPath,omitempty"`
}

// NewHypervisorResourcePoolResponseModelAllOfDefaultNetwork instantiates a new HypervisorResourcePoolResponseModelAllOfDefaultNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolResponseModelAllOfDefaultNetwork() *HypervisorResourcePoolResponseModelAllOfDefaultNetwork {
	this := HypervisorResourcePoolResponseModelAllOfDefaultNetwork{}
	return &this
}

// NewHypervisorResourcePoolResponseModelAllOfDefaultNetworkWithDefaults instantiates a new HypervisorResourcePoolResponseModelAllOfDefaultNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolResponseModelAllOfDefaultNetworkWithDefaults() *HypervisorResourcePoolResponseModelAllOfDefaultNetwork {
	this := HypervisorResourcePoolResponseModelAllOfDefaultNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetName(v string) {
	o.Name = &v
}

// GetXDPath returns the XDPath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetXDPath() string {
	if o == nil || IsNil(o.XDPath) {
		var ret string
		return ret
	}
	return *o.XDPath
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.XDPath) {
		return nil, false
	}
	return o.XDPath, true
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasXDPath() bool {
	if o != nil && !IsNil(o.XDPath) {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given string and assigns it to the XDPath field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetXDPath(v string) {
	o.XDPath = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetRelativePath(v string) {
	o.RelativePath = &v
}

// GetFullRelativePath returns the FullRelativePath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetFullRelativePath() string {
	if o == nil || IsNil(o.FullRelativePath) {
		var ret string
		return ret
	}
	return *o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetFullRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.FullRelativePath) {
		return nil, false
	}
	return o.FullRelativePath, true
}

// HasFullRelativePath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasFullRelativePath() bool {
	if o != nil && !IsNil(o.FullRelativePath) {
		return true
	}

	return false
}

// SetFullRelativePath gets a reference to the given string and assigns it to the FullRelativePath field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetFullRelativePath(v string) {
	o.FullRelativePath = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetObjectTypeName() string {
	if o == nil || IsNil(o.ObjectTypeName) {
		var ret string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetObjectTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectTypeName) {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasObjectTypeName() bool {
	if o != nil && !IsNil(o.ObjectTypeName) {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given string and assigns it to the ObjectTypeName field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetObjectTypeName(v string) {
	o.ObjectTypeName = &v
}

// GetResourceContainer returns the ResourceContainer field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer {
	if o == nil || IsNil(o.ResourceContainer) {
		var ret HypervisorResourceRefResponseModelAllOfResourceContainer
		return ret
	}
	return *o.ResourceContainer
}

// GetResourceContainerOk returns a tuple with the ResourceContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool) {
	if o == nil || IsNil(o.ResourceContainer) {
		return nil, false
	}
	return o.ResourceContainer, true
}

// HasResourceContainer returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasResourceContainer() bool {
	if o != nil && !IsNil(o.ResourceContainer) {
		return true
	}

	return false
}

// SetResourceContainer gets a reference to the given HypervisorResourceRefResponseModelAllOfResourceContainer and assigns it to the ResourceContainer field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer) {
	o.ResourceContainer = &v
}

// GetResourcePoolXDPath returns the ResourcePoolXDPath field value if set, zero value otherwise.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourcePoolXDPath() string {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		var ret string
		return ret
	}
	return *o.ResourcePoolXDPath
}

// GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) GetResourcePoolXDPathOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolXDPath) {
		return nil, false
	}
	return o.ResourcePoolXDPath, true
}

// HasResourcePoolXDPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) HasResourcePoolXDPath() bool {
	if o != nil && !IsNil(o.ResourcePoolXDPath) {
		return true
	}

	return false
}

// SetResourcePoolXDPath gets a reference to the given string and assigns it to the ResourcePoolXDPath field.
func (o *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) SetResourcePoolXDPath(v string) {
	o.ResourcePoolXDPath = &v
}

func (o HypervisorResourcePoolResponseModelAllOfDefaultNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolResponseModelAllOfDefaultNetwork) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResourceType) {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if !IsNil(o.ObjectTypeName) {
		toSerialize["ObjectTypeName"] = o.ObjectTypeName
	}
	if !IsNil(o.ResourceContainer) {
		toSerialize["ResourceContainer"] = o.ResourceContainer
	}
	if !IsNil(o.ResourcePoolXDPath) {
		toSerialize["ResourcePoolXDPath"] = o.ResourcePoolXDPath
	}
	return toSerialize, nil
}

type NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork struct {
	value *HypervisorResourcePoolResponseModelAllOfDefaultNetwork
	isSet bool
}

func (v NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) Get() *HypervisorResourcePoolResponseModelAllOfDefaultNetwork {
	return v.value
}

func (v *NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) Set(val *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork(val *HypervisorResourcePoolResponseModelAllOfDefaultNetwork) *NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork {
	return &NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolResponseModelAllOfDefaultNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


