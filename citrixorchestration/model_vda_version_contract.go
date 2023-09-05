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

// checks if the VdaVersionContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VdaVersionContract{}

// VdaVersionContract struct for VdaVersionContract
type VdaVersionContract struct {
	// XA or XD, no other values.
	Product *string `json:"Product,omitempty"`
	// One of the values in XaVersions or XdVersions.
	Version *string `json:"Version,omitempty"`
	// MS or SS, no other values.
	Edition *string `json:"Edition,omitempty"`
	// Display name.
	DisplayName *string `json:"DisplayName,omitempty"`
}

// NewVdaVersionContract instantiates a new VdaVersionContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVdaVersionContract() *VdaVersionContract {
	this := VdaVersionContract{}
	return &this
}

// NewVdaVersionContractWithDefaults instantiates a new VdaVersionContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdaVersionContractWithDefaults() *VdaVersionContract {
	this := VdaVersionContract{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *VdaVersionContract) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdaVersionContract) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *VdaVersionContract) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *VdaVersionContract) SetProduct(v string) {
	o.Product = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VdaVersionContract) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdaVersionContract) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VdaVersionContract) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VdaVersionContract) SetVersion(v string) {
	o.Version = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *VdaVersionContract) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdaVersionContract) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *VdaVersionContract) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *VdaVersionContract) SetEdition(v string) {
	o.Edition = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *VdaVersionContract) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdaVersionContract) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *VdaVersionContract) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *VdaVersionContract) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o VdaVersionContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VdaVersionContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["Product"] = o.Product
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.Edition) {
		toSerialize["Edition"] = o.Edition
	}
	if !IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableVdaVersionContract struct {
	value *VdaVersionContract
	isSet bool
}

func (v NullableVdaVersionContract) Get() *VdaVersionContract {
	return v.value
}

func (v *NullableVdaVersionContract) Set(val *VdaVersionContract) {
	v.value = val
	v.isSet = true
}

func (v NullableVdaVersionContract) IsSet() bool {
	return v.isSet
}

func (v *NullableVdaVersionContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdaVersionContract(val *VdaVersionContract) *NullableVdaVersionContract {
	return &NullableVdaVersionContract{value: val, isSet: true}
}

func (v NullableVdaVersionContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdaVersionContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


