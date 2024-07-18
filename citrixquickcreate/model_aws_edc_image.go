/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the AwsEdcImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcImage{}

// AwsEdcImage struct for AwsEdcImage
type AwsEdcImage struct {
	Image
	// Provides additional status on image like  error message
	Status NullableString `json:"status,omitempty"`
	IngestionProcess NullableAwsEdcWorkspaceImageIngestionProcess `json:"ingestionProcess,omitempty"`
	WorkspaceImageTenancy NullableAwsEdcWorkspaceImageTenancy `json:"workspaceImageTenancy,omitempty"`
	WorkspaceImageState NullableAwsEdcWorkspaceImageState `json:"workspaceImageState,omitempty"`
	// The list of installed applications
	ApplicationList []AwsEdcAmiImportApplications `json:"applicationList,omitempty"`
}

// NewAwsEdcImage instantiates a new AwsEdcImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcImage(accountType AccountType) *AwsEdcImage {
	this := AwsEdcImage{}
	this.AccountType = accountType
	return &this
}

// NewAwsEdcImageWithDefaults instantiates a new AwsEdcImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcImageWithDefaults() *AwsEdcImage {
	this := AwsEdcImage{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcImage) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcImage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsEdcImage) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *AwsEdcImage) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AwsEdcImage) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AwsEdcImage) UnsetStatus() {
	o.Status.Unset()
}

// GetIngestionProcess returns the IngestionProcess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcImage) GetIngestionProcess() AwsEdcWorkspaceImageIngestionProcess {
	if o == nil || IsNil(o.IngestionProcess.Get()) {
		var ret AwsEdcWorkspaceImageIngestionProcess
		return ret
	}
	return *o.IngestionProcess.Get()
}

// GetIngestionProcessOk returns a tuple with the IngestionProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcImage) GetIngestionProcessOk() (*AwsEdcWorkspaceImageIngestionProcess, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestionProcess.Get(), o.IngestionProcess.IsSet()
}

// HasIngestionProcess returns a boolean if a field has been set.
func (o *AwsEdcImage) HasIngestionProcess() bool {
	if o != nil && o.IngestionProcess.IsSet() {
		return true
	}

	return false
}

// SetIngestionProcess gets a reference to the given NullableAwsEdcWorkspaceImageIngestionProcess and assigns it to the IngestionProcess field.
func (o *AwsEdcImage) SetIngestionProcess(v AwsEdcWorkspaceImageIngestionProcess) {
	o.IngestionProcess.Set(&v)
}
// SetIngestionProcessNil sets the value for IngestionProcess to be an explicit nil
func (o *AwsEdcImage) SetIngestionProcessNil() {
	o.IngestionProcess.Set(nil)
}

// UnsetIngestionProcess ensures that no value is present for IngestionProcess, not even an explicit nil
func (o *AwsEdcImage) UnsetIngestionProcess() {
	o.IngestionProcess.Unset()
}

// GetWorkspaceImageTenancy returns the WorkspaceImageTenancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcImage) GetWorkspaceImageTenancy() AwsEdcWorkspaceImageTenancy {
	if o == nil || IsNil(o.WorkspaceImageTenancy.Get()) {
		var ret AwsEdcWorkspaceImageTenancy
		return ret
	}
	return *o.WorkspaceImageTenancy.Get()
}

// GetWorkspaceImageTenancyOk returns a tuple with the WorkspaceImageTenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcImage) GetWorkspaceImageTenancyOk() (*AwsEdcWorkspaceImageTenancy, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceImageTenancy.Get(), o.WorkspaceImageTenancy.IsSet()
}

