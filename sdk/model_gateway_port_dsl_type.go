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

// GatewayPortDslType if `wan_type`==`lte`
type GatewayPortDslType string

// List of gateway_port_dsl_type
const (
	GATEWAYPORTDSLTYPE_EMPTY GatewayPortDslType = ""
	GATEWAYPORTDSLTYPE_VDSL GatewayPortDslType = "vdsl"
	GATEWAYPORTDSLTYPE_ADSL GatewayPortDslType = "adsl"
)

// All allowed values of GatewayPortDslType enum
var AllowedGatewayPortDslTypeEnumValues = []GatewayPortDslType{
	"",
	"vdsl",
	"adsl",
}

func (v *GatewayPortDslType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortDslType(value)
	for _, existing := range AllowedGatewayPortDslTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortDslType", value)
}

// NewGatewayPortDslTypeFromValue returns a pointer to a valid GatewayPortDslType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortDslTypeFromValue(v string) (*GatewayPortDslType, error) {
	ev := GatewayPortDslType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortDslType: valid values are %v", v, AllowedGatewayPortDslTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortDslType) IsValid() bool {
	for _, existing := range AllowedGatewayPortDslTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_dsl_type value
func (v GatewayPortDslType) Ptr() *GatewayPortDslType {
	return &v
}

type NullableGatewayPortDslType struct {
	value *GatewayPortDslType
	isSet bool
}

func (v NullableGatewayPortDslType) Get() *GatewayPortDslType {
	return v.value
}

func (v *NullableGatewayPortDslType) Set(val *GatewayPortDslType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortDslType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortDslType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortDslType(val *GatewayPortDslType) *NullableGatewayPortDslType {
	return &NullableGatewayPortDslType{value: val, isSet: true}
}

func (v NullableGatewayPortDslType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortDslType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

