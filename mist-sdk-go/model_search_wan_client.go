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

// checks if the SearchWanClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchWanClient{}

// SearchWanClient struct for SearchWanClient
type SearchWanClient struct {
	End *int32 `json:"end,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Results *ClientsWanStats `json:"results,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchWanClient SearchWanClient

// NewSearchWanClient instantiates a new SearchWanClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchWanClient() *SearchWanClient {
	this := SearchWanClient{}
	return &this
}

// NewSearchWanClientWithDefaults instantiates a new SearchWanClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWanClientWithDefaults() *SearchWanClient {
	this := SearchWanClient{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SearchWanClient) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWanClient) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SearchWanClient) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *SearchWanClient) SetEnd(v int32) {
	o.End = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SearchWanClient) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWanClient) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SearchWanClient) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SearchWanClient) SetLimit(v int32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SearchWanClient) GetResults() ClientsWanStats {
	if o == nil || IsNil(o.Results) {
		var ret ClientsWanStats
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWanClient) GetResultsOk() (*ClientsWanStats, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SearchWanClient) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given ClientsWanStats and assigns it to the Results field.
func (o *SearchWanClient) SetResults(v ClientsWanStats) {
	o.Results = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SearchWanClient) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWanClient) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SearchWanClient) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *SearchWanClient) SetStart(v int32) {
	o.Start = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SearchWanClient) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchWanClient) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SearchWanClient) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SearchWanClient) SetTotal(v int32) {
	o.Total = &v
}

func (o SearchWanClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchWanClient) ToMap() (map[string]interface{}, error) {
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

func (o *SearchWanClient) UnmarshalJSON(data []byte) (err error) {
	varSearchWanClient := _SearchWanClient{}

	err = json.Unmarshal(data, &varSearchWanClient)

	if err != nil {
		return err
	}

	*o = SearchWanClient(varSearchWanClient)

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

type NullableSearchWanClient struct {
	value *SearchWanClient
	isSet bool
}

func (v NullableSearchWanClient) Get() *SearchWanClient {
	return v.value
}

func (v *NullableSearchWanClient) Set(val *SearchWanClient) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchWanClient) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchWanClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchWanClient(val *SearchWanClient) *NullableSearchWanClient {
	return &NullableSearchWanClient{value: val, isSet: true}
}

func (v NullableSearchWanClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchWanClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


