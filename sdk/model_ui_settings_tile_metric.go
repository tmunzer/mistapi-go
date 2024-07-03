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

// checks if the UiSettingsTileMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiSettingsTileMetric{}

// UiSettingsTileMetric struct for UiSettingsTileMetric
type UiSettingsTileMetric struct {
	ApiName *string `json:"apiName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiSettingsTileMetric UiSettingsTileMetric

// NewUiSettingsTileMetric instantiates a new UiSettingsTileMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiSettingsTileMetric() *UiSettingsTileMetric {
	this := UiSettingsTileMetric{}
	return &this
}

// NewUiSettingsTileMetricWithDefaults instantiates a new UiSettingsTileMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiSettingsTileMetricWithDefaults() *UiSettingsTileMetric {
	this := UiSettingsTileMetric{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *UiSettingsTileMetric) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileMetric) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *UiSettingsTileMetric) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *UiSettingsTileMetric) SetApiName(v string) {
	o.ApiName = &v
}

func (o UiSettingsTileMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiSettingsTileMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiName) {
		toSerialize["apiName"] = o.ApiName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiSettingsTileMetric) UnmarshalJSON(data []byte) (err error) {
	varUiSettingsTileMetric := _UiSettingsTileMetric{}

	err = json.Unmarshal(data, &varUiSettingsTileMetric)

	if err != nil {
		return err
	}

	*o = UiSettingsTileMetric(varUiSettingsTileMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiSettingsTileMetric struct {
	value *UiSettingsTileMetric
	isSet bool
}

func (v NullableUiSettingsTileMetric) Get() *UiSettingsTileMetric {
	return v.value
}

func (v *NullableUiSettingsTileMetric) Set(val *UiSettingsTileMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableUiSettingsTileMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableUiSettingsTileMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiSettingsTileMetric(val *UiSettingsTileMetric) *NullableUiSettingsTileMetric {
	return &NullableUiSettingsTileMetric{value: val, isSet: true}
}

func (v NullableUiSettingsTileMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiSettingsTileMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


