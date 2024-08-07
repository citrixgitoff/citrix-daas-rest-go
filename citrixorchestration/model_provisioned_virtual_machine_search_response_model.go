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

// checks if the ProvisionedVirtualMachineSearchResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionedVirtualMachineSearchResponseModel{}

// ProvisionedVirtualMachineSearchResponseModel Response model for provisioned virtual machine.
type ProvisionedVirtualMachineSearchResponseModel struct {
	// Active Operation.
	ActiveOperation NullableString `json:"ActiveOperation,omitempty"`
	// Provisioned Virtual Machine Sid.
	VMSid NullableString `json:"VMSid,omitempty"`
	// Cpu Count.
	CpuCount *int32 `json:"CpuCount,omitempty"`
	// Identity Type:ActiveDirectory, AzureAD, HybridAzureAD, Workgroup.
	IdentityType NullableString `json:"IdentityType,omitempty"`
	// ActiveDirectory, HybridAzureAD - Domain. AzureAD - TenantId. Workgroup - null.
	Identities NullableString `json:"Identities,omitempty"`
	// Booted image isn't same as assigned image.
	ImageOutOfDate *bool `json:"ImageOutOfDate,omitempty"`
	// Last boot time.
	LastBootTime NullableString `json:"LastBootTime,omitempty"`
	// Memory(MB).
	MemoryMB *int32 `json:"MemoryMB,omitempty"`
	// Persistency.
	Persistency *bool `json:"Persistency,omitempty"`
	// Provision Scheme Name.
	ProvisioningSchemeName NullableString `json:"ProvisioningSchemeName,omitempty"`
	// Provision Scheme Version.
	ProvisioningSchemeVersion *int32 `json:"ProvisioningSchemeVersion,omitempty"`
	// Provisioned Virtual Machine Update Configuration Version.
	ProvVMConfigurationUpdateVersion NullableInt32 `json:"ProvVMConfigurationUpdateVersion,omitempty"`
	// Provisioned virtual machine name on hypervisor.
	VMName NullableString `json:"VMName,omitempty"`
	ActivationType *WindowsActivationType `json:"ActivationType,omitempty"`
	// Whether use write back cache. 
	UseWriteBackCache *bool `json:"UseWriteBackCache,omitempty"`
	// Write back cache disk size.
	WriteBackCacheDiskSize *int32 `json:"WriteBackCacheDiskSize,omitempty"`
	// Write back cache memory size.
	WriteBackCacheMemorySize *int32 `json:"WriteBackCacheMemorySize,omitempty"`
	// List of ProvisioningOperationEventSearchResponseModel.
	OperationEvents []ProvisioningOperationEventSearchResponseModel `json:"OperationEvents,omitempty"`
}

// NewProvisionedVirtualMachineSearchResponseModel instantiates a new ProvisionedVirtualMachineSearchResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionedVirtualMachineSearchResponseModel() *ProvisionedVirtualMachineSearchResponseModel {
	this := ProvisionedVirtualMachineSearchResponseModel{}
	return &this
}

// NewProvisionedVirtualMachineSearchResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineSearchResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionedVirtualMachineSearchResponseModelWithDefaults() *ProvisionedVirtualMachineSearchResponseModel {
	this := ProvisionedVirtualMachineSearchResponseModel{}
	return &this
}

// GetActiveOperation returns the ActiveOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetActiveOperation() string {
	if o == nil || IsNil(o.ActiveOperation.Get()) {
		var ret string
		return ret
	}
	return *o.ActiveOperation.Get()
}

// GetActiveOperationOk returns a tuple with the ActiveOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetActiveOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveOperation.Get(), o.ActiveOperation.IsSet()
}

// HasActiveOperation returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasActiveOperation() bool {
	if o != nil && o.ActiveOperation.IsSet() {
		return true
	}

	return false
}

// SetActiveOperation gets a reference to the given NullableString and assigns it to the ActiveOperation field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetActiveOperation(v string) {
	o.ActiveOperation.Set(&v)
}
// SetActiveOperationNil sets the value for ActiveOperation to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetActiveOperationNil() {
	o.ActiveOperation.Set(nil)
}

// UnsetActiveOperation ensures that no value is present for ActiveOperation, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetActiveOperation() {
	o.ActiveOperation.Unset()
}

