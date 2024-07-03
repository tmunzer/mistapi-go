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

// SearchSiteSwOrGwPortsStpState the model 'SearchSiteSwOrGwPortsStpState'
type SearchSiteSwOrGwPortsStpState string

// List of search_site_sw_or_gw_ports_stp_state
const (
	SEARCHSITESWORGWPORTSSTPSTATE_EMPTY SearchSiteSwOrGwPortsStpState = ""
	SEARCHSITESWORGWPORTSSTPSTATE_FORWARDING SearchSiteSwOrGwPortsStpState = "forwarding"
	SEARCHSITESWORGWPORTSSTPSTATE_BLOCKING SearchSiteSwOrGwPortsStpState = "blocking"
	SEARCHSITESWORGWPORTSSTPSTATE_LEARNING SearchSiteSwOrGwPortsStpState = "learning"
	SEARCHSITESWORGWPORTSSTPSTATE_LISTENING SearchSiteSwOrGwPortsStpState = "listening"
	SEARCHSITESWORGWPORTSSTPSTATE_DISABLED SearchSiteSwOrGwPortsStpState = "disabled"
)

// All allowed values of SearchSiteSwOrGwPortsStpState enum
var AllowedSearchSiteSwOrGwPortsStpStateEnumValues = []SearchSiteSwOrGwPortsStpState{
	"",
	"forwarding",
	"blocking",
	"learning",
	"listening",
	"disabled",
}

func (v *SearchSiteSwOrGwPortsStpState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchSiteSwOrGwPortsStpState(value)
	for _, existing := range AllowedSearchSiteSwOrGwPortsStpStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchSiteSwOrGwPortsStpState", value)
}

// NewSearchSiteSwOrGwPortsStpStateFromValue returns a pointer to a valid SearchSiteSwOrGwPortsStpState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchSiteSwOrGwPortsStpStateFromValue(v string) (*SearchSiteSwOrGwPortsStpState, error) {
	ev := SearchSiteSwOrGwPortsStpState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchSiteSwOrGwPortsStpState: valid values are %v", v, AllowedSearchSiteSwOrGwPortsStpStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchSiteSwOrGwPortsStpState) IsValid() bool {
	for _, existing := range AllowedSearchSiteSwOrGwPortsStpStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to search_site_sw_or_gw_ports_stp_state value
func (v SearchSiteSwOrGwPortsStpState) Ptr() *SearchSiteSwOrGwPortsStpState {
	return &v
}

type NullableSearchSiteSwOrGwPortsStpState struct {
	value *SearchSiteSwOrGwPortsStpState
	isSet bool
}

func (v NullableSearchSiteSwOrGwPortsStpState) Get() *SearchSiteSwOrGwPortsStpState {
	return v.value
}

func (v *NullableSearchSiteSwOrGwPortsStpState) Set(val *SearchSiteSwOrGwPortsStpState) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSiteSwOrGwPortsStpState) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSiteSwOrGwPortsStpState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSiteSwOrGwPortsStpState(val *SearchSiteSwOrGwPortsStpState) *NullableSearchSiteSwOrGwPortsStpState {
	return &NullableSearchSiteSwOrGwPortsStpState{value: val, isSet: true}
}

func (v NullableSearchSiteSwOrGwPortsStpState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSiteSwOrGwPortsStpState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

