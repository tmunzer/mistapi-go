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
)

// checks if the ResponseTroubleshoot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTroubleshoot{}

// ResponseTroubleshoot struct for ResponseTroubleshoot
type ResponseTroubleshoot struct {
	End *int32 `json:"end,omitempty"`
	Results []ResponseTroubleshootItem `json:"results,omitempty"`
	Start *int32 `json:"start,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseTroubleshoot ResponseTroubleshoot

// NewResponseTroubleshoot instantiates a new ResponseTroubleshoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTroubleshoot() *ResponseTroubleshoot {
	this := ResponseTroubleshoot{}
	return &this
}

// NewResponseTroubleshootWithDefaults instantiates a new ResponseTroubleshoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTroubleshootWithDefaults() *ResponseTroubleshoot {
	this := ResponseTroubleshoot{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ResponseTroubleshoot) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTroubleshoot) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ResponseTroubleshoot) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ResponseTroubleshoot) SetEnd(v int32) {
	o.End = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResponseTroubleshoot) GetResults() []ResponseTroubleshootItem {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseTroubleshootItem
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTroubleshoot) GetResultsOk() ([]ResponseTroubleshootItem, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResponseTroubleshoot) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseTroubleshootItem and assigns it to the Results field.
func (o *ResponseTroubleshoot) SetResults(v []ResponseTroubleshootItem) {
	o.Results = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResponseTroubleshoot) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTroubleshoot) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResponseTroubleshoot) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ResponseTroubleshoot) SetStart(v int32) {
	o.Start = &v
}

func (o ResponseTroubleshoot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTroubleshoot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseTroubleshoot) UnmarshalJSON(data []byte) (err error) {
	varResponseTroubleshoot := _ResponseTroubleshoot{}

	err = json.Unmarshal(data, &varResponseTroubleshoot)

	if err != nil {
		return err
	}

	*o = ResponseTroubleshoot(varResponseTroubleshoot)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseTroubleshoot struct {
	value *ResponseTroubleshoot
	isSet bool
}

func (v NullableResponseTroubleshoot) Get() *ResponseTroubleshoot {
	return v.value
}

func (v *NullableResponseTroubleshoot) Set(val *ResponseTroubleshoot) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTroubleshoot) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTroubleshoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTroubleshoot(val *ResponseTroubleshoot) *NullableResponseTroubleshoot {
	return &NullableResponseTroubleshoot{value: val, isSet: true}
}

func (v NullableResponseTroubleshoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTroubleshoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


