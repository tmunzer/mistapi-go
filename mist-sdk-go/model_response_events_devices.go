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

// checks if the ResponseEventsDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEventsDevices{}

// ResponseEventsDevices struct for ResponseEventsDevices
type ResponseEventsDevices struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Next string `json:"next"`
	Results []EventsDeviceAp `json:"results"`
	Start int32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _ResponseEventsDevices ResponseEventsDevices

// NewResponseEventsDevices instantiates a new ResponseEventsDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEventsDevices(end int32, limit int32, next string, results []EventsDeviceAp, start int32, total int32) *ResponseEventsDevices {
	this := ResponseEventsDevices{}
	this.End = end
	this.Limit = limit
	this.Next = next
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewResponseEventsDevicesWithDefaults instantiates a new ResponseEventsDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEventsDevicesWithDefaults() *ResponseEventsDevices {
	this := ResponseEventsDevices{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseEventsDevices) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseEventsDevices) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseEventsDevices) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseEventsDevices) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value
func (o *ResponseEventsDevices) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *ResponseEventsDevices) SetNext(v string) {
	o.Next = v
}

// GetResults returns the Results field value
func (o *ResponseEventsDevices) GetResults() []EventsDeviceAp {
	if o == nil {
		var ret []EventsDeviceAp
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetResultsOk() ([]EventsDeviceAp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseEventsDevices) SetResults(v []EventsDeviceAp) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseEventsDevices) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseEventsDevices) SetStart(v int32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *ResponseEventsDevices) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsDevices) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseEventsDevices) SetTotal(v int32) {
	o.Total = v
}

func (o ResponseEventsDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEventsDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	toSerialize["next"] = o.Next
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseEventsDevices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
		"next",
		"results",
		"start",
		"total",
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

	varResponseEventsDevices := _ResponseEventsDevices{}

	err = json.Unmarshal(data, &varResponseEventsDevices)

	if err != nil {
		return err
	}

	*o = ResponseEventsDevices(varResponseEventsDevices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseEventsDevices struct {
	value *ResponseEventsDevices
	isSet bool
}

func (v NullableResponseEventsDevices) Get() *ResponseEventsDevices {
	return v.value
}

func (v *NullableResponseEventsDevices) Set(val *ResponseEventsDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEventsDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEventsDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEventsDevices(val *ResponseEventsDevices) *NullableResponseEventsDevices {
	return &NullableResponseEventsDevices{value: val, isSet: true}
}

func (v NullableResponseEventsDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEventsDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


