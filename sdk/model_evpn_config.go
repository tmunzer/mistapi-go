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

// checks if the EvpnConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvpnConfig{}

// EvpnConfig EVPN Junos settings
type EvpnConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	Role *EvpnConfigRole `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EvpnConfig EvpnConfig

// NewEvpnConfig instantiates a new EvpnConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvpnConfig() *EvpnConfig {
	this := EvpnConfig{}
	return &this
}

// NewEvpnConfigWithDefaults instantiates a new EvpnConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvpnConfigWithDefaults() *EvpnConfig {
	this := EvpnConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EvpnConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EvpnConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EvpnConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *EvpnConfig) GetRole() EvpnConfigRole {
	if o == nil || IsNil(o.Role) {
		var ret EvpnConfigRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnConfig) GetRoleOk() (*EvpnConfigRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *EvpnConfig) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given EvpnConfigRole and assigns it to the Role field.
func (o *EvpnConfig) SetRole(v EvpnConfigRole) {
	o.Role = &v
}

func (o EvpnConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvpnConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EvpnConfig) UnmarshalJSON(data []byte) (err error) {
	varEvpnConfig := _EvpnConfig{}

	err = json.Unmarshal(data, &varEvpnConfig)

	if err != nil {
		return err
	}

	*o = EvpnConfig(varEvpnConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvpnConfig struct {
	value *EvpnConfig
	isSet bool
}

func (v NullableEvpnConfig) Get() *EvpnConfig {
	return v.value
}

func (v *NullableEvpnConfig) Set(val *EvpnConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEvpnConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEvpnConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvpnConfig(val *EvpnConfig) *NullableEvpnConfig {
	return &NullableEvpnConfig{value: val, isSet: true}
}

func (v NullableEvpnConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvpnConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


