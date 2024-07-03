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

// checks if the ResponseSwitchMetricsVersionCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSwitchMetricsVersionCompliance{}

// ResponseSwitchMetricsVersionCompliance struct for ResponseSwitchMetricsVersionCompliance
type ResponseSwitchMetricsVersionCompliance struct {
	Details *ResponseSwitchMetricsVersionComplianceDetails `json:"details,omitempty"`
	Score *int32 `json:"score,omitempty"`
	TotalSwitchCount *int32 `json:"total_switch_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSwitchMetricsVersionCompliance ResponseSwitchMetricsVersionCompliance

// NewResponseSwitchMetricsVersionCompliance instantiates a new ResponseSwitchMetricsVersionCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSwitchMetricsVersionCompliance() *ResponseSwitchMetricsVersionCompliance {
	this := ResponseSwitchMetricsVersionCompliance{}
	return &this
}

// NewResponseSwitchMetricsVersionComplianceWithDefaults instantiates a new ResponseSwitchMetricsVersionCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSwitchMetricsVersionComplianceWithDefaults() *ResponseSwitchMetricsVersionCompliance {
	this := ResponseSwitchMetricsVersionCompliance{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsVersionCompliance) GetDetails() ResponseSwitchMetricsVersionComplianceDetails {
	if o == nil || IsNil(o.Details) {
		var ret ResponseSwitchMetricsVersionComplianceDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsVersionCompliance) GetDetailsOk() (*ResponseSwitchMetricsVersionComplianceDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsVersionCompliance) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ResponseSwitchMetricsVersionComplianceDetails and assigns it to the Details field.
func (o *ResponseSwitchMetricsVersionCompliance) SetDetails(v ResponseSwitchMetricsVersionComplianceDetails) {
	o.Details = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsVersionCompliance) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsVersionCompliance) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsVersionCompliance) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *ResponseSwitchMetricsVersionCompliance) SetScore(v int32) {
	o.Score = &v
}

// GetTotalSwitchCount returns the TotalSwitchCount field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsVersionCompliance) GetTotalSwitchCount() int32 {
	if o == nil || IsNil(o.TotalSwitchCount) {
		var ret int32
		return ret
	}
	return *o.TotalSwitchCount
}

// GetTotalSwitchCountOk returns a tuple with the TotalSwitchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsVersionCompliance) GetTotalSwitchCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSwitchCount) {
		return nil, false
	}
	return o.TotalSwitchCount, true
}

// HasTotalSwitchCount returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsVersionCompliance) HasTotalSwitchCount() bool {
	if o != nil && !IsNil(o.TotalSwitchCount) {
		return true
	}

	return false
}

// SetTotalSwitchCount gets a reference to the given int32 and assigns it to the TotalSwitchCount field.
func (o *ResponseSwitchMetricsVersionCompliance) SetTotalSwitchCount(v int32) {
	o.TotalSwitchCount = &v
}

func (o ResponseSwitchMetricsVersionCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSwitchMetricsVersionCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.TotalSwitchCount) {
		toSerialize["total_switch_count"] = o.TotalSwitchCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSwitchMetricsVersionCompliance) UnmarshalJSON(data []byte) (err error) {
	varResponseSwitchMetricsVersionCompliance := _ResponseSwitchMetricsVersionCompliance{}

	err = json.Unmarshal(data, &varResponseSwitchMetricsVersionCompliance)

	if err != nil {
		return err
	}

	*o = ResponseSwitchMetricsVersionCompliance(varResponseSwitchMetricsVersionCompliance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")
		delete(additionalProperties, "score")
		delete(additionalProperties, "total_switch_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSwitchMetricsVersionCompliance struct {
	value *ResponseSwitchMetricsVersionCompliance
	isSet bool
}

func (v NullableResponseSwitchMetricsVersionCompliance) Get() *ResponseSwitchMetricsVersionCompliance {
	return v.value
}

func (v *NullableResponseSwitchMetricsVersionCompliance) Set(val *ResponseSwitchMetricsVersionCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSwitchMetricsVersionCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSwitchMetricsVersionCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSwitchMetricsVersionCompliance(val *ResponseSwitchMetricsVersionCompliance) *NullableResponseSwitchMetricsVersionCompliance {
	return &NullableResponseSwitchMetricsVersionCompliance{value: val, isSet: true}
}

func (v NullableResponseSwitchMetricsVersionCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSwitchMetricsVersionCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


