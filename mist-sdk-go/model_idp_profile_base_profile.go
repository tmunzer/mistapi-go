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

// IdpProfileBaseProfile the model 'IdpProfileBaseProfile'
type IdpProfileBaseProfile string

// List of idp_profile_base_profile
const (
	IDPPROFILEBASEPROFILE_EMPTY IdpProfileBaseProfile = ""
	IDPPROFILEBASEPROFILE_STRICT IdpProfileBaseProfile = "strict"
	IDPPROFILEBASEPROFILE_STANDARD IdpProfileBaseProfile = "standard"
)

// All allowed values of IdpProfileBaseProfile enum
var AllowedIdpProfileBaseProfileEnumValues = []IdpProfileBaseProfile{
	"",
	"strict",
	"standard",
}

func (v *IdpProfileBaseProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdpProfileBaseProfile(value)
	for _, existing := range AllowedIdpProfileBaseProfileEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdpProfileBaseProfile", value)
}

// NewIdpProfileBaseProfileFromValue returns a pointer to a valid IdpProfileBaseProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdpProfileBaseProfileFromValue(v string) (*IdpProfileBaseProfile, error) {
	ev := IdpProfileBaseProfile(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdpProfileBaseProfile: valid values are %v", v, AllowedIdpProfileBaseProfileEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdpProfileBaseProfile) IsValid() bool {
	for _, existing := range AllowedIdpProfileBaseProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to idp_profile_base_profile value
func (v IdpProfileBaseProfile) Ptr() *IdpProfileBaseProfile {
	return &v
}

type NullableIdpProfileBaseProfile struct {
	value *IdpProfileBaseProfile
	isSet bool
}

func (v NullableIdpProfileBaseProfile) Get() *IdpProfileBaseProfile {
	return v.value
}

func (v *NullableIdpProfileBaseProfile) Set(val *IdpProfileBaseProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpProfileBaseProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpProfileBaseProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpProfileBaseProfile(val *IdpProfileBaseProfile) *NullableIdpProfileBaseProfile {
	return &NullableIdpProfileBaseProfile{value: val, isSet: true}
}

func (v NullableIdpProfileBaseProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpProfileBaseProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

