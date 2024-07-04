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

// DiscoveredSwitchesMetricScope the model 'DiscoveredSwitchesMetricScope'
type DiscoveredSwitchesMetricScope string

// List of discovered_switches_metric_scope
const (
	DISCOVEREDSWITCHESMETRICSCOPE_EMPTY DiscoveredSwitchesMetricScope = ""
	DISCOVEREDSWITCHESMETRICSCOPE_SITE DiscoveredSwitchesMetricScope = "site"
	DISCOVEREDSWITCHESMETRICSCOPE_SWITCH DiscoveredSwitchesMetricScope = "switch"
)

// All allowed values of DiscoveredSwitchesMetricScope enum
var AllowedDiscoveredSwitchesMetricScopeEnumValues = []DiscoveredSwitchesMetricScope{
	"",
	"site",
	"switch",
}

func (v *DiscoveredSwitchesMetricScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiscoveredSwitchesMetricScope(value)
	for _, existing := range AllowedDiscoveredSwitchesMetricScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiscoveredSwitchesMetricScope", value)
}

// NewDiscoveredSwitchesMetricScopeFromValue returns a pointer to a valid DiscoveredSwitchesMetricScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiscoveredSwitchesMetricScopeFromValue(v string) (*DiscoveredSwitchesMetricScope, error) {
	ev := DiscoveredSwitchesMetricScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiscoveredSwitchesMetricScope: valid values are %v", v, AllowedDiscoveredSwitchesMetricScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiscoveredSwitchesMetricScope) IsValid() bool {
	for _, existing := range AllowedDiscoveredSwitchesMetricScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to discovered_switches_metric_scope value
func (v DiscoveredSwitchesMetricScope) Ptr() *DiscoveredSwitchesMetricScope {
	return &v
}

type NullableDiscoveredSwitchesMetricScope struct {
	value *DiscoveredSwitchesMetricScope
	isSet bool
}

func (v NullableDiscoveredSwitchesMetricScope) Get() *DiscoveredSwitchesMetricScope {
	return v.value
}

func (v *NullableDiscoveredSwitchesMetricScope) Set(val *DiscoveredSwitchesMetricScope) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredSwitchesMetricScope) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredSwitchesMetricScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredSwitchesMetricScope(val *DiscoveredSwitchesMetricScope) *NullableDiscoveredSwitchesMetricScope {
	return &NullableDiscoveredSwitchesMetricScope{value: val, isSet: true}
}

func (v NullableDiscoveredSwitchesMetricScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredSwitchesMetricScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
