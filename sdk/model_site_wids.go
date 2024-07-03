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

// checks if the SiteWids type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteWids{}

// SiteWids WIDS site settings
type SiteWids struct {
	RepeatedAuthFailures *SiteWidsRepeatedAuthFailures `json:"repeated_auth_failures,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteWids SiteWids

// NewSiteWids instantiates a new SiteWids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteWids() *SiteWids {
	this := SiteWids{}
	return &this
}

// NewSiteWidsWithDefaults instantiates a new SiteWids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWidsWithDefaults() *SiteWids {
	this := SiteWids{}
	return &this
}

// GetRepeatedAuthFailures returns the RepeatedAuthFailures field value if set, zero value otherwise.
func (o *SiteWids) GetRepeatedAuthFailures() SiteWidsRepeatedAuthFailures {
	if o == nil || IsNil(o.RepeatedAuthFailures) {
		var ret SiteWidsRepeatedAuthFailures
		return ret
	}
	return *o.RepeatedAuthFailures
}

// GetRepeatedAuthFailuresOk returns a tuple with the RepeatedAuthFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWids) GetRepeatedAuthFailuresOk() (*SiteWidsRepeatedAuthFailures, bool) {
	if o == nil || IsNil(o.RepeatedAuthFailures) {
		return nil, false
	}
	return o.RepeatedAuthFailures, true
}

// HasRepeatedAuthFailures returns a boolean if a field has been set.
func (o *SiteWids) HasRepeatedAuthFailures() bool {
	if o != nil && !IsNil(o.RepeatedAuthFailures) {
		return true
	}

	return false
}

// SetRepeatedAuthFailures gets a reference to the given SiteWidsRepeatedAuthFailures and assigns it to the RepeatedAuthFailures field.
func (o *SiteWids) SetRepeatedAuthFailures(v SiteWidsRepeatedAuthFailures) {
	o.RepeatedAuthFailures = &v
}

func (o SiteWids) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteWids) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepeatedAuthFailures) {
		toSerialize["repeated_auth_failures"] = o.RepeatedAuthFailures
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteWids) UnmarshalJSON(data []byte) (err error) {
	varSiteWids := _SiteWids{}

	err = json.Unmarshal(data, &varSiteWids)

	if err != nil {
		return err
	}

	*o = SiteWids(varSiteWids)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repeated_auth_failures")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteWids struct {
	value *SiteWids
	isSet bool
}

func (v NullableSiteWids) Get() *SiteWids {
	return v.value
}

func (v *NullableSiteWids) Set(val *SiteWids) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWids) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWids(val *SiteWids) *NullableSiteWids {
	return &NullableSiteWids{value: val, isSet: true}
}

func (v NullableSiteWids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


