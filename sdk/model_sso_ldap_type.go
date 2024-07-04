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

// SsoLdapType if `idp_type`==`ldap`
type SsoLdapType string

// List of sso_ldap_type
const (
	SSOLDAPTYPE_EMPTY SsoLdapType = ""
	SSOLDAPTYPE_AZURE SsoLdapType = "azure"
	SSOLDAPTYPE_OKTA SsoLdapType = "okta"
	SSOLDAPTYPE_GOOGLE SsoLdapType = "google"
	SSOLDAPTYPE_CUSTOM SsoLdapType = "custom"
)

// All allowed values of SsoLdapType enum
var AllowedSsoLdapTypeEnumValues = []SsoLdapType{
	"",
	"azure",
	"okta",
	"google",
	"custom",
}

func (v *SsoLdapType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SsoLdapType(value)
	for _, existing := range AllowedSsoLdapTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SsoLdapType", value)
}

// NewSsoLdapTypeFromValue returns a pointer to a valid SsoLdapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSsoLdapTypeFromValue(v string) (*SsoLdapType, error) {
	ev := SsoLdapType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SsoLdapType: valid values are %v", v, AllowedSsoLdapTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SsoLdapType) IsValid() bool {
	for _, existing := range AllowedSsoLdapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sso_ldap_type value
func (v SsoLdapType) Ptr() *SsoLdapType {
	return &v
}

type NullableSsoLdapType struct {
	value *SsoLdapType
	isSet bool
}

func (v NullableSsoLdapType) Get() *SsoLdapType {
	return v.value
}

func (v *NullableSsoLdapType) Set(val *SsoLdapType) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoLdapType) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoLdapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoLdapType(val *SsoLdapType) *NullableSsoLdapType {
	return &NullableSsoLdapType{value: val, isSet: true}
}

func (v NullableSsoLdapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoLdapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
