/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
	"fmt"
)

// CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType the model 'CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType'
type CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType string

// List of Citrix.CloudServices.Administrators.ApiModels.AdministratorNotificationType
const (
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORNOTIFICATIONTYPE_ERROR CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType = "Error"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORNOTIFICATIONTYPE_WARNING CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType = "Warning"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORNOTIFICATIONTYPE_INFORMATION CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType = "Information"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORNOTIFICATIONTYPE_CRITICAL CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType = "Critical"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORNOTIFICATIONTYPE_UNDEFINED CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType = "Undefined"
)

// All allowed values of CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType enum
var AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeEnumValues = []CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType{
	"Error",
	"Warning",
	"Information",
	"Critical",
	"Undefined",
}

func (v *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType(value)
	for _, existing := range AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType", value)
}

// NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeFromValue returns a pointer to a valid CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeFromValue(v string) (*CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType, error) {
	ev := CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType: valid values are %v", v, AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) IsValid() bool {
	for _, existing := range AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Citrix.CloudServices.Administrators.ApiModels.AdministratorNotificationType value
func (v CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) Ptr() *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType {
	return &v
}

type NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType struct {
	value *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) Get() *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) Set(val *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType(val *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType {
	return &NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

