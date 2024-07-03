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

// SiteDiscoveredSwitchesCountDistinct the model 'SiteDiscoveredSwitchesCountDistinct'
type SiteDiscoveredSwitchesCountDistinct string

// List of site_discovered_switches_count_distinct
const (
	SITEDISCOVEREDSWITCHESCOUNTDISTINCT_EMPTY SiteDiscoveredSwitchesCountDistinct = ""
	SITEDISCOVEREDSWITCHESCOUNTDISTINCT_SYSTEM_NAME SiteDiscoveredSwitchesCountDistinct = "system_name"
	SITEDISCOVEREDSWITCHESCOUNTDISTINCT_VERSION SiteDiscoveredSwitchesCountDistinct = "version"
	SITEDISCOVEREDSWITCHESCOUNTDISTINCT_MODEL SiteDiscoveredSwitchesCountDistinct = "model"
	SITEDISCOVEREDSWITCHESCOUNTDISTINCT_MGMT_ADDR SiteDiscoveredSwitchesCountDistinct = "mgmt_addr"
)

// All allowed values of SiteDiscoveredSwitchesCountDistinct enum
var AllowedSiteDiscoveredSwitchesCountDistinctEnumValues = []SiteDiscoveredSwitchesCountDistinct{
	"",
	"system_name",
	"version",
	"model",
	"mgmt_addr",
}

func (v *SiteDiscoveredSwitchesCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteDiscoveredSwitchesCountDistinct(value)
	for _, existing := range AllowedSiteDiscoveredSwitchesCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteDiscoveredSwitchesCountDistinct", value)
}

// NewSiteDiscoveredSwitchesCountDistinctFromValue returns a pointer to a valid SiteDiscoveredSwitchesCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteDiscoveredSwitchesCountDistinctFromValue(v string) (*SiteDiscoveredSwitchesCountDistinct, error) {
	ev := SiteDiscoveredSwitchesCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteDiscoveredSwitchesCountDistinct: valid values are %v", v, AllowedSiteDiscoveredSwitchesCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteDiscoveredSwitchesCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteDiscoveredSwitchesCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_discovered_switches_count_distinct value
func (v SiteDiscoveredSwitchesCountDistinct) Ptr() *SiteDiscoveredSwitchesCountDistinct {
	return &v
}

type NullableSiteDiscoveredSwitchesCountDistinct struct {
	value *SiteDiscoveredSwitchesCountDistinct
	isSet bool
}

func (v NullableSiteDiscoveredSwitchesCountDistinct) Get() *SiteDiscoveredSwitchesCountDistinct {
	return v.value
}

func (v *NullableSiteDiscoveredSwitchesCountDistinct) Set(val *SiteDiscoveredSwitchesCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteDiscoveredSwitchesCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteDiscoveredSwitchesCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteDiscoveredSwitchesCountDistinct(val *SiteDiscoveredSwitchesCountDistinct) *NullableSiteDiscoveredSwitchesCountDistinct {
	return &NullableSiteDiscoveredSwitchesCountDistinct{value: val, isSet: true}
}

func (v NullableSiteDiscoveredSwitchesCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteDiscoveredSwitchesCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
