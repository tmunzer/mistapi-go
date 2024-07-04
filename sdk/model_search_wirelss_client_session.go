/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the SearchWirelssClientSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchWirelssClientSession{}

// SearchWirelssClientSession struct for SearchWirelssClientSession
type SearchWirelssClientSession struct {
	End float32 `json:"end"`
	Limit int32 `json:"limit"`
	Next *string `json:"next,omitempty"`
	Results []WirelssClientSession `json:"results"`
	Start float32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _SearchWirelssClientSession SearchWirelssClientSession

// NewSearchWirelssClientSession instantiates a new SearchWirelssClientSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchWirelssClientSession(end float32, limit int32, results []WirelssClientSession, start float32, total int32) *SearchWirelssClientSession {
	this := SearchWirelssClientSession{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewSearchWirelssClientSessionWithDefaults instantiates a new SearchWirelssClientSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWirelssClientSessionWithDefaults() *SearchWirelssClientSession {
	this := SearchWirelssClientSession{}
	return &this
}

// GetEnd returns the End field value
func (o *SearchWirelssClientSession) GetEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *SearchWirelssClientSession) SetEnd(v float32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *SearchWirelssClientSession) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *SearchWirelssClientSession) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SearchWirelssClientSession) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SearchWirelssClientSession) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *SearchWirelssClientSession) SetNext(v string) {
	o.Next = &v
}

// GetResults returns the Results field value
func (o *SearchWirelssClientSession) GetResults() []WirelssClientSession {
	if o == nil {
		var ret []WirelssClientSession
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetResultsOk() ([]WirelssClientSession, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *SearchWirelssClientSession) SetResults(v []WirelssClientSession) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *SearchWirelssClientSession) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *SearchWirelssClientSession) SetStart(v float32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *SearchWirelssClientSession) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SearchWirelssClientSession) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SearchWirelssClientSession) SetTotal(v int32) {
	o.Total = v
}

func (o SearchWirelssClientSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchWirelssClientSession) ToMap() (map[string]interface{}, error) {
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

func (o *SearchWirelssClientSession) UnmarshalJSON(data []byte) (err error) {
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

	varSearchWirelssClientSession := _SearchWirelssClientSession{}

	err = json.Unmarshal(data, &varSearchWirelssClientSession)

	if err != nil {
		return err
	}

	*o = SearchWirelssClientSession(varSearchWirelssClientSession)

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

type NullableSearchWirelssClientSession struct {
	value *SearchWirelssClientSession
	isSet bool
}

func (v NullableSearchWirelssClientSession) Get() *SearchWirelssClientSession {
	return v.value
}

func (v *NullableSearchWirelssClientSession) Set(val *SearchWirelssClientSession) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchWirelssClientSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchWirelssClientSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchWirelssClientSession(val *SearchWirelssClientSession) *NullableSearchWirelssClientSession {
	return &NullableSearchWirelssClientSession{value: val, isSet: true}
}

func (v NullableSearchWirelssClientSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchWirelssClientSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


