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

// checks if the CallTroubleshootSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallTroubleshootSummary{}

// CallTroubleshootSummary struct for CallTroubleshootSummary
type CallTroubleshootSummary struct {
	Mac *string `json:"mac,omitempty"`
	MeetingId *string `json:"meeting_id,omitempty"`
	Results []CallTroubleshootSummar `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallTroubleshootSummary CallTroubleshootSummary

// NewCallTroubleshootSummary instantiates a new CallTroubleshootSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallTroubleshootSummary() *CallTroubleshootSummary {
	this := CallTroubleshootSummary{}
	return &this
}

// NewCallTroubleshootSummaryWithDefaults instantiates a new CallTroubleshootSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallTroubleshootSummaryWithDefaults() *CallTroubleshootSummary {
	this := CallTroubleshootSummary{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *CallTroubleshootSummary) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummary) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *CallTroubleshootSummary) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *CallTroubleshootSummary) SetMac(v string) {
	o.Mac = &v
}

// GetMeetingId returns the MeetingId field value if set, zero value otherwise.
func (o *CallTroubleshootSummary) GetMeetingId() string {
	if o == nil || IsNil(o.MeetingId) {
		var ret string
		return ret
	}
	return *o.MeetingId
}

// GetMeetingIdOk returns a tuple with the MeetingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummary) GetMeetingIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeetingId) {
		return nil, false
	}
	return o.MeetingId, true
}

// HasMeetingId returns a boolean if a field has been set.
func (o *CallTroubleshootSummary) HasMeetingId() bool {
	if o != nil && !IsNil(o.MeetingId) {
		return true
	}

	return false
}

// SetMeetingId gets a reference to the given string and assigns it to the MeetingId field.
func (o *CallTroubleshootSummary) SetMeetingId(v string) {
	o.MeetingId = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CallTroubleshootSummary) GetResults() []CallTroubleshootSummar {
	if o == nil || IsNil(o.Results) {
		var ret []CallTroubleshootSummar
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummary) GetResultsOk() ([]CallTroubleshootSummar, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CallTroubleshootSummary) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CallTroubleshootSummar and assigns it to the Results field.
func (o *CallTroubleshootSummary) SetResults(v []CallTroubleshootSummar) {
	o.Results = v
}

func (o CallTroubleshootSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallTroubleshootSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.MeetingId) {
		toSerialize["meeting_id"] = o.MeetingId
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallTroubleshootSummary) UnmarshalJSON(data []byte) (err error) {
	varCallTroubleshootSummary := _CallTroubleshootSummary{}

	err = json.Unmarshal(data, &varCallTroubleshootSummary)

	if err != nil {
		return err
	}

	*o = CallTroubleshootSummary(varCallTroubleshootSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "meeting_id")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallTroubleshootSummary struct {
	value *CallTroubleshootSummary
	isSet bool
}

func (v NullableCallTroubleshootSummary) Get() *CallTroubleshootSummary {
	return v.value
}

func (v *NullableCallTroubleshootSummary) Set(val *CallTroubleshootSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCallTroubleshootSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCallTroubleshootSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallTroubleshootSummary(val *CallTroubleshootSummary) *NullableCallTroubleshootSummary {
	return &NullableCallTroubleshootSummary{value: val, isSet: true}
}

func (v NullableCallTroubleshootSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallTroubleshootSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


