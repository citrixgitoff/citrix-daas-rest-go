/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the AwsEdcAccountResourceDirectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountResourceDirectory{}

// AwsEdcAccountResourceDirectory struct for AwsEdcAccountResourceDirectory
type AwsEdcAccountResourceDirectory struct {
	AwsEdcAccountResource
	// Directory ID
	DirectoryId NullableString `json:"directoryId,omitempty"`
	// Directory Name
	Name NullableString `json:"name,omitempty"`
	// Directory Description
	Description NullableString `json:"description,omitempty"`
	// Directory Domain
	Domain NullableString `json:"domain,omitempty"`
	// Directory Organizational Unit
	DefaultOU NullableString `json:"defaultOU,omitempty"`
	Type NullableAwsEdcDirectoryType `json:"type,omitempty"`
	Size NullableAwsEdcDirectorySize `json:"size,omitempty"`
	Status NullableAwsEdcDirectoryStatus `json:"status,omitempty"`
	RegistrationStatus NullableAwsEdcDirectoryRegistrationStatus `json:"registrationStatus,omitempty"`
	Tenancy NullableAwsEdcDirectoryTenancy `json:"tenancy,omitempty"`
	// Directory associated VPC
	VpcId NullableString `json:"vpcId,omitempty"`
	// Directory associated Security Group
	SecurityGroupId NullableString `json:"securityGroupId,omitempty"`
	// Directory associated Subnets
	SubnetIds []string `json:"subnetIds,omitempty"`
}

// NewAwsEdcAccountResourceDirectory instantiates a new AwsEdcAccountResourceDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountResourceDirectory(accountType AccountType) *AwsEdcAccountResourceDirectory {
	this := AwsEdcAccountResourceDirectory{}
	this.AccountType = accountType
	return &this
}

// NewAwsEdcAccountResourceDirectoryWithDefaults instantiates a new AwsEdcAccountResourceDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountResourceDirectoryWithDefaults() *AwsEdcAccountResourceDirectory {
	this := AwsEdcAccountResourceDirectory{}
	return &this
}

// GetDirectoryId returns the DirectoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetDirectoryId() string {
	if o == nil || IsNil(o.DirectoryId.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryId.Get()
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryId.Get(), o.DirectoryId.IsSet()
}

// HasDirectoryId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasDirectoryId() bool {
	if o != nil && o.DirectoryId.IsSet() {
		return true
	}

	return false
}

// SetDirectoryId gets a reference to the given NullableString and assigns it to the DirectoryId field.
func (o *AwsEdcAccountResourceDirectory) SetDirectoryId(v string) {
	o.DirectoryId.Set(&v)
}
// SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetDirectoryIdNil() {
	o.DirectoryId.Set(nil)
}

// UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetDirectoryId() {
	o.DirectoryId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AwsEdcAccountResourceDirectory) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AwsEdcAccountResourceDirectory) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetDescription() {
	o.Description.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *AwsEdcAccountResourceDirectory) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetDomain() {
	o.Domain.Unset()
}

// GetDefaultOU returns the DefaultOU field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetDefaultOU() string {
	if o == nil || IsNil(o.DefaultOU.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultOU.Get()
}

// GetDefaultOUOk returns a tuple with the DefaultOU field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetDefaultOUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultOU.Get(), o.DefaultOU.IsSet()
}

// HasDefaultOU returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasDefaultOU() bool {
	if o != nil && o.DefaultOU.IsSet() {
		return true
	}

	return false
}

// SetDefaultOU gets a reference to the given NullableString and assigns it to the DefaultOU field.
func (o *AwsEdcAccountResourceDirectory) SetDefaultOU(v string) {
	o.DefaultOU.Set(&v)
}
// SetDefaultOUNil sets the value for DefaultOU to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetDefaultOUNil() {
	o.DefaultOU.Set(nil)
}

// UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetDefaultOU() {
	o.DefaultOU.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetType() AwsEdcDirectoryType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret AwsEdcDirectoryType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetTypeOk() (*AwsEdcDirectoryType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableAwsEdcDirectoryType and assigns it to the Type field.
func (o *AwsEdcAccountResourceDirectory) SetType(v AwsEdcDirectoryType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetType() {
	o.Type.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetSize() AwsEdcDirectorySize {
	if o == nil || IsNil(o.Size.Get()) {
		var ret AwsEdcDirectorySize
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetSizeOk() (*AwsEdcDirectorySize, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableAwsEdcDirectorySize and assigns it to the Size field.
func (o *AwsEdcAccountResourceDirectory) SetSize(v AwsEdcDirectorySize) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetSize() {
	o.Size.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetStatus() AwsEdcDirectoryStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret AwsEdcDirectoryStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetStatusOk() (*AwsEdcDirectoryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAwsEdcDirectoryStatus and assigns it to the Status field.
func (o *AwsEdcAccountResourceDirectory) SetStatus(v AwsEdcDirectoryStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetStatus() {
	o.Status.Unset()
}

// GetRegistrationStatus returns the RegistrationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetRegistrationStatus() AwsEdcDirectoryRegistrationStatus {
	if o == nil || IsNil(o.RegistrationStatus.Get()) {
		var ret AwsEdcDirectoryRegistrationStatus
		return ret
	}
	return *o.RegistrationStatus.Get()
}

// GetRegistrationStatusOk returns a tuple with the RegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetRegistrationStatusOk() (*AwsEdcDirectoryRegistrationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationStatus.Get(), o.RegistrationStatus.IsSet()
}

// HasRegistrationStatus returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasRegistrationStatus() bool {
	if o != nil && o.RegistrationStatus.IsSet() {
		return true
	}

	return false
}

// SetRegistrationStatus gets a reference to the given NullableAwsEdcDirectoryRegistrationStatus and assigns it to the RegistrationStatus field.
func (o *AwsEdcAccountResourceDirectory) SetRegistrationStatus(v AwsEdcDirectoryRegistrationStatus) {
	o.RegistrationStatus.Set(&v)
}
// SetRegistrationStatusNil sets the value for RegistrationStatus to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetRegistrationStatusNil() {
	o.RegistrationStatus.Set(nil)
}

// UnsetRegistrationStatus ensures that no value is present for RegistrationStatus, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetRegistrationStatus() {
	o.RegistrationStatus.Unset()
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetTenancy() AwsEdcDirectoryTenancy {
	if o == nil || IsNil(o.Tenancy.Get()) {
		var ret AwsEdcDirectoryTenancy
		return ret
	}
	return *o.Tenancy.Get()
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetTenancyOk() (*AwsEdcDirectoryTenancy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenancy.Get(), o.Tenancy.IsSet()
}

// HasTenancy returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasTenancy() bool {
	if o != nil && o.Tenancy.IsSet() {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given NullableAwsEdcDirectoryTenancy and assigns it to the Tenancy field.
func (o *AwsEdcAccountResourceDirectory) SetTenancy(v AwsEdcDirectoryTenancy) {
	o.Tenancy.Set(&v)
}
// SetTenancyNil sets the value for Tenancy to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetTenancyNil() {
	o.Tenancy.Set(nil)
}

// UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetTenancy() {
	o.Tenancy.Unset()
}

// GetVpcId returns the VpcId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetVpcId() string {
	if o == nil || IsNil(o.VpcId.Get()) {
		var ret string
		return ret
	}
	return *o.VpcId.Get()
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VpcId.Get(), o.VpcId.IsSet()
}

// HasVpcId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasVpcId() bool {
	if o != nil && o.VpcId.IsSet() {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given NullableString and assigns it to the VpcId field.
func (o *AwsEdcAccountResourceDirectory) SetVpcId(v string) {
	o.VpcId.Set(&v)
}
// SetVpcIdNil sets the value for VpcId to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetVpcIdNil() {
	o.VpcId.Set(nil)
}

// UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetVpcId() {
	o.VpcId.Unset()
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId.Get()
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityGroupId.Get(), o.SecurityGroupId.IsSet()
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId.IsSet() {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given NullableString and assigns it to the SecurityGroupId field.
func (o *AwsEdcAccountResourceDirectory) SetSecurityGroupId(v string) {
	o.SecurityGroupId.Set(&v)
}
// SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil
func (o *AwsEdcAccountResourceDirectory) SetSecurityGroupIdNil() {
	o.SecurityGroupId.Set(nil)
}

// UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
func (o *AwsEdcAccountResourceDirectory) UnsetSecurityGroupId() {
	o.SecurityGroupId.Unset()
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceDirectory) GetSubnetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceDirectory) GetSubnetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubnetIds) {
		return nil, false
	}
	return o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceDirectory) HasSubnetIds() bool {
	if o != nil && IsNil(o.SubnetIds) {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *AwsEdcAccountResourceDirectory) SetSubnetIds(v []string) {
	o.SubnetIds = v
}

func (o AwsEdcAccountResourceDirectory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountResourceDirectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAwsEdcAccountResource, errAwsEdcAccountResource := json.Marshal(o.AwsEdcAccountResource)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	errAwsEdcAccountResource = json.Unmarshal([]byte(serializedAwsEdcAccountResource), &toSerialize)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	if o.DirectoryId.IsSet() {
		toSerialize["directoryId"] = o.DirectoryId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.DefaultOU.IsSet() {
		toSerialize["defaultOU"] = o.DefaultOU.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.RegistrationStatus.IsSet() {
		toSerialize["registrationStatus"] = o.RegistrationStatus.Get()
	}
	if o.Tenancy.IsSet() {
		toSerialize["tenancy"] = o.Tenancy.Get()
	}
	if o.VpcId.IsSet() {
		toSerialize["vpcId"] = o.VpcId.Get()
	}
	if o.SecurityGroupId.IsSet() {
		toSerialize["securityGroupId"] = o.SecurityGroupId.Get()
	}
	if o.SubnetIds != nil {
		toSerialize["subnetIds"] = o.SubnetIds
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountResourceDirectory struct {
	value *AwsEdcAccountResourceDirectory
	isSet bool
}

func (v NullableAwsEdcAccountResourceDirectory) Get() *AwsEdcAccountResourceDirectory {
	return v.value
}

func (v *NullableAwsEdcAccountResourceDirectory) Set(val *AwsEdcAccountResourceDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountResourceDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountResourceDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountResourceDirectory(val *AwsEdcAccountResourceDirectory) *NullableAwsEdcAccountResourceDirectory {
	return &NullableAwsEdcAccountResourceDirectory{value: val, isSet: true}
}

func (v NullableAwsEdcAccountResourceDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountResourceDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


