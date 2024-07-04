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

// WlanRoamMode the model 'WlanRoamMode'
type WlanRoamMode string

// List of wlan_roam_mode
const (
	WLANROAMMODE_EMPTY WlanRoamMode = ""
	WLANROAMMODE_NONE WlanRoamMode = "none"
	WLANROAMMODE_OKC WlanRoamMode = "OKC"
	WLANROAMMODE__11R WlanRoamMode = "11r"
)

// All allowed values of WlanRoamMode enum
var AllowedWlanRoamModeEnumValues = []WlanRoamMode{
	"",
	"none",
	"OKC",
	"11r",
}

func (v *WlanRoamMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanRoamMode(value)
	for _, existing := range AllowedWlanRoamModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanRoamMode", value)
}

// NewWlanRoamModeFromValue returns a pointer to a valid WlanRoamMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanRoamModeFromValue(v string) (*WlanRoamMode, error) {
	ev := WlanRoamMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanRoamMode: valid values are %v", v, AllowedWlanRoamModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanRoamMode) IsValid() bool {
	for _, existing := range AllowedWlanRoamModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_roam_mode value
func (v WlanRoamMode) Ptr() *WlanRoamMode {
	return &v
}

type NullableWlanRoamMode struct {
	value *WlanRoamMode
	isSet bool
}

func (v NullableWlanRoamMode) Get() *WlanRoamMode {
	return v.value
}

func (v *NullableWlanRoamMode) Set(val *WlanRoamMode) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanRoamMode) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanRoamMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanRoamMode(val *WlanRoamMode) *NullableWlanRoamMode {
	return &NullableWlanRoamMode{value: val, isSet: true}
}

func (v NullableWlanRoamMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanRoamMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

