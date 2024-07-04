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

// GatewayTemplateProbeType the model 'GatewayTemplateProbeType'
type GatewayTemplateProbeType string

// List of gateway_template_probe_type
const (
	GATEWAYTEMPLATEPROBETYPE_EMPTY GatewayTemplateProbeType = ""
	GATEWAYTEMPLATEPROBETYPE_ICMP GatewayTemplateProbeType = "icmp"
	GATEWAYTEMPLATEPROBETYPE_HTTP GatewayTemplateProbeType = "http"
)

// All allowed values of GatewayTemplateProbeType enum
var AllowedGatewayTemplateProbeTypeEnumValues = []GatewayTemplateProbeType{
	"",
	"icmp",
	"http",
}

func (v *GatewayTemplateProbeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayTemplateProbeType(value)
	for _, existing := range AllowedGatewayTemplateProbeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayTemplateProbeType", value)
}

// NewGatewayTemplateProbeTypeFromValue returns a pointer to a valid GatewayTemplateProbeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTemplateProbeTypeFromValue(v string) (*GatewayTemplateProbeType, error) {
	ev := GatewayTemplateProbeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayTemplateProbeType: valid values are %v", v, AllowedGatewayTemplateProbeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayTemplateProbeType) IsValid() bool {
	for _, existing := range AllowedGatewayTemplateProbeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_template_probe_type value
func (v GatewayTemplateProbeType) Ptr() *GatewayTemplateProbeType {
	return &v
}

type NullableGatewayTemplateProbeType struct {
	value *GatewayTemplateProbeType
	isSet bool
}

func (v NullableGatewayTemplateProbeType) Get() *GatewayTemplateProbeType {
	return v.value
}

func (v *NullableGatewayTemplateProbeType) Set(val *GatewayTemplateProbeType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateProbeType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateProbeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateProbeType(val *GatewayTemplateProbeType) *NullableGatewayTemplateProbeType {
	return &NullableGatewayTemplateProbeType{value: val, isSet: true}
}

func (v NullableGatewayTemplateProbeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateProbeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

