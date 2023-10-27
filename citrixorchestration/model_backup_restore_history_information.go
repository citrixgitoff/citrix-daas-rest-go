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

// checks if the BackupRestoreHistoryInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreHistoryInformation{}

// BackupRestoreHistoryInformation Single history for backup / restore operations
type BackupRestoreHistoryInformation struct {
	// Unique Identifier             
	Uid *int32 `json:"Uid,omitempty"`
	Action BackupRestoreActions `json:"Action"`
	// History entry ID
	HistoryId string `json:"HistoryId"`
	// True if successful, false if not successful
	Successful bool `json:"Successful"`
	// Duration in seconds
	Duration *int32 `json:"Duration,omitempty"`
	// TimeZone offset from UTC
	TimeZoneOffset *int32 `json:"TimeZoneOffset,omitempty"`
	// Date and time the action was started
	DateTime *time.Time `json:"DateTime,omitempty"`
	// Notes about the action
	Notes NullableString `json:"Notes,omitempty"`
	RestoreType *BackupRestoreRestoreTypes `json:"RestoreType,omitempty"`
	// Filters used (applicable only when the action was restore)
	Filters NullableString `json:"Filters,omitempty"`
	// With Prerequisites (applicable only when the action was restore)
	WithPrerequisites *bool `json:"WithPrerequisites,omitempty"`
	// Check mode (applicable only when the action was restore)
	CheckMode *bool `json:"CheckMode,omitempty"`
	// Name of schedule to perform backup (applicable only when the action was backup)
	ScheduleName NullableString `json:"ScheduleName,omitempty"`
	Component *BckRstrAutoConfigComponents `json:"Component,omitempty"`
	// Execution Id
	ExecutionId NullableString `json:"ExecutionId,omitempty"`
	// Backup name
	BackupName NullableString `json:"BackupName,omitempty"`
	// Backup size
	BackupSize *int64 `json:"BackupSize,omitempty"`
	// Backup File Specification
	BackupFileSpec NullableString `json:"BackupFileSpec,omitempty"`
	// Related History UID for restore with checkmode set to true
	RelatedUid *int32 `json:"RelatedUid,omitempty"`
	// Related History date for restore with checkmode set to true
	RelatedDate *time.Time `json:"RelatedDate,omitempty"`
	// Related is run as check mode
	RelatedIsCheckMode *bool `json:"RelatedIsCheckMode,omitempty"`
	// Backup is pinned
	Pinned *bool `json:"Pinned,omitempty"`
	// Administrator Name
	AdministratorName NullableString `json:"AdministratorName,omitempty"`
	// Backup Status Details
	BackupDetails map[string]string `json:"BackupDetails,omitempty"`
	// Restore Status Details
	RestoreDetails []BackupRestoreRestoreSingleMemberModel `json:"RestoreDetails,omitempty"`
	// Simple Results (such as Get backed up member names)
	SimpleResults []string `json:"SimpleResults,omitempty"`
	// Fixups 
	Fixups []BackupRestoreFixupModel `json:"Fixups,omitempty"`
}

// NewBackupRestoreHistoryInformation instantiates a new BackupRestoreHistoryInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreHistoryInformation(action BackupRestoreActions, historyId string, successful bool) *BackupRestoreHistoryInformation {
	this := BackupRestoreHistoryInformation{}
	this.Action = action
	this.HistoryId = historyId
	this.Successful = successful
	return &this
}

// NewBackupRestoreHistoryInformationWithDefaults instantiates a new BackupRestoreHistoryInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreHistoryInformationWithDefaults() *BackupRestoreHistoryInformation {
	this := BackupRestoreHistoryInformation{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *BackupRestoreHistoryInformation) SetUid(v int32) {
	o.Uid = &v
}

