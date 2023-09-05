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

// checks if the EditHypervisorResourcePoolAzureRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHypervisorResourcePoolAzureRequestModel{}

// EditHypervisorResourcePoolAzureRequestModel struct for EditHypervisorResourcePoolAzureRequestModel
type EditHypervisorResourcePoolAzureRequestModel struct {
	// Name of the resource pool.  Optional.  If not specified, will not be changed.
	Name *string `json:"Name,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// Indicates whether VMs created by Virtual Apps & Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default.
	VmTagging *bool `json:"VmTagging,omitempty"`
	// Metadata for hypervisor resource pool. When set the property value equal to null/empty means to remove this property.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	// Path to the storage resource(s) that are available for provisioning operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	Storage []HypervisorResourcePoolStorageRequestModel `json:"Storage,omitempty"`
	// Path to the storage resource(s) that are used for temporary operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	TemporaryStorage []HypervisorResourcePoolStorageRequestModel `json:"TemporaryStorage,omitempty"`
	// Path to the personal virtual disk storage resource(s). Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	PersonalvDiskStorage []HypervisorResourcePoolStorageRequestModel `json:"PersonalvDiskStorage,omitempty"`
	// Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom.
	CustomProperties *string `json:"CustomProperties,omitempty"`
	// Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to `false`.
	UseLocalStorageCaching *bool `json:"UseLocalStorageCaching,omitempty"`
	// Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Networks []string `json:"Networks,omitempty"`
	// Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Subnets []string `json:"Subnets,omitempty"`
}

// NewEditHypervisorResourcePoolAzureRequestModel instantiates a new EditHypervisorResourcePoolAzureRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHypervisorResourcePoolAzureRequestModel(connectionType HypervisorConnectionType) *EditHypervisorResourcePoolAzureRequestModel {
	this := EditHypervisorResourcePoolAzureRequestModel{}
	this.ConnectionType = connectionType
	return &this
}

// NewEditHypervisorResourcePoolAzureRequestModelWithDefaults instantiates a new EditHypervisorResourcePoolAzureRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHypervisorResourcePoolAzureRequestModelWithDefaults() *EditHypervisorResourcePoolAzureRequestModel {
	this := EditHypervisorResourcePoolAzureRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetName(v string) {
	o.Name = &v
}

// GetConnectionType returns the ConnectionType field value
func (o *EditHypervisorResourcePoolAzureRequestModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *EditHypervisorResourcePoolAzureRequestModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetVmTagging returns the VmTagging field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetVmTagging() bool {
	if o == nil || IsNil(o.VmTagging) {
		var ret bool
		return ret
	}
	return *o.VmTagging
}

// GetVmTaggingOk returns a tuple with the VmTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetVmTaggingOk() (*bool, bool) {
	if o == nil || IsNil(o.VmTagging) {
		return nil, false
	}
	return o.VmTagging, true
}

// HasVmTagging returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasVmTagging() bool {
	if o != nil && !IsNil(o.VmTagging) {
		return true
	}

	return false
}

// SetVmTagging gets a reference to the given bool and assigns it to the VmTagging field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetVmTagging(v bool) {
	o.VmTagging = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetMetadata() []NameValueStringPairModel {
	if o == nil || IsNil(o.Metadata) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.Storage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the Storage field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.Storage = v
}

// GetTemporaryStorage returns the TemporaryStorage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.TemporaryStorage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.TemporaryStorage
}

// GetTemporaryStorageOk returns a tuple with the TemporaryStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetTemporaryStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.TemporaryStorage) {
		return nil, false
	}
	return o.TemporaryStorage, true
}

// HasTemporaryStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasTemporaryStorage() bool {
	if o != nil && !IsNil(o.TemporaryStorage) {
		return true
	}

	return false
}

// SetTemporaryStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the TemporaryStorage field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.TemporaryStorage = v
}

// GetPersonalvDiskStorage returns the PersonalvDiskStorage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.PersonalvDiskStorage
}

// GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		return nil, false
	}
	return o.PersonalvDiskStorage, true
}

// HasPersonalvDiskStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasPersonalvDiskStorage() bool {
	if o != nil && !IsNil(o.PersonalvDiskStorage) {
		return true
	}

	return false
}

// SetPersonalvDiskStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the PersonalvDiskStorage field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.PersonalvDiskStorage = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties) {
		var ret string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given string and assigns it to the CustomProperties field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetCustomProperties(v string) {
	o.CustomProperties = &v
}

// GetUseLocalStorageCaching returns the UseLocalStorageCaching field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCaching() bool {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		var ret bool
		return ret
	}
	return *o.UseLocalStorageCaching
}

// GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		return nil, false
	}
	return o.UseLocalStorageCaching, true
}

// HasUseLocalStorageCaching returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasUseLocalStorageCaching() bool {
	if o != nil && !IsNil(o.UseLocalStorageCaching) {
		return true
	}

	return false
}

// SetUseLocalStorageCaching gets a reference to the given bool and assigns it to the UseLocalStorageCaching field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetUseLocalStorageCaching(v bool) {
	o.UseLocalStorageCaching = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetNetworks() []string {
	if o == nil || IsNil(o.Networks) {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetNetworks(v []string) {
	o.Networks = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetSubnets() []string {
	if o == nil || IsNil(o.Subnets) {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAzureRequestModel) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *EditHypervisorResourcePoolAzureRequestModel) SetSubnets(v []string) {
	o.Subnets = v
}

func (o EditHypervisorResourcePoolAzureRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHypervisorResourcePoolAzureRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	if !IsNil(o.VmTagging) {
		toSerialize["VmTagging"] = o.VmTagging
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Storage) {
		toSerialize["Storage"] = o.Storage
	}
	if !IsNil(o.TemporaryStorage) {
		toSerialize["TemporaryStorage"] = o.TemporaryStorage
	}
	if !IsNil(o.PersonalvDiskStorage) {
		toSerialize["PersonalvDiskStorage"] = o.PersonalvDiskStorage
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["CustomProperties"] = o.CustomProperties
	}
	if !IsNil(o.UseLocalStorageCaching) {
		toSerialize["UseLocalStorageCaching"] = o.UseLocalStorageCaching
	}
	if !IsNil(o.Networks) {
		toSerialize["Networks"] = o.Networks
	}
	if !IsNil(o.Subnets) {
		toSerialize["Subnets"] = o.Subnets
	}
	return toSerialize, nil
}

type NullableEditHypervisorResourcePoolAzureRequestModel struct {
	value *EditHypervisorResourcePoolAzureRequestModel
	isSet bool
}

func (v NullableEditHypervisorResourcePoolAzureRequestModel) Get() *EditHypervisorResourcePoolAzureRequestModel {
	return v.value
}

func (v *NullableEditHypervisorResourcePoolAzureRequestModel) Set(val *EditHypervisorResourcePoolAzureRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHypervisorResourcePoolAzureRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHypervisorResourcePoolAzureRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHypervisorResourcePoolAzureRequestModel(val *EditHypervisorResourcePoolAzureRequestModel) *NullableEditHypervisorResourcePoolAzureRequestModel {
	return &NullableEditHypervisorResourcePoolAzureRequestModel{value: val, isSet: true}
}

func (v NullableEditHypervisorResourcePoolAzureRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHypervisorResourcePoolAzureRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


