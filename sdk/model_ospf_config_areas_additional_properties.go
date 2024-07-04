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

// checks if the OspfConfigAreasAdditionalProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspfConfigAreasAdditionalProperties{}

// OspfConfigAreasAdditionalProperties struct for OspfConfigAreasAdditionalProperties
type OspfConfigAreasAdditionalProperties struct {
	// for a stub/nssa area, where to avoid forwarding type-3 LSA to this area
	NoSummary *bool `json:"no_summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OspfConfigAreasAdditionalProperties OspfConfigAreasAdditionalProperties

// NewOspfConfigAreasAdditionalProperties instantiates a new OspfConfigAreasAdditionalProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspfConfigAreasAdditionalProperties() *OspfConfigAreasAdditionalProperties {
	this := OspfConfigAreasAdditionalProperties{}
	return &this
}

// NewOspfConfigAreasAdditionalPropertiesWithDefaults instantiates a new OspfConfigAreasAdditionalProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspfConfigAreasAdditionalPropertiesWithDefaults() *OspfConfigAreasAdditionalProperties {
	this := OspfConfigAreasAdditionalProperties{}
	return &this
}

// GetNoSummary returns the NoSummary field value if set, zero value otherwise.
func (o *OspfConfigAreasAdditionalProperties) GetNoSummary() bool {
	if o == nil || IsNil(o.NoSummary) {
		var ret bool
		return ret
	}
	return *o.NoSummary
}

// GetNoSummaryOk returns a tuple with the NoSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspfConfigAreasAdditionalProperties) GetNoSummaryOk() (*bool, bool) {
	if o == nil || IsNil(o.NoSummary) {
		return nil, false
	}
	return o.NoSummary, true
}

// HasNoSummary returns a boolean if a field has been set.
func (o *OspfConfigAreasAdditionalProperties) HasNoSummary() bool {
	if o != nil && !IsNil(o.NoSummary) {
		return true
	}

	return false
}

// SetNoSummary gets a reference to the given bool and assigns it to the NoSummary field.
func (o *OspfConfigAreasAdditionalProperties) SetNoSummary(v bool) {
	o.NoSummary = &v
}

func (o OspfConfigAreasAdditionalProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspfConfigAreasAdditionalProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoSummary) {
		toSerialize["no_summary"] = o.NoSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OspfConfigAreasAdditionalProperties) UnmarshalJSON(data []byte) (err error) {
	varOspfConfigAreasAdditionalProperties := _OspfConfigAreasAdditionalProperties{}

	err = json.Unmarshal(data, &varOspfConfigAreasAdditionalProperties)

	if err != nil {
		return err
	}

	*o = OspfConfigAreasAdditionalProperties(varOspfConfigAreasAdditionalProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "no_summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOspfConfigAreasAdditionalProperties struct {
	value *OspfConfigAreasAdditionalProperties
	isSet bool
}

func (v NullableOspfConfigAreasAdditionalProperties) Get() *OspfConfigAreasAdditionalProperties {
	return v.value
}

func (v *NullableOspfConfigAreasAdditionalProperties) Set(val *OspfConfigAreasAdditionalProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableOspfConfigAreasAdditionalProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableOspfConfigAreasAdditionalProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspfConfigAreasAdditionalProperties(val *OspfConfigAreasAdditionalProperties) *NullableOspfConfigAreasAdditionalProperties {
	return &NullableOspfConfigAreasAdditionalProperties{value: val, isSet: true}
}

func (v NullableOspfConfigAreasAdditionalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspfConfigAreasAdditionalProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

