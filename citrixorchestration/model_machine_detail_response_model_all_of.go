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

// checks if the MachineDetailResponseModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineDetailResponseModelAllOf{}

// MachineDetailResponseModelAllOf Response model for machine details.
type MachineDetailResponseModelAllOf struct {
	// The name of the endpoint client device that the machine has been assigned to.
	AssignedClientName *string `json:"AssignedClientName,omitempty"`
	// The IP address of the endpoint client device that the machine has been assigned to.
	AssignedIPAddress *string `json:"AssignedIPAddress,omitempty"`
	// Site-wide unique name identifying associated desktop to other components (for example StoreFront). This is typically non-null only for machines backing assigned private desktops.
	BrowserName *string `json:"BrowserName,omitempty"`
	ColorDepth *ColorDepth `json:"ColorDepth,omitempty"`
	// The machine's icon that is displayed in Receiver.
	IconId *string `json:"IconId,omitempty"`
	// Indicates if machine is reserved for special use, for example for AppDisk preparation. A reserved machine cannot be a member of a delivery group.
	IsReserved bool `json:"IsReserved"`
	// Gives the last reported individual load indexes that were used in the calculation of the LoadIndex value. Note that the LoadIndex value may have been subsequently adjusted due to session brokering operations. This value is only set when SessionSupport is equal to MultiSession.
	LoadIndexes []int32 `json:"LoadIndexes,omitempty"`
	// Flag indicating whether SecureICA is required or not when starting a session on the machine.
	SecureIcaRequired *bool `json:"SecureIcaRequired,omitempty"`
	// Number of established sessions on this machine.  When SessionSupport is equal to MultiSession, this excludes established sessions which have not yet completed their logon processing.
	SessionsEstablished int32 `json:"SessionsEstablished"`
	// Number of pending (brokered but not yet established) sessions on this machine.  When SessionSupport is equal to MultiSession, this also includes established sessions which have not yet completed their logon processing.
	SessionsPending int32 `json:"SessionsPending"`
	// StoreFront servers to use for Receiver when it is hosted on the machine.
	StoreFrontServersForHostedReceiver []StoreFrontServerResponseModel `json:"StoreFrontServersForHostedReceiver,omitempty"`
	VMToolsState *VMToolsState `json:"VMToolsState,omitempty"`
	// Failure reason of power action.
	FailureReason *string `json:"FailureReason,omitempty"`
}

// NewMachineDetailResponseModelAllOf instantiates a new MachineDetailResponseModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineDetailResponseModelAllOf(isReserved bool, sessionsEstablished int32, sessionsPending int32) *MachineDetailResponseModelAllOf {
	this := MachineDetailResponseModelAllOf{}
	this.IsReserved = isReserved
	this.SessionsEstablished = sessionsEstablished
	this.SessionsPending = sessionsPending
	return &this
}

// NewMachineDetailResponseModelAllOfWithDefaults instantiates a new MachineDetailResponseModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineDetailResponseModelAllOfWithDefaults() *MachineDetailResponseModelAllOf {
	this := MachineDetailResponseModelAllOf{}
	return &this
}

// GetAssignedClientName returns the AssignedClientName field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetAssignedClientName() string {
	if o == nil || IsNil(o.AssignedClientName) {
		var ret string
		return ret
	}
	return *o.AssignedClientName
}

// GetAssignedClientNameOk returns a tuple with the AssignedClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetAssignedClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedClientName) {
		return nil, false
	}
	return o.AssignedClientName, true
}

// HasAssignedClientName returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasAssignedClientName() bool {
	if o != nil && !IsNil(o.AssignedClientName) {
		return true
	}

	return false
}

// SetAssignedClientName gets a reference to the given string and assigns it to the AssignedClientName field.
func (o *MachineDetailResponseModelAllOf) SetAssignedClientName(v string) {
	o.AssignedClientName = &v
}

// GetAssignedIPAddress returns the AssignedIPAddress field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetAssignedIPAddress() string {
	if o == nil || IsNil(o.AssignedIPAddress) {
		var ret string
		return ret
	}
	return *o.AssignedIPAddress
}

// GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetAssignedIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedIPAddress) {
		return nil, false
	}
	return o.AssignedIPAddress, true
}

// HasAssignedIPAddress returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasAssignedIPAddress() bool {
	if o != nil && !IsNil(o.AssignedIPAddress) {
		return true
	}

	return false
}

// SetAssignedIPAddress gets a reference to the given string and assigns it to the AssignedIPAddress field.
func (o *MachineDetailResponseModelAllOf) SetAssignedIPAddress(v string) {
	o.AssignedIPAddress = &v
}

