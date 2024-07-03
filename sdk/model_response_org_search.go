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

// checks if the ResponseOrgSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseOrgSearch{}

// ResponseOrgSearch struct for ResponseOrgSearch
type ResponseOrgSearch struct {
	End float32 `json:"end"`
	Limit int32 `json:"limit"`
	Next *string `json:"next,omitempty"`
	Results []ResponseOrgSearchItem `json:"results"`
	Start float32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _ResponseOrgSearch ResponseOrgSearch

// NewResponseOrgSearch instantiates a new ResponseOrgSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseOrgSearch(end float32, limit int32, results []ResponseOrgSearchItem, start float32, total int32) *ResponseOrgSearch {
	this := ResponseOrgSearch{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewResponseOrgSearchWithDefaults instantiates a new ResponseOrgSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseOrgSearchWithDefaults() *ResponseOrgSearch {
	this := ResponseOrgSearch{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseOrgSearch) GetEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseOrgSearch) SetEnd(v float32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseOrgSearch) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseOrgSearch) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResponseOrgSearch) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseOrgSearch) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ResponseOrgSearch) SetNext(v string) {
	o.Next = &v
}

// GetResults returns the Results field value
func (o *ResponseOrgSearch) GetResults() []ResponseOrgSearchItem {
	if o == nil {
		var ret []ResponseOrgSearchItem
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetResultsOk() ([]ResponseOrgSearchItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseOrgSearch) SetResults(v []ResponseOrgSearchItem) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseOrgSearch) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseOrgSearch) SetStart(v float32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *ResponseOrgSearch) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgSearch) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseOrgSearch) SetTotal(v int32) {
	o.Total = v
}

func (o ResponseOrgSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseOrgSearch) ToMap() (map[string]interface{}, error) {
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

func (o *ResponseOrgSearch) UnmarshalJSON(data []byte) (err error) {
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

	varResponseOrgSearch := _ResponseOrgSearch{}

	err = json.Unmarshal(data, &varResponseOrgSearch)

	if err != nil {
		return err
	}

	*o = ResponseOrgSearch(varResponseOrgSearch)

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

type NullableResponseOrgSearch struct {
	value *ResponseOrgSearch
	isSet bool
}

func (v NullableResponseOrgSearch) Get() *ResponseOrgSearch {
	return v.value
}

func (v *NullableResponseOrgSearch) Set(val *ResponseOrgSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseOrgSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseOrgSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseOrgSearch(val *ResponseOrgSearch) *NullableResponseOrgSearch {
	return &NullableResponseOrgSearch{value: val, isSet: true}
}

func (v NullableResponseOrgSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseOrgSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


