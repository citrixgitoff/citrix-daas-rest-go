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

// checks if the ImageVersionSpecResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageVersionSpecResponseModel{}

// ImageVersionSpecResponseModel Response object for image version specification.
type ImageVersionSpecResponseModel struct {
	// The Id of the image version specification.
	Id string `json:"Id"`
	Context *ImageVersionSpecContextResponseModel `json:"Context,omitempty"`
	// The create time of the image version specification.
	CreationTime NullableString `json:"CreationTime,omitempty"`
	// The errors in this image version specification
	Errors []string `json:"Errors,omitempty"`
	ImageRuntimeEnvironment *ImageRuntimeEnvironmentResponseModel `json:"ImageRuntimeEnvironment,omitempty"`
	ImageVersionSpecStatus *ImageVersionSpecStatus `json:"ImageVersionSpecStatus,omitempty"`
	MasterImage *HypervisorResourceRefResponseModel `json:"MasterImage,omitempty"`
	PreparationType PreparationType `json:"PreparationType"`
	// Number of provisioning scheme created from this image version specification.
	ProvisioningSchemeCount *int32 `json:"ProvisioningSchemeCount,omitempty"`
	ResourcePool HypervisorResourcePoolRefResponseModel `json:"ResourcePool"`
	SourceImageVersionSpec *ImageVersionSpecRefResponseModel `json:"SourceImageVersionSpec,omitempty"`
	// The disk size
	DiskSize *int32 `json:"DiskSize,omitempty"`
	// List of warnings that are currently active on the image version specification, if any.  If there are no warnings this will not be specified.
	Warnings []ImageVersionSpecWarningResponseModel `json:"Warnings,omitempty"`
	Zone *RefResponseModel `json:"Zone,omitempty"`
}

// NewImageVersionSpecResponseModel instantiates a new ImageVersionSpecResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageVersionSpecResponseModel(id string, preparationType PreparationType, resourcePool HypervisorResourcePoolRefResponseModel) *ImageVersionSpecResponseModel {
	this := ImageVersionSpecResponseModel{}
	this.Id = id
	this.PreparationType = preparationType
	this.ResourcePool = resourcePool
	return &this
}

// NewImageVersionSpecResponseModelWithDefaults instantiates a new ImageVersionSpecResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVersionSpecResponseModelWithDefaults() *ImageVersionSpecResponseModel {
	this := ImageVersionSpecResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ImageVersionSpecResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageVersionSpecResponseModel) SetId(v string) {
	o.Id = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetContext() ImageVersionSpecContextResponseModel {
	if o == nil || IsNil(o.Context) {
		var ret ImageVersionSpecContextResponseModel
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetContextOk() (*ImageVersionSpecContextResponseModel, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given ImageVersionSpecContextResponseModel and assigns it to the Context field.
func (o *ImageVersionSpecResponseModel) SetContext(v ImageVersionSpecContextResponseModel) {
	o.Context = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResponseModel) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime.Get()) {
		var ret string
		return ret
	}
	return *o.CreationTime.Get()
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResponseModel) GetCreationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationTime.Get(), o.CreationTime.IsSet()
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasCreationTime() bool {
	if o != nil && o.CreationTime.IsSet() {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given NullableString and assigns it to the CreationTime field.
func (o *ImageVersionSpecResponseModel) SetCreationTime(v string) {
	o.CreationTime.Set(&v)
}
// SetCreationTimeNil sets the value for CreationTime to be an explicit nil
func (o *ImageVersionSpecResponseModel) SetCreationTimeNil() {
	o.CreationTime.Set(nil)
}

// UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
func (o *ImageVersionSpecResponseModel) UnsetCreationTime() {
	o.CreationTime.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResponseModel) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResponseModel) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ImageVersionSpecResponseModel) SetErrors(v []string) {
	o.Errors = v
}

// GetImageRuntimeEnvironment returns the ImageRuntimeEnvironment field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetImageRuntimeEnvironment() ImageRuntimeEnvironmentResponseModel {
	if o == nil || IsNil(o.ImageRuntimeEnvironment) {
		var ret ImageRuntimeEnvironmentResponseModel
		return ret
	}
	return *o.ImageRuntimeEnvironment
}

// GetImageRuntimeEnvironmentOk returns a tuple with the ImageRuntimeEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetImageRuntimeEnvironmentOk() (*ImageRuntimeEnvironmentResponseModel, bool) {
	if o == nil || IsNil(o.ImageRuntimeEnvironment) {
		return nil, false
	}
	return o.ImageRuntimeEnvironment, true
}

