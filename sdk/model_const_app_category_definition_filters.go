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

// checks if the ConstAppCategoryDefinitionFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstAppCategoryDefinitionFilters{}

// ConstAppCategoryDefinitionFilters struct for ConstAppCategoryDefinitionFilters
type ConstAppCategoryDefinitionFilters struct {
	Srx []string `json:"srx,omitempty"`
	Ssr []string `json:"ssr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstAppCategoryDefinitionFilters ConstAppCategoryDefinitionFilters

// NewConstAppCategoryDefinitionFilters instantiates a new ConstAppCategoryDefinitionFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstAppCategoryDefinitionFilters() *ConstAppCategoryDefinitionFilters {
	this := ConstAppCategoryDefinitionFilters{}
	return &this
}

// NewConstAppCategoryDefinitionFiltersWithDefaults instantiates a new ConstAppCategoryDefinitionFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstAppCategoryDefinitionFiltersWithDefaults() *ConstAppCategoryDefinitionFilters {
	this := ConstAppCategoryDefinitionFilters{}
	return &this
}

// GetSrx returns the Srx field value if set, zero value otherwise.
func (o *ConstAppCategoryDefinitionFilters) GetSrx() []string {
	if o == nil || IsNil(o.Srx) {
		var ret []string
		return ret
	}
	return o.Srx
}

// GetSrxOk returns a tuple with the Srx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinitionFilters) GetSrxOk() ([]string, bool) {
	if o == nil || IsNil(o.Srx) {
		return nil, false
	}
	return o.Srx, true
}

// HasSrx returns a boolean if a field has been set.
func (o *ConstAppCategoryDefinitionFilters) HasSrx() bool {
	if o != nil && !IsNil(o.Srx) {
		return true
	}

	return false
}

// SetSrx gets a reference to the given []string and assigns it to the Srx field.
func (o *ConstAppCategoryDefinitionFilters) SetSrx(v []string) {
	o.Srx = v
}

// GetSsr returns the Ssr field value if set, zero value otherwise.
func (o *ConstAppCategoryDefinitionFilters) GetSsr() []string {
	if o == nil || IsNil(o.Ssr) {
		var ret []string
		return ret
	}
	return o.Ssr
}

// GetSsrOk returns a tuple with the Ssr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstAppCategoryDefinitionFilters) GetSsrOk() ([]string, bool) {
	if o == nil || IsNil(o.Ssr) {
		return nil, false
	}
	return o.Ssr, true
}

// HasSsr returns a boolean if a field has been set.
func (o *ConstAppCategoryDefinitionFilters) HasSsr() bool {
	if o != nil && !IsNil(o.Ssr) {
		return true
	}

	return false
}

// SetSsr gets a reference to the given []string and assigns it to the Ssr field.
func (o *ConstAppCategoryDefinitionFilters) SetSsr(v []string) {
	o.Ssr = v
}

func (o ConstAppCategoryDefinitionFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstAppCategoryDefinitionFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Srx) {
		toSerialize["srx"] = o.Srx
	}
	if !IsNil(o.Ssr) {
		toSerialize["ssr"] = o.Ssr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstAppCategoryDefinitionFilters) UnmarshalJSON(data []byte) (err error) {
	varConstAppCategoryDefinitionFilters := _ConstAppCategoryDefinitionFilters{}

	err = json.Unmarshal(data, &varConstAppCategoryDefinitionFilters)

	if err != nil {
		return err
	}

	*o = ConstAppCategoryDefinitionFilters(varConstAppCategoryDefinitionFilters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "srx")
		delete(additionalProperties, "ssr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstAppCategoryDefinitionFilters struct {
	value *ConstAppCategoryDefinitionFilters
	isSet bool
}

func (v NullableConstAppCategoryDefinitionFilters) Get() *ConstAppCategoryDefinitionFilters {
	return v.value
}

func (v *NullableConstAppCategoryDefinitionFilters) Set(val *ConstAppCategoryDefinitionFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableConstAppCategoryDefinitionFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableConstAppCategoryDefinitionFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstAppCategoryDefinitionFilters(val *ConstAppCategoryDefinitionFilters) *NullableConstAppCategoryDefinitionFilters {
	return &NullableConstAppCategoryDefinitionFilters{value: val, isSet: true}
}

func (v NullableConstAppCategoryDefinitionFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstAppCategoryDefinitionFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


