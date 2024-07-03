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

// checks if the RoutingPolicyTermMatchingVpnPathSla type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingPolicyTermMatchingVpnPathSla{}

// RoutingPolicyTermMatchingVpnPathSla struct for RoutingPolicyTermMatchingVpnPathSla
type RoutingPolicyTermMatchingVpnPathSla struct {
	MaxJitter NullableInt32 `json:"max_jitter,omitempty"`
	MaxLatency NullableInt32 `json:"max_latency,omitempty"`
	MaxLoss NullableInt32 `json:"max_loss,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingPolicyTermMatchingVpnPathSla RoutingPolicyTermMatchingVpnPathSla

// NewRoutingPolicyTermMatchingVpnPathSla instantiates a new RoutingPolicyTermMatchingVpnPathSla object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingPolicyTermMatchingVpnPathSla() *RoutingPolicyTermMatchingVpnPathSla {
	this := RoutingPolicyTermMatchingVpnPathSla{}
	return &this
}

// NewRoutingPolicyTermMatchingVpnPathSlaWithDefaults instantiates a new RoutingPolicyTermMatchingVpnPathSla object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingPolicyTermMatchingVpnPathSlaWithDefaults() *RoutingPolicyTermMatchingVpnPathSla {
	this := RoutingPolicyTermMatchingVpnPathSla{}
	return &this
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxJitter() int32 {
	if o == nil || IsNil(o.MaxJitter.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxJitter.Get()
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxJitterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxJitter.Get(), o.MaxJitter.IsSet()
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *RoutingPolicyTermMatchingVpnPathSla) HasMaxJitter() bool {
	if o != nil && o.MaxJitter.IsSet() {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given NullableInt32 and assigns it to the MaxJitter field.
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxJitter(v int32) {
	o.MaxJitter.Set(&v)
}
// SetMaxJitterNil sets the value for MaxJitter to be an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxJitterNil() {
	o.MaxJitter.Set(nil)
}

// UnsetMaxJitter ensures that no value is present for MaxJitter, not even an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) UnsetMaxJitter() {
	o.MaxJitter.Unset()
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxLatency() int32 {
	if o == nil || IsNil(o.MaxLatency.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLatency.Get()
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxLatencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLatency.Get(), o.MaxLatency.IsSet()
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *RoutingPolicyTermMatchingVpnPathSla) HasMaxLatency() bool {
	if o != nil && o.MaxLatency.IsSet() {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given NullableInt32 and assigns it to the MaxLatency field.
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxLatency(v int32) {
	o.MaxLatency.Set(&v)
}
// SetMaxLatencyNil sets the value for MaxLatency to be an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxLatencyNil() {
	o.MaxLatency.Set(nil)
}

// UnsetMaxLatency ensures that no value is present for MaxLatency, not even an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) UnsetMaxLatency() {
	o.MaxLatency.Unset()
}

// GetMaxLoss returns the MaxLoss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxLoss() int32 {
	if o == nil || IsNil(o.MaxLoss.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLoss.Get()
}

// GetMaxLossOk returns a tuple with the MaxLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingPolicyTermMatchingVpnPathSla) GetMaxLossOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLoss.Get(), o.MaxLoss.IsSet()
}

// HasMaxLoss returns a boolean if a field has been set.
func (o *RoutingPolicyTermMatchingVpnPathSla) HasMaxLoss() bool {
	if o != nil && o.MaxLoss.IsSet() {
		return true
	}

	return false
}

// SetMaxLoss gets a reference to the given NullableInt32 and assigns it to the MaxLoss field.
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxLoss(v int32) {
	o.MaxLoss.Set(&v)
}
// SetMaxLossNil sets the value for MaxLoss to be an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) SetMaxLossNil() {
	o.MaxLoss.Set(nil)
}

// UnsetMaxLoss ensures that no value is present for MaxLoss, not even an explicit nil
func (o *RoutingPolicyTermMatchingVpnPathSla) UnsetMaxLoss() {
	o.MaxLoss.Unset()
}

func (o RoutingPolicyTermMatchingVpnPathSla) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingPolicyTermMatchingVpnPathSla) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxJitter.IsSet() {
		toSerialize["max_jitter"] = o.MaxJitter.Get()
	}
	if o.MaxLatency.IsSet() {
		toSerialize["max_latency"] = o.MaxLatency.Get()
	}
	if o.MaxLoss.IsSet() {
		toSerialize["max_loss"] = o.MaxLoss.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingPolicyTermMatchingVpnPathSla) UnmarshalJSON(data []byte) (err error) {
	varRoutingPolicyTermMatchingVpnPathSla := _RoutingPolicyTermMatchingVpnPathSla{}

	err = json.Unmarshal(data, &varRoutingPolicyTermMatchingVpnPathSla)

	if err != nil {
		return err
	}

	*o = RoutingPolicyTermMatchingVpnPathSla(varRoutingPolicyTermMatchingVpnPathSla)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_jitter")
		delete(additionalProperties, "max_latency")
		delete(additionalProperties, "max_loss")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingPolicyTermMatchingVpnPathSla struct {
	value *RoutingPolicyTermMatchingVpnPathSla
	isSet bool
}

func (v NullableRoutingPolicyTermMatchingVpnPathSla) Get() *RoutingPolicyTermMatchingVpnPathSla {
	return v.value
}

func (v *NullableRoutingPolicyTermMatchingVpnPathSla) Set(val *RoutingPolicyTermMatchingVpnPathSla) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingPolicyTermMatchingVpnPathSla) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingPolicyTermMatchingVpnPathSla) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingPolicyTermMatchingVpnPathSla(val *RoutingPolicyTermMatchingVpnPathSla) *NullableRoutingPolicyTermMatchingVpnPathSla {
	return &NullableRoutingPolicyTermMatchingVpnPathSla{value: val, isSet: true}
}

func (v NullableRoutingPolicyTermMatchingVpnPathSla) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingPolicyTermMatchingVpnPathSla) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


