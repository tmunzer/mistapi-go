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

// checks if the AlarmAck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmAck{}

// AlarmAck struct for AlarmAck
type AlarmAck struct {
	AlarmIds []string `json:"alarm_ids"`
	// Some text note describing the intent
	Note *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlarmAck AlarmAck

// NewAlarmAck instantiates a new AlarmAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmAck(alarmIds []string) *AlarmAck {
	this := AlarmAck{}
	this.AlarmIds = alarmIds
	return &this
}

// NewAlarmAckWithDefaults instantiates a new AlarmAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmAckWithDefaults() *AlarmAck {
	this := AlarmAck{}
	return &this
}

// GetAlarmIds returns the AlarmIds field value
func (o *AlarmAck) GetAlarmIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlarmIds
}

// GetAlarmIdsOk returns a tuple with the AlarmIds field value
// and a boolean to check if the value has been set.
func (o *AlarmAck) GetAlarmIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmIds, true
}

// SetAlarmIds sets field value
func (o *AlarmAck) SetAlarmIds(v []string) {
	o.AlarmIds = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *AlarmAck) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmAck) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *AlarmAck) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *AlarmAck) SetNote(v string) {
	o.Note = &v
}

func (o AlarmAck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmAck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alarm_ids"] = o.AlarmIds
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlarmAck) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alarm_ids",
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

	varAlarmAck := _AlarmAck{}

	err = json.Unmarshal(data, &varAlarmAck)

	if err != nil {
		return err
	}

	*o = AlarmAck(varAlarmAck)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alarm_ids")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlarmAck struct {
	value *AlarmAck
	isSet bool
}

func (v NullableAlarmAck) Get() *AlarmAck {
	return v.value
}

func (v *NullableAlarmAck) Set(val *AlarmAck) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmAck) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmAck(val *AlarmAck) *NullableAlarmAck {
	return &NullableAlarmAck{value: val, isSet: true}
}

func (v NullableAlarmAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


