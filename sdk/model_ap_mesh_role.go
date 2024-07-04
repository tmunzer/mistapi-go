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

// ApMeshRole the model 'ApMeshRole'
type ApMeshRole string

// List of ap_mesh_role
const (
	APMESHROLE_EMPTY ApMeshRole = ""
	APMESHROLE_BASE ApMeshRole = "base"
	APMESHROLE_REMOTE ApMeshRole = "remote"
)

// All allowed values of ApMeshRole enum
var AllowedApMeshRoleEnumValues = []ApMeshRole{
	"",
	"base",
	"remote",
}

func (v *ApMeshRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApMeshRole(value)
	for _, existing := range AllowedApMeshRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApMeshRole", value)
}

// NewApMeshRoleFromValue returns a pointer to a valid ApMeshRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApMeshRoleFromValue(v string) (*ApMeshRole, error) {
	ev := ApMeshRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApMeshRole: valid values are %v", v, AllowedApMeshRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApMeshRole) IsValid() bool {
	for _, existing := range AllowedApMeshRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ap_mesh_role value
func (v ApMeshRole) Ptr() *ApMeshRole {
	return &v
}

type NullableApMeshRole struct {
	value *ApMeshRole
	isSet bool
}

func (v NullableApMeshRole) Get() *ApMeshRole {
	return v.value
}

func (v *NullableApMeshRole) Set(val *ApMeshRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApMeshRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApMeshRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApMeshRole(val *ApMeshRole) *NullableApMeshRole {
	return &NullableApMeshRole{value: val, isSet: true}
}

func (v NullableApMeshRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApMeshRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

