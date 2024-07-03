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

// Snmpv3ConfigNotifyType the model 'Snmpv3ConfigNotifyType'
type Snmpv3ConfigNotifyType string

// List of snmpv3_config_notify_type
const (
	SNMPV3CONFIGNOTIFYTYPE_EMPTY Snmpv3ConfigNotifyType = ""
	SNMPV3CONFIGNOTIFYTYPE_TRAP Snmpv3ConfigNotifyType = "trap"
	SNMPV3CONFIGNOTIFYTYPE_INFORM Snmpv3ConfigNotifyType = "inform"
)

// All allowed values of Snmpv3ConfigNotifyType enum
var AllowedSnmpv3ConfigNotifyTypeEnumValues = []Snmpv3ConfigNotifyType{
	"",
	"trap",
	"inform",
}

func (v *Snmpv3ConfigNotifyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Snmpv3ConfigNotifyType(value)
	for _, existing := range AllowedSnmpv3ConfigNotifyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Snmpv3ConfigNotifyType", value)
}

// NewSnmpv3ConfigNotifyTypeFromValue returns a pointer to a valid Snmpv3ConfigNotifyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnmpv3ConfigNotifyTypeFromValue(v string) (*Snmpv3ConfigNotifyType, error) {
	ev := Snmpv3ConfigNotifyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Snmpv3ConfigNotifyType: valid values are %v", v, AllowedSnmpv3ConfigNotifyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Snmpv3ConfigNotifyType) IsValid() bool {
	for _, existing := range AllowedSnmpv3ConfigNotifyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to snmpv3_config_notify_type value
func (v Snmpv3ConfigNotifyType) Ptr() *Snmpv3ConfigNotifyType {
	return &v
}

type NullableSnmpv3ConfigNotifyType struct {
	value *Snmpv3ConfigNotifyType
	isSet bool
}

func (v NullableSnmpv3ConfigNotifyType) Get() *Snmpv3ConfigNotifyType {
	return v.value
}

func (v *NullableSnmpv3ConfigNotifyType) Set(val *Snmpv3ConfigNotifyType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpv3ConfigNotifyType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpv3ConfigNotifyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpv3ConfigNotifyType(val *Snmpv3ConfigNotifyType) *NullableSnmpv3ConfigNotifyType {
	return &NullableSnmpv3ConfigNotifyType{value: val, isSet: true}
}

func (v NullableSnmpv3ConfigNotifyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpv3ConfigNotifyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

