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

// checks if the EditHypervisorResourcePoolAWSOrGcpRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHypervisorResourcePoolAWSOrGcpRequestModel{}

// EditHypervisorResourcePoolAWSOrGcpRequestModel struct for EditHypervisorResourcePoolAWSOrGcpRequestModel
type EditHypervisorResourcePoolAWSOrGcpRequestModel struct {
	// Name of the resource pool.  Optional.  If not specified, will not be changed.
	Name NullableString `json:"Name,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// Indicates whether VMs created by Virtual Apps & Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default.
	VmTagging NullableBool `json:"VmTagging,omitempty"`
	// Metadata for hypervisor resource pool. When set the property value equal to null/empty means to remove this property.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	// Path to the storage resource(s) that are available for provisioning operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	Storage []HypervisorResourcePoolStorageRequestModel `json:"Storage,omitempty"`
	// Path to the storage resource(s) that are used for temporary operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	TemporaryStorage []HypervisorResourcePoolStorageRequestModel `json:"TemporaryStorage,omitempty"`
	// Path to the personal virtual disk storage resource(s). Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced.
	PersonalvDiskStorage []HypervisorResourcePoolStorageRequestModel `json:"PersonalvDiskStorage,omitempty"`
	// Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom.
	CustomProperties NullableString `json:"CustomProperties,omitempty"`
	// Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to `false`.
	UseLocalStorageCaching NullableBool `json:"UseLocalStorageCaching,omitempty"`
	// Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Networks []string `json:"Networks,omitempty"`
	// Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Subnets []string `json:"Subnets,omitempty"`
}

// NewEditHypervisorResourcePoolAWSOrGcpRequestModel instantiates a new EditHypervisorResourcePoolAWSOrGcpRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHypervisorResourcePoolAWSOrGcpRequestModel(connectionType HypervisorConnectionType) *EditHypervisorResourcePoolAWSOrGcpRequestModel {
	this := EditHypervisorResourcePoolAWSOrGcpRequestModel{}
	this.ConnectionType = connectionType
	return &this
}

// NewEditHypervisorResourcePoolAWSOrGcpRequestModelWithDefaults instantiates a new EditHypervisorResourcePoolAWSOrGcpRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHypervisorResourcePoolAWSOrGcpRequestModelWithDefaults() *EditHypervisorResourcePoolAWSOrGcpRequestModel {
	this := EditHypervisorResourcePoolAWSOrGcpRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetConnectionType returns the ConnectionType field value
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetVmTagging returns the VmTagging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetVmTagging() bool {
	if o == nil || IsNil(o.VmTagging.Get()) {
		var ret bool
		return ret
	}
	return *o.VmTagging.Get()
}

// GetVmTaggingOk returns a tuple with the VmTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetVmTaggingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTagging.Get(), o.VmTagging.IsSet()
}

// HasVmTagging returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasVmTagging() bool {
	if o != nil && o.VmTagging.IsSet() {
		return true
	}

	return false
}

// SetVmTagging gets a reference to the given NullableBool and assigns it to the VmTagging field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetVmTagging(v bool) {
	o.VmTagging.Set(&v)
}
// SetVmTaggingNil sets the value for VmTagging to be an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetVmTaggingNil() {
	o.VmTagging.Set(nil)
}

// UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) UnsetVmTagging() {
	o.VmTagging.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetMetadata() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasStorage() bool {
	if o != nil && IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the Storage field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.Storage = v
}

// GetTemporaryStorage returns the TemporaryStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.TemporaryStorage
}

// GetTemporaryStorageOk returns a tuple with the TemporaryStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetTemporaryStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.TemporaryStorage) {
		return nil, false
	}
	return o.TemporaryStorage, true
}

// HasTemporaryStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasTemporaryStorage() bool {
	if o != nil && IsNil(o.TemporaryStorage) {
		return true
	}

	return false
}

// SetTemporaryStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the TemporaryStorage field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.TemporaryStorage = v
}

// GetPersonalvDiskStorage returns the PersonalvDiskStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.PersonalvDiskStorage
}

// GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetPersonalvDiskStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		return nil, false
	}
	return o.PersonalvDiskStorage, true
}

// HasPersonalvDiskStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasPersonalvDiskStorage() bool {
	if o != nil && IsNil(o.PersonalvDiskStorage) {
		return true
	}

	return false
}

// SetPersonalvDiskStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the PersonalvDiskStorage field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.PersonalvDiskStorage = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties.Get()) {
		var ret string
		return ret
	}
	return *o.CustomProperties.Get()
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomProperties.Get(), o.CustomProperties.IsSet()
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasCustomProperties() bool {
	if o != nil && o.CustomProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given NullableString and assigns it to the CustomProperties field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetCustomProperties(v string) {
	o.CustomProperties.Set(&v)
}
// SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetCustomPropertiesNil() {
	o.CustomProperties.Set(nil)
}

// UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) UnsetCustomProperties() {
	o.CustomProperties.Unset()
}

// GetUseLocalStorageCaching returns the UseLocalStorageCaching field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetUseLocalStorageCaching() bool {
	if o == nil || IsNil(o.UseLocalStorageCaching.Get()) {
		var ret bool
		return ret
	}
	return *o.UseLocalStorageCaching.Get()
}

// GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetUseLocalStorageCachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseLocalStorageCaching.Get(), o.UseLocalStorageCaching.IsSet()
}

// HasUseLocalStorageCaching returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasUseLocalStorageCaching() bool {
	if o != nil && o.UseLocalStorageCaching.IsSet() {
		return true
	}

	return false
}

// SetUseLocalStorageCaching gets a reference to the given NullableBool and assigns it to the UseLocalStorageCaching field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetUseLocalStorageCaching(v bool) {
	o.UseLocalStorageCaching.Set(&v)
}
// SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetUseLocalStorageCachingNil() {
	o.UseLocalStorageCaching.Set(nil)
}

// UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) UnsetUseLocalStorageCaching() {
	o.UseLocalStorageCaching.Unset()
}

// GetNetworks returns the Networks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetNetworks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasNetworks() bool {
	if o != nil && IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetNetworks(v []string) {
	o.Networks = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) HasSubnets() bool {
	if o != nil && IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *EditHypervisorResourcePoolAWSOrGcpRequestModel) SetSubnets(v []string) {
	o.Subnets = v
}

func (o EditHypervisorResourcePoolAWSOrGcpRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHypervisorResourcePoolAWSOrGcpRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	if o.VmTagging.IsSet() {
		toSerialize["VmTagging"] = o.VmTagging.Get()
	}
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if o.Storage != nil {
		toSerialize["Storage"] = o.Storage
	}
	if o.TemporaryStorage != nil {
		toSerialize["TemporaryStorage"] = o.TemporaryStorage
	}
	if o.PersonalvDiskStorage != nil {
		toSerialize["PersonalvDiskStorage"] = o.PersonalvDiskStorage
	}
	if o.CustomProperties.IsSet() {
		toSerialize["CustomProperties"] = o.CustomProperties.Get()
	}
	if o.UseLocalStorageCaching.IsSet() {
		toSerialize["UseLocalStorageCaching"] = o.UseLocalStorageCaching.Get()
	}
	if o.Networks != nil {
		toSerialize["Networks"] = o.Networks
	}
	if o.Subnets != nil {
		toSerialize["Subnets"] = o.Subnets
	}
	return toSerialize, nil
}

type NullableEditHypervisorResourcePoolAWSOrGcpRequestModel struct {
	value *EditHypervisorResourcePoolAWSOrGcpRequestModel
	isSet bool
}

func (v NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) Get() *EditHypervisorResourcePoolAWSOrGcpRequestModel {
	return v.value
}

func (v *NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) Set(val *EditHypervisorResourcePoolAWSOrGcpRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHypervisorResourcePoolAWSOrGcpRequestModel(val *EditHypervisorResourcePoolAWSOrGcpRequestModel) *NullableEditHypervisorResourcePoolAWSOrGcpRequestModel {
	return &NullableEditHypervisorResourcePoolAWSOrGcpRequestModel{value: val, isSet: true}
}

func (v NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHypervisorResourcePoolAWSOrGcpRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


