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

// checks if the IdentityMachineResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityMachineResponseModel{}

// IdentityMachineResponseModel struct for IdentityMachineResponseModel
type IdentityMachineResponseModel struct {
	// The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched.
	Domain NullableString `json:"Domain,omitempty"`
	// The forest the object belongs to.
	Forest NullableString `json:"Forest,omitempty"`
	// The GUID of the object, the GUID is unique across the enterprise and anywhere else.
	Guid NullableString `json:"Guid,omitempty"`
	// The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory.
	DistinguishedName NullableString `json:"DistinguishedName,omitempty"`
	// The unique distinguished name of the object in canonical format.
	CanonicalName NullableString `json:"CanonicalName,omitempty"`
	// Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown.
	PossibleLookupFailure *bool `json:"PossibleLookupFailure,omitempty"`
	// Domain controller that performed the add/update operation on this object
	DirectoryServer NullableString `json:"DirectoryServer,omitempty"`
	// Fully qualified sAMAccountName of the Machine object (including domain).
	SamName NullableString `json:"SamName,omitempty"`
	// Unqualified SAM name (Directory property) of the Machine object.  Does not include the domain. The sAMAccountName is unique within a domain.
	SamAccountName NullableString `json:"SamAccountName,omitempty"`
	// The full qualified domain name associated with the Machine object.
	DnsName NullableString `json:"DnsName,omitempty"`
	// The Security Identifier associated with the Machine object.
	Sid NullableString `json:"Sid,omitempty"`
	// The service principal names (SPN) associated with the Machine object.
	ServicePrincipalNames []string `json:"ServicePrincipalNames,omitempty"`
	// The IP addresses associated with the Machine object.
	IPAddress []string `json:"IPAddress,omitempty"`
	IPAddressResolveMethod *IdentityMachineIPAddressResolveMethod `json:"IPAddressResolveMethod,omitempty"`
	// Indicates whether the Machine object is enabled. NOTE that this is opposite of the low-level SDK, which has \"IsDisabled\". By changing to \"Enabled\" it avoids a confusing double-negative.  It also matches all other SDK objects by using \"Enabled\" rather than \"IsEnabled\".
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the Machine object is locked. low-level has \"IsLocked\".
	Locked *bool `json:"Locked,omitempty"`
	// Properties fetched and populated in the Machine object.  This is a bitfield indicating the fetched properties.
	PropertiesFetched int32 `json:"PropertiesFetched"`
}

// NewIdentityMachineResponseModel instantiates a new IdentityMachineResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMachineResponseModel(propertiesFetched int32) *IdentityMachineResponseModel {
	this := IdentityMachineResponseModel{}
	this.PropertiesFetched = propertiesFetched
	return &this
}

// NewIdentityMachineResponseModelWithDefaults instantiates a new IdentityMachineResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMachineResponseModelWithDefaults() *IdentityMachineResponseModel {
	this := IdentityMachineResponseModel{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *IdentityMachineResponseModel) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *IdentityMachineResponseModel) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetDomain() {
	o.Domain.Unset()
}

// GetForest returns the Forest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetForest() string {
	if o == nil || IsNil(o.Forest.Get()) {
		var ret string
		return ret
	}
	return *o.Forest.Get()
}

// GetForestOk returns a tuple with the Forest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetForestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forest.Get(), o.Forest.IsSet()
}

// HasForest returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasForest() bool {
	if o != nil && o.Forest.IsSet() {
		return true
	}

	return false
}

// SetForest gets a reference to the given NullableString and assigns it to the Forest field.
func (o *IdentityMachineResponseModel) SetForest(v string) {
	o.Forest.Set(&v)
}
// SetForestNil sets the value for Forest to be an explicit nil
func (o *IdentityMachineResponseModel) SetForestNil() {
	o.Forest.Set(nil)
}

