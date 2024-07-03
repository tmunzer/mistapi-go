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

// VpnPathBfdProfile the model 'VpnPathBfdProfile'
type VpnPathBfdProfile string

// List of vpn_path_bfd_profile
const (
	VPNPATHBFDPROFILE_EMPTY VpnPathBfdProfile = ""
	VPNPATHBFDPROFILE_BROADBAND VpnPathBfdProfile = "broadband"
	VPNPATHBFDPROFILE_LTE VpnPathBfdProfile = "lte"
)

// All allowed values of VpnPathBfdProfile enum
var AllowedVpnPathBfdProfileEnumValues = []VpnPathBfdProfile{
	"",
	"broadband",
	"lte",
}

func (v *VpnPathBfdProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpnPathBfdProfile(value)
	for _, existing := range AllowedVpnPathBfdProfileEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpnPathBfdProfile", value)
}

// NewVpnPathBfdProfileFromValue returns a pointer to a valid VpnPathBfdProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpnPathBfdProfileFromValue(v string) (*VpnPathBfdProfile, error) {
	ev := VpnPathBfdProfile(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpnPathBfdProfile: valid values are %v", v, AllowedVpnPathBfdProfileEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpnPathBfdProfile) IsValid() bool {
	for _, existing := range AllowedVpnPathBfdProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpn_path_bfd_profile value
func (v VpnPathBfdProfile) Ptr() *VpnPathBfdProfile {
	return &v
}

type NullableVpnPathBfdProfile struct {
	value *VpnPathBfdProfile
	isSet bool
}

func (v NullableVpnPathBfdProfile) Get() *VpnPathBfdProfile {
	return v.value
}

func (v *NullableVpnPathBfdProfile) Set(val *VpnPathBfdProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnPathBfdProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnPathBfdProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnPathBfdProfile(val *VpnPathBfdProfile) *NullableVpnPathBfdProfile {
	return &NullableVpnPathBfdProfile{value: val, isSet: true}
}

func (v NullableVpnPathBfdProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnPathBfdProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

