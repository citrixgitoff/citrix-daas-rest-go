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

// checks if the ApplicationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResponseModel{}

// ApplicationResponseModel Response object for application.
type ApplicationResponseModel struct {
	// Id of the application. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility
	Id string `json:"Id"`
	// `DEPRECATED.  Use <see cref='Id'/>.` DEPRECATED. Use Id.
	Uid *int32 `json:"Uid,omitempty"`
	ApplicationFolder ApplicationResponseModelApplicationFolder `json:"ApplicationFolder"`
	ApplicationType ApplicationType `json:"ApplicationType"`
	// The folder that the application belongs to as the user sees it.
	ClientFolder *string `json:"ClientFolder,omitempty"`
	// Delegated admin scopes in which the containers of the application reside.
	ContainerScopes []ContainerScopeResponseModel `json:"ContainerScopes"`
	// The description of the application.
	Description *string `json:"Description,omitempty"`
	// Indicates whether or not this application is enumerable
	DoNotEnumerate *bool `json:"DoNotEnumerate,omitempty"`
	// Indicates whether or not this application can be launched.
	Enabled bool `json:"Enabled"`
	// Id of the icon used for the application. Used to be: IconUid (and it was not globally unique) Needs to be globally unique Might be constructed from site ID + internal Uid
	IconId string `json:"IconId"`
	InstalledAppProperties *ApplicationResponseModelInstalledAppProperties `json:"InstalledAppProperties,omitempty"`
	AppVAppProperties *ApplicationResponseModelAppVAppProperties `json:"AppVAppProperties,omitempty"`
	// Location of published content.
	ContentLocation *string `json:"ContentLocation,omitempty"`
	// Name of the application.  Only seen by administrators.
	Name string `json:"Name"`
	// The name seen by end users who have access to the application.
	PublishedName string `json:"PublishedName"`
	// Indicates whether or not this application is visible to users.
	Visible bool `json:"Visible"`
	SharingKind SharingKind `json:"SharingKind"`
	// Tags associated with this application.
	Tags []string `json:"Tags,omitempty"`
	// The tenant(s) that the application is assigned to.  If `null`, the application is not assigned to any tenants, and may be used by any tenant.
	Tenants []RefResponseModel `json:"Tenants,omitempty"`
	// Indicates whether users are managed in the Citrix Cloud Library, or within Studio.
	CloudWorkspaceManaged *bool `json:"CloudWorkspaceManaged,omitempty"`
	// Number of delivery groups that the application is associated with.
	NumAssociatedDeliveryGroups *int32 `json:"NumAssociatedDeliveryGroups,omitempty"`
	// Number of application groups that the application is associated with.
	NumAssociatedApplicationGroups *int32 `json:"NumAssociatedApplicationGroups,omitempty"`
	// Delivery group Uuids that the application is associated with.
	AssociatedDeliveryGroupUuids []string `json:"AssociatedDeliveryGroupUuids,omitempty"`
	// Application group Uuids that the application is associated with.
	AssociatedApplicationGroupUuids []string `json:"AssociatedApplicationGroupUuids,omitempty"`
	// Application Zone info.
	ZoneId *string `json:"ZoneId,omitempty"`
}

// NewApplicationResponseModel instantiates a new ApplicationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponseModel(id string, applicationFolder ApplicationResponseModelApplicationFolder, applicationType ApplicationType, containerScopes []ContainerScopeResponseModel, enabled bool, iconId string, name string, publishedName string, visible bool, sharingKind SharingKind) *ApplicationResponseModel {
	this := ApplicationResponseModel{}
	this.Id = id
	this.ApplicationFolder = applicationFolder
	this.ApplicationType = applicationType
	this.ContainerScopes = containerScopes
	this.Enabled = enabled
	this.IconId = iconId
	this.Name = name
	this.PublishedName = publishedName
	this.Visible = visible
	this.SharingKind = sharingKind
	return &this
}

// NewApplicationResponseModelWithDefaults instantiates a new ApplicationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseModelWithDefaults() *ApplicationResponseModel {
	this := ApplicationResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResponseModel) SetId(v string) {
	o.Id = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *ApplicationResponseModel) SetUid(v int32) {
	o.Uid = &v
}

