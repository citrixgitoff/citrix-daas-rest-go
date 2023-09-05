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

// checks if the IdentityValidateUserRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityValidateUserRequestModel{}

// IdentityValidateUserRequestModel Validate user credentials and if successful, retrieve the properties of the user.
type IdentityValidateUserRequestModel struct {
	// Forest in which the user account resides.
	Forest *string `json:"Forest,omitempty"`
	// Domain in which the user account resides.
	Domain *string `json:"Domain,omitempty"`
	LogOnType *IdentityLogonType `json:"LogOnType,omitempty"`
	// Username to validate.  Required.
	UserName string `json:"UserName"`
	// Password to validate.  Required. Must be specified in the format indicated by PasswordFormat.
	Password string `json:"Password"`
	PasswordFormat *IdentityPasswordFormat `json:"PasswordFormat,omitempty"`
}

// NewIdentityValidateUserRequestModel instantiates a new IdentityValidateUserRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityValidateUserRequestModel(userName string, password string) *IdentityValidateUserRequestModel {
	this := IdentityValidateUserRequestModel{}
	this.UserName = userName
	this.Password = password
	return &this
}

// NewIdentityValidateUserRequestModelWithDefaults instantiates a new IdentityValidateUserRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityValidateUserRequestModelWithDefaults() *IdentityValidateUserRequestModel {
	this := IdentityValidateUserRequestModel{}
	return &this
}

// GetForest returns the Forest field value if set, zero value otherwise.
func (o *IdentityValidateUserRequestModel) GetForest() string {
	if o == nil || IsNil(o.Forest) {
		var ret string
		return ret
	}
	return *o.Forest
}

// GetForestOk returns a tuple with the Forest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetForestOk() (*string, bool) {
	if o == nil || IsNil(o.Forest) {
		return nil, false
	}
	return o.Forest, true
}

// HasForest returns a boolean if a field has been set.
func (o *IdentityValidateUserRequestModel) HasForest() bool {
	if o != nil && !IsNil(o.Forest) {
		return true
	}

	return false
}

// SetForest gets a reference to the given string and assigns it to the Forest field.
func (o *IdentityValidateUserRequestModel) SetForest(v string) {
	o.Forest = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *IdentityValidateUserRequestModel) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *IdentityValidateUserRequestModel) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *IdentityValidateUserRequestModel) SetDomain(v string) {
	o.Domain = &v
}

// GetLogOnType returns the LogOnType field value if set, zero value otherwise.
func (o *IdentityValidateUserRequestModel) GetLogOnType() IdentityLogonType {
	if o == nil || IsNil(o.LogOnType) {
		var ret IdentityLogonType
		return ret
	}
	return *o.LogOnType
}

// GetLogOnTypeOk returns a tuple with the LogOnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetLogOnTypeOk() (*IdentityLogonType, bool) {
	if o == nil || IsNil(o.LogOnType) {
		return nil, false
	}
	return o.LogOnType, true
}

// HasLogOnType returns a boolean if a field has been set.
func (o *IdentityValidateUserRequestModel) HasLogOnType() bool {
	if o != nil && !IsNil(o.LogOnType) {
		return true
	}

	return false
}

// SetLogOnType gets a reference to the given IdentityLogonType and assigns it to the LogOnType field.
func (o *IdentityValidateUserRequestModel) SetLogOnType(v IdentityLogonType) {
	o.LogOnType = &v
}

// GetUserName returns the UserName field value
func (o *IdentityValidateUserRequestModel) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *IdentityValidateUserRequestModel) SetUserName(v string) {
	o.UserName = v
}

// GetPassword returns the Password field value
func (o *IdentityValidateUserRequestModel) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *IdentityValidateUserRequestModel) SetPassword(v string) {
	o.Password = v
}

// GetPasswordFormat returns the PasswordFormat field value if set, zero value otherwise.
func (o *IdentityValidateUserRequestModel) GetPasswordFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.PasswordFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.PasswordFormat
}

// GetPasswordFormatOk returns a tuple with the PasswordFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityValidateUserRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.PasswordFormat) {
		return nil, false
	}
	return o.PasswordFormat, true
}

// HasPasswordFormat returns a boolean if a field has been set.
func (o *IdentityValidateUserRequestModel) HasPasswordFormat() bool {
	if o != nil && !IsNil(o.PasswordFormat) {
		return true
	}

	return false
}

// SetPasswordFormat gets a reference to the given IdentityPasswordFormat and assigns it to the PasswordFormat field.
func (o *IdentityValidateUserRequestModel) SetPasswordFormat(v IdentityPasswordFormat) {
	o.PasswordFormat = &v
}

func (o IdentityValidateUserRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityValidateUserRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Forest) {
		toSerialize["Forest"] = o.Forest
	}
	if !IsNil(o.Domain) {
		toSerialize["Domain"] = o.Domain
	}
	if !IsNil(o.LogOnType) {
		toSerialize["LogOnType"] = o.LogOnType
	}
	toSerialize["UserName"] = o.UserName
	toSerialize["Password"] = o.Password
	if !IsNil(o.PasswordFormat) {
		toSerialize["PasswordFormat"] = o.PasswordFormat
	}
	return toSerialize, nil
}

type NullableIdentityValidateUserRequestModel struct {
	value *IdentityValidateUserRequestModel
	isSet bool
}

func (v NullableIdentityValidateUserRequestModel) Get() *IdentityValidateUserRequestModel {
	return v.value
}

func (v *NullableIdentityValidateUserRequestModel) Set(val *IdentityValidateUserRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityValidateUserRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityValidateUserRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityValidateUserRequestModel(val *IdentityValidateUserRequestModel) *NullableIdentityValidateUserRequestModel {
	return &NullableIdentityValidateUserRequestModel{value: val, isSet: true}
}

func (v NullableIdentityValidateUserRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityValidateUserRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