// GetVMSid returns the VMSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMSid() string {
	if o == nil || IsNil(o.VMSid.Get()) {
		var ret string
		return ret
	}
	return *o.VMSid.Get()
}

// GetVMSidOk returns a tuple with the VMSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VMSid.Get(), o.VMSid.IsSet()
}

// HasVMSid returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasVMSid() bool {
	if o != nil && o.VMSid.IsSet() {
		return true
	}

	return false
}

// SetVMSid gets a reference to the given NullableString and assigns it to the VMSid field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMSid(v string) {
	o.VMSid.Set(&v)
}
// SetVMSidNil sets the value for VMSid to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMSidNil() {
	o.VMSid.Set(nil)
}

// UnsetVMSid ensures that no value is present for VMSid, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetVMSid() {
	o.VMSid.Unset()
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetCpuCount() int32 {
	if o == nil || IsNil(o.CpuCount) {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetCpuCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuCount) {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasCpuCount() bool {
	if o != nil && !IsNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityType.Get()
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityType.Get(), o.IdentityType.IsSet()
}

// HasIdentityType returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasIdentityType() bool {
	if o != nil && o.IdentityType.IsSet() {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given NullableString and assigns it to the IdentityType field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentityType(v string) {
	o.IdentityType.Set(&v)
}
// SetIdentityTypeNil sets the value for IdentityType to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentityTypeNil() {
	o.IdentityType.Set(nil)
}

// UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetIdentityType() {
	o.IdentityType.Unset()
}

// GetIdentities returns the Identities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentities() string {
	if o == nil || IsNil(o.Identities.Get()) {
		var ret string
		return ret
	}
	return *o.Identities.Get()
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentitiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identities.Get(), o.Identities.IsSet()
}

// HasIdentities returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasIdentities() bool {
	if o != nil && o.Identities.IsSet() {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given NullableString and assigns it to the Identities field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentities(v string) {
	o.Identities.Set(&v)
}
// SetIdentitiesNil sets the value for Identities to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentitiesNil() {
	o.Identities.Set(nil)
}

// UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetIdentities() {
	o.Identities.Unset()
}

// GetImageOutOfDate returns the ImageOutOfDate field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetImageOutOfDate() bool {
	if o == nil || IsNil(o.ImageOutOfDate) {
		var ret bool
		return ret
	}
	return *o.ImageOutOfDate
}

// GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetImageOutOfDateOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageOutOfDate) {
		return nil, false
	}
	return o.ImageOutOfDate, true
}

// HasImageOutOfDate returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasImageOutOfDate() bool {
	if o != nil && !IsNil(o.ImageOutOfDate) {
		return true
	}

	return false
}

// SetImageOutOfDate gets a reference to the given bool and assigns it to the ImageOutOfDate field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetImageOutOfDate(v bool) {
	o.ImageOutOfDate = &v
}

// GetLastBootTime returns the LastBootTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetLastBootTime() string {
	if o == nil || IsNil(o.LastBootTime.Get()) {
		var ret string
		return ret
	}
	return *o.LastBootTime.Get()
}

// GetLastBootTimeOk returns a tuple with the LastBootTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetLastBootTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastBootTime.Get(), o.LastBootTime.IsSet()
}

// HasLastBootTime returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasLastBootTime() bool {
	if o != nil && o.LastBootTime.IsSet() {
		return true
	}

	return false
}

// SetLastBootTime gets a reference to the given NullableString and assigns it to the LastBootTime field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetLastBootTime(v string) {
	o.LastBootTime.Set(&v)
}
// SetLastBootTimeNil sets the value for LastBootTime to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetLastBootTimeNil() {
	o.LastBootTime.Set(nil)
}

