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

// checks if the ActionMachineRemovalResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionMachineRemovalResponseModel{}

// ActionMachineRemovalResponseModel struct for ActionMachineRemovalResponseModel
type ActionMachineRemovalResponseModel struct {
	// The action id.
	ActionId *string `json:"ActionId,omitempty"`
	// The action target name, it's the related catalog name.
	ActionTargetName *string `json:"ActionTargetName,omitempty"`
	// The action target uid, it's the related catalog uid.
	ActionTargetUid *string `json:"ActionTargetUid,omitempty"`
	// The action target id, it's related on catalog id.
	ActionTargetId *string `json:"ActionTargetId,omitempty"`
	ActionType *ActionType `json:"ActionType,omitempty"`
	// The action creation time.
	CreationTime *string `json:"CreationTime,omitempty"`
	// The action start time.
	StartTime *string `json:"StartTime,omitempty"`
	// The action finish time.
	FinishTime *string `json:"FinishTime,omitempty"`
	State *ActionState `json:"State,omitempty"`
	ErrorState *ActionErrorStatus `json:"ErrorState,omitempty"`
	// The progress of the action.
	Progress *float64 `json:"Progress,omitempty"`
	// The progress message.
	ProgressMessage *string `json:"ProgressMessage,omitempty"`
	TerminatingError *ActionResponseModelTerminatingError `json:"TerminatingError,omitempty"`
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
	MasterImage *string `json:"MasterImage,omitempty"`
	MachineCreationData *ActionCatalogCreationResponseModelAllOfMachineCreationData `json:"MachineCreationData,omitempty"`
	// The snapshot.
	Snapshot *string `json:"Snapshot,omitempty"`
	MachineRemovalData *ActionMachineRemovalResponseModelAllOfMachineRemovalData `json:"MachineRemovalData,omitempty"`
	// The delete virtual machines. 
	DeleteVirtualMachines *string `json:"DeleteVirtualMachines,omitempty"`
}

// NewActionMachineRemovalResponseModel instantiates a new ActionMachineRemovalResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionMachineRemovalResponseModel() *ActionMachineRemovalResponseModel {
	this := ActionMachineRemovalResponseModel{}
	return &this
}

// NewActionMachineRemovalResponseModelWithDefaults instantiates a new ActionMachineRemovalResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionMachineRemovalResponseModelWithDefaults() *ActionMachineRemovalResponseModel {
	this := ActionMachineRemovalResponseModel{}
	return &this
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetActionId() string {
	if o == nil || IsNil(o.ActionId) {
		var ret string
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetActionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasActionId() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given string and assigns it to the ActionId field.
func (o *ActionMachineRemovalResponseModel) SetActionId(v string) {
	o.ActionId = &v
}

// GetActionTargetName returns the ActionTargetName field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetActionTargetName() string {
	if o == nil || IsNil(o.ActionTargetName) {
		var ret string
		return ret
	}
	return *o.ActionTargetName
}

// GetActionTargetNameOk returns a tuple with the ActionTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetActionTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTargetName) {
		return nil, false
	}
	return o.ActionTargetName, true
}

// HasActionTargetName returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasActionTargetName() bool {
	if o != nil && !IsNil(o.ActionTargetName) {
		return true
	}

	return false
}

// SetActionTargetName gets a reference to the given string and assigns it to the ActionTargetName field.
func (o *ActionMachineRemovalResponseModel) SetActionTargetName(v string) {
	o.ActionTargetName = &v
}

// GetActionTargetUid returns the ActionTargetUid field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetActionTargetUid() string {
	if o == nil || IsNil(o.ActionTargetUid) {
		var ret string
		return ret
	}
	return *o.ActionTargetUid
}

// GetActionTargetUidOk returns a tuple with the ActionTargetUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetActionTargetUidOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTargetUid) {
		return nil, false
	}
	return o.ActionTargetUid, true
}

// HasActionTargetUid returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasActionTargetUid() bool {
	if o != nil && !IsNil(o.ActionTargetUid) {
		return true
	}

	return false
}

// SetActionTargetUid gets a reference to the given string and assigns it to the ActionTargetUid field.
func (o *ActionMachineRemovalResponseModel) SetActionTargetUid(v string) {
	o.ActionTargetUid = &v
}

