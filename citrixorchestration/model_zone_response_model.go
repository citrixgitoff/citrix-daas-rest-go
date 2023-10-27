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

// checks if the ZoneResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneResponseModel{}

// ZoneResponseModel Zone object.
type ZoneResponseModel struct {
	// ID of the zone.
	Id string `json:"Id"`
	// Name of the zone.
	Name string `json:"Name"`
	// Description of the zone.
	Description NullableString `json:"Description,omitempty"`
	// Indicates whether the zone is the primary zone for the site.  Resources in the site default to the primary zone if not set otherwise.
	IsPrimary bool `json:"IsPrimary"`
	// The flag to indicate whether the zone is healthy.
	IsHealthy *bool `json:"IsHealthy,omitempty"`
	// The LastStateChangeTimeInUtc of the zone.
	LastStateChangeTimeInUtc NullableString `json:"LastStateChangeTimeInUtc,omitempty"`
	// The metadata of Zone.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	ResourceLocation *RefResponseModel `json:"ResourceLocation,omitempty"`
}

// NewZoneResponseModel instantiates a new ZoneResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneResponseModel(id string, name string, isPrimary bool) *ZoneResponseModel {
	this := ZoneResponseModel{}
	this.Id = id
	this.Name = name
	this.IsPrimary = isPrimary
	return &this
}

// NewZoneResponseModelWithDefaults instantiates a new ZoneResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneResponseModelWithDefaults() *ZoneResponseModel {
	this := ZoneResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ZoneResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoneResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ZoneResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ZoneResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoneResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZoneResponseModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZoneResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ZoneResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ZoneResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ZoneResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ZoneResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetIsPrimary returns the IsPrimary field value
func (o *ZoneResponseModel) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *ZoneResponseModel) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *ZoneResponseModel) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetIsHealthy returns the IsHealthy field value if set, zero value otherwise.
func (o *ZoneResponseModel) GetIsHealthy() bool {
	if o == nil || IsNil(o.IsHealthy) {
		var ret bool
		return ret
	}
	return *o.IsHealthy
}

// GetIsHealthyOk returns a tuple with the IsHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneResponseModel) GetIsHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHealthy) {
		return nil, false
	}
	return o.IsHealthy, true
}

// HasIsHealthy returns a boolean if a field has been set.
func (o *ZoneResponseModel) HasIsHealthy() bool {
	if o != nil && !IsNil(o.IsHealthy) {
		return true
	}

	return false
}

// SetIsHealthy gets a reference to the given bool and assigns it to the IsHealthy field.
func (o *ZoneResponseModel) SetIsHealthy(v bool) {
	o.IsHealthy = &v
}

// GetLastStateChangeTimeInUtc returns the LastStateChangeTimeInUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZoneResponseModel) GetLastStateChangeTimeInUtc() string {
	if o == nil || IsNil(o.LastStateChangeTimeInUtc.Get()) {
		var ret string
		return ret
	}
	return *o.LastStateChangeTimeInUtc.Get()
}

// GetLastStateChangeTimeInUtcOk returns a tuple with the LastStateChangeTimeInUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneResponseModel) GetLastStateChangeTimeInUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastStateChangeTimeInUtc.Get(), o.LastStateChangeTimeInUtc.IsSet()
}

// HasLastStateChangeTimeInUtc returns a boolean if a field has been set.
func (o *ZoneResponseModel) HasLastStateChangeTimeInUtc() bool {
	if o != nil && o.LastStateChangeTimeInUtc.IsSet() {
		return true
	}

	return false
}

// SetLastStateChangeTimeInUtc gets a reference to the given NullableString and assigns it to the LastStateChangeTimeInUtc field.
func (o *ZoneResponseModel) SetLastStateChangeTimeInUtc(v string) {
	o.LastStateChangeTimeInUtc.Set(&v)
}
// SetLastStateChangeTimeInUtcNil sets the value for LastStateChangeTimeInUtc to be an explicit nil
func (o *ZoneResponseModel) SetLastStateChangeTimeInUtcNil() {
	o.LastStateChangeTimeInUtc.Set(nil)
}

// UnsetLastStateChangeTimeInUtc ensures that no value is present for LastStateChangeTimeInUtc, not even an explicit nil
func (o *ZoneResponseModel) UnsetLastStateChangeTimeInUtc() {
	o.LastStateChangeTimeInUtc.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZoneResponseModel) GetMetadata() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneResponseModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ZoneResponseModel) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *ZoneResponseModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetResourceLocation returns the ResourceLocation field value if set, zero value otherwise.
func (o *ZoneResponseModel) GetResourceLocation() RefResponseModel {
	if o == nil || IsNil(o.ResourceLocation) {
		var ret RefResponseModel
		return ret
	}
	return *o.ResourceLocation
}

// GetResourceLocationOk returns a tuple with the ResourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneResponseModel) GetResourceLocationOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.ResourceLocation) {
		return nil, false
	}
	return o.ResourceLocation, true
}

// HasResourceLocation returns a boolean if a field has been set.
func (o *ZoneResponseModel) HasResourceLocation() bool {
	if o != nil && !IsNil(o.ResourceLocation) {
		return true
	}

	return false
}

// SetResourceLocation gets a reference to the given RefResponseModel and assigns it to the ResourceLocation field.
func (o *ZoneResponseModel) SetResourceLocation(v RefResponseModel) {
	o.ResourceLocation = &v
}

func (o ZoneResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	toSerialize["IsPrimary"] = o.IsPrimary
	if !IsNil(o.IsHealthy) {
		toSerialize["IsHealthy"] = o.IsHealthy
	}
	if o.LastStateChangeTimeInUtc.IsSet() {
		toSerialize["LastStateChangeTimeInUtc"] = o.LastStateChangeTimeInUtc.Get()
	}
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.ResourceLocation) {
		toSerialize["ResourceLocation"] = o.ResourceLocation
	}
	return toSerialize, nil
}

type NullableZoneResponseModel struct {
	value *ZoneResponseModel
	isSet bool
}

func (v NullableZoneResponseModel) Get() *ZoneResponseModel {
	return v.value
}

func (v *NullableZoneResponseModel) Set(val *ZoneResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneResponseModel(val *ZoneResponseModel) *NullableZoneResponseModel {
	return &NullableZoneResponseModel{value: val, isSet: true}
}

func (v NullableZoneResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


