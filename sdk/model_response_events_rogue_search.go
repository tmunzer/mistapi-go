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

// checks if the ResponseEventsRogueSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEventsRogueSearch{}

// ResponseEventsRogueSearch struct for ResponseEventsRogueSearch
type ResponseEventsRogueSearch struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Next *string `json:"next,omitempty"`
	Results []EventsRogue `json:"results"`
	Start int32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _ResponseEventsRogueSearch ResponseEventsRogueSearch

// NewResponseEventsRogueSearch instantiates a new ResponseEventsRogueSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEventsRogueSearch(end int32, limit int32, results []EventsRogue, start int32, total int32) *ResponseEventsRogueSearch {
	this := ResponseEventsRogueSearch{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewResponseEventsRogueSearchWithDefaults instantiates a new ResponseEventsRogueSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEventsRogueSearchWithDefaults() *ResponseEventsRogueSearch {
	this := ResponseEventsRogueSearch{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseEventsRogueSearch) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseEventsRogueSearch) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseEventsRogueSearch) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseEventsRogueSearch) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResponseEventsRogueSearch) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseEventsRogueSearch) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ResponseEventsRogueSearch) SetNext(v string) {
	o.Next = &v
}

// GetResults returns the Results field value
func (o *ResponseEventsRogueSearch) GetResults() []EventsRogue {
	if o == nil {
		var ret []EventsRogue
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetResultsOk() ([]EventsRogue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseEventsRogueSearch) SetResults(v []EventsRogue) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseEventsRogueSearch) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseEventsRogueSearch) SetStart(v int32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *ResponseEventsRogueSearch) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseEventsRogueSearch) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseEventsRogueSearch) SetTotal(v int32) {
	o.Total = v
}

func (o ResponseEventsRogueSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEventsRogueSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseEventsRogueSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
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

	varResponseEventsRogueSearch := _ResponseEventsRogueSearch{}

	err = json.Unmarshal(data, &varResponseEventsRogueSearch)

	if err != nil {
		return err
	}

	*o = ResponseEventsRogueSearch(varResponseEventsRogueSearch)

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

type NullableResponseEventsRogueSearch struct {
	value *ResponseEventsRogueSearch
	isSet bool
}

func (v NullableResponseEventsRogueSearch) Get() *ResponseEventsRogueSearch {
	return v.value
}

func (v *NullableResponseEventsRogueSearch) Set(val *ResponseEventsRogueSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEventsRogueSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEventsRogueSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEventsRogueSearch(val *ResponseEventsRogueSearch) *NullableResponseEventsRogueSearch {
	return &NullableResponseEventsRogueSearch{value: val, isSet: true}
}

func (v NullableResponseEventsRogueSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEventsRogueSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


