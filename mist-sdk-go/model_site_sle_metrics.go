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

// checks if the SiteSleMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSleMetrics{}

// SiteSleMetrics struct for SiteSleMetrics
type SiteSleMetrics struct {
	Enabled []string `json:"enabled"`
	Supported []string `json:"supported"`
	AdditionalProperties map[string]interface{}
}

type _SiteSleMetrics SiteSleMetrics

// NewSiteSleMetrics instantiates a new SiteSleMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSleMetrics(enabled []string, supported []string) *SiteSleMetrics {
	this := SiteSleMetrics{}
	this.Enabled = enabled
	this.Supported = supported
	return &this
}

// NewSiteSleMetricsWithDefaults instantiates a new SiteSleMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSleMetricsWithDefaults() *SiteSleMetrics {
	this := SiteSleMetrics{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SiteSleMetrics) GetEnabled() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SiteSleMetrics) GetEnabledOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *SiteSleMetrics) SetEnabled(v []string) {
	o.Enabled = v
}

// GetSupported returns the Supported field value
func (o *SiteSleMetrics) GetSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *SiteSleMetrics) GetSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Supported, true
}

// SetSupported sets field value
func (o *SiteSleMetrics) SetSupported(v []string) {
	o.Supported = v
}

func (o SiteSleMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSleMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["supported"] = o.Supported

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSleMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"supported",
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

	varSiteSleMetrics := _SiteSleMetrics{}

	err = json.Unmarshal(data, &varSiteSleMetrics)

	if err != nil {
		return err
	}

	*o = SiteSleMetrics(varSiteSleMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "supported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSleMetrics struct {
	value *SiteSleMetrics
	isSet bool
}

func (v NullableSiteSleMetrics) Get() *SiteSleMetrics {
	return v.value
}

func (v *NullableSiteSleMetrics) Set(val *SiteSleMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleMetrics(val *SiteSleMetrics) *NullableSiteSleMetrics {
	return &NullableSiteSleMetrics{value: val, isSet: true}
}

func (v NullableSiteSleMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


