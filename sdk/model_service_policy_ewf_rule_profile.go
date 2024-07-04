/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// ServicePolicyEwfRuleProfile the model 'ServicePolicyEwfRuleProfile'
type ServicePolicyEwfRuleProfile string

// List of service_policy_ewf_rule_profile
const (
	SERVICEPOLICYEWFRULEPROFILE_EMPTY ServicePolicyEwfRuleProfile = ""
	SERVICEPOLICYEWFRULEPROFILE_STRICT ServicePolicyEwfRuleProfile = "strict"
	SERVICEPOLICYEWFRULEPROFILE_STANDARD ServicePolicyEwfRuleProfile = "standard"
)

// All allowed values of ServicePolicyEwfRuleProfile enum
var AllowedServicePolicyEwfRuleProfileEnumValues = []ServicePolicyEwfRuleProfile{
	"",
	"strict",
	"standard",
}

func (v *ServicePolicyEwfRuleProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServicePolicyEwfRuleProfile(value)
	for _, existing := range AllowedServicePolicyEwfRuleProfileEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServicePolicyEwfRuleProfile", value)
}

// NewServicePolicyEwfRuleProfileFromValue returns a pointer to a valid ServicePolicyEwfRuleProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServicePolicyEwfRuleProfileFromValue(v string) (*ServicePolicyEwfRuleProfile, error) {
	ev := ServicePolicyEwfRuleProfile(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServicePolicyEwfRuleProfile: valid values are %v", v, AllowedServicePolicyEwfRuleProfileEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServicePolicyEwfRuleProfile) IsValid() bool {
	for _, existing := range AllowedServicePolicyEwfRuleProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to service_policy_ewf_rule_profile value
func (v ServicePolicyEwfRuleProfile) Ptr() *ServicePolicyEwfRuleProfile {
	return &v
}

type NullableServicePolicyEwfRuleProfile struct {
	value *ServicePolicyEwfRuleProfile
	isSet bool
}

func (v NullableServicePolicyEwfRuleProfile) Get() *ServicePolicyEwfRuleProfile {
	return v.value
}

func (v *NullableServicePolicyEwfRuleProfile) Set(val *ServicePolicyEwfRuleProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePolicyEwfRuleProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePolicyEwfRuleProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePolicyEwfRuleProfile(val *ServicePolicyEwfRuleProfile) *NullableServicePolicyEwfRuleProfile {
	return &NullableServicePolicyEwfRuleProfile{value: val, isSet: true}
}

func (v NullableServicePolicyEwfRuleProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePolicyEwfRuleProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

