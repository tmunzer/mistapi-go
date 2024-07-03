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

// SiteSleMetricSummaryScopeParameters the model 'SiteSleMetricSummaryScopeParameters'
type SiteSleMetricSummaryScopeParameters string

// List of site_sle_metric_summary_scope_parameters
const (
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_EMPTY SiteSleMetricSummaryScopeParameters = ""
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_SITE SiteSleMetricSummaryScopeParameters = "site"
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_AP SiteSleMetricSummaryScopeParameters = "ap"
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_SWITCH SiteSleMetricSummaryScopeParameters = "switch"
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_GATEWAY SiteSleMetricSummaryScopeParameters = "gateway"
	SITESLEMETRICSUMMARYSCOPEPARAMETERS_CLIENT SiteSleMetricSummaryScopeParameters = "client"
)

// All allowed values of SiteSleMetricSummaryScopeParameters enum
var AllowedSiteSleMetricSummaryScopeParametersEnumValues = []SiteSleMetricSummaryScopeParameters{
	"",
	"site",
	"ap",
	"switch",
	"gateway",
	"client",
}

func (v *SiteSleMetricSummaryScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleMetricSummaryScopeParameters(value)
	for _, existing := range AllowedSiteSleMetricSummaryScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleMetricSummaryScopeParameters", value)
}

// NewSiteSleMetricSummaryScopeParametersFromValue returns a pointer to a valid SiteSleMetricSummaryScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleMetricSummaryScopeParametersFromValue(v string) (*SiteSleMetricSummaryScopeParameters, error) {
	ev := SiteSleMetricSummaryScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleMetricSummaryScopeParameters: valid values are %v", v, AllowedSiteSleMetricSummaryScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleMetricSummaryScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleMetricSummaryScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_metric_summary_scope_parameters value
func (v SiteSleMetricSummaryScopeParameters) Ptr() *SiteSleMetricSummaryScopeParameters {
	return &v
}

type NullableSiteSleMetricSummaryScopeParameters struct {
	value *SiteSleMetricSummaryScopeParameters
	isSet bool
}

func (v NullableSiteSleMetricSummaryScopeParameters) Get() *SiteSleMetricSummaryScopeParameters {
	return v.value
}

func (v *NullableSiteSleMetricSummaryScopeParameters) Set(val *SiteSleMetricSummaryScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleMetricSummaryScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleMetricSummaryScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleMetricSummaryScopeParameters(val *SiteSleMetricSummaryScopeParameters) *NullableSiteSleMetricSummaryScopeParameters {
	return &NullableSiteSleMetricSummaryScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleMetricSummaryScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleMetricSummaryScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

