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

// SiteSleMetricsScopeParameters the model 'SiteSleMetricsScopeParameters'
type SiteSleMetricsScopeParameters string

// List of site_sle_metrics_scope_parameters
const (
	SITESLEMETRICSSCOPEPARAMETERS_EMPTY SiteSleMetricsScopeParameters = ""
	SITESLEMETRICSSCOPEPARAMETERS_SITE SiteSleMetricsScopeParameters = "site"
	SITESLEMETRICSSCOPEPARAMETERS_AP SiteSleMetricsScopeParameters = "ap"
	SITESLEMETRICSSCOPEPARAMETERS_SWITCH SiteSleMetricsScopeParameters = "switch"
	SITESLEMETRICSSCOPEPARAMETERS_GATEWAY SiteSleMetricsScopeParameters = "gateway"
	SITESLEMETRICSSCOPEPARAMETERS_CLIENT SiteSleMetricsScopeParameters = "client"
)

// All allowed values of SiteSleMetricsScopeParameters enum
var AllowedSiteSleMetricsScopeParametersEnumValues = []SiteSleMetricsScopeParameters{
	"",
	"site",
	"ap",
	"switch",
	"gateway",
	"client",
}

func (v *SiteSleMetricsScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleMetricsScopeParameters(value)
	for _, existing := range AllowedSiteSleMetricsScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleMetricsScopeParameters", value)
}

// NewSiteSleMetricsScopeParametersFromValue returns a pointer to a valid SiteSleMetricsScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleMetricsScopeParametersFromValue(v string) (*SiteSleMetricsScopeParameters, error) {
	ev := SiteSleMetricsScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleMetricsScopeParameters: valid values are %v", v, AllowedSiteSleMetricsScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleMetricsScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleMetricsScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_metrics_scope_parameters value
func (v SiteSleMetricsScopeParameters) Ptr() *SiteSleMetricsScopeParameters {
	return &v
}

type NullableSiteSleMetricsScopeParameters struct {
	value *SiteSleMetricsScopeParameters
	isSet bool
}

func (v NullableSiteSleMetricsScopeParameters) Get() *SiteSleMetricsScopeParameters {
	return v.value
}

func (v *NullableSiteSleMetricsScopeParameters) Set(val *SiteSleMetricsScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleMetricsScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleMetricsScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleMetricsScopeParameters(val *SiteSleMetricsScopeParameters) *NullableSiteSleMetricsScopeParameters {
	return &NullableSiteSleMetricsScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleMetricsScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleMetricsScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

