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

// WlanDynamicVlanType standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco)
type WlanDynamicVlanType string

// List of wlan_dynamic_vlan_type
const (
	WLANDYNAMICVLANTYPE_EMPTY WlanDynamicVlanType = ""
	WLANDYNAMICVLANTYPE_STANDARD WlanDynamicVlanType = "standard"
	WLANDYNAMICVLANTYPE_AIRESPACE_INTERFACE_NAME WlanDynamicVlanType = "airespace-interface-name"
)

// All allowed values of WlanDynamicVlanType enum
var AllowedWlanDynamicVlanTypeEnumValues = []WlanDynamicVlanType{
	"",
	"standard",
	"airespace-interface-name",
}

func (v *WlanDynamicVlanType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanDynamicVlanType(value)
	for _, existing := range AllowedWlanDynamicVlanTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanDynamicVlanType", value)
}

// NewWlanDynamicVlanTypeFromValue returns a pointer to a valid WlanDynamicVlanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanDynamicVlanTypeFromValue(v string) (*WlanDynamicVlanType, error) {
	ev := WlanDynamicVlanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanDynamicVlanType: valid values are %v", v, AllowedWlanDynamicVlanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanDynamicVlanType) IsValid() bool {
	for _, existing := range AllowedWlanDynamicVlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_dynamic_vlan_type value
func (v WlanDynamicVlanType) Ptr() *WlanDynamicVlanType {
	return &v
}

type NullableWlanDynamicVlanType struct {
	value *WlanDynamicVlanType
	isSet bool
}

func (v NullableWlanDynamicVlanType) Get() *WlanDynamicVlanType {
	return v.value
}

func (v *NullableWlanDynamicVlanType) Set(val *WlanDynamicVlanType) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanDynamicVlanType) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanDynamicVlanType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanDynamicVlanType(val *WlanDynamicVlanType) *NullableWlanDynamicVlanType {
	return &NullableWlanDynamicVlanType{value: val, isSet: true}
}

func (v NullableWlanDynamicVlanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanDynamicVlanType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
