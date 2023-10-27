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

// checks if the ActionCatalogCreationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionCatalogCreationResponseModel{}

// ActionCatalogCreationResponseModel struct for ActionCatalogCreationResponseModel
type ActionCatalogCreationResponseModel struct {
	// The action id.
	ActionId *string `json:"ActionId,omitempty"`
	// The action target name, it's the related catalog name.
	ActionTargetName NullableString `json:"ActionTargetName,omitempty"`
	// The action target uid, it's the related catalog uid.
	ActionTargetUid NullableString `json:"ActionTargetUid,omitempty"`
	// The action target id, it's related on catalog id.
	ActionTargetId NullableString `json:"ActionTargetId,omitempty"`
	ActionType *ActionType `json:"ActionType,omitempty"`
	// The action creation time.
	CreationTime NullableString `json:"CreationTime,omitempty"`
	// The action start time.
	StartTime NullableString `json:"StartTime,omitempty"`
	// The action finish time.
	FinishTime NullableString `json:"FinishTime,omitempty"`
	State *ActionState `json:"State,omitempty"`
	ErrorState *ActionErrorStatus `json:"ErrorState,omitempty"`
	// The progress of the action.
	Progress *float64 `json:"Progress,omitempty"`
	// The progress message.
	ProgressMessage NullableString `json:"ProgressMessage,omitempty"`
	TerminatingError *ActionError `json:"TerminatingError,omitempty"`
	// The non terminating errors.
	NonTerminatingErrors []ActionError `json:"NonTerminatingErrors,omitempty"`
	// The metadata of the action.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	// The scope of the catalog.
	Scopes []ScopeResponseModel `json:"Scopes,omitempty"`
	// The tenant(s) that the hypervisor is assigned to.  If `null`, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants.
	Tenants []RefResponseModel `json:"Tenants,omitempty"`
	// If the task is running.
	IsRunning *bool `json:"IsRunning,omitempty"`
	// The master image.
	MasterImage NullableString `json:"MasterImage,omitempty"`
	MachineCreationData *ActionMachineCreationDetailsResponseModel `json:"MachineCreationData,omitempty"`
	// The snapshot.
	Snapshot NullableString `json:"Snapshot,omitempty"`
	MachineRemovalData *ActionMachineRemovalDetailsResponseModel `json:"MachineRemovalData,omitempty"`
	// The delete virtual machines. 
	DeleteVirtualMachines NullableString `json:"DeleteVirtualMachines,omitempty"`
}

// NewActionCatalogCreationResponseModel instantiates a new ActionCatalogCreationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionCatalogCreationResponseModel() *ActionCatalogCreationResponseModel {
	this := ActionCatalogCreationResponseModel{}
	return &this
}

// NewActionCatalogCreationResponseModelWithDefaults instantiates a new ActionCatalogCreationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionCatalogCreationResponseModelWithDefaults() *ActionCatalogCreationResponseModel {
	this := ActionCatalogCreationResponseModel{}
	return &this
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetActionId() string {
	if o == nil || IsNil(o.ActionId) {
		var ret string
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetActionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasActionId() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given string and assigns it to the ActionId field.
func (o *ActionCatalogCreationResponseModel) SetActionId(v string) {
	o.ActionId = &v
}

// GetActionTargetName returns the ActionTargetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetActionTargetName() string {
	if o == nil || IsNil(o.ActionTargetName.Get()) {
		var ret string
		return ret
	}
	return *o.ActionTargetName.Get()
}

// GetActionTargetNameOk returns a tuple with the ActionTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetActionTargetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionTargetName.Get(), o.ActionTargetName.IsSet()
}

// HasActionTargetName returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasActionTargetName() bool {
	if o != nil && o.ActionTargetName.IsSet() {
		return true
	}

	return false
}

// SetActionTargetName gets a reference to the given NullableString and assigns it to the ActionTargetName field.
func (o *ActionCatalogCreationResponseModel) SetActionTargetName(v string) {
	o.ActionTargetName.Set(&v)
}
// SetActionTargetNameNil sets the value for ActionTargetName to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetActionTargetNameNil() {
	o.ActionTargetName.Set(nil)
}

