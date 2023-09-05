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

// checks if the OperatingSystemInfoResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatingSystemInfoResponseModel{}

// OperatingSystemInfoResponseModel The operating sysytem information
type OperatingSystemInfoResponseModel struct {
	OperatingSystemType *OperatingSystemType `json:"OperatingSystemType,omitempty"`
	// OS major version.
	MajorVersion *int32 `json:"MajorVersion,omitempty"`
	// OS minor version.
	MinorVersion *int32 `json:"MinorVersion,omitempty"`
	// OS build number.
	BuildNumber *int32 `json:"BuildNumber,omitempty"`
	ProductType *ProductType `json:"ProductType,omitempty"`
	SuiteMask *SuiteMask `json:"SuiteMask,omitempty"`
	// OS release Id. It is read from registry and available for Windows 10 and higher.
	ReleaseId *string `json:"ReleaseId,omitempty"`
	// OS update build revision number. It is read from registry and available for Windows 10 and higher.
	UpdateBuildRevision *string `json:"UpdateBuildRevision,omitempty"`
	// OS display version. It is read from registry and available for Windows 10 and higher.
	DisplayVersion *string `json:"DisplayVersion,omitempty"`
	SessionSupport *SessionSupport `json:"SessionSupport,omitempty"`
}

// NewOperatingSystemInfoResponseModel instantiates a new OperatingSystemInfoResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystemInfoResponseModel() *OperatingSystemInfoResponseModel {
	this := OperatingSystemInfoResponseModel{}
	return &this
}

// NewOperatingSystemInfoResponseModelWithDefaults instantiates a new OperatingSystemInfoResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemInfoResponseModelWithDefaults() *OperatingSystemInfoResponseModel {
	this := OperatingSystemInfoResponseModel{}
	return &this
}

// GetOperatingSystemType returns the OperatingSystemType field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetOperatingSystemType() OperatingSystemType {
	if o == nil || IsNil(o.OperatingSystemType) {
		var ret OperatingSystemType
		return ret
	}
	return *o.OperatingSystemType
}

// GetOperatingSystemTypeOk returns a tuple with the OperatingSystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetOperatingSystemTypeOk() (*OperatingSystemType, bool) {
	if o == nil || IsNil(o.OperatingSystemType) {
		return nil, false
	}
	return o.OperatingSystemType, true
}

// HasOperatingSystemType returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasOperatingSystemType() bool {
	if o != nil && !IsNil(o.OperatingSystemType) {
		return true
	}

	return false
}

// SetOperatingSystemType gets a reference to the given OperatingSystemType and assigns it to the OperatingSystemType field.
func (o *OperatingSystemInfoResponseModel) SetOperatingSystemType(v OperatingSystemType) {
	o.OperatingSystemType = &v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetMajorVersion() int32 {
	if o == nil || IsNil(o.MajorVersion) {
		var ret int32
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetMajorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MajorVersion) {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasMajorVersion() bool {
	if o != nil && !IsNil(o.MajorVersion) {
		return true
	}

	return false
}

// SetMajorVersion gets a reference to the given int32 and assigns it to the MajorVersion field.
func (o *OperatingSystemInfoResponseModel) SetMajorVersion(v int32) {
	o.MajorVersion = &v
}

// GetMinorVersion returns the MinorVersion field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetMinorVersion() int32 {
	if o == nil || IsNil(o.MinorVersion) {
		var ret int32
		return ret
	}
	return *o.MinorVersion
}

// GetMinorVersionOk returns a tuple with the MinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetMinorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinorVersion) {
		return nil, false
	}
	return o.MinorVersion, true
}

// HasMinorVersion returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasMinorVersion() bool {
	if o != nil && !IsNil(o.MinorVersion) {
		return true
	}

	return false
}

