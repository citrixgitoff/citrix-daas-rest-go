/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// ConfigLogGetOperationsSearchDateOptionParameter - struct for ConfigLogGetOperationsSearchDateOptionParameter
type ConfigLogGetOperationsSearchDateOptionParameter struct {
	SearchDateOption *SearchDateOption
}

// SearchDateOptionAsConfigLogGetOperationsSearchDateOptionParameter is a convenience function that returns SearchDateOption wrapped in ConfigLogGetOperationsSearchDateOptionParameter
func SearchDateOptionAsConfigLogGetOperationsSearchDateOptionParameter(v *SearchDateOption) ConfigLogGetOperationsSearchDateOptionParameter {
	return ConfigLogGetOperationsSearchDateOptionParameter{
		SearchDateOption: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfigLogGetOperationsSearchDateOptionParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SearchDateOption
	err = newStrictDecoder(data).Decode(&dst.SearchDateOption)
	if err == nil {
		jsonSearchDateOption, _ := json.Marshal(dst.SearchDateOption)
		if string(jsonSearchDateOption) == "{}" { // empty struct
			dst.SearchDateOption = nil
		} else {
			match++
		}
	} else {
		dst.SearchDateOption = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SearchDateOption = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ConfigLogGetOperationsSearchDateOptionParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConfigLogGetOperationsSearchDateOptionParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfigLogGetOperationsSearchDateOptionParameter) MarshalJSON() ([]byte, error) {
	if src.SearchDateOption != nil {
		return json.Marshal(&src.SearchDateOption)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConfigLogGetOperationsSearchDateOptionParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SearchDateOption != nil {
		return obj.SearchDateOption
	}

	// all schemas are nil
	return nil
}

type NullableConfigLogGetOperationsSearchDateOptionParameter struct {
	value *ConfigLogGetOperationsSearchDateOptionParameter
	isSet bool
}

func (v NullableConfigLogGetOperationsSearchDateOptionParameter) Get() *ConfigLogGetOperationsSearchDateOptionParameter {
	return v.value
}

func (v *NullableConfigLogGetOperationsSearchDateOptionParameter) Set(val *ConfigLogGetOperationsSearchDateOptionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigLogGetOperationsSearchDateOptionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigLogGetOperationsSearchDateOptionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigLogGetOperationsSearchDateOptionParameter(val *ConfigLogGetOperationsSearchDateOptionParameter) *NullableConfigLogGetOperationsSearchDateOptionParameter {
	return &NullableConfigLogGetOperationsSearchDateOptionParameter{value: val, isSet: true}
}

func (v NullableConfigLogGetOperationsSearchDateOptionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigLogGetOperationsSearchDateOptionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


