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

// OrgLogsCountDistinct the model 'OrgLogsCountDistinct'
type OrgLogsCountDistinct string

// List of org_logs_count_distinct
const (
	ORGLOGSCOUNTDISTINCT_EMPTY OrgLogsCountDistinct = ""
	ORGLOGSCOUNTDISTINCT_ADMIN_ID OrgLogsCountDistinct = "admin_id"
	ORGLOGSCOUNTDISTINCT_ADMIN_NAME OrgLogsCountDistinct = "admin_name"
	ORGLOGSCOUNTDISTINCT_MESSAGE OrgLogsCountDistinct = "message"
	ORGLOGSCOUNTDISTINCT_SITE_ID OrgLogsCountDistinct = "site_id"
)

// All allowed values of OrgLogsCountDistinct enum
var AllowedOrgLogsCountDistinctEnumValues = []OrgLogsCountDistinct{
	"",
	"admin_id",
	"admin_name",
	"message",
	"site_id",
}

func (v *OrgLogsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgLogsCountDistinct(value)
	for _, existing := range AllowedOrgLogsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgLogsCountDistinct", value)
}

// NewOrgLogsCountDistinctFromValue returns a pointer to a valid OrgLogsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgLogsCountDistinctFromValue(v string) (*OrgLogsCountDistinct, error) {
	ev := OrgLogsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgLogsCountDistinct: valid values are %v", v, AllowedOrgLogsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgLogsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgLogsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_logs_count_distinct value
func (v OrgLogsCountDistinct) Ptr() *OrgLogsCountDistinct {
	return &v
}

type NullableOrgLogsCountDistinct struct {
	value *OrgLogsCountDistinct
	isSet bool
}

func (v NullableOrgLogsCountDistinct) Get() *OrgLogsCountDistinct {
	return v.value
}

func (v *NullableOrgLogsCountDistinct) Set(val *OrgLogsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgLogsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgLogsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgLogsCountDistinct(val *OrgLogsCountDistinct) *NullableOrgLogsCountDistinct {
	return &NullableOrgLogsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgLogsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgLogsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

