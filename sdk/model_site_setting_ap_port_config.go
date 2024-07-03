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

// checks if the SiteSettingApPortConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingApPortConfig{}

// SiteSettingApPortConfig struct for SiteSettingApPortConfig
type SiteSettingApPortConfig struct {
	// Property key is the AP model (e.g \"AP32\")
	ModelSpecific *map[string]map[string]ApPortConfig `json:"model_specific,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingApPortConfig SiteSettingApPortConfig

// NewSiteSettingApPortConfig instantiates a new SiteSettingApPortConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingApPortConfig() *SiteSettingApPortConfig {
	this := SiteSettingApPortConfig{}
	return &this
}

// NewSiteSettingApPortConfigWithDefaults instantiates a new SiteSettingApPortConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingApPortConfigWithDefaults() *SiteSettingApPortConfig {
	this := SiteSettingApPortConfig{}
	return &this
}

// GetModelSpecific returns the ModelSpecific field value if set, zero value otherwise.
func (o *SiteSettingApPortConfig) GetModelSpecific() map[string]map[string]ApPortConfig {
	if o == nil || IsNil(o.ModelSpecific) {
		var ret map[string]map[string]ApPortConfig
		return ret
	}
	return *o.ModelSpecific
}

// GetModelSpecificOk returns a tuple with the ModelSpecific field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApPortConfig) GetModelSpecificOk() (*map[string]map[string]ApPortConfig, bool) {
	if o == nil || IsNil(o.ModelSpecific) {
		return nil, false
	}
	return o.ModelSpecific, true
}

// HasModelSpecific returns a boolean if a field has been set.
func (o *SiteSettingApPortConfig) HasModelSpecific() bool {
	if o != nil && !IsNil(o.ModelSpecific) {
		return true
	}

	return false
}

// SetModelSpecific gets a reference to the given map[string]map[string]ApPortConfig and assigns it to the ModelSpecific field.
func (o *SiteSettingApPortConfig) SetModelSpecific(v map[string]map[string]ApPortConfig) {
	o.ModelSpecific = &v
}

func (o SiteSettingApPortConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingApPortConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModelSpecific) {
		toSerialize["model_specific"] = o.ModelSpecific
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingApPortConfig) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingApPortConfig := _SiteSettingApPortConfig{}

	err = json.Unmarshal(data, &varSiteSettingApPortConfig)

	if err != nil {
		return err
	}

	*o = SiteSettingApPortConfig(varSiteSettingApPortConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "model_specific")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingApPortConfig struct {
	value *SiteSettingApPortConfig
	isSet bool
}

func (v NullableSiteSettingApPortConfig) Get() *SiteSettingApPortConfig {
	return v.value
}

func (v *NullableSiteSettingApPortConfig) Set(val *SiteSettingApPortConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingApPortConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingApPortConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingApPortConfig(val *SiteSettingApPortConfig) *NullableSiteSettingApPortConfig {
	return &NullableSiteSettingApPortConfig{value: val, isSet: true}
}

func (v NullableSiteSettingApPortConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingApPortConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