// UnsetForest ensures that no value is present for Forest, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetForest() {
	o.Forest.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *IdentityMachineResponseModel) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *IdentityMachineResponseModel) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetGuid() {
	o.Guid.Unset()
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetDistinguishedName() string {
	if o == nil || IsNil(o.DistinguishedName.Get()) {
		var ret string
		return ret
	}
	return *o.DistinguishedName.Get()
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetDistinguishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistinguishedName.Get(), o.DistinguishedName.IsSet()
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName.IsSet() {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given NullableString and assigns it to the DistinguishedName field.
func (o *IdentityMachineResponseModel) SetDistinguishedName(v string) {
	o.DistinguishedName.Set(&v)
}
// SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil
func (o *IdentityMachineResponseModel) SetDistinguishedNameNil() {
	o.DistinguishedName.Set(nil)
}

// UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetDistinguishedName() {
	o.DistinguishedName.Unset()
}

// GetCanonicalName returns the CanonicalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetCanonicalName() string {
	if o == nil || IsNil(o.CanonicalName.Get()) {
		var ret string
		return ret
	}
	return *o.CanonicalName.Get()
}

// GetCanonicalNameOk returns a tuple with the CanonicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetCanonicalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanonicalName.Get(), o.CanonicalName.IsSet()
}

// HasCanonicalName returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasCanonicalName() bool {
	if o != nil && o.CanonicalName.IsSet() {
		return true
	}

	return false
}

// SetCanonicalName gets a reference to the given NullableString and assigns it to the CanonicalName field.
func (o *IdentityMachineResponseModel) SetCanonicalName(v string) {
	o.CanonicalName.Set(&v)
}
// SetCanonicalNameNil sets the value for CanonicalName to be an explicit nil
func (o *IdentityMachineResponseModel) SetCanonicalNameNil() {
	o.CanonicalName.Set(nil)
}

// UnsetCanonicalName ensures that no value is present for CanonicalName, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetCanonicalName() {
	o.CanonicalName.Unset()
}

// GetPossibleLookupFailure returns the PossibleLookupFailure field value if set, zero value otherwise.
func (o *IdentityMachineResponseModel) GetPossibleLookupFailure() bool {
	if o == nil || IsNil(o.PossibleLookupFailure) {
		var ret bool
		return ret
	}
	return *o.PossibleLookupFailure
}

// GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMachineResponseModel) GetPossibleLookupFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.PossibleLookupFailure) {
		return nil, false
	}
	return o.PossibleLookupFailure, true
}

// HasPossibleLookupFailure returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasPossibleLookupFailure() bool {
	if o != nil && !IsNil(o.PossibleLookupFailure) {
		return true
	}

	return false
}

// SetPossibleLookupFailure gets a reference to the given bool and assigns it to the PossibleLookupFailure field.
func (o *IdentityMachineResponseModel) SetPossibleLookupFailure(v bool) {
	o.PossibleLookupFailure = &v
}

// GetDirectoryServer returns the DirectoryServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetDirectoryServer() string {
	if o == nil || IsNil(o.DirectoryServer.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryServer.Get()
}

// GetDirectoryServerOk returns a tuple with the DirectoryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetDirectoryServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryServer.Get(), o.DirectoryServer.IsSet()
}

// HasDirectoryServer returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasDirectoryServer() bool {
	if o != nil && o.DirectoryServer.IsSet() {
		return true
	}

	return false
}

// SetDirectoryServer gets a reference to the given NullableString and assigns it to the DirectoryServer field.
func (o *IdentityMachineResponseModel) SetDirectoryServer(v string) {
	o.DirectoryServer.Set(&v)
}
// SetDirectoryServerNil sets the value for DirectoryServer to be an explicit nil
func (o *IdentityMachineResponseModel) SetDirectoryServerNil() {
	o.DirectoryServer.Set(nil)
}

// UnsetDirectoryServer ensures that no value is present for DirectoryServer, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetDirectoryServer() {
	o.DirectoryServer.Unset()
}

// GetSamName returns the SamName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetSamName() string {
	if o == nil || IsNil(o.SamName.Get()) {
		var ret string
		return ret
	}
	return *o.SamName.Get()
}

// GetSamNameOk returns a tuple with the SamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetSamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamName.Get(), o.SamName.IsSet()
}

