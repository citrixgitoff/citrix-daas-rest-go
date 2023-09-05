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

// checks if the UpdateAppVIsolationGroupRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAppVIsolationGroupRequestModel{}

// UpdateAppVIsolationGroupRequestModel Request object for an App-V IsolationGroup.
type UpdateAppVIsolationGroupRequestModel struct {
	// Name of IsolationGroup
	Name *string `json:"Name,omitempty"`
	// Description of IsolationGroup
	Description *string `json:"Description,omitempty"`
	// Included AppV packages
	IncludedAppVPackages []AppVIsolationGroupPackageRequestModel `json:"IncludedAppVPackages,omitempty"`
}

// NewUpdateAppVIsolationGroupRequestModel instantiates a new UpdateAppVIsolationGroupRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAppVIsolationGroupRequestModel() *UpdateAppVIsolationGroupRequestModel {
	this := UpdateAppVIsolationGroupRequestModel{}
	return &this
}

// NewUpdateAppVIsolationGroupRequestModelWithDefaults instantiates a new UpdateAppVIsolationGroupRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppVIsolationGroupRequestModelWithDefaults() *UpdateAppVIsolationGroupRequestModel {
	this := UpdateAppVIsolationGroupRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAppVIsolationGroupRequestModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppVIsolationGroupRequestModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAppVIsolationGroupRequestModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAppVIsolationGroupRequestModel) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateAppVIsolationGroupRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppVIsolationGroupRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAppVIsolationGroupRequestModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateAppVIsolationGroupRequestModel) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedAppVPackages returns the IncludedAppVPackages field value if set, zero value otherwise.
func (o *UpdateAppVIsolationGroupRequestModel) GetIncludedAppVPackages() []AppVIsolationGroupPackageRequestModel {
	if o == nil || IsNil(o.IncludedAppVPackages) {
		var ret []AppVIsolationGroupPackageRequestModel
		return ret
	}
	return o.IncludedAppVPackages
}

// GetIncludedAppVPackagesOk returns a tuple with the IncludedAppVPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppVIsolationGroupRequestModel) GetIncludedAppVPackagesOk() ([]AppVIsolationGroupPackageRequestModel, bool) {
	if o == nil || IsNil(o.IncludedAppVPackages) {
		return nil, false
	}
	return o.IncludedAppVPackages, true
}

// HasIncludedAppVPackages returns a boolean if a field has been set.
func (o *UpdateAppVIsolationGroupRequestModel) HasIncludedAppVPackages() bool {
	if o != nil && !IsNil(o.IncludedAppVPackages) {
		return true
	}

	return false
}

// SetIncludedAppVPackages gets a reference to the given []AppVIsolationGroupPackageRequestModel and assigns it to the IncludedAppVPackages field.
func (o *UpdateAppVIsolationGroupRequestModel) SetIncludedAppVPackages(v []AppVIsolationGroupPackageRequestModel) {
	o.IncludedAppVPackages = v
}

func (o UpdateAppVIsolationGroupRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAppVIsolationGroupRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.IncludedAppVPackages) {
		toSerialize["IncludedAppVPackages"] = o.IncludedAppVPackages
	}
	return toSerialize, nil
}

type NullableUpdateAppVIsolationGroupRequestModel struct {
	value *UpdateAppVIsolationGroupRequestModel
	isSet bool
}

func (v NullableUpdateAppVIsolationGroupRequestModel) Get() *UpdateAppVIsolationGroupRequestModel {
	return v.value
}

func (v *NullableUpdateAppVIsolationGroupRequestModel) Set(val *UpdateAppVIsolationGroupRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAppVIsolationGroupRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAppVIsolationGroupRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAppVIsolationGroupRequestModel(val *UpdateAppVIsolationGroupRequestModel) *NullableUpdateAppVIsolationGroupRequestModel {
	return &NullableUpdateAppVIsolationGroupRequestModel{value: val, isSet: true}
}

func (v NullableUpdateAppVIsolationGroupRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAppVIsolationGroupRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


