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

// checks if the ConstAppCategoryDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstAppCategoryDefinition{}

// ConstAppCategoryDefinition struct for ConstAppCategoryDefinition
type ConstAppCategoryDefinition struct {
	// Description of the app category
	Display string `json:"display"`
	Filters *ConstAppCategoryDefinitionFilters `json:"filters,omitempty"`
	// List of other App Categories contained by this one
	Includes []string `json:"includes,omitempty"`
	// Key name of the app category
	Key string `json:"key"`
	AdditionalProperties map[string]interface{}
}

type _ConstAppCategoryDefinition ConstAppCategoryDefinition

// NewConstAppCategoryDefinition instantiates a new ConstAppCategoryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstAppCategoryDefinition(display string, key string) *ConstAppCategoryDefinition {
	this := ConstAppCategoryDefinition{}
	this.Display = display
	this.Key = key
	return &this
}

// NewConstAppCategoryDefinitionWithDefaults instantiates a new ConstAppCategoryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstAppCategoryDefinitionWithDefaults() *ConstAppCategoryDefinition {
	this := ConstAppCategoryDefinition{}
	return &this
}

// GetDisplay returns the Display field value
func (o *ConstAppCategoryDefinition) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinition) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConstAppCategoryDefinition) SetDisplay(v string) {
	o.Display = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ConstAppCategoryDefinition) GetFilters() ConstAppCategoryDefinitionFilters {
	if o == nil || IsNil(o.Filters) {
		var ret ConstAppCategoryDefinitionFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinition) GetFiltersOk() (*ConstAppCategoryDefinitionFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ConstAppCategoryDefinition) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ConstAppCategoryDefinitionFilters and assigns it to the Filters field.
func (o *ConstAppCategoryDefinition) SetFilters(v ConstAppCategoryDefinitionFilters) {
	o.Filters = &v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *ConstAppCategoryDefinition) GetIncludes() []string {
	if o == nil || IsNil(o.Includes) {
		var ret []string
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinition) GetIncludesOk() ([]string, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *ConstAppCategoryDefinition) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given []string and assigns it to the Includes field.
func (o *ConstAppCategoryDefinition) SetIncludes(v []string) {
	o.Includes = v
}

// GetKey returns the Key field value
func (o *ConstAppCategoryDefinition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinition) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ConstAppCategoryDefinition) SetKey(v string) {
	o.Key = v
}

func (o ConstAppCategoryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstAppCategoryDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display"] = o.Display
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	toSerialize["key"] = o.Key

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstAppCategoryDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"key",
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

	varConstAppCategoryDefinition := _ConstAppCategoryDefinition{}

	err = json.Unmarshal(data, &varConstAppCategoryDefinition)

	if err != nil {
		return err
	}

	*o = ConstAppCategoryDefinition(varConstAppCategoryDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "includes")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstAppCategoryDefinition struct {
	value *ConstAppCategoryDefinition
	isSet bool
}

func (v NullableConstAppCategoryDefinition) Get() *ConstAppCategoryDefinition {
	return v.value
}

func (v *NullableConstAppCategoryDefinition) Set(val *ConstAppCategoryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableConstAppCategoryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableConstAppCategoryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstAppCategoryDefinition(val *ConstAppCategoryDefinition) *NullableConstAppCategoryDefinition {
	return &NullableConstAppCategoryDefinition{value: val, isSet: true}
}

func (v NullableConstAppCategoryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstAppCategoryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


