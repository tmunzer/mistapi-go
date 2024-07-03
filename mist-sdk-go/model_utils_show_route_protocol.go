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

// UtilsShowRouteProtocol the model 'UtilsShowRouteProtocol'
type UtilsShowRouteProtocol string

// List of utils_show_route_protocol
const (
	UTILSSHOWROUTEPROTOCOL_EMPTY UtilsShowRouteProtocol = ""
	UTILSSHOWROUTEPROTOCOL_BGP UtilsShowRouteProtocol = "bgp"
	UTILSSHOWROUTEPROTOCOL_ANY UtilsShowRouteProtocol = "any"
	UTILSSHOWROUTEPROTOCOL_OSPF UtilsShowRouteProtocol = "ospf"
	UTILSSHOWROUTEPROTOCOL_STATIC UtilsShowRouteProtocol = "static"
	UTILSSHOWROUTEPROTOCOL_DIRECT UtilsShowRouteProtocol = "direct"
	UTILSSHOWROUTEPROTOCOL_EVPN UtilsShowRouteProtocol = "evpn"
)

// All allowed values of UtilsShowRouteProtocol enum
var AllowedUtilsShowRouteProtocolEnumValues = []UtilsShowRouteProtocol{
	"",
	"bgp",
	"any",
	"ospf",
	"static",
	"direct",
	"evpn",
}

func (v *UtilsShowRouteProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UtilsShowRouteProtocol(value)
	for _, existing := range AllowedUtilsShowRouteProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UtilsShowRouteProtocol", value)
}

// NewUtilsShowRouteProtocolFromValue returns a pointer to a valid UtilsShowRouteProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUtilsShowRouteProtocolFromValue(v string) (*UtilsShowRouteProtocol, error) {
	ev := UtilsShowRouteProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UtilsShowRouteProtocol: valid values are %v", v, AllowedUtilsShowRouteProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UtilsShowRouteProtocol) IsValid() bool {
	for _, existing := range AllowedUtilsShowRouteProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to utils_show_route_protocol value
func (v UtilsShowRouteProtocol) Ptr() *UtilsShowRouteProtocol {
	return &v
}

type NullableUtilsShowRouteProtocol struct {
	value *UtilsShowRouteProtocol
	isSet bool
}

func (v NullableUtilsShowRouteProtocol) Get() *UtilsShowRouteProtocol {
	return v.value
}

func (v *NullableUtilsShowRouteProtocol) Set(val *UtilsShowRouteProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowRouteProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowRouteProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowRouteProtocol(val *UtilsShowRouteProtocol) *NullableUtilsShowRouteProtocol {
	return &NullableUtilsShowRouteProtocol{value: val, isSet: true}
}

func (v NullableUtilsShowRouteProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowRouteProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

