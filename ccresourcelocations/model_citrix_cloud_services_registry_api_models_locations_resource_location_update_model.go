/*
Resource Locations APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccresourcelocations

import (
	"encoding/json"
)

// checks if the CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel{}

// CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel Information used to update a customer resource location.
type CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel struct {
	// The new Resource Location Name.
	Name *string `json:"name,omitempty"`
	// Resource Location Connectivity
	InternalOnly *bool `json:"internalOnly,omitempty"`
	// Time zone.
	TimeZone *string `json:"timeZone,omitempty"`
	// If true means the resource location is managed by citrix.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel instantiates a new CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel() *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel {
	this := CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel{}
	return &this
}

// NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModelWithDefaults instantiates a new CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModelWithDefaults() *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel {
	this := CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) SetName(v string) {
	o.Name = &v
}

// GetInternalOnly returns the InternalOnly field value if set, zero value otherwise.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetInternalOnly() bool {
	if o == nil || IsNil(o.InternalOnly) {
		var ret bool
		return ret
	}
	return *o.InternalOnly
}

// GetInternalOnlyOk returns a tuple with the InternalOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetInternalOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.InternalOnly) {
		return nil, false
	}
	return o.InternalOnly, true
}

// HasInternalOnly returns a boolean if a field has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) HasInternalOnly() bool {
	if o != nil && !IsNil(o.InternalOnly) {
		return true
	}

	return false
}

// SetInternalOnly gets a reference to the given bool and assigns it to the InternalOnly field.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) SetInternalOnly(v bool) {
	o.InternalOnly = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.InternalOnly) {
		toSerialize["internalOnly"] = o.InternalOnly
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel struct {
	value *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel
	isSet bool
}

func (v NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) Get() *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel {
	return v.value
}

func (v *NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) Set(val *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel(val *CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) *NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel {
	return &NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