// GetActionTargetId returns the ActionTargetId field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetActionTargetId() string {
	if o == nil || IsNil(o.ActionTargetId) {
		var ret string
		return ret
	}
	return *o.ActionTargetId
}

// GetActionTargetIdOk returns a tuple with the ActionTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetActionTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTargetId) {
		return nil, false
	}
	return o.ActionTargetId, true
}

// HasActionTargetId returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasActionTargetId() bool {
	if o != nil && !IsNil(o.ActionTargetId) {
		return true
	}

	return false
}

// SetActionTargetId gets a reference to the given string and assigns it to the ActionTargetId field.
func (o *ActionMachineRemovalResponseModel) SetActionTargetId(v string) {
	o.ActionTargetId = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetActionType() ActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret ActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetActionTypeOk() (*ActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given ActionType and assigns it to the ActionType field.
func (o *ActionMachineRemovalResponseModel) SetActionType(v ActionType) {
	o.ActionType = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *ActionMachineRemovalResponseModel) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ActionMachineRemovalResponseModel) SetStartTime(v string) {
	o.StartTime = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetFinishTime() string {
	if o == nil || IsNil(o.FinishTime) {
		var ret string
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetFinishTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FinishTime) {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasFinishTime() bool {
	if o != nil && !IsNil(o.FinishTime) {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given string and assigns it to the FinishTime field.
func (o *ActionMachineRemovalResponseModel) SetFinishTime(v string) {
	o.FinishTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetState() ActionState {
	if o == nil || IsNil(o.State) {
		var ret ActionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetStateOk() (*ActionState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ActionState and assigns it to the State field.
func (o *ActionMachineRemovalResponseModel) SetState(v ActionState) {
	o.State = &v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetErrorState() ActionErrorStatus {
	if o == nil || IsNil(o.ErrorState) {
		var ret ActionErrorStatus
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool) {
	if o == nil || IsNil(o.ErrorState) {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasErrorState() bool {
	if o != nil && !IsNil(o.ErrorState) {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given ActionErrorStatus and assigns it to the ErrorState field.
func (o *ActionMachineRemovalResponseModel) SetErrorState(v ActionErrorStatus) {
	o.ErrorState = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *ActionMachineRemovalResponseModel) SetProgress(v float64) {
	o.Progress = &v
}

// GetProgressMessage returns the ProgressMessage field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetProgressMessage() string {
	if o == nil || IsNil(o.ProgressMessage) {
		var ret string
		return ret
	}
	return *o.ProgressMessage
}

// GetProgressMessageOk returns a tuple with the ProgressMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetProgressMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ProgressMessage) {
		return nil, false
	}
	return o.ProgressMessage, true
}

// HasProgressMessage returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasProgressMessage() bool {
	if o != nil && !IsNil(o.ProgressMessage) {
		return true
	}

	return false
}

// SetProgressMessage gets a reference to the given string and assigns it to the ProgressMessage field.
func (o *ActionMachineRemovalResponseModel) SetProgressMessage(v string) {
	o.ProgressMessage = &v
}

// GetTerminatingError returns the TerminatingError field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetTerminatingError() ActionResponseModelTerminatingError {
	if o == nil || IsNil(o.TerminatingError) {
		var ret ActionResponseModelTerminatingError
		return ret
	}
	return *o.TerminatingError
}

// GetTerminatingErrorOk returns a tuple with the TerminatingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetTerminatingErrorOk() (*ActionResponseModelTerminatingError, bool) {
	if o == nil || IsNil(o.TerminatingError) {
		return nil, false
	}
	return o.TerminatingError, true
}

// HasTerminatingError returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasTerminatingError() bool {
	if o != nil && !IsNil(o.TerminatingError) {
		return true
	}

	return false
}

// SetTerminatingError gets a reference to the given ActionResponseModelTerminatingError and assigns it to the TerminatingError field.
func (o *ActionMachineRemovalResponseModel) SetTerminatingError(v ActionResponseModelTerminatingError) {
	o.TerminatingError = &v
}

// GetNonTerminatingErrors returns the NonTerminatingErrors field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetNonTerminatingErrors() []ActionError {
	if o == nil || IsNil(o.NonTerminatingErrors) {
		var ret []ActionError
		return ret
	}
	return o.NonTerminatingErrors
}

// GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetNonTerminatingErrorsOk() ([]ActionError, bool) {
	if o == nil || IsNil(o.NonTerminatingErrors) {
		return nil, false
	}
	return o.NonTerminatingErrors, true
}

// HasNonTerminatingErrors returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasNonTerminatingErrors() bool {
	if o != nil && !IsNil(o.NonTerminatingErrors) {
		return true
	}

	return false
}

// SetNonTerminatingErrors gets a reference to the given []ActionError and assigns it to the NonTerminatingErrors field.
func (o *ActionMachineRemovalResponseModel) SetNonTerminatingErrors(v []ActionError) {
	o.NonTerminatingErrors = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetMetadata() []NameValueStringPairModel {
	if o == nil || IsNil(o.Metadata) {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *ActionMachineRemovalResponseModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetScopes() []ScopeResponseModel {
	if o == nil || IsNil(o.Scopes) {
		var ret []ScopeResponseModel
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetScopesOk() ([]ScopeResponseModel, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ScopeResponseModel and assigns it to the Scopes field.
func (o *ActionMachineRemovalResponseModel) SetScopes(v []ScopeResponseModel) {
	o.Scopes = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetTenants() []RefResponseModel {
	if o == nil || IsNil(o.Tenants) {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *ActionMachineRemovalResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetIsRunning() bool {
	if o == nil || IsNil(o.IsRunning) {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetIsRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRunning) {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasIsRunning() bool {
	if o != nil && !IsNil(o.IsRunning) {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *ActionMachineRemovalResponseModel) SetIsRunning(v bool) {
	o.IsRunning = &v
}

// GetMasterImage returns the MasterImage field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetMasterImage() string {
	if o == nil || IsNil(o.MasterImage) {
		var ret string
		return ret
	}
	return *o.MasterImage
}

// GetMasterImageOk returns a tuple with the MasterImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetMasterImageOk() (*string, bool) {
	if o == nil || IsNil(o.MasterImage) {
		return nil, false
	}
	return o.MasterImage, true
}

// HasMasterImage returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasMasterImage() bool {
	if o != nil && !IsNil(o.MasterImage) {
		return true
	}

	return false
}

// SetMasterImage gets a reference to the given string and assigns it to the MasterImage field.
func (o *ActionMachineRemovalResponseModel) SetMasterImage(v string) {
	o.MasterImage = &v
}

// GetMachineCreationData returns the MachineCreationData field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetMachineCreationData() ActionCatalogCreationResponseModelAllOfMachineCreationData {
	if o == nil || IsNil(o.MachineCreationData) {
		var ret ActionCatalogCreationResponseModelAllOfMachineCreationData
		return ret
	}
	return *o.MachineCreationData
}

// GetMachineCreationDataOk returns a tuple with the MachineCreationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetMachineCreationDataOk() (*ActionCatalogCreationResponseModelAllOfMachineCreationData, bool) {
	if o == nil || IsNil(o.MachineCreationData) {
		return nil, false
	}
	return o.MachineCreationData, true
}

// HasMachineCreationData returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasMachineCreationData() bool {
	if o != nil && !IsNil(o.MachineCreationData) {
		return true
	}

	return false
}

// SetMachineCreationData gets a reference to the given ActionCatalogCreationResponseModelAllOfMachineCreationData and assigns it to the MachineCreationData field.
func (o *ActionMachineRemovalResponseModel) SetMachineCreationData(v ActionCatalogCreationResponseModelAllOfMachineCreationData) {
	o.MachineCreationData = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetSnapshot() string {
	if o == nil || IsNil(o.Snapshot) {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *ActionMachineRemovalResponseModel) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetMachineRemovalData returns the MachineRemovalData field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetMachineRemovalData() ActionMachineRemovalResponseModelAllOfMachineRemovalData {
	if o == nil || IsNil(o.MachineRemovalData) {
		var ret ActionMachineRemovalResponseModelAllOfMachineRemovalData
		return ret
	}
	return *o.MachineRemovalData
}

// GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalResponseModelAllOfMachineRemovalData, bool) {
	if o == nil || IsNil(o.MachineRemovalData) {
		return nil, false
	}
	return o.MachineRemovalData, true
}

// HasMachineRemovalData returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasMachineRemovalData() bool {
	if o != nil && !IsNil(o.MachineRemovalData) {
		return true
	}

	return false
}

// SetMachineRemovalData gets a reference to the given ActionMachineRemovalResponseModelAllOfMachineRemovalData and assigns it to the MachineRemovalData field.
func (o *ActionMachineRemovalResponseModel) SetMachineRemovalData(v ActionMachineRemovalResponseModelAllOfMachineRemovalData) {
	o.MachineRemovalData = &v
}

// GetDeleteVirtualMachines returns the DeleteVirtualMachines field value if set, zero value otherwise.
func (o *ActionMachineRemovalResponseModel) GetDeleteVirtualMachines() string {
	if o == nil || IsNil(o.DeleteVirtualMachines) {
		var ret string
		return ret
	}
	return *o.DeleteVirtualMachines
}

// GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalResponseModel) GetDeleteVirtualMachinesOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteVirtualMachines) {
		return nil, false
	}
	return o.DeleteVirtualMachines, true
}

// HasDeleteVirtualMachines returns a boolean if a field has been set.
func (o *ActionMachineRemovalResponseModel) HasDeleteVirtualMachines() bool {
	if o != nil && !IsNil(o.DeleteVirtualMachines) {
		return true
	}

	return false
}

// SetDeleteVirtualMachines gets a reference to the given string and assigns it to the DeleteVirtualMachines field.
func (o *ActionMachineRemovalResponseModel) SetDeleteVirtualMachines(v string) {
	o.DeleteVirtualMachines = &v
}

func (o ActionMachineRemovalResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionMachineRemovalResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionId) {
		toSerialize["ActionId"] = o.ActionId
	}
	if !IsNil(o.ActionTargetName) {
		toSerialize["ActionTargetName"] = o.ActionTargetName
	}
	if !IsNil(o.ActionTargetUid) {
		toSerialize["ActionTargetUid"] = o.ActionTargetUid
	}
	if !IsNil(o.ActionTargetId) {
		toSerialize["ActionTargetId"] = o.ActionTargetId
	}
	if !IsNil(o.ActionType) {
		toSerialize["ActionType"] = o.ActionType
	}
	if !IsNil(o.CreationTime) {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if !IsNil(o.FinishTime) {
		toSerialize["FinishTime"] = o.FinishTime
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
	if !IsNil(o.ProgressMessage) {
		toSerialize["ProgressMessage"] = o.ProgressMessage
	}
	if !IsNil(o.TerminatingError) {
		toSerialize["TerminatingError"] = o.TerminatingError
	}
	if !IsNil(o.NonTerminatingErrors) {
		toSerialize["NonTerminatingErrors"] = o.NonTerminatingErrors
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Scopes) {
		toSerialize["Scopes"] = o.Scopes
	}
	if !IsNil(o.Tenants) {
		toSerialize["Tenants"] = o.Tenants
	}
	if !IsNil(o.IsRunning) {
		toSerialize["IsRunning"] = o.IsRunning
	}
	if !IsNil(o.MasterImage) {
		toSerialize["MasterImage"] = o.MasterImage
	}
	if !IsNil(o.MachineCreationData) {
		toSerialize["MachineCreationData"] = o.MachineCreationData
	}
	if !IsNil(o.Snapshot) {
		toSerialize["Snapshot"] = o.Snapshot
	}
	if !IsNil(o.MachineRemovalData) {
		toSerialize["MachineRemovalData"] = o.MachineRemovalData
	}
	if !IsNil(o.DeleteVirtualMachines) {
		toSerialize["DeleteVirtualMachines"] = o.DeleteVirtualMachines
	}
	return toSerialize, nil
}

type NullableActionMachineRemovalResponseModel struct {
	value *ActionMachineRemovalResponseModel
	isSet bool
}

func (v NullableActionMachineRemovalResponseModel) Get() *ActionMachineRemovalResponseModel {
	return v.value
}

func (v *NullableActionMachineRemovalResponseModel) Set(val *ActionMachineRemovalResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableActionMachineRemovalResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableActionMachineRemovalResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionMachineRemovalResponseModel(val *ActionMachineRemovalResponseModel) *NullableActionMachineRemovalResponseModel {
	return &NullableActionMachineRemovalResponseModel{value: val, isSet: true}
}

func (v NullableActionMachineRemovalResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionMachineRemovalResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


