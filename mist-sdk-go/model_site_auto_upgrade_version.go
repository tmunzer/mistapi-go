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

// SiteAutoUpgradeVersion desired version
type SiteAutoUpgradeVersion string

// List of site_auto_upgrade_version
const (
	SITEAUTOUPGRADEVERSION_EMPTY SiteAutoUpgradeVersion = ""
	SITEAUTOUPGRADEVERSION_BETA SiteAutoUpgradeVersion = "beta"
	SITEAUTOUPGRADEVERSION_STABLE SiteAutoUpgradeVersion = "stable"
	SITEAUTOUPGRADEVERSION_CUSTOM SiteAutoUpgradeVersion = "custom"
)

// All allowed values of SiteAutoUpgradeVersion enum
var AllowedSiteAutoUpgradeVersionEnumValues = []SiteAutoUpgradeVersion{
	"",
	"beta",
	"stable",
	"custom",
}

func (v *SiteAutoUpgradeVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteAutoUpgradeVersion(value)
	for _, existing := range AllowedSiteAutoUpgradeVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteAutoUpgradeVersion", value)
}

// NewSiteAutoUpgradeVersionFromValue returns a pointer to a valid SiteAutoUpgradeVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteAutoUpgradeVersionFromValue(v string) (*SiteAutoUpgradeVersion, error) {
	ev := SiteAutoUpgradeVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteAutoUpgradeVersion: valid values are %v", v, AllowedSiteAutoUpgradeVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteAutoUpgradeVersion) IsValid() bool {
	for _, existing := range AllowedSiteAutoUpgradeVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_auto_upgrade_version value
func (v SiteAutoUpgradeVersion) Ptr() *SiteAutoUpgradeVersion {
	return &v
}

type NullableSiteAutoUpgradeVersion struct {
	value *SiteAutoUpgradeVersion
	isSet bool
}

func (v NullableSiteAutoUpgradeVersion) Get() *SiteAutoUpgradeVersion {
	return v.value
}

func (v *NullableSiteAutoUpgradeVersion) Set(val *SiteAutoUpgradeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAutoUpgradeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAutoUpgradeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAutoUpgradeVersion(val *SiteAutoUpgradeVersion) *NullableSiteAutoUpgradeVersion {
	return &NullableSiteAutoUpgradeVersion{value: val, isSet: true}
}

func (v NullableSiteAutoUpgradeVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAutoUpgradeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

