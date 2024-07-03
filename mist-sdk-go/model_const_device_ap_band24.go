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

// checks if the ConstDeviceApBand24 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceApBand24{}

// ConstDeviceApBand24 struct for ConstDeviceApBand24
type ConstDeviceApBand24 struct {
	Band5ChannelsOp *string `json:"band5_channels_op,omitempty"`
	MaxClients *int32 `json:"max_clients,omitempty"`
	MaxPower *int32 `json:"max_power,omitempty"`
	MinPower *int32 `json:"min_power,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceApBand24 ConstDeviceApBand24

// NewConstDeviceApBand24 instantiates a new ConstDeviceApBand24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceApBand24() *ConstDeviceApBand24 {
	this := ConstDeviceApBand24{}
	return &this
}

// NewConstDeviceApBand24WithDefaults instantiates a new ConstDeviceApBand24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceApBand24WithDefaults() *ConstDeviceApBand24 {
	this := ConstDeviceApBand24{}
	return &this
}

// GetBand5ChannelsOp returns the Band5ChannelsOp field value if set, zero value otherwise.
func (o *ConstDeviceApBand24) GetBand5ChannelsOp() string {
	if o == nil || IsNil(o.Band5ChannelsOp) {
		var ret string
		return ret
	}
	return *o.Band5ChannelsOp
}

// GetBand5ChannelsOpOk returns a tuple with the Band5ChannelsOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand24) GetBand5ChannelsOpOk() (*string, bool) {
	if o == nil || IsNil(o.Band5ChannelsOp) {
		return nil, false
	}
	return o.Band5ChannelsOp, true
}

// HasBand5ChannelsOp returns a boolean if a field has been set.
func (o *ConstDeviceApBand24) HasBand5ChannelsOp() bool {
	if o != nil && !IsNil(o.Band5ChannelsOp) {
		return true
	}

	return false
}

// SetBand5ChannelsOp gets a reference to the given string and assigns it to the Band5ChannelsOp field.
func (o *ConstDeviceApBand24) SetBand5ChannelsOp(v string) {
	o.Band5ChannelsOp = &v
}

// GetMaxClients returns the MaxClients field value if set, zero value otherwise.
func (o *ConstDeviceApBand24) GetMaxClients() int32 {
	if o == nil || IsNil(o.MaxClients) {
		var ret int32
		return ret
	}
	return *o.MaxClients
}

// GetMaxClientsOk returns a tuple with the MaxClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand24) GetMaxClientsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxClients) {
		return nil, false
	}
	return o.MaxClients, true
}

// HasMaxClients returns a boolean if a field has been set.
func (o *ConstDeviceApBand24) HasMaxClients() bool {
	if o != nil && !IsNil(o.MaxClients) {
		return true
	}

	return false
}

// SetMaxClients gets a reference to the given int32 and assigns it to the MaxClients field.
func (o *ConstDeviceApBand24) SetMaxClients(v int32) {
	o.MaxClients = &v
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *ConstDeviceApBand24) GetMaxPower() int32 {
	if o == nil || IsNil(o.MaxPower) {
		var ret int32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand24) GetMaxPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPower) {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *ConstDeviceApBand24) HasMaxPower() bool {
	if o != nil && !IsNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int32 and assigns it to the MaxPower field.
func (o *ConstDeviceApBand24) SetMaxPower(v int32) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *ConstDeviceApBand24) GetMinPower() int32 {
	if o == nil || IsNil(o.MinPower) {
		var ret int32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand24) GetMinPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.MinPower) {
		return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *ConstDeviceApBand24) HasMinPower() bool {
	if o != nil && !IsNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int32 and assigns it to the MinPower field.
func (o *ConstDeviceApBand24) SetMinPower(v int32) {
	o.MinPower = &v
}

func (o ConstDeviceApBand24) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceApBand24) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band5ChannelsOp) {
		toSerialize["band5_channels_op"] = o.Band5ChannelsOp
	}
	if !IsNil(o.MaxClients) {
		toSerialize["max_clients"] = o.MaxClients
	}
	if !IsNil(o.MaxPower) {
		toSerialize["max_power"] = o.MaxPower
	}
	if !IsNil(o.MinPower) {
		toSerialize["min_power"] = o.MinPower
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceApBand24) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceApBand24 := _ConstDeviceApBand24{}

	err = json.Unmarshal(data, &varConstDeviceApBand24)

	if err != nil {
		return err
	}

	*o = ConstDeviceApBand24(varConstDeviceApBand24)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band5_channels_op")
		delete(additionalProperties, "max_clients")
		delete(additionalProperties, "max_power")
		delete(additionalProperties, "min_power")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceApBand24 struct {
	value *ConstDeviceApBand24
	isSet bool
}

func (v NullableConstDeviceApBand24) Get() *ConstDeviceApBand24 {
	return v.value
}

func (v *NullableConstDeviceApBand24) Set(val *ConstDeviceApBand24) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceApBand24) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceApBand24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceApBand24(val *ConstDeviceApBand24) *NullableConstDeviceApBand24 {
	return &NullableConstDeviceApBand24{value: val, isSet: true}
}

func (v NullableConstDeviceApBand24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceApBand24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

