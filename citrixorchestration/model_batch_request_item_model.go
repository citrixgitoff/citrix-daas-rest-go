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

// checks if the BatchRequestItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequestItemModel{}

// BatchRequestItemModel Request object for a single item in a batch request.
type BatchRequestItemModel struct {
	// Reference.  Required, and must be unique within all items in a single batch request.
	Reference string `json:"Reference"`
	// HTTP method of the request endpoint.  (GET, HEAD, PUT, PATCH, POST, DELETE)
	Method NullableString `json:"Method,omitempty"`
	// Relative URL of the request endpoint vs. the site root. Must start with the API version i.e. `\"v1/{customerid}/{siteid}/...\"`
	RelativeUrl string `json:"RelativeUrl"`
	// List of request headers.
	Headers []NameValueStringPairModel `json:"Headers,omitempty"`
	// Request body.  Optional for PUT, PATCH, and POST. Cannot be specified for GET, HEAD, and DELETE.
	Body NullableString `json:"Body,omitempty"`
}

// NewBatchRequestItemModel instantiates a new BatchRequestItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequestItemModel(reference string, relativeUrl string) *BatchRequestItemModel {
	this := BatchRequestItemModel{}
	this.Reference = reference
	var method string = "GET"
	this.Method = *NewNullableString(&method)
	this.RelativeUrl = relativeUrl
	return &this
}

// NewBatchRequestItemModelWithDefaults instantiates a new BatchRequestItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestItemModelWithDefaults() *BatchRequestItemModel {
	this := BatchRequestItemModel{}
	var method string = "GET"
	this.Method = *NewNullableString(&method)
	return &this
}

// GetReference returns the Reference field value
func (o *BatchRequestItemModel) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *BatchRequestItemModel) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *BatchRequestItemModel) SetReference(v string) {
	o.Reference = v
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestItemModel) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestItemModel) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *BatchRequestItemModel) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *BatchRequestItemModel) SetMethod(v string) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *BatchRequestItemModel) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *BatchRequestItemModel) UnsetMethod() {
	o.Method.Unset()
}

// GetRelativeUrl returns the RelativeUrl field value
func (o *BatchRequestItemModel) GetRelativeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativeUrl
}

// GetRelativeUrlOk returns a tuple with the RelativeUrl field value
// and a boolean to check if the value has been set.
func (o *BatchRequestItemModel) GetRelativeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativeUrl, true
}

// SetRelativeUrl sets field value
func (o *BatchRequestItemModel) SetRelativeUrl(v string) {
	o.RelativeUrl = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestItemModel) GetHeaders() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestItemModel) GetHeadersOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BatchRequestItemModel) HasHeaders() bool {
	if o != nil && IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NameValueStringPairModel and assigns it to the Headers field.
func (o *BatchRequestItemModel) SetHeaders(v []NameValueStringPairModel) {
	o.Headers = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestItemModel) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestItemModel) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *BatchRequestItemModel) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *BatchRequestItemModel) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *BatchRequestItemModel) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *BatchRequestItemModel) UnsetBody() {
	o.Body.Unset()
}

func (o BatchRequestItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequestItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Reference"] = o.Reference
	if o.Method.IsSet() {
		toSerialize["Method"] = o.Method.Get()
	}
	toSerialize["RelativeUrl"] = o.RelativeUrl
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.Body.IsSet() {
		toSerialize["Body"] = o.Body.Get()
	}
	return toSerialize, nil
}

type NullableBatchRequestItemModel struct {
	value *BatchRequestItemModel
	isSet bool
}

func (v NullableBatchRequestItemModel) Get() *BatchRequestItemModel {
	return v.value
}

func (v *NullableBatchRequestItemModel) Set(val *BatchRequestItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequestItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequestItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequestItemModel(val *BatchRequestItemModel) *NullableBatchRequestItemModel {
	return &NullableBatchRequestItemModel{value: val, isSet: true}
}

func (v NullableBatchRequestItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequestItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


