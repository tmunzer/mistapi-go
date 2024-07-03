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

// WlanInterface where this WLAN will be connected to
type WlanInterface string

// List of wlan_interface
const (
	WLANINTERFACE_EMPTY WlanInterface = ""
	WLANINTERFACE_ALL WlanInterface = "all"
	WLANINTERFACE_ETH0 WlanInterface = "eth0"
	WLANINTERFACE_ETH1 WlanInterface = "eth1"
	WLANINTERFACE_ETH2 WlanInterface = "eth2"
	WLANINTERFACE_ETH3 WlanInterface = "eth3"
	WLANINTERFACE_WXTUNNEL WlanInterface = "wxtunnel"
	WLANINTERFACE_MXTUNNEL WlanInterface = "mxtunnel"
	WLANINTERFACE_SITE_MXEDGE WlanInterface = "site_mxedge"
)

// All allowed values of WlanInterface enum
var AllowedWlanInterfaceEnumValues = []WlanInterface{
	"",
	"all",
	"eth0",
	"eth1",
	"eth2",
	"eth3",
	"wxtunnel",
	"mxtunnel",
	"site_mxedge",
}

func (v *WlanInterface) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanInterface(value)
	for _, existing := range AllowedWlanInterfaceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanInterface", value)
}

// NewWlanInterfaceFromValue returns a pointer to a valid WlanInterface
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanInterfaceFromValue(v string) (*WlanInterface, error) {
	ev := WlanInterface(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanInterface: valid values are %v", v, AllowedWlanInterfaceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanInterface) IsValid() bool {
	for _, existing := range AllowedWlanInterfaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_interface value
func (v WlanInterface) Ptr() *WlanInterface {
	return &v
}

type NullableWlanInterface struct {
	value *WlanInterface
	isSet bool
}

func (v NullableWlanInterface) Get() *WlanInterface {
	return v.value
}

func (v *NullableWlanInterface) Set(val *WlanInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanInterface(val *WlanInterface) *NullableWlanInterface {
	return &NullableWlanInterface{value: val, isSet: true}
}

func (v NullableWlanInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
