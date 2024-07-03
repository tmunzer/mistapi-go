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

// checks if the UtilsShowBgpRummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsShowBgpRummary{}

// UtilsShowBgpRummary struct for UtilsShowBgpRummary
type UtilsShowBgpRummary struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsShowBgpRummary UtilsShowBgpRummary

// NewUtilsShowBgpRummary instantiates a new UtilsShowBgpRummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsShowBgpRummary() *UtilsShowBgpRummary {
	this := UtilsShowBgpRummary{}
	return &this
}

// NewUtilsShowBgpRummaryWithDefaults instantiates a new UtilsShowBgpRummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsShowBgpRummaryWithDefaults() *UtilsShowBgpRummary {
	this := UtilsShowBgpRummary{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *UtilsShowBgpRummary) GetNode() HaClusterNodeEnum {
	if o == nil || IsNil(o.Node) {
		var ret HaClusterNodeEnum
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowBgpRummary) GetNodeOk() (*HaClusterNodeEnum, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *UtilsShowBgpRummary) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given HaClusterNodeEnum and assigns it to the Node field.
func (o *UtilsShowBgpRummary) SetNode(v HaClusterNodeEnum) {
	o.Node = &v
}

func (o UtilsShowBgpRummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsShowBgpRummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsShowBgpRummary) UnmarshalJSON(data []byte) (err error) {
	varUtilsShowBgpRummary := _UtilsShowBgpRummary{}

	err = json.Unmarshal(data, &varUtilsShowBgpRummary)

	if err != nil {
		return err
	}

	*o = UtilsShowBgpRummary(varUtilsShowBgpRummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "node")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsShowBgpRummary struct {
	value *UtilsShowBgpRummary
	isSet bool
}

func (v NullableUtilsShowBgpRummary) Get() *UtilsShowBgpRummary {
	return v.value
}

func (v *NullableUtilsShowBgpRummary) Set(val *UtilsShowBgpRummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowBgpRummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowBgpRummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowBgpRummary(val *UtilsShowBgpRummary) *NullableUtilsShowBgpRummary {
	return &NullableUtilsShowBgpRummary{value: val, isSet: true}
}

func (v NullableUtilsShowBgpRummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowBgpRummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


