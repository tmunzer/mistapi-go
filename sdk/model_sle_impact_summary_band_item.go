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

// checks if the SleImpactSummaryBandItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactSummaryBandItem{}

// SleImpactSummaryBandItem struct for SleImpactSummaryBandItem
type SleImpactSummaryBandItem struct {
	Band string `json:"band"`
	Degraded float32 `json:"degraded"`
	Duration float32 `json:"duration"`
	Name string `json:"name"`
	Total float32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactSummaryBandItem SleImpactSummaryBandItem

// NewSleImpactSummaryBandItem instantiates a new SleImpactSummaryBandItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactSummaryBandItem(band string, degraded float32, duration float32, name string, total float32) *SleImpactSummaryBandItem {
	this := SleImpactSummaryBandItem{}
	this.Band = band
	this.Degraded = degraded
	this.Duration = duration
	this.Name = name
	this.Total = total
	return &this
}

// NewSleImpactSummaryBandItemWithDefaults instantiates a new SleImpactSummaryBandItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactSummaryBandItemWithDefaults() *SleImpactSummaryBandItem {
	this := SleImpactSummaryBandItem{}
	return &this
}

// GetBand returns the Band field value
func (o *SleImpactSummaryBandItem) GetBand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Band
}

// GetBandOk returns a tuple with the Band field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryBandItem) GetBandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band, true
}

// SetBand sets field value
func (o *SleImpactSummaryBandItem) SetBand(v string) {
	o.Band = v
}

// GetDegraded returns the Degraded field value
func (o *SleImpactSummaryBandItem) GetDegraded() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryBandItem) GetDegradedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Degraded, true
}

// SetDegraded sets field value
func (o *SleImpactSummaryBandItem) SetDegraded(v float32) {
	o.Degraded = v
}

// GetDuration returns the Duration field value
func (o *SleImpactSummaryBandItem) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryBandItem) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SleImpactSummaryBandItem) SetDuration(v float32) {
	o.Duration = v
}

// GetName returns the Name field value
func (o *SleImpactSummaryBandItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryBandItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SleImpactSummaryBandItem) SetName(v string) {
	o.Name = v
}

// GetTotal returns the Total field value
func (o *SleImpactSummaryBandItem) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryBandItem) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SleImpactSummaryBandItem) SetTotal(v float32) {
	o.Total = v
}

func (o SleImpactSummaryBandItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactSummaryBandItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["band"] = o.Band
	toSerialize["degraded"] = o.Degraded
	toSerialize["duration"] = o.Duration
	toSerialize["name"] = o.Name
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactSummaryBandItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"band",
		"degraded",
		"duration",
		"name",
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

	varSleImpactSummaryBandItem := _SleImpactSummaryBandItem{}

	err = json.Unmarshal(data, &varSleImpactSummaryBandItem)

	if err != nil {
		return err
	}

	*o = SleImpactSummaryBandItem(varSleImpactSummaryBandItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band")
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactSummaryBandItem struct {
	value *SleImpactSummaryBandItem
	isSet bool
}

func (v NullableSleImpactSummaryBandItem) Get() *SleImpactSummaryBandItem {
	return v.value
}

func (v *NullableSleImpactSummaryBandItem) Set(val *SleImpactSummaryBandItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactSummaryBandItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactSummaryBandItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactSummaryBandItem(val *SleImpactSummaryBandItem) *NullableSleImpactSummaryBandItem {
	return &NullableSleImpactSummaryBandItem{value: val, isSet: true}
}

func (v NullableSleImpactSummaryBandItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactSummaryBandItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