// UnsetLastBootTime ensures that no value is present for LastBootTime, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetLastBootTime() {
	o.LastBootTime.Unset()
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetMemoryMB() int32 {
	if o == nil || IsNil(o.MemoryMB) {
		var ret int32
		return ret
	}
	return *o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetMemoryMBOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryMB) {
		return nil, false
	}
	return o.MemoryMB, true
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasMemoryMB() bool {
	if o != nil && !IsNil(o.MemoryMB) {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given int32 and assigns it to the MemoryMB field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetMemoryMB(v int32) {
	o.MemoryMB = &v
}

// GetPersistency returns the Persistency field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetPersistency() bool {
	if o == nil || IsNil(o.Persistency) {
		var ret bool
		return ret
	}
	return *o.Persistency
}

// GetPersistencyOk returns a tuple with the Persistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetPersistencyOk() (*bool, bool) {
	if o == nil || IsNil(o.Persistency) {
		return nil, false
	}
	return o.Persistency, true
}

// HasPersistency returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasPersistency() bool {
	if o != nil && !IsNil(o.Persistency) {
		return true
	}

	return false
}

// SetPersistency gets a reference to the given bool and assigns it to the Persistency field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetPersistency(v bool) {
	o.Persistency = &v
}

// GetProvisioningSchemeName returns the ProvisioningSchemeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeName() string {
	if o == nil || IsNil(o.ProvisioningSchemeName.Get()) {
		var ret string
		return ret
	}
	return *o.ProvisioningSchemeName.Get()
}

// GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisioningSchemeName.Get(), o.ProvisioningSchemeName.IsSet()
}

// HasProvisioningSchemeName returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvisioningSchemeName() bool {
	if o != nil && o.ProvisioningSchemeName.IsSet() {
		return true
	}

	return false
}

// SetProvisioningSchemeName gets a reference to the given NullableString and assigns it to the ProvisioningSchemeName field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeName(v string) {
	o.ProvisioningSchemeName.Set(&v)
}
// SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeNameNil() {
	o.ProvisioningSchemeName.Set(nil)
}

// UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetProvisioningSchemeName() {
	o.ProvisioningSchemeName.Unset()
}

// GetProvisioningSchemeVersion returns the ProvisioningSchemeVersion field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeVersion() int32 {
	if o == nil || IsNil(o.ProvisioningSchemeVersion) {
		var ret int32
		return ret
	}
	return *o.ProvisioningSchemeVersion
}

// GetProvisioningSchemeVersionOk returns a tuple with the ProvisioningSchemeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.ProvisioningSchemeVersion) {
		return nil, false
	}
	return o.ProvisioningSchemeVersion, true
}

// HasProvisioningSchemeVersion returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvisioningSchemeVersion() bool {
	if o != nil && !IsNil(o.ProvisioningSchemeVersion) {
		return true
	}

	return false
}

// SetProvisioningSchemeVersion gets a reference to the given int32 and assigns it to the ProvisioningSchemeVersion field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeVersion(v int32) {
	o.ProvisioningSchemeVersion = &v
}

// GetProvVMConfigurationUpdateVersion returns the ProvVMConfigurationUpdateVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvVMConfigurationUpdateVersion() int32 {
	if o == nil || IsNil(o.ProvVMConfigurationUpdateVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.ProvVMConfigurationUpdateVersion.Get()
}

// GetProvVMConfigurationUpdateVersionOk returns a tuple with the ProvVMConfigurationUpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvVMConfigurationUpdateVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvVMConfigurationUpdateVersion.Get(), o.ProvVMConfigurationUpdateVersion.IsSet()
}

// HasProvVMConfigurationUpdateVersion returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvVMConfigurationUpdateVersion() bool {
	if o != nil && o.ProvVMConfigurationUpdateVersion.IsSet() {
		return true
	}

	return false
}

// SetProvVMConfigurationUpdateVersion gets a reference to the given NullableInt32 and assigns it to the ProvVMConfigurationUpdateVersion field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvVMConfigurationUpdateVersion(v int32) {
	o.ProvVMConfigurationUpdateVersion.Set(&v)
}
// SetProvVMConfigurationUpdateVersionNil sets the value for ProvVMConfigurationUpdateVersion to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvVMConfigurationUpdateVersionNil() {
	o.ProvVMConfigurationUpdateVersion.Set(nil)
}

// UnsetProvVMConfigurationUpdateVersion ensures that no value is present for ProvVMConfigurationUpdateVersion, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetProvVMConfigurationUpdateVersion() {
	o.ProvVMConfigurationUpdateVersion.Unset()
}

// GetVMName returns the VMName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMName() string {
	if o == nil || IsNil(o.VMName.Get()) {
		var ret string
		return ret
	}
	return *o.VMName.Get()
}

