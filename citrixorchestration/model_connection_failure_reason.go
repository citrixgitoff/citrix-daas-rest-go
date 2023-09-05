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

// ConnectionFailureReason The reason why the last connection failure occurred.
type ConnectionFailureReason string

// List of ConnectionFailureReason
const (
	CONNECTIONFAILUREREASON_UNKNOWN ConnectionFailureReason = "Unknown"
	CONNECTIONFAILUREREASON_NONE ConnectionFailureReason = "None"
	CONNECTIONFAILUREREASON_SESSION_PREPARATION ConnectionFailureReason = "SessionPreparation"
	CONNECTIONFAILUREREASON_REGISTRATION_TIMEOUT ConnectionFailureReason = "RegistrationTimeout"
	CONNECTIONFAILUREREASON_CONNECTION_TIMEOUT ConnectionFailureReason = "ConnectionTimeout"
	CONNECTIONFAILUREREASON_LICENSING ConnectionFailureReason = "Licensing"
	CONNECTIONFAILUREREASON_TICKETING ConnectionFailureReason = "Ticketing"
	CONNECTIONFAILUREREASON_OTHER ConnectionFailureReason = "Other"
)

// All allowed values of ConnectionFailureReason enum
var AllowedConnectionFailureReasonEnumValues = []ConnectionFailureReason{
	"Unknown",
	"None",
	"SessionPreparation",
	"RegistrationTimeout",
	"ConnectionTimeout",
	"Licensing",
	"Ticketing",
	"Other",
}

func (v *ConnectionFailureReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionFailureReason(value)
	for _, existing := range AllowedConnectionFailureReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionFailureReason", value)
}

// NewConnectionFailureReasonFromValue returns a pointer to a valid ConnectionFailureReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionFailureReasonFromValue(v string) (*ConnectionFailureReason, error) {
	ev := ConnectionFailureReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionFailureReason: valid values are %v", v, AllowedConnectionFailureReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionFailureReason) IsValid() bool {
	for _, existing := range AllowedConnectionFailureReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionFailureReason value
func (v ConnectionFailureReason) Ptr() *ConnectionFailureReason {
	return &v
}

type NullableConnectionFailureReason struct {
	value *ConnectionFailureReason
	isSet bool
}

func (v NullableConnectionFailureReason) Get() *ConnectionFailureReason {
	return v.value
}

func (v *NullableConnectionFailureReason) Set(val *ConnectionFailureReason) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionFailureReason) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionFailureReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionFailureReason(val *ConnectionFailureReason) *NullableConnectionFailureReason {
	return &NullableConnectionFailureReason{value: val, isSet: true}
}

func (v NullableConnectionFailureReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionFailureReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

