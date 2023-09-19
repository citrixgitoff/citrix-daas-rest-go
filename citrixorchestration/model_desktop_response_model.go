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

// checks if the DesktopResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesktopResponseModel{}

// DesktopResponseModel Response object for a published desktop.
type DesktopResponseModel struct {
	// Id of the desktop.
	Id string `json:"Id"`
	// `DEPRECATED` DEPRECATED. Use Id.
	// Deprecated
	Uid *int32 `json:"Uid,omitempty"`
	ColorDepth ColorDepth `json:"ColorDepth"`
	// Optional description of the desktop. The text may be visible to the end user, for example, as a tooltip associated with the desktop within receiver.
	Description NullableString `json:"Description,omitempty"`
	// Whether the published desktop is enabled. A disabled desktop is ignored when determining which desktops are available for a user.
	Enabled bool `json:"Enabled"`
	// Indicates whether the ExcludedUsers filter is enabled.  If the filter is disabled then any user entries in the filter are ignored when determining which desktops are available for a user.
	ExcludedUserFilterEnabled bool `json:"ExcludedUserFilterEnabled"`
	// The excluded users filter of the desktop; that is, the users and groups who are explicitly denied access to the published desktop.
	ExcludedUsers []IdentityUserResponseModel `json:"ExcludedUsers,omitempty"`
	// Id of the icon used to display the published desktop to the user, and of assigned desktop(s) in the case where SharingKind is equal to Private.
	IconId string `json:"IconId"`
	// Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group's access policy is implicitly granted an entitlement to the published desktop.
	IncludedUserFilterEnabled bool `json:"IncludedUserFilterEnabled"`
	// The included users filter of the desktop; that is, the users and groups who are explicitly granted access to the published desktop.
	IncludedUsers []IdentityUserResponseModel `json:"IncludedUsers,omitempty"`
	LeasingBehavior *LeasingBehavior `json:"LeasingBehavior,omitempty"`
	// The number of machines from the delivery group which a user may privately allocate.
	MaxDesktops NullableInt32 `json:"MaxDesktops,omitempty"`
	// The administrative name of the desktop.
	Name NullableString `json:"Name,omitempty"`
	// The name of the published desktop as seen by the user, and of assigned desktop(s) in the case where SharingKind is equal to Private.
	PublishedName string `json:"PublishedName"`
	RestrictToTag *RefResponseModel `json:"RestrictToTag,omitempty"`
	// Indicates whether the desktop requires the SecureICA protocol for desktop sessions.
	SecureIcaRequired bool `json:"SecureIcaRequired"`
	SessionReconnection *SessionReconnection `json:"SessionReconnection,omitempty"`
	// Indicates the number of machines which are available for assignment based on this desktop configuration.
	MachinesForAssignment NullableInt32 `json:"MachinesForAssignment,omitempty"`
	// The tenant(s) that the desktop is assigned to.  If `null`, the desktop is not assigned to any tenants, and may be used by any tenant.
	Tenants []RefResponseModel `json:"Tenants,omitempty"`
}

// NewDesktopResponseModel instantiates a new DesktopResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopResponseModel(id string, colorDepth ColorDepth, enabled bool, excludedUserFilterEnabled bool, iconId string, includedUserFilterEnabled bool, publishedName string, secureIcaRequired bool) *DesktopResponseModel {
	this := DesktopResponseModel{}
	this.Id = id
	this.ColorDepth = colorDepth
	this.Enabled = enabled
	this.ExcludedUserFilterEnabled = excludedUserFilterEnabled
	this.IconId = iconId
	this.IncludedUserFilterEnabled = includedUserFilterEnabled
	this.PublishedName = publishedName
	this.SecureIcaRequired = secureIcaRequired
	return &this
}

