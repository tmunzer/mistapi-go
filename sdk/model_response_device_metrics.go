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

// checks if the ResponseDeviceMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceMetrics{}

// ResponseDeviceMetrics struct for ResponseDeviceMetrics
type ResponseDeviceMetrics struct {
	End int32 `json:"end"`
	Interval int32 `json:"interval"`
	Results []ResponseDeviceMetricsResultsItems `json:"results"`
	Rt []string `json:"rt,omitempty"`
	Start int32 `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceMetrics ResponseDeviceMetrics

// NewResponseDeviceMetrics instantiates a new ResponseDeviceMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceMetrics(end int32, interval int32, results []ResponseDeviceMetricsResultsItems, start int32) *ResponseDeviceMetrics {
	this := ResponseDeviceMetrics{}
	this.End = end
	this.Interval = interval
	this.Results = results
	this.Start = start
	return &this
}

// NewResponseDeviceMetricsWithDefaults instantiates a new ResponseDeviceMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceMetricsWithDefaults() *ResponseDeviceMetrics {
	this := ResponseDeviceMetrics{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseDeviceMetrics) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceMetrics) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseDeviceMetrics) SetEnd(v int32) {
	o.End = v
}

// GetInterval returns the Interval field value
func (o *ResponseDeviceMetrics) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceMetrics) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ResponseDeviceMetrics) SetInterval(v int32) {
	o.Interval = v
}

// GetResults returns the Results field value
func (o *ResponseDeviceMetrics) GetResults() []ResponseDeviceMetricsResultsItems {
	if o == nil {
		var ret []ResponseDeviceMetricsResultsItems
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceMetrics) GetResultsOk() ([]ResponseDeviceMetricsResultsItems, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseDeviceMetrics) SetResults(v []ResponseDeviceMetricsResultsItems) {
	o.Results = v
}

// GetRt returns the Rt field value if set, zero value otherwise.
func (o *ResponseDeviceMetrics) GetRt() []string {
	if o == nil || IsNil(o.Rt) {
		var ret []string
		return ret
	}
	return o.Rt
}

// GetRtOk returns a tuple with the Rt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceMetrics) GetRtOk() ([]string, bool) {
	if o == nil || IsNil(o.Rt) {
		return nil, false
	}
	return o.Rt, true
}

// HasRt returns a boolean if a field has been set.
func (o *ResponseDeviceMetrics) HasRt() bool {
	if o != nil && !IsNil(o.Rt) {
		return true
	}

	return false
}

// SetRt gets a reference to the given []string and assigns it to the Rt field.
func (o *ResponseDeviceMetrics) SetRt(v []string) {
	o.Rt = v
}

// GetStart returns the Start field value
func (o *ResponseDeviceMetrics) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceMetrics) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseDeviceMetrics) SetStart(v int32) {
	o.Start = v
}

func (o ResponseDeviceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["interval"] = o.Interval
	toSerialize["results"] = o.Results
	if !IsNil(o.Rt) {
		toSerialize["rt"] = o.Rt
	}
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"interval",
		"results",
		"start",
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

	varResponseDeviceMetrics := _ResponseDeviceMetrics{}

	err = json.Unmarshal(data, &varResponseDeviceMetrics)

	if err != nil {
		return err
	}

	*o = ResponseDeviceMetrics(varResponseDeviceMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "results")
		delete(additionalProperties, "rt")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceMetrics struct {
	value *ResponseDeviceMetrics
	isSet bool
}

func (v NullableResponseDeviceMetrics) Get() *ResponseDeviceMetrics {
	return v.value
}

func (v *NullableResponseDeviceMetrics) Set(val *ResponseDeviceMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceMetrics(val *ResponseDeviceMetrics) *NullableResponseDeviceMetrics {
	return &NullableResponseDeviceMetrics{value: val, isSet: true}
}

func (v NullableResponseDeviceMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

