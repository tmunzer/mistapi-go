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

// checks if the ApPortConfigDynamicVlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApPortConfigDynamicVlan{}

// ApPortConfigDynamicVlan optional dynamic vlan
type ApPortConfigDynamicVlan struct {
	DefaultVlanId *int32 `json:"default_vlan_id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Type *string `json:"type,omitempty"`
	Vlans *map[string]string `json:"vlans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApPortConfigDynamicVlan ApPortConfigDynamicVlan

// NewApPortConfigDynamicVlan instantiates a new ApPortConfigDynamicVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApPortConfigDynamicVlan() *ApPortConfigDynamicVlan {
	this := ApPortConfigDynamicVlan{}
	return &this
}

// NewApPortConfigDynamicVlanWithDefaults instantiates a new ApPortConfigDynamicVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApPortConfigDynamicVlanWithDefaults() *ApPortConfigDynamicVlan {
	this := ApPortConfigDynamicVlan{}
	return &this
}

// GetDefaultVlanId returns the DefaultVlanId field value if set, zero value otherwise.
func (o *ApPortConfigDynamicVlan) GetDefaultVlanId() int32 {
	if o == nil || IsNil(o.DefaultVlanId) {
		var ret int32
		return ret
	}
	return *o.DefaultVlanId
}

// GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPortConfigDynamicVlan) GetDefaultVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultVlanId) {
		return nil, false
	}
	return o.DefaultVlanId, true
}

// HasDefaultVlanId returns a boolean if a field has been set.
func (o *ApPortConfigDynamicVlan) HasDefaultVlanId() bool {
	if o != nil && !IsNil(o.DefaultVlanId) {
		return true
	}

	return false
}

// SetDefaultVlanId gets a reference to the given int32 and assigns it to the DefaultVlanId field.
func (o *ApPortConfigDynamicVlan) SetDefaultVlanId(v int32) {
	o.DefaultVlanId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApPortConfigDynamicVlan) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPortConfigDynamicVlan) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApPortConfigDynamicVlan) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApPortConfigDynamicVlan) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApPortConfigDynamicVlan) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPortConfigDynamicVlan) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApPortConfigDynamicVlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApPortConfigDynamicVlan) SetType(v string) {
	o.Type = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *ApPortConfigDynamicVlan) GetVlans() map[string]string {
	if o == nil || IsNil(o.Vlans) {
		var ret map[string]string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPortConfigDynamicVlan) GetVlansOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Vlans) {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *ApPortConfigDynamicVlan) HasVlans() bool {
	if o != nil && !IsNil(o.Vlans) {
		return true
	}

	return false
}

// SetVlans gets a reference to the given map[string]string and assigns it to the Vlans field.
func (o *ApPortConfigDynamicVlan) SetVlans(v map[string]string) {
	o.Vlans = &v
}

func (o ApPortConfigDynamicVlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApPortConfigDynamicVlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultVlanId) {
		toSerialize["default_vlan_id"] = o.DefaultVlanId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vlans) {
		toSerialize["vlans"] = o.Vlans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApPortConfigDynamicVlan) UnmarshalJSON(data []byte) (err error) {
	varApPortConfigDynamicVlan := _ApPortConfigDynamicVlan{}

	err = json.Unmarshal(data, &varApPortConfigDynamicVlan)

	if err != nil {
		return err
	}

	*o = ApPortConfigDynamicVlan(varApPortConfigDynamicVlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default_vlan_id")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vlans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApPortConfigDynamicVlan struct {
	value *ApPortConfigDynamicVlan
	isSet bool
}

func (v NullableApPortConfigDynamicVlan) Get() *ApPortConfigDynamicVlan {
	return v.value
}

func (v *NullableApPortConfigDynamicVlan) Set(val *ApPortConfigDynamicVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableApPortConfigDynamicVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableApPortConfigDynamicVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApPortConfigDynamicVlan(val *ApPortConfigDynamicVlan) *NullableApPortConfigDynamicVlan {
	return &NullableApPortConfigDynamicVlan{value: val, isSet: true}
}

func (v NullableApPortConfigDynamicVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApPortConfigDynamicVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

