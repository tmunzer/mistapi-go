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

// OrgSwitchPortCountDistinct the model 'OrgSwitchPortCountDistinct'
type OrgSwitchPortCountDistinct string

// List of org_switch_port_count_distinct
const (
	ORGSWITCHPORTCOUNTDISTINCT_EMPTY OrgSwitchPortCountDistinct = ""
	ORGSWITCHPORTCOUNTDISTINCT_PORT_ID OrgSwitchPortCountDistinct = "port_id"
	ORGSWITCHPORTCOUNTDISTINCT_PORT_MAC OrgSwitchPortCountDistinct = "port_mac"
	ORGSWITCHPORTCOUNTDISTINCT_FULL_DUPLEX OrgSwitchPortCountDistinct = "full_duplex"
	ORGSWITCHPORTCOUNTDISTINCT_MAC OrgSwitchPortCountDistinct = "mac"
	ORGSWITCHPORTCOUNTDISTINCT_NEIGHBOR_MAC OrgSwitchPortCountDistinct = "neighbor_mac"
	ORGSWITCHPORTCOUNTDISTINCT_NEIGHBOR_PORT_DESC OrgSwitchPortCountDistinct = "neighbor_port_desc"
	ORGSWITCHPORTCOUNTDISTINCT_NEIGHBOR_SYSTEM_NAME OrgSwitchPortCountDistinct = "neighbor_system_name"
	ORGSWITCHPORTCOUNTDISTINCT_POE_DISABLED OrgSwitchPortCountDistinct = "poe_disabled"
	ORGSWITCHPORTCOUNTDISTINCT_POE_MODE OrgSwitchPortCountDistinct = "poe_mode"
	ORGSWITCHPORTCOUNTDISTINCT_POE_ON OrgSwitchPortCountDistinct = "poe_on"
	ORGSWITCHPORTCOUNTDISTINCT_SPEED OrgSwitchPortCountDistinct = "speed"
	ORGSWITCHPORTCOUNTDISTINCT_UP OrgSwitchPortCountDistinct = "up"
)

// All allowed values of OrgSwitchPortCountDistinct enum
var AllowedOrgSwitchPortCountDistinctEnumValues = []OrgSwitchPortCountDistinct{
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

func (v *OrgSwitchPortCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgSwitchPortCountDistinct(value)
	for _, existing := range AllowedOrgSwitchPortCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgSwitchPortCountDistinct", value)
}

// NewOrgSwitchPortCountDistinctFromValue returns a pointer to a valid OrgSwitchPortCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgSwitchPortCountDistinctFromValue(v string) (*OrgSwitchPortCountDistinct, error) {
	ev := OrgSwitchPortCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgSwitchPortCountDistinct: valid values are %v", v, AllowedOrgSwitchPortCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgSwitchPortCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgSwitchPortCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_switch_port_count_distinct value
func (v OrgSwitchPortCountDistinct) Ptr() *OrgSwitchPortCountDistinct {
	return &v
}

type NullableOrgSwitchPortCountDistinct struct {
	value *OrgSwitchPortCountDistinct
	isSet bool
}

func (v NullableOrgSwitchPortCountDistinct) Get() *OrgSwitchPortCountDistinct {
	return v.value
}

func (v *NullableOrgSwitchPortCountDistinct) Set(val *OrgSwitchPortCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSwitchPortCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSwitchPortCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSwitchPortCountDistinct(val *OrgSwitchPortCountDistinct) *NullableOrgSwitchPortCountDistinct {
	return &NullableOrgSwitchPortCountDistinct{value: val, isSet: true}
}

func (v NullableOrgSwitchPortCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSwitchPortCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

