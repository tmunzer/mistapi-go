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

// checks if the ResponseEventsFastroam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEventsFastroam{}

// ResponseEventsFastroam struct for ResponseEventsFastroam
type ResponseEventsFastroam struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	// the link to query next set of results. value is null if no next page exists.
	Next *string `json:"next,omitempty"`
	Results []EventFastroam `json:"results"`
	Start int32 `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _ResponseEventsFastroam ResponseEventsFastroam

// NewResponseEventsFastroam instantiates a new ResponseEventsFastroam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEventsFastroam(end int32, limit int32, results []EventFastroam, start int32) *ResponseEventsFastroam {
	this := ResponseEventsFastroam{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	return &this
}

// NewResponseEventsFastroamWithDefaults instantiates a new ResponseEventsFastroam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEventsFastroamWithDefaults() *ResponseEventsFastroam {
	this := ResponseEventsFastroam{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseEventsFastroam) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsFastroam) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseEventsFastroam) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseEventsFastroam) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsFastroam) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseEventsFastroam) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResponseEventsFastroam) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsFastroam) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseEventsFastroam) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ResponseEventsFastroam) SetNext(v string) {
	o.Next = &v
}

// GetResults returns the Results field value
func (o *ResponseEventsFastroam) GetResults() []EventFastroam {
	if o == nil {
		var ret []EventFastroam
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsFastroam) GetResultsOk() ([]EventFastroam, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseEventsFastroam) SetResults(v []EventFastroam) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseEventsFastroam) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsFastroam) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseEventsFastroam) SetStart(v int32) {
	o.Start = v
}

func (o ResponseEventsFastroam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEventsFastroam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseEventsFastroam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
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

	varResponseEventsFastroam := _ResponseEventsFastroam{}

	err = json.Unmarshal(data, &varResponseEventsFastroam)

	if err != nil {
		return err
	}

	*o = ResponseEventsFastroam(varResponseEventsFastroam)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseEventsFastroam struct {
	value *ResponseEventsFastroam
	isSet bool
}

func (v NullableResponseEventsFastroam) Get() *ResponseEventsFastroam {
	return v.value
}

func (v *NullableResponseEventsFastroam) Set(val *ResponseEventsFastroam) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEventsFastroam) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEventsFastroam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEventsFastroam(val *ResponseEventsFastroam) *NullableResponseEventsFastroam {
	return &NullableResponseEventsFastroam{value: val, isSet: true}
}

func (v NullableResponseEventsFastroam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEventsFastroam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


