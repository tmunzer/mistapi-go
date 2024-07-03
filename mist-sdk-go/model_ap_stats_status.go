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

// ApStatsStatus the model 'ApStatsStatus'
type ApStatsStatus string

// List of ap_stats_status
const (
	APSTATSSTATUS_EMPTY ApStatsStatus = ""
	APSTATSSTATUS_CONNECTED ApStatsStatus = "connected"
	APSTATSSTATUS_DISCONNECTED ApStatsStatus = "disconnected"
	APSTATSSTATUS_RESTARTING ApStatsStatus = "restarting"
	APSTATSSTATUS_UPGRADING ApStatsStatus = "upgrading"
)

// All allowed values of ApStatsStatus enum
var AllowedApStatsStatusEnumValues = []ApStatsStatus{
	"",
	"connected",
	"disconnected",
	"restarting",
	"upgrading",
}

func (v *ApStatsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApStatsStatus(value)
	for _, existing := range AllowedApStatsStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApStatsStatus", value)
}

// NewApStatsStatusFromValue returns a pointer to a valid ApStatsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApStatsStatusFromValue(v string) (*ApStatsStatus, error) {
	ev := ApStatsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApStatsStatus: valid values are %v", v, AllowedApStatsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApStatsStatus) IsValid() bool {
	for _, existing := range AllowedApStatsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ap_stats_status value
func (v ApStatsStatus) Ptr() *ApStatsStatus {
	return &v
}

type NullableApStatsStatus struct {
	value *ApStatsStatus
	isSet bool
}

func (v NullableApStatsStatus) Get() *ApStatsStatus {
	return v.value
}

func (v *NullableApStatsStatus) Set(val *ApStatsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsStatus(val *ApStatsStatus) *NullableApStatsStatus {
	return &NullableApStatsStatus{value: val, isSet: true}
}

func (v NullableApStatsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

