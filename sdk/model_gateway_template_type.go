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

// GatewayTemplateType the model 'GatewayTemplateType'
type GatewayTemplateType string

// List of gateway_template_type
const (
	GATEWAYTEMPLATETYPE_EMPTY GatewayTemplateType = ""
	GATEWAYTEMPLATETYPE_STANDALONE GatewayTemplateType = "standalone"
	GATEWAYTEMPLATETYPE_SPOKE GatewayTemplateType = "spoke"
)

// All allowed values of GatewayTemplateType enum
var AllowedGatewayTemplateTypeEnumValues = []GatewayTemplateType{
	"",
	"standalone",
	"spoke",
}

func (v *GatewayTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayTemplateType(value)
	for _, existing := range AllowedGatewayTemplateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayTemplateType", value)
}

// NewGatewayTemplateTypeFromValue returns a pointer to a valid GatewayTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTemplateTypeFromValue(v string) (*GatewayTemplateType, error) {
	ev := GatewayTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayTemplateType: valid values are %v", v, AllowedGatewayTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayTemplateType) IsValid() bool {
	for _, existing := range AllowedGatewayTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_template_type value
func (v GatewayTemplateType) Ptr() *GatewayTemplateType {
	return &v
}

type NullableGatewayTemplateType struct {
	value *GatewayTemplateType
	isSet bool
}

func (v NullableGatewayTemplateType) Get() *GatewayTemplateType {
	return v.value
}

func (v *NullableGatewayTemplateType) Set(val *GatewayTemplateType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateType(val *GatewayTemplateType) *NullableGatewayTemplateType {
	return &NullableGatewayTemplateType{value: val, isSet: true}
}

func (v NullableGatewayTemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

