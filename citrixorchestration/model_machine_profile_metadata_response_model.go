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

// checks if the MachineProfileMetadataResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineProfileMetadataResponseModel{}

// MachineProfileMetadataResponseModel The opaque VMMetadata response model
type MachineProfileMetadataResponseModel struct {
	// VmMetadata machine size
	MachineSize NullableString `json:"MachineSize,omitempty"`
	// VmMetadata OS disk cache
	OsDiskCaching NullableString `json:"OsDiskCaching,omitempty"`
	// VmMetadata tags
	Tags NullableString `json:"Tags,omitempty"`
	// VmMetadata boot diagnostics configuration
	BootDiagnostics NullableBool `json:"BootDiagnostics,omitempty"`
	// VmMetadata accelerated network configuration
	AcceleratedNetwork NullableBool `json:"AcceleratedNetwork,omitempty"`
	// VmMetadata hibernation configuration.
	SupportsHibernation NullableBool `json:"SupportsHibernation,omitempty"`
	// VmMetadata zone name
	ZoneName NullableString `json:"ZoneName,omitempty"`
	// VmMetadata labels
	Labels NullableString `json:"Labels,omitempty"`
}

// NewMachineProfileMetadataResponseModel instantiates a new MachineProfileMetadataResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineProfileMetadataResponseModel() *MachineProfileMetadataResponseModel {
	this := MachineProfileMetadataResponseModel{}
	return &this
}

// NewMachineProfileMetadataResponseModelWithDefaults instantiates a new MachineProfileMetadataResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineProfileMetadataResponseModelWithDefaults() *MachineProfileMetadataResponseModel {
	this := MachineProfileMetadataResponseModel{}
	return &this
}

// GetMachineSize returns the MachineSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetMachineSize() string {
	if o == nil || IsNil(o.MachineSize.Get()) {
		var ret string
		return ret
	}
	return *o.MachineSize.Get()
}

// GetMachineSizeOk returns a tuple with the MachineSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetMachineSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineSize.Get(), o.MachineSize.IsSet()
}

// HasMachineSize returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasMachineSize() bool {
	if o != nil && o.MachineSize.IsSet() {
		return true
	}

	return false
}

// SetMachineSize gets a reference to the given NullableString and assigns it to the MachineSize field.
func (o *MachineProfileMetadataResponseModel) SetMachineSize(v string) {
	o.MachineSize.Set(&v)
}
// SetMachineSizeNil sets the value for MachineSize to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetMachineSizeNil() {
	o.MachineSize.Set(nil)
}

// UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetMachineSize() {
	o.MachineSize.Unset()
}

// GetOsDiskCaching returns the OsDiskCaching field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetOsDiskCaching() string {
	if o == nil || IsNil(o.OsDiskCaching.Get()) {
		var ret string
		return ret
	}
	return *o.OsDiskCaching.Get()
}

// GetOsDiskCachingOk returns a tuple with the OsDiskCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetOsDiskCachingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsDiskCaching.Get(), o.OsDiskCaching.IsSet()
}

// HasOsDiskCaching returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasOsDiskCaching() bool {
	if o != nil && o.OsDiskCaching.IsSet() {
		return true
	}

	return false
}

// SetOsDiskCaching gets a reference to the given NullableString and assigns it to the OsDiskCaching field.
func (o *MachineProfileMetadataResponseModel) SetOsDiskCaching(v string) {
	o.OsDiskCaching.Set(&v)
}
// SetOsDiskCachingNil sets the value for OsDiskCaching to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetOsDiskCachingNil() {
	o.OsDiskCaching.Set(nil)
}

// UnsetOsDiskCaching ensures that no value is present for OsDiskCaching, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetOsDiskCaching() {
	o.OsDiskCaching.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *MachineProfileMetadataResponseModel) SetTags(v string) {
	o.Tags.Set(&v)
}
// SetTagsNil sets the value for Tags to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetTags() {
	o.Tags.Unset()
}

// GetBootDiagnostics returns the BootDiagnostics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetBootDiagnostics() bool {
	if o == nil || IsNil(o.BootDiagnostics.Get()) {
		var ret bool
		return ret
	}
	return *o.BootDiagnostics.Get()
}

// GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetBootDiagnosticsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootDiagnostics.Get(), o.BootDiagnostics.IsSet()
}

// HasBootDiagnostics returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasBootDiagnostics() bool {
	if o != nil && o.BootDiagnostics.IsSet() {
		return true
	}

	return false
}

// SetBootDiagnostics gets a reference to the given NullableBool and assigns it to the BootDiagnostics field.
func (o *MachineProfileMetadataResponseModel) SetBootDiagnostics(v bool) {
	o.BootDiagnostics.Set(&v)
}
// SetBootDiagnosticsNil sets the value for BootDiagnostics to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetBootDiagnosticsNil() {
	o.BootDiagnostics.Set(nil)
}

// UnsetBootDiagnostics ensures that no value is present for BootDiagnostics, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetBootDiagnostics() {
	o.BootDiagnostics.Unset()
}

// GetAcceleratedNetwork returns the AcceleratedNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetAcceleratedNetwork() bool {
	if o == nil || IsNil(o.AcceleratedNetwork.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceleratedNetwork.Get()
}

// GetAcceleratedNetworkOk returns a tuple with the AcceleratedNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetAcceleratedNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceleratedNetwork.Get(), o.AcceleratedNetwork.IsSet()
}

