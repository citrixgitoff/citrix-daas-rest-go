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

// checks if the TestReportResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestReportResultModel{}

// TestReportResultModel The result of running a test.  The result is composed of a test id, scope and collection of 1 or more component results
type TestReportResultModel struct {
	// The localized display name of the test.
	TestName NullableString `json:"TestName,omitempty"`
	// The description of the localized end user description of the test.
	TestDescription NullableString `json:"TestDescription,omitempty"`
	// Gets or sets the test start time. If the test has not yet been run, this may be null.
	TestStartTime NullableString `json:"TestStartTime,omitempty"`
	// Gets or sets the formatted test start time. If the test has not yet been run, this may be null. RFC 3339 compatible format.
	FormattedTestStartTime NullableString `json:"FormattedTestStartTime,omitempty"`
	// Gets or sets the test end time. If the test has not yet been run, this may be null.
	TestEndTime NullableString `json:"TestEndTime,omitempty"`
	// Gets or sets the formatted test end time. If the test has not yet been run, this may be null. RFC 3339 compatible format.
	FormattedTestEndTime NullableString `json:"FormattedTestEndTime,omitempty"`
	// Gets or sets the test service target, generally the name of the service the test is being run against.
	TestServiceTarget NullableString `json:"TestServiceTarget,omitempty"`
	// Gets or sets the overall status of this test run.  If the test was run against more than one controller, it is the result of or'ing together of the statuses of each component result.
	TestComponentStatus NullableString `json:"TestComponentStatus,omitempty"`
	TestScope *TestScope `json:"TestScope,omitempty"`
	// Gets or sets the list of component results. Each machine that the test is run against will be represented by a component result and their statuses aggregated into TestComponentStatus
	TestComponents []TestComponentResultModel `json:"TestComponents,omitempty"`
}

// NewTestReportResultModel instantiates a new TestReportResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestReportResultModel() *TestReportResultModel {
	this := TestReportResultModel{}
	return &this
}

// NewTestReportResultModelWithDefaults instantiates a new TestReportResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestReportResultModelWithDefaults() *TestReportResultModel {
	this := TestReportResultModel{}
	return &this
}

// GetTestName returns the TestName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestName() string {
	if o == nil || IsNil(o.TestName.Get()) {
		var ret string
		return ret
	}
	return *o.TestName.Get()
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestName.Get(), o.TestName.IsSet()
}

// HasTestName returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestName() bool {
	if o != nil && o.TestName.IsSet() {
		return true
	}

	return false
}

// SetTestName gets a reference to the given NullableString and assigns it to the TestName field.
func (o *TestReportResultModel) SetTestName(v string) {
	o.TestName.Set(&v)
}
// SetTestNameNil sets the value for TestName to be an explicit nil
func (o *TestReportResultModel) SetTestNameNil() {
	o.TestName.Set(nil)
}

// UnsetTestName ensures that no value is present for TestName, not even an explicit nil
func (o *TestReportResultModel) UnsetTestName() {
	o.TestName.Unset()
}

// GetTestDescription returns the TestDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestDescription() string {
	if o == nil || IsNil(o.TestDescription.Get()) {
		var ret string
		return ret
	}
	return *o.TestDescription.Get()
}

// GetTestDescriptionOk returns a tuple with the TestDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestDescription.Get(), o.TestDescription.IsSet()
}

// HasTestDescription returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestDescription() bool {
	if o != nil && o.TestDescription.IsSet() {
		return true
	}

	return false
}

// SetTestDescription gets a reference to the given NullableString and assigns it to the TestDescription field.
func (o *TestReportResultModel) SetTestDescription(v string) {
	o.TestDescription.Set(&v)
}
// SetTestDescriptionNil sets the value for TestDescription to be an explicit nil
func (o *TestReportResultModel) SetTestDescriptionNil() {
	o.TestDescription.Set(nil)
}

// UnsetTestDescription ensures that no value is present for TestDescription, not even an explicit nil
func (o *TestReportResultModel) UnsetTestDescription() {
	o.TestDescription.Unset()
}

// GetTestStartTime returns the TestStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestStartTime() string {
	if o == nil || IsNil(o.TestStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.TestStartTime.Get()
}

// GetTestStartTimeOk returns a tuple with the TestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestStartTime.Get(), o.TestStartTime.IsSet()
}

// HasTestStartTime returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestStartTime() bool {
	if o != nil && o.TestStartTime.IsSet() {
		return true
	}

	return false
}

// SetTestStartTime gets a reference to the given NullableString and assigns it to the TestStartTime field.
func (o *TestReportResultModel) SetTestStartTime(v string) {
	o.TestStartTime.Set(&v)
}
// SetTestStartTimeNil sets the value for TestStartTime to be an explicit nil
func (o *TestReportResultModel) SetTestStartTimeNil() {
	o.TestStartTime.Set(nil)
}

