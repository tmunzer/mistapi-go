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

// checks if the EventFastroam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFastroam{}

// EventFastroam struct for EventFastroam
type EventFastroam struct {
	ApMac string `json:"ap_mac"`
	ClientMac string `json:"client_mac"`
	Fromap string `json:"fromap"`
	Latency float32 `json:"latency"`
	Ssid string `json:"ssid"`
	Subtype *string `json:"subtype,omitempty"`
	// timestamp of the event in nsec
	Timestamp float32 `json:"timestamp"`
	Type *EventFastroamType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventFastroam EventFastroam

// NewEventFastroam instantiates a new EventFastroam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFastroam(apMac string, clientMac string, fromap string, latency float32, ssid string, timestamp float32) *EventFastroam {
	this := EventFastroam{}
	this.ApMac = apMac
	this.ClientMac = clientMac
	this.Fromap = fromap
	this.Latency = latency
	this.Ssid = ssid
	this.Timestamp = timestamp
	return &this
}

// NewEventFastroamWithDefaults instantiates a new EventFastroam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFastroamWithDefaults() *EventFastroam {
	this := EventFastroam{}
	return &this
}

// GetApMac returns the ApMac field value
func (o *EventFastroam) GetApMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApMac
}

// GetApMacOk returns a tuple with the ApMac field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetApMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApMac, true
}

// SetApMac sets field value
func (o *EventFastroam) SetApMac(v string) {
	o.ApMac = v
}

// GetClientMac returns the ClientMac field value
func (o *EventFastroam) GetClientMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetClientMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientMac, true
}

// SetClientMac sets field value
func (o *EventFastroam) SetClientMac(v string) {
	o.ClientMac = v
}

// GetFromap returns the Fromap field value
func (o *EventFastroam) GetFromap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fromap
}

// GetFromapOk returns a tuple with the Fromap field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetFromapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fromap, true
}

// SetFromap sets field value
func (o *EventFastroam) SetFromap(v string) {
	o.Fromap = v
}

// GetLatency returns the Latency field value
func (o *EventFastroam) GetLatency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetLatencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latency, true
}

// SetLatency sets field value
func (o *EventFastroam) SetLatency(v float32) {
	o.Latency = v
}

// GetSsid returns the Ssid field value
func (o *EventFastroam) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *EventFastroam) SetSsid(v string) {
	o.Ssid = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *EventFastroam) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *EventFastroam) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *EventFastroam) SetSubtype(v string) {
	o.Subtype = &v
}

// GetTimestamp returns the Timestamp field value
func (o *EventFastroam) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EventFastroam) SetTimestamp(v float32) {
	o.Timestamp = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventFastroam) GetType() EventFastroamType {
	if o == nil || IsNil(o.Type) {
		var ret EventFastroamType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFastroam) GetTypeOk() (*EventFastroamType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventFastroam) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EventFastroamType and assigns it to the Type field.
func (o *EventFastroam) SetType(v EventFastroamType) {
	o.Type = &v
}

func (o EventFastroam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFastroam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ap_mac"] = o.ApMac
	toSerialize["client_mac"] = o.ClientMac
	toSerialize["fromap"] = o.Fromap
	toSerialize["latency"] = o.Latency
	toSerialize["ssid"] = o.Ssid
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventFastroam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ap_mac",
		"client_mac",
		"fromap",
		"latency",
		"ssid",
		"timestamp",
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

	varEventFastroam := _EventFastroam{}

	err = json.Unmarshal(data, &varEventFastroam)

	if err != nil {
		return err
	}

	*o = EventFastroam(varEventFastroam)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "client_mac")
		delete(additionalProperties, "fromap")
		delete(additionalProperties, "latency")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventFastroam struct {
	value *EventFastroam
	isSet bool
}

func (v NullableEventFastroam) Get() *EventFastroam {
	return v.value
}

func (v *NullableEventFastroam) Set(val *EventFastroam) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFastroam) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFastroam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFastroam(val *EventFastroam) *NullableEventFastroam {
	return &NullableEventFastroam{value: val, isSet: true}
}

func (v NullableEventFastroam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFastroam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