// GetBrowserName returns the BrowserName field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetBrowserName() string {
	if o == nil || IsNil(o.BrowserName) {
		var ret string
		return ret
	}
	return *o.BrowserName
}

// GetBrowserNameOk returns a tuple with the BrowserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetBrowserNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserName) {
		return nil, false
	}
	return o.BrowserName, true
}

// HasBrowserName returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasBrowserName() bool {
	if o != nil && !IsNil(o.BrowserName) {
		return true
	}

	return false
}

// SetBrowserName gets a reference to the given string and assigns it to the BrowserName field.
func (o *MachineDetailResponseModelAllOf) SetBrowserName(v string) {
	o.BrowserName = &v
}

// GetColorDepth returns the ColorDepth field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetColorDepth() ColorDepth {
	if o == nil || IsNil(o.ColorDepth) {
		var ret ColorDepth
		return ret
	}
	return *o.ColorDepth
}

// GetColorDepthOk returns a tuple with the ColorDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetColorDepthOk() (*ColorDepth, bool) {
	if o == nil || IsNil(o.ColorDepth) {
		return nil, false
	}
	return o.ColorDepth, true
}

// HasColorDepth returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasColorDepth() bool {
	if o != nil && !IsNil(o.ColorDepth) {
		return true
	}

	return false
}

// SetColorDepth gets a reference to the given ColorDepth and assigns it to the ColorDepth field.
func (o *MachineDetailResponseModelAllOf) SetColorDepth(v ColorDepth) {
	o.ColorDepth = &v
}

// GetIconId returns the IconId field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetIconId() string {
	if o == nil || IsNil(o.IconId) {
		var ret string
		return ret
	}
	return *o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetIconIdOk() (*string, bool) {
	if o == nil || IsNil(o.IconId) {
		return nil, false
	}
	return o.IconId, true
}

// HasIconId returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasIconId() bool {
	if o != nil && !IsNil(o.IconId) {
		return true
	}

	return false
}

// SetIconId gets a reference to the given string and assigns it to the IconId field.
func (o *MachineDetailResponseModelAllOf) SetIconId(v string) {
	o.IconId = &v
}

// GetIsReserved returns the IsReserved field value
func (o *MachineDetailResponseModelAllOf) GetIsReserved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReserved
}

// GetIsReservedOk returns a tuple with the IsReserved field value
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetIsReservedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReserved, true
}

// SetIsReserved sets field value
func (o *MachineDetailResponseModelAllOf) SetIsReserved(v bool) {
	o.IsReserved = v
}

// GetLoadIndexes returns the LoadIndexes field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetLoadIndexes() []int32 {
	if o == nil || IsNil(o.LoadIndexes) {
		var ret []int32
		return ret
	}
	return o.LoadIndexes
}

// GetLoadIndexesOk returns a tuple with the LoadIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetLoadIndexesOk() ([]int32, bool) {
	if o == nil || IsNil(o.LoadIndexes) {
		return nil, false
	}
	return o.LoadIndexes, true
}

// HasLoadIndexes returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasLoadIndexes() bool {
	if o != nil && !IsNil(o.LoadIndexes) {
		return true
	}

	return false
}

// SetLoadIndexes gets a reference to the given []int32 and assigns it to the LoadIndexes field.
func (o *MachineDetailResponseModelAllOf) SetLoadIndexes(v []int32) {
	o.LoadIndexes = v
}

// GetSecureIcaRequired returns the SecureIcaRequired field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetSecureIcaRequired() bool {
	if o == nil || IsNil(o.SecureIcaRequired) {
		var ret bool
		return ret
	}
	return *o.SecureIcaRequired
}

// GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetSecureIcaRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureIcaRequired) {
		return nil, false
	}
	return o.SecureIcaRequired, true
}

// HasSecureIcaRequired returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasSecureIcaRequired() bool {
	if o != nil && !IsNil(o.SecureIcaRequired) {
		return true
	}

	return false
}

// SetSecureIcaRequired gets a reference to the given bool and assigns it to the SecureIcaRequired field.
func (o *MachineDetailResponseModelAllOf) SetSecureIcaRequired(v bool) {
	o.SecureIcaRequired = &v
}

// GetSessionsEstablished returns the SessionsEstablished field value
func (o *MachineDetailResponseModelAllOf) GetSessionsEstablished() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessionsEstablished
}

// GetSessionsEstablishedOk returns a tuple with the SessionsEstablished field value
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetSessionsEstablishedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsEstablished, true
}

// SetSessionsEstablished sets field value
func (o *MachineDetailResponseModelAllOf) SetSessionsEstablished(v int32) {
	o.SessionsEstablished = v
}

