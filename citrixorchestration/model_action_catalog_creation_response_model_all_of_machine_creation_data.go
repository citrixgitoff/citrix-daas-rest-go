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

// checks if the ActionCatalogCreationResponseModelAllOfMachineCreationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionCatalogCreationResponseModelAllOfMachineCreationData{}

// ActionCatalogCreationResponseModelAllOfMachineCreationData The machine creation data.
type ActionCatalogCreationResponseModelAllOfMachineCreationData struct {
	// Active directory account action.
	AdAccountAction *string `json:"AdAccountAction,omitempty"`
	// Successful accounts.
	SuccessfulAccounts []string `json:"SuccessfulAccounts,omitempty"`
	// Successful machines. 
	SuccessfulMachines []string `json:"SuccessfulMachines,omitempty"`
	// Failed machines and Action errors.
	FailedMachines []ActionFailedMachineOrAccountResponseModel `json:"FailedMachines,omitempty"`
	// Failed accounts and action errors.
	FailedAccounts []ActionFailedMachineOrAccountResponseModel `json:"FailedAccounts,omitempty"`
	// The machine requested
	MachineRequested *int32 `json:"MachineRequested,omitempty"`
}

// NewActionCatalogCreationResponseModelAllOfMachineCreationData instantiates a new ActionCatalogCreationResponseModelAllOfMachineCreationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionCatalogCreationResponseModelAllOfMachineCreationData() *ActionCatalogCreationResponseModelAllOfMachineCreationData {
	this := ActionCatalogCreationResponseModelAllOfMachineCreationData{}
	return &this
}

// NewActionCatalogCreationResponseModelAllOfMachineCreationDataWithDefaults instantiates a new ActionCatalogCreationResponseModelAllOfMachineCreationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionCatalogCreationResponseModelAllOfMachineCreationDataWithDefaults() *ActionCatalogCreationResponseModelAllOfMachineCreationData {
	this := ActionCatalogCreationResponseModelAllOfMachineCreationData{}
	return &this
}

// GetAdAccountAction returns the AdAccountAction field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetAdAccountAction() string {
	if o == nil || IsNil(o.AdAccountAction) {
		var ret string
		return ret
	}
	return *o.AdAccountAction
}

// GetAdAccountActionOk returns a tuple with the AdAccountAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetAdAccountActionOk() (*string, bool) {
	if o == nil || IsNil(o.AdAccountAction) {
		return nil, false
	}
	return o.AdAccountAction, true
}

// HasAdAccountAction returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasAdAccountAction() bool {
	if o != nil && !IsNil(o.AdAccountAction) {
		return true
	}

	return false
}

// SetAdAccountAction gets a reference to the given string and assigns it to the AdAccountAction field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetAdAccountAction(v string) {
	o.AdAccountAction = &v
}

// GetSuccessfulAccounts returns the SuccessfulAccounts field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetSuccessfulAccounts() []string {
	if o == nil || IsNil(o.SuccessfulAccounts) {
		var ret []string
		return ret
	}
	return o.SuccessfulAccounts
}

// GetSuccessfulAccountsOk returns a tuple with the SuccessfulAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetSuccessfulAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.SuccessfulAccounts) {
		return nil, false
	}
	return o.SuccessfulAccounts, true
}

// HasSuccessfulAccounts returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasSuccessfulAccounts() bool {
	if o != nil && !IsNil(o.SuccessfulAccounts) {
		return true
	}

	return false
}

// SetSuccessfulAccounts gets a reference to the given []string and assigns it to the SuccessfulAccounts field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetSuccessfulAccounts(v []string) {
	o.SuccessfulAccounts = v
}

// GetSuccessfulMachines returns the SuccessfulMachines field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetSuccessfulMachines() []string {
	if o == nil || IsNil(o.SuccessfulMachines) {
		var ret []string
		return ret
	}
	return o.SuccessfulMachines
}

// GetSuccessfulMachinesOk returns a tuple with the SuccessfulMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetSuccessfulMachinesOk() ([]string, bool) {
	if o == nil || IsNil(o.SuccessfulMachines) {
		return nil, false
	}
	return o.SuccessfulMachines, true
}

// HasSuccessfulMachines returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasSuccessfulMachines() bool {
	if o != nil && !IsNil(o.SuccessfulMachines) {
		return true
	}

	return false
}

// SetSuccessfulMachines gets a reference to the given []string and assigns it to the SuccessfulMachines field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetSuccessfulMachines(v []string) {
	o.SuccessfulMachines = v
}

// GetFailedMachines returns the FailedMachines field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetFailedMachines() []ActionFailedMachineOrAccountResponseModel {
	if o == nil || IsNil(o.FailedMachines) {
		var ret []ActionFailedMachineOrAccountResponseModel
		return ret
	}
	return o.FailedMachines
}

