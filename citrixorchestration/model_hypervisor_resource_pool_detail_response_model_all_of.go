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

// checks if the HypervisorResourcePoolDetailResponseModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolDetailResponseModelAllOf{}

// HypervisorResourcePoolDetailResponseModelAllOf struct for HypervisorResourcePoolDetailResponseModelAllOf
type HypervisorResourcePoolDetailResponseModelAllOf struct {
	// GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types.
	GpuTypes []HypervisorResourceRefResponseModel `json:"GpuTypes,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// If the hypervisor resource pool use ExplicitStorage.
	UsesExplicitStorage *bool `json:"UsesExplicitStorage,omitempty"`
	// Metadata for hypervisor resource pool. 
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
}

// NewHypervisorResourcePoolDetailResponseModelAllOf instantiates a new HypervisorResourcePoolDetailResponseModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolDetailResponseModelAllOf(connectionType HypervisorConnectionType) *HypervisorResourcePoolDetailResponseModelAllOf {
	this := HypervisorResourcePoolDetailResponseModelAllOf{}
	this.ConnectionType = connectionType
	return &this
}

// NewHypervisorResourcePoolDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolDetailResponseModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolDetailResponseModelAllOf {
	this := HypervisorResourcePoolDetailResponseModelAllOf{}
	return &this
}

// GetGpuTypes returns the GpuTypes field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetGpuTypes() []HypervisorResourceRefResponseModel {
	if o == nil || IsNil(o.GpuTypes) {
		var ret []HypervisorResourceRefResponseModel
		return ret
	}
	return o.GpuTypes
}

// GetGpuTypesOk returns a tuple with the GpuTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetGpuTypesOk() ([]HypervisorResourceRefResponseModel, bool) {
	if o == nil || IsNil(o.GpuTypes) {
		return nil, false
	}
	return o.GpuTypes, true
}

// HasGpuTypes returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasGpuTypes() bool {
	if o != nil && !IsNil(o.GpuTypes) {
		return true
	}

	return false
}

// SetGpuTypes gets a reference to the given []HypervisorResourceRefResponseModel and assigns it to the GpuTypes field.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetGpuTypes(v []HypervisorResourceRefResponseModel) {
	o.GpuTypes = v
}

// GetConnectionType returns the ConnectionType field value
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetUsesExplicitStorage returns the UsesExplicitStorage field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetUsesExplicitStorage() bool {
	if o == nil || IsNil(o.UsesExplicitStorage) {
		var ret bool
		return ret
	}
	return *o.UsesExplicitStorage
}

// GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetUsesExplicitStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesExplicitStorage) {
		return nil, false
	}
	return o.UsesExplicitStorage, true
}

// HasUsesExplicitStorage returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasUsesExplicitStorage() bool {
	if o != nil && !IsNil(o.UsesExplicitStorage) {
		return true
	}

	return false
}

// SetUsesExplicitStorage gets a reference to the given bool and assigns it to the UsesExplicitStorage field.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetUsesExplicitStorage(v bool) {
	o.UsesExplicitStorage = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetMetadata() []NameValueStringPairModel {
	if o == nil || IsNil(o.Metadata) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

func (o HypervisorResourcePoolDetailResponseModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolDetailResponseModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GpuTypes) {
		toSerialize["GpuTypes"] = o.GpuTypes
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	if !IsNil(o.UsesExplicitStorage) {
		toSerialize["UsesExplicitStorage"] = o.UsesExplicitStorage
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableHypervisorResourcePoolDetailResponseModelAllOf struct {
	value *HypervisorResourcePoolDetailResponseModelAllOf
	isSet bool
}

func (v NullableHypervisorResourcePoolDetailResponseModelAllOf) Get() *HypervisorResourcePoolDetailResponseModelAllOf {
	return v.value
}

func (v *NullableHypervisorResourcePoolDetailResponseModelAllOf) Set(val *HypervisorResourcePoolDetailResponseModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolDetailResponseModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolDetailResponseModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolDetailResponseModelAllOf(val *HypervisorResourcePoolDetailResponseModelAllOf) *NullableHypervisorResourcePoolDetailResponseModelAllOf {
	return &NullableHypervisorResourcePoolDetailResponseModelAllOf{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolDetailResponseModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolDetailResponseModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


