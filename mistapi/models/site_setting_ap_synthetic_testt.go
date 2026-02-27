// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteSettingApSyntheticTest represents a SiteSettingApSyntheticTest struct.
// AP Synthetic Test configuration
type SiteSettingApSyntheticTest struct {
	// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
	AdditionalVlanIds    *AdditionalVlanIds     `json:"additional_vlan_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingApSyntheticTest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingApSyntheticTest) String() string {
	return fmt.Sprintf(
		"SiteSettingApSyntheticTest[AdditionalVlanIds=%v, AdditionalProperties=%v]",
		s.AdditionalVlanIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApSyntheticTest.
// It customizes the JSON marshaling process for SiteSettingApSyntheticTest objects.
func (s SiteSettingApSyntheticTest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"additional_vlan_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApSyntheticTest object to a map representation for JSON marshaling.
func (s SiteSettingApSyntheticTest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AdditionalVlanIds != nil {
		structMap["additional_vlan_ids"] = s.AdditionalVlanIds.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingApSyntheticTest.
// It customizes the JSON unmarshaling process for SiteSettingApSyntheticTest objects.
func (s *SiteSettingApSyntheticTest) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingApSyntheticTest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_vlan_ids")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AdditionalVlanIds = temp.AdditionalVlanIds
	return nil
}

// tempSiteSettingApSyntheticTest is a temporary struct used for validating the fields of SiteSettingApSyntheticTest.
type tempSiteSettingApSyntheticTest struct {
	AdditionalVlanIds *AdditionalVlanIds `json:"additional_vlan_ids,omitempty"`
}
