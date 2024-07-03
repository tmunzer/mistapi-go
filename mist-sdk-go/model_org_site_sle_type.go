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

// OrgSiteSleType the model 'OrgSiteSleType'
type OrgSiteSleType string

// List of org_site_sle_type
const (
	ORGSITESLETYPE_EMPTY OrgSiteSleType = ""
	ORGSITESLETYPE_WAN OrgSiteSleType = "wan"
	ORGSITESLETYPE_WIRED OrgSiteSleType = "wired"
	ORGSITESLETYPE_WIFI OrgSiteSleType = "wifi"
)

// All allowed values of OrgSiteSleType enum
var AllowedOrgSiteSleTypeEnumValues = []OrgSiteSleType{
	"",
	"wan",
	"wired",
	"wifi",
}

func (v *OrgSiteSleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgSiteSleType(value)
	for _, existing := range AllowedOrgSiteSleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgSiteSleType", value)
}

// NewOrgSiteSleTypeFromValue returns a pointer to a valid OrgSiteSleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgSiteSleTypeFromValue(v string) (*OrgSiteSleType, error) {
	ev := OrgSiteSleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgSiteSleType: valid values are %v", v, AllowedOrgSiteSleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgSiteSleType) IsValid() bool {
	for _, existing := range AllowedOrgSiteSleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_site_sle_type value
func (v OrgSiteSleType) Ptr() *OrgSiteSleType {
	return &v
}

type NullableOrgSiteSleType struct {
	value *OrgSiteSleType
	isSet bool
}

func (v NullableOrgSiteSleType) Get() *OrgSiteSleType {
	return v.value
}

func (v *NullableOrgSiteSleType) Set(val *OrgSiteSleType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSiteSleType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSiteSleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSiteSleType(val *OrgSiteSleType) *NullableOrgSiteSleType {
	return &NullableOrgSiteSleType{value: val, isSet: true}
}

func (v NullableOrgSiteSleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSiteSleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

