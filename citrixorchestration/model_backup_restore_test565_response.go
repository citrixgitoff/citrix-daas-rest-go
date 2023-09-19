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

// checks if the BackupRestoreTest565Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreTest565Response{}

// BackupRestoreTest565Response Response for Backup/Restore Test result
type BackupRestoreTest565Response struct {
	// String Result 1             
	Result1 NullableString `json:"Result1,omitempty"`
	// String Result 2             
	Result2 NullableString `json:"Result2,omitempty"`
	// String Result 3             
	Result3 NullableString `json:"Result3,omitempty"`
	// Int Resujlt             
	ResultInt *int32 `json:"ResultInt,omitempty"`
	// Over all test result             
	TestResult *bool `json:"TestResult,omitempty"`
}

// NewBackupRestoreTest565Response instantiates a new BackupRestoreTest565Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreTest565Response() *BackupRestoreTest565Response {
	this := BackupRestoreTest565Response{}
	return &this
}

// NewBackupRestoreTest565ResponseWithDefaults instantiates a new BackupRestoreTest565Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreTest565ResponseWithDefaults() *BackupRestoreTest565Response {
	this := BackupRestoreTest565Response{}
	return &this
}

// GetResult1 returns the Result1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreTest565Response) GetResult1() string {
	if o == nil || IsNil(o.Result1.Get()) {
		var ret string
		return ret
	}
	return *o.Result1.Get()
}

// GetResult1Ok returns a tuple with the Result1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreTest565Response) GetResult1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result1.Get(), o.Result1.IsSet()
}

// HasResult1 returns a boolean if a field has been set.
func (o *BackupRestoreTest565Response) HasResult1() bool {
	if o != nil && o.Result1.IsSet() {
		return true
	}

	return false
}

// SetResult1 gets a reference to the given NullableString and assigns it to the Result1 field.
func (o *BackupRestoreTest565Response) SetResult1(v string) {
	o.Result1.Set(&v)
}
// SetResult1Nil sets the value for Result1 to be an explicit nil
func (o *BackupRestoreTest565Response) SetResult1Nil() {
	o.Result1.Set(nil)
}

// UnsetResult1 ensures that no value is present for Result1, not even an explicit nil
func (o *BackupRestoreTest565Response) UnsetResult1() {
	o.Result1.Unset()
}

// GetResult2 returns the Result2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreTest565Response) GetResult2() string {
	if o == nil || IsNil(o.Result2.Get()) {
		var ret string
		return ret
	}
	return *o.Result2.Get()
}

// GetResult2Ok returns a tuple with the Result2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreTest565Response) GetResult2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result2.Get(), o.Result2.IsSet()
}

// HasResult2 returns a boolean if a field has been set.
func (o *BackupRestoreTest565Response) HasResult2() bool {
	if o != nil && o.Result2.IsSet() {
		return true
	}

	return false
}

// SetResult2 gets a reference to the given NullableString and assigns it to the Result2 field.
func (o *BackupRestoreTest565Response) SetResult2(v string) {
	o.Result2.Set(&v)
}
// SetResult2Nil sets the value for Result2 to be an explicit nil
func (o *BackupRestoreTest565Response) SetResult2Nil() {
	o.Result2.Set(nil)
}

// UnsetResult2 ensures that no value is present for Result2, not even an explicit nil
func (o *BackupRestoreTest565Response) UnsetResult2() {
	o.Result2.Unset()
}

// GetResult3 returns the Result3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreTest565Response) GetResult3() string {
	if o == nil || IsNil(o.Result3.Get()) {
		var ret string
		return ret
	}
	return *o.Result3.Get()
}

// GetResult3Ok returns a tuple with the Result3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreTest565Response) GetResult3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result3.Get(), o.Result3.IsSet()
}

// HasResult3 returns a boolean if a field has been set.
func (o *BackupRestoreTest565Response) HasResult3() bool {
	if o != nil && o.Result3.IsSet() {
		return true
	}

	return false
}

// SetResult3 gets a reference to the given NullableString and assigns it to the Result3 field.
func (o *BackupRestoreTest565Response) SetResult3(v string) {
	o.Result3.Set(&v)
}
// SetResult3Nil sets the value for Result3 to be an explicit nil
func (o *BackupRestoreTest565Response) SetResult3Nil() {
	o.Result3.Set(nil)
}

// UnsetResult3 ensures that no value is present for Result3, not even an explicit nil
func (o *BackupRestoreTest565Response) UnsetResult3() {
	o.Result3.Unset()
}

// GetResultInt returns the ResultInt field value if set, zero value otherwise.
func (o *BackupRestoreTest565Response) GetResultInt() int32 {
	if o == nil || IsNil(o.ResultInt) {
		var ret int32
		return ret
	}
	return *o.ResultInt
}

// GetResultIntOk returns a tuple with the ResultInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreTest565Response) GetResultIntOk() (*int32, bool) {
	if o == nil || IsNil(o.ResultInt) {
		return nil, false
	}
	return o.ResultInt, true
}

// HasResultInt returns a boolean if a field has been set.
func (o *BackupRestoreTest565Response) HasResultInt() bool {
	if o != nil && !IsNil(o.ResultInt) {
		return true
	}

	return false
}

// SetResultInt gets a reference to the given int32 and assigns it to the ResultInt field.
func (o *BackupRestoreTest565Response) SetResultInt(v int32) {
	o.ResultInt = &v
}

// GetTestResult returns the TestResult field value if set, zero value otherwise.
func (o *BackupRestoreTest565Response) GetTestResult() bool {
	if o == nil || IsNil(o.TestResult) {
		var ret bool
		return ret
	}
	return *o.TestResult
}

// GetTestResultOk returns a tuple with the TestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreTest565Response) GetTestResultOk() (*bool, bool) {
	if o == nil || IsNil(o.TestResult) {
		return nil, false
	}
	return o.TestResult, true
}

// HasTestResult returns a boolean if a field has been set.
func (o *BackupRestoreTest565Response) HasTestResult() bool {
	if o != nil && !IsNil(o.TestResult) {
		return true
	}

	return false
}

// SetTestResult gets a reference to the given bool and assigns it to the TestResult field.
func (o *BackupRestoreTest565Response) SetTestResult(v bool) {
	o.TestResult = &v
}

func (o BackupRestoreTest565Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreTest565Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result1.IsSet() {
		toSerialize["Result1"] = o.Result1.Get()
	}
	if o.Result2.IsSet() {
		toSerialize["Result2"] = o.Result2.Get()
	}
	if o.Result3.IsSet() {
		toSerialize["Result3"] = o.Result3.Get()
	}
	if !IsNil(o.ResultInt) {
		toSerialize["ResultInt"] = o.ResultInt
	}
	if !IsNil(o.TestResult) {
		toSerialize["TestResult"] = o.TestResult
	}
	return toSerialize, nil
}

type NullableBackupRestoreTest565Response struct {
	value *BackupRestoreTest565Response
	isSet bool
}

func (v NullableBackupRestoreTest565Response) Get() *BackupRestoreTest565Response {
	return v.value
}

func (v *NullableBackupRestoreTest565Response) Set(val *BackupRestoreTest565Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreTest565Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreTest565Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreTest565Response(val *BackupRestoreTest565Response) *NullableBackupRestoreTest565Response {
	return &NullableBackupRestoreTest565Response{value: val, isSet: true}
}

func (v NullableBackupRestoreTest565Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreTest565Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


