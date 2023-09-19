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

// checks if the PredefinedPermissionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPermissionResponseModel{}

// PredefinedPermissionResponseModel Response object for a predefined permission
type PredefinedPermissionResponseModel struct {
	// The description of the permission. It is usually missed (null or empty).
	Description NullableString `json:"Description,omitempty"`
	// The group id of the permission.
	GroupId string `json:"GroupId"`
	// The group name of the permission.
	GroupName string `json:"GroupName"`
	// The Id of the permission. It is globally unique.
	Id string `json:"Id"`
	// The name of the permission. It is a friendly name of permission.
	Name string `json:"Name"`
	// The list of operations of the permission.
	Operations []string `json:"Operations"`
	// Whether any of the operations of the permission may change object status.
	IsReadOnly bool `json:"IsReadOnly"`
}

// NewPredefinedPermissionResponseModel instantiates a new PredefinedPermissionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPermissionResponseModel(groupId string, groupName string, id string, name string, operations []string, isReadOnly bool) *PredefinedPermissionResponseModel {
	this := PredefinedPermissionResponseModel{}
	this.GroupId = groupId
	this.GroupName = groupName
	this.Id = id
	this.Name = name
	this.Operations = operations
	this.IsReadOnly = isReadOnly
	return &this
}

// NewPredefinedPermissionResponseModelWithDefaults instantiates a new PredefinedPermissionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPermissionResponseModelWithDefaults() *PredefinedPermissionResponseModel {
	this := PredefinedPermissionResponseModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredefinedPermissionResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredefinedPermissionResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PredefinedPermissionResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PredefinedPermissionResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PredefinedPermissionResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PredefinedPermissionResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetGroupId returns the GroupId field value
func (o *PredefinedPermissionResponseModel) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *PredefinedPermissionResponseModel) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupName returns the GroupName field value
func (o *PredefinedPermissionResponseModel) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *PredefinedPermissionResponseModel) SetGroupName(v string) {
	o.GroupName = v
}

// GetId returns the Id field value
func (o *PredefinedPermissionResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PredefinedPermissionResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PredefinedPermissionResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PredefinedPermissionResponseModel) SetName(v string) {
	o.Name = v
}

// GetOperations returns the Operations field value
func (o *PredefinedPermissionResponseModel) GetOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetOperationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *PredefinedPermissionResponseModel) SetOperations(v []string) {
	o.Operations = v
}

// GetIsReadOnly returns the IsReadOnly field value
func (o *PredefinedPermissionResponseModel) GetIsReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value
// and a boolean to check if the value has been set.
func (o *PredefinedPermissionResponseModel) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReadOnly, true
}

// SetIsReadOnly sets field value
func (o *PredefinedPermissionResponseModel) SetIsReadOnly(v bool) {
	o.IsReadOnly = v
}

func (o PredefinedPermissionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPermissionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	toSerialize["GroupId"] = o.GroupId
	toSerialize["GroupName"] = o.GroupName
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	toSerialize["Operations"] = o.Operations
	toSerialize["IsReadOnly"] = o.IsReadOnly
	return toSerialize, nil
}

type NullablePredefinedPermissionResponseModel struct {
	value *PredefinedPermissionResponseModel
	isSet bool
}

func (v NullablePredefinedPermissionResponseModel) Get() *PredefinedPermissionResponseModel {
	return v.value
}

func (v *NullablePredefinedPermissionResponseModel) Set(val *PredefinedPermissionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPermissionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPermissionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPermissionResponseModel(val *PredefinedPermissionResponseModel) *NullablePredefinedPermissionResponseModel {
	return &NullablePredefinedPermissionResponseModel{value: val, isSet: true}
}

func (v NullablePredefinedPermissionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPermissionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