// HasImageRuntimeEnvironment returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasImageRuntimeEnvironment() bool {
	if o != nil && !IsNil(o.ImageRuntimeEnvironment) {
		return true
	}

	return false
}

// SetImageRuntimeEnvironment gets a reference to the given ImageRuntimeEnvironmentResponseModel and assigns it to the ImageRuntimeEnvironment field.
func (o *ImageVersionSpecResponseModel) SetImageRuntimeEnvironment(v ImageRuntimeEnvironmentResponseModel) {
	o.ImageRuntimeEnvironment = &v
}

// GetImageVersionSpecStatus returns the ImageVersionSpecStatus field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetImageVersionSpecStatus() ImageVersionSpecStatus {
	if o == nil || IsNil(o.ImageVersionSpecStatus) {
		var ret ImageVersionSpecStatus
		return ret
	}
	return *o.ImageVersionSpecStatus
}

// GetImageVersionSpecStatusOk returns a tuple with the ImageVersionSpecStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetImageVersionSpecStatusOk() (*ImageVersionSpecStatus, bool) {
	if o == nil || IsNil(o.ImageVersionSpecStatus) {
		return nil, false
	}
	return o.ImageVersionSpecStatus, true
}

// HasImageVersionSpecStatus returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasImageVersionSpecStatus() bool {
	if o != nil && !IsNil(o.ImageVersionSpecStatus) {
		return true
	}

	return false
}

// SetImageVersionSpecStatus gets a reference to the given ImageVersionSpecStatus and assigns it to the ImageVersionSpecStatus field.
func (o *ImageVersionSpecResponseModel) SetImageVersionSpecStatus(v ImageVersionSpecStatus) {
	o.ImageVersionSpecStatus = &v
}

// GetMasterImage returns the MasterImage field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetMasterImage() HypervisorResourceRefResponseModel {
	if o == nil || IsNil(o.MasterImage) {
		var ret HypervisorResourceRefResponseModel
		return ret
	}
	return *o.MasterImage
}

// GetMasterImageOk returns a tuple with the MasterImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetMasterImageOk() (*HypervisorResourceRefResponseModel, bool) {
	if o == nil || IsNil(o.MasterImage) {
		return nil, false
	}
	return o.MasterImage, true
}

// HasMasterImage returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasMasterImage() bool {
	if o != nil && !IsNil(o.MasterImage) {
		return true
	}

	return false
}

// SetMasterImage gets a reference to the given HypervisorResourceRefResponseModel and assigns it to the MasterImage field.
func (o *ImageVersionSpecResponseModel) SetMasterImage(v HypervisorResourceRefResponseModel) {
	o.MasterImage = &v
}

// GetPreparationType returns the PreparationType field value
func (o *ImageVersionSpecResponseModel) GetPreparationType() PreparationType {
	if o == nil {
		var ret PreparationType
		return ret
	}

	return o.PreparationType
}

// GetPreparationTypeOk returns a tuple with the PreparationType field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetPreparationTypeOk() (*PreparationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreparationType, true
}

// SetPreparationType sets field value
func (o *ImageVersionSpecResponseModel) SetPreparationType(v PreparationType) {
	o.PreparationType = v
}

// GetProvisioningSchemeCount returns the ProvisioningSchemeCount field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetProvisioningSchemeCount() int32 {
	if o == nil || IsNil(o.ProvisioningSchemeCount) {
		var ret int32
		return ret
	}
	return *o.ProvisioningSchemeCount
}

// GetProvisioningSchemeCountOk returns a tuple with the ProvisioningSchemeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetProvisioningSchemeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProvisioningSchemeCount) {
		return nil, false
	}
	return o.ProvisioningSchemeCount, true
}

// HasProvisioningSchemeCount returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasProvisioningSchemeCount() bool {
	if o != nil && !IsNil(o.ProvisioningSchemeCount) {
		return true
	}

	return false
}

// SetProvisioningSchemeCount gets a reference to the given int32 and assigns it to the ProvisioningSchemeCount field.
func (o *ImageVersionSpecResponseModel) SetProvisioningSchemeCount(v int32) {
	o.ProvisioningSchemeCount = &v
}

