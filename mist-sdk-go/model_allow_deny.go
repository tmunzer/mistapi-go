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

// AllowDeny the model 'AllowDeny'
type AllowDeny string

// List of allow_deny
const (
	ALLOWDENY_EMPTY AllowDeny = ""
	ALLOWDENY_ALLOW AllowDeny = "allow"
	ALLOWDENY_DENY AllowDeny = "deny"
)

// All allowed values of AllowDeny enum
var AllowedAllowDenyEnumValues = []AllowDeny{
	"",
	"allow",
	"deny",
}

func (v *AllowDeny) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllowDeny(value)
	for _, existing := range AllowedAllowDenyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllowDeny", value)
}

// NewAllowDenyFromValue returns a pointer to a valid AllowDeny
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowDenyFromValue(v string) (*AllowDeny, error) {
	ev := AllowDeny(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowDeny: valid values are %v", v, AllowedAllowDenyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowDeny) IsValid() bool {
	for _, existing := range AllowedAllowDenyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to allow_deny value
func (v AllowDeny) Ptr() *AllowDeny {
	return &v
}

type NullableAllowDeny struct {
	value *AllowDeny
	isSet bool
}

func (v NullableAllowDeny) Get() *AllowDeny {
	return v.value
}

func (v *NullableAllowDeny) Set(val *AllowDeny) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowDeny) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowDeny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowDeny(val *AllowDeny) *NullableAllowDeny {
	return &NullableAllowDeny{value: val, isSet: true}
}

func (v NullableAllowDeny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowDeny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
