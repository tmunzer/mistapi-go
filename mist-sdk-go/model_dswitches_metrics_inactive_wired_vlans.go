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

// checks if the DswitchesMetricsInactiveWiredVlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DswitchesMetricsInactiveWiredVlans{}

// DswitchesMetricsInactiveWiredVlans struct for DswitchesMetricsInactiveWiredVlans
type DswitchesMetricsInactiveWiredVlans struct {
	Details map[string]interface{} `json:"details"`
	Score float32 `json:"score"`
	AdditionalProperties map[string]interface{}
}

type _DswitchesMetricsInactiveWiredVlans DswitchesMetricsInactiveWiredVlans

// NewDswitchesMetricsInactiveWiredVlans instantiates a new DswitchesMetricsInactiveWiredVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDswitchesMetricsInactiveWiredVlans(details map[string]interface{}, score float32) *DswitchesMetricsInactiveWiredVlans {
	this := DswitchesMetricsInactiveWiredVlans{}
	this.Details = details
	this.Score = score
	return &this
}

// NewDswitchesMetricsInactiveWiredVlansWithDefaults instantiates a new DswitchesMetricsInactiveWiredVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDswitchesMetricsInactiveWiredVlansWithDefaults() *DswitchesMetricsInactiveWiredVlans {
	this := DswitchesMetricsInactiveWiredVlans{}
	return &this
}

// GetDetails returns the Details field value
func (o *DswitchesMetricsInactiveWiredVlans) GetDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsInactiveWiredVlans) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *DswitchesMetricsInactiveWiredVlans) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetScore returns the Score field value
func (o *DswitchesMetricsInactiveWiredVlans) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsInactiveWiredVlans) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *DswitchesMetricsInactiveWiredVlans) SetScore(v float32) {
	o.Score = v
}

func (o DswitchesMetricsInactiveWiredVlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DswitchesMetricsInactiveWiredVlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["details"] = o.Details
	toSerialize["score"] = o.Score

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DswitchesMetricsInactiveWiredVlans) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"details",
		"score",
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

	varDswitchesMetricsInactiveWiredVlans := _DswitchesMetricsInactiveWiredVlans{}

	err = json.Unmarshal(data, &varDswitchesMetricsInactiveWiredVlans)

	if err != nil {
		return err
	}

	*o = DswitchesMetricsInactiveWiredVlans(varDswitchesMetricsInactiveWiredVlans)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDswitchesMetricsInactiveWiredVlans struct {
	value *DswitchesMetricsInactiveWiredVlans
	isSet bool
}

func (v NullableDswitchesMetricsInactiveWiredVlans) Get() *DswitchesMetricsInactiveWiredVlans {
	return v.value
}

func (v *NullableDswitchesMetricsInactiveWiredVlans) Set(val *DswitchesMetricsInactiveWiredVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableDswitchesMetricsInactiveWiredVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableDswitchesMetricsInactiveWiredVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDswitchesMetricsInactiveWiredVlans(val *DswitchesMetricsInactiveWiredVlans) *NullableDswitchesMetricsInactiveWiredVlans {
	return &NullableDswitchesMetricsInactiveWiredVlans{value: val, isSet: true}
}

func (v NullableDswitchesMetricsInactiveWiredVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDswitchesMetricsInactiveWiredVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