// HasAcceleratedNetwork returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasAcceleratedNetwork() bool {
	if o != nil && o.AcceleratedNetwork.IsSet() {
		return true
	}

	return false
}

// SetAcceleratedNetwork gets a reference to the given NullableBool and assigns it to the AcceleratedNetwork field.
func (o *MachineProfileMetadataResponseModel) SetAcceleratedNetwork(v bool) {
	o.AcceleratedNetwork.Set(&v)
}
// SetAcceleratedNetworkNil sets the value for AcceleratedNetwork to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetAcceleratedNetworkNil() {
	o.AcceleratedNetwork.Set(nil)
}

// UnsetAcceleratedNetwork ensures that no value is present for AcceleratedNetwork, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetAcceleratedNetwork() {
	o.AcceleratedNetwork.Unset()
}

// GetSupportsHibernation returns the SupportsHibernation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetSupportsHibernation() bool {
	if o == nil || IsNil(o.SupportsHibernation.Get()) {
		var ret bool
		return ret
	}
	return *o.SupportsHibernation.Get()
}

// GetSupportsHibernationOk returns a tuple with the SupportsHibernation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetSupportsHibernationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportsHibernation.Get(), o.SupportsHibernation.IsSet()
}

// HasSupportsHibernation returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasSupportsHibernation() bool {
	if o != nil && o.SupportsHibernation.IsSet() {
		return true
	}

	return false
}

// SetSupportsHibernation gets a reference to the given NullableBool and assigns it to the SupportsHibernation field.
func (o *MachineProfileMetadataResponseModel) SetSupportsHibernation(v bool) {
	o.SupportsHibernation.Set(&v)
}
// SetSupportsHibernationNil sets the value for SupportsHibernation to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetSupportsHibernationNil() {
	o.SupportsHibernation.Set(nil)
}

// UnsetSupportsHibernation ensures that no value is present for SupportsHibernation, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetSupportsHibernation() {
	o.SupportsHibernation.Unset()
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName.Get()) {
		var ret string
		return ret
	}
	return *o.ZoneName.Get()
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetZoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneName.Get(), o.ZoneName.IsSet()
}

// HasZoneName returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasZoneName() bool {
	if o != nil && o.ZoneName.IsSet() {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given NullableString and assigns it to the ZoneName field.
func (o *MachineProfileMetadataResponseModel) SetZoneName(v string) {
	o.ZoneName.Set(&v)
}
// SetZoneNameNil sets the value for ZoneName to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetZoneNameNil() {
	o.ZoneName.Set(nil)
}

// UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetZoneName() {
	o.ZoneName.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineProfileMetadataResponseModel) GetLabels() string {
	if o == nil || IsNil(o.Labels.Get()) {
		var ret string
		return ret
	}
	return *o.Labels.Get()
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineProfileMetadataResponseModel) GetLabelsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels.Get(), o.Labels.IsSet()
}

// HasLabels returns a boolean if a field has been set.
func (o *MachineProfileMetadataResponseModel) HasLabels() bool {
	if o != nil && o.Labels.IsSet() {
		return true
	}

	return false
}

// SetLabels gets a reference to the given NullableString and assigns it to the Labels field.
func (o *MachineProfileMetadataResponseModel) SetLabels(v string) {
	o.Labels.Set(&v)
}
// SetLabelsNil sets the value for Labels to be an explicit nil
func (o *MachineProfileMetadataResponseModel) SetLabelsNil() {
	o.Labels.Set(nil)
}

// UnsetLabels ensures that no value is present for Labels, not even an explicit nil
func (o *MachineProfileMetadataResponseModel) UnsetLabels() {
	o.Labels.Unset()
}

func (o MachineProfileMetadataResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineProfileMetadataResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MachineSize.IsSet() {
		toSerialize["MachineSize"] = o.MachineSize.Get()
	}
	if o.OsDiskCaching.IsSet() {
		toSerialize["OsDiskCaching"] = o.OsDiskCaching.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
	}
	if o.BootDiagnostics.IsSet() {
		toSerialize["BootDiagnostics"] = o.BootDiagnostics.Get()
	}
	if o.AcceleratedNetwork.IsSet() {
		toSerialize["AcceleratedNetwork"] = o.AcceleratedNetwork.Get()
	}
	if o.SupportsHibernation.IsSet() {
		toSerialize["SupportsHibernation"] = o.SupportsHibernation.Get()
	}
	if o.ZoneName.IsSet() {
		toSerialize["ZoneName"] = o.ZoneName.Get()
	}
	if o.Labels.IsSet() {
		toSerialize["Labels"] = o.Labels.Get()
	}
	return toSerialize, nil
}

type NullableMachineProfileMetadataResponseModel struct {
	value *MachineProfileMetadataResponseModel
	isSet bool
}

func (v NullableMachineProfileMetadataResponseModel) Get() *MachineProfileMetadataResponseModel {
	return v.value
}

func (v *NullableMachineProfileMetadataResponseModel) Set(val *MachineProfileMetadataResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineProfileMetadataResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineProfileMetadataResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineProfileMetadataResponseModel(val *MachineProfileMetadataResponseModel) *NullableMachineProfileMetadataResponseModel {
	return &NullableMachineProfileMetadataResponseModel{value: val, isSet: true}
}

func (v NullableMachineProfileMetadataResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineProfileMetadataResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


