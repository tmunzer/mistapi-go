/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the SdkclientStatNetworkConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkclientStatNetworkConnection{}

// SdkclientStatNetworkConnection various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as
type SdkclientStatNetworkConnection struct {
	Mac string `json:"mac"`
	Rssi float32 `json:"rssi"`
	SignalLevel float32 `json:"signal_level"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _SdkclientStatNetworkConnection SdkclientStatNetworkConnection

// NewSdkclientStatNetworkConnection instantiates a new SdkclientStatNetworkConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkclientStatNetworkConnection(mac string, rssi float32, signalLevel float32, type_ string) *SdkclientStatNetworkConnection {
	this := SdkclientStatNetworkConnection{}
	this.Mac = mac
	this.Rssi = rssi
	this.SignalLevel = signalLevel
	this.Type = type_
	return &this
}

// NewSdkclientStatNetworkConnectionWithDefaults instantiates a new SdkclientStatNetworkConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkclientStatNetworkConnectionWithDefaults() *SdkclientStatNetworkConnection {
	this := SdkclientStatNetworkConnection{}
	return &this
}

// GetMac returns the Mac field value
func (o *SdkclientStatNetworkConnection) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *SdkclientStatNetworkConnection) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *SdkclientStatNetworkConnection) SetMac(v string) {
	o.Mac = v
}

// GetRssi returns the Rssi field value
func (o *SdkclientStatNetworkConnection) GetRssi() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value
// and a boolean to check if the value has been set.
func (o *SdkclientStatNetworkConnection) GetRssiOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rssi, true
}

// SetRssi sets field value
func (o *SdkclientStatNetworkConnection) SetRssi(v float32) {
	o.Rssi = v
}

// GetSignalLevel returns the SignalLevel field value
func (o *SdkclientStatNetworkConnection) GetSignalLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SignalLevel
}

// GetSignalLevelOk returns a tuple with the SignalLevel field value
// and a boolean to check if the value has been set.
func (o *SdkclientStatNetworkConnection) GetSignalLevelOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalLevel, true
}

// SetSignalLevel sets field value
func (o *SdkclientStatNetworkConnection) SetSignalLevel(v float32) {
	o.SignalLevel = v
}

// GetType returns the Type field value
func (o *SdkclientStatNetworkConnection) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SdkclientStatNetworkConnection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SdkclientStatNetworkConnection) SetType(v string) {
	o.Type = v
}

func (o SdkclientStatNetworkConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkclientStatNetworkConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	toSerialize["rssi"] = o.Rssi
	toSerialize["signal_level"] = o.SignalLevel
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdkclientStatNetworkConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mac",
		"rssi",
		"signal_level",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSdkclientStatNetworkConnection := _SdkclientStatNetworkConnection{}

	err = json.Unmarshal(data, &varSdkclientStatNetworkConnection)

	if err != nil {
		return err
	}

	*o = SdkclientStatNetworkConnection(varSdkclientStatNetworkConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "rssi")
		delete(additionalProperties, "signal_level")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkclientStatNetworkConnection struct {
	value *SdkclientStatNetworkConnection
	isSet bool
}

func (v NullableSdkclientStatNetworkConnection) Get() *SdkclientStatNetworkConnection {
	return v.value
}

func (v *NullableSdkclientStatNetworkConnection) Set(val *SdkclientStatNetworkConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkclientStatNetworkConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkclientStatNetworkConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkclientStatNetworkConnection(val *SdkclientStatNetworkConnection) *NullableSdkclientStatNetworkConnection {
	return &NullableSdkclientStatNetworkConnection{value: val, isSet: true}
}

func (v NullableSdkclientStatNetworkConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkclientStatNetworkConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