// HasSamName returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasSamName() bool {
	if o != nil && o.SamName.IsSet() {
		return true
	}

	return false
}

// SetSamName gets a reference to the given NullableString and assigns it to the SamName field.
func (o *IdentityMachineResponseModel) SetSamName(v string) {
	o.SamName.Set(&v)
}
// SetSamNameNil sets the value for SamName to be an explicit nil
func (o *IdentityMachineResponseModel) SetSamNameNil() {
	o.SamName.Set(nil)
}

// UnsetSamName ensures that no value is present for SamName, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetSamName() {
	o.SamName.Unset()
}

// GetSamAccountName returns the SamAccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetSamAccountName() string {
	if o == nil || IsNil(o.SamAccountName.Get()) {
		var ret string
		return ret
	}
	return *o.SamAccountName.Get()
}

// GetSamAccountNameOk returns a tuple with the SamAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetSamAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamAccountName.Get(), o.SamAccountName.IsSet()
}

// HasSamAccountName returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasSamAccountName() bool {
	if o != nil && o.SamAccountName.IsSet() {
		return true
	}

	return false
}

// SetSamAccountName gets a reference to the given NullableString and assigns it to the SamAccountName field.
func (o *IdentityMachineResponseModel) SetSamAccountName(v string) {
	o.SamAccountName.Set(&v)
}
// SetSamAccountNameNil sets the value for SamAccountName to be an explicit nil
func (o *IdentityMachineResponseModel) SetSamAccountNameNil() {
	o.SamAccountName.Set(nil)
}

// UnsetSamAccountName ensures that no value is present for SamAccountName, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetSamAccountName() {
	o.SamAccountName.Unset()
}

// GetDnsName returns the DnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetDnsName() string {
	if o == nil || IsNil(o.DnsName.Get()) {
		var ret string
		return ret
	}
	return *o.DnsName.Get()
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsName.Get(), o.DnsName.IsSet()
}

// HasDnsName returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasDnsName() bool {
	if o != nil && o.DnsName.IsSet() {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given NullableString and assigns it to the DnsName field.
func (o *IdentityMachineResponseModel) SetDnsName(v string) {
	o.DnsName.Set(&v)
}
// SetDnsNameNil sets the value for DnsName to be an explicit nil
func (o *IdentityMachineResponseModel) SetDnsNameNil() {
	o.DnsName.Set(nil)
}

// UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetDnsName() {
	o.DnsName.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetSid() string {
	if o == nil || IsNil(o.Sid.Get()) {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *IdentityMachineResponseModel) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *IdentityMachineResponseModel) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *IdentityMachineResponseModel) UnsetSid() {
	o.Sid.Unset()
}

// GetServicePrincipalNames returns the ServicePrincipalNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetServicePrincipalNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ServicePrincipalNames
}

// GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetServicePrincipalNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServicePrincipalNames) {
		return nil, false
	}
	return o.ServicePrincipalNames, true
}

// HasServicePrincipalNames returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasServicePrincipalNames() bool {
	if o != nil && IsNil(o.ServicePrincipalNames) {
		return true
	}

	return false
}

// SetServicePrincipalNames gets a reference to the given []string and assigns it to the ServicePrincipalNames field.
func (o *IdentityMachineResponseModel) SetServicePrincipalNames(v []string) {
	o.ServicePrincipalNames = v
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMachineResponseModel) GetIPAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMachineResponseModel) GetIPAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.IPAddress) {
		return nil, false
	}
	return o.IPAddress, true
}

// HasIPAddress returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasIPAddress() bool {
	if o != nil && IsNil(o.IPAddress) {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given []string and assigns it to the IPAddress field.
func (o *IdentityMachineResponseModel) SetIPAddress(v []string) {
	o.IPAddress = v
}

// GetIPAddressResolveMethod returns the IPAddressResolveMethod field value if set, zero value otherwise.
func (o *IdentityMachineResponseModel) GetIPAddressResolveMethod() IdentityMachineIPAddressResolveMethod {
	if o == nil || IsNil(o.IPAddressResolveMethod) {
		var ret IdentityMachineIPAddressResolveMethod
		return ret
	}
	return *o.IPAddressResolveMethod
}

// GetIPAddressResolveMethodOk returns a tuple with the IPAddressResolveMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMachineResponseModel) GetIPAddressResolveMethodOk() (*IdentityMachineIPAddressResolveMethod, bool) {
	if o == nil || IsNil(o.IPAddressResolveMethod) {
		return nil, false
	}
	return o.IPAddressResolveMethod, true
}