// HasWorkspaceImageTenancy returns a boolean if a field has been set.
func (o *AwsEdcImage) HasWorkspaceImageTenancy() bool {
	if o != nil && o.WorkspaceImageTenancy.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceImageTenancy gets a reference to the given NullableAwsEdcWorkspaceImageTenancy and assigns it to the WorkspaceImageTenancy field.
func (o *AwsEdcImage) SetWorkspaceImageTenancy(v AwsEdcWorkspaceImageTenancy) {
	o.WorkspaceImageTenancy.Set(&v)
}
// SetWorkspaceImageTenancyNil sets the value for WorkspaceImageTenancy to be an explicit nil
func (o *AwsEdcImage) SetWorkspaceImageTenancyNil() {
	o.WorkspaceImageTenancy.Set(nil)
}

// UnsetWorkspaceImageTenancy ensures that no value is present for WorkspaceImageTenancy, not even an explicit nil
func (o *AwsEdcImage) UnsetWorkspaceImageTenancy() {
	o.WorkspaceImageTenancy.Unset()
}

// GetWorkspaceImageState returns the WorkspaceImageState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcImage) GetWorkspaceImageState() AwsEdcWorkspaceImageState {
	if o == nil || IsNil(o.WorkspaceImageState.Get()) {
		var ret AwsEdcWorkspaceImageState
		return ret
	}
	return *o.WorkspaceImageState.Get()
}

// GetWorkspaceImageStateOk returns a tuple with the WorkspaceImageState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcImage) GetWorkspaceImageStateOk() (*AwsEdcWorkspaceImageState, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceImageState.Get(), o.WorkspaceImageState.IsSet()
}

// HasWorkspaceImageState returns a boolean if a field has been set.
func (o *AwsEdcImage) HasWorkspaceImageState() bool {
	if o != nil && o.WorkspaceImageState.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceImageState gets a reference to the given NullableAwsEdcWorkspaceImageState and assigns it to the WorkspaceImageState field.
func (o *AwsEdcImage) SetWorkspaceImageState(v AwsEdcWorkspaceImageState) {
	o.WorkspaceImageState.Set(&v)
}
// SetWorkspaceImageStateNil sets the value for WorkspaceImageState to be an explicit nil
func (o *AwsEdcImage) SetWorkspaceImageStateNil() {
	o.WorkspaceImageState.Set(nil)
}

// UnsetWorkspaceImageState ensures that no value is present for WorkspaceImageState, not even an explicit nil
func (o *AwsEdcImage) UnsetWorkspaceImageState() {
	o.WorkspaceImageState.Unset()
}

// GetApplicationList returns the ApplicationList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcImage) GetApplicationList() []AwsEdcAmiImportApplications {
	if o == nil {
		var ret []AwsEdcAmiImportApplications
		return ret
	}
	return o.ApplicationList
}

// GetApplicationListOk returns a tuple with the ApplicationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcImage) GetApplicationListOk() ([]AwsEdcAmiImportApplications, bool) {
	if o == nil || IsNil(o.ApplicationList) {
		return nil, false
	}
	return o.ApplicationList, true
}

// HasApplicationList returns a boolean if a field has been set.
func (o *AwsEdcImage) HasApplicationList() bool {
	if o != nil && IsNil(o.ApplicationList) {
		return true
	}

	return false
}

// SetApplicationList gets a reference to the given []AwsEdcAmiImportApplications and assigns it to the ApplicationList field.
func (o *AwsEdcImage) SetApplicationList(v []AwsEdcAmiImportApplications) {
	o.ApplicationList = v
}

func (o AwsEdcImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedImage, errImage := json.Marshal(o.Image)
	if errImage != nil {
		return map[string]interface{}{}, errImage
	}
	errImage = json.Unmarshal([]byte(serializedImage), &toSerialize)
	if errImage != nil {
		return map[string]interface{}{}, errImage
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.IngestionProcess.IsSet() {
		toSerialize["ingestionProcess"] = o.IngestionProcess.Get()
	}
	if o.WorkspaceImageTenancy.IsSet() {
		toSerialize["workspaceImageTenancy"] = o.WorkspaceImageTenancy.Get()
	}
	if o.WorkspaceImageState.IsSet() {
		toSerialize["workspaceImageState"] = o.WorkspaceImageState.Get()
	}
	if o.ApplicationList != nil {
		toSerialize["applicationList"] = o.ApplicationList
	}
	return toSerialize, nil
}

type NullableAwsEdcImage struct {
	value *AwsEdcImage
	isSet bool
}

func (v NullableAwsEdcImage) Get() *AwsEdcImage {
	return v.value
}

func (v *NullableAwsEdcImage) Set(val *AwsEdcImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcImage(val *AwsEdcImage) *NullableAwsEdcImage {
	return &NullableAwsEdcImage{value: val, isSet: true}
}

func (v NullableAwsEdcImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


