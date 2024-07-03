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

// checks if the HaClusterConfigNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HaClusterConfigNode{}

// HaClusterConfigNode struct for HaClusterConfigNode
type HaClusterConfigNode struct {
	// node mac, should be unassigned
	Mac *string `json:"mac,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HaClusterConfigNode HaClusterConfigNode

// NewHaClusterConfigNode instantiates a new HaClusterConfigNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHaClusterConfigNode() *HaClusterConfigNode {
	this := HaClusterConfigNode{}
	return &this
}

// NewHaClusterConfigNodeWithDefaults instantiates a new HaClusterConfigNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHaClusterConfigNodeWithDefaults() *HaClusterConfigNode {
	this := HaClusterConfigNode{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *HaClusterConfigNode) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaClusterConfigNode) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *HaClusterConfigNode) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *HaClusterConfigNode) SetMac(v string) {
	o.Mac = &v
}

func (o HaClusterConfigNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HaClusterConfigNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HaClusterConfigNode) UnmarshalJSON(data []byte) (err error) {
	varHaClusterConfigNode := _HaClusterConfigNode{}

	err = json.Unmarshal(data, &varHaClusterConfigNode)

	if err != nil {
		return err
	}

	*o = HaClusterConfigNode(varHaClusterConfigNode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHaClusterConfigNode struct {
	value *HaClusterConfigNode
	isSet bool
}

func (v NullableHaClusterConfigNode) Get() *HaClusterConfigNode {
	return v.value
}

func (v *NullableHaClusterConfigNode) Set(val *HaClusterConfigNode) {
	v.value = val
	v.isSet = true
}

func (v NullableHaClusterConfigNode) IsSet() bool {
	return v.isSet
}

func (v *NullableHaClusterConfigNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHaClusterConfigNode(val *HaClusterConfigNode) *NullableHaClusterConfigNode {
	return &NullableHaClusterConfigNode{value: val, isSet: true}
}

func (v NullableHaClusterConfigNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHaClusterConfigNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


