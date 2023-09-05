/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: techpreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// MachineCatalogWarningSubtype Sub-types of warnings that may occur on a machine catalog.
type MachineCatalogWarningSubtype string

// List of MachineCatalogWarningSubtype
const (
	MACHINECATALOGWARNINGSUBTYPE_UNKNOWN MachineCatalogWarningSubtype = "Unknown"
	MACHINECATALOGWARNINGSUBTYPE_NO_WARNING MachineCatalogWarningSubtype = "NoWarning"
	MACHINECATALOGWARNINGSUBTYPE_NO_POWERED_ON_VM MachineCatalogWarningSubtype = "NoPoweredOnVm"
	MACHINECATALOGWARNINGSUBTYPE_LICENSING_TYPE_NOT_CONFIGURED MachineCatalogWarningSubtype = "LicensingTypeNotConfigured"
	MACHINECATALOGWARNINGSUBTYPE_CHECK_NOT_SUPPORTED MachineCatalogWarningSubtype = "CheckNotSupported"
	MACHINECATALOGWARNINGSUBTYPE_CHECK_NOT_COMPLETED MachineCatalogWarningSubtype = "CheckNotCompleted"
	MACHINECATALOGWARNINGSUBTYPE_CHECK_FAILED MachineCatalogWarningSubtype = "CheckFailed"
	MACHINECATALOGWARNINGSUBTYPE_NO_NEW_RDS_CONNECTIONS_ALLOWED MachineCatalogWarningSubtype = "NoNewRdsConnectionsAllowed"
	MACHINECATALOGWARNINGSUBTYPE_OUT_OF_GRACE_PERIOD MachineCatalogWarningSubtype = "OutOfGracePeriod"
	MACHINECATALOGWARNINGSUBTYPE_PERSONAL_TERMINAL_SERVER MachineCatalogWarningSubtype = "PersonalTerminalServer"
	MACHINECATALOGWARNINGSUBTYPE_REMOTE_DESKTOP_FOR_ADMINISTRATION MachineCatalogWarningSubtype = "RemoteDesktopForAdministration"
	MACHINECATALOGWARNINGSUBTYPE_PER_USER_ALLOWED MachineCatalogWarningSubtype = "PerUserAllowed"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_NO_LICENSE_SERVER MachineCatalogWarningSubtype = "PerDeviceNoLicenseServer"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_INCOMPATIBLE_LICENSE_SERVER MachineCatalogWarningSubtype = "PerDeviceIncompatibleLicenseServer"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_NO_REQUIRED_LICENSE_SERVER MachineCatalogWarningSubtype = "PerDeviceNoRequiredLicenseServer"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_NO_CAL_AND_NO_LICENSE_SERVER MachineCatalogWarningSubtype = "PerDeviceNoCalAndNoLicenseServer"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_ISSUE_MARKED_TEMPORARY_LICENSE MachineCatalogWarningSubtype = "PerDeviceIssueMarkedTemporaryLicense"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_UNMARKED_TEMPORARY_LICENSE MachineCatalogWarningSubtype = "PerDeviceUnmarkedTemporaryLicense"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_UNEXPIRED_TEMPORARY_LICENSE_NOT_UPGRADED MachineCatalogWarningSubtype = "PerDeviceUnexpiredTemporaryLicenseNotUpgraded"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_USE_UNEXPIRED_UNMARKED_TEMPORARY_LICENSE MachineCatalogWarningSubtype = "PerDeviceUseUnexpiredUnmarkedTemporaryLicense"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_EXPIRED_TEMPORARY_LICENSE_NOT_UPGRADED MachineCatalogWarningSubtype = "PerDeviceExpiredTemporaryLicenseNotUpgraded"
	MACHINECATALOGWARNINGSUBTYPE_PER_DEVICE_EXPIRED_UNMARKED_TEMPORARY_LICENSE MachineCatalogWarningSubtype = "PerDeviceExpiredUnmarkedTemporaryLicense"
	MACHINECATALOGWARNINGSUBTYPE_MACHINE_VDA_UPGRADE_FAILED MachineCatalogWarningSubtype = "MachineVdaUpgradeFailed"
)

// All allowed values of MachineCatalogWarningSubtype enum
var AllowedMachineCatalogWarningSubtypeEnumValues = []MachineCatalogWarningSubtype{
	"Unknown",
	"NoWarning",
	"NoPoweredOnVm",
	"LicensingTypeNotConfigured",
	"CheckNotSupported",
	"CheckNotCompleted",
	"CheckFailed",
	"NoNewRdsConnectionsAllowed",
	"OutOfGracePeriod",
	"PersonalTerminalServer",
	"RemoteDesktopForAdministration",
	"PerUserAllowed",
	"PerDeviceNoLicenseServer",
	"PerDeviceIncompatibleLicenseServer",
	"PerDeviceNoRequiredLicenseServer",
	"PerDeviceNoCalAndNoLicenseServer",
	"PerDeviceIssueMarkedTemporaryLicense",
	"PerDeviceUnmarkedTemporaryLicense",
	"PerDeviceUnexpiredTemporaryLicenseNotUpgraded",
	"PerDeviceUseUnexpiredUnmarkedTemporaryLicense",
	"PerDeviceExpiredTemporaryLicenseNotUpgraded",
	"PerDeviceExpiredUnmarkedTemporaryLicense",
	"MachineVdaUpgradeFailed",
}

func (v *MachineCatalogWarningSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MachineCatalogWarningSubtype(value)
	for _, existing := range AllowedMachineCatalogWarningSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MachineCatalogWarningSubtype", value)
}

// NewMachineCatalogWarningSubtypeFromValue returns a pointer to a valid MachineCatalogWarningSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineCatalogWarningSubtypeFromValue(v string) (*MachineCatalogWarningSubtype, error) {
	ev := MachineCatalogWarningSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineCatalogWarningSubtype: valid values are %v", v, AllowedMachineCatalogWarningSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineCatalogWarningSubtype) IsValid() bool {
	for _, existing := range AllowedMachineCatalogWarningSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineCatalogWarningSubtype value
func (v MachineCatalogWarningSubtype) Ptr() *MachineCatalogWarningSubtype {
	return &v
}

type NullableMachineCatalogWarningSubtype struct {
	value *MachineCatalogWarningSubtype
	isSet bool
}

func (v NullableMachineCatalogWarningSubtype) Get() *MachineCatalogWarningSubtype {
	return v.value
}

func (v *NullableMachineCatalogWarningSubtype) Set(val *MachineCatalogWarningSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineCatalogWarningSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineCatalogWarningSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineCatalogWarningSubtype(val *MachineCatalogWarningSubtype) *NullableMachineCatalogWarningSubtype {
	return &NullableMachineCatalogWarningSubtype{value: val, isSet: true}
}

func (v NullableMachineCatalogWarningSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineCatalogWarningSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

