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

// checks if the ImageVersionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageVersionResponseModel{}

// ImageVersionResponseModel struct for ImageVersionResponseModel
type ImageVersionResponseModel struct {
	// The version number associated with the image version.
	Version *string `json:"Version,omitempty"`
	// The Id of the image version.
	Id *string `json:"Id,omitempty"`
	// The MasterImageVM of the image version.
	MasterImagePath *string `json:"MasterImagePath,omitempty"`
	// The Description of the image version.
	Description *string `json:"Description,omitempty"`
	// Size of the VM's OS disk, in GB.
	DiskSizeGB *int32 `json:"DiskSizeGB,omitempty"`
	// The create time of the image version.
	CreateTime *string `json:"CreateTime,omitempty"`
	// The WriteBackCacheDiskSize of the image version.
	WriteBackCacheDiskSize *int32 `json:"WriteBackCacheDiskSize,omitempty"`
	// The WriteBackCacheMemorySize of the image version.
	WriteBackCacheMemorySize *int32 `json:"WriteBackCacheMemorySize,omitempty"`
	// The image status of the image version.
	ImageStatus *string `json:"ImageStatus,omitempty"`
	// The error info of the image version.
	Error *string `json:"Error,omitempty"`
	// The additional data of the image version.
	AdditionalData *map[string]string `json:"AdditionalData,omitempty"`
	ImageDefinition *ImageVersionResponseModelAllOfImageDefinition `json:"ImageDefinition,omitempty"`
	ImageScheme *ImageVersionDetailResponseModelAllOfImageScheme `json:"ImageScheme,omitempty"`
}

// NewImageVersionResponseModel instantiates a new ImageVersionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageVersionResponseModel() *ImageVersionResponseModel {
	this := ImageVersionResponseModel{}
	return &this
}

// NewImageVersionResponseModelWithDefaults instantiates a new ImageVersionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVersionResponseModelWithDefaults() *ImageVersionResponseModel {
	this := ImageVersionResponseModel{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ImageVersionResponseModel) SetVersion(v string) {
	o.Version = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImageVersionResponseModel) SetId(v string) {
	o.Id = &v
}

// GetMasterImagePath returns the MasterImagePath field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetMasterImagePath() string {
	if o == nil || IsNil(o.MasterImagePath) {
		var ret string
		return ret
	}
	return *o.MasterImagePath
}

// GetMasterImagePathOk returns a tuple with the MasterImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetMasterImagePathOk() (*string, bool) {
	if o == nil || IsNil(o.MasterImagePath) {
		return nil, false
	}
	return o.MasterImagePath, true
}

// HasMasterImagePath returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasMasterImagePath() bool {
	if o != nil && !IsNil(o.MasterImagePath) {
		return true
	}

	return false
}

// SetMasterImagePath gets a reference to the given string and assigns it to the MasterImagePath field.
func (o *ImageVersionResponseModel) SetMasterImagePath(v string) {
	o.MasterImagePath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImageVersionResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetDiskSizeGB() int32 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret int32
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetDiskSizeGBOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}
	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given int32 and assigns it to the DiskSizeGB field.
func (o *ImageVersionResponseModel) SetDiskSizeGB(v int32) {
	o.DiskSizeGB = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ImageVersionResponseModel) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetWriteBackCacheDiskSize() int32 {
	if o == nil || IsNil(o.WriteBackCacheDiskSize) {
		var ret int32
		return ret
	}
	return *o.WriteBackCacheDiskSize
}

// GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.WriteBackCacheDiskSize) {
		return nil, false
	}
	return o.WriteBackCacheDiskSize, true
}

// HasWriteBackCacheDiskSize returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasWriteBackCacheDiskSize() bool {
	if o != nil && !IsNil(o.WriteBackCacheDiskSize) {
		return true
	}

	return false
}

// SetWriteBackCacheDiskSize gets a reference to the given int32 and assigns it to the WriteBackCacheDiskSize field.
func (o *ImageVersionResponseModel) SetWriteBackCacheDiskSize(v int32) {
	o.WriteBackCacheDiskSize = &v
}

// GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetWriteBackCacheMemorySize() int32 {
	if o == nil || IsNil(o.WriteBackCacheMemorySize) {
		var ret int32
		return ret
	}
	return *o.WriteBackCacheMemorySize
}

// GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool) {
	if o == nil || IsNil(o.WriteBackCacheMemorySize) {
		return nil, false
	}
	return o.WriteBackCacheMemorySize, true
}

// HasWriteBackCacheMemorySize returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasWriteBackCacheMemorySize() bool {
	if o != nil && !IsNil(o.WriteBackCacheMemorySize) {
		return true
	}

	return false
}

// SetWriteBackCacheMemorySize gets a reference to the given int32 and assigns it to the WriteBackCacheMemorySize field.
func (o *ImageVersionResponseModel) SetWriteBackCacheMemorySize(v int32) {
	o.WriteBackCacheMemorySize = &v
}

