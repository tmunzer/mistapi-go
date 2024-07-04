/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the VrrpConfigGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrrpConfigGroup{}

// VrrpConfigGroup struct for VrrpConfigGroup
type VrrpConfigGroup struct {
	Priority *int32 `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrrpConfigGroup VrrpConfigGroup

// NewVrrpConfigGroup instantiates a new VrrpConfigGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrrpConfigGroup() *VrrpConfigGroup {
	this := VrrpConfigGroup{}
	return &this
}

// NewVrrpConfigGroupWithDefaults instantiates a new VrrpConfigGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrrpConfigGroupWithDefaults() *VrrpConfigGroup {
	this := VrrpConfigGroup{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VrrpConfigGroup) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpConfigGroup) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VrrpConfigGroup) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *VrrpConfigGroup) SetPriority(v int32) {
	o.Priority = &v
}

func (o VrrpConfigGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrrpConfigGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrrpConfigGroup) UnmarshalJSON(data []byte) (err error) {
	varVrrpConfigGroup := _VrrpConfigGroup{}

	err = json.Unmarshal(data, &varVrrpConfigGroup)

	if err != nil {
		return err
	}

	*o = VrrpConfigGroup(varVrrpConfigGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrrpConfigGroup struct {
	value *VrrpConfigGroup
	isSet bool
}

func (v NullableVrrpConfigGroup) Get() *VrrpConfigGroup {
	return v.value
}

func (v *NullableVrrpConfigGroup) Set(val *VrrpConfigGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVrrpConfigGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVrrpConfigGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrrpConfigGroup(val *VrrpConfigGroup) *NullableVrrpConfigGroup {
	return &NullableVrrpConfigGroup{value: val, isSet: true}
}

func (v NullableVrrpConfigGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrrpConfigGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


