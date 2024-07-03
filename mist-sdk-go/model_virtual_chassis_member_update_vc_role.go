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

// VirtualChassisMemberUpdateVcRole Required if `op`==`add` or `op`==`preprovision`
type VirtualChassisMemberUpdateVcRole string

// List of virtual_chassis_member_update_vc_role
const (
	VIRTUALCHASSISMEMBERUPDATEVCROLE_EMPTY VirtualChassisMemberUpdateVcRole = ""
	VIRTUALCHASSISMEMBERUPDATEVCROLE_MASTER VirtualChassisMemberUpdateVcRole = "master"
	VIRTUALCHASSISMEMBERUPDATEVCROLE_BACKUP VirtualChassisMemberUpdateVcRole = "backup"
	VIRTUALCHASSISMEMBERUPDATEVCROLE_LINECARD VirtualChassisMemberUpdateVcRole = "linecard"
)

// All allowed values of VirtualChassisMemberUpdateVcRole enum
var AllowedVirtualChassisMemberUpdateVcRoleEnumValues = []VirtualChassisMemberUpdateVcRole{
	"",
	"master",
	"backup",
	"linecard",
}

func (v *VirtualChassisMemberUpdateVcRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualChassisMemberUpdateVcRole(value)
	for _, existing := range AllowedVirtualChassisMemberUpdateVcRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualChassisMemberUpdateVcRole", value)
}

// NewVirtualChassisMemberUpdateVcRoleFromValue returns a pointer to a valid VirtualChassisMemberUpdateVcRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualChassisMemberUpdateVcRoleFromValue(v string) (*VirtualChassisMemberUpdateVcRole, error) {
	ev := VirtualChassisMemberUpdateVcRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualChassisMemberUpdateVcRole: valid values are %v", v, AllowedVirtualChassisMemberUpdateVcRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualChassisMemberUpdateVcRole) IsValid() bool {
	for _, existing := range AllowedVirtualChassisMemberUpdateVcRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to virtual_chassis_member_update_vc_role value
func (v VirtualChassisMemberUpdateVcRole) Ptr() *VirtualChassisMemberUpdateVcRole {
	return &v
}

type NullableVirtualChassisMemberUpdateVcRole struct {
	value *VirtualChassisMemberUpdateVcRole
	isSet bool
}

func (v NullableVirtualChassisMemberUpdateVcRole) Get() *VirtualChassisMemberUpdateVcRole {
	return v.value
}

func (v *NullableVirtualChassisMemberUpdateVcRole) Set(val *VirtualChassisMemberUpdateVcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualChassisMemberUpdateVcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualChassisMemberUpdateVcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualChassisMemberUpdateVcRole(val *VirtualChassisMemberUpdateVcRole) *NullableVirtualChassisMemberUpdateVcRole {
	return &NullableVirtualChassisMemberUpdateVcRole{value: val, isSet: true}
}

func (v NullableVirtualChassisMemberUpdateVcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualChassisMemberUpdateVcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
