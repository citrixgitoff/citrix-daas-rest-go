/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// JobType Types of long-running jobs that may be executing.
type JobType string

// List of JobType
const (
	JOBTYPE_UNKNOWN JobType = "Unknown"
	JOBTYPE_BATCH JobType = "Batch"
	JOBTYPE_CREATE_DELIVERY_GROUP JobType = "CreateDeliveryGroup"
	JOBTYPE_UPDATE_DELIVERY_GROUP JobType = "UpdateDeliveryGroup"
	JOBTYPE_DELETE_DELIVERY_GROUP JobType = "DeleteDeliveryGroup"
	JOBTYPE_ADD_DELIVERY_GROUP_MACHINES JobType = "AddDeliveryGroupMachines"
	JOBTYPE_REMOVE_DELIVERY_GROUP_MACHINES JobType = "RemoveDeliveryGroupMachines"
	JOBTYPE_ADD_DELIVERY_GROUP_APPLICATIONS JobType = "AddDeliveryGroupApplications"
	JOBTYPE_REMOVE_DELIVERY_GROUP_APPLICATIONS JobType = "RemoveDeliveryGroupApplications"
	JOBTYPE_REMOVE_DELIVERY_GROUP_APPLICATION_GROUPS JobType = "RemoveDeliveryGroupApplicationGroups"
	JOBTYPE_GET_ADMIN_FOLDERS JobType = "GetAdminFolders"
	JOBTYPE_GET_APPLICATION_FOLDERS JobType = "GetApplicationFolders"
	JOBTYPE_GET_ADMIN_FOLDER_APPLICATIONS JobType = "GetAdminFolderApplications"
	JOBTYPE_GET_ADMIN_FOLDER_MACHINE_CATALOGS JobType = "GetAdminFolderMachineCatalogs"
	JOBTYPE_GET_APPLICATION_FOLDER_APPLICATIONS JobType = "GetApplicationFolderApplications"
	JOBTYPE_GET_APPLICATIONS JobType = "GetApplications"
	JOBTYPE_GET_APPLICATION_SESSIONS JobType = "GetApplicationSessions"
	JOBTYPE_GET_DELIVERY_GROUPS JobType = "GetDeliveryGroups"
	JOBTYPE_SEARCH_APPLICATIONS JobType = "SearchApplications"
	JOBTYPE_REFRESH_APP_V_APPLICATIONS JobType = "RefreshAppVApplications"
	JOBTYPE_TEST_DELIVERY_GROUP JobType = "TestDeliveryGroup"
	JOBTYPE_CREATE_TAG JobType = "CreateTag"
	JOBTYPE_UPDATE_TAG JobType = "UpdateTag"
	JOBTYPE_DELETE_TAG JobType = "DeleteTag"
	JOBTYPE_ADD_ICON JobType = "AddIcon"
	JOBTYPE_REMOVE_ICON JobType = "RemoveIcon"
	JOBTYPE_SET_TAG_APPLICATIONS JobType = "SetTagApplications"
	JOBTYPE_SET_TAG_APPLICATION_GROUPS JobType = "SetTagApplicationGroups"
	JOBTYPE_SET_TAG_DELIVERY_GROUPS JobType = "SetTagDeliveryGroups"
	JOBTYPE_SET_TAG_MACHINE_CATALOGS JobType = "SetTagMachineCatalogs"
	JOBTYPE_SET_TAG_MACHINES JobType = "SetTagMachines"
	JOBTYPE_GET_START_MENU_APPLICATIONS JobType = "GetStartMenuApplications"
	JOBTYPE_GET_DELIVERY_GROUP_APPLICATIONS JobType = "GetDeliveryGroupApplications"
	JOBTYPE_GET_DELIVERY_GROUPS_DESKTOPS JobType = "GetDeliveryGroupsDesktops"
	JOBTYPE_REBOOT_MACHINES JobType = "RebootMachines"
	JOBTYPE_SHUT_DOWN_MACHINES JobType = "ShutDownMachines"
	JOBTYPE_START_MACHINES JobType = "StartMachines"
	JOBTYPE_SUSPEND_MACHINES JobType = "SuspendMachines"
	JOBTYPE_RESUME_MACHINES JobType = "ResumeMachines"
	JOBTYPE_CREATE_ADMIN_FOLDER JobType = "CreateAdminFolder"
	JOBTYPE_CREATE_APPLICATION_FOLDER JobType = "CreateApplicationFolder"
	JOBTYPE_CREATE_REBOOT_SCHEDULE JobType = "CreateRebootSchedule"
	JOBTYPE_UPDATE_REBOOT_SCHEDULE JobType = "UpdateRebootSchedule"
	JOBTYPE_DELETE_REBOOT_SCHEDULE JobType = "DeleteRebootSchedule"
	JOBTYPE_CREATE_POWER_TIME_SCHEME JobType = "CreatePowerTimeScheme"
	JOBTYPE_UPDATE_POWER_TIME_SCHEME JobType = "UpdatePowerTimeScheme"
	JOBTYPE_DELETE_POWER_TIME_SCHEME JobType = "DeletePowerTimeScheme"
	JOBTYPE_CREATE_ADVANCED_ACCESS_POLICY JobType = "CreateAdvancedAccessPolicy"
	JOBTYPE_UPDATE_ADVANCED_ACCESS_POLICY JobType = "UpdateAdvancedAccessPolicy"
	JOBTYPE_DELETE_ADVANCED_ACCESS_POLICY JobType = "DeleteAdvancedAccessPolicy"
	JOBTYPE_ADD_MACHINE_CATALOG_MACHINE JobType = "AddMachineCatalogMachine"
	JOBTYPE_UPDATE_MACHINE_CATALOG_PROVISIONING_SCHEME JobType = "UpdateMachineCatalogProvisioningScheme"
	JOBTYPE_CREATE_MACHINE_CATALOG JobType = "CreateMachineCatalog"
	JOBTYPE_CREATE_MACHINE_IDENTITY JobType = "CreateMachineIdentity"
	JOBTYPE_DELETE_MACHINE_IDENTITY JobType = "DeleteMachineIdentity"
	JOBTYPE_UPDATE_MACHINE_IDENTITY JobType = "UpdateMachineIdentity"
	JOBTYPE_UPDATE_USER_IDENTITY JobType = "UpdateUserIdentity"
	JOBTYPE_UPDATE_MACHINE_CATALOG JobType = "UpdateMachineCatalog"
	JOBTYPE_DELETE_MACHINE_CATALOG JobType = "DeleteMachineCatalog"
	JOBTYPE_TEST_MACHINE_CATALOG JobType = "TestMachineCatalog"
	JOBTYPE_LOGOFF_SESSIONS JobType = "LogoffSessions"
	JOBTYPE_DISCONNECT_SESSIONS JobType = "DisconnectSessions"
	JOBTYPE_LOGOFF_MACHINE_SESSIONS JobType = "LogoffMachineSessions"
	JOBTYPE_DELETE_MACHINE JobType = "DeleteMachine"
	JOBTYPE_SEND_SESSION_MESSAGE JobType = "SendSessionMessage"
	JOBTYPE_ADD_APPLICATIONS JobType = "AddApplications"
	JOBTYPE_GET_APPLICATION_GROUPS JobType = "GetApplicationGroups"
	JOBTYPE_GET_APPLICATION_GROUP_APPLICATIONS JobType = "GetApplicationGroupApplications"
	JOBTYPE_GET_APPLICATION_GROUP_DELIVERY_GROUPS JobType = "GetApplicationGroupDeliveryGroups"
	JOBTYPE_CREATE_APPLICATION_GROUP JobType = "CreateApplicationGroup"
	JOBTYPE_DELETE_APPLICATION JobType = "DeleteApplication"
	JOBTYPE_DELETE_APPLICATION_GROUP JobType = "DeleteApplicationGroup"
	JOBTYPE_SEARCH_APPLICATION_GROUPS JobType = "SearchApplicationGroups"
	JOBTYPE_DELETE_ADMIN_FOLDER JobType = "DeleteAdminFolder"
	JOBTYPE_DELETE_APPLICATION_FOLDER JobType = "DeleteApplicationFolder"
	JOBTYPE_REMOVE_APPLICATION_GROUP_APPLICATIONS JobType = "RemoveApplicationGroupApplications"
	JOBTYPE_GET_APPLICATION_APPLICATION_GROUPS JobType = "GetApplicationApplicationGroups"
	JOBTYPE_GET_APPLICATION_DELIVERY_GROUPS JobType = "GetApplicationDeliveryGroups"
	JOBTYPE_UPDATE_APPLICATION JobType = "UpdateApplication"
	JOBTYPE_UPDATE_APPLICATION_GROUP JobType = "UpdateApplicationGroup"
	JOBTYPE_UPDATE_ADMIN_FOLDER JobType = "UpdateAdminFolder"
	JOBTYPE_UPDATE_APPLICATION_FOLDER JobType = "UpdateApplicationFolder"
	JOBTYPE_SEARCH_START_MENU_APPLICATIONS JobType = "SearchStartMenuApplications"
	JOBTYPE_APPLICATIONS_DISCOVERY JobType = "ApplicationsDiscovery"
	JOBTYPE_MODIFY_FTA JobType = "ModifyFTA"
	JOBTYPE_IMPORT_FTA JobType = "ImportFTA"
	JOBTYPE_GET_SITE_STATUS JobType = "GetSiteStatus"
	JOBTYPE_TEST_SITE JobType = "TestSite"
	JOBTYPE_CREATE_HYPERVISOR JobType = "CreateHypervisor"
	JOBTYPE_CREATE_HYPERVISOR_RESOURCE_POOL JobType = "CreateHypervisorResourcePool"
	JOBTYPE_UPDATE_HYPERVISOR JobType = "UpdateHypervisor"
	JOBTYPE_UPDATE_HYPERVISOR_RESOURCE_POOL JobType = "UpdateHypervisorResourcePool"
	JOBTYPE_HYPERVISOR_VM_VALIDATION JobType = "HypervisorVmValidation"
	JOBTYPE_TEST_HYPERVISOR JobType = "TestHypervisor"
	JOBTYPE_GET_HYPERVISOR_ORPHANED_RESOURCES JobType = "GetHypervisorOrphanedResources"
	JOBTYPE_TEST_HYPERVISOR_RESOURCE_POOL JobType = "TestHypervisorResourcePool"
	JOBTYPE_CREATE_STORE_FRONT_SERVER JobType = "CreateStoreFrontServer"
	JOBTYPE_DELETE_STORE_FRONT_SERVER JobType = "DeleteStoreFrontServer"
	JOBTYPE_UPDATE_STORE_FRONT_SERVER JobType = "UpdateStoreFrontServer"
	JOBTYPE_UPDATE_IMAGE JobType = "UpdateImage"
	JOBTYPE_SET_PROV_SCHEME JobType = "SetProvScheme"
	JOBTYPE_ADD_MACHINES JobType = "AddMachines"
	JOBTYPE_UPDATE_MACHINE_CATALOG_MACHINE JobType = "UpdateMachineCatalogMachine"
	JOBTYPE_ROLLBACK_MACHINE_CATALOG_PROVISIONING_SCHEME JobType = "RollbackMachineCatalogProvisioningScheme"
	JOBTYPE_GET_HYPERVISOR_RESOURCE_POOL JobType = "GetHypervisorResourcePool"
	JOBTYPE_GET_HYPERVISORS_AND_RESOURCE_POOLS JobType = "GetHypervisorsAndResourcePools"
	JOBTYPE_GET_HYPERVISOR_RESOURCE_POOL_RESOURCES JobType = "GetHypervisorResourcePoolResources"
	JOBTYPE_SEARCH_HYPERVISOR_RESOURCE_POOL_RESOURCES JobType = "SearchHypervisorResourcePoolResources"
	JOBTYPE_GET_HYPERVISOR_ALL_RESOURCES JobType = "GetHypervisorAllResources"
	JOBTYPE_GET_HYPERVISOR_ALL_RESOURCES_WITHOUT_CONNECTION JobType = "GetHypervisorAllResourcesWithoutConnection"
	JOBTYPE_GET_HYPERVISOR_RESOURCE_POOLS JobType = "GetHypervisorResourcePools"
	JOBTYPE_GET_HYPERVISOR JobType = "GetHypervisor"
	JOBTYPE_GET_HYPERVISORS JobType = "GetHypervisors"
	JOBTYPE_DELETE_HYPERVISOR JobType = "DeleteHypervisor"
	JOBTYPE_DELETE_RESOURCE_POOL JobType = "DeleteResourcePool"
	JOBTYPE_GET_HYPERVISOR_SUPPORTED_PLUGINS JobType = "GetHypervisorSupportedPlugins"
	JOBTYPE_GET_HYPERVISOR_SERVER_HA_ADDRESSES JobType = "GetHypervisorServerHAAddresses"
	JOBTYPE_GET_SESSION_APPLICATIONS JobType = "GetSessionApplications"
	JOBTYPE_GET_OPERATIONS JobType = "GetOperations"
	JOBTYPE_DO_OPERATION_SEARCH JobType = "DoOperationSearch"
	JOBTYPE_FETCH_EXISTING_OPERATION_LABELS JobType = "FetchExistingOperationLabels"
	JOBTYPE_UPDATE_OPERATION JobType = "UpdateOperation"
	JOBTYPE_GET_SESSIONS JobType = "GetSessions"
	JOBTYPE_DO_SESSION_SEARCH JobType = "DoSessionSearch"
	JOBTYPE_GET_MACHINES JobType = "GetMachines"
	JOBTYPE_DO_MACHINE_SEARCH JobType = "DoMachineSearch"
	JOBTYPE_GET_PVS_SITES JobType = "GetPvsSites"
	JOBTYPE_GET_PVS_COLLECTIONS JobType = "GetPvsCollections"
	JOBTYPE_GET_PVS_MACHINES_FOR_CATALOG JobType = "GetPvsMachinesForCatalog"
	JOBTYPE_GET_MACHINE_CATALOG_MACHINE_ACCOUNTS JobType = "GetMachineCatalogMachineAccounts"
	JOBTYPE_GET_MACHINE_START_MENU_SHORTCUT_ICON JobType = "GetMachineStartMenuShortcutIcon"
	JOBTYPE_GET_MACHINE_CATALOGS JobType = "GetMachineCatalogs"
	JOBTYPE_SEARCH_MACHINE_CATALOGS JobType = "SearchMachineCatalogs"
	JOBTYPE_GET_MACHINE_CATALOG_DETAILS JobType = "GetMachineCatalogDetails"
	JOBTYPE_GET_APP_V_SERVERS JobType = "GetAppVServers"
	JOBTYPE_IMPORT_APP_V_PACKAGES JobType = "ImportAppVPackages"
	JOBTYPE_ADD_APP_V_SERVER JobType = "AddAppVServer"
	JOBTYPE_UPDATE_APP_V_SERVER JobType = "UpdateAppVServer"
	JOBTYPE_REMOVE_APP_V_SERVER JobType = "RemoveAppVServer"
	JOBTYPE_GET_APP_V_PACKAGES JobType = "GetAppVPackages"
	JOBTYPE_REMOVE_APP_V_PACKAGE JobType = "RemoveAppVPackage"
	JOBTYPE_GET_APP_V_ISOLATION_GROUPS JobType = "GetAppVIsolationGroups"
	JOBTYPE_GET_APP_V_PACKAGE_BROKER_APPLICATIONS JobType = "GetAppVPackageBrokerApplications"
	JOBTYPE_GET_APP_V_PACKAGE_DELVERY_GROUPS JobType = "GetAppVPackageDelveryGroups"
	JOBTYPE_DELETE_APP_V_ISOLATION_GROUP JobType = "DeleteAppVIsolationGroup"
	JOBTYPE_CREATE_APP_V_ISOLATION_GROUP JobType = "CreateAppVIsolationGroup"
	JOBTYPE_UPDATE_APP_V_ISOLATION_GROUP JobType = "UpdateAppVIsolationGroup"
	JOBTYPE_GET_APP_V_ISOLATION_GROUP JobType = "GetAppVIsolationGroup"
	JOBTYPE_GET_APPLICATION_FTAS JobType = "GetApplicationFtas"
	JOBTYPE_GET_MACHINE_CATALOG_MACHINES JobType = "GetMachineCatalogMachines"
	JOBTYPE_GET_MACHINE_CATALOG_DELIVERY_GROUP_ASSOCIATIONS JobType = "GetMachineCatalogDeliveryGroupAssociations"
	JOBTYPE_DO_DELIVERY_GROUP_SEARCH JobType = "DoDeliveryGroupSearch"
	JOBTYPE_GET_MACHINE_CATALOG_LAST_MASTER_IMAGE JobType = "GetMachineCatalogLastMasterImage"
	JOBTYPE_DO_ZONE_SEARCH JobType = "DoZoneSearch"
	JOBTYPE_CREATE_ZONE JobType = "CreateZone"
	JOBTYPE_EDIT_ZONE JobType = "EditZone"
	JOBTYPE_DELETE_ZONE JobType = "DeleteZone"
	JOBTYPE_ADD_ITEMS_INTO_ZONE JobType = "AddItemsIntoZone"
	JOBTYPE_MOVE_ITEMS_INTO_ZONE JobType = "MoveItemsIntoZone"
	JOBTYPE_REMOVE_ITEMS_FROM_ZONE JobType = "RemoveItemsFromZone"
	JOBTYPE_START_DELIVERY_GROUP_REBOOT_CYCLE JobType = "StartDeliveryGroupRebootCycle"
	JOBTYPE_GET_ADMIN_ADMINISTRATORS JobType = "GetAdminAdministrators"
	JOBTYPE_CREATE_ADMIN_SCOPE JobType = "CreateAdminScope"
	JOBTYPE_DELETE_ADMIN_SCOPE JobType = "DeleteAdminScope"
	JOBTYPE_CREATE_ADMIN_ROLE JobType = "CreateAdminRole"
	JOBTYPE_UPDATE_ADMIN_ROLE JobType = "UpdateAdminRole"
	JOBTYPE_DELETE_ADMIN_ROLE JobType = "DeleteAdminRole"
	JOBTYPE_UPDATE_ADMIN_SCOPE JobType = "UpdateAdminScope"
	JOBTYPE_SET_BROKER_HYPERVISOR_CONNECTION JobType = "SetBrokerHypervisorConnection"
	JOBTYPE_RESET_BROKER_HYPERVISOR_CONNECTION JobType = "ResetBrokerHypervisorConnection"
	JOBTYPE_GET_ALL_CATALOG_ACTIONS JobType = "GetAllCatalogActions"
	JOBTYPE_GET_CATALOG_ACTION JobType = "GetCatalogAction"
	JOBTYPE_GET_SINGLE_CATALOG_ACTION JobType = "GetSingleCatalogAction"
	JOBTYPE_CANCEL_CATALOG_ACTION JobType = "CancelCatalogAction"
	JOBTYPE_DELETE_ALL_CATALOG_ACTIONS JobType = "DeleteAllCatalogActions"
	JOBTYPE_DELETE_CATALOG_ACTIONS JobType = "DeleteCatalogActions"
	JOBTYPE_HIDE_SESSION JobType = "HideSession"
	JOBTYPE_UNHIDE_SESSION JobType = "UnhideSession"
	JOBTYPE_CREATE_MACHINE_CATALOG_UPGRADE_SCHEDULE JobType = "CreateMachineCatalogUpgradeSchedule"
	JOBTYPE_REMOVE_MACHINE_CATALOG_UPGRADE_SCHEDULE JobType = "RemoveMachineCatalogUpgradeSchedule"
	JOBTYPE_UPDATE_MACHINE_CATALOG_UPGRADE_SCHEDULE JobType = "UpdateMachineCatalogUpgradeSchedule"
	JOBTYPE_GET_HYPERVISOR_MACHINE_CATALOGS JobType = "GetHypervisorMachineCatalogs"
	JOBTYPE_GET_HYPERVISOR_RESOURCE_POOL_MACHINE_CATALOGS JobType = "GetHypervisorResourcePoolMachineCatalogs"
	JOBTYPE_GET_HYPERVISOR_ADMINISTRATORS JobType = "GetHypervisorAdministrators"
	JOBTYPE_GET_HYPERVISOR_RESOURCE_POOL_ADMINISTRATORS JobType = "GetHypervisorResourcePoolAdministrators"
	JOBTYPE_CREATE_MACHINE_UPGRADE_SCHEDULE JobType = "CreateMachineUpgradeSchedule"
	JOBTYPE_UPDATE_MACHINE_UPGRADE_SCHEDULE JobType = "UpdateMachineUpgradeSchedule"
	JOBTYPE_REMOVE_MACHINE_UPGRADE_SCHEDULE JobType = "RemoveMachineUpgradeSchedule"
	JOBTYPE_CREATE_MACHINES_UPGRADE_SCHEUDLE JobType = "CreateMachinesUpgradeScheudle"
	JOBTYPE_GET_APP_LIB_PACKAGE_DISCOVERY_PROFILES JobType = "GetAppLibPackageDiscoveryProfiles"
	JOBTYPE_GET_APP_LIB_PACKAGE_DISCOVERY_PROFILE JobType = "GetAppLibPackageDiscoveryProfile"
	JOBTYPE_CREATE_APP_LIB_PACKAGE_DISCOVERY_PROFILE JobType = "CreateAppLibPackageDiscoveryProfile"
	JOBTYPE_UPDATE_APP_LIB_PACKAGE_DISCOVERY_PROFILE JobType = "UpdateAppLibPackageDiscoveryProfile"
	JOBTYPE_REMOVE_APP_LIB_PACKAGE_DISCOVERY_PROFILE JobType = "RemoveAppLibPackageDiscoveryProfile"
	JOBTYPE_VALIDATE_HYPERVISOR_RESOURCE_POOL_RESOURCE JobType = "ValidateHypervisorResourcePoolResource"
	JOBTYPE_CREATE_APP_LIB_PACKAGE_DISCOVERY JobType = "CreateAppLibPackageDiscovery"
	JOBTYPE_GET_APP_LIB_PACKAGE_DISCOVERY JobType = "GetAppLibPackageDiscovery"
	JOBTYPE_GET_APP_LIB_PACKAGE_DISCOVERIES JobType = "GetAppLibPackageDiscoveries"
	JOBTYPE_SEND_CAS_EVENT JobType = "SendCasEvent"
	JOBTYPE_ADD_DELIVERY_GROUP_POLICY_SET JobType = "AddDeliveryGroupPolicySet"
	JOBTYPE_REMOVE_DELIVERY_GROUP_POLICY_SET JobType = "RemoveDeliveryGroupPolicySet"
	JOBTYPE_START_SESSION_RECORDING JobType = "StartSessionRecording"
	JOBTYPE_STOP_SESSION_RECORDING JobType = "StopSessionRecording"
	JOBTYPE_GET_LICENSE_SUMMARY JobType = "GetLicenseSummary"
	JOBTYPE_GET_DOMAINS_IDENTITY JobType = "GetDomainsIdentity"
	JOBTYPE_GET_LICENSE_PERMISSION JobType = "GetLicensePermission"
	JOBTYPE_REMOVE_MACHINE_POWER_ACTION_SCHEDULE JobType = "RemoveMachinePowerActionSchedule"
	JOBTYPE_GET_LICENSE_CERTIFICATE JobType = "GetLicenseCertificate"
	JOBTYPE_SET_LICENSE_SERVER JobType = "SetLicenseServer"
	JOBTYPE_GET_LICENSE_ALERT JobType = "GetLicenseAlert"
	JOBTYPE_ADD_LICENSE JobType = "AddLicense"
	JOBTYPE_DELETE_DELIVERY_CONTROLLER JobType = "DeleteDeliveryController"
	JOBTYPE_GENERATE_DELIVERY_CONTROLLER_DELETE_SCRIPT JobType = "GenerateDeliveryControllerDeleteScript"
	JOBTYPE_CHANGE_DATABASE JobType = "ChangeDatabase"
	JOBTYPE_GET_LICENSE_ENTITLEMENT JobType = "GetLicenseEntitlement"
	JOBTYPE_ALLOCATE_LICENSE JobType = "AllocateLicense"
	JOBTYPE_VALIDATE_MACHINE_CATALOG_VDA_COMPONENTS_AND_FEATURES_SELECTION JobType = "ValidateMachineCatalogVDAComponentsAndFeaturesSelection"
	JOBTYPE_REMOVE_IMAGE_DEFINITION JobType = "RemoveImageDefinition"
	JOBTYPE_CREATE_IMAGE_VERSION JobType = "CreateImageVersion"
	JOBTYPE_SET_IMAGE_VERSION JobType = "SetImageVersion"
	JOBTYPE_DELETE_IMAGE_VERSION JobType = "DeleteImageVersion"
	JOBTYPE_GET_ALL_MACHINE_IDENTITY_POOLS JobType = "GetAllMachineIdentityPools"
	JOBTYPE_GET_ALL_MACHINE_IDENTITY_POOL JobType = "GetAllMachineIdentityPool"
	JOBTYPE_GET_PROV_SCHEME_REFERENCES JobType = "GetProvSchemeReferences"
	JOBTYPE_REQUEST_PROV_VM_UPDATE JobType = "RequestProvVmUpdate"
	JOBTYPE_CREATE_ADMIN JobType = "CreateAdmin"
	JOBTYPE_UPDATE_ADMIN JobType = "UpdateAdmin"
	JOBTYPE_DELETE_ADMIN JobType = "DeleteAdmin"
	JOBTYPE_GET_BACKUPS JobType = "GetBackups"
	JOBTYPE_GET_BACKUP_HISTORY JobType = "GetBackupHistory"
	JOBTYPE_GET_BACKUP_MEMBERS JobType = "GetBackupMembers"
	JOBTYPE_PIN_BACKUP JobType = "PinBackup"
	JOBTYPE_GET_PVS_SITE JobType = "GetPvsSite"
	JOBTYPE_GET_PVS_V_DISK JobType = "GetPvsVDisk"
	JOBTYPE_GET_PVS_STORE JobType = "GetPvsStore"
	JOBTYPE_RESET_VM_DISK JobType = "ResetVMDisk"
	JOBTYPE_GET_GPO_POLICY_SETS JobType = "GetGpoPolicySets"
	JOBTYPE_GET_GPO_POLICIES JobType = "GetGpoPolicies"
	JOBTYPE_GET_GPO_FILTERS JobType = "GetGpoFilters"
	JOBTYPE_GET_GPO_SETTINGS JobType = "GetGpoSettings"
	JOBTYPE_REPAIR_MACHINE_CATALOG_ACCOUNTS JobType = "RepairMachineCatalogAccounts"
	JOBTYPE_IMPORT_PROVISIONED_VIRTUAL_MACHINES JobType = "ImportProvisionedVirtualMachines"
	JOBTYPE_VALIDATE_MACHINE_CATALOG_CREATION JobType = "ValidateMachineCatalogCreation"
)

