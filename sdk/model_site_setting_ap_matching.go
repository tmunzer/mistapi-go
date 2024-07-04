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

// checks if the SiteSettingApMatching type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingApMatching{}

// SiteSettingApMatching struct for SiteSettingApMatching
type SiteSettingApMatching struct {
	Enabled *bool `json:"enabled,omitempty"`
	Rules []SiteSettingApMatchingRule `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingApMatching SiteSettingApMatching

// NewSiteSettingApMatching instantiates a new SiteSettingApMatching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingApMatching() *SiteSettingApMatching {
	this := SiteSettingApMatching{}
	return &this
}

// NewSiteSettingApMatchingWithDefaults instantiates a new SiteSettingApMatching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingApMatchingWithDefaults() *SiteSettingApMatching {
	this := SiteSettingApMatching{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SiteSettingApMatching) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApMatching) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SiteSettingApMatching) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SiteSettingApMatching) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SiteSettingApMatching) GetRules() []SiteSettingApMatchingRule {
	if o == nil || IsNil(o.Rules) {
		var ret []SiteSettingApMatchingRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApMatching) GetRulesOk() ([]SiteSettingApMatchingRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SiteSettingApMatching) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SiteSettingApMatchingRule and assigns it to the Rules field.
func (o *SiteSettingApMatching) SetRules(v []SiteSettingApMatchingRule) {
	o.Rules = v
}

func (o SiteSettingApMatching) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingApMatching) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingApMatching) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingApMatching := _SiteSettingApMatching{}

	err = json.Unmarshal(data, &varSiteSettingApMatching)

	if err != nil {
		return err
	}

	*o = SiteSettingApMatching(varSiteSettingApMatching)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingApMatching struct {
	value *SiteSettingApMatching
	isSet bool
}

func (v NullableSiteSettingApMatching) Get() *SiteSettingApMatching {
	return v.value
}

func (v *NullableSiteSettingApMatching) Set(val *SiteSettingApMatching) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingApMatching) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingApMatching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingApMatching(val *SiteSettingApMatching) *NullableSiteSettingApMatching {
	return &NullableSiteSettingApMatching{value: val, isSet: true}
}

func (v NullableSiteSettingApMatching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingApMatching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


