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

// LicenseType the model 'LicenseType'
type LicenseType string

// List of license_type
const (
	LICENSETYPE_EMPTY LicenseType = ""
	LICENSETYPE_SUB_MAN LicenseType = "SUB-MAN"
	LICENSETYPE_SUB_AST LicenseType = "SUB-AST"
	LICENSETYPE_SUB_VNA LicenseType = "SUB-VNA"
	LICENSETYPE_SUB_DATA LicenseType = "SUB-DATA"
	LICENSETYPE_SUB_ENG LicenseType = "SUB-ENG"
	LICENSETYPE_SUB_PMA LicenseType = "SUB-PMA"
	LICENSETYPE_SUB_EX12 LicenseType = "SUB-EX12"
	LICENSETYPE_SUB_EX24 LicenseType = "SUB-EX24"
	LICENSETYPE_SUB_EX48 LicenseType = "SUB-EX48"
	LICENSETYPE_SUB_SVNA LicenseType = "SUB-SVNA"
	LICENSETYPE_SUB_ME LicenseType = "SUB-ME"
	LICENSETYPE_SUB_WAN1 LicenseType = "SUB-WAN1"
	LICENSETYPE_SUB_WAN2 LicenseType = "SUB-WAN2"
	LICENSETYPE_SUB_WVNA1 LicenseType = "SUB-WVNA1"
	LICENSETYPE_SUB_WVNA2 LicenseType = "SUB-WVNA2"
	LICENSETYPE_SUB_SRX1 LicenseType = "SUB-SRX1"
	LICENSETYPE_SUB_SRX2 LicenseType = "SUB-SRX2"
)

// All allowed values of LicenseType enum
var AllowedLicenseTypeEnumValues = []LicenseType{
	"",
	"SUB-MAN",
	"SUB-AST",
	"SUB-VNA",
	"SUB-DATA",
	"SUB-ENG",
	"SUB-PMA",
	"SUB-EX12",
	"SUB-EX24",
	"SUB-EX48",
	"SUB-SVNA",
	"SUB-ME",
	"SUB-WAN1",
	"SUB-WAN2",
	"SUB-WVNA1",
	"SUB-WVNA2",
	"SUB-SRX1",
	"SUB-SRX2",
}

func (v *LicenseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseType(value)
	for _, existing := range AllowedLicenseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseType", value)
}

// NewLicenseTypeFromValue returns a pointer to a valid LicenseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseTypeFromValue(v string) (*LicenseType, error) {
	ev := LicenseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseType: valid values are %v", v, AllowedLicenseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseType) IsValid() bool {
	for _, existing := range AllowedLicenseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to license_type value
func (v LicenseType) Ptr() *LicenseType {
	return &v
}

type NullableLicenseType struct {
	value *LicenseType
	isSet bool
}

func (v NullableLicenseType) Get() *LicenseType {
	return v.value
}

func (v *NullableLicenseType) Set(val *LicenseType) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseType) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseType(val *LicenseType) *NullableLicenseType {
	return &NullableLicenseType{value: val, isSet: true}
}

func (v NullableLicenseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

