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

// WlanAuthPairwiseItem the model 'WlanAuthPairwiseItem'
type WlanAuthPairwiseItem string

// List of wlan_auth_pairwise_item
const (
	WLANAUTHPAIRWISEITEM_EMPTY WlanAuthPairwiseItem = ""
	WLANAUTHPAIRWISEITEM_WPA1_CCMP WlanAuthPairwiseItem = "wpa1-ccmp"
	WLANAUTHPAIRWISEITEM_WPA2_TKIP WlanAuthPairwiseItem = "wpa2-tkip"
	WLANAUTHPAIRWISEITEM_WPA1_TKIP WlanAuthPairwiseItem = "wpa1-tkip"
	WLANAUTHPAIRWISEITEM_WPA2_CCMP WlanAuthPairwiseItem = "wpa2-ccmp"
	WLANAUTHPAIRWISEITEM_WPA3 WlanAuthPairwiseItem = "wpa3"
)

// All allowed values of WlanAuthPairwiseItem enum
var AllowedWlanAuthPairwiseItemEnumValues = []WlanAuthPairwiseItem{
	"",
	"wpa1-ccmp",
	"wpa2-tkip",
	"wpa1-tkip",
	"wpa2-ccmp",
	"wpa3",
}

func (v *WlanAuthPairwiseItem) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanAuthPairwiseItem(value)
	for _, existing := range AllowedWlanAuthPairwiseItemEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanAuthPairwiseItem", value)
}

// NewWlanAuthPairwiseItemFromValue returns a pointer to a valid WlanAuthPairwiseItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanAuthPairwiseItemFromValue(v string) (*WlanAuthPairwiseItem, error) {
	ev := WlanAuthPairwiseItem(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanAuthPairwiseItem: valid values are %v", v, AllowedWlanAuthPairwiseItemEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanAuthPairwiseItem) IsValid() bool {
	for _, existing := range AllowedWlanAuthPairwiseItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_auth_pairwise_item value
func (v WlanAuthPairwiseItem) Ptr() *WlanAuthPairwiseItem {
	return &v
}

type NullableWlanAuthPairwiseItem struct {
	value *WlanAuthPairwiseItem
	isSet bool
}

func (v NullableWlanAuthPairwiseItem) Get() *WlanAuthPairwiseItem {
	return v.value
}

func (v *NullableWlanAuthPairwiseItem) Set(val *WlanAuthPairwiseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAuthPairwiseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAuthPairwiseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAuthPairwiseItem(val *WlanAuthPairwiseItem) *NullableWlanAuthPairwiseItem {
	return &NullableWlanAuthPairwiseItem{value: val, isSet: true}
}

func (v NullableWlanAuthPairwiseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAuthPairwiseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