// GetResourcePool returns the ResourcePool field value
func (o *ImageVersionSpecResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel {
	if o == nil {
		var ret HypervisorResourcePoolRefResponseModel
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *ImageVersionSpecResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel) {
	o.ResourcePool = v
}

// GetSourceImageVersionSpec returns the SourceImageVersionSpec field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetSourceImageVersionSpec() ImageVersionSpecRefResponseModel {
	if o == nil || IsNil(o.SourceImageVersionSpec) {
		var ret ImageVersionSpecRefResponseModel
		return ret
	}
	return *o.SourceImageVersionSpec
}

// GetSourceImageVersionSpecOk returns a tuple with the SourceImageVersionSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetSourceImageVersionSpecOk() (*ImageVersionSpecRefResponseModel, bool) {
	if o == nil || IsNil(o.SourceImageVersionSpec) {
		return nil, false
	}
	return o.SourceImageVersionSpec, true
}

// HasSourceImageVersionSpec returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasSourceImageVersionSpec() bool {
	if o != nil && !IsNil(o.SourceImageVersionSpec) {
		return true
	}

	return false
}

// SetSourceImageVersionSpec gets a reference to the given ImageVersionSpecRefResponseModel and assigns it to the SourceImageVersionSpec field.
func (o *ImageVersionSpecResponseModel) SetSourceImageVersionSpec(v ImageVersionSpecRefResponseModel) {
	o.SourceImageVersionSpec = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetDiskSize() int32 {
	if o == nil || IsNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *ImageVersionSpecResponseModel) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResponseModel) GetWarnings() []ImageVersionSpecWarningResponseModel {
	if o == nil {
		var ret []ImageVersionSpecWarningResponseModel
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResponseModel) GetWarningsOk() ([]ImageVersionSpecWarningResponseModel, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasWarnings() bool {
	if o != nil && IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ImageVersionSpecWarningResponseModel and assigns it to the Warnings field.
func (o *ImageVersionSpecResponseModel) SetWarnings(v []ImageVersionSpecWarningResponseModel) {
	o.Warnings = v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ImageVersionSpecResponseModel) GetZone() RefResponseModel {
	if o == nil || IsNil(o.Zone) {
		var ret RefResponseModel
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResponseModel) GetZoneOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ImageVersionSpecResponseModel) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given RefResponseModel and assigns it to the Zone field.
func (o *ImageVersionSpecResponseModel) SetZone(v RefResponseModel) {
	o.Zone = &v
}

func (o ImageVersionSpecResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageVersionSpecResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.Context) {
		toSerialize["Context"] = o.Context
	}
	if o.CreationTime.IsSet() {
		toSerialize["CreationTime"] = o.CreationTime.Get()
	}
	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}
	if !IsNil(o.ImageRuntimeEnvironment) {
		toSerialize["ImageRuntimeEnvironment"] = o.ImageRuntimeEnvironment
	}
	if !IsNil(o.ImageVersionSpecStatus) {
		toSerialize["ImageVersionSpecStatus"] = o.ImageVersionSpecStatus
	}
	if !IsNil(o.MasterImage) {
		toSerialize["MasterImage"] = o.MasterImage
	}
	toSerialize["PreparationType"] = o.PreparationType
	if !IsNil(o.ProvisioningSchemeCount) {
		toSerialize["ProvisioningSchemeCount"] = o.ProvisioningSchemeCount
	}
	toSerialize["ResourcePool"] = o.ResourcePool
	if !IsNil(o.SourceImageVersionSpec) {
		toSerialize["SourceImageVersionSpec"] = o.SourceImageVersionSpec
	}
	if !IsNil(o.DiskSize) {
		toSerialize["DiskSize"] = o.DiskSize
	}
	if o.Warnings != nil {
		toSerialize["Warnings"] = o.Warnings
	}
	if !IsNil(o.Zone) {
		toSerialize["Zone"] = o.Zone
	}
	return toSerialize, nil
}

type NullableImageVersionSpecResponseModel struct {
	value *ImageVersionSpecResponseModel
	isSet bool
}

func (v NullableImageVersionSpecResponseModel) Get() *ImageVersionSpecResponseModel {
	return v.value
}

func (v *NullableImageVersionSpecResponseModel) Set(val *ImageVersionSpecResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionSpecResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionSpecResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionSpecResponseModel(val *ImageVersionSpecResponseModel) *NullableImageVersionSpecResponseModel {
	return &NullableImageVersionSpecResponseModel{value: val, isSet: true}
}

func (v NullableImageVersionSpecResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionSpecResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