// UnsetActionTargetName ensures that no value is present for ActionTargetName, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetActionTargetName() {
	o.ActionTargetName.Unset()
}

// GetActionTargetUid returns the ActionTargetUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetActionTargetUid() string {
	if o == nil || IsNil(o.ActionTargetUid.Get()) {
		var ret string
		return ret
	}
	return *o.ActionTargetUid.Get()
}

// GetActionTargetUidOk returns a tuple with the ActionTargetUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetActionTargetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionTargetUid.Get(), o.ActionTargetUid.IsSet()
}

// HasActionTargetUid returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasActionTargetUid() bool {
	if o != nil && o.ActionTargetUid.IsSet() {
		return true
	}

	return false
}

// SetActionTargetUid gets a reference to the given NullableString and assigns it to the ActionTargetUid field.
func (o *ActionCatalogCreationResponseModel) SetActionTargetUid(v string) {
	o.ActionTargetUid.Set(&v)
}
// SetActionTargetUidNil sets the value for ActionTargetUid to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetActionTargetUidNil() {
	o.ActionTargetUid.Set(nil)
}

// UnsetActionTargetUid ensures that no value is present for ActionTargetUid, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetActionTargetUid() {
	o.ActionTargetUid.Unset()
}

// GetActionTargetId returns the ActionTargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetActionTargetId() string {
	if o == nil || IsNil(o.ActionTargetId.Get()) {
		var ret string
		return ret
	}
	return *o.ActionTargetId.Get()
}

// GetActionTargetIdOk returns a tuple with the ActionTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetActionTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionTargetId.Get(), o.ActionTargetId.IsSet()
}

// HasActionTargetId returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasActionTargetId() bool {
	if o != nil && o.ActionTargetId.IsSet() {
		return true
	}

	return false
}

// SetActionTargetId gets a reference to the given NullableString and assigns it to the ActionTargetId field.
func (o *ActionCatalogCreationResponseModel) SetActionTargetId(v string) {
	o.ActionTargetId.Set(&v)
}
// SetActionTargetIdNil sets the value for ActionTargetId to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetActionTargetIdNil() {
	o.ActionTargetId.Set(nil)
}

