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

// ClaimType what to claim
type ClaimType string

// List of claim_type
const (
	CLAIMTYPE_EMPTY ClaimType = ""
	CLAIMTYPE_ALL ClaimType = "all"
	CLAIMTYPE_LICENSE ClaimType = "license"
	CLAIMTYPE_INVENTORY ClaimType = "inventory"
)

// All allowed values of ClaimType enum
var AllowedClaimTypeEnumValues = []ClaimType{
	"",
	"all",
	"license",
	"inventory",
}

func (v *ClaimType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClaimType(value)
	for _, existing := range AllowedClaimTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClaimType", value)
}

// NewClaimTypeFromValue returns a pointer to a valid ClaimType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClaimTypeFromValue(v string) (*ClaimType, error) {
	ev := ClaimType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClaimType: valid values are %v", v, AllowedClaimTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClaimType) IsValid() bool {
	for _, existing := range AllowedClaimTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to claim_type value
func (v ClaimType) Ptr() *ClaimType {
	return &v
}

type NullableClaimType struct {
	value *ClaimType
	isSet bool
}

func (v NullableClaimType) Get() *ClaimType {
	return v.value
}

func (v *NullableClaimType) Set(val *ClaimType) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimType) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimType(val *ClaimType) *NullableClaimType {
	return &NullableClaimType{value: val, isSet: true}
}

func (v NullableClaimType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