// GetImageStatus returns the ImageStatus field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetImageStatus() string {
	if o == nil || IsNil(o.ImageStatus) {
		var ret string
		return ret
	}
	return *o.ImageStatus
}

// GetImageStatusOk returns a tuple with the ImageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetImageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ImageStatus) {
		return nil, false
	}
	return o.ImageStatus, true
}

// HasImageStatus returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasImageStatus() bool {
	if o != nil && !IsNil(o.ImageStatus) {
		return true
	}

	return false
}

// SetImageStatus gets a reference to the given string and assigns it to the ImageStatus field.
func (o *ImageVersionResponseModel) SetImageStatus(v string) {
	o.ImageStatus = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ImageVersionResponseModel) SetError(v string) {
	o.Error = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetAdditionalData() map[string]string {
	if o == nil || IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ImageVersionResponseModel) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetImageDefinition returns the ImageDefinition field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetImageDefinition() ImageVersionResponseModelAllOfImageDefinition {
	if o == nil || IsNil(o.ImageDefinition) {
		var ret ImageVersionResponseModelAllOfImageDefinition
		return ret
	}
	return *o.ImageDefinition
}

// GetImageDefinitionOk returns a tuple with the ImageDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetImageDefinitionOk() (*ImageVersionResponseModelAllOfImageDefinition, bool) {
	if o == nil || IsNil(o.ImageDefinition) {
		return nil, false
	}
	return o.ImageDefinition, true
}

// HasImageDefinition returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasImageDefinition() bool {
	if o != nil && !IsNil(o.ImageDefinition) {
		return true
	}

	return false
}

// SetImageDefinition gets a reference to the given ImageVersionResponseModelAllOfImageDefinition and assigns it to the ImageDefinition field.
func (o *ImageVersionResponseModel) SetImageDefinition(v ImageVersionResponseModelAllOfImageDefinition) {
	o.ImageDefinition = &v
}

// GetImageScheme returns the ImageScheme field value if set, zero value otherwise.
func (o *ImageVersionResponseModel) GetImageScheme() ImageVersionDetailResponseModelAllOfImageScheme {
	if o == nil || IsNil(o.ImageScheme) {
		var ret ImageVersionDetailResponseModelAllOfImageScheme
		return ret
	}
	return *o.ImageScheme
}

// GetImageSchemeOk returns a tuple with the ImageScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionResponseModel) GetImageSchemeOk() (*ImageVersionDetailResponseModelAllOfImageScheme, bool) {
	if o == nil || IsNil(o.ImageScheme) {
		return nil, false
	}
	return o.ImageScheme, true
}

// HasImageScheme returns a boolean if a field has been set.
func (o *ImageVersionResponseModel) HasImageScheme() bool {
	if o != nil && !IsNil(o.ImageScheme) {
		return true
	}

	return false
}

// SetImageScheme gets a reference to the given ImageVersionDetailResponseModelAllOfImageScheme and assigns it to the ImageScheme field.
func (o *ImageVersionResponseModel) SetImageScheme(v ImageVersionDetailResponseModelAllOfImageScheme) {
	o.ImageScheme = &v
}

func (o ImageVersionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageVersionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.MasterImagePath) {
		toSerialize["MasterImagePath"] = o.MasterImagePath
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DiskSizeGB) {
		toSerialize["DiskSizeGB"] = o.DiskSizeGB
	}
	if !IsNil(o.CreateTime) {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if !IsNil(o.WriteBackCacheDiskSize) {
		toSerialize["WriteBackCacheDiskSize"] = o.WriteBackCacheDiskSize
	}
	if !IsNil(o.WriteBackCacheMemorySize) {
		toSerialize["WriteBackCacheMemorySize"] = o.WriteBackCacheMemorySize
	}
	if !IsNil(o.ImageStatus) {
		toSerialize["ImageStatus"] = o.ImageStatus
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	if !IsNil(o.AdditionalData) {
		toSerialize["AdditionalData"] = o.AdditionalData
	}
	if !IsNil(o.ImageDefinition) {
		toSerialize["ImageDefinition"] = o.ImageDefinition
	}
	if !IsNil(o.ImageScheme) {
		toSerialize["ImageScheme"] = o.ImageScheme
	}
	return toSerialize, nil
}

type NullableImageVersionResponseModel struct {
	value *ImageVersionResponseModel
	isSet bool
}

func (v NullableImageVersionResponseModel) Get() *ImageVersionResponseModel {
	return v.value
}

func (v *NullableImageVersionResponseModel) Set(val *ImageVersionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionResponseModel(val *ImageVersionResponseModel) *NullableImageVersionResponseModel {
	return &NullableImageVersionResponseModel{value: val, isSet: true}
}

func (v NullableImageVersionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