// GetApplicationFolder returns the ApplicationFolder field value
func (o *ApplicationResponseModel) GetApplicationFolder() ApplicationResponseModelApplicationFolder {
	if o == nil {
		var ret ApplicationResponseModelApplicationFolder
		return ret
	}

	return o.ApplicationFolder
}

// GetApplicationFolderOk returns a tuple with the ApplicationFolder field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetApplicationFolderOk() (*ApplicationResponseModelApplicationFolder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationFolder, true
}

// SetApplicationFolder sets field value
func (o *ApplicationResponseModel) SetApplicationFolder(v ApplicationResponseModelApplicationFolder) {
	o.ApplicationFolder = v
}

// GetApplicationType returns the ApplicationType field value
func (o *ApplicationResponseModel) GetApplicationType() ApplicationType {
	if o == nil {
		var ret ApplicationType
		return ret
	}

	return o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetApplicationTypeOk() (*ApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationType, true
}

// SetApplicationType sets field value
func (o *ApplicationResponseModel) SetApplicationType(v ApplicationType) {
	o.ApplicationType = v
}

// GetClientFolder returns the ClientFolder field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetClientFolder() string {
	if o == nil || IsNil(o.ClientFolder) {
		var ret string
		return ret
	}
	return *o.ClientFolder
}

// GetClientFolderOk returns a tuple with the ClientFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetClientFolderOk() (*string, bool) {
	if o == nil || IsNil(o.ClientFolder) {
		return nil, false
	}
	return o.ClientFolder, true
}

// HasClientFolder returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasClientFolder() bool {
	if o != nil && !IsNil(o.ClientFolder) {
		return true
	}

	return false
}

// SetClientFolder gets a reference to the given string and assigns it to the ClientFolder field.
func (o *ApplicationResponseModel) SetClientFolder(v string) {
	o.ClientFolder = &v
}

// GetContainerScopes returns the ContainerScopes field value
func (o *ApplicationResponseModel) GetContainerScopes() []ContainerScopeResponseModel {
	if o == nil {
		var ret []ContainerScopeResponseModel
		return ret
	}

	return o.ContainerScopes
}

// GetContainerScopesOk returns a tuple with the ContainerScopes field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetContainerScopesOk() ([]ContainerScopeResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerScopes, true
}

// SetContainerScopes sets field value
func (o *ApplicationResponseModel) SetContainerScopes(v []ContainerScopeResponseModel) {
	o.ContainerScopes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationResponseModel) SetDescription(v string) {
	o.Description = &v
}

// GetDoNotEnumerate returns the DoNotEnumerate field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetDoNotEnumerate() bool {
	if o == nil || IsNil(o.DoNotEnumerate) {
		var ret bool
		return ret
	}
	return *o.DoNotEnumerate
}

// GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetDoNotEnumerateOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotEnumerate) {
		return nil, false
	}
	return o.DoNotEnumerate, true
}

// HasDoNotEnumerate returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasDoNotEnumerate() bool {
	if o != nil && !IsNil(o.DoNotEnumerate) {
		return true
	}

	return false
}

// SetDoNotEnumerate gets a reference to the given bool and assigns it to the DoNotEnumerate field.
func (o *ApplicationResponseModel) SetDoNotEnumerate(v bool) {
	o.DoNotEnumerate = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationResponseModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationResponseModel) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIconId returns the IconId field value
func (o *ApplicationResponseModel) GetIconId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetIconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconId, true
}

// SetIconId sets field value
func (o *ApplicationResponseModel) SetIconId(v string) {
	o.IconId = v
}

// GetInstalledAppProperties returns the InstalledAppProperties field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetInstalledAppProperties() ApplicationResponseModelInstalledAppProperties {
	if o == nil || IsNil(o.InstalledAppProperties) {
		var ret ApplicationResponseModelInstalledAppProperties
		return ret
	}
	return *o.InstalledAppProperties
}

// GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetInstalledAppPropertiesOk() (*ApplicationResponseModelInstalledAppProperties, bool) {
	if o == nil || IsNil(o.InstalledAppProperties) {
		return nil, false
	}
	return o.InstalledAppProperties, true
}

// HasInstalledAppProperties returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasInstalledAppProperties() bool {
	if o != nil && !IsNil(o.InstalledAppProperties) {
		return true
	}

	return false
}

// SetInstalledAppProperties gets a reference to the given ApplicationResponseModelInstalledAppProperties and assigns it to the InstalledAppProperties field.
func (o *ApplicationResponseModel) SetInstalledAppProperties(v ApplicationResponseModelInstalledAppProperties) {
	o.InstalledAppProperties = &v
}

// GetAppVAppProperties returns the AppVAppProperties field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetAppVAppProperties() ApplicationResponseModelAppVAppProperties {
	if o == nil || IsNil(o.AppVAppProperties) {
		var ret ApplicationResponseModelAppVAppProperties
		return ret
	}
	return *o.AppVAppProperties
}

// GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetAppVAppPropertiesOk() (*ApplicationResponseModelAppVAppProperties, bool) {
	if o == nil || IsNil(o.AppVAppProperties) {
		return nil, false
	}
	return o.AppVAppProperties, true
}

// HasAppVAppProperties returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasAppVAppProperties() bool {
	if o != nil && !IsNil(o.AppVAppProperties) {
		return true
	}

	return false
}

// SetAppVAppProperties gets a reference to the given ApplicationResponseModelAppVAppProperties and assigns it to the AppVAppProperties field.
func (o *ApplicationResponseModel) SetAppVAppProperties(v ApplicationResponseModelAppVAppProperties) {
	o.AppVAppProperties = &v
}

// GetContentLocation returns the ContentLocation field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetContentLocation() string {
	if o == nil || IsNil(o.ContentLocation) {
		var ret string
		return ret
	}
	return *o.ContentLocation
}

// GetContentLocationOk returns a tuple with the ContentLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetContentLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ContentLocation) {
		return nil, false
	}
	return o.ContentLocation, true
}

// HasContentLocation returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasContentLocation() bool {
	if o != nil && !IsNil(o.ContentLocation) {
		return true
	}

	return false
}

// SetContentLocation gets a reference to the given string and assigns it to the ContentLocation field.
func (o *ApplicationResponseModel) SetContentLocation(v string) {
	o.ContentLocation = &v
}

// GetName returns the Name field value
func (o *ApplicationResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationResponseModel) SetName(v string) {
	o.Name = v
}

// GetPublishedName returns the PublishedName field value
func (o *ApplicationResponseModel) GetPublishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedName
}

// GetPublishedNameOk returns a tuple with the PublishedName field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetPublishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedName, true
}

// SetPublishedName sets field value
func (o *ApplicationResponseModel) SetPublishedName(v string) {
	o.PublishedName = v
}

// GetVisible returns the Visible field value
func (o *ApplicationResponseModel) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *ApplicationResponseModel) SetVisible(v bool) {
	o.Visible = v
}

// GetSharingKind returns the SharingKind field value
func (o *ApplicationResponseModel) GetSharingKind() SharingKind {
	if o == nil {
		var ret SharingKind
		return ret
	}

	return o.SharingKind
}

// GetSharingKindOk returns a tuple with the SharingKind field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetSharingKindOk() (*SharingKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharingKind, true
}

// SetSharingKind sets field value
func (o *ApplicationResponseModel) SetSharingKind(v SharingKind) {
	o.SharingKind = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApplicationResponseModel) SetTags(v []string) {
	o.Tags = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetTenants() []RefResponseModel {
	if o == nil || IsNil(o.Tenants) {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *ApplicationResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

// GetCloudWorkspaceManaged returns the CloudWorkspaceManaged field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetCloudWorkspaceManaged() bool {
	if o == nil || IsNil(o.CloudWorkspaceManaged) {
		var ret bool
		return ret
	}
	return *o.CloudWorkspaceManaged
}

// GetCloudWorkspaceManagedOk returns a tuple with the CloudWorkspaceManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetCloudWorkspaceManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.CloudWorkspaceManaged) {
		return nil, false
	}
	return o.CloudWorkspaceManaged, true
}

