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

// checks if the AddApplicationsRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddApplicationsRequestModel{}

// AddApplicationsRequestModel Request object for adding applications to the site.
type AddApplicationsRequestModel struct {
	// List of existing applications to be added to the DeliveryGroups and/or ApplicationGroups.
	ExistingApplications []string `json:"ExistingApplications,omitempty"`
	// List of applications which should be created and added to the DeliveryGroups and/or ApplicationGroups.
	NewApplications []CreateApplicationRequestModel `json:"NewApplications,omitempty"`
	// List of application groups to add the ExistingApplications and NewApplications to.
	ApplicationGroups []string `json:"ApplicationGroups,omitempty"`
	// List of delivery groups to add the ExistingApplications and NewApplications to.
	DeliveryGroups []PriorityRefRequestModel `json:"DeliveryGroups,omitempty"`
}

// NewAddApplicationsRequestModel instantiates a new AddApplicationsRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddApplicationsRequestModel() *AddApplicationsRequestModel {
	this := AddApplicationsRequestModel{}
	return &this
}

// NewAddApplicationsRequestModelWithDefaults instantiates a new AddApplicationsRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddApplicationsRequestModelWithDefaults() *AddApplicationsRequestModel {
	this := AddApplicationsRequestModel{}
	return &this
}

// GetExistingApplications returns the ExistingApplications field value if set, zero value otherwise.
func (o *AddApplicationsRequestModel) GetExistingApplications() []string {
	if o == nil || IsNil(o.ExistingApplications) {
		var ret []string
		return ret
	}
	return o.ExistingApplications
}

// GetExistingApplicationsOk returns a tuple with the ExistingApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddApplicationsRequestModel) GetExistingApplicationsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExistingApplications) {
		return nil, false
	}
	return o.ExistingApplications, true
}

// HasExistingApplications returns a boolean if a field has been set.
func (o *AddApplicationsRequestModel) HasExistingApplications() bool {
	if o != nil && !IsNil(o.ExistingApplications) {
		return true
	}

	return false
}

// SetExistingApplications gets a reference to the given []string and assigns it to the ExistingApplications field.
func (o *AddApplicationsRequestModel) SetExistingApplications(v []string) {
	o.ExistingApplications = v
}

// GetNewApplications returns the NewApplications field value if set, zero value otherwise.
func (o *AddApplicationsRequestModel) GetNewApplications() []CreateApplicationRequestModel {
	if o == nil || IsNil(o.NewApplications) {
		var ret []CreateApplicationRequestModel
		return ret
	}
	return o.NewApplications
}

// GetNewApplicationsOk returns a tuple with the NewApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddApplicationsRequestModel) GetNewApplicationsOk() ([]CreateApplicationRequestModel, bool) {
	if o == nil || IsNil(o.NewApplications) {
		return nil, false
	}
	return o.NewApplications, true
}

// HasNewApplications returns a boolean if a field has been set.
func (o *AddApplicationsRequestModel) HasNewApplications() bool {
	if o != nil && !IsNil(o.NewApplications) {
		return true
	}

	return false
}

// SetNewApplications gets a reference to the given []CreateApplicationRequestModel and assigns it to the NewApplications field.
func (o *AddApplicationsRequestModel) SetNewApplications(v []CreateApplicationRequestModel) {
	o.NewApplications = v
}

// GetApplicationGroups returns the ApplicationGroups field value if set, zero value otherwise.
func (o *AddApplicationsRequestModel) GetApplicationGroups() []string {
	if o == nil || IsNil(o.ApplicationGroups) {
		var ret []string
		return ret
	}
	return o.ApplicationGroups
}

// GetApplicationGroupsOk returns a tuple with the ApplicationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddApplicationsRequestModel) GetApplicationGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationGroups) {
		return nil, false
	}
	return o.ApplicationGroups, true
}

// HasApplicationGroups returns a boolean if a field has been set.
func (o *AddApplicationsRequestModel) HasApplicationGroups() bool {
	if o != nil && !IsNil(o.ApplicationGroups) {
		return true
	}

	return false
}

// SetApplicationGroups gets a reference to the given []string and assigns it to the ApplicationGroups field.
func (o *AddApplicationsRequestModel) SetApplicationGroups(v []string) {
	o.ApplicationGroups = v
}

// GetDeliveryGroups returns the DeliveryGroups field value if set, zero value otherwise.
func (o *AddApplicationsRequestModel) GetDeliveryGroups() []PriorityRefRequestModel {
	if o == nil || IsNil(o.DeliveryGroups) {
		var ret []PriorityRefRequestModel
		return ret
	}
	return o.DeliveryGroups
}

// GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddApplicationsRequestModel) GetDeliveryGroupsOk() ([]PriorityRefRequestModel, bool) {
	if o == nil || IsNil(o.DeliveryGroups) {
		return nil, false
	}
	return o.DeliveryGroups, true
}

// HasDeliveryGroups returns a boolean if a field has been set.
func (o *AddApplicationsRequestModel) HasDeliveryGroups() bool {
	if o != nil && !IsNil(o.DeliveryGroups) {
		return true
	}

	return false
}

// SetDeliveryGroups gets a reference to the given []PriorityRefRequestModel and assigns it to the DeliveryGroups field.
func (o *AddApplicationsRequestModel) SetDeliveryGroups(v []PriorityRefRequestModel) {
	o.DeliveryGroups = v
}

func (o AddApplicationsRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddApplicationsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExistingApplications) {
		toSerialize["ExistingApplications"] = o.ExistingApplications
	}
	if !IsNil(o.NewApplications) {
		toSerialize["NewApplications"] = o.NewApplications
	}
	if !IsNil(o.ApplicationGroups) {
		toSerialize["ApplicationGroups"] = o.ApplicationGroups
	}
	if !IsNil(o.DeliveryGroups) {
		toSerialize["DeliveryGroups"] = o.DeliveryGroups
	}
	return toSerialize, nil
}

type NullableAddApplicationsRequestModel struct {
	value *AddApplicationsRequestModel
	isSet bool
}

func (v NullableAddApplicationsRequestModel) Get() *AddApplicationsRequestModel {
	return v.value
}

func (v *NullableAddApplicationsRequestModel) Set(val *AddApplicationsRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddApplicationsRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddApplicationsRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddApplicationsRequestModel(val *AddApplicationsRequestModel) *NullableAddApplicationsRequestModel {
	return &NullableAddApplicationsRequestModel{value: val, isSet: true}
}

func (v NullableAddApplicationsRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddApplicationsRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


