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

// checks if the UtilsClearMacs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsClearMacs{}

// UtilsClearMacs struct for UtilsClearMacs
type UtilsClearMacs struct {
	// a list of ports on which to clear mac addresses. must include logical unit number
	Ports []string `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsClearMacs UtilsClearMacs

// NewUtilsClearMacs instantiates a new UtilsClearMacs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsClearMacs() *UtilsClearMacs {
	this := UtilsClearMacs{}
	return &this
}

// NewUtilsClearMacsWithDefaults instantiates a new UtilsClearMacs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsClearMacsWithDefaults() *UtilsClearMacs {
	this := UtilsClearMacs{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *UtilsClearMacs) GetPorts() []string {
	if o == nil || IsNil(o.Ports) {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsClearMacs) GetPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *UtilsClearMacs) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *UtilsClearMacs) SetPorts(v []string) {
	o.Ports = v
}

func (o UtilsClearMacs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsClearMacs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsClearMacs) UnmarshalJSON(data []byte) (err error) {
	varUtilsClearMacs := _UtilsClearMacs{}

	err = json.Unmarshal(data, &varUtilsClearMacs)

	if err != nil {
		return err
	}

	*o = UtilsClearMacs(varUtilsClearMacs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsClearMacs struct {
	value *UtilsClearMacs
	isSet bool
}

func (v NullableUtilsClearMacs) Get() *UtilsClearMacs {
	return v.value
}

func (v *NullableUtilsClearMacs) Set(val *UtilsClearMacs) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsClearMacs) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsClearMacs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsClearMacs(val *UtilsClearMacs) *NullableUtilsClearMacs {
	return &NullableUtilsClearMacs{value: val, isSet: true}
}

func (v NullableUtilsClearMacs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsClearMacs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


