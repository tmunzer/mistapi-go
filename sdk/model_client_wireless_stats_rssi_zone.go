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

// checks if the ClientWirelessStatsRssiZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientWirelessStatsRssiZone{}

// ClientWirelessStatsRssiZone struct for ClientWirelessStatsRssiZone
type ClientWirelessStatsRssiZone struct {
	Id *string `json:"id,omitempty"`
	Since *int32 `json:"since,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientWirelessStatsRssiZone ClientWirelessStatsRssiZone

// NewClientWirelessStatsRssiZone instantiates a new ClientWirelessStatsRssiZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientWirelessStatsRssiZone() *ClientWirelessStatsRssiZone {
	this := ClientWirelessStatsRssiZone{}
	return &this
}

// NewClientWirelessStatsRssiZoneWithDefaults instantiates a new ClientWirelessStatsRssiZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWirelessStatsRssiZoneWithDefaults() *ClientWirelessStatsRssiZone {
	this := ClientWirelessStatsRssiZone{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientWirelessStatsRssiZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsRssiZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientWirelessStatsRssiZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientWirelessStatsRssiZone) SetId(v string) {
	o.Id = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *ClientWirelessStatsRssiZone) GetSince() int32 {
	if o == nil || IsNil(o.Since) {
		var ret int32
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsRssiZone) GetSinceOk() (*int32, bool) {
	if o == nil || IsNil(o.Since) {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *ClientWirelessStatsRssiZone) HasSince() bool {
	if o != nil && !IsNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given int32 and assigns it to the Since field.
func (o *ClientWirelessStatsRssiZone) SetSince(v int32) {
	o.Since = &v
}

func (o ClientWirelessStatsRssiZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientWirelessStatsRssiZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Since) {
		toSerialize["since"] = o.Since
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientWirelessStatsRssiZone) UnmarshalJSON(data []byte) (err error) {
	varClientWirelessStatsRssiZone := _ClientWirelessStatsRssiZone{}

	err = json.Unmarshal(data, &varClientWirelessStatsRssiZone)

	if err != nil {
		return err
	}

	*o = ClientWirelessStatsRssiZone(varClientWirelessStatsRssiZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "since")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientWirelessStatsRssiZone struct {
	value *ClientWirelessStatsRssiZone
	isSet bool
}

func (v NullableClientWirelessStatsRssiZone) Get() *ClientWirelessStatsRssiZone {
	return v.value
}

func (v *NullableClientWirelessStatsRssiZone) Set(val *ClientWirelessStatsRssiZone) {
	v.value = val
	v.isSet = true
}

func (v NullableClientWirelessStatsRssiZone) IsSet() bool {
	return v.isSet
}

func (v *NullableClientWirelessStatsRssiZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientWirelessStatsRssiZone(val *ClientWirelessStatsRssiZone) *NullableClientWirelessStatsRssiZone {
	return &NullableClientWirelessStatsRssiZone{value: val, isSet: true}
}

func (v NullableClientWirelessStatsRssiZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientWirelessStatsRssiZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


