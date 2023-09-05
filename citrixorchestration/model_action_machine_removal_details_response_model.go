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

// checks if the ActionMachineRemovalDetailsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionMachineRemovalDetailsResponseModel{}

// ActionMachineRemovalDetailsResponseModel struct for ActionMachineRemovalDetailsResponseModel
type ActionMachineRemovalDetailsResponseModel struct {
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
	// The virtual machine actions. 
	VMAction *string `json:"VMAction,omitempty"`
}

// NewActionMachineRemovalDetailsResponseModel instantiates a new ActionMachineRemovalDetailsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionMachineRemovalDetailsResponseModel() *ActionMachineRemovalDetailsResponseModel {
	this := ActionMachineRemovalDetailsResponseModel{}
	return &this
}

// NewActionMachineRemovalDetailsResponseModelWithDefaults instantiates a new ActionMachineRemovalDetailsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionMachineRemovalDetailsResponseModelWithDefaults() *ActionMachineRemovalDetailsResponseModel {
	this := ActionMachineRemovalDetailsResponseModel{}
	return &this
}

// GetAdAccountAction returns the AdAccountAction field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetAdAccountAction() string {
	if o == nil || IsNil(o.AdAccountAction) {
		var ret string
		return ret
	}
	return *o.AdAccountAction
}

// GetAdAccountActionOk returns a tuple with the AdAccountAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetAdAccountActionOk() (*string, bool) {
	if o == nil || IsNil(o.AdAccountAction) {
		return nil, false
	}
	return o.AdAccountAction, true
}

// HasAdAccountAction returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasAdAccountAction() bool {
	if o != nil && !IsNil(o.AdAccountAction) {
		return true
	}

	return false
}

// SetAdAccountAction gets a reference to the given string and assigns it to the AdAccountAction field.
func (o *ActionMachineRemovalDetailsResponseModel) SetAdAccountAction(v string) {
	o.AdAccountAction = &v
}

// GetSuccessfulAccounts returns the SuccessfulAccounts field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulAccounts() []string {
	if o == nil || IsNil(o.SuccessfulAccounts) {
		var ret []string
		return ret
	}
	return o.SuccessfulAccounts
}

// GetSuccessfulAccountsOk returns a tuple with the SuccessfulAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.SuccessfulAccounts) {
		return nil, false
	}
	return o.SuccessfulAccounts, true
}

// HasSuccessfulAccounts returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasSuccessfulAccounts() bool {
	if o != nil && !IsNil(o.SuccessfulAccounts) {
		return true
	}

	return false
}

// SetSuccessfulAccounts gets a reference to the given []string and assigns it to the SuccessfulAccounts field.
func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulAccounts(v []string) {
	o.SuccessfulAccounts = v
}

// GetSuccessfulMachines returns the SuccessfulMachines field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulMachines() []string {
	if o == nil || IsNil(o.SuccessfulMachines) {
		var ret []string
		return ret
	}
	return o.SuccessfulMachines
}

// GetSuccessfulMachinesOk returns a tuple with the SuccessfulMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulMachinesOk() ([]string, bool) {
	if o == nil || IsNil(o.SuccessfulMachines) {
		return nil, false
	}
	return o.SuccessfulMachines, true
}

// HasSuccessfulMachines returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasSuccessfulMachines() bool {
	if o != nil && !IsNil(o.SuccessfulMachines) {
		return true
	}

	return false
}

// SetSuccessfulMachines gets a reference to the given []string and assigns it to the SuccessfulMachines field.
func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulMachines(v []string) {
	o.SuccessfulMachines = v
}

// GetFailedMachines returns the FailedMachines field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetFailedMachines() []ActionFailedMachineOrAccountResponseModel {
	if o == nil || IsNil(o.FailedMachines) {
		var ret []ActionFailedMachineOrAccountResponseModel
		return ret
	}
	return o.FailedMachines
}

// GetFailedMachinesOk returns a tuple with the FailedMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetFailedMachinesOk() ([]ActionFailedMachineOrAccountResponseModel, bool) {
	if o == nil || IsNil(o.FailedMachines) {
		return nil, false
	}
	return o.FailedMachines, true
}

// HasFailedMachines returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasFailedMachines() bool {
	if o != nil && !IsNil(o.FailedMachines) {
		return true
	}

	return false
}

// SetFailedMachines gets a reference to the given []ActionFailedMachineOrAccountResponseModel and assigns it to the FailedMachines field.
func (o *ActionMachineRemovalDetailsResponseModel) SetFailedMachines(v []ActionFailedMachineOrAccountResponseModel) {
	o.FailedMachines = v
}

// GetFailedAccounts returns the FailedAccounts field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetFailedAccounts() []ActionFailedMachineOrAccountResponseModel {
	if o == nil || IsNil(o.FailedAccounts) {
		var ret []ActionFailedMachineOrAccountResponseModel
		return ret
	}
	return o.FailedAccounts
}

// GetFailedAccountsOk returns a tuple with the FailedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetFailedAccountsOk() ([]ActionFailedMachineOrAccountResponseModel, bool) {
	if o == nil || IsNil(o.FailedAccounts) {
		return nil, false
	}
	return o.FailedAccounts, true
}

// HasFailedAccounts returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasFailedAccounts() bool {
	if o != nil && !IsNil(o.FailedAccounts) {
		return true
	}

	return false
}

// SetFailedAccounts gets a reference to the given []ActionFailedMachineOrAccountResponseModel and assigns it to the FailedAccounts field.
func (o *ActionMachineRemovalDetailsResponseModel) SetFailedAccounts(v []ActionFailedMachineOrAccountResponseModel) {
	o.FailedAccounts = v
}

// GetVMAction returns the VMAction field value if set, zero value otherwise.
func (o *ActionMachineRemovalDetailsResponseModel) GetVMAction() string {
	if o == nil || IsNil(o.VMAction) {
		var ret string
		return ret
	}
	return *o.VMAction
}

// GetVMActionOk returns a tuple with the VMAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionMachineRemovalDetailsResponseModel) GetVMActionOk() (*string, bool) {
	if o == nil || IsNil(o.VMAction) {
		return nil, false
	}
	return o.VMAction, true
}

// HasVMAction returns a boolean if a field has been set.
func (o *ActionMachineRemovalDetailsResponseModel) HasVMAction() bool {
	if o != nil && !IsNil(o.VMAction) {
		return true
	}

	return false
}

// SetVMAction gets a reference to the given string and assigns it to the VMAction field.
func (o *ActionMachineRemovalDetailsResponseModel) SetVMAction(v string) {
	o.VMAction = &v
}

func (o ActionMachineRemovalDetailsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionMachineRemovalDetailsResponseModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.VMAction) {
		toSerialize["VMAction"] = o.VMAction
	}
	return toSerialize, nil
}

type NullableActionMachineRemovalDetailsResponseModel struct {
	value *ActionMachineRemovalDetailsResponseModel
	isSet bool
}

func (v NullableActionMachineRemovalDetailsResponseModel) Get() *ActionMachineRemovalDetailsResponseModel {
	return v.value
}

func (v *NullableActionMachineRemovalDetailsResponseModel) Set(val *ActionMachineRemovalDetailsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableActionMachineRemovalDetailsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableActionMachineRemovalDetailsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionMachineRemovalDetailsResponseModel(val *ActionMachineRemovalDetailsResponseModel) *NullableActionMachineRemovalDetailsResponseModel {
	return &NullableActionMachineRemovalDetailsResponseModel{value: val, isSet: true}
}

func (v NullableActionMachineRemovalDetailsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionMachineRemovalDetailsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