// GetAction returns the Action field value
func (o *BackupRestoreHistoryInformation) GetAction() BackupRestoreActions {
	if o == nil {
		var ret BackupRestoreActions
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetActionOk() (*BackupRestoreActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BackupRestoreHistoryInformation) SetAction(v BackupRestoreActions) {
	o.Action = v
}

// GetHistoryId returns the HistoryId field value
func (o *BackupRestoreHistoryInformation) GetHistoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HistoryId
}

// GetHistoryIdOk returns a tuple with the HistoryId field value
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetHistoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoryId, true
}

// SetHistoryId sets field value
func (o *BackupRestoreHistoryInformation) SetHistoryId(v string) {
	o.HistoryId = v
}

// GetSuccessful returns the Successful field value
func (o *BackupRestoreHistoryInformation) GetSuccessful() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Successful, true
}

// SetSuccessful sets field value
func (o *BackupRestoreHistoryInformation) SetSuccessful(v bool) {
	o.Successful = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *BackupRestoreHistoryInformation) SetDuration(v int32) {
	o.Duration = &v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetTimeZoneOffset() int32 {
	if o == nil || IsNil(o.TimeZoneOffset) {
		var ret int32
		return ret
	}
	return *o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetTimeZoneOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeZoneOffset) {
		return nil, false
	}
	return o.TimeZoneOffset, true
}

// HasTimeZoneOffset returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasTimeZoneOffset() bool {
	if o != nil && !IsNil(o.TimeZoneOffset) {
		return true
	}

	return false
}

// SetTimeZoneOffset gets a reference to the given int32 and assigns it to the TimeZoneOffset field.
func (o *BackupRestoreHistoryInformation) SetTimeZoneOffset(v int32) {
	o.TimeZoneOffset = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *BackupRestoreHistoryInformation) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *BackupRestoreHistoryInformation) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetNotes() {
	o.Notes.Unset()
}

// GetRestoreType returns the RestoreType field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetRestoreType() BackupRestoreRestoreTypes {
	if o == nil || IsNil(o.RestoreType) {
		var ret BackupRestoreRestoreTypes
		return ret
	}
	return *o.RestoreType
}

// GetRestoreTypeOk returns a tuple with the RestoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetRestoreTypeOk() (*BackupRestoreRestoreTypes, bool) {
	if o == nil || IsNil(o.RestoreType) {
		return nil, false
	}
	return o.RestoreType, true
}

// HasRestoreType returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasRestoreType() bool {
	if o != nil && !IsNil(o.RestoreType) {
		return true
	}

	return false
}

// SetRestoreType gets a reference to the given BackupRestoreRestoreTypes and assigns it to the RestoreType field.
func (o *BackupRestoreHistoryInformation) SetRestoreType(v BackupRestoreRestoreTypes) {
	o.RestoreType = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetFilters() string {
	if o == nil || IsNil(o.Filters.Get()) {
		var ret string
		return ret
	}
	return *o.Filters.Get()
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetFiltersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters.Get(), o.Filters.IsSet()
}

// HasFilters returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasFilters() bool {
	if o != nil && o.Filters.IsSet() {
		return true
	}

	return false
}

// SetFilters gets a reference to the given NullableString and assigns it to the Filters field.
func (o *BackupRestoreHistoryInformation) SetFilters(v string) {
	o.Filters.Set(&v)
}
// SetFiltersNil sets the value for Filters to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetFiltersNil() {
	o.Filters.Set(nil)
}

// UnsetFilters ensures that no value is present for Filters, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetFilters() {
	o.Filters.Unset()
}

// GetWithPrerequisites returns the WithPrerequisites field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetWithPrerequisites() bool {
	if o == nil || IsNil(o.WithPrerequisites) {
		var ret bool
		return ret
	}
	return *o.WithPrerequisites
}

// GetWithPrerequisitesOk returns a tuple with the WithPrerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetWithPrerequisitesOk() (*bool, bool) {
	if o == nil || IsNil(o.WithPrerequisites) {
		return nil, false
	}
	return o.WithPrerequisites, true
}

