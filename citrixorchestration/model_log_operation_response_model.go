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

// checks if the LogOperationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogOperationResponseModel{}

// LogOperationResponseModel Log operation response model
type LogOperationResponseModel struct {
	// ID of the logged operation.
	Id string `json:"Id"`
	// IP address of the admin machine from which the operation was performed.
	AdminMachineIP NullableString `json:"AdminMachineIP,omitempty"`
	// Time when the operation ended. If the operation is incomplete, will be null.
	EndTime NullableString `json:"EndTime,omitempty"`
	// Formatted time when the operation ended. If the operation is incomplete, will be null. RFC 3339 compatible format.
	FormattedEndTime NullableString `json:"FormattedEndTime,omitempty"`
	// Indicates whether the operation completed successfully.  If the operation is incomplete, will be null.
	IsSuccessful NullableBool `json:"IsSuccessful,omitempty"`
	OperationType LogOperationType `json:"OperationType"`
	// Operation parameters. It is a list of parameter names.
	Parameters []NameValueStringPairModel `json:"Parameters"`
	// Source of the operation.
	Source string `json:"Source"`
	// Time when the operation started.
	StartTime string `json:"StartTime"`
	// Formatted time when the operation started. RFC 3339 compatible format.
	FormattedStartTime string `json:"FormattedStartTime"`
	// The type(s) of object which were the target of the configuration change. For example, \"Session\" or \"Machine\".
	TargetTypes []string `json:"TargetTypes"`
	// Human-readable description of the change.
	Text string `json:"Text"`
	// User who performed the change.
	User NullableString `json:"User,omitempty"`
}

// NewLogOperationResponseModel instantiates a new LogOperationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogOperationResponseModel(id string, operationType LogOperationType, parameters []NameValueStringPairModel, source string, startTime string, formattedStartTime string, targetTypes []string, text string) *LogOperationResponseModel {
	this := LogOperationResponseModel{}
	this.Id = id
	this.OperationType = operationType
	this.Parameters = parameters
	this.Source = source
	this.StartTime = startTime
	this.FormattedStartTime = formattedStartTime
	this.TargetTypes = targetTypes
	this.Text = text
	return &this
}

// NewLogOperationResponseModelWithDefaults instantiates a new LogOperationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogOperationResponseModelWithDefaults() *LogOperationResponseModel {
	this := LogOperationResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *LogOperationResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogOperationResponseModel) SetId(v string) {
	o.Id = v
}

// GetAdminMachineIP returns the AdminMachineIP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationResponseModel) GetAdminMachineIP() string {
	if o == nil || IsNil(o.AdminMachineIP.Get()) {
		var ret string
		return ret
	}
	return *o.AdminMachineIP.Get()
}

// GetAdminMachineIPOk returns a tuple with the AdminMachineIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationResponseModel) GetAdminMachineIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdminMachineIP.Get(), o.AdminMachineIP.IsSet()
}

// HasAdminMachineIP returns a boolean if a field has been set.
func (o *LogOperationResponseModel) HasAdminMachineIP() bool {
	if o != nil && o.AdminMachineIP.IsSet() {
		return true
	}

	return false
}

// SetAdminMachineIP gets a reference to the given NullableString and assigns it to the AdminMachineIP field.
func (o *LogOperationResponseModel) SetAdminMachineIP(v string) {
	o.AdminMachineIP.Set(&v)
}
// SetAdminMachineIPNil sets the value for AdminMachineIP to be an explicit nil
func (o *LogOperationResponseModel) SetAdminMachineIPNil() {
	o.AdminMachineIP.Set(nil)
}

// UnsetAdminMachineIP ensures that no value is present for AdminMachineIP, not even an explicit nil
func (o *LogOperationResponseModel) UnsetAdminMachineIP() {
	o.AdminMachineIP.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationResponseModel) GetEndTime() string {
	if o == nil || IsNil(o.EndTime.Get()) {
		var ret string
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationResponseModel) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *LogOperationResponseModel) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableString and assigns it to the EndTime field.
func (o *LogOperationResponseModel) SetEndTime(v string) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *LogOperationResponseModel) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *LogOperationResponseModel) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetFormattedEndTime returns the FormattedEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationResponseModel) GetFormattedEndTime() string {
	if o == nil || IsNil(o.FormattedEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.FormattedEndTime.Get()
}

// GetFormattedEndTimeOk returns a tuple with the FormattedEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationResponseModel) GetFormattedEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormattedEndTime.Get(), o.FormattedEndTime.IsSet()
}

// HasFormattedEndTime returns a boolean if a field has been set.
func (o *LogOperationResponseModel) HasFormattedEndTime() bool {
	if o != nil && o.FormattedEndTime.IsSet() {
		return true
	}

	return false
}

// SetFormattedEndTime gets a reference to the given NullableString and assigns it to the FormattedEndTime field.
func (o *LogOperationResponseModel) SetFormattedEndTime(v string) {
	o.FormattedEndTime.Set(&v)
}
// SetFormattedEndTimeNil sets the value for FormattedEndTime to be an explicit nil
func (o *LogOperationResponseModel) SetFormattedEndTimeNil() {
	o.FormattedEndTime.Set(nil)
}

