/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the SessionClientResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionClientResponseModel{}

// SessionClientResponseModel Response object for details about the client connected to a session.
type SessionClientResponseModel struct {
	// Unique identifier for the client device that has most recently been associated with the session.
	DeviceId NullableString `json:"DeviceId,omitempty"`
	// Unique identifier for the client hardware that has been most recently associated with the session.
	HardwareId NullableString `json:"HardwareId,omitempty"`
	// The IP address of the client connected to the session.
	IPAddress NullableString `json:"IPAddress,omitempty"`
	// The host name of the client connected to the session.
	Name NullableString `json:"Name,omitempty"`
	// The name of client platform, as indicated by client product ID.
	Platform NullableString `json:"Platform,omitempty"`
	// The product ID of the client connected to the session.
	ProductId NullableInt32 `json:"ProductId,omitempty"`
	// The IP address of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected.
	ReceiverIPAddress NullableString `json:"ReceiverIPAddress,omitempty"`
	// The name of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected.
	ReceiverName NullableString `json:"ReceiverName,omitempty"`
	// The version of the Citrix Receiver running on the client connected to the session.
	Version NullableString `json:"Version,omitempty"`
}

// NewSessionClientResponseModel instantiates a new SessionClientResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionClientResponseModel() *SessionClientResponseModel {
	this := SessionClientResponseModel{}
	return &this
}

// NewSessionClientResponseModelWithDefaults instantiates a new SessionClientResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionClientResponseModelWithDefaults() *SessionClientResponseModel {
	this := SessionClientResponseModel{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *SessionClientResponseModel) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *SessionClientResponseModel) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *SessionClientResponseModel) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetHardwareId returns the HardwareId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetHardwareId() string {
	if o == nil || IsNil(o.HardwareId.Get()) {
		var ret string
		return ret
	}
	return *o.HardwareId.Get()
}

// GetHardwareIdOk returns a tuple with the HardwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetHardwareIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardwareId.Get(), o.HardwareId.IsSet()
}

// HasHardwareId returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasHardwareId() bool {
	if o != nil && o.HardwareId.IsSet() {
		return true
	}

	return false
}

// SetHardwareId gets a reference to the given NullableString and assigns it to the HardwareId field.
func (o *SessionClientResponseModel) SetHardwareId(v string) {
	o.HardwareId.Set(&v)
}
// SetHardwareIdNil sets the value for HardwareId to be an explicit nil
func (o *SessionClientResponseModel) SetHardwareIdNil() {
	o.HardwareId.Set(nil)
}

// UnsetHardwareId ensures that no value is present for HardwareId, not even an explicit nil
func (o *SessionClientResponseModel) UnsetHardwareId() {
	o.HardwareId.Unset()
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetIPAddress() string {
	if o == nil || IsNil(o.IPAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IPAddress.Get()
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IPAddress.Get(), o.IPAddress.IsSet()
}

// HasIPAddress returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasIPAddress() bool {
	if o != nil && o.IPAddress.IsSet() {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given NullableString and assigns it to the IPAddress field.
func (o *SessionClientResponseModel) SetIPAddress(v string) {
	o.IPAddress.Set(&v)
}
// SetIPAddressNil sets the value for IPAddress to be an explicit nil
func (o *SessionClientResponseModel) SetIPAddressNil() {
	o.IPAddress.Set(nil)
}

// UnsetIPAddress ensures that no value is present for IPAddress, not even an explicit nil
func (o *SessionClientResponseModel) UnsetIPAddress() {
	o.IPAddress.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SessionClientResponseModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SessionClientResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SessionClientResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *SessionClientResponseModel) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *SessionClientResponseModel) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *SessionClientResponseModel) UnsetPlatform() {
	o.Platform.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableInt32 and assigns it to the ProductId field.
func (o *SessionClientResponseModel) SetProductId(v int32) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *SessionClientResponseModel) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *SessionClientResponseModel) UnsetProductId() {
	o.ProductId.Unset()
}

// GetReceiverIPAddress returns the ReceiverIPAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetReceiverIPAddress() string {
	if o == nil || IsNil(o.ReceiverIPAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverIPAddress.Get()
}

// GetReceiverIPAddressOk returns a tuple with the ReceiverIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetReceiverIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverIPAddress.Get(), o.ReceiverIPAddress.IsSet()
}

// HasReceiverIPAddress returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasReceiverIPAddress() bool {
	if o != nil && o.ReceiverIPAddress.IsSet() {
		return true
	}

	return false
}

// SetReceiverIPAddress gets a reference to the given NullableString and assigns it to the ReceiverIPAddress field.
func (o *SessionClientResponseModel) SetReceiverIPAddress(v string) {
	o.ReceiverIPAddress.Set(&v)
}
// SetReceiverIPAddressNil sets the value for ReceiverIPAddress to be an explicit nil
func (o *SessionClientResponseModel) SetReceiverIPAddressNil() {
	o.ReceiverIPAddress.Set(nil)
}

// UnsetReceiverIPAddress ensures that no value is present for ReceiverIPAddress, not even an explicit nil
func (o *SessionClientResponseModel) UnsetReceiverIPAddress() {
	o.ReceiverIPAddress.Unset()
}

// GetReceiverName returns the ReceiverName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetReceiverName() string {
	if o == nil || IsNil(o.ReceiverName.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverName.Get()
}

// GetReceiverNameOk returns a tuple with the ReceiverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetReceiverNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverName.Get(), o.ReceiverName.IsSet()
}

// HasReceiverName returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasReceiverName() bool {
	if o != nil && o.ReceiverName.IsSet() {
		return true
	}

	return false
}

// SetReceiverName gets a reference to the given NullableString and assigns it to the ReceiverName field.
func (o *SessionClientResponseModel) SetReceiverName(v string) {
	o.ReceiverName.Set(&v)
}
// SetReceiverNameNil sets the value for ReceiverName to be an explicit nil
func (o *SessionClientResponseModel) SetReceiverNameNil() {
	o.ReceiverName.Set(nil)
}

// UnsetReceiverName ensures that no value is present for ReceiverName, not even an explicit nil
func (o *SessionClientResponseModel) UnsetReceiverName() {
	o.ReceiverName.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionClientResponseModel) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionClientResponseModel) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *SessionClientResponseModel) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *SessionClientResponseModel) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *SessionClientResponseModel) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *SessionClientResponseModel) UnsetVersion() {
	o.Version.Unset()
}

func (o SessionClientResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionClientResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId.IsSet() {
		toSerialize["DeviceId"] = o.DeviceId.Get()
	}
	if o.HardwareId.IsSet() {
		toSerialize["HardwareId"] = o.HardwareId.Get()
	}
	if o.IPAddress.IsSet() {
		toSerialize["IPAddress"] = o.IPAddress.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["Platform"] = o.Platform.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["ProductId"] = o.ProductId.Get()
	}
	if o.ReceiverIPAddress.IsSet() {
		toSerialize["ReceiverIPAddress"] = o.ReceiverIPAddress.Get()
	}
	if o.ReceiverName.IsSet() {
		toSerialize["ReceiverName"] = o.ReceiverName.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableSessionClientResponseModel struct {
	value *SessionClientResponseModel
	isSet bool
}

func (v NullableSessionClientResponseModel) Get() *SessionClientResponseModel {
	return v.value
}

func (v *NullableSessionClientResponseModel) Set(val *SessionClientResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionClientResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionClientResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionClientResponseModel(val *SessionClientResponseModel) *NullableSessionClientResponseModel {
	return &NullableSessionClientResponseModel{value: val, isSet: true}
}

func (v NullableSessionClientResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionClientResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