// HasWithPrerequisites returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasWithPrerequisites() bool {
	if o != nil && !IsNil(o.WithPrerequisites) {
		return true
	}

	return false
}

// SetWithPrerequisites gets a reference to the given bool and assigns it to the WithPrerequisites field.
func (o *BackupRestoreHistoryInformation) SetWithPrerequisites(v bool) {
	o.WithPrerequisites = &v
}

// GetCheckMode returns the CheckMode field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetCheckMode() bool {
	if o == nil || IsNil(o.CheckMode) {
		var ret bool
		return ret
	}
	return *o.CheckMode
}

// GetCheckModeOk returns a tuple with the CheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetCheckModeOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckMode) {
		return nil, false
	}
	return o.CheckMode, true
}

// HasCheckMode returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasCheckMode() bool {
	if o != nil && !IsNil(o.CheckMode) {
		return true
	}

	return false
}

// SetCheckMode gets a reference to the given bool and assigns it to the CheckMode field.
func (o *BackupRestoreHistoryInformation) SetCheckMode(v bool) {
	o.CheckMode = &v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetScheduleName() string {
	if o == nil || IsNil(o.ScheduleName.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduleName.Get()
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetScheduleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleName.Get(), o.ScheduleName.IsSet()
}

// HasScheduleName returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasScheduleName() bool {
	if o != nil && o.ScheduleName.IsSet() {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given NullableString and assigns it to the ScheduleName field.
func (o *BackupRestoreHistoryInformation) SetScheduleName(v string) {
	o.ScheduleName.Set(&v)
}
// SetScheduleNameNil sets the value for ScheduleName to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetScheduleNameNil() {
	o.ScheduleName.Set(nil)
}

// UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetScheduleName() {
	o.ScheduleName.Unset()
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetComponent() BckRstrAutoConfigComponents {
	if o == nil || IsNil(o.Component) {
		var ret BckRstrAutoConfigComponents
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetComponentOk() (*BckRstrAutoConfigComponents, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given BckRstrAutoConfigComponents and assigns it to the Component field.
func (o *BackupRestoreHistoryInformation) SetComponent(v BckRstrAutoConfigComponents) {
	o.Component = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *BackupRestoreHistoryInformation) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetBackupName returns the BackupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetBackupName() string {
	if o == nil || IsNil(o.BackupName.Get()) {
		var ret string
		return ret
	}
	return *o.BackupName.Get()
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetBackupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupName.Get(), o.BackupName.IsSet()
}

// HasBackupName returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasBackupName() bool {
	if o != nil && o.BackupName.IsSet() {
		return true
	}

	return false
}

// SetBackupName gets a reference to the given NullableString and assigns it to the BackupName field.
func (o *BackupRestoreHistoryInformation) SetBackupName(v string) {
	o.BackupName.Set(&v)
}
// SetBackupNameNil sets the value for BackupName to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetBackupNameNil() {
	o.BackupName.Set(nil)
}

// UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetBackupName() {
	o.BackupName.Unset()
}

// GetBackupSize returns the BackupSize field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetBackupSize() int64 {
	if o == nil || IsNil(o.BackupSize) {
		var ret int64
		return ret
	}
	return *o.BackupSize
}

// GetBackupSizeOk returns a tuple with the BackupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetBackupSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.BackupSize) {
		return nil, false
	}
	return o.BackupSize, true
}

// HasBackupSize returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasBackupSize() bool {
	if o != nil && !IsNil(o.BackupSize) {
		return true
	}

	return false
}

// SetBackupSize gets a reference to the given int64 and assigns it to the BackupSize field.
func (o *BackupRestoreHistoryInformation) SetBackupSize(v int64) {
	o.BackupSize = &v
}

// GetBackupFileSpec returns the BackupFileSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetBackupFileSpec() string {
	if o == nil || IsNil(o.BackupFileSpec.Get()) {
		var ret string
		return ret
	}
	return *o.BackupFileSpec.Get()
}

