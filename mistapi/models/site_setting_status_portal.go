// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingStatusPortal represents a SiteSettingStatusPortal struct.
type SiteSettingStatusPortal struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    Hostnames            []string               `json:"hostnames,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingStatusPortal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingStatusPortal) String() string {
    return fmt.Sprintf(
    	"SiteSettingStatusPortal[Enabled=%v, Hostnames=%v, AdditionalProperties=%v]",
    	s.Enabled, s.Hostnames, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingStatusPortal.
// It customizes the JSON marshaling process for SiteSettingStatusPortal objects.
func (s SiteSettingStatusPortal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "hostnames"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingStatusPortal object to a map representation for JSON marshaling.
func (s SiteSettingStatusPortal) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Hostnames != nil {
        structMap["hostnames"] = s.Hostnames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingStatusPortal.
// It customizes the JSON unmarshaling process for SiteSettingStatusPortal objects.
func (s *SiteSettingStatusPortal) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingStatusPortal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "hostnames")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.Hostnames = temp.Hostnames
    return nil
}

// tempSiteSettingStatusPortal is a temporary struct used for validating the fields of SiteSettingStatusPortal.
type tempSiteSettingStatusPortal  struct {
    Enabled   *bool    `json:"enabled,omitempty"`
    Hostnames []string `json:"hostnames,omitempty"`
}
