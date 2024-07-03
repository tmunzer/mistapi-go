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

// checks if the ResponsePcapAp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePcapAp{}

// ResponsePcapAp struct for ResponsePcapAp
type ResponsePcapAp struct {
	Band *int32 `json:"band,omitempty"`
	Bandwidth *int32 `json:"bandwidth,omitempty"`
	Channel *int32 `json:"channel,omitempty"`
	TcpdumpExpresssion NullableString `json:"tcpdump_expresssion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponsePcapAp ResponsePcapAp

// NewResponsePcapAp instantiates a new ResponsePcapAp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePcapAp() *ResponsePcapAp {
	this := ResponsePcapAp{}
	return &this
}

// NewResponsePcapApWithDefaults instantiates a new ResponsePcapAp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePcapApWithDefaults() *ResponsePcapAp {
	this := ResponsePcapAp{}
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *ResponsePcapAp) GetBand() int32 {
	if o == nil || IsNil(o.Band) {
		var ret int32
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapAp) GetBandOk() (*int32, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *ResponsePcapAp) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given int32 and assigns it to the Band field.
func (o *ResponsePcapAp) SetBand(v int32) {
	o.Band = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *ResponsePcapAp) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapAp) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *ResponsePcapAp) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *ResponsePcapAp) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ResponsePcapAp) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapAp) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ResponsePcapAp) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ResponsePcapAp) SetChannel(v int32) {
	o.Channel = &v
}

// GetTcpdumpExpresssion returns the TcpdumpExpresssion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponsePcapAp) GetTcpdumpExpresssion() string {
	if o == nil || IsNil(o.TcpdumpExpresssion.Get()) {
		var ret string
		return ret
	}
	return *o.TcpdumpExpresssion.Get()
}

// GetTcpdumpExpresssionOk returns a tuple with the TcpdumpExpresssion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsePcapAp) GetTcpdumpExpresssionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TcpdumpExpresssion.Get(), o.TcpdumpExpresssion.IsSet()
}

// HasTcpdumpExpresssion returns a boolean if a field has been set.
func (o *ResponsePcapAp) HasTcpdumpExpresssion() bool {
	if o != nil && o.TcpdumpExpresssion.IsSet() {
		return true
	}

	return false
}

// SetTcpdumpExpresssion gets a reference to the given NullableString and assigns it to the TcpdumpExpresssion field.
func (o *ResponsePcapAp) SetTcpdumpExpresssion(v string) {
	o.TcpdumpExpresssion.Set(&v)
}
// SetTcpdumpExpresssionNil sets the value for TcpdumpExpresssion to be an explicit nil
func (o *ResponsePcapAp) SetTcpdumpExpresssionNil() {
	o.TcpdumpExpresssion.Set(nil)
}

// UnsetTcpdumpExpresssion ensures that no value is present for TcpdumpExpresssion, not even an explicit nil
func (o *ResponsePcapAp) UnsetTcpdumpExpresssion() {
	o.TcpdumpExpresssion.Unset()
}

func (o ResponsePcapAp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePcapAp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if o.TcpdumpExpresssion.IsSet() {
		toSerialize["tcpdump_expresssion"] = o.TcpdumpExpresssion.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponsePcapAp) UnmarshalJSON(data []byte) (err error) {
	varResponsePcapAp := _ResponsePcapAp{}

	err = json.Unmarshal(data, &varResponsePcapAp)

	if err != nil {
		return err
	}

	*o = ResponsePcapAp(varResponsePcapAp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "tcpdump_expresssion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponsePcapAp struct {
	value *ResponsePcapAp
	isSet bool
}

func (v NullableResponsePcapAp) Get() *ResponsePcapAp {
	return v.value
}

func (v *NullableResponsePcapAp) Set(val *ResponsePcapAp) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePcapAp) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePcapAp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePcapAp(val *ResponsePcapAp) *NullableResponsePcapAp {
	return &NullableResponsePcapAp{value: val, isSet: true}
}

func (v NullableResponsePcapAp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePcapAp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

