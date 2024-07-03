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

// PrivilegeSelfRole access permissions
type PrivilegeSelfRole string

// List of privilege_self_role
const (
	PRIVILEGESELFROLE_EMPTY PrivilegeSelfRole = ""
	PRIVILEGESELFROLE_ADMIN PrivilegeSelfRole = "admin"
	PRIVILEGESELFROLE_WRITE PrivilegeSelfRole = "write"
	PRIVILEGESELFROLE_READ PrivilegeSelfRole = "read"
	PRIVILEGESELFROLE_HELPDESK PrivilegeSelfRole = "helpdesk"
	PRIVILEGESELFROLE_INSTALLER PrivilegeSelfRole = "installer"
)

// All allowed values of PrivilegeSelfRole enum
var AllowedPrivilegeSelfRoleEnumValues = []PrivilegeSelfRole{
	"",
	"admin",
	"write",
	"read",
	"helpdesk",
	"installer",
}

func (v *PrivilegeSelfRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivilegeSelfRole(value)
	for _, existing := range AllowedPrivilegeSelfRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivilegeSelfRole", value)
}

// NewPrivilegeSelfRoleFromValue returns a pointer to a valid PrivilegeSelfRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivilegeSelfRoleFromValue(v string) (*PrivilegeSelfRole, error) {
	ev := PrivilegeSelfRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivilegeSelfRole: valid values are %v", v, AllowedPrivilegeSelfRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivilegeSelfRole) IsValid() bool {
	for _, existing := range AllowedPrivilegeSelfRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to privilege_self_role value
func (v PrivilegeSelfRole) Ptr() *PrivilegeSelfRole {
	return &v
}

type NullablePrivilegeSelfRole struct {
	value *PrivilegeSelfRole
	isSet bool
}

func (v NullablePrivilegeSelfRole) Get() *PrivilegeSelfRole {
	return v.value
}

func (v *NullablePrivilegeSelfRole) Set(val *PrivilegeSelfRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegeSelfRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegeSelfRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegeSelfRole(val *PrivilegeSelfRole) *NullablePrivilegeSelfRole {
	return &NullablePrivilegeSelfRole{value: val, isSet: true}
}

func (v NullablePrivilegeSelfRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegeSelfRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