// All allowed values of JobType enum
var AllowedJobTypeEnumValues = []JobType{
	"Unknown",
	"Batch",
	"CreateDeliveryGroup",
	"UpdateDeliveryGroup",
	"DeleteDeliveryGroup",
	"AddDeliveryGroupMachines",
	"RemoveDeliveryGroupMachines",
	"AddDeliveryGroupApplications",
	"RemoveDeliveryGroupApplications",
	"RemoveDeliveryGroupApplicationGroups",
	"GetAdminFolders",
	"GetApplicationFolders",
	"GetAdminFolderApplications",
	"GetAdminFolderMachineCatalogs",
	"GetApplicationFolderApplications",
	"GetApplications",
	"GetApplicationSessions",
	"GetDeliveryGroups",
	"SearchApplications",
	"RefreshAppVApplications",
	"TestDeliveryGroup",
	"CreateTag",
	"UpdateTag",
	"DeleteTag",
	"AddIcon",
	"RemoveIcon",
	"SetTagApplications",
	"SetTagApplicationGroups",
	"SetTagDeliveryGroups",
	"SetTagMachineCatalogs",
	"SetTagMachines",
	"GetStartMenuApplications",
	"GetDeliveryGroupApplications",
	"GetDeliveryGroupsDesktops",
	"RebootMachines",
	"ShutDownMachines",
	"StartMachines",
	"SuspendMachines",
	"ResumeMachines",
	"CreateAdminFolder",
	"CreateApplicationFolder",
	"CreateRebootSchedule",
	"UpdateRebootSchedule",
	"DeleteRebootSchedule",
	"CreatePowerTimeScheme",
	"UpdatePowerTimeScheme",
	"DeletePowerTimeScheme",
	"CreateAdvancedAccessPolicy",
	"UpdateAdvancedAccessPolicy",
	"DeleteAdvancedAccessPolicy",
	"AddMachineCatalogMachine",
	"UpdateMachineCatalogProvisioningScheme",
	"CreateMachineCatalog",
	"CreateMachineIdentity",
	"DeleteMachineIdentity",
	"UpdateMachineIdentity",
	"UpdateUserIdentity",
	"UpdateMachineCatalog",
	"DeleteMachineCatalog",
	"TestMachineCatalog",
	"LogoffSessions",
	"DisconnectSessions",
	"LogoffMachineSessions",
	"DeleteMachine",
	"SendSessionMessage",
	"AddApplications",
	"GetApplicationGroups",
	"GetApplicationGroupApplications",
	"GetApplicationGroupDeliveryGroups",
	"CreateApplicationGroup",
	"DeleteApplication",
	"DeleteApplicationGroup",
	"SearchApplicationGroups",
	"DeleteAdminFolder",
	"DeleteApplicationFolder",
	"RemoveApplicationGroupApplications",
	"GetApplicationApplicationGroups",
	"GetApplicationDeliveryGroups",
	"UpdateApplication",
	"UpdateApplicationGroup",
	"UpdateAdminFolder",
	"UpdateApplicationFolder",
	"SearchStartMenuApplications",
	"ApplicationsDiscovery",
	"ModifyFTA",
	"ImportFTA",
	"GetSiteStatus",
	"TestSite",
	"CreateHypervisor",
	"CreateHypervisorResourcePool",
	"UpdateHypervisor",
	"UpdateHypervisorResourcePool",
	"HypervisorVmValidation",
	"TestHypervisor",
	"GetHypervisorOrphanedResources",
	"TestHypervisorResourcePool",
	"CreateStoreFrontServer",
	"DeleteStoreFrontServer",
	"UpdateStoreFrontServer",
	"UpdateImage",
	"SetProvScheme",
	"AddMachines",
	"UpdateMachineCatalogMachine",
	"RollbackMachineCatalogProvisioningScheme",
	"GetHypervisorResourcePool",
	"GetHypervisorsAndResourcePools",
	"GetHypervisorResourcePoolResources",
	"SearchHypervisorResourcePoolResources",
	"GetHypervisorAllResources",
	"GetHypervisorAllResourcesWithoutConnection",
	"GetHypervisorResourcePools",
	"GetHypervisor",
	"GetHypervisors",
	"DeleteHypervisor",
	"DeleteResourcePool",
	"GetHypervisorSupportedPlugins",
	"GetHypervisorServerHAAddresses",
	"GetSessionApplications",
	"GetOperations",
	"DoOperationSearch",
	"FetchExistingOperationLabels",
	"UpdateOperation",
	"GetSessions",
	"DoSessionSearch",
	"GetMachines",
	"DoMachineSearch",
	"GetPvsSites",
	"GetPvsCollections",
	"GetPvsMachinesForCatalog",
	"GetMachineCatalogMachineAccounts",
	"GetMachineStartMenuShortcutIcon",
	"GetMachineCatalogs",
	"SearchMachineCatalogs",
	"GetMachineCatalogDetails",
	"GetAppVServers",
	"ImportAppVPackages",
	"AddAppVServer",
	"UpdateAppVServer",
	"RemoveAppVServer",
	"GetAppVPackages",
	"RemoveAppVPackage",
	"GetAppVIsolationGroups",
	"GetAppVPackageBrokerApplications",
	"GetAppVPackageDelveryGroups",
	"DeleteAppVIsolationGroup",
	"CreateAppVIsolationGroup",
	"UpdateAppVIsolationGroup",
	"GetAppVIsolationGroup",
	"GetApplicationFtas",
	"GetMachineCatalogMachines",
	"GetMachineCatalogDeliveryGroupAssociations",
	"DoDeliveryGroupSearch",
	"GetMachineCatalogLastMasterImage",
	"DoZoneSearch",
	"CreateZone",
	"EditZone",
	"DeleteZone",
	"AddItemsIntoZone",
	"MoveItemsIntoZone",
	"RemoveItemsFromZone",
	"StartDeliveryGroupRebootCycle",
	"GetAdminAdministrators",
	"CreateAdminScope",
	"DeleteAdminScope",
	"CreateAdminRole",
	"UpdateAdminRole",
	"DeleteAdminRole",
	"UpdateAdminScope",
	"SetBrokerHypervisorConnection",
	"ResetBrokerHypervisorConnection",
	"GetAllCatalogActions",
	"GetCatalogAction",
	"GetSingleCatalogAction",
	"CancelCatalogAction",
	"DeleteAllCatalogActions",
	"DeleteCatalogActions",
	"HideSession",
	"UnhideSession",
	"CreateMachineCatalogUpgradeSchedule",
	"RemoveMachineCatalogUpgradeSchedule",
	"UpdateMachineCatalogUpgradeSchedule",
	"GetHypervisorMachineCatalogs",
	"GetHypervisorResourcePoolMachineCatalogs",
	"GetHypervisorAdministrators",
	"GetHypervisorResourcePoolAdministrators",
	"CreateMachineUpgradeSchedule",
	"UpdateMachineUpgradeSchedule",
	"RemoveMachineUpgradeSchedule",
	"CreateMachinesUpgradeScheudle",
	"GetAppLibPackageDiscoveryProfiles",
	"GetAppLibPackageDiscoveryProfile",
	"CreateAppLibPackageDiscoveryProfile",
	"UpdateAppLibPackageDiscoveryProfile",
	"RemoveAppLibPackageDiscoveryProfile",
	"ValidateHypervisorResourcePoolResource",
	"CreateAppLibPackageDiscovery",
	"GetAppLibPackageDiscovery",
	"GetAppLibPackageDiscoveries",
	"SendCasEvent",
	"AddDeliveryGroupPolicySet",
	"RemoveDeliveryGroupPolicySet",
	"StartSessionRecording",
	"StopSessionRecording",
	"GetLicenseSummary",
	"GetDomainsIdentity",
	"GetLicensePermission",
	"RemoveMachinePowerActionSchedule",
	"GetLicenseCertificate",
	"SetLicenseServer",
	"GetLicenseAlert",
	"AddLicense",
	"DeleteDeliveryController",
	"GenerateDeliveryControllerDeleteScript",
	"ChangeDatabase",
	"GetLicenseEntitlement",
	"AllocateLicense",
	"ValidateMachineCatalogVDAComponentsAndFeaturesSelection",
	"RemoveImageDefinition",
	"CreateImageVersion",
	"SetImageVersion",
	"DeleteImageVersion",
	"GetAllMachineIdentityPools",
	"GetAllMachineIdentityPool",
	"GetProvSchemeReferences",
	"RequestProvVmUpdate",
	"CreateAdmin",
	"UpdateAdmin",
	"DeleteAdmin",
	"GetBackups",
	"GetBackupHistory",
	"GetBackupMembers",
	"PinBackup",
	"GetPvsSite",
	"GetPvsVDisk",
	"GetPvsStore",
	"ResetVMDisk",
	"GetGpoPolicySets",
	"GetGpoPolicies",
	"GetGpoFilters",
	"GetGpoSettings",
	"RepairMachineCatalogAccounts",
	"ImportProvisionedVirtualMachines",
	"ValidateMachineCatalogCreation",
}

func (v *JobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobType(value)
	for _, existing := range AllowedJobTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobType", value)
}

// NewJobTypeFromValue returns a pointer to a valid JobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTypeFromValue(v string) (*JobType, error) {
	ev := JobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobType: valid values are %v", v, AllowedJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobType) IsValid() bool {
	for _, existing := range AllowedJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobType value
func (v JobType) Ptr() *JobType {
	return &v
}

type NullableJobType struct {
	value *JobType
	isSet bool
}

func (v NullableJobType) Get() *JobType {
	return v.value
}

func (v *NullableJobType) Set(val *JobType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobType(val *JobType) *NullableJobType {
	return &NullableJobType{value: val, isSet: true}
}

func (v NullableJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

