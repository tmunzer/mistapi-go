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

// IpConfigType the model 'IpConfigType'
type IpConfigType string

// List of ip_config_type
const (
	IPCONFIGTYPE_EMPTY IpConfigType = ""
	IPCONFIGTYPE_STATIC IpConfigType = "static"
	IPCONFIGTYPE_DYNAMIC IpConfigType = "dynamic"
)

// All allowed values of IpConfigType enum
var AllowedIpConfigTypeEnumValues = []IpConfigType{
	"",
	"static",
	"dynamic",
}

func (v *IpConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpConfigType(value)
	for _, existing := range AllowedIpConfigTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpConfigType", value)
}

// NewIpConfigTypeFromValue returns a pointer to a valid IpConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpConfigTypeFromValue(v string) (*IpConfigType, error) {
	ev := IpConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpConfigType: valid values are %v", v, AllowedIpConfigTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpConfigType) IsValid() bool {
	for _, existing := range AllowedIpConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ip_config_type value
func (v IpConfigType) Ptr() *IpConfigType {
	return &v
}

type NullableIpConfigType struct {
	value *IpConfigType
	isSet bool
}

func (v NullableIpConfigType) Get() *IpConfigType {
	return v.value
}

func (v *NullableIpConfigType) Set(val *IpConfigType) {
	v.value = val
	v.isSet = true
}

func (v NullableIpConfigType) IsSet() bool {
	return v.isSet
}

func (v *NullableIpConfigType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpConfigType(val *IpConfigType) *NullableIpConfigType {
	return &NullableIpConfigType{value: val, isSet: true}
}

func (v NullableIpConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpConfigType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
