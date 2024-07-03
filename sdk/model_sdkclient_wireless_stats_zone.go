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

// checks if the SdkclientWirelessStatsZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkclientWirelessStatsZone{}

// SdkclientWirelessStatsZone struct for SdkclientWirelessStatsZone
type SdkclientWirelessStatsZone struct {
	Id string `json:"id"`
	Since float32 `json:"since"`
	AdditionalProperties map[string]interface{}
}

type _SdkclientWirelessStatsZone SdkclientWirelessStatsZone

// NewSdkclientWirelessStatsZone instantiates a new SdkclientWirelessStatsZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkclientWirelessStatsZone(id string, since float32) *SdkclientWirelessStatsZone {
	this := SdkclientWirelessStatsZone{}
	this.Id = id
	this.Since = since
	return &this
}

// NewSdkclientWirelessStatsZoneWithDefaults instantiates a new SdkclientWirelessStatsZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkclientWirelessStatsZoneWithDefaults() *SdkclientWirelessStatsZone {
	this := SdkclientWirelessStatsZone{}
	return &this
}

// GetId returns the Id field value
func (o *SdkclientWirelessStatsZone) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SdkclientWirelessStatsZone) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SdkclientWirelessStatsZone) SetId(v string) {
	o.Id = v
}

// GetSince returns the Since field value
func (o *SdkclientWirelessStatsZone) GetSince() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Since
}

// GetSinceOk returns a tuple with the Since field value
// and a boolean to check if the value has been set.
func (o *SdkclientWirelessStatsZone) GetSinceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Since, true
}

// SetSince sets field value
func (o *SdkclientWirelessStatsZone) SetSince(v float32) {
	o.Since = v
}

func (o SdkclientWirelessStatsZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkclientWirelessStatsZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["since"] = o.Since

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdkclientWirelessStatsZone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"since",
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

	varSdkclientWirelessStatsZone := _SdkclientWirelessStatsZone{}

	err = json.Unmarshal(data, &varSdkclientWirelessStatsZone)

	if err != nil {
		return err
	}

	*o = SdkclientWirelessStatsZone(varSdkclientWirelessStatsZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "since")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkclientWirelessStatsZone struct {
	value *SdkclientWirelessStatsZone
	isSet bool
}

func (v NullableSdkclientWirelessStatsZone) Get() *SdkclientWirelessStatsZone {
	return v.value
}

func (v *NullableSdkclientWirelessStatsZone) Set(val *SdkclientWirelessStatsZone) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkclientWirelessStatsZone) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkclientWirelessStatsZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkclientWirelessStatsZone(val *SdkclientWirelessStatsZone) *NullableSdkclientWirelessStatsZone {
	return &NullableSdkclientWirelessStatsZone{value: val, isSet: true}
}

func (v NullableSdkclientWirelessStatsZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkclientWirelessStatsZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