// GetVMNameOk returns a tuple with the VMName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VMName.Get(), o.VMName.IsSet()
}

// HasVMName returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasVMName() bool {
	if o != nil && o.VMName.IsSet() {
		return true
	}

	return false
}

// SetVMName gets a reference to the given NullableString and assigns it to the VMName field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMName(v string) {
	o.VMName.Set(&v)
}
// SetVMNameNil sets the value for VMName to be an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMNameNil() {
	o.VMName.Set(nil)
}

// UnsetVMName ensures that no value is present for VMName, not even an explicit nil
func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetVMName() {
	o.VMName.Unset()
}

// GetActivationType returns the ActivationType field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetActivationType() WindowsActivationType {
	if o == nil || IsNil(o.ActivationType) {
		var ret WindowsActivationType
		return ret
	}
	return *o.ActivationType
}

// GetActivationTypeOk returns a tuple with the ActivationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetActivationTypeOk() (*WindowsActivationType, bool) {
	if o == nil || IsNil(o.ActivationType) {
		return nil, false
	}
	return o.ActivationType, true
}

// HasActivationType returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasActivationType() bool {
	if o != nil && !IsNil(o.ActivationType) {
		return true
	}

	return false
}

// SetActivationType gets a reference to the given WindowsActivationType and assigns it to the ActivationType field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetActivationType(v WindowsActivationType) {
	o.ActivationType = &v
}

// GetUseWriteBackCache returns the UseWriteBackCache field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetUseWriteBackCache() bool {
	if o == nil || IsNil(o.UseWriteBackCache) {
		var ret bool
		return ret
	}
	return *o.UseWriteBackCache
}

// GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetUseWriteBackCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.UseWriteBackCache) {
		return nil, false
	}
	return o.UseWriteBackCache, true
}

// HasUseWriteBackCache returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasUseWriteBackCache() bool {
	if o != nil && !IsNil(o.UseWriteBackCache) {
		return true
	}

	return false
}

// SetUseWriteBackCache gets a reference to the given bool and assigns it to the UseWriteBackCache field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetUseWriteBackCache(v bool) {
	o.UseWriteBackCache = &v
}

// GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheDiskSize() int32 {
	if o == nil || IsNil(o.WriteBackCacheDiskSize) {
		var ret int32
		return ret
	}
	return *o.WriteBackCacheDiskSize
}

// GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.WriteBackCacheDiskSize) {
		return nil, false
	}
	return o.WriteBackCacheDiskSize, true
}

// HasWriteBackCacheDiskSize returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasWriteBackCacheDiskSize() bool {
	if o != nil && !IsNil(o.WriteBackCacheDiskSize) {
		return true
	}

	return false
}

// SetWriteBackCacheDiskSize gets a reference to the given int32 and assigns it to the WriteBackCacheDiskSize field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetWriteBackCacheDiskSize(v int32) {
	o.WriteBackCacheDiskSize = &v
}

// GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field value if set, zero value otherwise.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheMemorySize() int32 {
	if o == nil || IsNil(o.WriteBackCacheMemorySize) {
		var ret int32
		return ret
	}
	return *o.WriteBackCacheMemorySize
}

// GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool) {
	if o == nil || IsNil(o.WriteBackCacheMemorySize) {
		return nil, false
	}
	return o.WriteBackCacheMemorySize, true
}

// HasWriteBackCacheMemorySize returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasWriteBackCacheMemorySize() bool {
	if o != nil && !IsNil(o.WriteBackCacheMemorySize) {
		return true
	}

	return false
}

// SetWriteBackCacheMemorySize gets a reference to the given int32 and assigns it to the WriteBackCacheMemorySize field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetWriteBackCacheMemorySize(v int32) {
	o.WriteBackCacheMemorySize = &v
}

// GetOperationEvents returns the OperationEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedVirtualMachineSearchResponseModel) GetOperationEvents() []ProvisioningOperationEventSearchResponseModel {
	if o == nil {
		var ret []ProvisioningOperationEventSearchResponseModel
		return ret
	}
	return o.OperationEvents
}

