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

// checks if the TagResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagResponseModel{}

// TagResponseModel Response object for a tag.
type TagResponseModel struct {
	// Name of the tag.
	Name string `json:"Name"`
	// Id of the tag. Used to be: Guid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility
	Id string `json:"Id"`
	// Description of the tag.
	Description *string `json:"Description,omitempty"`
	// `DEPRECATED` DEPRECATED.  Use Id.
	Uid *int32 `json:"Uid,omitempty"`
	// Used by Gpo or not.
	IsUsedByGpo *bool `json:"IsUsedByGpo,omitempty"`
	// Number of machines the tag is applied to, within the context of the query.
	NumMachines int32 `json:"NumMachines"`
	// Number of applications the tag is applied to, within the context of the query.
	NumApplications int32 `json:"NumApplications"`
	// Number of application groups the tag is applied to, within the context of the query.
	NumApplicationGroups int32 `json:"NumApplicationGroups"`
	// Number of machine catalogs the tag is applied to, within the context of the query.
	NumMachineCatalogs int32 `json:"NumMachineCatalogs"`
	// Number of delivery groups the tag is applied to, within the context of the query.
	NumDeliveryGroups int32 `json:"NumDeliveryGroups"`
	// Number of delivery groups using this tag for Autoscale, within the context of the query.
	NumAutoscale *int32 `json:"NumAutoscale,omitempty"`
	// The number of objects of all types that are tagged with this tag and that are *not* visible to the calling delegated administrator.
	NumUnknownObjects int32 `json:"NumUnknownObjects"`
	// The list of the delegated admin scopes to which the tag belongs.
	ScopeReferences []ScopeReferenceModel `json:"ScopeReferences,omitempty"`
}

// NewTagResponseModel instantiates a new TagResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResponseModel(name string, id string, numMachines int32, numApplications int32, numApplicationGroups int32, numMachineCatalogs int32, numDeliveryGroups int32, numUnknownObjects int32) *TagResponseModel {
	this := TagResponseModel{}
	this.Name = name
	this.Id = id
	this.NumMachines = numMachines
	this.NumApplications = numApplications
	this.NumApplicationGroups = numApplicationGroups
	this.NumMachineCatalogs = numMachineCatalogs
	this.NumDeliveryGroups = numDeliveryGroups
	this.NumUnknownObjects = numUnknownObjects
	return &this
}

// NewTagResponseModelWithDefaults instantiates a new TagResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResponseModelWithDefaults() *TagResponseModel {
	this := TagResponseModel{}
	return &this
}

// GetName returns the Name field value
func (o *TagResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagResponseModel) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *TagResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagResponseModel) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *TagResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *TagResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *TagResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetIsUsedByGpo returns the IsUsedByGpo field value if set, zero value otherwise.
func (o *TagResponseModel) GetIsUsedByGpo() bool {
	if o == nil || IsNil(o.IsUsedByGpo) {
		var ret bool
		return ret
	}
	return *o.IsUsedByGpo
}

// GetIsUsedByGpoOk returns a tuple with the IsUsedByGpo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetIsUsedByGpoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUsedByGpo) {
		return nil, false
	}
	return o.IsUsedByGpo, true
}

// HasIsUsedByGpo returns a boolean if a field has been set.
func (o *TagResponseModel) HasIsUsedByGpo() bool {
	if o != nil && !IsNil(o.IsUsedByGpo) {
		return true
	}

	return false
}

// SetIsUsedByGpo gets a reference to the given bool and assigns it to the IsUsedByGpo field.
func (o *TagResponseModel) SetIsUsedByGpo(v bool) {
	o.IsUsedByGpo = &v
}

// GetNumMachines returns the NumMachines field value
func (o *TagResponseModel) GetNumMachines() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumMachines
}

// GetNumMachinesOk returns a tuple with the NumMachines field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumMachinesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumMachines, true
}

// SetNumMachines sets field value
func (o *TagResponseModel) SetNumMachines(v int32) {
	o.NumMachines = v
}

// GetNumApplications returns the NumApplications field value
func (o *TagResponseModel) GetNumApplications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumApplications
}

// GetNumApplicationsOk returns a tuple with the NumApplications field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumApplicationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumApplications, true
}

// SetNumApplications sets field value
func (o *TagResponseModel) SetNumApplications(v int32) {
	o.NumApplications = v
}

// GetNumApplicationGroups returns the NumApplicationGroups field value
func (o *TagResponseModel) GetNumApplicationGroups() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumApplicationGroups
}

// GetNumApplicationGroupsOk returns a tuple with the NumApplicationGroups field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumApplicationGroupsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumApplicationGroups, true
}

// SetNumApplicationGroups sets field value
func (o *TagResponseModel) SetNumApplicationGroups(v int32) {
	o.NumApplicationGroups = v
}

