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

// SiteSleImpactedClientsScopeParameters the model 'SiteSleImpactedClientsScopeParameters'
type SiteSleImpactedClientsScopeParameters string

// List of site_sle_impacted_clients_scope_parameters
const (
	SITESLEIMPACTEDCLIENTSSCOPEPARAMETERS_EMPTY SiteSleImpactedClientsScopeParameters = ""
	SITESLEIMPACTEDCLIENTSSCOPEPARAMETERS_SITE SiteSleImpactedClientsScopeParameters = "site"
	SITESLEIMPACTEDCLIENTSSCOPEPARAMETERS_SWITCH SiteSleImpactedClientsScopeParameters = "switch"
	SITESLEIMPACTEDCLIENTSSCOPEPARAMETERS_GATEWAY SiteSleImpactedClientsScopeParameters = "gateway"
)

// All allowed values of SiteSleImpactedClientsScopeParameters enum
var AllowedSiteSleImpactedClientsScopeParametersEnumValues = []SiteSleImpactedClientsScopeParameters{
	"",
	"site",
	"switch",
	"gateway",
}

func (v *SiteSleImpactedClientsScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleImpactedClientsScopeParameters(value)
	for _, existing := range AllowedSiteSleImpactedClientsScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleImpactedClientsScopeParameters", value)
}

// NewSiteSleImpactedClientsScopeParametersFromValue returns a pointer to a valid SiteSleImpactedClientsScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleImpactedClientsScopeParametersFromValue(v string) (*SiteSleImpactedClientsScopeParameters, error) {
	ev := SiteSleImpactedClientsScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleImpactedClientsScopeParameters: valid values are %v", v, AllowedSiteSleImpactedClientsScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleImpactedClientsScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleImpactedClientsScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_impacted_clients_scope_parameters value
func (v SiteSleImpactedClientsScopeParameters) Ptr() *SiteSleImpactedClientsScopeParameters {
	return &v
}

type NullableSiteSleImpactedClientsScopeParameters struct {
	value *SiteSleImpactedClientsScopeParameters
	isSet bool
}

func (v NullableSiteSleImpactedClientsScopeParameters) Get() *SiteSleImpactedClientsScopeParameters {
	return v.value
}

func (v *NullableSiteSleImpactedClientsScopeParameters) Set(val *SiteSleImpactedClientsScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleImpactedClientsScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleImpactedClientsScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleImpactedClientsScopeParameters(val *SiteSleImpactedClientsScopeParameters) *NullableSiteSleImpactedClientsScopeParameters {
	return &NullableSiteSleImpactedClientsScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleImpactedClientsScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleImpactedClientsScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

