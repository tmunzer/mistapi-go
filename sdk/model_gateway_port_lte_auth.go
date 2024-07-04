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

// GatewayPortLteAuth if `wan_type`==`lte`
type GatewayPortLteAuth string

// List of gateway_port_lte_auth
const (
	GATEWAYPORTLTEAUTH_EMPTY GatewayPortLteAuth = ""
	GATEWAYPORTLTEAUTH_NONE GatewayPortLteAuth = "none"
	GATEWAYPORTLTEAUTH_CHAP GatewayPortLteAuth = "chap"
	GATEWAYPORTLTEAUTH_PAP GatewayPortLteAuth = "pap"
)

// All allowed values of GatewayPortLteAuth enum
var AllowedGatewayPortLteAuthEnumValues = []GatewayPortLteAuth{
	"",
	"none",
	"chap",
	"pap",
}

func (v *GatewayPortLteAuth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortLteAuth(value)
	for _, existing := range AllowedGatewayPortLteAuthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortLteAuth", value)
}

// NewGatewayPortLteAuthFromValue returns a pointer to a valid GatewayPortLteAuth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortLteAuthFromValue(v string) (*GatewayPortLteAuth, error) {
	ev := GatewayPortLteAuth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortLteAuth: valid values are %v", v, AllowedGatewayPortLteAuthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortLteAuth) IsValid() bool {
	for _, existing := range AllowedGatewayPortLteAuthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_lte_auth value
func (v GatewayPortLteAuth) Ptr() *GatewayPortLteAuth {
	return &v
}

type NullableGatewayPortLteAuth struct {
	value *GatewayPortLteAuth
	isSet bool
}

func (v NullableGatewayPortLteAuth) Get() *GatewayPortLteAuth {
	return v.value
}

func (v *NullableGatewayPortLteAuth) Set(val *GatewayPortLteAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortLteAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortLteAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortLteAuth(val *GatewayPortLteAuth) *NullableGatewayPortLteAuth {
	return &NullableGatewayPortLteAuth{value: val, isSet: true}
}

func (v NullableGatewayPortLteAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortLteAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

