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

// checks if the MxedgeStatsIpStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeStatsIpStat{}

// MxedgeStatsIpStat OOBM IP stats
type MxedgeStatsIpStat struct {
	Ip *string `json:"ip,omitempty"`
	// Property key is the interface name. IPs for each net interface
	Ips *map[string]string `json:"ips,omitempty"`
	// Property key is the interface name. MAC for each net interface
	Macs *map[string]string `json:"macs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeStatsIpStat MxedgeStatsIpStat

// NewMxedgeStatsIpStat instantiates a new MxedgeStatsIpStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeStatsIpStat() *MxedgeStatsIpStat {
	this := MxedgeStatsIpStat{}
	return &this
}

// NewMxedgeStatsIpStatWithDefaults instantiates a new MxedgeStatsIpStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeStatsIpStatWithDefaults() *MxedgeStatsIpStat {
	this := MxedgeStatsIpStat{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *MxedgeStatsIpStat) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsIpStat) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *MxedgeStatsIpStat) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *MxedgeStatsIpStat) SetIp(v string) {
	o.Ip = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *MxedgeStatsIpStat) GetIps() map[string]string {
	if o == nil || IsNil(o.Ips) {
		var ret map[string]string
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsIpStat) GetIpsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *MxedgeStatsIpStat) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given map[string]string and assigns it to the Ips field.
func (o *MxedgeStatsIpStat) SetIps(v map[string]string) {
	o.Ips = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *MxedgeStatsIpStat) GetMacs() map[string]string {
	if o == nil || IsNil(o.Macs) {
		var ret map[string]string
		return ret
	}
	return *o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsIpStat) GetMacsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *MxedgeStatsIpStat) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given map[string]string and assigns it to the Macs field.
func (o *MxedgeStatsIpStat) SetMacs(v map[string]string) {
	o.Macs = &v
}

func (o MxedgeStatsIpStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeStatsIpStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeStatsIpStat) UnmarshalJSON(data []byte) (err error) {
	varMxedgeStatsIpStat := _MxedgeStatsIpStat{}

	err = json.Unmarshal(data, &varMxedgeStatsIpStat)

	if err != nil {
		return err
	}

	*o = MxedgeStatsIpStat(varMxedgeStatsIpStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "ips")
		delete(additionalProperties, "macs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeStatsIpStat struct {
	value *MxedgeStatsIpStat
	isSet bool
}

func (v NullableMxedgeStatsIpStat) Get() *MxedgeStatsIpStat {
	return v.value
}

func (v *NullableMxedgeStatsIpStat) Set(val *MxedgeStatsIpStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeStatsIpStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeStatsIpStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeStatsIpStat(val *MxedgeStatsIpStat) *NullableMxedgeStatsIpStat {
	return &NullableMxedgeStatsIpStat{value: val, isSet: true}
}

func (v NullableMxedgeStatsIpStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeStatsIpStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