// GetNumMachineCatalogs returns the NumMachineCatalogs field value
func (o *TagResponseModel) GetNumMachineCatalogs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumMachineCatalogs
}

// GetNumMachineCatalogsOk returns a tuple with the NumMachineCatalogs field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumMachineCatalogsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumMachineCatalogs, true
}

// SetNumMachineCatalogs sets field value
func (o *TagResponseModel) SetNumMachineCatalogs(v int32) {
	o.NumMachineCatalogs = v
}

// GetNumDeliveryGroups returns the NumDeliveryGroups field value
func (o *TagResponseModel) GetNumDeliveryGroups() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumDeliveryGroups
}

// GetNumDeliveryGroupsOk returns a tuple with the NumDeliveryGroups field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumDeliveryGroupsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumDeliveryGroups, true
}

// SetNumDeliveryGroups sets field value
func (o *TagResponseModel) SetNumDeliveryGroups(v int32) {
	o.NumDeliveryGroups = v
}

// GetNumAutoscale returns the NumAutoscale field value if set, zero value otherwise.
func (o *TagResponseModel) GetNumAutoscale() int32 {
	if o == nil || IsNil(o.NumAutoscale) {
		var ret int32
		return ret
	}
	return *o.NumAutoscale
}

// GetNumAutoscaleOk returns a tuple with the NumAutoscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumAutoscaleOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAutoscale) {
		return nil, false
	}
	return o.NumAutoscale, true
}

// HasNumAutoscale returns a boolean if a field has been set.
func (o *TagResponseModel) HasNumAutoscale() bool {
	if o != nil && !IsNil(o.NumAutoscale) {
		return true
	}

	return false
}

// SetNumAutoscale gets a reference to the given int32 and assigns it to the NumAutoscale field.
func (o *TagResponseModel) SetNumAutoscale(v int32) {
	o.NumAutoscale = &v
}

// GetNumUnknownObjects returns the NumUnknownObjects field value
func (o *TagResponseModel) GetNumUnknownObjects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumUnknownObjects
}

// GetNumUnknownObjectsOk returns a tuple with the NumUnknownObjects field value
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetNumUnknownObjectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumUnknownObjects, true
}

// SetNumUnknownObjects sets field value
func (o *TagResponseModel) SetNumUnknownObjects(v int32) {
	o.NumUnknownObjects = v
}

// GetScopeReferences returns the ScopeReferences field value if set, zero value otherwise.
func (o *TagResponseModel) GetScopeReferences() []ScopeReferenceModel {
	if o == nil || IsNil(o.ScopeReferences) {
		var ret []ScopeReferenceModel
		return ret
	}
	return o.ScopeReferences
}

// GetScopeReferencesOk returns a tuple with the ScopeReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseModel) GetScopeReferencesOk() ([]ScopeReferenceModel, bool) {
	if o == nil || IsNil(o.ScopeReferences) {
		return nil, false
	}
	return o.ScopeReferences, true
}

// HasScopeReferences returns a boolean if a field has been set.
func (o *TagResponseModel) HasScopeReferences() bool {
	if o != nil && !IsNil(o.ScopeReferences) {
		return true
	}

	return false
}

// SetScopeReferences gets a reference to the given []ScopeReferenceModel and assigns it to the ScopeReferences field.
func (o *TagResponseModel) SetScopeReferences(v []ScopeReferenceModel) {
	o.ScopeReferences = v
}

func (o TagResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	toSerialize["Id"] = o.Id
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	if !IsNil(o.IsUsedByGpo) {
		toSerialize["IsUsedByGpo"] = o.IsUsedByGpo
	}
	toSerialize["NumMachines"] = o.NumMachines
	toSerialize["NumApplications"] = o.NumApplications
	toSerialize["NumApplicationGroups"] = o.NumApplicationGroups
	toSerialize["NumMachineCatalogs"] = o.NumMachineCatalogs
	toSerialize["NumDeliveryGroups"] = o.NumDeliveryGroups
	if !IsNil(o.NumAutoscale) {
		toSerialize["NumAutoscale"] = o.NumAutoscale
	}
	toSerialize["NumUnknownObjects"] = o.NumUnknownObjects
	if !IsNil(o.ScopeReferences) {
		toSerialize["ScopeReferences"] = o.ScopeReferences
	}
	return toSerialize, nil
}

type NullableTagResponseModel struct {
	value *TagResponseModel
	isSet bool
}

func (v NullableTagResponseModel) Get() *TagResponseModel {
	return v.value
}

func (v *NullableTagResponseModel) Set(val *TagResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResponseModel(val *TagResponseModel) *NullableTagResponseModel {
	return &NullableTagResponseModel{value: val, isSet: true}
}

func (v NullableTagResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


