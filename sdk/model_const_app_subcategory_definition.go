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
	"fmt"
)

// checks if the ConstAppSubcategoryDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstAppSubcategoryDefinition{}

// ConstAppSubcategoryDefinition struct for ConstAppSubcategoryDefinition
type ConstAppSubcategoryDefinition struct {
	// Description of the app subcategory
	Display string `json:"display"`
	// Key name of the app subcategory
	Key string `json:"key"`
	// Type of traffic (QoS) of the app subcategory
	TrafficType string `json:"traffic_type"`
	AdditionalProperties map[string]interface{}
}

type _ConstAppSubcategoryDefinition ConstAppSubcategoryDefinition

// NewConstAppSubcategoryDefinition instantiates a new ConstAppSubcategoryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstAppSubcategoryDefinition(display string, key string, trafficType string) *ConstAppSubcategoryDefinition {
	this := ConstAppSubcategoryDefinition{}
	this.Display = display
	this.Key = key
	this.TrafficType = trafficType
	return &this
}

// NewConstAppSubcategoryDefinitionWithDefaults instantiates a new ConstAppSubcategoryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstAppSubcategoryDefinitionWithDefaults() *ConstAppSubcategoryDefinition {
	this := ConstAppSubcategoryDefinition{}
	return &this
}

// GetDisplay returns the Display field value
func (o *ConstAppSubcategoryDefinition) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConstAppSubcategoryDefinition) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConstAppSubcategoryDefinition) SetDisplay(v string) {
	o.Display = v
}

// GetKey returns the Key field value
func (o *ConstAppSubcategoryDefinition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConstAppSubcategoryDefinition) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ConstAppSubcategoryDefinition) SetKey(v string) {
	o.Key = v
}

// GetTrafficType returns the TrafficType field value
func (o *ConstAppSubcategoryDefinition) GetTrafficType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrafficType
}

// GetTrafficTypeOk returns a tuple with the TrafficType field value
// and a boolean to check if the value has been set.
func (o *ConstAppSubcategoryDefinition) GetTrafficTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficType, true
}

// SetTrafficType sets field value
func (o *ConstAppSubcategoryDefinition) SetTrafficType(v string) {
	o.TrafficType = v
}

func (o ConstAppSubcategoryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstAppSubcategoryDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display"] = o.Display
	toSerialize["key"] = o.Key
	toSerialize["traffic_type"] = o.TrafficType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstAppSubcategoryDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"key",
		"traffic_type",
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

	varConstAppSubcategoryDefinition := _ConstAppSubcategoryDefinition{}

	err = json.Unmarshal(data, &varConstAppSubcategoryDefinition)

	if err != nil {
		return err
	}

	*o = ConstAppSubcategoryDefinition(varConstAppSubcategoryDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "key")
		delete(additionalProperties, "traffic_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstAppSubcategoryDefinition struct {
	value *ConstAppSubcategoryDefinition
	isSet bool
}

func (v NullableConstAppSubcategoryDefinition) Get() *ConstAppSubcategoryDefinition {
	return v.value
}

func (v *NullableConstAppSubcategoryDefinition) Set(val *ConstAppSubcategoryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableConstAppSubcategoryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableConstAppSubcategoryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstAppSubcategoryDefinition(val *ConstAppSubcategoryDefinition) *NullableConstAppSubcategoryDefinition {
	return &NullableConstAppSubcategoryDefinition{value: val, isSet: true}
}

func (v NullableConstAppSubcategoryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstAppSubcategoryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