// GetBackupFileSpecOk returns a tuple with the BackupFileSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetBackupFileSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupFileSpec.Get(), o.BackupFileSpec.IsSet()
}

// HasBackupFileSpec returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasBackupFileSpec() bool {
	if o != nil && o.BackupFileSpec.IsSet() {
		return true
	}

	return false
}

// SetBackupFileSpec gets a reference to the given NullableString and assigns it to the BackupFileSpec field.
func (o *BackupRestoreHistoryInformation) SetBackupFileSpec(v string) {
	o.BackupFileSpec.Set(&v)
}
// SetBackupFileSpecNil sets the value for BackupFileSpec to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetBackupFileSpecNil() {
	o.BackupFileSpec.Set(nil)
}

// UnsetBackupFileSpec ensures that no value is present for BackupFileSpec, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetBackupFileSpec() {
	o.BackupFileSpec.Unset()
}

// GetRelatedUid returns the RelatedUid field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetRelatedUid() int32 {
	if o == nil || IsNil(o.RelatedUid) {
		var ret int32
		return ret
	}
	return *o.RelatedUid
}

// GetRelatedUidOk returns a tuple with the RelatedUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetRelatedUidOk() (*int32, bool) {
	if o == nil || IsNil(o.RelatedUid) {
		return nil, false
	}
	return o.RelatedUid, true
}

// HasRelatedUid returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasRelatedUid() bool {
	if o != nil && !IsNil(o.RelatedUid) {
		return true
	}

	return false
}

// SetRelatedUid gets a reference to the given int32 and assigns it to the RelatedUid field.
func (o *BackupRestoreHistoryInformation) SetRelatedUid(v int32) {
	o.RelatedUid = &v
}

// GetRelatedDate returns the RelatedDate field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetRelatedDate() time.Time {
	if o == nil || IsNil(o.RelatedDate) {
		var ret time.Time
		return ret
	}
	return *o.RelatedDate
}

// GetRelatedDateOk returns a tuple with the RelatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetRelatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RelatedDate) {
		return nil, false
	}
	return o.RelatedDate, true
}

// HasRelatedDate returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasRelatedDate() bool {
	if o != nil && !IsNil(o.RelatedDate) {
		return true
	}

	return false
}

// SetRelatedDate gets a reference to the given time.Time and assigns it to the RelatedDate field.
func (o *BackupRestoreHistoryInformation) SetRelatedDate(v time.Time) {
	o.RelatedDate = &v
}

// GetRelatedIsCheckMode returns the RelatedIsCheckMode field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetRelatedIsCheckMode() bool {
	if o == nil || IsNil(o.RelatedIsCheckMode) {
		var ret bool
		return ret
	}
	return *o.RelatedIsCheckMode
}

// GetRelatedIsCheckModeOk returns a tuple with the RelatedIsCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetRelatedIsCheckModeOk() (*bool, bool) {
	if o == nil || IsNil(o.RelatedIsCheckMode) {
		return nil, false
	}
	return o.RelatedIsCheckMode, true
}

// HasRelatedIsCheckMode returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasRelatedIsCheckMode() bool {
	if o != nil && !IsNil(o.RelatedIsCheckMode) {
		return true
	}

	return false
}

// SetRelatedIsCheckMode gets a reference to the given bool and assigns it to the RelatedIsCheckMode field.
func (o *BackupRestoreHistoryInformation) SetRelatedIsCheckMode(v bool) {
	o.RelatedIsCheckMode = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *BackupRestoreHistoryInformation) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreHistoryInformation) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasPinned() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *BackupRestoreHistoryInformation) SetPinned(v bool) {
	o.Pinned = &v
}

// GetAdministratorName returns the AdministratorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetAdministratorName() string {
	if o == nil || IsNil(o.AdministratorName.Get()) {
		var ret string
		return ret
	}
	return *o.AdministratorName.Get()
}

