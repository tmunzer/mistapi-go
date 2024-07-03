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

// checks if the SearchWxtagAppsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchWxtagAppsItem{}

// SearchWxtagAppsItem struct for SearchWxtagAppsItem
type SearchWxtagAppsItem struct {
	Group string `json:"group"`
	Key string `json:"key"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _SearchWxtagAppsItem SearchWxtagAppsItem

// NewSearchWxtagAppsItem instantiates a new SearchWxtagAppsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchWxtagAppsItem(group string, key string, name string) *SearchWxtagAppsItem {
	this := SearchWxtagAppsItem{}
	this.Group = group
	this.Key = key
	this.Name = name
	return &this
}

// NewSearchWxtagAppsItemWithDefaults instantiates a new SearchWxtagAppsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWxtagAppsItemWithDefaults() *SearchWxtagAppsItem {
	this := SearchWxtagAppsItem{}
	return &this
}

// GetGroup returns the Group field value
func (o *SearchWxtagAppsItem) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SearchWxtagAppsItem) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SearchWxtagAppsItem) SetGroup(v string) {
	o.Group = v
}

// GetKey returns the Key field value
func (o *SearchWxtagAppsItem) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SearchWxtagAppsItem) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SearchWxtagAppsItem) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *SearchWxtagAppsItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchWxtagAppsItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchWxtagAppsItem) SetName(v string) {
	o.Name = v
}

func (o SearchWxtagAppsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchWxtagAppsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchWxtagAppsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"key",
		"name",
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

	varSearchWxtagAppsItem := _SearchWxtagAppsItem{}

	err = json.Unmarshal(data, &varSearchWxtagAppsItem)

	if err != nil {
		return err
	}

	*o = SearchWxtagAppsItem(varSearchWxtagAppsItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchWxtagAppsItem struct {
	value *SearchWxtagAppsItem
	isSet bool
}

func (v NullableSearchWxtagAppsItem) Get() *SearchWxtagAppsItem {
	return v.value
}

func (v *NullableSearchWxtagAppsItem) Set(val *SearchWxtagAppsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchWxtagAppsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchWxtagAppsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchWxtagAppsItem(val *SearchWxtagAppsItem) *NullableSearchWxtagAppsItem {
	return &NullableSearchWxtagAppsItem{value: val, isSet: true}
}

func (v NullableSearchWxtagAppsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchWxtagAppsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