// UnsetTestStartTime ensures that no value is present for TestStartTime, not even an explicit nil
func (o *TestReportResultModel) UnsetTestStartTime() {
	o.TestStartTime.Unset()
}

// GetFormattedTestStartTime returns the FormattedTestStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetFormattedTestStartTime() string {
	if o == nil || IsNil(o.FormattedTestStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.FormattedTestStartTime.Get()
}

// GetFormattedTestStartTimeOk returns a tuple with the FormattedTestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetFormattedTestStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormattedTestStartTime.Get(), o.FormattedTestStartTime.IsSet()
}

// HasFormattedTestStartTime returns a boolean if a field has been set.
func (o *TestReportResultModel) HasFormattedTestStartTime() bool {
	if o != nil && o.FormattedTestStartTime.IsSet() {
		return true
	}

	return false
}

// SetFormattedTestStartTime gets a reference to the given NullableString and assigns it to the FormattedTestStartTime field.
func (o *TestReportResultModel) SetFormattedTestStartTime(v string) {
	o.FormattedTestStartTime.Set(&v)
}
// SetFormattedTestStartTimeNil sets the value for FormattedTestStartTime to be an explicit nil
func (o *TestReportResultModel) SetFormattedTestStartTimeNil() {
	o.FormattedTestStartTime.Set(nil)
}

// UnsetFormattedTestStartTime ensures that no value is present for FormattedTestStartTime, not even an explicit nil
func (o *TestReportResultModel) UnsetFormattedTestStartTime() {
	o.FormattedTestStartTime.Unset()
}

// GetTestEndTime returns the TestEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestEndTime() string {
	if o == nil || IsNil(o.TestEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.TestEndTime.Get()
}

// GetTestEndTimeOk returns a tuple with the TestEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestEndTime.Get(), o.TestEndTime.IsSet()
}

// HasTestEndTime returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestEndTime() bool {
	if o != nil && o.TestEndTime.IsSet() {
		return true
	}

	return false
}

// SetTestEndTime gets a reference to the given NullableString and assigns it to the TestEndTime field.
func (o *TestReportResultModel) SetTestEndTime(v string) {
	o.TestEndTime.Set(&v)
}
// SetTestEndTimeNil sets the value for TestEndTime to be an explicit nil
func (o *TestReportResultModel) SetTestEndTimeNil() {
	o.TestEndTime.Set(nil)
}

// UnsetTestEndTime ensures that no value is present for TestEndTime, not even an explicit nil
func (o *TestReportResultModel) UnsetTestEndTime() {
	o.TestEndTime.Unset()
}

// GetFormattedTestEndTime returns the FormattedTestEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetFormattedTestEndTime() string {
	if o == nil || IsNil(o.FormattedTestEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.FormattedTestEndTime.Get()
}

// GetFormattedTestEndTimeOk returns a tuple with the FormattedTestEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetFormattedTestEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormattedTestEndTime.Get(), o.FormattedTestEndTime.IsSet()
}

// HasFormattedTestEndTime returns a boolean if a field has been set.
func (o *TestReportResultModel) HasFormattedTestEndTime() bool {
	if o != nil && o.FormattedTestEndTime.IsSet() {
		return true
	}

	return false
}

// SetFormattedTestEndTime gets a reference to the given NullableString and assigns it to the FormattedTestEndTime field.
func (o *TestReportResultModel) SetFormattedTestEndTime(v string) {
	o.FormattedTestEndTime.Set(&v)
}
// SetFormattedTestEndTimeNil sets the value for FormattedTestEndTime to be an explicit nil
func (o *TestReportResultModel) SetFormattedTestEndTimeNil() {
	o.FormattedTestEndTime.Set(nil)
}

// UnsetFormattedTestEndTime ensures that no value is present for FormattedTestEndTime, not even an explicit nil
func (o *TestReportResultModel) UnsetFormattedTestEndTime() {
	o.FormattedTestEndTime.Unset()
}

// GetTestServiceTarget returns the TestServiceTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestServiceTarget() string {
	if o == nil || IsNil(o.TestServiceTarget.Get()) {
		var ret string
		return ret
	}
	return *o.TestServiceTarget.Get()
}

// GetTestServiceTargetOk returns a tuple with the TestServiceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestServiceTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestServiceTarget.Get(), o.TestServiceTarget.IsSet()
}

// HasTestServiceTarget returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestServiceTarget() bool {
	if o != nil && o.TestServiceTarget.IsSet() {
		return true
	}

	return false
}

// SetTestServiceTarget gets a reference to the given NullableString and assigns it to the TestServiceTarget field.
func (o *TestReportResultModel) SetTestServiceTarget(v string) {
	o.TestServiceTarget.Set(&v)
}
// SetTestServiceTargetNil sets the value for TestServiceTarget to be an explicit nil
func (o *TestReportResultModel) SetTestServiceTargetNil() {
	o.TestServiceTarget.Set(nil)
}

