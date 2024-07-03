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

// SitePortsCountDistinct the model 'SitePortsCountDistinct'
type SitePortsCountDistinct string

// List of site_ports_count_distinct
const (
	SITEPORTSCOUNTDISTINCT_EMPTY SitePortsCountDistinct = ""
	SITEPORTSCOUNTDISTINCT_PORT_ID SitePortsCountDistinct = "port_id"
	SITEPORTSCOUNTDISTINCT_PORT_MAC SitePortsCountDistinct = "port_mac"
	SITEPORTSCOUNTDISTINCT_FULL_DUPLEX SitePortsCountDistinct = "full_duplex"
	SITEPORTSCOUNTDISTINCT_MAC SitePortsCountDistinct = "mac"
	SITEPORTSCOUNTDISTINCT_NEIGHBOR_MAC SitePortsCountDistinct = "neighbor_mac"
	SITEPORTSCOUNTDISTINCT_NEIGHBOR_PORT_DESC SitePortsCountDistinct = "neighbor_port_desc"
	SITEPORTSCOUNTDISTINCT_NEIGHBOR_SYSTEM_NAME SitePortsCountDistinct = "neighbor_system_name"
	SITEPORTSCOUNTDISTINCT_POE_DISABLED SitePortsCountDistinct = "poe_disabled"
	SITEPORTSCOUNTDISTINCT_POE_MODE SitePortsCountDistinct = "poe_mode"
	SITEPORTSCOUNTDISTINCT_POE_ON SitePortsCountDistinct = "poe_on"
	SITEPORTSCOUNTDISTINCT_SPEED SitePortsCountDistinct = "speed"
	SITEPORTSCOUNTDISTINCT_UP SitePortsCountDistinct = "up"
)

// All allowed values of SitePortsCountDistinct enum
var AllowedSitePortsCountDistinctEnumValues = []SitePortsCountDistinct{
	"",
	"port_id",
	"port_mac",
	"full_duplex",
	"mac",
	"neighbor_mac",
	"neighbor_port_desc",
	"neighbor_system_name",
	"poe_disabled",
	"poe_mode",
	"poe_on",
	"speed",
	"up",
}

func (v *SitePortsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SitePortsCountDistinct(value)
	for _, existing := range AllowedSitePortsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SitePortsCountDistinct", value)
}

// NewSitePortsCountDistinctFromValue returns a pointer to a valid SitePortsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSitePortsCountDistinctFromValue(v string) (*SitePortsCountDistinct, error) {
	ev := SitePortsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SitePortsCountDistinct: valid values are %v", v, AllowedSitePortsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SitePortsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSitePortsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_ports_count_distinct value
func (v SitePortsCountDistinct) Ptr() *SitePortsCountDistinct {
	return &v
}

type NullableSitePortsCountDistinct struct {
	value *SitePortsCountDistinct
	isSet bool
}

func (v NullableSitePortsCountDistinct) Get() *SitePortsCountDistinct {
	return v.value
}

func (v *NullableSitePortsCountDistinct) Set(val *SitePortsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSitePortsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSitePortsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSitePortsCountDistinct(val *SitePortsCountDistinct) *NullableSitePortsCountDistinct {
	return &NullableSitePortsCountDistinct{value: val, isSet: true}
}

func (v NullableSitePortsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSitePortsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