// NewDesktopResponseModelWithDefaults instantiates a new DesktopResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopResponseModelWithDefaults() *DesktopResponseModel {
	this := DesktopResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *DesktopResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DesktopResponseModel) SetId(v string) {
	o.Id = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
// Deprecated
func (o *DesktopResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DesktopResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
// Deprecated
func (o *DesktopResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetColorDepth returns the ColorDepth field value
func (o *DesktopResponseModel) GetColorDepth() ColorDepth {
	if o == nil {
		var ret ColorDepth
		return ret
	}

	return o.ColorDepth
}

// GetColorDepthOk returns a tuple with the ColorDepth field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetColorDepthOk() (*ColorDepth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColorDepth, true
}

// SetColorDepth sets field value
func (o *DesktopResponseModel) SetColorDepth(v ColorDepth) {
	o.ColorDepth = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DesktopResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DesktopResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DesktopResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value
func (o *DesktopResponseModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DesktopResponseModel) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field value
func (o *DesktopResponseModel) GetExcludedUserFilterEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExcludedUserFilterEnabled
}

// GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetExcludedUserFilterEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludedUserFilterEnabled, true
}

// SetExcludedUserFilterEnabled sets field value
func (o *DesktopResponseModel) SetExcludedUserFilterEnabled(v bool) {
	o.ExcludedUserFilterEnabled = v
}

// GetExcludedUsers returns the ExcludedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetExcludedUsers() []IdentityUserResponseModel {
	if o == nil {
		var ret []IdentityUserResponseModel
		return ret
	}
	return o.ExcludedUsers
}

// GetExcludedUsersOk returns a tuple with the ExcludedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetExcludedUsersOk() ([]IdentityUserResponseModel, bool) {
	if o == nil || IsNil(o.ExcludedUsers) {
		return nil, false
	}
	return o.ExcludedUsers, true
}

// HasExcludedUsers returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasExcludedUsers() bool {
	if o != nil && IsNil(o.ExcludedUsers) {
		return true
	}

	return false
}

// SetExcludedUsers gets a reference to the given []IdentityUserResponseModel and assigns it to the ExcludedUsers field.
func (o *DesktopResponseModel) SetExcludedUsers(v []IdentityUserResponseModel) {
	o.ExcludedUsers = v
}

// GetIconId returns the IconId field value
func (o *DesktopResponseModel) GetIconId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetIconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconId, true
}

// SetIconId sets field value
func (o *DesktopResponseModel) SetIconId(v string) {
	o.IconId = v
}

// GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field value
func (o *DesktopResponseModel) GetIncludedUserFilterEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludedUserFilterEnabled
}

// GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludedUserFilterEnabled, true
}

// SetIncludedUserFilterEnabled sets field value
func (o *DesktopResponseModel) SetIncludedUserFilterEnabled(v bool) {
	o.IncludedUserFilterEnabled = v
}

// GetIncludedUsers returns the IncludedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetIncludedUsers() []IdentityUserResponseModel {
	if o == nil {
		var ret []IdentityUserResponseModel
		return ret
	}
	return o.IncludedUsers
}

// GetIncludedUsersOk returns a tuple with the IncludedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetIncludedUsersOk() ([]IdentityUserResponseModel, bool) {
	if o == nil || IsNil(o.IncludedUsers) {
		return nil, false
	}
	return o.IncludedUsers, true
}

// HasIncludedUsers returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasIncludedUsers() bool {
	if o != nil && IsNil(o.IncludedUsers) {
		return true
	}

	return false
}

// SetIncludedUsers gets a reference to the given []IdentityUserResponseModel and assigns it to the IncludedUsers field.
func (o *DesktopResponseModel) SetIncludedUsers(v []IdentityUserResponseModel) {
	o.IncludedUsers = v
}

// GetLeasingBehavior returns the LeasingBehavior field value if set, zero value otherwise.
func (o *DesktopResponseModel) GetLeasingBehavior() LeasingBehavior {
	if o == nil || IsNil(o.LeasingBehavior) {
		var ret LeasingBehavior
		return ret
	}
	return *o.LeasingBehavior
}

// GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetLeasingBehaviorOk() (*LeasingBehavior, bool) {
	if o == nil || IsNil(o.LeasingBehavior) {
		return nil, false
	}
	return o.LeasingBehavior, true
}

// HasLeasingBehavior returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasLeasingBehavior() bool {
	if o != nil && !IsNil(o.LeasingBehavior) {
		return true
	}

	return false
}

// SetLeasingBehavior gets a reference to the given LeasingBehavior and assigns it to the LeasingBehavior field.
func (o *DesktopResponseModel) SetLeasingBehavior(v LeasingBehavior) {
	o.LeasingBehavior = &v
}

// GetMaxDesktops returns the MaxDesktops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetMaxDesktops() int32 {
	if o == nil || IsNil(o.MaxDesktops.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxDesktops.Get()
}

// GetMaxDesktopsOk returns a tuple with the MaxDesktops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetMaxDesktopsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDesktops.Get(), o.MaxDesktops.IsSet()
}

// HasMaxDesktops returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasMaxDesktops() bool {
	if o != nil && o.MaxDesktops.IsSet() {
		return true
	}

	return false
}

// SetMaxDesktops gets a reference to the given NullableInt32 and assigns it to the MaxDesktops field.
func (o *DesktopResponseModel) SetMaxDesktops(v int32) {
	o.MaxDesktops.Set(&v)
}
// SetMaxDesktopsNil sets the value for MaxDesktops to be an explicit nil
func (o *DesktopResponseModel) SetMaxDesktopsNil() {
	o.MaxDesktops.Set(nil)
}

// UnsetMaxDesktops ensures that no value is present for MaxDesktops, not even an explicit nil
func (o *DesktopResponseModel) UnsetMaxDesktops() {
	o.MaxDesktops.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DesktopResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DesktopResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DesktopResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetPublishedName returns the PublishedName field value
func (o *DesktopResponseModel) GetPublishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedName
}

// GetPublishedNameOk returns a tuple with the PublishedName field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetPublishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedName, true
}

// SetPublishedName sets field value
func (o *DesktopResponseModel) SetPublishedName(v string) {
	o.PublishedName = v
}

// GetRestrictToTag returns the RestrictToTag field value if set, zero value otherwise.
func (o *DesktopResponseModel) GetRestrictToTag() RefResponseModel {
	if o == nil || IsNil(o.RestrictToTag) {
		var ret RefResponseModel
		return ret
	}
	return *o.RestrictToTag
}

// GetRestrictToTagOk returns a tuple with the RestrictToTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetRestrictToTagOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.RestrictToTag) {
		return nil, false
	}
	return o.RestrictToTag, true
}

// HasRestrictToTag returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasRestrictToTag() bool {
	if o != nil && !IsNil(o.RestrictToTag) {
		return true
	}

	return false
}

// SetRestrictToTag gets a reference to the given RefResponseModel and assigns it to the RestrictToTag field.
func (o *DesktopResponseModel) SetRestrictToTag(v RefResponseModel) {
	o.RestrictToTag = &v
}

// GetSecureIcaRequired returns the SecureIcaRequired field value
func (o *DesktopResponseModel) GetSecureIcaRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SecureIcaRequired
}

// GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field value
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetSecureIcaRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureIcaRequired, true
}

// SetSecureIcaRequired sets field value
func (o *DesktopResponseModel) SetSecureIcaRequired(v bool) {
	o.SecureIcaRequired = v
}

// GetSessionReconnection returns the SessionReconnection field value if set, zero value otherwise.
func (o *DesktopResponseModel) GetSessionReconnection() SessionReconnection {
	if o == nil || IsNil(o.SessionReconnection) {
		var ret SessionReconnection
		return ret
	}
	return *o.SessionReconnection
}

// GetSessionReconnectionOk returns a tuple with the SessionReconnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopResponseModel) GetSessionReconnectionOk() (*SessionReconnection, bool) {
	if o == nil || IsNil(o.SessionReconnection) {
		return nil, false
	}
	return o.SessionReconnection, true
}