// SetMinorVersion gets a reference to the given int32 and assigns it to the MinorVersion field.
func (o *OperatingSystemInfoResponseModel) SetMinorVersion(v int32) {
	o.MinorVersion = &v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetBuildNumber() int32 {
	if o == nil || IsNil(o.BuildNumber) {
		var ret int32
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetBuildNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BuildNumber) {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasBuildNumber() bool {
	if o != nil && !IsNil(o.BuildNumber) {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given int32 and assigns it to the BuildNumber field.
func (o *OperatingSystemInfoResponseModel) SetBuildNumber(v int32) {
	o.BuildNumber = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetProductType() ProductType {
	if o == nil || IsNil(o.ProductType) {
		var ret ProductType
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetProductTypeOk() (*ProductType, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given ProductType and assigns it to the ProductType field.
func (o *OperatingSystemInfoResponseModel) SetProductType(v ProductType) {
	o.ProductType = &v
}

// GetSuiteMask returns the SuiteMask field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetSuiteMask() SuiteMask {
	if o == nil || IsNil(o.SuiteMask) {
		var ret SuiteMask
		return ret
	}
	return *o.SuiteMask
}

// GetSuiteMaskOk returns a tuple with the SuiteMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetSuiteMaskOk() (*SuiteMask, bool) {
	if o == nil || IsNil(o.SuiteMask) {
		return nil, false
	}
	return o.SuiteMask, true
}

// HasSuiteMask returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasSuiteMask() bool {
	if o != nil && !IsNil(o.SuiteMask) {
		return true
	}

	return false
}

// SetSuiteMask gets a reference to the given SuiteMask and assigns it to the SuiteMask field.
func (o *OperatingSystemInfoResponseModel) SetSuiteMask(v SuiteMask) {
	o.SuiteMask = &v
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetReleaseId() string {
	if o == nil || IsNil(o.ReleaseId) {
		var ret string
		return ret
	}
	return *o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetReleaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseId) {
		return nil, false
	}
	return o.ReleaseId, true
}

// HasReleaseId returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasReleaseId() bool {
	if o != nil && !IsNil(o.ReleaseId) {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given string and assigns it to the ReleaseId field.
func (o *OperatingSystemInfoResponseModel) SetReleaseId(v string) {
	o.ReleaseId = &v
}

// GetUpdateBuildRevision returns the UpdateBuildRevision field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetUpdateBuildRevision() string {
	if o == nil || IsNil(o.UpdateBuildRevision) {
		var ret string
		return ret
	}
	return *o.UpdateBuildRevision
}

// GetUpdateBuildRevisionOk returns a tuple with the UpdateBuildRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetUpdateBuildRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateBuildRevision) {
		return nil, false
	}
	return o.UpdateBuildRevision, true
}

// HasUpdateBuildRevision returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasUpdateBuildRevision() bool {
	if o != nil && !IsNil(o.UpdateBuildRevision) {
		return true
	}

	return false
}

// SetUpdateBuildRevision gets a reference to the given string and assigns it to the UpdateBuildRevision field.
func (o *OperatingSystemInfoResponseModel) SetUpdateBuildRevision(v string) {
	o.UpdateBuildRevision = &v
}

// GetDisplayVersion returns the DisplayVersion field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetDisplayVersion() string {
	if o == nil || IsNil(o.DisplayVersion) {
		var ret string
		return ret
	}
	return *o.DisplayVersion
}

// GetDisplayVersionOk returns a tuple with the DisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetDisplayVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayVersion) {
		return nil, false
	}
	return o.DisplayVersion, true
}

// HasDisplayVersion returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasDisplayVersion() bool {
	if o != nil && !IsNil(o.DisplayVersion) {
		return true
	}

	return false
}

// SetDisplayVersion gets a reference to the given string and assigns it to the DisplayVersion field.
func (o *OperatingSystemInfoResponseModel) SetDisplayVersion(v string) {
	o.DisplayVersion = &v
}

// GetSessionSupport returns the SessionSupport field value if set, zero value otherwise.
func (o *OperatingSystemInfoResponseModel) GetSessionSupport() SessionSupport {
	if o == nil || IsNil(o.SessionSupport) {
		var ret SessionSupport
		return ret
	}
	return *o.SessionSupport
}

// GetSessionSupportOk returns a tuple with the SessionSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystemInfoResponseModel) GetSessionSupportOk() (*SessionSupport, bool) {
	if o == nil || IsNil(o.SessionSupport) {
		return nil, false
	}
	return o.SessionSupport, true
}

// HasSessionSupport returns a boolean if a field has been set.
func (o *OperatingSystemInfoResponseModel) HasSessionSupport() bool {
	if o != nil && !IsNil(o.SessionSupport) {
		return true
	}

	return false
}

// SetSessionSupport gets a reference to the given SessionSupport and assigns it to the SessionSupport field.
func (o *OperatingSystemInfoResponseModel) SetSessionSupport(v SessionSupport) {
	o.SessionSupport = &v
}

func (o OperatingSystemInfoResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatingSystemInfoResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperatingSystemType) {
		toSerialize["OperatingSystemType"] = o.OperatingSystemType
	}
	if !IsNil(o.MajorVersion) {
		toSerialize["MajorVersion"] = o.MajorVersion
	}
	if !IsNil(o.MinorVersion) {
		toSerialize["MinorVersion"] = o.MinorVersion
	}
	if !IsNil(o.BuildNumber) {
		toSerialize["BuildNumber"] = o.BuildNumber
	}
	if !IsNil(o.ProductType) {
		toSerialize["ProductType"] = o.ProductType
	}
	if !IsNil(o.SuiteMask) {
		toSerialize["SuiteMask"] = o.SuiteMask
	}
	if !IsNil(o.ReleaseId) {
		toSerialize["ReleaseId"] = o.ReleaseId
	}
	if !IsNil(o.UpdateBuildRevision) {
		toSerialize["UpdateBuildRevision"] = o.UpdateBuildRevision
	}
	if !IsNil(o.DisplayVersion) {
		toSerialize["DisplayVersion"] = o.DisplayVersion
	}
	if !IsNil(o.SessionSupport) {
		toSerialize["SessionSupport"] = o.SessionSupport
	}
	return toSerialize, nil
}

type NullableOperatingSystemInfoResponseModel struct {
	value *OperatingSystemInfoResponseModel
	isSet bool
}

func (v NullableOperatingSystemInfoResponseModel) Get() *OperatingSystemInfoResponseModel {
	return v.value
}

func (v *NullableOperatingSystemInfoResponseModel) Set(val *OperatingSystemInfoResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystemInfoResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystemInfoResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystemInfoResponseModel(val *OperatingSystemInfoResponseModel) *NullableOperatingSystemInfoResponseModel {
	return &NullableOperatingSystemInfoResponseModel{value: val, isSet: true}
}

func (v NullableOperatingSystemInfoResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystemInfoResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


