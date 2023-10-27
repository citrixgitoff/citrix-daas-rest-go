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

// checks if the CreateMachineConfigurationRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMachineConfigurationRequestModel{}

// CreateMachineConfigurationRequestModel Request object for creating a machine configuration.
type CreateMachineConfigurationRequestModel struct {
	// Id of the configuration slot to associate with this machine configuration.
	ConfigurationSlotId int32 `json:"ConfigurationSlotId"`
	// Name of the new machine configuration.
	LeafName string `json:"LeafName"`
	// Policy settings data created with the SDK snap-in that matches the SettingsGroup of the configuration slot. Base64 encrypted.
	Policy string `json:"Policy"`
	// Description of the new machine configuration.
	Description NullableString `json:"Description,omitempty"`
}

// NewCreateMachineConfigurationRequestModel instantiates a new CreateMachineConfigurationRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMachineConfigurationRequestModel(configurationSlotId int32, leafName string, policy string) *CreateMachineConfigurationRequestModel {
	this := CreateMachineConfigurationRequestModel{}
	this.ConfigurationSlotId = configurationSlotId
	this.LeafName = leafName
	this.Policy = policy
	return &this
}

// NewCreateMachineConfigurationRequestModelWithDefaults instantiates a new CreateMachineConfigurationRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMachineConfigurationRequestModelWithDefaults() *CreateMachineConfigurationRequestModel {
	this := CreateMachineConfigurationRequestModel{}
	return &this
}

// GetConfigurationSlotId returns the ConfigurationSlotId field value
func (o *CreateMachineConfigurationRequestModel) GetConfigurationSlotId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfigurationSlotId
}

// GetConfigurationSlotIdOk returns a tuple with the ConfigurationSlotId field value
// and a boolean to check if the value has been set.
func (o *CreateMachineConfigurationRequestModel) GetConfigurationSlotIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationSlotId, true
}

// SetConfigurationSlotId sets field value
func (o *CreateMachineConfigurationRequestModel) SetConfigurationSlotId(v int32) {
	o.ConfigurationSlotId = v
}

// GetLeafName returns the LeafName field value
func (o *CreateMachineConfigurationRequestModel) GetLeafName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeafName
}

// GetLeafNameOk returns a tuple with the LeafName field value
// and a boolean to check if the value has been set.
func (o *CreateMachineConfigurationRequestModel) GetLeafNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeafName, true
}

// SetLeafName sets field value
func (o *CreateMachineConfigurationRequestModel) SetLeafName(v string) {
	o.LeafName = v
}

// GetPolicy returns the Policy field value
func (o *CreateMachineConfigurationRequestModel) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *CreateMachineConfigurationRequestModel) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *CreateMachineConfigurationRequestModel) SetPolicy(v string) {
	o.Policy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineConfigurationRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineConfigurationRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMachineConfigurationRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateMachineConfigurationRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateMachineConfigurationRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateMachineConfigurationRequestModel) UnsetDescription() {
	o.Description.Unset()
}

func (o CreateMachineConfigurationRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMachineConfigurationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ConfigurationSlotId"] = o.ConfigurationSlotId
	toSerialize["LeafName"] = o.LeafName
	toSerialize["Policy"] = o.Policy
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableCreateMachineConfigurationRequestModel struct {
	value *CreateMachineConfigurationRequestModel
	isSet bool
}

func (v NullableCreateMachineConfigurationRequestModel) Get() *CreateMachineConfigurationRequestModel {
	return v.value
}

func (v *NullableCreateMachineConfigurationRequestModel) Set(val *CreateMachineConfigurationRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMachineConfigurationRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMachineConfigurationRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMachineConfigurationRequestModel(val *CreateMachineConfigurationRequestModel) *NullableCreateMachineConfigurationRequestModel {
	return &NullableCreateMachineConfigurationRequestModel{value: val, isSet: true}
}

func (v NullableCreateMachineConfigurationRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMachineConfigurationRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


