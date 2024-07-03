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

// checks if the ResponsePcapSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePcapSearch{}

// ResponsePcapSearch struct for ResponsePcapSearch
type ResponsePcapSearch struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Next string `json:"next"`
	Results []ResponsePcapSearchItem `json:"results"`
	Start int32 `json:"start"`
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponsePcapSearch ResponsePcapSearch

// NewResponsePcapSearch instantiates a new ResponsePcapSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePcapSearch(end int32, limit int32, next string, results []ResponsePcapSearchItem, start int32) *ResponsePcapSearch {
	this := ResponsePcapSearch{}
	this.End = end
	this.Limit = limit
	this.Next = next
	this.Results = results
	this.Start = start
	return &this
}

// NewResponsePcapSearchWithDefaults instantiates a new ResponsePcapSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePcapSearchWithDefaults() *ResponsePcapSearch {
	this := ResponsePcapSearch{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponsePcapSearch) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponsePcapSearch) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponsePcapSearch) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponsePcapSearch) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value
func (o *ResponsePcapSearch) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *ResponsePcapSearch) SetNext(v string) {
	o.Next = v
}

// GetResults returns the Results field value
func (o *ResponsePcapSearch) GetResults() []ResponsePcapSearchItem {
	if o == nil {
		var ret []ResponsePcapSearchItem
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetResultsOk() ([]ResponsePcapSearchItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponsePcapSearch) SetResults(v []ResponsePcapSearchItem) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponsePcapSearch) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponsePcapSearch) SetStart(v int32) {
	o.Start = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponsePcapSearch) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapSearch) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponsePcapSearch) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponsePcapSearch) SetTotal(v int32) {
	o.Total = &v
}

func (o ResponsePcapSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePcapSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	toSerialize["next"] = o.Next
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponsePcapSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
		"next",
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

	varResponsePcapSearch := _ResponsePcapSearch{}

	err = json.Unmarshal(data, &varResponsePcapSearch)

	if err != nil {
		return err
	}

	*o = ResponsePcapSearch(varResponsePcapSearch)

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

type NullableResponsePcapSearch struct {
	value *ResponsePcapSearch
	isSet bool
}

func (v NullableResponsePcapSearch) Get() *ResponsePcapSearch {
	return v.value
}

func (v *NullableResponsePcapSearch) Set(val *ResponsePcapSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePcapSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePcapSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePcapSearch(val *ResponsePcapSearch) *NullableResponsePcapSearch {
	return &NullableResponsePcapSearch{value: val, isSet: true}
}

func (v NullableResponsePcapSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePcapSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


