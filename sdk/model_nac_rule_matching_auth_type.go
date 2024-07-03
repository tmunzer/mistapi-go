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

// NacRuleMatchingAuthType the model 'NacRuleMatchingAuthType'
type NacRuleMatchingAuthType string

// List of nac_rule_matching_auth_type
const (
	NACRULEMATCHINGAUTHTYPE_EMPTY NacRuleMatchingAuthType = ""
	NACRULEMATCHINGAUTHTYPE_CERT NacRuleMatchingAuthType = "cert"
	NACRULEMATCHINGAUTHTYPE_IDP NacRuleMatchingAuthType = "idp"
	NACRULEMATCHINGAUTHTYPE_MAB NacRuleMatchingAuthType = "mab"
	NACRULEMATCHINGAUTHTYPE_PSK NacRuleMatchingAuthType = "psk"
	NACRULEMATCHINGAUTHTYPE_DEVICE_AUTH NacRuleMatchingAuthType = "device-auth"
	NACRULEMATCHINGAUTHTYPE_EAP_TLS NacRuleMatchingAuthType = "eap-tls"
	NACRULEMATCHINGAUTHTYPE_EAP_TTLS NacRuleMatchingAuthType = "eap-ttls"
	NACRULEMATCHINGAUTHTYPE_EAP_TEAP NacRuleMatchingAuthType = "eap-teap"
)

// All allowed values of NacRuleMatchingAuthType enum
var AllowedNacRuleMatchingAuthTypeEnumValues = []NacRuleMatchingAuthType{
	"",
	"cert",
	"idp",
	"mab",
	"psk",
	"device-auth",
	"eap-tls",
	"eap-ttls",
	"eap-teap",
}

func (v *NacRuleMatchingAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NacRuleMatchingAuthType(value)
	for _, existing := range AllowedNacRuleMatchingAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NacRuleMatchingAuthType", value)
}

// NewNacRuleMatchingAuthTypeFromValue returns a pointer to a valid NacRuleMatchingAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNacRuleMatchingAuthTypeFromValue(v string) (*NacRuleMatchingAuthType, error) {
	ev := NacRuleMatchingAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NacRuleMatchingAuthType: valid values are %v", v, AllowedNacRuleMatchingAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NacRuleMatchingAuthType) IsValid() bool {
	for _, existing := range AllowedNacRuleMatchingAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to nac_rule_matching_auth_type value
func (v NacRuleMatchingAuthType) Ptr() *NacRuleMatchingAuthType {
	return &v
}

type NullableNacRuleMatchingAuthType struct {
	value *NacRuleMatchingAuthType
	isSet bool
}

func (v NullableNacRuleMatchingAuthType) Get() *NacRuleMatchingAuthType {
	return v.value
}

func (v *NullableNacRuleMatchingAuthType) Set(val *NacRuleMatchingAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableNacRuleMatchingAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableNacRuleMatchingAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNacRuleMatchingAuthType(val *NacRuleMatchingAuthType) *NullableNacRuleMatchingAuthType {
	return &NullableNacRuleMatchingAuthType{value: val, isSet: true}
}

func (v NullableNacRuleMatchingAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNacRuleMatchingAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

