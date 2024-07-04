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

// SnmpUsmpUserAuthenticationType sha224, sha256, sha384, sha512 are supported in 21.1 and newer release
type SnmpUsmpUserAuthenticationType string

// List of snmp_usmp_user_authentication_type
const (
	SNMPUSMPUSERAUTHENTICATIONTYPE_EMPTY SnmpUsmpUserAuthenticationType = ""
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_MD5 SnmpUsmpUserAuthenticationType = "authentication_md5"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_SHA SnmpUsmpUserAuthenticationType = "authentication_sha"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_SHA224 SnmpUsmpUserAuthenticationType = "authentication_sha224"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_SHA256 SnmpUsmpUserAuthenticationType = "authentication_sha256"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_SHA384 SnmpUsmpUserAuthenticationType = "authentication_sha384"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_SHA512 SnmpUsmpUserAuthenticationType = "authentication_sha512"
	SNMPUSMPUSERAUTHENTICATIONTYPE_AUTHENTICATION_NONE SnmpUsmpUserAuthenticationType = "authentication_none"
)

// All allowed values of SnmpUsmpUserAuthenticationType enum
var AllowedSnmpUsmpUserAuthenticationTypeEnumValues = []SnmpUsmpUserAuthenticationType{
	"",
	"authentication_md5",
	"authentication_sha",
	"authentication_sha224",
	"authentication_sha256",
	"authentication_sha384",
	"authentication_sha512",
	"authentication_none",
}

func (v *SnmpUsmpUserAuthenticationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SnmpUsmpUserAuthenticationType(value)
	for _, existing := range AllowedSnmpUsmpUserAuthenticationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SnmpUsmpUserAuthenticationType", value)
}

// NewSnmpUsmpUserAuthenticationTypeFromValue returns a pointer to a valid SnmpUsmpUserAuthenticationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnmpUsmpUserAuthenticationTypeFromValue(v string) (*SnmpUsmpUserAuthenticationType, error) {
	ev := SnmpUsmpUserAuthenticationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnmpUsmpUserAuthenticationType: valid values are %v", v, AllowedSnmpUsmpUserAuthenticationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnmpUsmpUserAuthenticationType) IsValid() bool {
	for _, existing := range AllowedSnmpUsmpUserAuthenticationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to snmp_usmp_user_authentication_type value
func (v SnmpUsmpUserAuthenticationType) Ptr() *SnmpUsmpUserAuthenticationType {
	return &v
}

type NullableSnmpUsmpUserAuthenticationType struct {
	value *SnmpUsmpUserAuthenticationType
	isSet bool
}

func (v NullableSnmpUsmpUserAuthenticationType) Get() *SnmpUsmpUserAuthenticationType {
	return v.value
}

func (v *NullableSnmpUsmpUserAuthenticationType) Set(val *SnmpUsmpUserAuthenticationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUsmpUserAuthenticationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUsmpUserAuthenticationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUsmpUserAuthenticationType(val *SnmpUsmpUserAuthenticationType) *NullableSnmpUsmpUserAuthenticationType {
	return &NullableSnmpUsmpUserAuthenticationType{value: val, isSet: true}
}

func (v NullableSnmpUsmpUserAuthenticationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUsmpUserAuthenticationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

