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

// checks if the SleHistogramDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleHistogramDataItem{}

// SleHistogramDataItem struct for SleHistogramDataItem
type SleHistogramDataItem struct {
	Range []*float32 `json:"range,omitempty"`
	Value float32 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _SleHistogramDataItem SleHistogramDataItem

// NewSleHistogramDataItem instantiates a new SleHistogramDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleHistogramDataItem(value float32) *SleHistogramDataItem {
	this := SleHistogramDataItem{}
	this.Value = value
	return &this
}

// NewSleHistogramDataItemWithDefaults instantiates a new SleHistogramDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleHistogramDataItemWithDefaults() *SleHistogramDataItem {
	this := SleHistogramDataItem{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *SleHistogramDataItem) GetRange() []*float32 {
	if o == nil || IsNil(o.Range) {
		var ret []*float32
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleHistogramDataItem) GetRangeOk() ([]*float32, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *SleHistogramDataItem) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given []*float32 and assigns it to the Range field.
func (o *SleHistogramDataItem) SetRange(v []*float32) {
	o.Range = v
}

// GetValue returns the Value field value
func (o *SleHistogramDataItem) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SleHistogramDataItem) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SleHistogramDataItem) SetValue(v float32) {
	o.Value = v
}

func (o SleHistogramDataItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleHistogramDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleHistogramDataItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varSleHistogramDataItem := _SleHistogramDataItem{}

	err = json.Unmarshal(data, &varSleHistogramDataItem)

	if err != nil {
		return err
	}

	*o = SleHistogramDataItem(varSleHistogramDataItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "range")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleHistogramDataItem struct {
	value *SleHistogramDataItem
	isSet bool
}

func (v NullableSleHistogramDataItem) Get() *SleHistogramDataItem {
	return v.value
}

func (v *NullableSleHistogramDataItem) Set(val *SleHistogramDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSleHistogramDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSleHistogramDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleHistogramDataItem(val *SleHistogramDataItem) *NullableSleHistogramDataItem {
	return &NullableSleHistogramDataItem{value: val, isSet: true}
}

func (v NullableSleHistogramDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleHistogramDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