// HasCloudWorkspaceManaged returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasCloudWorkspaceManaged() bool {
	if o != nil && !IsNil(o.CloudWorkspaceManaged) {
		return true
	}

	return false
}

// SetCloudWorkspaceManaged gets a reference to the given bool and assigns it to the CloudWorkspaceManaged field.
func (o *ApplicationResponseModel) SetCloudWorkspaceManaged(v bool) {
	o.CloudWorkspaceManaged = &v
}

// GetNumAssociatedDeliveryGroups returns the NumAssociatedDeliveryGroups field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetNumAssociatedDeliveryGroups() int32 {
	if o == nil || IsNil(o.NumAssociatedDeliveryGroups) {
		var ret int32
		return ret
	}
	return *o.NumAssociatedDeliveryGroups
}

// GetNumAssociatedDeliveryGroupsOk returns a tuple with the NumAssociatedDeliveryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetNumAssociatedDeliveryGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAssociatedDeliveryGroups) {
		return nil, false
	}
	return o.NumAssociatedDeliveryGroups, true
}

// HasNumAssociatedDeliveryGroups returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasNumAssociatedDeliveryGroups() bool {
	if o != nil && !IsNil(o.NumAssociatedDeliveryGroups) {
		return true
	}

	return false
}

// SetNumAssociatedDeliveryGroups gets a reference to the given int32 and assigns it to the NumAssociatedDeliveryGroups field.
func (o *ApplicationResponseModel) SetNumAssociatedDeliveryGroups(v int32) {
	o.NumAssociatedDeliveryGroups = &v
}

// GetNumAssociatedApplicationGroups returns the NumAssociatedApplicationGroups field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetNumAssociatedApplicationGroups() int32 {
	if o == nil || IsNil(o.NumAssociatedApplicationGroups) {
		var ret int32
		return ret
	}
	return *o.NumAssociatedApplicationGroups
}

// GetNumAssociatedApplicationGroupsOk returns a tuple with the NumAssociatedApplicationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetNumAssociatedApplicationGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAssociatedApplicationGroups) {
		return nil, false
	}
	return o.NumAssociatedApplicationGroups, true
}

// HasNumAssociatedApplicationGroups returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasNumAssociatedApplicationGroups() bool {
	if o != nil && !IsNil(o.NumAssociatedApplicationGroups) {
		return true
	}

	return false
}

// SetNumAssociatedApplicationGroups gets a reference to the given int32 and assigns it to the NumAssociatedApplicationGroups field.
func (o *ApplicationResponseModel) SetNumAssociatedApplicationGroups(v int32) {
	o.NumAssociatedApplicationGroups = &v
}

// GetAssociatedDeliveryGroupUuids returns the AssociatedDeliveryGroupUuids field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetAssociatedDeliveryGroupUuids() []string {
	if o == nil || IsNil(o.AssociatedDeliveryGroupUuids) {
		var ret []string
		return ret
	}
	return o.AssociatedDeliveryGroupUuids
}

// GetAssociatedDeliveryGroupUuidsOk returns a tuple with the AssociatedDeliveryGroupUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetAssociatedDeliveryGroupUuidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedDeliveryGroupUuids) {
		return nil, false
	}
	return o.AssociatedDeliveryGroupUuids, true
}

// HasAssociatedDeliveryGroupUuids returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasAssociatedDeliveryGroupUuids() bool {
	if o != nil && !IsNil(o.AssociatedDeliveryGroupUuids) {
		return true
	}

	return false
}

// SetAssociatedDeliveryGroupUuids gets a reference to the given []string and assigns it to the AssociatedDeliveryGroupUuids field.
func (o *ApplicationResponseModel) SetAssociatedDeliveryGroupUuids(v []string) {
	o.AssociatedDeliveryGroupUuids = v
}

// GetAssociatedApplicationGroupUuids returns the AssociatedApplicationGroupUuids field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetAssociatedApplicationGroupUuids() []string {
	if o == nil || IsNil(o.AssociatedApplicationGroupUuids) {
		var ret []string
		return ret
	}
	return o.AssociatedApplicationGroupUuids
}

