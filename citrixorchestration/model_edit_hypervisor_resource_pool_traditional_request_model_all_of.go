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

// checks if the EditHypervisorResourcePoolTraditionalRequestModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHypervisorResourcePoolTraditionalRequestModelAllOf{}

// EditHypervisorResourcePoolTraditionalRequestModelAllOf struct for EditHypervisorResourcePoolTraditionalRequestModelAllOf
type EditHypervisorResourcePoolTraditionalRequestModelAllOf struct {
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
}

// NewEditHypervisorResourcePoolTraditionalRequestModelAllOf instantiates a new EditHypervisorResourcePoolTraditionalRequestModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHypervisorResourcePoolTraditionalRequestModelAllOf() *EditHypervisorResourcePoolTraditionalRequestModelAllOf {
	this := EditHypervisorResourcePoolTraditionalRequestModelAllOf{}
	return &this
}

// NewEditHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults instantiates a new EditHypervisorResourcePoolTraditionalRequestModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults() *EditHypervisorResourcePoolTraditionalRequestModelAllOf {
	this := EditHypervisorResourcePoolTraditionalRequestModelAllOf{}
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.Storage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the Storage field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.Storage = v
}

// GetTemporaryStorage returns the TemporaryStorage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.TemporaryStorage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.TemporaryStorage
}

// GetTemporaryStorageOk returns a tuple with the TemporaryStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.TemporaryStorage) {
		return nil, false
	}
	return o.TemporaryStorage, true
}

// HasTemporaryStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasTemporaryStorage() bool {
	if o != nil && !IsNil(o.TemporaryStorage) {
		return true
	}

	return false
}

// SetTemporaryStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the TemporaryStorage field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.TemporaryStorage = v
}

// GetPersonalvDiskStorage returns the PersonalvDiskStorage field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		var ret []HypervisorResourcePoolStorageRequestModel
		return ret
	}
	return o.PersonalvDiskStorage
}

// GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorageOk() ([]HypervisorResourcePoolStorageRequestModel, bool) {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		return nil, false
	}
	return o.PersonalvDiskStorage, true
}

// HasPersonalvDiskStorage returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasPersonalvDiskStorage() bool {
	if o != nil && !IsNil(o.PersonalvDiskStorage) {
		return true
	}

	return false
}

// SetPersonalvDiskStorage gets a reference to the given []HypervisorResourcePoolStorageRequestModel and assigns it to the PersonalvDiskStorage field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel) {
	o.PersonalvDiskStorage = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties) {
		var ret string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given string and assigns it to the CustomProperties field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetCustomProperties(v string) {
	o.CustomProperties = &v
}

// GetUseLocalStorageCaching returns the UseLocalStorageCaching field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCaching() bool {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		var ret bool
		return ret
	}
	return *o.UseLocalStorageCaching
}

// GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLocalStorageCaching) {
		return nil, false
	}
	return o.UseLocalStorageCaching, true
}

// HasUseLocalStorageCaching returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasUseLocalStorageCaching() bool {
	if o != nil && !IsNil(o.UseLocalStorageCaching) {
		return true
	}

	return false
}

// SetUseLocalStorageCaching gets a reference to the given bool and assigns it to the UseLocalStorageCaching field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetUseLocalStorageCaching(v bool) {
	o.UseLocalStorageCaching = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworks() []string {
	if o == nil || IsNil(o.Networks) {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetNetworks(v []string) {
	o.Networks = v
}

func (o EditHypervisorResourcePoolTraditionalRequestModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHypervisorResourcePoolTraditionalRequestModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf struct {
	value *EditHypervisorResourcePoolTraditionalRequestModelAllOf
	isSet bool
}

func (v NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) Get() *EditHypervisorResourcePoolTraditionalRequestModelAllOf {
	return v.value
}

func (v *NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) Set(val *EditHypervisorResourcePoolTraditionalRequestModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHypervisorResourcePoolTraditionalRequestModelAllOf(val *EditHypervisorResourcePoolTraditionalRequestModelAllOf) *NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf {
	return &NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf{value: val, isSet: true}
}

func (v NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHypervisorResourcePoolTraditionalRequestModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


