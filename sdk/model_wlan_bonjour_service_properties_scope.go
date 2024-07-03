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

// WlanBonjourServicePropertiesScope how bonjour services should be discovered for the same WLAN
type WlanBonjourServicePropertiesScope string

// List of wlan_bonjour_service_properties_scope
const (
	WLANBONJOURSERVICEPROPERTIESSCOPE_EMPTY WlanBonjourServicePropertiesScope = ""
	WLANBONJOURSERVICEPROPERTIESSCOPE_SAME_SITE WlanBonjourServicePropertiesScope = "same_site"
	WLANBONJOURSERVICEPROPERTIESSCOPE_SAME_MAP WlanBonjourServicePropertiesScope = "same_map"
	WLANBONJOURSERVICEPROPERTIESSCOPE_SAME_AP WlanBonjourServicePropertiesScope = "same_ap"
)

// All allowed values of WlanBonjourServicePropertiesScope enum
var AllowedWlanBonjourServicePropertiesScopeEnumValues = []WlanBonjourServicePropertiesScope{
	"",
	"same_site",
	"same_map",
	"same_ap",
}

func (v *WlanBonjourServicePropertiesScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanBonjourServicePropertiesScope(value)
	for _, existing := range AllowedWlanBonjourServicePropertiesScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanBonjourServicePropertiesScope", value)
}

// NewWlanBonjourServicePropertiesScopeFromValue returns a pointer to a valid WlanBonjourServicePropertiesScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanBonjourServicePropertiesScopeFromValue(v string) (*WlanBonjourServicePropertiesScope, error) {
	ev := WlanBonjourServicePropertiesScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanBonjourServicePropertiesScope: valid values are %v", v, AllowedWlanBonjourServicePropertiesScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanBonjourServicePropertiesScope) IsValid() bool {
	for _, existing := range AllowedWlanBonjourServicePropertiesScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_bonjour_service_properties_scope value
func (v WlanBonjourServicePropertiesScope) Ptr() *WlanBonjourServicePropertiesScope {
	return &v
}

type NullableWlanBonjourServicePropertiesScope struct {
	value *WlanBonjourServicePropertiesScope
	isSet bool
}

func (v NullableWlanBonjourServicePropertiesScope) Get() *WlanBonjourServicePropertiesScope {
	return v.value
}

func (v *NullableWlanBonjourServicePropertiesScope) Set(val *WlanBonjourServicePropertiesScope) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanBonjourServicePropertiesScope) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanBonjourServicePropertiesScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanBonjourServicePropertiesScope(val *WlanBonjourServicePropertiesScope) *NullableWlanBonjourServicePropertiesScope {
	return &NullableWlanBonjourServicePropertiesScope{value: val, isSet: true}
}

func (v NullableWlanBonjourServicePropertiesScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanBonjourServicePropertiesScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

