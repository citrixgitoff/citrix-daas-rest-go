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

// checks if the ScopeResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeResponseModel{}

// ScopeResponseModel Response object representing an administrative scope.
type ScopeResponseModel struct {
	// Id of the administrative scope. Used to be: ScopeId (and it was not globally unique) Needs to be globally unique Might be constructed from site ID + internal Uid
	Id string `json:"Id"`
	// `DEPRECATED.  Use <see cref='Id'/>.` DEPRECATED.  Use Id.
	Uid *int32 `json:"Uid,omitempty"`
	// Name of the administrative scope.
	Name string `json:"Name"`
	// Description of the administrative scope.
	Description *string `json:"Description,omitempty"`
	// Whether the administrative scope is built-in.
	IsBuiltIn bool `json:"IsBuiltIn"`
	// Indicates the built-in \"All\" scope.  There will be exactly one scope with this property set to `true`.
	IsAllScope bool `json:"IsAllScope"`
	// Whether the scope is created for CSP tenant.
	IsTenantScope bool `json:"IsTenantScope"`
	// Id of the CSP tenant. Valid when IsTenantScope is true.
	TenantId *string `json:"TenantId,omitempty"`
	// Name of the CSP tenant. Valid when IsTenantScope is true.
	TenantName *string `json:"TenantName,omitempty"`
}

// NewScopeResponseModel instantiates a new ScopeResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeResponseModel(id string, name string, isBuiltIn bool, isAllScope bool, isTenantScope bool) *ScopeResponseModel {
	this := ScopeResponseModel{}
	this.Id = id
	this.Name = name
	this.IsBuiltIn = isBuiltIn
	this.IsAllScope = isAllScope
	this.IsTenantScope = isTenantScope
	return &this
}

// NewScopeResponseModelWithDefaults instantiates a new ScopeResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeResponseModelWithDefaults() *ScopeResponseModel {
	this := ScopeResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ScopeResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScopeResponseModel) SetId(v string) {
	o.Id = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ScopeResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ScopeResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *ScopeResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetName returns the Name field value
func (o *ScopeResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopeResponseModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScopeResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScopeResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScopeResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetIsBuiltIn returns the IsBuiltIn field value
func (o *ScopeResponseModel) GetIsBuiltIn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBuiltIn
}

// GetIsBuiltInOk returns a tuple with the IsBuiltIn field value
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetIsBuiltInOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBuiltIn, true
}

// SetIsBuiltIn sets field value
func (o *ScopeResponseModel) SetIsBuiltIn(v bool) {
	o.IsBuiltIn = v
}

// GetIsAllScope returns the IsAllScope field value
func (o *ScopeResponseModel) GetIsAllScope() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAllScope
}

// GetIsAllScopeOk returns a tuple with the IsAllScope field value
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetIsAllScopeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAllScope, true
}

// SetIsAllScope sets field value
func (o *ScopeResponseModel) SetIsAllScope(v bool) {
	o.IsAllScope = v
}

// GetIsTenantScope returns the IsTenantScope field value
func (o *ScopeResponseModel) GetIsTenantScope() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTenantScope
}

// GetIsTenantScopeOk returns a tuple with the IsTenantScope field value
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetIsTenantScopeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTenantScope, true
}

// SetIsTenantScope sets field value
func (o *ScopeResponseModel) SetIsTenantScope(v bool) {
	o.IsTenantScope = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ScopeResponseModel) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ScopeResponseModel) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ScopeResponseModel) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *ScopeResponseModel) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponseModel) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *ScopeResponseModel) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *ScopeResponseModel) SetTenantName(v string) {
	o.TenantName = &v
}

func (o ScopeResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	toSerialize["Name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["IsBuiltIn"] = o.IsBuiltIn
	toSerialize["IsAllScope"] = o.IsAllScope
	toSerialize["IsTenantScope"] = o.IsTenantScope
	if !IsNil(o.TenantId) {
		toSerialize["TenantId"] = o.TenantId
	}
	if !IsNil(o.TenantName) {
		toSerialize["TenantName"] = o.TenantName
	}
	return toSerialize, nil
}

type NullableScopeResponseModel struct {
	value *ScopeResponseModel
	isSet bool
}

func (v NullableScopeResponseModel) Get() *ScopeResponseModel {
	return v.value
}

func (v *NullableScopeResponseModel) Set(val *ScopeResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeResponseModel(val *ScopeResponseModel) *NullableScopeResponseModel {
	return &NullableScopeResponseModel{value: val, isSet: true}
}

func (v NullableScopeResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


