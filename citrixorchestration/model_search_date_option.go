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

// SearchDateOption Enum options for search time filter
type SearchDateOption string

// List of SearchDateOption
const (
	SEARCHDATEOPTION_LAST_MINUTE SearchDateOption = "LastMinute"
	SEARCHDATEOPTION_LAST5_MINUTES SearchDateOption = "Last5Minutes"
	SEARCHDATEOPTION_LAST30_MINUTES SearchDateOption = "Last30Minutes"
	SEARCHDATEOPTION_LAST_HOUR SearchDateOption = "LastHour"
	SEARCHDATEOPTION_LAST12_HOURS SearchDateOption = "Last12Hours"
	SEARCHDATEOPTION_LAST24_HOURS SearchDateOption = "Last24Hours"
	SEARCHDATEOPTION_TODAY SearchDateOption = "Today"
	SEARCHDATEOPTION_YESTERDAY SearchDateOption = "Yesterday"
	SEARCHDATEOPTION_LAST_MONTH SearchDateOption = "LastMonth"
	SEARCHDATEOPTION_LAST_SIX_MONTHS SearchDateOption = "LastSixMonths"
	SEARCHDATEOPTION_LAST_THREE_MONTHS SearchDateOption = "LastThreeMonths"
	SEARCHDATEOPTION_LAST28_DAYS SearchDateOption = "Last28Days"
	SEARCHDATEOPTION_LAST7_DAYS SearchDateOption = "Last7Days"
)

// All allowed values of SearchDateOption enum
var AllowedSearchDateOptionEnumValues = []SearchDateOption{
	"LastMinute",
	"Last5Minutes",
	"Last30Minutes",
	"LastHour",
	"Last12Hours",
	"Last24Hours",
	"Today",
	"Yesterday",
	"LastMonth",
	"LastSixMonths",
	"LastThreeMonths",
	"Last28Days",
	"Last7Days",
}

func (v *SearchDateOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchDateOption(value)
	for _, existing := range AllowedSearchDateOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchDateOption", value)
}

// NewSearchDateOptionFromValue returns a pointer to a valid SearchDateOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchDateOptionFromValue(v string) (*SearchDateOption, error) {
	ev := SearchDateOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchDateOption: valid values are %v", v, AllowedSearchDateOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchDateOption) IsValid() bool {
	for _, existing := range AllowedSearchDateOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchDateOption value
func (v SearchDateOption) Ptr() *SearchDateOption {
	return &v
}

type NullableSearchDateOption struct {
	value *SearchDateOption
	isSet bool
}

func (v NullableSearchDateOption) Get() *SearchDateOption {
	return v.value
}

func (v *NullableSearchDateOption) Set(val *SearchDateOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDateOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDateOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDateOption(val *SearchDateOption) *NullableSearchDateOption {
	return &NullableSearchDateOption{value: val, isSet: true}
}

func (v NullableSearchDateOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDateOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

