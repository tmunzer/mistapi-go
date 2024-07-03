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

// checks if the WlanInjectDhcpOption82 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanInjectDhcpOption82{}

// WlanInjectDhcpOption82 struct for WlanInjectDhcpOption82
type WlanInjectDhcpOption82 struct {
	CircuitId *string `json:"circuit_id,omitempty"`
	// whether to inject option 82 when forwarding DHCP packets
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanInjectDhcpOption82 WlanInjectDhcpOption82

// NewWlanInjectDhcpOption82 instantiates a new WlanInjectDhcpOption82 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanInjectDhcpOption82() *WlanInjectDhcpOption82 {
	this := WlanInjectDhcpOption82{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWlanInjectDhcpOption82WithDefaults instantiates a new WlanInjectDhcpOption82 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanInjectDhcpOption82WithDefaults() *WlanInjectDhcpOption82 {
	this := WlanInjectDhcpOption82{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCircuitId returns the CircuitId field value if set, zero value otherwise.
func (o *WlanInjectDhcpOption82) GetCircuitId() string {
	if o == nil || IsNil(o.CircuitId) {
		var ret string
		return ret
	}
	return *o.CircuitId
}

// GetCircuitIdOk returns a tuple with the CircuitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanInjectDhcpOption82) GetCircuitIdOk() (*string, bool) {
	if o == nil || IsNil(o.CircuitId) {
		return nil, false
	}
	return o.CircuitId, true
}

// HasCircuitId returns a boolean if a field has been set.
func (o *WlanInjectDhcpOption82) HasCircuitId() bool {
	if o != nil && !IsNil(o.CircuitId) {
		return true
	}

	return false
}

// SetCircuitId gets a reference to the given string and assigns it to the CircuitId field.
func (o *WlanInjectDhcpOption82) SetCircuitId(v string) {
	o.CircuitId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanInjectDhcpOption82) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanInjectDhcpOption82) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanInjectDhcpOption82) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanInjectDhcpOption82) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o WlanInjectDhcpOption82) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanInjectDhcpOption82) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CircuitId) {
		toSerialize["circuit_id"] = o.CircuitId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanInjectDhcpOption82) UnmarshalJSON(data []byte) (err error) {
	varWlanInjectDhcpOption82 := _WlanInjectDhcpOption82{}

	err = json.Unmarshal(data, &varWlanInjectDhcpOption82)

	if err != nil {
		return err
	}

	*o = WlanInjectDhcpOption82(varWlanInjectDhcpOption82)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "circuit_id")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanInjectDhcpOption82 struct {
	value *WlanInjectDhcpOption82
	isSet bool
}

func (v NullableWlanInjectDhcpOption82) Get() *WlanInjectDhcpOption82 {
	return v.value
}

func (v *NullableWlanInjectDhcpOption82) Set(val *WlanInjectDhcpOption82) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanInjectDhcpOption82) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanInjectDhcpOption82) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanInjectDhcpOption82(val *WlanInjectDhcpOption82) *NullableWlanInjectDhcpOption82 {
	return &NullableWlanInjectDhcpOption82{value: val, isSet: true}
}

func (v NullableWlanInjectDhcpOption82) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanInjectDhcpOption82) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


