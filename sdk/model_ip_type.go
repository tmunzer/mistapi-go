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

// IpType the model 'IpType'
type IpType string

// List of ip_type
const (
	IPTYPE_EMPTY IpType = ""
	IPTYPE_STATIC IpType = "static"
	IPTYPE_DHCP IpType = "dhcp"
)

// All allowed values of IpType enum
var AllowedIpTypeEnumValues = []IpType{
	"",
	"static",
	"dhcp",
}

func (v *IpType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpType(value)
	for _, existing := range AllowedIpTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpType", value)
}

// NewIpTypeFromValue returns a pointer to a valid IpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpTypeFromValue(v string) (*IpType, error) {
	ev := IpType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpType: valid values are %v", v, AllowedIpTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpType) IsValid() bool {
	for _, existing := range AllowedIpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ip_type value
func (v IpType) Ptr() *IpType {
	return &v
}

type NullableIpType struct {
	value *IpType
	isSet bool
}

func (v NullableIpType) Get() *IpType {
	return v.value
}

func (v *NullableIpType) Set(val *IpType) {
	v.value = val
	v.isSet = true
}

func (v NullableIpType) IsSet() bool {
	return v.isSet
}

func (v *NullableIpType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpType(val *IpType) *NullableIpType {
	return &NullableIpType{value: val, isSet: true}
}

func (v NullableIpType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

