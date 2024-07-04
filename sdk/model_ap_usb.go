/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the ApUsb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApUsb{}

// ApUsb USB AP settings Note: if native imagotag is enabled, BLE will be disabled automatically Note: legacy, new config moved to ESL Config.
type ApUsb struct {
	// only if `type`==`imagotag`
	Cacert NullableString `json:"cacert,omitempty"`
	// only if `type`==`imagotag` channel selection, not needed by default, required for manual channel override only
	Channel *int32 `json:"channel,omitempty"`
	// whether to enable any usb config
	Enabled *bool `json:"enabled,omitempty"`
	// only if `type`==`imagotag`
	Host *string `json:"host,omitempty"`
	// only if `type`==`imagotag`
	Port *int32 `json:"port,omitempty"`
	Type *ApUsbType `json:"type,omitempty"`
	// only if `type`==`imagotag`, whether to turn on SSL verification
	VerifyCert *bool `json:"verify_cert,omitempty"`
	// only if `type`==`solum` or `type`==`hanshow`
	VlanId *int32 `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApUsb ApUsb

// NewApUsb instantiates a new ApUsb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApUsb() *ApUsb {
	this := ApUsb{}
	var port int32 = 0
	this.Port = &port
	var vlanId int32 = 1
	this.VlanId = &vlanId
	return &this
}

// NewApUsbWithDefaults instantiates a new ApUsb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApUsbWithDefaults() *ApUsb {
	this := ApUsb{}
	var port int32 = 0
	this.Port = &port
	var vlanId int32 = 1
	this.VlanId = &vlanId
	return &this
}

// GetCacert returns the Cacert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApUsb) GetCacert() string {
	if o == nil || IsNil(o.Cacert.Get()) {
		var ret string
		return ret
	}
	return *o.Cacert.Get()
}

// GetCacertOk returns a tuple with the Cacert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApUsb) GetCacertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cacert.Get(), o.Cacert.IsSet()
}

// HasCacert returns a boolean if a field has been set.
func (o *ApUsb) HasCacert() bool {
	if o != nil && o.Cacert.IsSet() {
		return true
	}

	return false
}

// SetCacert gets a reference to the given NullableString and assigns it to the Cacert field.
func (o *ApUsb) SetCacert(v string) {
	o.Cacert.Set(&v)
}
// SetCacertNil sets the value for Cacert to be an explicit nil
func (o *ApUsb) SetCacertNil() {
	o.Cacert.Set(nil)
}

// UnsetCacert ensures that no value is present for Cacert, not even an explicit nil
func (o *ApUsb) UnsetCacert() {
	o.Cacert.Unset()
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ApUsb) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ApUsb) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ApUsb) SetChannel(v int32) {
	o.Channel = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApUsb) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApUsb) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApUsb) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ApUsb) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ApUsb) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ApUsb) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApUsb) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApUsb) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApUsb) SetPort(v int32) {
	o.Port = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApUsb) GetType() ApUsbType {
	if o == nil || IsNil(o.Type) {
		var ret ApUsbType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetTypeOk() (*ApUsbType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApUsb) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ApUsbType and assigns it to the Type field.
func (o *ApUsb) SetType(v ApUsbType) {
	o.Type = &v
}

// GetVerifyCert returns the VerifyCert field value if set, zero value otherwise.
func (o *ApUsb) GetVerifyCert() bool {
	if o == nil || IsNil(o.VerifyCert) {
		var ret bool
		return ret
	}
	return *o.VerifyCert
}

// GetVerifyCertOk returns a tuple with the VerifyCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetVerifyCertOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyCert) {
		return nil, false
	}
	return o.VerifyCert, true
}

// HasVerifyCert returns a boolean if a field has been set.
func (o *ApUsb) HasVerifyCert() bool {
	if o != nil && !IsNil(o.VerifyCert) {
		return true
	}

	return false
}

// SetVerifyCert gets a reference to the given bool and assigns it to the VerifyCert field.
func (o *ApUsb) SetVerifyCert(v bool) {
	o.VerifyCert = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *ApUsb) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApUsb) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *ApUsb) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *ApUsb) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o ApUsb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApUsb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cacert.IsSet() {
		toSerialize["cacert"] = o.Cacert.Get()
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VerifyCert) {
		toSerialize["verify_cert"] = o.VerifyCert
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApUsb) UnmarshalJSON(data []byte) (err error) {
	varApUsb := _ApUsb{}

	err = json.Unmarshal(data, &varApUsb)

	if err != nil {
		return err
	}

	*o = ApUsb(varApUsb)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cacert")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "type")
		delete(additionalProperties, "verify_cert")
		delete(additionalProperties, "vlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApUsb struct {
	value *ApUsb
	isSet bool
}

func (v NullableApUsb) Get() *ApUsb {
	return v.value
}

func (v *NullableApUsb) Set(val *ApUsb) {
	v.value = val
	v.isSet = true
}

func (v NullableApUsb) IsSet() bool {
	return v.isSet
}

func (v *NullableApUsb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApUsb(val *ApUsb) *NullableApUsb {
	return &NullableApUsb{value: val, isSet: true}
}

func (v NullableApUsb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApUsb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


