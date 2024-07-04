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

// SearchSiteSwOrGwPortsDeviceType the model 'SearchSiteSwOrGwPortsDeviceType'
type SearchSiteSwOrGwPortsDeviceType string

// List of search_site_sw_or_gw_ports_device_type
const (
	SEARCHSITESWORGWPORTSDEVICETYPE_EMPTY SearchSiteSwOrGwPortsDeviceType = ""
	SEARCHSITESWORGWPORTSDEVICETYPE_AP SearchSiteSwOrGwPortsDeviceType = "ap"
	SEARCHSITESWORGWPORTSDEVICETYPE_BLE SearchSiteSwOrGwPortsDeviceType = "ble"
	SEARCHSITESWORGWPORTSDEVICETYPE_SWITCH SearchSiteSwOrGwPortsDeviceType = "switch"
	SEARCHSITESWORGWPORTSDEVICETYPE_GATEWAY SearchSiteSwOrGwPortsDeviceType = "gateway"
	SEARCHSITESWORGWPORTSDEVICETYPE_MXEDGE SearchSiteSwOrGwPortsDeviceType = "mxedge"
	SEARCHSITESWORGWPORTSDEVICETYPE_NAC SearchSiteSwOrGwPortsDeviceType = "nac"
)

// All allowed values of SearchSiteSwOrGwPortsDeviceType enum
var AllowedSearchSiteSwOrGwPortsDeviceTypeEnumValues = []SearchSiteSwOrGwPortsDeviceType{
	"",
	"ap",
	"ble",
	"switch",
	"gateway",
	"mxedge",
	"nac",
}

func (v *SearchSiteSwOrGwPortsDeviceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchSiteSwOrGwPortsDeviceType(value)
	for _, existing := range AllowedSearchSiteSwOrGwPortsDeviceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchSiteSwOrGwPortsDeviceType", value)
}

// NewSearchSiteSwOrGwPortsDeviceTypeFromValue returns a pointer to a valid SearchSiteSwOrGwPortsDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchSiteSwOrGwPortsDeviceTypeFromValue(v string) (*SearchSiteSwOrGwPortsDeviceType, error) {
	ev := SearchSiteSwOrGwPortsDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchSiteSwOrGwPortsDeviceType: valid values are %v", v, AllowedSearchSiteSwOrGwPortsDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchSiteSwOrGwPortsDeviceType) IsValid() bool {
	for _, existing := range AllowedSearchSiteSwOrGwPortsDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to search_site_sw_or_gw_ports_device_type value
func (v SearchSiteSwOrGwPortsDeviceType) Ptr() *SearchSiteSwOrGwPortsDeviceType {
	return &v
}

type NullableSearchSiteSwOrGwPortsDeviceType struct {
	value *SearchSiteSwOrGwPortsDeviceType
	isSet bool
}

func (v NullableSearchSiteSwOrGwPortsDeviceType) Get() *SearchSiteSwOrGwPortsDeviceType {
	return v.value
}

func (v *NullableSearchSiteSwOrGwPortsDeviceType) Set(val *SearchSiteSwOrGwPortsDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSiteSwOrGwPortsDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSiteSwOrGwPortsDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSiteSwOrGwPortsDeviceType(val *SearchSiteSwOrGwPortsDeviceType) *NullableSearchSiteSwOrGwPortsDeviceType {
	return &NullableSearchSiteSwOrGwPortsDeviceType{value: val, isSet: true}
}

func (v NullableSearchSiteSwOrGwPortsDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSiteSwOrGwPortsDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