// GetOperationEventsOk returns a tuple with the OperationEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedVirtualMachineSearchResponseModel) GetOperationEventsOk() ([]ProvisioningOperationEventSearchResponseModel, bool) {
	if o == nil || IsNil(o.OperationEvents) {
		return nil, false
	}
	return o.OperationEvents, true
}

// HasOperationEvents returns a boolean if a field has been set.
func (o *ProvisionedVirtualMachineSearchResponseModel) HasOperationEvents() bool {
	if o != nil && IsNil(o.OperationEvents) {
		return true
	}

	return false
}

// SetOperationEvents gets a reference to the given []ProvisioningOperationEventSearchResponseModel and assigns it to the OperationEvents field.
func (o *ProvisionedVirtualMachineSearchResponseModel) SetOperationEvents(v []ProvisioningOperationEventSearchResponseModel) {
	o.OperationEvents = v
}

func (o ProvisionedVirtualMachineSearchResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionedVirtualMachineSearchResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveOperation.IsSet() {
		toSerialize["ActiveOperation"] = o.ActiveOperation.Get()
	}
	if o.VMSid.IsSet() {
		toSerialize["VMSid"] = o.VMSid.Get()
	}
	if !IsNil(o.CpuCount) {
		toSerialize["CpuCount"] = o.CpuCount
	}
	if o.IdentityType.IsSet() {
		toSerialize["IdentityType"] = o.IdentityType.Get()
	}
	if o.Identities.IsSet() {
		toSerialize["Identities"] = o.Identities.Get()
	}
	if !IsNil(o.ImageOutOfDate) {
		toSerialize["ImageOutOfDate"] = o.ImageOutOfDate
	}
	if o.LastBootTime.IsSet() {
		toSerialize["LastBootTime"] = o.LastBootTime.Get()
	}
	if !IsNil(o.MemoryMB) {
		toSerialize["MemoryMB"] = o.MemoryMB
	}
	if !IsNil(o.Persistency) {
		toSerialize["Persistency"] = o.Persistency
	}
	if o.ProvisioningSchemeName.IsSet() {
		toSerialize["ProvisioningSchemeName"] = o.ProvisioningSchemeName.Get()
	}
	if !IsNil(o.ProvisioningSchemeVersion) {
		toSerialize["ProvisioningSchemeVersion"] = o.ProvisioningSchemeVersion
	}
	if o.ProvVMConfigurationUpdateVersion.IsSet() {
		toSerialize["ProvVMConfigurationUpdateVersion"] = o.ProvVMConfigurationUpdateVersion.Get()
	}
	if o.VMName.IsSet() {
		toSerialize["VMName"] = o.VMName.Get()
	}
	if !IsNil(o.ActivationType) {
		toSerialize["ActivationType"] = o.ActivationType
	}
	if !IsNil(o.UseWriteBackCache) {
		toSerialize["UseWriteBackCache"] = o.UseWriteBackCache
	}
	if !IsNil(o.WriteBackCacheDiskSize) {
		toSerialize["WriteBackCacheDiskSize"] = o.WriteBackCacheDiskSize
	}
	if !IsNil(o.WriteBackCacheMemorySize) {
		toSerialize["WriteBackCacheMemorySize"] = o.WriteBackCacheMemorySize
	}
	if o.OperationEvents != nil {
		toSerialize["OperationEvents"] = o.OperationEvents
	}
	return toSerialize, nil
}

type NullableProvisionedVirtualMachineSearchResponseModel struct {
	value *ProvisionedVirtualMachineSearchResponseModel
	isSet bool
}

func (v NullableProvisionedVirtualMachineSearchResponseModel) Get() *ProvisionedVirtualMachineSearchResponseModel {
	return v.value
}

func (v *NullableProvisionedVirtualMachineSearchResponseModel) Set(val *ProvisionedVirtualMachineSearchResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionedVirtualMachineSearchResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionedVirtualMachineSearchResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionedVirtualMachineSearchResponseModel(val *ProvisionedVirtualMachineSearchResponseModel) *NullableProvisionedVirtualMachineSearchResponseModel {
	return &NullableProvisionedVirtualMachineSearchResponseModel{value: val, isSet: true}
}

func (v NullableProvisionedVirtualMachineSearchResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionedVirtualMachineSearchResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


