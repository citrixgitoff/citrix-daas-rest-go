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

// checks if the HypervisorSecurityGroupRuleResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorSecurityGroupRuleResponseModel{}

// HypervisorSecurityGroupRuleResponseModel Security group rule.
type HypervisorSecurityGroupRuleResponseModel struct {
	// Start of the port number range, required for any protocol that uses ports. In AWS this can also be an ICMP port number, where -1 specifies all ICMP types.
	FromPort NullableFloat32 `json:"FromPort,omitempty"`
	// IDs of source or destination security groups. Empty when CIDR ranges are specified.
	GroupIds []string `json:"GroupIds,omitempty"`
	// Source or destination CIDR address ranges. Empty if group IDs are specified.
	IPRanges []string `json:"IPRanges,omitempty"`
	// IP protocol name or number. In AWS this is the IP protocol name or number, where -1 specifies all protocols. Protocol numbers: http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xml
	Protocol string `json:"Protocol"`
	// End of the port number range, required for any protocol that uses ports. In AWS this can also be an ICMP port number, where -1 specfies all ICMP types.
	ToPort NullableFloat32 `json:"ToPort,omitempty"`
}

// NewHypervisorSecurityGroupRuleResponseModel instantiates a new HypervisorSecurityGroupRuleResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorSecurityGroupRuleResponseModel(protocol string) *HypervisorSecurityGroupRuleResponseModel {
	this := HypervisorSecurityGroupRuleResponseModel{}
	this.Protocol = protocol
	return &this
}

// NewHypervisorSecurityGroupRuleResponseModelWithDefaults instantiates a new HypervisorSecurityGroupRuleResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorSecurityGroupRuleResponseModelWithDefaults() *HypervisorSecurityGroupRuleResponseModel {
	this := HypervisorSecurityGroupRuleResponseModel{}
	return &this
}

// GetFromPort returns the FromPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorSecurityGroupRuleResponseModel) GetFromPort() float32 {
	if o == nil || IsNil(o.FromPort.Get()) {
		var ret float32
		return ret
	}
	return *o.FromPort.Get()
}

// GetFromPortOk returns a tuple with the FromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorSecurityGroupRuleResponseModel) GetFromPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromPort.Get(), o.FromPort.IsSet()
}

// HasFromPort returns a boolean if a field has been set.
func (o *HypervisorSecurityGroupRuleResponseModel) HasFromPort() bool {
	if o != nil && o.FromPort.IsSet() {
		return true
	}

	return false
}

// SetFromPort gets a reference to the given NullableFloat32 and assigns it to the FromPort field.
func (o *HypervisorSecurityGroupRuleResponseModel) SetFromPort(v float32) {
	o.FromPort.Set(&v)
}
// SetFromPortNil sets the value for FromPort to be an explicit nil
func (o *HypervisorSecurityGroupRuleResponseModel) SetFromPortNil() {
	o.FromPort.Set(nil)
}

// UnsetFromPort ensures that no value is present for FromPort, not even an explicit nil
func (o *HypervisorSecurityGroupRuleResponseModel) UnsetFromPort() {
	o.FromPort.Unset()
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorSecurityGroupRuleResponseModel) GetGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorSecurityGroupRuleResponseModel) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *HypervisorSecurityGroupRuleResponseModel) HasGroupIds() bool {
	if o != nil && IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *HypervisorSecurityGroupRuleResponseModel) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetIPRanges returns the IPRanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorSecurityGroupRuleResponseModel) GetIPRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IPRanges
}

// GetIPRangesOk returns a tuple with the IPRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorSecurityGroupRuleResponseModel) GetIPRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IPRanges) {
		return nil, false
	}
	return o.IPRanges, true
}

// HasIPRanges returns a boolean if a field has been set.
func (o *HypervisorSecurityGroupRuleResponseModel) HasIPRanges() bool {
	if o != nil && IsNil(o.IPRanges) {
		return true
	}

	return false
}

// SetIPRanges gets a reference to the given []string and assigns it to the IPRanges field.
func (o *HypervisorSecurityGroupRuleResponseModel) SetIPRanges(v []string) {
	o.IPRanges = v
}

// GetProtocol returns the Protocol field value
func (o *HypervisorSecurityGroupRuleResponseModel) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *HypervisorSecurityGroupRuleResponseModel) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *HypervisorSecurityGroupRuleResponseModel) SetProtocol(v string) {
	o.Protocol = v
}

// GetToPort returns the ToPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorSecurityGroupRuleResponseModel) GetToPort() float32 {
	if o == nil || IsNil(o.ToPort.Get()) {
		var ret float32
		return ret
	}
	return *o.ToPort.Get()
}

// GetToPortOk returns a tuple with the ToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorSecurityGroupRuleResponseModel) GetToPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToPort.Get(), o.ToPort.IsSet()
}

// HasToPort returns a boolean if a field has been set.
func (o *HypervisorSecurityGroupRuleResponseModel) HasToPort() bool {
	if o != nil && o.ToPort.IsSet() {
		return true
	}

	return false
}

// SetToPort gets a reference to the given NullableFloat32 and assigns it to the ToPort field.
func (o *HypervisorSecurityGroupRuleResponseModel) SetToPort(v float32) {
	o.ToPort.Set(&v)
}
// SetToPortNil sets the value for ToPort to be an explicit nil
func (o *HypervisorSecurityGroupRuleResponseModel) SetToPortNil() {
	o.ToPort.Set(nil)
}

// UnsetToPort ensures that no value is present for ToPort, not even an explicit nil
func (o *HypervisorSecurityGroupRuleResponseModel) UnsetToPort() {
	o.ToPort.Unset()
}

func (o HypervisorSecurityGroupRuleResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorSecurityGroupRuleResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FromPort.IsSet() {
		toSerialize["FromPort"] = o.FromPort.Get()
	}
	if o.GroupIds != nil {
		toSerialize["GroupIds"] = o.GroupIds
	}
	if o.IPRanges != nil {
		toSerialize["IPRanges"] = o.IPRanges
	}
	toSerialize["Protocol"] = o.Protocol
	if o.ToPort.IsSet() {
		toSerialize["ToPort"] = o.ToPort.Get()
	}
	return toSerialize, nil
}

type NullableHypervisorSecurityGroupRuleResponseModel struct {
	value *HypervisorSecurityGroupRuleResponseModel
	isSet bool
}

func (v NullableHypervisorSecurityGroupRuleResponseModel) Get() *HypervisorSecurityGroupRuleResponseModel {
	return v.value
}

func (v *NullableHypervisorSecurityGroupRuleResponseModel) Set(val *HypervisorSecurityGroupRuleResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorSecurityGroupRuleResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorSecurityGroupRuleResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorSecurityGroupRuleResponseModel(val *HypervisorSecurityGroupRuleResponseModel) *NullableHypervisorSecurityGroupRuleResponseModel {
	return &NullableHypervisorSecurityGroupRuleResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorSecurityGroupRuleResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorSecurityGroupRuleResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


