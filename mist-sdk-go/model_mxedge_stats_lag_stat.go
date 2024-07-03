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

// checks if the MxedgeStatsLagStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeStatsLagStat{}

// MxedgeStatsLagStat struct for MxedgeStatsLagStat
type MxedgeStatsLagStat struct {
	// list of ports active on the LAG defined by the LACP
	ActivePorts []string `json:"active_ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeStatsLagStat MxedgeStatsLagStat

// NewMxedgeStatsLagStat instantiates a new MxedgeStatsLagStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeStatsLagStat() *MxedgeStatsLagStat {
	this := MxedgeStatsLagStat{}
	return &this
}

// NewMxedgeStatsLagStatWithDefaults instantiates a new MxedgeStatsLagStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeStatsLagStatWithDefaults() *MxedgeStatsLagStat {
	this := MxedgeStatsLagStat{}
	return &this
}

// GetActivePorts returns the ActivePorts field value if set, zero value otherwise.
func (o *MxedgeStatsLagStat) GetActivePorts() []string {
	if o == nil || IsNil(o.ActivePorts) {
		var ret []string
		return ret
	}
	return o.ActivePorts
}

// GetActivePortsOk returns a tuple with the ActivePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsLagStat) GetActivePortsOk() ([]string, bool) {
	if o == nil || IsNil(o.ActivePorts) {
		return nil, false
	}
	return o.ActivePorts, true
}

// HasActivePorts returns a boolean if a field has been set.
func (o *MxedgeStatsLagStat) HasActivePorts() bool {
	if o != nil && !IsNil(o.ActivePorts) {
		return true
	}

	return false
}

// SetActivePorts gets a reference to the given []string and assigns it to the ActivePorts field.
func (o *MxedgeStatsLagStat) SetActivePorts(v []string) {
	o.ActivePorts = v
}

func (o MxedgeStatsLagStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeStatsLagStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivePorts) {
		toSerialize["active_ports"] = o.ActivePorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeStatsLagStat) UnmarshalJSON(data []byte) (err error) {
	varMxedgeStatsLagStat := _MxedgeStatsLagStat{}

	err = json.Unmarshal(data, &varMxedgeStatsLagStat)

	if err != nil {
		return err
	}

	*o = MxedgeStatsLagStat(varMxedgeStatsLagStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeStatsLagStat struct {
	value *MxedgeStatsLagStat
	isSet bool
}

func (v NullableMxedgeStatsLagStat) Get() *MxedgeStatsLagStat {
	return v.value
}

func (v *NullableMxedgeStatsLagStat) Set(val *MxedgeStatsLagStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeStatsLagStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeStatsLagStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeStatsLagStat(val *MxedgeStatsLagStat) *NullableMxedgeStatsLagStat {
	return &NullableMxedgeStatsLagStat{value: val, isSet: true}
}

func (v NullableMxedgeStatsLagStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeStatsLagStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