// HasSessionReconnection returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasSessionReconnection() bool {
	if o != nil && !IsNil(o.SessionReconnection) {
		return true
	}

	return false
}

// SetSessionReconnection gets a reference to the given SessionReconnection and assigns it to the SessionReconnection field.
func (o *DesktopResponseModel) SetSessionReconnection(v SessionReconnection) {
	o.SessionReconnection = &v
}

// GetMachinesForAssignment returns the MachinesForAssignment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetMachinesForAssignment() int32 {
	if o == nil || IsNil(o.MachinesForAssignment.Get()) {
		var ret int32
		return ret
	}
	return *o.MachinesForAssignment.Get()
}

// GetMachinesForAssignmentOk returns a tuple with the MachinesForAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetMachinesForAssignmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachinesForAssignment.Get(), o.MachinesForAssignment.IsSet()
}

// HasMachinesForAssignment returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasMachinesForAssignment() bool {
	if o != nil && o.MachinesForAssignment.IsSet() {
		return true
	}

	return false
}

// SetMachinesForAssignment gets a reference to the given NullableInt32 and assigns it to the MachinesForAssignment field.
func (o *DesktopResponseModel) SetMachinesForAssignment(v int32) {
	o.MachinesForAssignment.Set(&v)
}
// SetMachinesForAssignmentNil sets the value for MachinesForAssignment to be an explicit nil
func (o *DesktopResponseModel) SetMachinesForAssignmentNil() {
	o.MachinesForAssignment.Set(nil)
}

// UnsetMachinesForAssignment ensures that no value is present for MachinesForAssignment, not even an explicit nil
func (o *DesktopResponseModel) UnsetMachinesForAssignment() {
	o.MachinesForAssignment.Unset()
}

// GetTenants returns the Tenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopResponseModel) GetTenants() []RefResponseModel {
	if o == nil {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *DesktopResponseModel) HasTenants() bool {
	if o != nil && IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *DesktopResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

func (o DesktopResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesktopResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	toSerialize["ColorDepth"] = o.ColorDepth
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	toSerialize["Enabled"] = o.Enabled
	toSerialize["ExcludedUserFilterEnabled"] = o.ExcludedUserFilterEnabled
	if o.ExcludedUsers != nil {
		toSerialize["ExcludedUsers"] = o.ExcludedUsers
	}
	toSerialize["IconId"] = o.IconId
	toSerialize["IncludedUserFilterEnabled"] = o.IncludedUserFilterEnabled
	if o.IncludedUsers != nil {
		toSerialize["IncludedUsers"] = o.IncludedUsers
	}
	if !IsNil(o.LeasingBehavior) {
		toSerialize["LeasingBehavior"] = o.LeasingBehavior
	}
	if o.MaxDesktops.IsSet() {
		toSerialize["MaxDesktops"] = o.MaxDesktops.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	toSerialize["PublishedName"] = o.PublishedName
	if !IsNil(o.RestrictToTag) {
		toSerialize["RestrictToTag"] = o.RestrictToTag
	}
	toSerialize["SecureIcaRequired"] = o.SecureIcaRequired
	if !IsNil(o.SessionReconnection) {
		toSerialize["SessionReconnection"] = o.SessionReconnection
	}
	if o.MachinesForAssignment.IsSet() {
		toSerialize["MachinesForAssignment"] = o.MachinesForAssignment.Get()
	}
	if o.Tenants != nil {
		toSerialize["Tenants"] = o.Tenants
	}
	return toSerialize, nil
}

type NullableDesktopResponseModel struct {
	value *DesktopResponseModel
	isSet bool
}

func (v NullableDesktopResponseModel) Get() *DesktopResponseModel {
	return v.value
}

func (v *NullableDesktopResponseModel) Set(val *DesktopResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopResponseModel(val *DesktopResponseModel) *NullableDesktopResponseModel {
	return &NullableDesktopResponseModel{value: val, isSet: true}
}

func (v NullableDesktopResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


