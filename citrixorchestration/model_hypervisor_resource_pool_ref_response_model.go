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

// checks if the HypervisorResourcePoolRefResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePoolRefResponseModel{}

// HypervisorResourcePoolRefResponseModel struct for HypervisorResourcePoolRefResponseModel
type HypervisorResourcePoolRefResponseModel struct {
	// Id of the resource.
	Id NullableString `json:"Id,omitempty"`
	// Name of the resource.
	Name NullableString `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath NullableString `json:"XDPath,omitempty"`
	// Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: `Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources`
	FullRelativePath string `json:"FullRelativePath"`
	Hypervisor RefResponseModel `json:"Hypervisor"`
}

// NewHypervisorResourcePoolRefResponseModel instantiates a new HypervisorResourcePoolRefResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePoolRefResponseModel(fullRelativePath string, hypervisor RefResponseModel) *HypervisorResourcePoolRefResponseModel {
	this := HypervisorResourcePoolRefResponseModel{}
	this.FullRelativePath = fullRelativePath
	this.Hypervisor = hypervisor
	return &this
}

// NewHypervisorResourcePoolRefResponseModelWithDefaults instantiates a new HypervisorResourcePoolRefResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePoolRefResponseModelWithDefaults() *HypervisorResourcePoolRefResponseModel {
	this := HypervisorResourcePoolRefResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePoolRefResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePoolRefResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorResourcePoolRefResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HypervisorResourcePoolRefResponseModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePoolRefResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePoolRefResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorResourcePoolRefResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorResourcePoolRefResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetXDPath returns the XDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePoolRefResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath.Get()) {
		var ret string
		return ret
	}
	return *o.XDPath.Get()
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePoolRefResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XDPath.Get(), o.XDPath.IsSet()
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorResourcePoolRefResponseModel) HasXDPath() bool {
	if o != nil && o.XDPath.IsSet() {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given NullableString and assigns it to the XDPath field.
func (o *HypervisorResourcePoolRefResponseModel) SetXDPath(v string) {
	o.XDPath.Set(&v)
}
// SetXDPathNil sets the value for XDPath to be an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) SetXDPathNil() {
	o.XDPath.Set(nil)
}

// UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
func (o *HypervisorResourcePoolRefResponseModel) UnsetXDPath() {
	o.XDPath.Unset()
}

// GetFullRelativePath returns the FullRelativePath field value
func (o *HypervisorResourcePoolRefResponseModel) GetFullRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolRefResponseModel) GetFullRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullRelativePath, true
}

// SetFullRelativePath sets field value
func (o *HypervisorResourcePoolRefResponseModel) SetFullRelativePath(v string) {
	o.FullRelativePath = v
}

// GetHypervisor returns the Hypervisor field value
func (o *HypervisorResourcePoolRefResponseModel) GetHypervisor() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value
// and a boolean to check if the value has been set.
func (o *HypervisorResourcePoolRefResponseModel) GetHypervisorOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hypervisor, true
}

// SetHypervisor sets field value
func (o *HypervisorResourcePoolRefResponseModel) SetHypervisor(v RefResponseModel) {
	o.Hypervisor = v
}

func (o HypervisorResourcePoolRefResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePoolRefResponseModel) ToMap() (map[string]interface{}, error) {
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
	toSerialize["FullRelativePath"] = o.FullRelativePath
	toSerialize["Hypervisor"] = o.Hypervisor
	return toSerialize, nil
}

type NullableHypervisorResourcePoolRefResponseModel struct {
	value *HypervisorResourcePoolRefResponseModel
	isSet bool
}

func (v NullableHypervisorResourcePoolRefResponseModel) Get() *HypervisorResourcePoolRefResponseModel {
	return v.value
}

func (v *NullableHypervisorResourcePoolRefResponseModel) Set(val *HypervisorResourcePoolRefResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePoolRefResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePoolRefResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePoolRefResponseModel(val *HypervisorResourcePoolRefResponseModel) *NullableHypervisorResourcePoolRefResponseModel {
	return &NullableHypervisorResourcePoolRefResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorResourcePoolRefResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePoolRefResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


