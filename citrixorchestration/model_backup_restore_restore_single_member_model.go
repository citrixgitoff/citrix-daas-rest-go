/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"time"
)

// checks if the BackupRestoreRestoreSingleMemberModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreRestoreSingleMemberModel{}

// BackupRestoreRestoreSingleMemberModel struct for BackupRestoreRestoreSingleMemberModel
type BackupRestoreRestoreSingleMemberModel struct {
	// Component Name
	ComponentName NullableString `json:"ComponentName,omitempty"`
	// Member Name
	MemberName NullableString `json:"MemberName,omitempty"`
	// Member Type (Host Connections Only)
	MemberType NullableString `json:"MemberType,omitempty"`
	PlaybookActionState *BackupRestorePlaybookActionState `json:"PlaybookActionState,omitempty"`
	// Result
	Result *bool `json:"Result,omitempty"`
	// Details or Error Message
	Details NullableString `json:"Details,omitempty"`
	// Transaction Id
	TransactionId NullableString `json:"TransactionId,omitempty"`
	// Date Time stamps
	DateTime *time.Time `json:"DateTime,omitempty"`
	// Member Count
	MemberCount *int32 `json:"MemberCount,omitempty"`
	CurrentMemberIndex *int32 `json:"CurrentMemberIndex,omitempty"`
	// Json string containing ACT internal Fixups data. Must be exchanged with ACT to obtain human readable fixups.
	Fixups NullableString `json:"Fixups,omitempty"`
	// Check Mode
	CheckMode *bool `json:"CheckMode,omitempty"`
}

// NewBackupRestoreRestoreSingleMemberModel instantiates a new BackupRestoreRestoreSingleMemberModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreRestoreSingleMemberModel() *BackupRestoreRestoreSingleMemberModel {
	this := BackupRestoreRestoreSingleMemberModel{}
	return &this
}

// NewBackupRestoreRestoreSingleMemberModelWithDefaults instantiates a new BackupRestoreRestoreSingleMemberModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreRestoreSingleMemberModelWithDefaults() *BackupRestoreRestoreSingleMemberModel {
	this := BackupRestoreRestoreSingleMemberModel{}
	return &this
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetComponentName() string {
	if o == nil || IsNil(o.ComponentName.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentName.Get()
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentName.Get(), o.ComponentName.IsSet()
}

// HasComponentName returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasComponentName() bool {
	if o != nil && o.ComponentName.IsSet() {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given NullableString and assigns it to the ComponentName field.
func (o *BackupRestoreRestoreSingleMemberModel) SetComponentName(v string) {
	o.ComponentName.Set(&v)
}
// SetComponentNameNil sets the value for ComponentName to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetComponentNameNil() {
	o.ComponentName.Set(nil)
}

// UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetComponentName() {
	o.ComponentName.Unset()
}

// GetMemberName returns the MemberName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberName() string {
	if o == nil || IsNil(o.MemberName.Get()) {
		var ret string
		return ret
	}
	return *o.MemberName.Get()
}

// GetMemberNameOk returns a tuple with the MemberName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberName.Get(), o.MemberName.IsSet()
}

// HasMemberName returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasMemberName() bool {
	if o != nil && o.MemberName.IsSet() {
		return true
	}

	return false
}

// SetMemberName gets a reference to the given NullableString and assigns it to the MemberName field.
func (o *BackupRestoreRestoreSingleMemberModel) SetMemberName(v string) {
	o.MemberName.Set(&v)
}
// SetMemberNameNil sets the value for MemberName to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetMemberNameNil() {
	o.MemberName.Set(nil)
}

// UnsetMemberName ensures that no value is present for MemberName, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetMemberName() {
	o.MemberName.Unset()
}

// GetMemberType returns the MemberType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberType() string {
	if o == nil || IsNil(o.MemberType.Get()) {
		var ret string
		return ret
	}
	return *o.MemberType.Get()
}

// GetMemberTypeOk returns a tuple with the MemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberType.Get(), o.MemberType.IsSet()
}

// HasMemberType returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasMemberType() bool {
	if o != nil && o.MemberType.IsSet() {
		return true
	}

	return false
}

// SetMemberType gets a reference to the given NullableString and assigns it to the MemberType field.
func (o *BackupRestoreRestoreSingleMemberModel) SetMemberType(v string) {
	o.MemberType.Set(&v)
}
// SetMemberTypeNil sets the value for MemberType to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetMemberTypeNil() {
	o.MemberType.Set(nil)
}

// UnsetMemberType ensures that no value is present for MemberType, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetMemberType() {
	o.MemberType.Unset()
}

// GetPlaybookActionState returns the PlaybookActionState field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetPlaybookActionState() BackupRestorePlaybookActionState {
	if o == nil || IsNil(o.PlaybookActionState) {
		var ret BackupRestorePlaybookActionState
		return ret
	}
	return *o.PlaybookActionState
}

// GetPlaybookActionStateOk returns a tuple with the PlaybookActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetPlaybookActionStateOk() (*BackupRestorePlaybookActionState, bool) {
	if o == nil || IsNil(o.PlaybookActionState) {
		return nil, false
	}
	return o.PlaybookActionState, true
}

// HasPlaybookActionState returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasPlaybookActionState() bool {
	if o != nil && !IsNil(o.PlaybookActionState) {
		return true
	}

	return false
}

