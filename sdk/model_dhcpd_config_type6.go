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

// DhcpdConfigType6 DHCP Server (local) or DHCP Relay (relay)
type DhcpdConfigType6 string

// List of dhcpd_config_type6
const (
	DHCPDCONFIGTYPE6_EMPTY DhcpdConfigType6 = ""
	DHCPDCONFIGTYPE6_LOCAL DhcpdConfigType6 = "local"
	DHCPDCONFIGTYPE6_RELAY DhcpdConfigType6 = "relay"
	DHCPDCONFIGTYPE6_NONE DhcpdConfigType6 = "none"
)

// All allowed values of DhcpdConfigType6 enum
var AllowedDhcpdConfigType6EnumValues = []DhcpdConfigType6{
	"",
	"local",
	"relay",
	"none",
}

func (v *DhcpdConfigType6) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DhcpdConfigType6(value)
	for _, existing := range AllowedDhcpdConfigType6EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DhcpdConfigType6", value)
}

// NewDhcpdConfigType6FromValue returns a pointer to a valid DhcpdConfigType6
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDhcpdConfigType6FromValue(v string) (*DhcpdConfigType6, error) {
	ev := DhcpdConfigType6(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DhcpdConfigType6: valid values are %v", v, AllowedDhcpdConfigType6EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DhcpdConfigType6) IsValid() bool {
	for _, existing := range AllowedDhcpdConfigType6EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dhcpd_config_type6 value
func (v DhcpdConfigType6) Ptr() *DhcpdConfigType6 {
	return &v
}

type NullableDhcpdConfigType6 struct {
	value *DhcpdConfigType6
	isSet bool
}

func (v NullableDhcpdConfigType6) Get() *DhcpdConfigType6 {
	return v.value
}

func (v *NullableDhcpdConfigType6) Set(val *DhcpdConfigType6) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpdConfigType6) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpdConfigType6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpdConfigType6(val *DhcpdConfigType6) *NullableDhcpdConfigType6 {
	return &NullableDhcpdConfigType6{value: val, isSet: true}
}

func (v NullableDhcpdConfigType6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpdConfigType6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

