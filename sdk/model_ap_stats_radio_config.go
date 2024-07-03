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

// checks if the ApStatsRadioConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsRadioConfig{}

// ApStatsRadioConfig struct for ApStatsRadioConfig
type ApStatsRadioConfig struct {
	Band24 *ApStatsRadioConfigBand `json:"band_24,omitempty"`
	Band5 *ApStatsRadioConfigBand `json:"band_5,omitempty"`
	Band6 *ApStatsRadioConfigBand `json:"band_6,omitempty"`
	ScanningEnabled *bool `json:"scanning_enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsRadioConfig ApStatsRadioConfig

// NewApStatsRadioConfig instantiates a new ApStatsRadioConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsRadioConfig() *ApStatsRadioConfig {
	this := ApStatsRadioConfig{}
	return &this
}

// NewApStatsRadioConfigWithDefaults instantiates a new ApStatsRadioConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsRadioConfigWithDefaults() *ApStatsRadioConfig {
	this := ApStatsRadioConfig{}
	return &this
}

// GetBand24 returns the Band24 field value if set, zero value otherwise.
func (o *ApStatsRadioConfig) GetBand24() ApStatsRadioConfigBand {
	if o == nil || IsNil(o.Band24) {
		var ret ApStatsRadioConfigBand
		return ret
	}
	return *o.Band24
}

// GetBand24Ok returns a tuple with the Band24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfig) GetBand24Ok() (*ApStatsRadioConfigBand, bool) {
	if o == nil || IsNil(o.Band24) {
		return nil, false
	}
	return o.Band24, true
}

// HasBand24 returns a boolean if a field has been set.
func (o *ApStatsRadioConfig) HasBand24() bool {
	if o != nil && !IsNil(o.Band24) {
		return true
	}

	return false
}

// SetBand24 gets a reference to the given ApStatsRadioConfigBand and assigns it to the Band24 field.
func (o *ApStatsRadioConfig) SetBand24(v ApStatsRadioConfigBand) {
	o.Band24 = &v
}

// GetBand5 returns the Band5 field value if set, zero value otherwise.
func (o *ApStatsRadioConfig) GetBand5() ApStatsRadioConfigBand {
	if o == nil || IsNil(o.Band5) {
		var ret ApStatsRadioConfigBand
		return ret
	}
	return *o.Band5
}

// GetBand5Ok returns a tuple with the Band5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfig) GetBand5Ok() (*ApStatsRadioConfigBand, bool) {
	if o == nil || IsNil(o.Band5) {
		return nil, false
	}
	return o.Band5, true
}

// HasBand5 returns a boolean if a field has been set.
func (o *ApStatsRadioConfig) HasBand5() bool {
	if o != nil && !IsNil(o.Band5) {
		return true
	}

	return false
}

// SetBand5 gets a reference to the given ApStatsRadioConfigBand and assigns it to the Band5 field.
func (o *ApStatsRadioConfig) SetBand5(v ApStatsRadioConfigBand) {
	o.Band5 = &v
}

// GetBand6 returns the Band6 field value if set, zero value otherwise.
func (o *ApStatsRadioConfig) GetBand6() ApStatsRadioConfigBand {
	if o == nil || IsNil(o.Band6) {
		var ret ApStatsRadioConfigBand
		return ret
	}
	return *o.Band6
}

// GetBand6Ok returns a tuple with the Band6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfig) GetBand6Ok() (*ApStatsRadioConfigBand, bool) {
	if o == nil || IsNil(o.Band6) {
		return nil, false
	}
	return o.Band6, true
}

// HasBand6 returns a boolean if a field has been set.
func (o *ApStatsRadioConfig) HasBand6() bool {
	if o != nil && !IsNil(o.Band6) {
		return true
	}

	return false
}

// SetBand6 gets a reference to the given ApStatsRadioConfigBand and assigns it to the Band6 field.
func (o *ApStatsRadioConfig) SetBand6(v ApStatsRadioConfigBand) {
	o.Band6 = &v
}

// GetScanningEnabled returns the ScanningEnabled field value if set, zero value otherwise.
func (o *ApStatsRadioConfig) GetScanningEnabled() bool {
	if o == nil || IsNil(o.ScanningEnabled) {
		var ret bool
		return ret
	}
	return *o.ScanningEnabled
}

// GetScanningEnabledOk returns a tuple with the ScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfig) GetScanningEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScanningEnabled) {
		return nil, false
	}
	return o.ScanningEnabled, true
}

// HasScanningEnabled returns a boolean if a field has been set.
func (o *ApStatsRadioConfig) HasScanningEnabled() bool {
	if o != nil && !IsNil(o.ScanningEnabled) {
		return true
	}

	return false
}

// SetScanningEnabled gets a reference to the given bool and assigns it to the ScanningEnabled field.
func (o *ApStatsRadioConfig) SetScanningEnabled(v bool) {
	o.ScanningEnabled = &v
}

func (o ApStatsRadioConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsRadioConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band24) {
		toSerialize["band_24"] = o.Band24
	}
	if !IsNil(o.Band5) {
		toSerialize["band_5"] = o.Band5
	}
	if !IsNil(o.Band6) {
		toSerialize["band_6"] = o.Band6
	}
	if !IsNil(o.ScanningEnabled) {
		toSerialize["scanning_enabled"] = o.ScanningEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsRadioConfig) UnmarshalJSON(data []byte) (err error) {
	varApStatsRadioConfig := _ApStatsRadioConfig{}

	err = json.Unmarshal(data, &varApStatsRadioConfig)

	if err != nil {
		return err
	}

	*o = ApStatsRadioConfig(varApStatsRadioConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band_24")
		delete(additionalProperties, "band_5")
		delete(additionalProperties, "band_6")
		delete(additionalProperties, "scanning_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsRadioConfig struct {
	value *ApStatsRadioConfig
	isSet bool
}

func (v NullableApStatsRadioConfig) Get() *ApStatsRadioConfig {
	return v.value
}

func (v *NullableApStatsRadioConfig) Set(val *ApStatsRadioConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsRadioConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsRadioConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsRadioConfig(val *ApStatsRadioConfig) *NullableApStatsRadioConfig {
	return &NullableApStatsRadioConfig{value: val, isSet: true}
}

func (v NullableApStatsRadioConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsRadioConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


