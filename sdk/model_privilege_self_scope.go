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

// PrivilegeSelfScope the model 'PrivilegeSelfScope'
type PrivilegeSelfScope string

// List of privilege_self_scope
const (
	PRIVILEGESELFSCOPE_EMPTY PrivilegeSelfScope = ""
	PRIVILEGESELFSCOPE_MSP PrivilegeSelfScope = "msp"
	PRIVILEGESELFSCOPE_ORG PrivilegeSelfScope = "org"
	PRIVILEGESELFSCOPE_ORGGROUP PrivilegeSelfScope = "orggroup"
	PRIVILEGESELFSCOPE_SITE PrivilegeSelfScope = "site"
	PRIVILEGESELFSCOPE_SITEGROUP PrivilegeSelfScope = "sitegroup"
)

// All allowed values of PrivilegeSelfScope enum
var AllowedPrivilegeSelfScopeEnumValues = []PrivilegeSelfScope{
	"",
	"msp",
	"org",
	"orggroup",
	"site",
	"sitegroup",
}

func (v *PrivilegeSelfScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivilegeSelfScope(value)
	for _, existing := range AllowedPrivilegeSelfScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivilegeSelfScope", value)
}

// NewPrivilegeSelfScopeFromValue returns a pointer to a valid PrivilegeSelfScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivilegeSelfScopeFromValue(v string) (*PrivilegeSelfScope, error) {
	ev := PrivilegeSelfScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivilegeSelfScope: valid values are %v", v, AllowedPrivilegeSelfScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivilegeSelfScope) IsValid() bool {
	for _, existing := range AllowedPrivilegeSelfScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to privilege_self_scope value
func (v PrivilegeSelfScope) Ptr() *PrivilegeSelfScope {
	return &v
}

type NullablePrivilegeSelfScope struct {
	value *PrivilegeSelfScope
	isSet bool
}

func (v NullablePrivilegeSelfScope) Get() *PrivilegeSelfScope {
	return v.value
}

func (v *NullablePrivilegeSelfScope) Set(val *PrivilegeSelfScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegeSelfScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegeSelfScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegeSelfScope(val *PrivilegeSelfScope) *NullablePrivilegeSelfScope {
	return &NullablePrivilegeSelfScope{value: val, isSet: true}
}

func (v NullablePrivilegeSelfScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegeSelfScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

