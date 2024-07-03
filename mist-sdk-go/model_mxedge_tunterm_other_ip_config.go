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

// checks if the MxedgeTuntermOtherIpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeTuntermOtherIpConfig{}

// MxedgeTuntermOtherIpConfig struct for MxedgeTuntermOtherIpConfig
type MxedgeTuntermOtherIpConfig struct {
	Ip string `json:"ip"`
	Netmask string `json:"netmask"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeTuntermOtherIpConfig MxedgeTuntermOtherIpConfig

// NewMxedgeTuntermOtherIpConfig instantiates a new MxedgeTuntermOtherIpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeTuntermOtherIpConfig(ip string, netmask string) *MxedgeTuntermOtherIpConfig {
	this := MxedgeTuntermOtherIpConfig{}
	this.Ip = ip
	this.Netmask = netmask
	return &this
}

// NewMxedgeTuntermOtherIpConfigWithDefaults instantiates a new MxedgeTuntermOtherIpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeTuntermOtherIpConfigWithDefaults() *MxedgeTuntermOtherIpConfig {
	this := MxedgeTuntermOtherIpConfig{}
	return &this
}

// GetIp returns the Ip field value
func (o *MxedgeTuntermOtherIpConfig) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermOtherIpConfig) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *MxedgeTuntermOtherIpConfig) SetIp(v string) {
	o.Ip = v
}

// GetNetmask returns the Netmask field value
func (o *MxedgeTuntermOtherIpConfig) GetNetmask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermOtherIpConfig) GetNetmaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netmask, true
}

// SetNetmask sets field value
func (o *MxedgeTuntermOtherIpConfig) SetNetmask(v string) {
	o.Netmask = v
}

func (o MxedgeTuntermOtherIpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeTuntermOtherIpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["netmask"] = o.Netmask

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeTuntermOtherIpConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"netmask",
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

	varMxedgeTuntermOtherIpConfig := _MxedgeTuntermOtherIpConfig{}

	err = json.Unmarshal(data, &varMxedgeTuntermOtherIpConfig)

	if err != nil {
		return err
	}

	*o = MxedgeTuntermOtherIpConfig(varMxedgeTuntermOtherIpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "netmask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeTuntermOtherIpConfig struct {
	value *MxedgeTuntermOtherIpConfig
	isSet bool
}

func (v NullableMxedgeTuntermOtherIpConfig) Get() *MxedgeTuntermOtherIpConfig {
	return v.value
}

func (v *NullableMxedgeTuntermOtherIpConfig) Set(val *MxedgeTuntermOtherIpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermOtherIpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermOtherIpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermOtherIpConfig(val *MxedgeTuntermOtherIpConfig) *NullableMxedgeTuntermOtherIpConfig {
	return &NullableMxedgeTuntermOtherIpConfig{value: val, isSet: true}
}

func (v NullableMxedgeTuntermOtherIpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermOtherIpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


