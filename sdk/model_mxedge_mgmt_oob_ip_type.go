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

// MxedgeMgmtOobIpType the model 'MxedgeMgmtOobIpType'
type MxedgeMgmtOobIpType string

// List of mxedge_mgmt_oob_ip_type
const (
	MXEDGEMGMTOOBIPTYPE_EMPTY MxedgeMgmtOobIpType = ""
	MXEDGEMGMTOOBIPTYPE_DHCP MxedgeMgmtOobIpType = "dhcp"
	MXEDGEMGMTOOBIPTYPE_STATIC MxedgeMgmtOobIpType = "static"
	MXEDGEMGMTOOBIPTYPE_DISABLED MxedgeMgmtOobIpType = "disabled"
)

// All allowed values of MxedgeMgmtOobIpType enum
var AllowedMxedgeMgmtOobIpTypeEnumValues = []MxedgeMgmtOobIpType{
	"",
	"dhcp",
	"static",
	"disabled",
}

func (v *MxedgeMgmtOobIpType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxedgeMgmtOobIpType(value)
	for _, existing := range AllowedMxedgeMgmtOobIpTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxedgeMgmtOobIpType", value)
}

// NewMxedgeMgmtOobIpTypeFromValue returns a pointer to a valid MxedgeMgmtOobIpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxedgeMgmtOobIpTypeFromValue(v string) (*MxedgeMgmtOobIpType, error) {
	ev := MxedgeMgmtOobIpType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxedgeMgmtOobIpType: valid values are %v", v, AllowedMxedgeMgmtOobIpTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxedgeMgmtOobIpType) IsValid() bool {
	for _, existing := range AllowedMxedgeMgmtOobIpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxedge_mgmt_oob_ip_type value
func (v MxedgeMgmtOobIpType) Ptr() *MxedgeMgmtOobIpType {
	return &v
}

type NullableMxedgeMgmtOobIpType struct {
	value *MxedgeMgmtOobIpType
	isSet bool
}

func (v NullableMxedgeMgmtOobIpType) Get() *MxedgeMgmtOobIpType {
	return v.value
}

func (v *NullableMxedgeMgmtOobIpType) Set(val *MxedgeMgmtOobIpType) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeMgmtOobIpType) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeMgmtOobIpType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeMgmtOobIpType(val *MxedgeMgmtOobIpType) *NullableMxedgeMgmtOobIpType {
	return &NullableMxedgeMgmtOobIpType{value: val, isSet: true}
}

func (v NullableMxedgeMgmtOobIpType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeMgmtOobIpType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

