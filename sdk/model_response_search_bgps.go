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

// checks if the ResponseSearchBgps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSearchBgps{}

// ResponseSearchBgps struct for ResponseSearchBgps
type ResponseSearchBgps struct {
	End *float32 `json:"end,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Results []BgpStats `json:"results,omitempty"`
	Start *float32 `json:"start,omitempty"`
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSearchBgps ResponseSearchBgps

// NewResponseSearchBgps instantiates a new ResponseSearchBgps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSearchBgps() *ResponseSearchBgps {
	this := ResponseSearchBgps{}
	return &this
}

// NewResponseSearchBgpsWithDefaults instantiates a new ResponseSearchBgps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSearchBgpsWithDefaults() *ResponseSearchBgps {
	this := ResponseSearchBgps{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ResponseSearchBgps) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchBgps) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ResponseSearchBgps) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *ResponseSearchBgps) SetEnd(v float32) {
	o.End = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseSearchBgps) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchBgps) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseSearchBgps) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseSearchBgps) SetLimit(v int32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResponseSearchBgps) GetResults() []BgpStats {
	if o == nil || IsNil(o.Results) {
		var ret []BgpStats
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchBgps) GetResultsOk() ([]BgpStats, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResponseSearchBgps) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BgpStats and assigns it to the Results field.
func (o *ResponseSearchBgps) SetResults(v []BgpStats) {
	o.Results = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResponseSearchBgps) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchBgps) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResponseSearchBgps) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *ResponseSearchBgps) SetStart(v float32) {
	o.Start = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseSearchBgps) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchBgps) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseSearchBgps) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponseSearchBgps) SetTotal(v int32) {
	o.Total = &v
}

func (o ResponseSearchBgps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSearchBgps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSearchBgps) UnmarshalJSON(data []byte) (err error) {
	varResponseSearchBgps := _ResponseSearchBgps{}

	err = json.Unmarshal(data, &varResponseSearchBgps)

	if err != nil {
		return err
	}

	*o = ResponseSearchBgps(varResponseSearchBgps)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSearchBgps struct {
	value *ResponseSearchBgps
	isSet bool
}

func (v NullableResponseSearchBgps) Get() *ResponseSearchBgps {
	return v.value
}

func (v *NullableResponseSearchBgps) Set(val *ResponseSearchBgps) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSearchBgps) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSearchBgps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSearchBgps(val *ResponseSearchBgps) *NullableResponseSearchBgps {
	return &NullableResponseSearchBgps{value: val, isSet: true}
}

func (v NullableResponseSearchBgps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSearchBgps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

