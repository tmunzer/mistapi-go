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

// checks if the ApPwrConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApPwrConfig{}

// ApPwrConfig power related configs
type ApPwrConfig struct {
	// additional power to request during negotiating with PSE over PoE, in mW
	Base *float32 `json:"base,omitempty"`
	// whether to enable power out to peripheral, meanwhile will reduce power to wifi (only for AP45 at power mode)
	PreferUsbOverWifi *bool `json:"prefer_usb_over_wifi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApPwrConfig ApPwrConfig

// NewApPwrConfig instantiates a new ApPwrConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApPwrConfig() *ApPwrConfig {
	this := ApPwrConfig{}
	var base float32 = 0
	this.Base = &base
	var preferUsbOverWifi bool = false
	this.PreferUsbOverWifi = &preferUsbOverWifi
	return &this
}

// NewApPwrConfigWithDefaults instantiates a new ApPwrConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApPwrConfigWithDefaults() *ApPwrConfig {
	this := ApPwrConfig{}
	var base float32 = 0
	this.Base = &base
	var preferUsbOverWifi bool = false
	this.PreferUsbOverWifi = &preferUsbOverWifi
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *ApPwrConfig) GetBase() float32 {
	if o == nil || IsNil(o.Base) {
		var ret float32
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPwrConfig) GetBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *ApPwrConfig) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given float32 and assigns it to the Base field.
func (o *ApPwrConfig) SetBase(v float32) {
	o.Base = &v
}

// GetPreferUsbOverWifi returns the PreferUsbOverWifi field value if set, zero value otherwise.
func (o *ApPwrConfig) GetPreferUsbOverWifi() bool {
	if o == nil || IsNil(o.PreferUsbOverWifi) {
		var ret bool
		return ret
	}
	return *o.PreferUsbOverWifi
}

// GetPreferUsbOverWifiOk returns a tuple with the PreferUsbOverWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApPwrConfig) GetPreferUsbOverWifiOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferUsbOverWifi) {
		return nil, false
	}
	return o.PreferUsbOverWifi, true
}

// HasPreferUsbOverWifi returns a boolean if a field has been set.
func (o *ApPwrConfig) HasPreferUsbOverWifi() bool {
	if o != nil && !IsNil(o.PreferUsbOverWifi) {
		return true
	}

	return false
}

// SetPreferUsbOverWifi gets a reference to the given bool and assigns it to the PreferUsbOverWifi field.
func (o *ApPwrConfig) SetPreferUsbOverWifi(v bool) {
	o.PreferUsbOverWifi = &v
}

func (o ApPwrConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApPwrConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.PreferUsbOverWifi) {
		toSerialize["prefer_usb_over_wifi"] = o.PreferUsbOverWifi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApPwrConfig) UnmarshalJSON(data []byte) (err error) {
	varApPwrConfig := _ApPwrConfig{}

	err = json.Unmarshal(data, &varApPwrConfig)

	if err != nil {
		return err
	}

	*o = ApPwrConfig(varApPwrConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base")
		delete(additionalProperties, "prefer_usb_over_wifi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApPwrConfig struct {
	value *ApPwrConfig
	isSet bool
}

func (v NullableApPwrConfig) Get() *ApPwrConfig {
	return v.value
}

func (v *NullableApPwrConfig) Set(val *ApPwrConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApPwrConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApPwrConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApPwrConfig(val *ApPwrConfig) *NullableApPwrConfig {
	return &NullableApPwrConfig{value: val, isSet: true}
}

func (v NullableApPwrConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApPwrConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