// UnsetActionTargetId ensures that no value is present for ActionTargetId, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetActionTargetId() {
	o.ActionTargetId.Unset()
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetActionType() ActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret ActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetActionTypeOk() (*ActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given ActionType and assigns it to the ActionType field.
func (o *ActionCatalogCreationResponseModel) SetActionType(v ActionType) {
	o.ActionType = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime.Get()) {
		var ret string
		return ret
	}
	return *o.CreationTime.Get()
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetCreationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationTime.Get(), o.CreationTime.IsSet()
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasCreationTime() bool {
	if o != nil && o.CreationTime.IsSet() {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given NullableString and assigns it to the CreationTime field.
func (o *ActionCatalogCreationResponseModel) SetCreationTime(v string) {
	o.CreationTime.Set(&v)
}
// SetCreationTimeNil sets the value for CreationTime to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetCreationTimeNil() {
	o.CreationTime.Set(nil)
}

// UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetCreationTime() {
	o.CreationTime.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime.Get()) {
		var ret string
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableString and assigns it to the StartTime field.
func (o *ActionCatalogCreationResponseModel) SetStartTime(v string) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetFinishTime() string {
	if o == nil || IsNil(o.FinishTime.Get()) {
		var ret string
		return ret
	}
	return *o.FinishTime.Get()
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetFinishTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishTime.Get(), o.FinishTime.IsSet()
}

// HasFinishTime returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasFinishTime() bool {
	if o != nil && o.FinishTime.IsSet() {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given NullableString and assigns it to the FinishTime field.
func (o *ActionCatalogCreationResponseModel) SetFinishTime(v string) {
	o.FinishTime.Set(&v)
}
// SetFinishTimeNil sets the value for FinishTime to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetFinishTimeNil() {
	o.FinishTime.Set(nil)
}

// UnsetFinishTime ensures that no value is present for FinishTime, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetFinishTime() {
	o.FinishTime.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetState() ActionState {
	if o == nil || IsNil(o.State) {
		var ret ActionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetStateOk() (*ActionState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ActionState and assigns it to the State field.
func (o *ActionCatalogCreationResponseModel) SetState(v ActionState) {
	o.State = &v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetErrorState() ActionErrorStatus {
	if o == nil || IsNil(o.ErrorState) {
		var ret ActionErrorStatus
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool) {
	if o == nil || IsNil(o.ErrorState) {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasErrorState() bool {
	if o != nil && !IsNil(o.ErrorState) {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given ActionErrorStatus and assigns it to the ErrorState field.
func (o *ActionCatalogCreationResponseModel) SetErrorState(v ActionErrorStatus) {
	o.ErrorState = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *ActionCatalogCreationResponseModel) SetProgress(v float64) {
	o.Progress = &v
}

// GetProgressMessage returns the ProgressMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetProgressMessage() string {
	if o == nil || IsNil(o.ProgressMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ProgressMessage.Get()
}

// GetProgressMessageOk returns a tuple with the ProgressMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetProgressMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgressMessage.Get(), o.ProgressMessage.IsSet()
}

// HasProgressMessage returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasProgressMessage() bool {
	if o != nil && o.ProgressMessage.IsSet() {
		return true
	}

	return false
}

// SetProgressMessage gets a reference to the given NullableString and assigns it to the ProgressMessage field.
func (o *ActionCatalogCreationResponseModel) SetProgressMessage(v string) {
	o.ProgressMessage.Set(&v)
}
// SetProgressMessageNil sets the value for ProgressMessage to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetProgressMessageNil() {
	o.ProgressMessage.Set(nil)
}

// UnsetProgressMessage ensures that no value is present for ProgressMessage, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetProgressMessage() {
	o.ProgressMessage.Unset()
}

// GetTerminatingError returns the TerminatingError field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetTerminatingError() ActionError {
	if o == nil || IsNil(o.TerminatingError) {
		var ret ActionError
		return ret
	}
	return *o.TerminatingError
}

// GetTerminatingErrorOk returns a tuple with the TerminatingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetTerminatingErrorOk() (*ActionError, bool) {
	if o == nil || IsNil(o.TerminatingError) {
		return nil, false
	}
	return o.TerminatingError, true
}

// HasTerminatingError returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasTerminatingError() bool {
	if o != nil && !IsNil(o.TerminatingError) {
		return true
	}

	return false
}

// SetTerminatingError gets a reference to the given ActionError and assigns it to the TerminatingError field.
func (o *ActionCatalogCreationResponseModel) SetTerminatingError(v ActionError) {
	o.TerminatingError = &v
}

// GetNonTerminatingErrors returns the NonTerminatingErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetNonTerminatingErrors() []ActionError {
	if o == nil {
		var ret []ActionError
		return ret
	}
	return o.NonTerminatingErrors
}

// GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetNonTerminatingErrorsOk() ([]ActionError, bool) {
	if o == nil || IsNil(o.NonTerminatingErrors) {
		return nil, false
	}
	return o.NonTerminatingErrors, true
}

// HasNonTerminatingErrors returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasNonTerminatingErrors() bool {
	if o != nil && IsNil(o.NonTerminatingErrors) {
		return true
	}

	return false
}

// SetNonTerminatingErrors gets a reference to the given []ActionError and assigns it to the NonTerminatingErrors field.
func (o *ActionCatalogCreationResponseModel) SetNonTerminatingErrors(v []ActionError) {
	o.NonTerminatingErrors = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetMetadata() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *ActionCatalogCreationResponseModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetScopes() []ScopeResponseModel {
	if o == nil {
		var ret []ScopeResponseModel
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetScopesOk() ([]ScopeResponseModel, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ScopeResponseModel and assigns it to the Scopes field.
func (o *ActionCatalogCreationResponseModel) SetScopes(v []ScopeResponseModel) {
	o.Scopes = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetTenants() []RefResponseModel {
	if o == nil {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasTenants() bool {
	if o != nil && IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *ActionCatalogCreationResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetIsRunning() bool {
	if o == nil || IsNil(o.IsRunning) {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetIsRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRunning) {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasIsRunning() bool {
	if o != nil && !IsNil(o.IsRunning) {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *ActionCatalogCreationResponseModel) SetIsRunning(v bool) {
	o.IsRunning = &v
}

// GetMasterImage returns the MasterImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetMasterImage() string {
	if o == nil || IsNil(o.MasterImage.Get()) {
		var ret string
		return ret
	}
	return *o.MasterImage.Get()
}

// GetMasterImageOk returns a tuple with the MasterImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetMasterImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MasterImage.Get(), o.MasterImage.IsSet()
}

// HasMasterImage returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasMasterImage() bool {
	if o != nil && o.MasterImage.IsSet() {
		return true
	}

	return false
}

// SetMasterImage gets a reference to the given NullableString and assigns it to the MasterImage field.
func (o *ActionCatalogCreationResponseModel) SetMasterImage(v string) {
	o.MasterImage.Set(&v)
}
// SetMasterImageNil sets the value for MasterImage to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetMasterImageNil() {
	o.MasterImage.Set(nil)
}

// UnsetMasterImage ensures that no value is present for MasterImage, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetMasterImage() {
	o.MasterImage.Unset()
}

// GetMachineCreationData returns the MachineCreationData field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetMachineCreationData() ActionMachineCreationDetailsResponseModel {
	if o == nil || IsNil(o.MachineCreationData) {
		var ret ActionMachineCreationDetailsResponseModel
		return ret
	}
	return *o.MachineCreationData
}

// GetMachineCreationDataOk returns a tuple with the MachineCreationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetMachineCreationDataOk() (*ActionMachineCreationDetailsResponseModel, bool) {
	if o == nil || IsNil(o.MachineCreationData) {
		return nil, false
	}
	return o.MachineCreationData, true
}

// HasMachineCreationData returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasMachineCreationData() bool {
	if o != nil && !IsNil(o.MachineCreationData) {
		return true
	}

	return false
}

// SetMachineCreationData gets a reference to the given ActionMachineCreationDetailsResponseModel and assigns it to the MachineCreationData field.
func (o *ActionCatalogCreationResponseModel) SetMachineCreationData(v ActionMachineCreationDetailsResponseModel) {
	o.MachineCreationData = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetSnapshot() string {
	if o == nil || IsNil(o.Snapshot.Get()) {
		var ret string
		return ret
	}
	return *o.Snapshot.Get()
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetSnapshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Snapshot.Get(), o.Snapshot.IsSet()
}

// HasSnapshot returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasSnapshot() bool {
	if o != nil && o.Snapshot.IsSet() {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given NullableString and assigns it to the Snapshot field.
func (o *ActionCatalogCreationResponseModel) SetSnapshot(v string) {
	o.Snapshot.Set(&v)
}
// SetSnapshotNil sets the value for Snapshot to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetSnapshotNil() {
	o.Snapshot.Set(nil)
}

// UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetSnapshot() {
	o.Snapshot.Unset()
}

// GetMachineRemovalData returns the MachineRemovalData field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModel) GetMachineRemovalData() ActionMachineRemovalDetailsResponseModel {
	if o == nil || IsNil(o.MachineRemovalData) {
		var ret ActionMachineRemovalDetailsResponseModel
		return ret
	}
	return *o.MachineRemovalData
}

// GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalDetailsResponseModel, bool) {
	if o == nil || IsNil(o.MachineRemovalData) {
		return nil, false
	}
	return o.MachineRemovalData, true
}

// HasMachineRemovalData returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasMachineRemovalData() bool {
	if o != nil && !IsNil(o.MachineRemovalData) {
		return true
	}

	return false
}

// SetMachineRemovalData gets a reference to the given ActionMachineRemovalDetailsResponseModel and assigns it to the MachineRemovalData field.
func (o *ActionCatalogCreationResponseModel) SetMachineRemovalData(v ActionMachineRemovalDetailsResponseModel) {
	o.MachineRemovalData = &v
}

// GetDeleteVirtualMachines returns the DeleteVirtualMachines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionCatalogCreationResponseModel) GetDeleteVirtualMachines() string {
	if o == nil || IsNil(o.DeleteVirtualMachines.Get()) {
		var ret string
		return ret
	}
	return *o.DeleteVirtualMachines.Get()
}

// GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionCatalogCreationResponseModel) GetDeleteVirtualMachinesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteVirtualMachines.Get(), o.DeleteVirtualMachines.IsSet()
}

// HasDeleteVirtualMachines returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModel) HasDeleteVirtualMachines() bool {
	if o != nil && o.DeleteVirtualMachines.IsSet() {
		return true
	}

	return false
}

// SetDeleteVirtualMachines gets a reference to the given NullableString and assigns it to the DeleteVirtualMachines field.
func (o *ActionCatalogCreationResponseModel) SetDeleteVirtualMachines(v string) {
	o.DeleteVirtualMachines.Set(&v)
}
// SetDeleteVirtualMachinesNil sets the value for DeleteVirtualMachines to be an explicit nil
func (o *ActionCatalogCreationResponseModel) SetDeleteVirtualMachinesNil() {
	o.DeleteVirtualMachines.Set(nil)
}

// UnsetDeleteVirtualMachines ensures that no value is present for DeleteVirtualMachines, not even an explicit nil
func (o *ActionCatalogCreationResponseModel) UnsetDeleteVirtualMachines() {
	o.DeleteVirtualMachines.Unset()
}

func (o ActionCatalogCreationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionCatalogCreationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionId) {
		toSerialize["ActionId"] = o.ActionId
	}
	if o.ActionTargetName.IsSet() {
		toSerialize["ActionTargetName"] = o.ActionTargetName.Get()
	}
	if o.ActionTargetUid.IsSet() {
		toSerialize["ActionTargetUid"] = o.ActionTargetUid.Get()
	}
	if o.ActionTargetId.IsSet() {
		toSerialize["ActionTargetId"] = o.ActionTargetId.Get()
	}
	if !IsNil(o.ActionType) {
		toSerialize["ActionType"] = o.ActionType
	}
	if o.CreationTime.IsSet() {
		toSerialize["CreationTime"] = o.CreationTime.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["StartTime"] = o.StartTime.Get()
	}
	if o.FinishTime.IsSet() {
		toSerialize["FinishTime"] = o.FinishTime.Get()
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.ErrorState) {
		toSerialize["ErrorState"] = o.ErrorState
	}
	if !IsNil(o.Progress) {
		toSerialize["Progress"] = o.Progress
	}
	if o.ProgressMessage.IsSet() {
		toSerialize["ProgressMessage"] = o.ProgressMessage.Get()
	}
	if !IsNil(o.TerminatingError) {
		toSerialize["TerminatingError"] = o.TerminatingError
	}
	if o.NonTerminatingErrors != nil {
		toSerialize["NonTerminatingErrors"] = o.NonTerminatingErrors
	}
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if o.Scopes != nil {
		toSerialize["Scopes"] = o.Scopes
	}
	if o.Tenants != nil {
		toSerialize["Tenants"] = o.Tenants
	}
	if !IsNil(o.IsRunning) {
		toSerialize["IsRunning"] = o.IsRunning
	}
	if o.MasterImage.IsSet() {
		toSerialize["MasterImage"] = o.MasterImage.Get()
	}
	if !IsNil(o.MachineCreationData) {
		toSerialize["MachineCreationData"] = o.MachineCreationData
	}
	if o.Snapshot.IsSet() {
		toSerialize["Snapshot"] = o.Snapshot.Get()
	}
	if !IsNil(o.MachineRemovalData) {
		toSerialize["MachineRemovalData"] = o.MachineRemovalData
	}
	if o.DeleteVirtualMachines.IsSet() {
		toSerialize["DeleteVirtualMachines"] = o.DeleteVirtualMachines.Get()
	}
	return toSerialize, nil
}

type NullableActionCatalogCreationResponseModel struct {
	value *ActionCatalogCreationResponseModel
	isSet bool
}

func (v NullableActionCatalogCreationResponseModel) Get() *ActionCatalogCreationResponseModel {
	return v.value
}

func (v *NullableActionCatalogCreationResponseModel) Set(val *ActionCatalogCreationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableActionCatalogCreationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableActionCatalogCreationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionCatalogCreationResponseModel(val *ActionCatalogCreationResponseModel) *NullableActionCatalogCreationResponseModel {
	return &NullableActionCatalogCreationResponseModel{value: val, isSet: true}
}

func (v NullableActionCatalogCreationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionCatalogCreationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