// GetAssociatedApplicationGroupUuidsOk returns a tuple with the AssociatedApplicationGroupUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetAssociatedApplicationGroupUuidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedApplicationGroupUuids) {
		return nil, false
	}
	return o.AssociatedApplicationGroupUuids, true
}

// HasAssociatedApplicationGroupUuids returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasAssociatedApplicationGroupUuids() bool {
	if o != nil && !IsNil(o.AssociatedApplicationGroupUuids) {
		return true
	}

	return false
}

// SetAssociatedApplicationGroupUuids gets a reference to the given []string and assigns it to the AssociatedApplicationGroupUuids field.
func (o *ApplicationResponseModel) SetAssociatedApplicationGroupUuids(v []string) {
	o.AssociatedApplicationGroupUuids = v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *ApplicationResponseModel) GetZoneId() string {
	if o == nil || IsNil(o.ZoneId) {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseModel) GetZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *ApplicationResponseModel) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *ApplicationResponseModel) SetZoneId(v string) {
	o.ZoneId = &v
}

func (o ApplicationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	if !IsNil(o.Uid) {
		toSerialize["Uid"] = o.Uid
	}
	toSerialize["ApplicationFolder"] = o.ApplicationFolder
	toSerialize["ApplicationType"] = o.ApplicationType
	if !IsNil(o.ClientFolder) {
		toSerialize["ClientFolder"] = o.ClientFolder
	}
	toSerialize["ContainerScopes"] = o.ContainerScopes
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DoNotEnumerate) {
		toSerialize["DoNotEnumerate"] = o.DoNotEnumerate
	}
	toSerialize["Enabled"] = o.Enabled
	toSerialize["IconId"] = o.IconId
	if !IsNil(o.InstalledAppProperties) {
		toSerialize["InstalledAppProperties"] = o.InstalledAppProperties
	}
	if !IsNil(o.AppVAppProperties) {
		toSerialize["AppVAppProperties"] = o.AppVAppProperties
	}
	if !IsNil(o.ContentLocation) {
		toSerialize["ContentLocation"] = o.ContentLocation
	}
	toSerialize["Name"] = o.Name
	toSerialize["PublishedName"] = o.PublishedName
	toSerialize["Visible"] = o.Visible
	toSerialize["SharingKind"] = o.SharingKind
	if !IsNil(o.Tags) {
		toSerialize["Tags"] = o.Tags
	}
	if !IsNil(o.Tenants) {
		toSerialize["Tenants"] = o.Tenants
	}
	if !IsNil(o.CloudWorkspaceManaged) {
		toSerialize["CloudWorkspaceManaged"] = o.CloudWorkspaceManaged
	}
	if !IsNil(o.NumAssociatedDeliveryGroups) {
		toSerialize["NumAssociatedDeliveryGroups"] = o.NumAssociatedDeliveryGroups
	}
	if !IsNil(o.NumAssociatedApplicationGroups) {
		toSerialize["NumAssociatedApplicationGroups"] = o.NumAssociatedApplicationGroups
	}
	if !IsNil(o.AssociatedDeliveryGroupUuids) {
		toSerialize["AssociatedDeliveryGroupUuids"] = o.AssociatedDeliveryGroupUuids
	}
	if !IsNil(o.AssociatedApplicationGroupUuids) {
		toSerialize["AssociatedApplicationGroupUuids"] = o.AssociatedApplicationGroupUuids
	}
	if !IsNil(o.ZoneId) {
		toSerialize["ZoneId"] = o.ZoneId
	}
	return toSerialize, nil
}

type NullableApplicationResponseModel struct {
	value *ApplicationResponseModel
	isSet bool
}

func (v NullableApplicationResponseModel) Get() *ApplicationResponseModel {
	return v.value
}

func (v *NullableApplicationResponseModel) Set(val *ApplicationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponseModel(val *ApplicationResponseModel) *NullableApplicationResponseModel {
	return &NullableApplicationResponseModel{value: val, isSet: true}
}

func (v NullableApplicationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


