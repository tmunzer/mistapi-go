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

// SiteNacClientEventsCountDistinct the model 'SiteNacClientEventsCountDistinct'
type SiteNacClientEventsCountDistinct string

// List of site_nac_client_events_count_distinct
const (
	SITENACCLIENTEVENTSCOUNTDISTINCT_EMPTY SiteNacClientEventsCountDistinct = ""
	SITENACCLIENTEVENTSCOUNTDISTINCT_TYPE SiteNacClientEventsCountDistinct = "type"
	SITENACCLIENTEVENTSCOUNTDISTINCT_NACRULE_ID SiteNacClientEventsCountDistinct = "nacrule_id"
	SITENACCLIENTEVENTSCOUNTDISTINCT_DRYRUN_NACRULE_ID SiteNacClientEventsCountDistinct = "dryrun_nacrule_id"
	SITENACCLIENTEVENTSCOUNTDISTINCT_AUTH_TYPE SiteNacClientEventsCountDistinct = "auth_type"
	SITENACCLIENTEVENTSCOUNTDISTINCT_VLAN SiteNacClientEventsCountDistinct = "vlan"
	SITENACCLIENTEVENTSCOUNTDISTINCT_NAS_VENDOR SiteNacClientEventsCountDistinct = "nas_vendor"
	SITENACCLIENTEVENTSCOUNTDISTINCT_USERNAME SiteNacClientEventsCountDistinct = "username"
	SITENACCLIENTEVENTSCOUNTDISTINCT_AP SiteNacClientEventsCountDistinct = "ap"
	SITENACCLIENTEVENTSCOUNTDISTINCT_MAC SiteNacClientEventsCountDistinct = "mac"
	SITENACCLIENTEVENTSCOUNTDISTINCT_SSID SiteNacClientEventsCountDistinct = "ssid"
)

// All allowed values of SiteNacClientEventsCountDistinct enum
var AllowedSiteNacClientEventsCountDistinctEnumValues = []SiteNacClientEventsCountDistinct{
	"",
	"type",
	"nacrule_id",
	"dryrun_nacrule_id",
	"auth_type",
	"vlan",
	"nas_vendor",
	"username",
	"ap",
	"mac",
	"ssid",
}

func (v *SiteNacClientEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteNacClientEventsCountDistinct(value)
	for _, existing := range AllowedSiteNacClientEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteNacClientEventsCountDistinct", value)
}

// NewSiteNacClientEventsCountDistinctFromValue returns a pointer to a valid SiteNacClientEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteNacClientEventsCountDistinctFromValue(v string) (*SiteNacClientEventsCountDistinct, error) {
	ev := SiteNacClientEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteNacClientEventsCountDistinct: valid values are %v", v, AllowedSiteNacClientEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteNacClientEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteNacClientEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_nac_client_events_count_distinct value
func (v SiteNacClientEventsCountDistinct) Ptr() *SiteNacClientEventsCountDistinct {
	return &v
}

type NullableSiteNacClientEventsCountDistinct struct {
	value *SiteNacClientEventsCountDistinct
	isSet bool
}

func (v NullableSiteNacClientEventsCountDistinct) Get() *SiteNacClientEventsCountDistinct {
	return v.value
}

func (v *NullableSiteNacClientEventsCountDistinct) Set(val *SiteNacClientEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteNacClientEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteNacClientEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteNacClientEventsCountDistinct(val *SiteNacClientEventsCountDistinct) *NullableSiteNacClientEventsCountDistinct {
	return &NullableSiteNacClientEventsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteNacClientEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteNacClientEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
