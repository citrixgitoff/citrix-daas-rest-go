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

// checks if the MachineBaseResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineBaseResponseModel{}

// MachineBaseResponseModel Default response field (Only return the fields specified there if supported in  API ): Id,MachineCatalog,Name. 
type MachineBaseResponseModel struct {
	// Id of machine. Used to be: DesktopUid (and wasn't globally unique) OR UUID, depending on context Needs to be globally unique Might be constructed from site ID + internal Uid?  or use uuid
	Id string `json:"Id"`
	MachineCatalog *RefResponseModel `json:"MachineCatalog,omitempty"`
	// DNS host name of the machine. Used to be: MachineName
	Name NullableString `json:"Name,omitempty"`
}

// NewMachineBaseResponseModel instantiates a new MachineBaseResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineBaseResponseModel(id string) *MachineBaseResponseModel {
	this := MachineBaseResponseModel{}
	this.Id = id
	return &this
}

// NewMachineBaseResponseModelWithDefaults instantiates a new MachineBaseResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineBaseResponseModelWithDefaults() *MachineBaseResponseModel {
	this := MachineBaseResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *MachineBaseResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MachineBaseResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MachineBaseResponseModel) SetId(v string) {
	o.Id = v
}

// GetMachineCatalog returns the MachineCatalog field value if set, zero value otherwise.
func (o *MachineBaseResponseModel) GetMachineCatalog() RefResponseModel {
	if o == nil || IsNil(o.MachineCatalog) {
		var ret RefResponseModel
		return ret
	}
	return *o.MachineCatalog
}

// GetMachineCatalogOk returns a tuple with the MachineCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineBaseResponseModel) GetMachineCatalogOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.MachineCatalog) {
		return nil, false
	}
	return o.MachineCatalog, true
}

// HasMachineCatalog returns a boolean if a field has been set.
func (o *MachineBaseResponseModel) HasMachineCatalog() bool {
	if o != nil && !IsNil(o.MachineCatalog) {
		return true
	}

	return false
}

// SetMachineCatalog gets a reference to the given RefResponseModel and assigns it to the MachineCatalog field.
func (o *MachineBaseResponseModel) SetMachineCatalog(v RefResponseModel) {
	o.MachineCatalog = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineBaseResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineBaseResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MachineBaseResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MachineBaseResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MachineBaseResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MachineBaseResponseModel) UnsetName() {
	o.Name.Unset()
}

func (o MachineBaseResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineBaseResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.MachineCatalog) {
		toSerialize["MachineCatalog"] = o.MachineCatalog
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableMachineBaseResponseModel struct {
	value *MachineBaseResponseModel
	isSet bool
}

func (v NullableMachineBaseResponseModel) Get() *MachineBaseResponseModel {
	return v.value
}

func (v *NullableMachineBaseResponseModel) Set(val *MachineBaseResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineBaseResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineBaseResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineBaseResponseModel(val *MachineBaseResponseModel) *NullableMachineBaseResponseModel {
	return &NullableMachineBaseResponseModel{value: val, isSet: true}
}

func (v NullableMachineBaseResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineBaseResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


