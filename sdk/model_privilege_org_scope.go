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

// PrivilegeOrgScope the model 'PrivilegeOrgScope'
type PrivilegeOrgScope string

// List of privilege_org_scope
const (
	PRIVILEGEORGSCOPE_EMPTY PrivilegeOrgScope = ""
	PRIVILEGEORGSCOPE_ORG PrivilegeOrgScope = "org"
	PRIVILEGEORGSCOPE_SITE PrivilegeOrgScope = "site"
	PRIVILEGEORGSCOPE_SITEGROUP PrivilegeOrgScope = "sitegroup"
)

// All allowed values of PrivilegeOrgScope enum
var AllowedPrivilegeOrgScopeEnumValues = []PrivilegeOrgScope{
	"",
	"org",
	"site",
	"sitegroup",
}

func (v *PrivilegeOrgScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivilegeOrgScope(value)
	for _, existing := range AllowedPrivilegeOrgScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivilegeOrgScope", value)
}

// NewPrivilegeOrgScopeFromValue returns a pointer to a valid PrivilegeOrgScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivilegeOrgScopeFromValue(v string) (*PrivilegeOrgScope, error) {
	ev := PrivilegeOrgScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivilegeOrgScope: valid values are %v", v, AllowedPrivilegeOrgScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivilegeOrgScope) IsValid() bool {
	for _, existing := range AllowedPrivilegeOrgScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to privilege_org_scope value
func (v PrivilegeOrgScope) Ptr() *PrivilegeOrgScope {
	return &v
}

type NullablePrivilegeOrgScope struct {
	value *PrivilegeOrgScope
	isSet bool
}

func (v NullablePrivilegeOrgScope) Get() *PrivilegeOrgScope {
	return v.value
}

func (v *NullablePrivilegeOrgScope) Set(val *PrivilegeOrgScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegeOrgScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegeOrgScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegeOrgScope(val *PrivilegeOrgScope) *NullablePrivilegeOrgScope {
	return &NullablePrivilegeOrgScope{value: val, isSet: true}
}

func (v NullablePrivilegeOrgScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegeOrgScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

