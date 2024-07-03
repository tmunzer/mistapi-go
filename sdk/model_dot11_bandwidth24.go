/*
Mist API

> Version: **2406.1.10** > > Date: **July 1, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.11
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// Dot11Bandwidth24 channel width for the 2.4GHz band
type Dot11Bandwidth24 int32

// List of dot11_bandwidth24
const (
	DOT11BANDWIDTH24__20 Dot11Bandwidth24 = 20
	DOT11BANDWIDTH24__40 Dot11Bandwidth24 = 40
)

// All allowed values of Dot11Bandwidth24 enum
var AllowedDot11Bandwidth24EnumValues = []Dot11Bandwidth24{
	20,
	40,
}

func (v *Dot11Bandwidth24) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Dot11Bandwidth24(value)
	for _, existing := range AllowedDot11Bandwidth24EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Dot11Bandwidth24", value)
}

// NewDot11Bandwidth24FromValue returns a pointer to a valid Dot11Bandwidth24
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDot11Bandwidth24FromValue(v int32) (*Dot11Bandwidth24, error) {
	ev := Dot11Bandwidth24(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Dot11Bandwidth24: valid values are %v", v, AllowedDot11Bandwidth24EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Dot11Bandwidth24) IsValid() bool {
	for _, existing := range AllowedDot11Bandwidth24EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dot11_bandwidth24 value
func (v Dot11Bandwidth24) Ptr() *Dot11Bandwidth24 {
	return &v
}

type NullableDot11Bandwidth24 struct {
	value *Dot11Bandwidth24
	isSet bool
}

func (v NullableDot11Bandwidth24) Get() *Dot11Bandwidth24 {
	return v.value
}

func (v *NullableDot11Bandwidth24) Set(val *Dot11Bandwidth24) {
	v.value = val
	v.isSet = true
}

func (v NullableDot11Bandwidth24) IsSet() bool {
	return v.isSet
}

func (v *NullableDot11Bandwidth24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDot11Bandwidth24(val *Dot11Bandwidth24) *NullableDot11Bandwidth24 {
	return &NullableDot11Bandwidth24{value: val, isSet: true}
}

func (v NullableDot11Bandwidth24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDot11Bandwidth24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