// UnsetFormattedEndTime ensures that no value is present for FormattedEndTime, not even an explicit nil
func (o *LogOperationResponseModel) UnsetFormattedEndTime() {
	o.FormattedEndTime.Unset()
}

// GetIsSuccessful returns the IsSuccessful field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationResponseModel) GetIsSuccessful() bool {
	if o == nil || IsNil(o.IsSuccessful.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSuccessful.Get()
}

// GetIsSuccessfulOk returns a tuple with the IsSuccessful field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationResponseModel) GetIsSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSuccessful.Get(), o.IsSuccessful.IsSet()
}

// HasIsSuccessful returns a boolean if a field has been set.
func (o *LogOperationResponseModel) HasIsSuccessful() bool {
	if o != nil && o.IsSuccessful.IsSet() {
		return true
	}

	return false
}

// SetIsSuccessful gets a reference to the given NullableBool and assigns it to the IsSuccessful field.
func (o *LogOperationResponseModel) SetIsSuccessful(v bool) {
	o.IsSuccessful.Set(&v)
}
// SetIsSuccessfulNil sets the value for IsSuccessful to be an explicit nil
func (o *LogOperationResponseModel) SetIsSuccessfulNil() {
	o.IsSuccessful.Set(nil)
}

// UnsetIsSuccessful ensures that no value is present for IsSuccessful, not even an explicit nil
func (o *LogOperationResponseModel) UnsetIsSuccessful() {
	o.IsSuccessful.Unset()
}

// GetOperationType returns the OperationType field value
func (o *LogOperationResponseModel) GetOperationType() LogOperationType {
	if o == nil {
		var ret LogOperationType
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetOperationTypeOk() (*LogOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *LogOperationResponseModel) SetOperationType(v LogOperationType) {
	o.OperationType = v
}

// GetParameters returns the Parameters field value
func (o *LogOperationResponseModel) GetParameters() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetParametersOk() ([]NameValueStringPairModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *LogOperationResponseModel) SetParameters(v []NameValueStringPairModel) {
	o.Parameters = v
}

// GetSource returns the Source field value
func (o *LogOperationResponseModel) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *LogOperationResponseModel) SetSource(v string) {
	o.Source = v
}

// GetStartTime returns the StartTime field value
func (o *LogOperationResponseModel) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *LogOperationResponseModel) SetStartTime(v string) {
	o.StartTime = v
}

// GetFormattedStartTime returns the FormattedStartTime field value
func (o *LogOperationResponseModel) GetFormattedStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedStartTime
}

// GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetFormattedStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormattedStartTime, true
}

// SetFormattedStartTime sets field value
func (o *LogOperationResponseModel) SetFormattedStartTime(v string) {
	o.FormattedStartTime = v
}

// GetTargetTypes returns the TargetTypes field value
func (o *LogOperationResponseModel) GetTargetTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetTypes
}

// GetTargetTypesOk returns a tuple with the TargetTypes field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetTargetTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetTypes, true
}

// SetTargetTypes sets field value
func (o *LogOperationResponseModel) SetTargetTypes(v []string) {
	o.TargetTypes = v
}

// GetText returns the Text field value
func (o *LogOperationResponseModel) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *LogOperationResponseModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *LogOperationResponseModel) SetText(v string) {
	o.Text = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationResponseModel) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationResponseModel) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *LogOperationResponseModel) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *LogOperationResponseModel) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *LogOperationResponseModel) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *LogOperationResponseModel) UnsetUser() {
	o.User.Unset()
}

func (o LogOperationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogOperationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if o.AdminMachineIP.IsSet() {
		toSerialize["AdminMachineIP"] = o.AdminMachineIP.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["EndTime"] = o.EndTime.Get()
	}
	if o.FormattedEndTime.IsSet() {
		toSerialize["FormattedEndTime"] = o.FormattedEndTime.Get()
	}
	if o.IsSuccessful.IsSet() {
		toSerialize["IsSuccessful"] = o.IsSuccessful.Get()
	}
	toSerialize["OperationType"] = o.OperationType
	toSerialize["Parameters"] = o.Parameters
	toSerialize["Source"] = o.Source
	toSerialize["StartTime"] = o.StartTime
	toSerialize["FormattedStartTime"] = o.FormattedStartTime
	toSerialize["TargetTypes"] = o.TargetTypes
	toSerialize["Text"] = o.Text
	if o.User.IsSet() {
		toSerialize["User"] = o.User.Get()
	}
	return toSerialize, nil
}

type NullableLogOperationResponseModel struct {
	value *LogOperationResponseModel
	isSet bool
}

func (v NullableLogOperationResponseModel) Get() *LogOperationResponseModel {
	return v.value
}

func (v *NullableLogOperationResponseModel) Set(val *LogOperationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLogOperationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLogOperationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogOperationResponseModel(val *LogOperationResponseModel) *NullableLogOperationResponseModel {
	return &NullableLogOperationResponseModel{value: val, isSet: true}
}

func (v NullableLogOperationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogOperationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


