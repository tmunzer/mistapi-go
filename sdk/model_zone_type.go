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

// ZoneType the model 'ZoneType'
type ZoneType string

// List of zone_type
const (
	ZONETYPE_EMPTY ZoneType = ""
	ZONETYPE_ZONES ZoneType = "zones"
	ZONETYPE_RSSIZONES ZoneType = "rssizones"
)

// All allowed values of ZoneType enum
var AllowedZoneTypeEnumValues = []ZoneType{
	"",
	"zones",
	"rssizones",
}

func (v *ZoneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZoneType(value)
	for _, existing := range AllowedZoneTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZoneType", value)
}

// NewZoneTypeFromValue returns a pointer to a valid ZoneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneTypeFromValue(v string) (*ZoneType, error) {
	ev := ZoneType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZoneType: valid values are %v", v, AllowedZoneTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneType) IsValid() bool {
	for _, existing := range AllowedZoneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to zone_type value
func (v ZoneType) Ptr() *ZoneType {
	return &v
}

type NullableZoneType struct {
	value *ZoneType
	isSet bool
}

func (v NullableZoneType) Get() *ZoneType {
	return v.value
}

func (v *NullableZoneType) Set(val *ZoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneType(val *ZoneType) *NullableZoneType {
	return &NullableZoneType{value: val, isSet: true}
}

func (v NullableZoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
