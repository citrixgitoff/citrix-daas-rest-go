/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"time"
)

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image Base class for image
type Image struct {
	AccountType AccountType `json:"accountType"`
	// The ID of the image
	ImageId NullableString `json:"imageId,omitempty"`
	// The name of the image
	Name NullableString `json:"name,omitempty"`
	// The description of the image
	Description NullableString `json:"description,omitempty"`
	// The notes of the image
	Notes NullableString `json:"notes,omitempty"`
	OperatingSystem NullableOperatingSystemType `json:"operatingSystem,omitempty"`
	AssociatedDeployments []AssociatedDeployment `json:"associatedDeployments,omitempty"`
	// Last time the status was modified
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(accountType AccountType) *Image {
	this := Image{}
	this.AccountType = accountType
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *Image) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *Image) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *Image) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *Image) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *Image) SetImageId(v string) {
	o.ImageId.Set(&v)
}
// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *Image) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *Image) UnsetImageId() {
	o.ImageId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Image) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Image) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Image) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Image) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Image) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Image) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Image) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Image) UnsetDescription() {
	o.Description.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *Image) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *Image) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *Image) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *Image) UnsetNotes() {
	o.Notes.Unset()
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetOperatingSystem() OperatingSystemType {
	if o == nil || IsNil(o.OperatingSystem.Get()) {
		var ret OperatingSystemType
		return ret
	}
	return *o.OperatingSystem.Get()
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetOperatingSystemOk() (*OperatingSystemType, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystem.Get(), o.OperatingSystem.IsSet()
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *Image) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given NullableOperatingSystemType and assigns it to the OperatingSystem field.
func (o *Image) SetOperatingSystem(v OperatingSystemType) {
	o.OperatingSystem.Set(&v)
}
// SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil
func (o *Image) SetOperatingSystemNil() {
	o.OperatingSystem.Set(nil)
}

// UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
func (o *Image) UnsetOperatingSystem() {
	o.OperatingSystem.Unset()
}

// GetAssociatedDeployments returns the AssociatedDeployments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetAssociatedDeployments() []AssociatedDeployment {
	if o == nil {
		var ret []AssociatedDeployment
		return ret
	}
	return o.AssociatedDeployments
}

// GetAssociatedDeploymentsOk returns a tuple with the AssociatedDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetAssociatedDeploymentsOk() ([]AssociatedDeployment, bool) {
	if o == nil || IsNil(o.AssociatedDeployments) {
		return nil, false
	}
	return o.AssociatedDeployments, true
}

// HasAssociatedDeployments returns a boolean if a field has been set.
func (o *Image) HasAssociatedDeployments() bool {
	if o != nil && IsNil(o.AssociatedDeployments) {
		return true
	}

	return false
}

// SetAssociatedDeployments gets a reference to the given []AssociatedDeployment and assigns it to the AssociatedDeployments field.
func (o *Image) SetAssociatedDeployments(v []AssociatedDeployment) {
	o.AssociatedDeployments = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *Image) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *Image) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *Image) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.OperatingSystem.IsSet() {
		toSerialize["operatingSystem"] = o.OperatingSystem.Get()
	}
	if o.AssociatedDeployments != nil {
		toSerialize["associatedDeployments"] = o.AssociatedDeployments
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