// GetSessionsPending returns the SessionsPending field value
func (o *MachineDetailResponseModelAllOf) GetSessionsPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessionsPending
}

// GetSessionsPendingOk returns a tuple with the SessionsPending field value
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetSessionsPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsPending, true
}

// SetSessionsPending sets field value
func (o *MachineDetailResponseModelAllOf) SetSessionsPending(v int32) {
	o.SessionsPending = v
}

// GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiver() []StoreFrontServerResponseModel {
	if o == nil || IsNil(o.StoreFrontServersForHostedReceiver) {
		var ret []StoreFrontServerResponseModel
		return ret
	}
	return o.StoreFrontServersForHostedReceiver
}

// GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiverOk() ([]StoreFrontServerResponseModel, bool) {
	if o == nil || IsNil(o.StoreFrontServersForHostedReceiver) {
		return nil, false
	}
	return o.StoreFrontServersForHostedReceiver, true
}

// HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasStoreFrontServersForHostedReceiver() bool {
	if o != nil && !IsNil(o.StoreFrontServersForHostedReceiver) {
		return true
	}

	return false
}

// SetStoreFrontServersForHostedReceiver gets a reference to the given []StoreFrontServerResponseModel and assigns it to the StoreFrontServersForHostedReceiver field.
func (o *MachineDetailResponseModelAllOf) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerResponseModel) {
	o.StoreFrontServersForHostedReceiver = v
}

// GetVMToolsState returns the VMToolsState field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetVMToolsState() VMToolsState {
	if o == nil || IsNil(o.VMToolsState) {
		var ret VMToolsState
		return ret
	}
	return *o.VMToolsState
}

// GetVMToolsStateOk returns a tuple with the VMToolsState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetVMToolsStateOk() (*VMToolsState, bool) {
	if o == nil || IsNil(o.VMToolsState) {
		return nil, false
	}
	return o.VMToolsState, true
}

// HasVMToolsState returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasVMToolsState() bool {
	if o != nil && !IsNil(o.VMToolsState) {
		return true
	}

	return false
}

// SetVMToolsState gets a reference to the given VMToolsState and assigns it to the VMToolsState field.
func (o *MachineDetailResponseModelAllOf) SetVMToolsState(v VMToolsState) {
	o.VMToolsState = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *MachineDetailResponseModelAllOf) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetailResponseModelAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MachineDetailResponseModelAllOf) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *MachineDetailResponseModelAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

func (o MachineDetailResponseModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineDetailResponseModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignedClientName) {
		toSerialize["AssignedClientName"] = o.AssignedClientName
	}
	if !IsNil(o.AssignedIPAddress) {
		toSerialize["AssignedIPAddress"] = o.AssignedIPAddress
	}
	if !IsNil(o.BrowserName) {
		toSerialize["BrowserName"] = o.BrowserName
	}
	if !IsNil(o.ColorDepth) {
		toSerialize["ColorDepth"] = o.ColorDepth
	}
	if !IsNil(o.IconId) {
		toSerialize["IconId"] = o.IconId
	}
	toSerialize["IsReserved"] = o.IsReserved
	if !IsNil(o.LoadIndexes) {
		toSerialize["LoadIndexes"] = o.LoadIndexes
	}
	if !IsNil(o.SecureIcaRequired) {
		toSerialize["SecureIcaRequired"] = o.SecureIcaRequired
	}
	toSerialize["SessionsEstablished"] = o.SessionsEstablished
	toSerialize["SessionsPending"] = o.SessionsPending
	if !IsNil(o.StoreFrontServersForHostedReceiver) {
		toSerialize["StoreFrontServersForHostedReceiver"] = o.StoreFrontServersForHostedReceiver
	}
	if !IsNil(o.VMToolsState) {
		toSerialize["VMToolsState"] = o.VMToolsState
	}
	if !IsNil(o.FailureReason) {
		toSerialize["FailureReason"] = o.FailureReason
	}
	return toSerialize, nil
}

type NullableMachineDetailResponseModelAllOf struct {
	value *MachineDetailResponseModelAllOf
	isSet bool
}

func (v NullableMachineDetailResponseModelAllOf) Get() *MachineDetailResponseModelAllOf {
	return v.value
}

func (v *NullableMachineDetailResponseModelAllOf) Set(val *MachineDetailResponseModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineDetailResponseModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineDetailResponseModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineDetailResponseModelAllOf(val *MachineDetailResponseModelAllOf) *NullableMachineDetailResponseModelAllOf {
	return &NullableMachineDetailResponseModelAllOf{value: val, isSet: true}
}

func (v NullableMachineDetailResponseModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineDetailResponseModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


