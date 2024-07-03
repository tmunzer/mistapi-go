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

// checks if the CaptureMxedge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureMxedge{}

// CaptureMxedge Initiate a Wireless Packet Capture
type CaptureMxedge struct {
	// duration of the capture, in seconds
	Duration *int32 `json:"duration,omitempty"`
	Format *CaptureMxedgeFormat `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen *int32 `json:"max_pkt_len,omitempty"`
	Mxedges *map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets *int32 `json:"num_packets,omitempty"`
	Type CaptureMxedgeType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _CaptureMxedge CaptureMxedge

// NewCaptureMxedge instantiates a new CaptureMxedge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureMxedge(type_ CaptureMxedgeType) *CaptureMxedge {
	this := CaptureMxedge{}
	var duration int32 = 600
	this.Duration = &duration
	var format CaptureMxedgeFormat = CAPTUREMXEDGEFORMAT_STREAM
	this.Format = &format
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	var numPackets int32 = 1024
	this.NumPackets = &numPackets
	this.Type = type_
	return &this
}

// NewCaptureMxedgeWithDefaults instantiates a new CaptureMxedge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureMxedgeWithDefaults() *CaptureMxedge {
	this := CaptureMxedge{}
	var duration int32 = 600
	this.Duration = &duration
	var format CaptureMxedgeFormat = CAPTUREMXEDGEFORMAT_STREAM
	this.Format = &format
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	var numPackets int32 = 1024
	this.NumPackets = &numPackets
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CaptureMxedge) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CaptureMxedge) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *CaptureMxedge) SetDuration(v int32) {
	o.Duration = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CaptureMxedge) GetFormat() CaptureMxedgeFormat {
	if o == nil || IsNil(o.Format) {
		var ret CaptureMxedgeFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetFormatOk() (*CaptureMxedgeFormat, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CaptureMxedge) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given CaptureMxedgeFormat and assigns it to the Format field.
func (o *CaptureMxedge) SetFormat(v CaptureMxedgeFormat) {
	o.Format = &v
}

// GetMaxPktLen returns the MaxPktLen field value if set, zero value otherwise.
func (o *CaptureMxedge) GetMaxPktLen() int32 {
	if o == nil || IsNil(o.MaxPktLen) {
		var ret int32
		return ret
	}
	return *o.MaxPktLen
}

// GetMaxPktLenOk returns a tuple with the MaxPktLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetMaxPktLenOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPktLen) {
		return nil, false
	}
	return o.MaxPktLen, true
}

// HasMaxPktLen returns a boolean if a field has been set.
func (o *CaptureMxedge) HasMaxPktLen() bool {
	if o != nil && !IsNil(o.MaxPktLen) {
		return true
	}

	return false
}

// SetMaxPktLen gets a reference to the given int32 and assigns it to the MaxPktLen field.
func (o *CaptureMxedge) SetMaxPktLen(v int32) {
	o.MaxPktLen = &v
}

// GetMxedges returns the Mxedges field value if set, zero value otherwise.
func (o *CaptureMxedge) GetMxedges() map[string]CaptureMxedgeMxedges {
	if o == nil || IsNil(o.Mxedges) {
		var ret map[string]CaptureMxedgeMxedges
		return ret
	}
	return *o.Mxedges
}

// GetMxedgesOk returns a tuple with the Mxedges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetMxedgesOk() (*map[string]CaptureMxedgeMxedges, bool) {
	if o == nil || IsNil(o.Mxedges) {
		return nil, false
	}
	return o.Mxedges, true
}

// HasMxedges returns a boolean if a field has been set.
func (o *CaptureMxedge) HasMxedges() bool {
	if o != nil && !IsNil(o.Mxedges) {
		return true
	}

	return false
}

// SetMxedges gets a reference to the given map[string]CaptureMxedgeMxedges and assigns it to the Mxedges field.
func (o *CaptureMxedge) SetMxedges(v map[string]CaptureMxedgeMxedges) {
	o.Mxedges = &v
}

// GetNumPackets returns the NumPackets field value if set, zero value otherwise.
func (o *CaptureMxedge) GetNumPackets() int32 {
	if o == nil || IsNil(o.NumPackets) {
		var ret int32
		return ret
	}
	return *o.NumPackets
}

// GetNumPacketsOk returns a tuple with the NumPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetNumPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumPackets) {
		return nil, false
	}
	return o.NumPackets, true
}

// HasNumPackets returns a boolean if a field has been set.
func (o *CaptureMxedge) HasNumPackets() bool {
	if o != nil && !IsNil(o.NumPackets) {
		return true
	}

	return false
}

// SetNumPackets gets a reference to the given int32 and assigns it to the NumPackets field.
func (o *CaptureMxedge) SetNumPackets(v int32) {
	o.NumPackets = &v
}

// GetType returns the Type field value
func (o *CaptureMxedge) GetType() CaptureMxedgeType {
	if o == nil {
		var ret CaptureMxedgeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureMxedge) GetTypeOk() (*CaptureMxedgeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureMxedge) SetType(v CaptureMxedgeType) {
	o.Type = v
}

func (o CaptureMxedge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureMxedge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.MaxPktLen) {
		toSerialize["max_pkt_len"] = o.MaxPktLen
	}
	if !IsNil(o.Mxedges) {
		toSerialize["mxedges"] = o.Mxedges
	}
	if !IsNil(o.NumPackets) {
		toSerialize["num_packets"] = o.NumPackets
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaptureMxedge) UnmarshalJSON(data []byte) (err error) {
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

	varCaptureMxedge := _CaptureMxedge{}

	err = json.Unmarshal(data, &varCaptureMxedge)

	if err != nil {
		return err
	}

	*o = CaptureMxedge(varCaptureMxedge)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "format")
		delete(additionalProperties, "max_pkt_len")
		delete(additionalProperties, "mxedges")
		delete(additionalProperties, "num_packets")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaptureMxedge struct {
	value *CaptureMxedge
	isSet bool
}

func (v NullableCaptureMxedge) Get() *CaptureMxedge {
	return v.value
}

func (v *NullableCaptureMxedge) Set(val *CaptureMxedge) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureMxedge) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureMxedge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureMxedge(val *CaptureMxedge) *NullableCaptureMxedge {
	return &NullableCaptureMxedge{value: val, isSet: true}
}

func (v NullableCaptureMxedge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureMxedge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


