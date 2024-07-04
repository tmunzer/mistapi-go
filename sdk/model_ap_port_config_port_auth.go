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

// ApPortConfigPortAuth When doing port auth
type ApPortConfigPortAuth string

// List of ap_port_config_port_auth
const (
	APPORTCONFIGPORTAUTH_EMPTY ApPortConfigPortAuth = ""
	APPORTCONFIGPORTAUTH_NONE ApPortConfigPortAuth = "none"
	APPORTCONFIGPORTAUTH_DOT1X ApPortConfigPortAuth = "dot1x"
)

// All allowed values of ApPortConfigPortAuth enum
var AllowedApPortConfigPortAuthEnumValues = []ApPortConfigPortAuth{
	"",
	"none",
	"dot1x",
}

func (v *ApPortConfigPortAuth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApPortConfigPortAuth(value)
	for _, existing := range AllowedApPortConfigPortAuthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApPortConfigPortAuth", value)
}

// NewApPortConfigPortAuthFromValue returns a pointer to a valid ApPortConfigPortAuth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApPortConfigPortAuthFromValue(v string) (*ApPortConfigPortAuth, error) {
	ev := ApPortConfigPortAuth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApPortConfigPortAuth: valid values are %v", v, AllowedApPortConfigPortAuthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApPortConfigPortAuth) IsValid() bool {
	for _, existing := range AllowedApPortConfigPortAuthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ap_port_config_port_auth value
func (v ApPortConfigPortAuth) Ptr() *ApPortConfigPortAuth {
	return &v
}

type NullableApPortConfigPortAuth struct {
	value *ApPortConfigPortAuth
	isSet bool
}

func (v NullableApPortConfigPortAuth) Get() *ApPortConfigPortAuth {
	return v.value
}

func (v *NullableApPortConfigPortAuth) Set(val *ApPortConfigPortAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableApPortConfigPortAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableApPortConfigPortAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApPortConfigPortAuth(val *ApPortConfigPortAuth) *NullableApPortConfigPortAuth {
	return &NullableApPortConfigPortAuth{value: val, isSet: true}
}

func (v NullableApPortConfigPortAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApPortConfigPortAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

