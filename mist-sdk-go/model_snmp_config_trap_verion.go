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

// SnmpConfigTrapVerion the model 'SnmpConfigTrapVerion'
type SnmpConfigTrapVerion string

// List of snmp_config_trap_verion
const (
	SNMPCONFIGTRAPVERION_EMPTY SnmpConfigTrapVerion = ""
	SNMPCONFIGTRAPVERION_V1 SnmpConfigTrapVerion = "v1"
	SNMPCONFIGTRAPVERION_V2 SnmpConfigTrapVerion = "v2"
	SNMPCONFIGTRAPVERION_ALL SnmpConfigTrapVerion = "all"
)

// All allowed values of SnmpConfigTrapVerion enum
var AllowedSnmpConfigTrapVerionEnumValues = []SnmpConfigTrapVerion{
	"",
	"v1",
	"v2",
	"all",
}

func (v *SnmpConfigTrapVerion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SnmpConfigTrapVerion(value)
	for _, existing := range AllowedSnmpConfigTrapVerionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SnmpConfigTrapVerion", value)
}

// NewSnmpConfigTrapVerionFromValue returns a pointer to a valid SnmpConfigTrapVerion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnmpConfigTrapVerionFromValue(v string) (*SnmpConfigTrapVerion, error) {
	ev := SnmpConfigTrapVerion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnmpConfigTrapVerion: valid values are %v", v, AllowedSnmpConfigTrapVerionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnmpConfigTrapVerion) IsValid() bool {
	for _, existing := range AllowedSnmpConfigTrapVerionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to snmp_config_trap_verion value
func (v SnmpConfigTrapVerion) Ptr() *SnmpConfigTrapVerion {
	return &v
}

type NullableSnmpConfigTrapVerion struct {
	value *SnmpConfigTrapVerion
	isSet bool
}

func (v NullableSnmpConfigTrapVerion) Get() *SnmpConfigTrapVerion {
	return v.value
}

func (v *NullableSnmpConfigTrapVerion) Set(val *SnmpConfigTrapVerion) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpConfigTrapVerion) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpConfigTrapVerion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpConfigTrapVerion(val *SnmpConfigTrapVerion) *NullableSnmpConfigTrapVerion {
	return &NullableSnmpConfigTrapVerion{value: val, isSet: true}
}

func (v NullableSnmpConfigTrapVerion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpConfigTrapVerion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

