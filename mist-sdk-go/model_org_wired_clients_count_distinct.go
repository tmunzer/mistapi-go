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

// OrgWiredClientsCountDistinct the model 'OrgWiredClientsCountDistinct'
type OrgWiredClientsCountDistinct string

// List of org_wired_clients_count_distinct
const (
	ORGWIREDCLIENTSCOUNTDISTINCT_EMPTY OrgWiredClientsCountDistinct = ""
	ORGWIREDCLIENTSCOUNTDISTINCT_PORT_ID OrgWiredClientsCountDistinct = "port_id"
	ORGWIREDCLIENTSCOUNTDISTINCT_VLAN OrgWiredClientsCountDistinct = "vlan"
	ORGWIREDCLIENTSCOUNTDISTINCT_MAC OrgWiredClientsCountDistinct = "mac"
	ORGWIREDCLIENTSCOUNTDISTINCT_DEVICE_MAC OrgWiredClientsCountDistinct = "device_mac"
	ORGWIREDCLIENTSCOUNTDISTINCT_SITE_ID OrgWiredClientsCountDistinct = "site_id"
	ORGWIREDCLIENTSCOUNTDISTINCT_TYPE OrgWiredClientsCountDistinct = "type"
)

// All allowed values of OrgWiredClientsCountDistinct enum
var AllowedOrgWiredClientsCountDistinctEnumValues = []OrgWiredClientsCountDistinct{
	"",
	"port_id",
	"vlan",
	"mac",
	"device_mac",
	"site_id",
	"type",
}

func (v *OrgWiredClientsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgWiredClientsCountDistinct(value)
	for _, existing := range AllowedOrgWiredClientsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgWiredClientsCountDistinct", value)
}

// NewOrgWiredClientsCountDistinctFromValue returns a pointer to a valid OrgWiredClientsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgWiredClientsCountDistinctFromValue(v string) (*OrgWiredClientsCountDistinct, error) {
	ev := OrgWiredClientsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgWiredClientsCountDistinct: valid values are %v", v, AllowedOrgWiredClientsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgWiredClientsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgWiredClientsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_wired_clients_count_distinct value
func (v OrgWiredClientsCountDistinct) Ptr() *OrgWiredClientsCountDistinct {
	return &v
}

type NullableOrgWiredClientsCountDistinct struct {
	value *OrgWiredClientsCountDistinct
	isSet bool
}

func (v NullableOrgWiredClientsCountDistinct) Get() *OrgWiredClientsCountDistinct {
	return v.value
}

func (v *NullableOrgWiredClientsCountDistinct) Set(val *OrgWiredClientsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgWiredClientsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgWiredClientsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgWiredClientsCountDistinct(val *OrgWiredClientsCountDistinct) *NullableOrgWiredClientsCountDistinct {
	return &NullableOrgWiredClientsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgWiredClientsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgWiredClientsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