// GetAdministratorNameOk returns a tuple with the AdministratorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetAdministratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdministratorName.Get(), o.AdministratorName.IsSet()
}

// HasAdministratorName returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasAdministratorName() bool {
	if o != nil && o.AdministratorName.IsSet() {
		return true
	}

	return false
}

// SetAdministratorName gets a reference to the given NullableString and assigns it to the AdministratorName field.
func (o *BackupRestoreHistoryInformation) SetAdministratorName(v string) {
	o.AdministratorName.Set(&v)
}
// SetAdministratorNameNil sets the value for AdministratorName to be an explicit nil
func (o *BackupRestoreHistoryInformation) SetAdministratorNameNil() {
	o.AdministratorName.Set(nil)
}

// UnsetAdministratorName ensures that no value is present for AdministratorName, not even an explicit nil
func (o *BackupRestoreHistoryInformation) UnsetAdministratorName() {
	o.AdministratorName.Unset()
}

// GetBackupDetails returns the BackupDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetBackupDetails() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.BackupDetails
}

// GetBackupDetailsOk returns a tuple with the BackupDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetBackupDetailsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.BackupDetails) {
		return nil, false
	}
	return &o.BackupDetails, true
}

// HasBackupDetails returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasBackupDetails() bool {
	if o != nil && IsNil(o.BackupDetails) {
		return true
	}

	return false
}

// SetBackupDetails gets a reference to the given map[string]string and assigns it to the BackupDetails field.
func (o *BackupRestoreHistoryInformation) SetBackupDetails(v map[string]string) {
	o.BackupDetails = v
}

// GetRestoreDetails returns the RestoreDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetRestoreDetails() []BackupRestoreRestoreSingleMemberModel {
	if o == nil {
		var ret []BackupRestoreRestoreSingleMemberModel
		return ret
	}
	return o.RestoreDetails
}

// GetRestoreDetailsOk returns a tuple with the RestoreDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetRestoreDetailsOk() ([]BackupRestoreRestoreSingleMemberModel, bool) {
	if o == nil || IsNil(o.RestoreDetails) {
		return nil, false
	}
	return o.RestoreDetails, true
}

// HasRestoreDetails returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasRestoreDetails() bool {
	if o != nil && IsNil(o.RestoreDetails) {
		return true
	}

	return false
}

// SetRestoreDetails gets a reference to the given []BackupRestoreRestoreSingleMemberModel and assigns it to the RestoreDetails field.
func (o *BackupRestoreHistoryInformation) SetRestoreDetails(v []BackupRestoreRestoreSingleMemberModel) {
	o.RestoreDetails = v
}

// GetSimpleResults returns the SimpleResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetSimpleResults() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SimpleResults
}

// GetSimpleResultsOk returns a tuple with the SimpleResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetSimpleResultsOk() ([]string, bool) {
	if o == nil || IsNil(o.SimpleResults) {
		return nil, false
	}
	return o.SimpleResults, true
}

// HasSimpleResults returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasSimpleResults() bool {
	if o != nil && IsNil(o.SimpleResults) {
		return true
	}

	return false
}

// SetSimpleResults gets a reference to the given []string and assigns it to the SimpleResults field.
func (o *BackupRestoreHistoryInformation) SetSimpleResults(v []string) {
	o.SimpleResults = v
}

// GetFixups returns the Fixups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreHistoryInformation) GetFixups() []BackupRestoreFixupModel {
	if o == nil {
		var ret []BackupRestoreFixupModel
		return ret
	}
	return o.Fixups
}

// GetFixupsOk returns a tuple with the Fixups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreHistoryInformation) GetFixupsOk() ([]BackupRestoreFixupModel, bool) {
	if o == nil || IsNil(o.Fixups) {
		return nil, false
	}
	return o.Fixups, true
}