// HasIPAddressResolveMethod returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasIPAddressResolveMethod() bool {
	if o != nil && !IsNil(o.IPAddressResolveMethod) {
		return true
	}

	return false
}

// SetIPAddressResolveMethod gets a reference to the given IdentityMachineIPAddressResolveMethod and assigns it to the IPAddressResolveMethod field.
func (o *IdentityMachineResponseModel) SetIPAddressResolveMethod(v IdentityMachineIPAddressResolveMethod) {
	o.IPAddressResolveMethod = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IdentityMachineResponseModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMachineResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IdentityMachineResponseModel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *IdentityMachineResponseModel) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMachineResponseModel) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *IdentityMachineResponseModel) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *IdentityMachineResponseModel) SetLocked(v bool) {
	o.Locked = &v
}

// GetPropertiesFetched returns the PropertiesFetched field value
func (o *IdentityMachineResponseModel) GetPropertiesFetched() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertiesFetched
}

// GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field value
// and a boolean to check if the value has been set.
func (o *IdentityMachineResponseModel) GetPropertiesFetchedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertiesFetched, true
}

// SetPropertiesFetched sets field value
func (o *IdentityMachineResponseModel) SetPropertiesFetched(v int32) {
	o.PropertiesFetched = v
}

func (o IdentityMachineResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityMachineResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain.IsSet() {
		toSerialize["Domain"] = o.Domain.Get()
	}
	if o.Forest.IsSet() {
		toSerialize["Forest"] = o.Forest.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["Guid"] = o.Guid.Get()
	}
	if o.DistinguishedName.IsSet() {
		toSerialize["DistinguishedName"] = o.DistinguishedName.Get()
	}
	if o.CanonicalName.IsSet() {
		toSerialize["CanonicalName"] = o.CanonicalName.Get()
	}
	if !IsNil(o.PossibleLookupFailure) {
		toSerialize["PossibleLookupFailure"] = o.PossibleLookupFailure
	}
	if o.DirectoryServer.IsSet() {
		toSerialize["DirectoryServer"] = o.DirectoryServer.Get()
	}
	if o.SamName.IsSet() {
		toSerialize["SamName"] = o.SamName.Get()
	}
	if o.SamAccountName.IsSet() {
		toSerialize["SamAccountName"] = o.SamAccountName.Get()
	}
	if o.DnsName.IsSet() {
		toSerialize["DnsName"] = o.DnsName.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["Sid"] = o.Sid.Get()
	}
	if o.ServicePrincipalNames != nil {
		toSerialize["ServicePrincipalNames"] = o.ServicePrincipalNames
	}
	if o.IPAddress != nil {
		toSerialize["IPAddress"] = o.IPAddress
	}
	if !IsNil(o.IPAddressResolveMethod) {
		toSerialize["IPAddressResolveMethod"] = o.IPAddressResolveMethod
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Locked) {
		toSerialize["Locked"] = o.Locked
	}
	toSerialize["PropertiesFetched"] = o.PropertiesFetched
	return toSerialize, nil
}

type NullableIdentityMachineResponseModel struct {
	value *IdentityMachineResponseModel
	isSet bool
}

func (v NullableIdentityMachineResponseModel) Get() *IdentityMachineResponseModel {
	return v.value
}

func (v *NullableIdentityMachineResponseModel) Set(val *IdentityMachineResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMachineResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMachineResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMachineResponseModel(val *IdentityMachineResponseModel) *NullableIdentityMachineResponseModel {
	return &NullableIdentityMachineResponseModel{value: val, isSet: true}
}

func (v NullableIdentityMachineResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMachineResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


