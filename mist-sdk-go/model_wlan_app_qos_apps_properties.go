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
)

// checks if the WlanAppQosAppsProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanAppQosAppsProperties{}

// WlanAppQosAppsProperties struct for WlanAppQosAppsProperties
type WlanAppQosAppsProperties struct {
	Dscp *int32 `json:"dscp,omitempty"`
	// subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
	DstSubnet *string `json:"dst_subnet,omitempty"`
	// subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
	SrcSubnet *string `json:"src_subnet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanAppQosAppsProperties WlanAppQosAppsProperties

// NewWlanAppQosAppsProperties instantiates a new WlanAppQosAppsProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanAppQosAppsProperties() *WlanAppQosAppsProperties {
	this := WlanAppQosAppsProperties{}
	return &this
}

// NewWlanAppQosAppsPropertiesWithDefaults instantiates a new WlanAppQosAppsProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanAppQosAppsPropertiesWithDefaults() *WlanAppQosAppsProperties {
	this := WlanAppQosAppsProperties{}
	return &this
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *WlanAppQosAppsProperties) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQosAppsProperties) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *WlanAppQosAppsProperties) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *WlanAppQosAppsProperties) SetDscp(v int32) {
	o.Dscp = &v
}

// GetDstSubnet returns the DstSubnet field value if set, zero value otherwise.
func (o *WlanAppQosAppsProperties) GetDstSubnet() string {
	if o == nil || IsNil(o.DstSubnet) {
		var ret string
		return ret
	}
	return *o.DstSubnet
}

// GetDstSubnetOk returns a tuple with the DstSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQosAppsProperties) GetDstSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.DstSubnet) {
		return nil, false
	}
	return o.DstSubnet, true
}

// HasDstSubnet returns a boolean if a field has been set.
func (o *WlanAppQosAppsProperties) HasDstSubnet() bool {
	if o != nil && !IsNil(o.DstSubnet) {
		return true
	}

	return false
}

// SetDstSubnet gets a reference to the given string and assigns it to the DstSubnet field.
func (o *WlanAppQosAppsProperties) SetDstSubnet(v string) {
	o.DstSubnet = &v
}

// GetSrcSubnet returns the SrcSubnet field value if set, zero value otherwise.
func (o *WlanAppQosAppsProperties) GetSrcSubnet() string {
	if o == nil || IsNil(o.SrcSubnet) {
		var ret string
		return ret
	}
	return *o.SrcSubnet
}

// GetSrcSubnetOk returns a tuple with the SrcSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQosAppsProperties) GetSrcSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.SrcSubnet) {
		return nil, false
	}
	return o.SrcSubnet, true
}

// HasSrcSubnet returns a boolean if a field has been set.
func (o *WlanAppQosAppsProperties) HasSrcSubnet() bool {
	if o != nil && !IsNil(o.SrcSubnet) {
		return true
	}

	return false
}

// SetSrcSubnet gets a reference to the given string and assigns it to the SrcSubnet field.
func (o *WlanAppQosAppsProperties) SetSrcSubnet(v string) {
	o.SrcSubnet = &v
}

func (o WlanAppQosAppsProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanAppQosAppsProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !IsNil(o.DstSubnet) {
		toSerialize["dst_subnet"] = o.DstSubnet
	}
	if !IsNil(o.SrcSubnet) {
		toSerialize["src_subnet"] = o.SrcSubnet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanAppQosAppsProperties) UnmarshalJSON(data []byte) (err error) {
	varWlanAppQosAppsProperties := _WlanAppQosAppsProperties{}

	err = json.Unmarshal(data, &varWlanAppQosAppsProperties)

	if err != nil {
		return err
	}

	*o = WlanAppQosAppsProperties(varWlanAppQosAppsProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dscp")
		delete(additionalProperties, "dst_subnet")
		delete(additionalProperties, "src_subnet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanAppQosAppsProperties struct {
	value *WlanAppQosAppsProperties
	isSet bool
}

func (v NullableWlanAppQosAppsProperties) Get() *WlanAppQosAppsProperties {
	return v.value
}

func (v *NullableWlanAppQosAppsProperties) Set(val *WlanAppQosAppsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAppQosAppsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAppQosAppsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAppQosAppsProperties(val *WlanAppQosAppsProperties) *NullableWlanAppQosAppsProperties {
	return &NullableWlanAppQosAppsProperties{value: val, isSet: true}
}

func (v NullableWlanAppQosAppsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAppQosAppsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


