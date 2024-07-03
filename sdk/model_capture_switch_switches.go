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

// checks if the CaptureSwitchSwitches type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureSwitchSwitches{}

// CaptureSwitchSwitches struct for CaptureSwitchSwitches
type CaptureSwitchSwitches struct {
	// Property key is the interface name
	Ports *map[string]map[string]interface{} `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaptureSwitchSwitches CaptureSwitchSwitches

// NewCaptureSwitchSwitches instantiates a new CaptureSwitchSwitches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureSwitchSwitches() *CaptureSwitchSwitches {
	this := CaptureSwitchSwitches{}
	return &this
}

// NewCaptureSwitchSwitchesWithDefaults instantiates a new CaptureSwitchSwitches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureSwitchSwitchesWithDefaults() *CaptureSwitchSwitches {
	this := CaptureSwitchSwitches{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *CaptureSwitchSwitches) GetPorts() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Ports) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureSwitchSwitches) GetPortsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *CaptureSwitchSwitches) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given map[string]map[string]interface{} and assigns it to the Ports field.
func (o *CaptureSwitchSwitches) SetPorts(v map[string]map[string]interface{}) {
	o.Ports = &v
}

func (o CaptureSwitchSwitches) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureSwitchSwitches) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaptureSwitchSwitches) UnmarshalJSON(data []byte) (err error) {
	varCaptureSwitchSwitches := _CaptureSwitchSwitches{}

	err = json.Unmarshal(data, &varCaptureSwitchSwitches)

	if err != nil {
		return err
	}

	*o = CaptureSwitchSwitches(varCaptureSwitchSwitches)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaptureSwitchSwitches struct {
	value *CaptureSwitchSwitches
	isSet bool
}

func (v NullableCaptureSwitchSwitches) Get() *CaptureSwitchSwitches {
	return v.value
}

func (v *NullableCaptureSwitchSwitches) Set(val *CaptureSwitchSwitches) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureSwitchSwitches) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureSwitchSwitches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureSwitchSwitches(val *CaptureSwitchSwitches) *NullableCaptureSwitchSwitches {
	return &NullableCaptureSwitchSwitches{value: val, isSet: true}
}

func (v NullableCaptureSwitchSwitches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureSwitchSwitches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