// SetPlaybookActionState gets a reference to the given BackupRestorePlaybookActionState and assigns it to the PlaybookActionState field.
func (o *BackupRestoreRestoreSingleMemberModel) SetPlaybookActionState(v BackupRestorePlaybookActionState) {
	o.PlaybookActionState = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *BackupRestoreRestoreSingleMemberModel) SetResult(v bool) {
	o.Result = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *BackupRestoreRestoreSingleMemberModel) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetDetails() {
	o.Details.Unset()
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *BackupRestoreRestoreSingleMemberModel) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *BackupRestoreRestoreSingleMemberModel) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberCount() int32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *BackupRestoreRestoreSingleMemberModel) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetCurrentMemberIndex returns the CurrentMemberIndex field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetCurrentMemberIndex() int32 {
	if o == nil || IsNil(o.CurrentMemberIndex) {
		var ret int32
		return ret
	}
	return *o.CurrentMemberIndex
}

// GetCurrentMemberIndexOk returns a tuple with the CurrentMemberIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetCurrentMemberIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentMemberIndex) {
		return nil, false
	}
	return o.CurrentMemberIndex, true
}

// HasCurrentMemberIndex returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasCurrentMemberIndex() bool {
	if o != nil && !IsNil(o.CurrentMemberIndex) {
		return true
	}

	return false
}

// SetCurrentMemberIndex gets a reference to the given int32 and assigns it to the CurrentMemberIndex field.
func (o *BackupRestoreRestoreSingleMemberModel) SetCurrentMemberIndex(v int32) {
	o.CurrentMemberIndex = &v
}

// GetFixups returns the Fixups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreRestoreSingleMemberModel) GetFixups() string {
	if o == nil || IsNil(o.Fixups.Get()) {
		var ret string
		return ret
	}
	return *o.Fixups.Get()
}

// GetFixupsOk returns a tuple with the Fixups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreRestoreSingleMemberModel) GetFixupsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fixups.Get(), o.Fixups.IsSet()
}

// HasFixups returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasFixups() bool {
	if o != nil && o.Fixups.IsSet() {
		return true
	}

	return false
}

// SetFixups gets a reference to the given NullableString and assigns it to the Fixups field.
func (o *BackupRestoreRestoreSingleMemberModel) SetFixups(v string) {
	o.Fixups.Set(&v)
}
// SetFixupsNil sets the value for Fixups to be an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) SetFixupsNil() {
	o.Fixups.Set(nil)
}

// UnsetFixups ensures that no value is present for Fixups, not even an explicit nil
func (o *BackupRestoreRestoreSingleMemberModel) UnsetFixups() {
	o.Fixups.Unset()
}

// GetCheckMode returns the CheckMode field value if set, zero value otherwise.
func (o *BackupRestoreRestoreSingleMemberModel) GetCheckMode() bool {
	if o == nil || IsNil(o.CheckMode) {
		var ret bool
		return ret
	}
	return *o.CheckMode
}

// GetCheckModeOk returns a tuple with the CheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreRestoreSingleMemberModel) GetCheckModeOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckMode) {
		return nil, false
	}
	return o.CheckMode, true
}

// HasCheckMode returns a boolean if a field has been set.
func (o *BackupRestoreRestoreSingleMemberModel) HasCheckMode() bool {
	if o != nil && !IsNil(o.CheckMode) {
		return true
	}

	return false
}

// SetCheckMode gets a reference to the given bool and assigns it to the CheckMode field.
func (o *BackupRestoreRestoreSingleMemberModel) SetCheckMode(v bool) {
	o.CheckMode = &v
}

func (o BackupRestoreRestoreSingleMemberModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreRestoreSingleMemberModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ComponentName.IsSet() {
		toSerialize["ComponentName"] = o.ComponentName.Get()
	}
	if o.MemberName.IsSet() {
		toSerialize["MemberName"] = o.MemberName.Get()
	}
	if o.MemberType.IsSet() {
		toSerialize["MemberType"] = o.MemberType.Get()
	}
	if !IsNil(o.PlaybookActionState) {
		toSerialize["PlaybookActionState"] = o.PlaybookActionState
	}
	if !IsNil(o.Result) {
		toSerialize["Result"] = o.Result
	}
	if o.Details.IsSet() {
		toSerialize["Details"] = o.Details.Get()
	}
	if o.TransactionId.IsSet() {
		toSerialize["TransactionId"] = o.TransactionId.Get()
	}
	if !IsNil(o.DateTime) {
		toSerialize["DateTime"] = o.DateTime
	}
	if !IsNil(o.MemberCount) {
		toSerialize["MemberCount"] = o.MemberCount
	}
	if !IsNil(o.CurrentMemberIndex) {
		toSerialize["CurrentMemberIndex"] = o.CurrentMemberIndex
	}
	if o.Fixups.IsSet() {
		toSerialize["Fixups"] = o.Fixups.Get()
	}
	if !IsNil(o.CheckMode) {
		toSerialize["CheckMode"] = o.CheckMode
	}
	return toSerialize, nil
}

type NullableBackupRestoreRestoreSingleMemberModel struct {
	value *BackupRestoreRestoreSingleMemberModel
	isSet bool
}

func (v NullableBackupRestoreRestoreSingleMemberModel) Get() *BackupRestoreRestoreSingleMemberModel {
	return v.value
}

func (v *NullableBackupRestoreRestoreSingleMemberModel) Set(val *BackupRestoreRestoreSingleMemberModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreRestoreSingleMemberModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreRestoreSingleMemberModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreRestoreSingleMemberModel(val *BackupRestoreRestoreSingleMemberModel) *NullableBackupRestoreRestoreSingleMemberModel {
	return &NullableBackupRestoreRestoreSingleMemberModel{value: val, isSet: true}
}

func (v NullableBackupRestoreRestoreSingleMemberModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreRestoreSingleMemberModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


