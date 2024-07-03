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

// checks if the AlarmTemplateRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmTemplateRule{}

// AlarmTemplateRule struct for AlarmTemplateRule
type AlarmTemplateRule struct {
	Delivery *Delivery `json:"delivery,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlarmTemplateRule AlarmTemplateRule

// NewAlarmTemplateRule instantiates a new AlarmTemplateRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmTemplateRule() *AlarmTemplateRule {
	this := AlarmTemplateRule{}
	return &this
}

// NewAlarmTemplateRuleWithDefaults instantiates a new AlarmTemplateRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmTemplateRuleWithDefaults() *AlarmTemplateRule {
	this := AlarmTemplateRule{}
	return &this
}

// GetDelivery returns the Delivery field value if set, zero value otherwise.
func (o *AlarmTemplateRule) GetDelivery() Delivery {
	if o == nil || IsNil(o.Delivery) {
		var ret Delivery
		return ret
	}
	return *o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmTemplateRule) GetDeliveryOk() (*Delivery, bool) {
	if o == nil || IsNil(o.Delivery) {
		return nil, false
	}
	return o.Delivery, true
}

// HasDelivery returns a boolean if a field has been set.
func (o *AlarmTemplateRule) HasDelivery() bool {
	if o != nil && !IsNil(o.Delivery) {
		return true
	}

	return false
}

// SetDelivery gets a reference to the given Delivery and assigns it to the Delivery field.
func (o *AlarmTemplateRule) SetDelivery(v Delivery) {
	o.Delivery = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlarmTemplateRule) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmTemplateRule) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlarmTemplateRule) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlarmTemplateRule) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AlarmTemplateRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmTemplateRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Delivery) {
		toSerialize["delivery"] = o.Delivery
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlarmTemplateRule) UnmarshalJSON(data []byte) (err error) {
	varAlarmTemplateRule := _AlarmTemplateRule{}

	err = json.Unmarshal(data, &varAlarmTemplateRule)

	if err != nil {
		return err
	}

	*o = AlarmTemplateRule(varAlarmTemplateRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delivery")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlarmTemplateRule struct {
	value *AlarmTemplateRule
	isSet bool
}

func (v NullableAlarmTemplateRule) Get() *AlarmTemplateRule {
	return v.value
}

func (v *NullableAlarmTemplateRule) Set(val *AlarmTemplateRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmTemplateRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmTemplateRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmTemplateRule(val *AlarmTemplateRule) *NullableAlarmTemplateRule {
	return &NullableAlarmTemplateRule{value: val, isSet: true}
}

func (v NullableAlarmTemplateRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmTemplateRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


