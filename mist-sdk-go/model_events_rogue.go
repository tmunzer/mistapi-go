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

// checks if the EventsRogue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsRogue{}

// EventsRogue rogue events
type EventsRogue struct {
	Ap string `json:"ap"`
	Bssid string `json:"bssid"`
	Channel int32 `json:"channel"`
	Rssi int32 `json:"rssi"`
	Ssid string `json:"ssid"`
	Timestamp float32 `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _EventsRogue EventsRogue

// NewEventsRogue instantiates a new EventsRogue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsRogue(ap string, bssid string, channel int32, rssi int32, ssid string, timestamp float32) *EventsRogue {
	this := EventsRogue{}
	this.Ap = ap
	this.Bssid = bssid
	this.Channel = channel
	this.Rssi = rssi
	this.Ssid = ssid
	this.Timestamp = timestamp
	return &this
}

// NewEventsRogueWithDefaults instantiates a new EventsRogue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsRogueWithDefaults() *EventsRogue {
	this := EventsRogue{}
	return &this
}

// GetAp returns the Ap field value
func (o *EventsRogue) GetAp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ap
}

// GetApOk returns a tuple with the Ap field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetApOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ap, true
}

// SetAp sets field value
func (o *EventsRogue) SetAp(v string) {
	o.Ap = v
}

// GetBssid returns the Bssid field value
func (o *EventsRogue) GetBssid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetBssidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bssid, true
}

// SetBssid sets field value
func (o *EventsRogue) SetBssid(v string) {
	o.Bssid = v
}

// GetChannel returns the Channel field value
func (o *EventsRogue) GetChannel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetChannelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *EventsRogue) SetChannel(v int32) {
	o.Channel = v
}

// GetRssi returns the Rssi field value
func (o *EventsRogue) GetRssi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetRssiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rssi, true
}

// SetRssi sets field value
func (o *EventsRogue) SetRssi(v int32) {
	o.Rssi = v
}

// GetSsid returns the Ssid field value
func (o *EventsRogue) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *EventsRogue) SetSsid(v string) {
	o.Ssid = v
}

// GetTimestamp returns the Timestamp field value
func (o *EventsRogue) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EventsRogue) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EventsRogue) SetTimestamp(v float32) {
	o.Timestamp = v
}

func (o EventsRogue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsRogue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ap"] = o.Ap
	toSerialize["bssid"] = o.Bssid
	toSerialize["channel"] = o.Channel
	toSerialize["rssi"] = o.Rssi
	toSerialize["ssid"] = o.Ssid
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventsRogue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ap",
		"bssid",
		"channel",
		"rssi",
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

	varEventsRogue := _EventsRogue{}

	err = json.Unmarshal(data, &varEventsRogue)

	if err != nil {
		return err
	}

	*o = EventsRogue(varEventsRogue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap")
		delete(additionalProperties, "bssid")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "rssi")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventsRogue struct {
	value *EventsRogue
	isSet bool
}

func (v NullableEventsRogue) Get() *EventsRogue {
	return v.value
}

func (v *NullableEventsRogue) Set(val *EventsRogue) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsRogue) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsRogue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsRogue(val *EventsRogue) *NullableEventsRogue {
	return &NullableEventsRogue{value: val, isSet: true}
}

func (v NullableEventsRogue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsRogue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


