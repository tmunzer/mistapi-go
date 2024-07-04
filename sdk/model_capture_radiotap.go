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
	"fmt"
)

// checks if the CaptureRadiotap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureRadiotap{}

// CaptureRadiotap Initiate a Radiotap Packet Capture
type CaptureRadiotap struct {
	ApMac *string `json:"ap_mac,omitempty"`
	Band *CaptureRadiotapBand `json:"band,omitempty"`
	ClientMac *string `json:"client_mac,omitempty"`
	// duration of the capture, in seconds
	Duration *int32 `json:"duration,omitempty"`
	Format *CaptureRadiotapFormat `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen *int32 `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets *int32 `json:"num_packets,omitempty"`
	Ssid *string `json:"ssid,omitempty"`
	// tcpdump expression specific to radiotap
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
	Type CaptureRadiotapType `json:"type"`
	// wlan id associated with the respective ssid.
	WlanId *string `json:"wlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaptureRadiotap CaptureRadiotap

// NewCaptureRadiotap instantiates a new CaptureRadiotap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureRadiotap(type_ CaptureRadiotapType) *CaptureRadiotap {
	this := CaptureRadiotap{}
	var band CaptureRadiotapBand = CAPTURERADIOTAPBAND__24
	this.Band = &band
	var duration int32 = 600
	this.Duration = &duration
	var format CaptureRadiotapFormat = CAPTURERADIOTAPFORMAT_PCAP
	this.Format = &format
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	var numPackets int32 = 1024
	this.NumPackets = &numPackets
	this.Type = type_
	return &this
}

// NewCaptureRadiotapWithDefaults instantiates a new CaptureRadiotap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureRadiotapWithDefaults() *CaptureRadiotap {
	this := CaptureRadiotap{}
	var band CaptureRadiotapBand = CAPTURERADIOTAPBAND__24
	this.Band = &band
	var duration int32 = 600
	this.Duration = &duration
	var format CaptureRadiotapFormat = CAPTURERADIOTAPFORMAT_PCAP
	this.Format = &format
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	var numPackets int32 = 1024
	this.NumPackets = &numPackets
	return &this
}

// GetApMac returns the ApMac field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetApMac() string {
	if o == nil || IsNil(o.ApMac) {
		var ret string
		return ret
	}
	return *o.ApMac
}

// GetApMacOk returns a tuple with the ApMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetApMacOk() (*string, bool) {
	if o == nil || IsNil(o.ApMac) {
		return nil, false
	}
	return o.ApMac, true
}

// HasApMac returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasApMac() bool {
	if o != nil && !IsNil(o.ApMac) {
		return true
	}

	return false
}

// SetApMac gets a reference to the given string and assigns it to the ApMac field.
func (o *CaptureRadiotap) SetApMac(v string) {
	o.ApMac = &v
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetBand() CaptureRadiotapBand {
	if o == nil || IsNil(o.Band) {
		var ret CaptureRadiotapBand
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetBandOk() (*CaptureRadiotapBand, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given CaptureRadiotapBand and assigns it to the Band field.
func (o *CaptureRadiotap) SetBand(v CaptureRadiotapBand) {
	o.Band = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetClientMac() string {
	if o == nil || IsNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetClientMacOk() (*string, bool) {
	if o == nil || IsNil(o.ClientMac) {
		return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasClientMac() bool {
	if o != nil && !IsNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *CaptureRadiotap) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *CaptureRadiotap) SetDuration(v int32) {
	o.Duration = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetFormat() CaptureRadiotapFormat {
	if o == nil || IsNil(o.Format) {
		var ret CaptureRadiotapFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetFormatOk() (*CaptureRadiotapFormat, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given CaptureRadiotapFormat and assigns it to the Format field.
func (o *CaptureRadiotap) SetFormat(v CaptureRadiotapFormat) {
	o.Format = &v
}

// GetMaxPktLen returns the MaxPktLen field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetMaxPktLen() int32 {
	if o == nil || IsNil(o.MaxPktLen) {
		var ret int32
		return ret
	}
	return *o.MaxPktLen
}

// GetMaxPktLenOk returns a tuple with the MaxPktLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetMaxPktLenOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPktLen) {
		return nil, false
	}
	return o.MaxPktLen, true
}

// HasMaxPktLen returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasMaxPktLen() bool {
	if o != nil && !IsNil(o.MaxPktLen) {
		return true
	}

	return false
}

// SetMaxPktLen gets a reference to the given int32 and assigns it to the MaxPktLen field.
func (o *CaptureRadiotap) SetMaxPktLen(v int32) {
	o.MaxPktLen = &v
}

// GetNumPackets returns the NumPackets field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetNumPackets() int32 {
	if o == nil || IsNil(o.NumPackets) {
		var ret int32
		return ret
	}
	return *o.NumPackets
}

// GetNumPacketsOk returns a tuple with the NumPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetNumPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumPackets) {
		return nil, false
	}
	return o.NumPackets, true
}

// HasNumPackets returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasNumPackets() bool {
	if o != nil && !IsNil(o.NumPackets) {
		return true
	}

	return false
}

// SetNumPackets gets a reference to the given int32 and assigns it to the NumPackets field.
func (o *CaptureRadiotap) SetNumPackets(v int32) {
	o.NumPackets = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *CaptureRadiotap) SetSsid(v string) {
	o.Ssid = &v
}

// GetTcpdumpExpression returns the TcpdumpExpression field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetTcpdumpExpression() string {
	if o == nil || IsNil(o.TcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.TcpdumpExpression
}

// GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.TcpdumpExpression) {
		return nil, false
	}
	return o.TcpdumpExpression, true
}

// HasTcpdumpExpression returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasTcpdumpExpression() bool {
	if o != nil && !IsNil(o.TcpdumpExpression) {
		return true
	}

	return false
}

// SetTcpdumpExpression gets a reference to the given string and assigns it to the TcpdumpExpression field.
func (o *CaptureRadiotap) SetTcpdumpExpression(v string) {
	o.TcpdumpExpression = &v
}

// GetType returns the Type field value
func (o *CaptureRadiotap) GetType() CaptureRadiotapType {
	if o == nil {
		var ret CaptureRadiotapType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetTypeOk() (*CaptureRadiotapType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureRadiotap) SetType(v CaptureRadiotapType) {
	o.Type = v
}

// GetWlanId returns the WlanId field value if set, zero value otherwise.
func (o *CaptureRadiotap) GetWlanId() string {
	if o == nil || IsNil(o.WlanId) {
		var ret string
		return ret
	}
	return *o.WlanId
}

// GetWlanIdOk returns a tuple with the WlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRadiotap) GetWlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.WlanId) {
		return nil, false
	}
	return o.WlanId, true
}

// HasWlanId returns a boolean if a field has been set.
func (o *CaptureRadiotap) HasWlanId() bool {
	if o != nil && !IsNil(o.WlanId) {
		return true
	}

	return false
}

// SetWlanId gets a reference to the given string and assigns it to the WlanId field.
func (o *CaptureRadiotap) SetWlanId(v string) {
	o.WlanId = &v
}

func (o CaptureRadiotap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureRadiotap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApMac) {
		toSerialize["ap_mac"] = o.ApMac
	}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.ClientMac) {
		toSerialize["client_mac"] = o.ClientMac
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.MaxPktLen) {
		toSerialize["max_pkt_len"] = o.MaxPktLen
	}
	if !IsNil(o.NumPackets) {
		toSerialize["num_packets"] = o.NumPackets
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.TcpdumpExpression) {
		toSerialize["tcpdump_expression"] = o.TcpdumpExpression
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.WlanId) {
		toSerialize["wlan_id"] = o.WlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaptureRadiotap) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCaptureRadiotap := _CaptureRadiotap{}

	err = json.Unmarshal(data, &varCaptureRadiotap)

	if err != nil {
		return err
	}

	*o = CaptureRadiotap(varCaptureRadiotap)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "band")
		delete(additionalProperties, "client_mac")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "format")
		delete(additionalProperties, "max_pkt_len")
		delete(additionalProperties, "num_packets")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "tcpdump_expression")
		delete(additionalProperties, "type")
		delete(additionalProperties, "wlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaptureRadiotap struct {
	value *CaptureRadiotap
	isSet bool
}

func (v NullableCaptureRadiotap) Get() *CaptureRadiotap {
	return v.value
}

func (v *NullableCaptureRadiotap) Set(val *CaptureRadiotap) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureRadiotap) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureRadiotap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureRadiotap(val *CaptureRadiotap) *NullableCaptureRadiotap {
	return &NullableCaptureRadiotap{value: val, isSet: true}
}

func (v NullableCaptureRadiotap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureRadiotap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


