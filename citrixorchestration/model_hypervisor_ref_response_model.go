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

// checks if the HypervisorRefResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorRefResponseModel{}

// HypervisorRefResponseModel struct for HypervisorRefResponseModel
type HypervisorRefResponseModel struct {
	// Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility
	Id NullableString `json:"Id,omitempty"`
	// DEPRECATED. Use Id.
	Uid NullableInt32 `json:"Uid,omitempty"`
	// Name of the object.
	Name NullableString `json:"Name,omitempty"`
	ConnectionType *HypervisorConnectionType `json:"ConnectionType,omitempty"`
	// The class name for the Citrix Managed Machine library that is used to access the hypervisor.
	PluginFactoryName NullableString `json:"PluginFactoryName,omitempty"`
}

// NewHypervisorRefResponseModel instantiates a new HypervisorRefResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorRefResponseModel() *HypervisorRefResponseModel {
	this := HypervisorRefResponseModel{}
	return &this
}

// NewHypervisorRefResponseModelWithDefaults instantiates a new HypervisorRefResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorRefResponseModelWithDefaults() *HypervisorRefResponseModel {
	this := HypervisorRefResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorRefResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorRefResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorRefResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HypervisorRefResponseModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HypervisorRefResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HypervisorRefResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorRefResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorRefResponseModel) GetUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *HypervisorRefResponseModel) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *HypervisorRefResponseModel) SetUid(v int32) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *HypervisorRefResponseModel) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *HypervisorRefResponseModel) UnsetUid() {
	o.Uid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorRefResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorRefResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorRefResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorRefResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorRefResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorRefResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *HypervisorRefResponseModel) GetConnectionType() HypervisorConnectionType {
	if o == nil || IsNil(o.ConnectionType) {
		var ret HypervisorConnectionType
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorRefResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *HypervisorRefResponseModel) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given HypervisorConnectionType and assigns it to the ConnectionType field.
func (o *HypervisorRefResponseModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = &v
}

// GetPluginFactoryName returns the PluginFactoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorRefResponseModel) GetPluginFactoryName() string {
	if o == nil || IsNil(o.PluginFactoryName.Get()) {
		var ret string
		return ret
	}
	return *o.PluginFactoryName.Get()
}

// GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorRefResponseModel) GetPluginFactoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginFactoryName.Get(), o.PluginFactoryName.IsSet()
}

// HasPluginFactoryName returns a boolean if a field has been set.
func (o *HypervisorRefResponseModel) HasPluginFactoryName() bool {
	if o != nil && o.PluginFactoryName.IsSet() {
		return true
	}

	return false
}

// SetPluginFactoryName gets a reference to the given NullableString and assigns it to the PluginFactoryName field.
func (o *HypervisorRefResponseModel) SetPluginFactoryName(v string) {
	o.PluginFactoryName.Set(&v)
}
// SetPluginFactoryNameNil sets the value for PluginFactoryName to be an explicit nil
func (o *HypervisorRefResponseModel) SetPluginFactoryNameNil() {
	o.PluginFactoryName.Set(nil)
}

// UnsetPluginFactoryName ensures that no value is present for PluginFactoryName, not even an explicit nil
func (o *HypervisorRefResponseModel) UnsetPluginFactoryName() {
	o.PluginFactoryName.Unset()
}

func (o HypervisorRefResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorRefResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["Uid"] = o.Uid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.ConnectionType) {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.PluginFactoryName.IsSet() {
		toSerialize["PluginFactoryName"] = o.PluginFactoryName.Get()
	}
	return toSerialize, nil
}

type NullableHypervisorRefResponseModel struct {
	value *HypervisorRefResponseModel
	isSet bool
}

func (v NullableHypervisorRefResponseModel) Get() *HypervisorRefResponseModel {
	return v.value
}

func (v *NullableHypervisorRefResponseModel) Set(val *HypervisorRefResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorRefResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorRefResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorRefResponseModel(val *HypervisorRefResponseModel) *NullableHypervisorRefResponseModel {
	return &NullableHypervisorRefResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorRefResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorRefResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


