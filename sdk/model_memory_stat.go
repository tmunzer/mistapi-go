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
	"fmt"
)

// checks if the MemoryStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryStat{}

// MemoryStat memory usage stat (for virtual chassis, memory usage of master RE)
type MemoryStat struct {
	Usage float32 `json:"usage"`
	AdditionalProperties map[string]interface{}
}

type _MemoryStat MemoryStat

// NewMemoryStat instantiates a new MemoryStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryStat(usage float32) *MemoryStat {
	this := MemoryStat{}
	this.Usage = usage
	return &this
}

// NewMemoryStatWithDefaults instantiates a new MemoryStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryStatWithDefaults() *MemoryStat {
	this := MemoryStat{}
	return &this
}

// GetUsage returns the Usage field value
func (o *MemoryStat) GetUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *MemoryStat) GetUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *MemoryStat) SetUsage(v float32) {
	o.Usage = v
}

func (o MemoryStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryStat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMemoryStat := _MemoryStat{}

	err = json.Unmarshal(data, &varMemoryStat)

	if err != nil {
		return err
	}

	*o = MemoryStat(varMemoryStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryStat struct {
	value *MemoryStat
	isSet bool
}

func (v NullableMemoryStat) Get() *MemoryStat {
	return v.value
}

func (v *NullableMemoryStat) Set(val *MemoryStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryStat(val *MemoryStat) *NullableMemoryStat {
	return &NullableMemoryStat{value: val, isSet: true}
}

func (v NullableMemoryStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