// HasFixups returns a boolean if a field has been set.
func (o *BackupRestoreHistoryInformation) HasFixups() bool {
	if o != nil && IsNil(o.Fixups) {
		return true
	}

	return false
}

// SetFixups gets a reference to the given []BackupRestoreFixupModel and assigns it to the Fixups field.
func (o *BackupRestoreHistoryInformation) SetFixups(v []BackupRestoreFixupModel) {
	o.Fixups = v
}

func (o BackupRestoreHistoryInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreHistoryInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	toSerialize["Action"] = o.Action
	toSerialize["HistoryId"] = o.HistoryId
	toSerialize["Successful"] = o.Successful
	if !IsNil(o.Duration) {
		toSerialize["Duration"] = o.Duration
	}
	if !IsNil(o.TimeZoneOffset) {
		toSerialize["TimeZoneOffset"] = o.TimeZoneOffset
	}
	if !IsNil(o.DateTime) {
		toSerialize["DateTime"] = o.DateTime
	}
	if o.Notes.IsSet() {
		toSerialize["Notes"] = o.Notes.Get()
	}
	if !IsNil(o.RestoreType) {
		toSerialize["RestoreType"] = o.RestoreType
	}
	if o.Filters.IsSet() {
		toSerialize["Filters"] = o.Filters.Get()
	}
	if !IsNil(o.WithPrerequisites) {
		toSerialize["WithPrerequisites"] = o.WithPrerequisites
	}
	if !IsNil(o.CheckMode) {
		toSerialize["CheckMode"] = o.CheckMode
	}
	if o.ScheduleName.IsSet() {
		toSerialize["ScheduleName"] = o.ScheduleName.Get()
	}
	if !IsNil(o.Component) {
		toSerialize["Component"] = o.Component
	}
	if o.ExecutionId.IsSet() {
		toSerialize["ExecutionId"] = o.ExecutionId.Get()
	}
	if o.BackupName.IsSet() {
		toSerialize["BackupName"] = o.BackupName.Get()
	}
	if !IsNil(o.BackupSize) {
		toSerialize["BackupSize"] = o.BackupSize
	}
	if o.BackupFileSpec.IsSet() {
		toSerialize["BackupFileSpec"] = o.BackupFileSpec.Get()
	}
	if !IsNil(o.RelatedUid) {
		toSerialize["RelatedUid"] = o.RelatedUid
	}
	if !IsNil(o.RelatedDate) {
		toSerialize["RelatedDate"] = o.RelatedDate
	}
	if !IsNil(o.RelatedIsCheckMode) {
		toSerialize["RelatedIsCheckMode"] = o.RelatedIsCheckMode
	}
	if !IsNil(o.Pinned) {
		toSerialize["Pinned"] = o.Pinned
	}
	if o.AdministratorName.IsSet() {
		toSerialize["AdministratorName"] = o.AdministratorName.Get()
	}
	if o.BackupDetails != nil {
		toSerialize["BackupDetails"] = o.BackupDetails
	}
	if o.RestoreDetails != nil {
		toSerialize["RestoreDetails"] = o.RestoreDetails
	}
	if o.SimpleResults != nil {
		toSerialize["SimpleResults"] = o.SimpleResults
	}
	if o.Fixups != nil {
		toSerialize["Fixups"] = o.Fixups
	}
	return toSerialize, nil
}

type NullableBackupRestoreHistoryInformation struct {
	value *BackupRestoreHistoryInformation
	isSet bool
}

func (v NullableBackupRestoreHistoryInformation) Get() *BackupRestoreHistoryInformation {
	return v.value
}

func (v *NullableBackupRestoreHistoryInformation) Set(val *BackupRestoreHistoryInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreHistoryInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreHistoryInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreHistoryInformation(val *BackupRestoreHistoryInformation) *NullableBackupRestoreHistoryInformation {
	return &NullableBackupRestoreHistoryInformation{value: val, isSet: true}
}

func (v NullableBackupRestoreHistoryInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreHistoryInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


