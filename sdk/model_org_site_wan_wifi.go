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

// checks if the OrgSiteWanWifi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSiteWanWifi{}

// OrgSiteWanWifi struct for OrgSiteWanWifi
type OrgSiteWanWifi struct {
	End float32 `json:"end"`
	Interval int32 `json:"interval"`
	Limit int32 `json:"limit"`
	Page int32 `json:"page"`
	Results []OrgSiteSleWanResult `json:"results"`
	Start float32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _OrgSiteWanWifi OrgSiteWanWifi

// NewOrgSiteWanWifi instantiates a new OrgSiteWanWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSiteWanWifi(end float32, interval int32, limit int32, page int32, results []OrgSiteSleWanResult, start float32, total int32) *OrgSiteWanWifi {
	this := OrgSiteWanWifi{}
	this.End = end
	this.Interval = interval
	this.Limit = limit
	this.Page = page
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewOrgSiteWanWifiWithDefaults instantiates a new OrgSiteWanWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSiteWanWifiWithDefaults() *OrgSiteWanWifi {
	this := OrgSiteWanWifi{}
	return &this
}

// GetEnd returns the End field value
func (o *OrgSiteWanWifi) GetEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *OrgSiteWanWifi) SetEnd(v float32) {
	o.End = v
}

// GetInterval returns the Interval field value
func (o *OrgSiteWanWifi) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *OrgSiteWanWifi) SetInterval(v int32) {
	o.Interval = v
}

// GetLimit returns the Limit field value
func (o *OrgSiteWanWifi) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *OrgSiteWanWifi) SetLimit(v int32) {
	o.Limit = v
}

// GetPage returns the Page field value
func (o *OrgSiteWanWifi) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *OrgSiteWanWifi) SetPage(v int32) {
	o.Page = v
}

// GetResults returns the Results field value
func (o *OrgSiteWanWifi) GetResults() []OrgSiteSleWanResult {
	if o == nil {
		var ret []OrgSiteSleWanResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetResultsOk() ([]OrgSiteSleWanResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *OrgSiteWanWifi) SetResults(v []OrgSiteSleWanResult) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *OrgSiteWanWifi) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *OrgSiteWanWifi) SetStart(v float32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *OrgSiteWanWifi) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OrgSiteWanWifi) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OrgSiteWanWifi) SetTotal(v int32) {
	o.Total = v
}

func (o OrgSiteWanWifi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSiteWanWifi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["interval"] = o.Interval
	toSerialize["limit"] = o.Limit
	toSerialize["page"] = o.Page
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSiteWanWifi) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"interval",
		"limit",
		"page",
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

	varOrgSiteWanWifi := _OrgSiteWanWifi{}

	err = json.Unmarshal(data, &varOrgSiteWanWifi)

	if err != nil {
		return err
	}

	*o = OrgSiteWanWifi(varOrgSiteWanWifi)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "page")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSiteWanWifi struct {
	value *OrgSiteWanWifi
	isSet bool
}

func (v NullableOrgSiteWanWifi) Get() *OrgSiteWanWifi {
	return v.value
}

func (v *NullableOrgSiteWanWifi) Set(val *OrgSiteWanWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSiteWanWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSiteWanWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSiteWanWifi(val *OrgSiteWanWifi) *NullableOrgSiteWanWifi {
	return &NullableOrgSiteWanWifi{value: val, isSet: true}
}

func (v NullableOrgSiteWanWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSiteWanWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