// GetFailedMachinesOk returns a tuple with the FailedMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetFailedMachinesOk() ([]ActionFailedMachineOrAccountResponseModel, bool) {
	if o == nil || IsNil(o.FailedMachines) {
		return nil, false
	}
	return o.FailedMachines, true
}

// HasFailedMachines returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasFailedMachines() bool {
	if o != nil && !IsNil(o.FailedMachines) {
		return true
	}

	return false
}

// SetFailedMachines gets a reference to the given []ActionFailedMachineOrAccountResponseModel and assigns it to the FailedMachines field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetFailedMachines(v []ActionFailedMachineOrAccountResponseModel) {
	o.FailedMachines = v
}

// GetFailedAccounts returns the FailedAccounts field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetFailedAccounts() []ActionFailedMachineOrAccountResponseModel {
	if o == nil || IsNil(o.FailedAccounts) {
		var ret []ActionFailedMachineOrAccountResponseModel
		return ret
	}
	return o.FailedAccounts
}

// GetFailedAccountsOk returns a tuple with the FailedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetFailedAccountsOk() ([]ActionFailedMachineOrAccountResponseModel, bool) {
	if o == nil || IsNil(o.FailedAccounts) {
		return nil, false
	}
	return o.FailedAccounts, true
}

// HasFailedAccounts returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasFailedAccounts() bool {
	if o != nil && !IsNil(o.FailedAccounts) {
		return true
	}

	return false
}

// SetFailedAccounts gets a reference to the given []ActionFailedMachineOrAccountResponseModel and assigns it to the FailedAccounts field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetFailedAccounts(v []ActionFailedMachineOrAccountResponseModel) {
	o.FailedAccounts = v
}

// GetMachineRequested returns the MachineRequested field value if set, zero value otherwise.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetMachineRequested() int32 {
	if o == nil || IsNil(o.MachineRequested) {
		var ret int32
		return ret
	}
	return *o.MachineRequested
}

// GetMachineRequestedOk returns a tuple with the MachineRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) GetMachineRequestedOk() (*int32, bool) {
	if o == nil || IsNil(o.MachineRequested) {
		return nil, false
	}
	return o.MachineRequested, true
}

// HasMachineRequested returns a boolean if a field has been set.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) HasMachineRequested() bool {
	if o != nil && !IsNil(o.MachineRequested) {
		return true
	}

	return false
}

// SetMachineRequested gets a reference to the given int32 and assigns it to the MachineRequested field.
func (o *ActionCatalogCreationResponseModelAllOfMachineCreationData) SetMachineRequested(v int32) {
	o.MachineRequested = &v
}

func (o ActionCatalogCreationResponseModelAllOfMachineCreationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionCatalogCreationResponseModelAllOfMachineCreationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdAccountAction) {
		toSerialize["AdAccountAction"] = o.AdAccountAction
	}
	if !IsNil(o.SuccessfulAccounts) {
		toSerialize["SuccessfulAccounts"] = o.SuccessfulAccounts
	}
	if !IsNil(o.SuccessfulMachines) {
		toSerialize["SuccessfulMachines"] = o.SuccessfulMachines
	}
	if !IsNil(o.FailedMachines) {
		toSerialize["FailedMachines"] = o.FailedMachines
	}
	if !IsNil(o.FailedAccounts) {
		toSerialize["FailedAccounts"] = o.FailedAccounts
	}
	if !IsNil(o.MachineRequested) {
		toSerialize["MachineRequested"] = o.MachineRequested
	}
	return toSerialize, nil
}

type NullableActionCatalogCreationResponseModelAllOfMachineCreationData struct {
	value *ActionCatalogCreationResponseModelAllOfMachineCreationData
	isSet bool
}

func (v NullableActionCatalogCreationResponseModelAllOfMachineCreationData) Get() *ActionCatalogCreationResponseModelAllOfMachineCreationData {
	return v.value
}

func (v *NullableActionCatalogCreationResponseModelAllOfMachineCreationData) Set(val *ActionCatalogCreationResponseModelAllOfMachineCreationData) {
	v.value = val
	v.isSet = true
}

func (v NullableActionCatalogCreationResponseModelAllOfMachineCreationData) IsSet() bool {
	return v.isSet
}

func (v *NullableActionCatalogCreationResponseModelAllOfMachineCreationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionCatalogCreationResponseModelAllOfMachineCreationData(val *ActionCatalogCreationResponseModelAllOfMachineCreationData) *NullableActionCatalogCreationResponseModelAllOfMachineCreationData {
	return &NullableActionCatalogCreationResponseModelAllOfMachineCreationData{value: val, isSet: true}
}

func (v NullableActionCatalogCreationResponseModelAllOfMachineCreationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionCatalogCreationResponseModelAllOfMachineCreationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