// UnsetTestServiceTarget ensures that no value is present for TestServiceTarget, not even an explicit nil
func (o *TestReportResultModel) UnsetTestServiceTarget() {
	o.TestServiceTarget.Unset()
}

// GetTestComponentStatus returns the TestComponentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestComponentStatus() string {
	if o == nil || IsNil(o.TestComponentStatus.Get()) {
		var ret string
		return ret
	}
	return *o.TestComponentStatus.Get()
}

// GetTestComponentStatusOk returns a tuple with the TestComponentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestComponentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestComponentStatus.Get(), o.TestComponentStatus.IsSet()
}

// HasTestComponentStatus returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestComponentStatus() bool {
	if o != nil && o.TestComponentStatus.IsSet() {
		return true
	}

	return false
}

// SetTestComponentStatus gets a reference to the given NullableString and assigns it to the TestComponentStatus field.
func (o *TestReportResultModel) SetTestComponentStatus(v string) {
	o.TestComponentStatus.Set(&v)
}
// SetTestComponentStatusNil sets the value for TestComponentStatus to be an explicit nil
func (o *TestReportResultModel) SetTestComponentStatusNil() {
	o.TestComponentStatus.Set(nil)
}

// UnsetTestComponentStatus ensures that no value is present for TestComponentStatus, not even an explicit nil
func (o *TestReportResultModel) UnsetTestComponentStatus() {
	o.TestComponentStatus.Unset()
}

// GetTestScope returns the TestScope field value if set, zero value otherwise.
func (o *TestReportResultModel) GetTestScope() TestScope {
	if o == nil || IsNil(o.TestScope) {
		var ret TestScope
		return ret
	}
	return *o.TestScope
}

// GetTestScopeOk returns a tuple with the TestScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestReportResultModel) GetTestScopeOk() (*TestScope, bool) {
	if o == nil || IsNil(o.TestScope) {
		return nil, false
	}
	return o.TestScope, true
}

// HasTestScope returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestScope() bool {
	if o != nil && !IsNil(o.TestScope) {
		return true
	}

	return false
}

// SetTestScope gets a reference to the given TestScope and assigns it to the TestScope field.
func (o *TestReportResultModel) SetTestScope(v TestScope) {
	o.TestScope = &v
}

// GetTestComponents returns the TestComponents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestReportResultModel) GetTestComponents() []TestComponentResultModel {
	if o == nil {
		var ret []TestComponentResultModel
		return ret
	}
	return o.TestComponents
}

// GetTestComponentsOk returns a tuple with the TestComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestReportResultModel) GetTestComponentsOk() ([]TestComponentResultModel, bool) {
	if o == nil || IsNil(o.TestComponents) {
		return nil, false
	}
	return o.TestComponents, true
}

// HasTestComponents returns a boolean if a field has been set.
func (o *TestReportResultModel) HasTestComponents() bool {
	if o != nil && IsNil(o.TestComponents) {
		return true
	}

	return false
}

// SetTestComponents gets a reference to the given []TestComponentResultModel and assigns it to the TestComponents field.
func (o *TestReportResultModel) SetTestComponents(v []TestComponentResultModel) {
	o.TestComponents = v
}

func (o TestReportResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestReportResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TestName.IsSet() {
		toSerialize["TestName"] = o.TestName.Get()
	}
	if o.TestDescription.IsSet() {
		toSerialize["TestDescription"] = o.TestDescription.Get()
	}
	if o.TestStartTime.IsSet() {
		toSerialize["TestStartTime"] = o.TestStartTime.Get()
	}
	if o.FormattedTestStartTime.IsSet() {
		toSerialize["FormattedTestStartTime"] = o.FormattedTestStartTime.Get()
	}
	if o.TestEndTime.IsSet() {
		toSerialize["TestEndTime"] = o.TestEndTime.Get()
	}
	if o.FormattedTestEndTime.IsSet() {
		toSerialize["FormattedTestEndTime"] = o.FormattedTestEndTime.Get()
	}
	if o.TestServiceTarget.IsSet() {
		toSerialize["TestServiceTarget"] = o.TestServiceTarget.Get()
	}
	if o.TestComponentStatus.IsSet() {
		toSerialize["TestComponentStatus"] = o.TestComponentStatus.Get()
	}
	if !IsNil(o.TestScope) {
		toSerialize["TestScope"] = o.TestScope
	}
	if o.TestComponents != nil {
		toSerialize["TestComponents"] = o.TestComponents
	}
	return toSerialize, nil
}

type NullableTestReportResultModel struct {
	value *TestReportResultModel
	isSet bool
}

func (v NullableTestReportResultModel) Get() *TestReportResultModel {
	return v.value
}

func (v *NullableTestReportResultModel) Set(val *TestReportResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestReportResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestReportResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestReportResultModel(val *TestReportResultModel) *NullableTestReportResultModel {
	return &NullableTestReportResultModel{value: val, isSet: true}
}

func (v NullableTestReportResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestReportResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


