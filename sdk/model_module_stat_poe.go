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

// checks if the ModuleStatPoe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatPoe{}

// ModuleStatPoe struct for ModuleStatPoe
type ModuleStatPoe struct {
	MaxPower *float32 `json:"max_power,omitempty"`
	PowerDraw *float32 `json:"power_draw,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStatPoe ModuleStatPoe

// NewModuleStatPoe instantiates a new ModuleStatPoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatPoe() *ModuleStatPoe {
	this := ModuleStatPoe{}
	return &this
}

// NewModuleStatPoeWithDefaults instantiates a new ModuleStatPoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatPoeWithDefaults() *ModuleStatPoe {
	this := ModuleStatPoe{}
	return &this
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *ModuleStatPoe) GetMaxPower() float32 {
	if o == nil || IsNil(o.MaxPower) {
		var ret float32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatPoe) GetMaxPowerOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxPower) {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *ModuleStatPoe) HasMaxPower() bool {
	if o != nil && !IsNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given float32 and assigns it to the MaxPower field.
func (o *ModuleStatPoe) SetMaxPower(v float32) {
	o.MaxPower = &v
}

// GetPowerDraw returns the PowerDraw field value if set, zero value otherwise.
func (o *ModuleStatPoe) GetPowerDraw() float32 {
	if o == nil || IsNil(o.PowerDraw) {
		var ret float32
		return ret
	}
	return *o.PowerDraw
}

// GetPowerDrawOk returns a tuple with the PowerDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatPoe) GetPowerDrawOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerDraw) {
		return nil, false
	}
	return o.PowerDraw, true
}

// HasPowerDraw returns a boolean if a field has been set.
func (o *ModuleStatPoe) HasPowerDraw() bool {
	if o != nil && !IsNil(o.PowerDraw) {
		return true
	}

	return false
}

// SetPowerDraw gets a reference to the given float32 and assigns it to the PowerDraw field.
func (o *ModuleStatPoe) SetPowerDraw(v float32) {
	o.PowerDraw = &v
}

func (o ModuleStatPoe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatPoe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxPower) {
		toSerialize["max_power"] = o.MaxPower
	}
	if !IsNil(o.PowerDraw) {
		toSerialize["power_draw"] = o.PowerDraw
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleStatPoe) UnmarshalJSON(data []byte) (err error) {
	varModuleStatPoe := _ModuleStatPoe{}

	err = json.Unmarshal(data, &varModuleStatPoe)

	if err != nil {
		return err
	}

	*o = ModuleStatPoe(varModuleStatPoe)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_power")
		delete(additionalProperties, "power_draw")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStatPoe struct {
	value *ModuleStatPoe
	isSet bool
}

func (v NullableModuleStatPoe) Get() *ModuleStatPoe {
	return v.value
}

func (v *NullableModuleStatPoe) Set(val *ModuleStatPoe) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatPoe) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatPoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatPoe(val *ModuleStatPoe) *NullableModuleStatPoe {
	return &NullableModuleStatPoe{value: val, isSet: true}
}

func (v NullableModuleStatPoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatPoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


