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

// Dot11Proto the model 'Dot11Proto'
type Dot11Proto string

// List of dot11_proto
const (
	DOT11PROTO_EMPTY Dot11Proto = ""
	DOT11PROTO_A Dot11Proto = "a"
	DOT11PROTO_B Dot11Proto = "b"
	DOT11PROTO_G Dot11Proto = "g"
	DOT11PROTO_N Dot11Proto = "n"
	DOT11PROTO_AC Dot11Proto = "ac"
	DOT11PROTO_AX Dot11Proto = "ax"
)

// All allowed values of Dot11Proto enum
var AllowedDot11ProtoEnumValues = []Dot11Proto{
	"",
	"a",
	"b",
	"g",
	"n",
	"ac",
	"ax",
}

func (v *Dot11Proto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Dot11Proto(value)
	for _, existing := range AllowedDot11ProtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Dot11Proto", value)
}

// NewDot11ProtoFromValue returns a pointer to a valid Dot11Proto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDot11ProtoFromValue(v string) (*Dot11Proto, error) {
	ev := Dot11Proto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Dot11Proto: valid values are %v", v, AllowedDot11ProtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Dot11Proto) IsValid() bool {
	for _, existing := range AllowedDot11ProtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dot11_proto value
func (v Dot11Proto) Ptr() *Dot11Proto {
	return &v
}

type NullableDot11Proto struct {
	value *Dot11Proto
	isSet bool
}

func (v NullableDot11Proto) Get() *Dot11Proto {
	return v.value
}

func (v *NullableDot11Proto) Set(val *Dot11Proto) {
	v.value = val
	v.isSet = true
}

func (v NullableDot11Proto) IsSet() bool {
	return v.isSet
}

func (v *NullableDot11Proto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDot11Proto(val *Dot11Proto) *NullableDot11Proto {
	return &NullableDot11Proto{value: val, isSet: true}
}

func (v NullableDot11Proto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDot11Proto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
