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

// SwitchPortStatsType device type
type SwitchPortStatsType string

// List of switch_port_stats_type
const (
	SWITCHPORTSTATSTYPE_EMPTY SwitchPortStatsType = ""
	SWITCHPORTSTATSTYPE_AP SwitchPortStatsType = "ap"
	SWITCHPORTSTATSTYPE_BLE SwitchPortStatsType = "ble"
	SWITCHPORTSTATSTYPE_SWITCH SwitchPortStatsType = "switch"
	SWITCHPORTSTATSTYPE_GATEWAY SwitchPortStatsType = "gateway"
	SWITCHPORTSTATSTYPE_MXEDGE SwitchPortStatsType = "mxedge"
	SWITCHPORTSTATSTYPE_NAC SwitchPortStatsType = "nac"
)

// All allowed values of SwitchPortStatsType enum
var AllowedSwitchPortStatsTypeEnumValues = []SwitchPortStatsType{
	"",
	"ap",
	"ble",
	"switch",
	"gateway",
	"mxedge",
	"nac",
}

func (v *SwitchPortStatsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchPortStatsType(value)
	for _, existing := range AllowedSwitchPortStatsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchPortStatsType", value)
}

// NewSwitchPortStatsTypeFromValue returns a pointer to a valid SwitchPortStatsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchPortStatsTypeFromValue(v string) (*SwitchPortStatsType, error) {
	ev := SwitchPortStatsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchPortStatsType: valid values are %v", v, AllowedSwitchPortStatsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchPortStatsType) IsValid() bool {
	for _, existing := range AllowedSwitchPortStatsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_port_stats_type value
func (v SwitchPortStatsType) Ptr() *SwitchPortStatsType {
	return &v
}

type NullableSwitchPortStatsType struct {
	value *SwitchPortStatsType
	isSet bool
}

func (v NullableSwitchPortStatsType) Get() *SwitchPortStatsType {
	return v.value
}

func (v *NullableSwitchPortStatsType) Set(val *SwitchPortStatsType) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchPortStatsType) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchPortStatsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchPortStatsType(val *SwitchPortStatsType) *NullableSwitchPortStatsType {
	return &NullableSwitchPortStatsType{value: val, isSet: true}
}

func (v NullableSwitchPortStatsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchPortStatsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

