// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingSrxApp represents a SiteSettingSrxApp struct.
type SiteSettingSrxApp struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingSrxApp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingSrxApp) String() string {
    return fmt.Sprintf(
    	"SiteSettingSrxApp[Enabled=%v, AdditionalProperties=%v]",
    	s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingSrxApp.
// It customizes the JSON marshaling process for SiteSettingSrxApp objects.
func (s SiteSettingSrxApp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingSrxApp object to a map representation for JSON marshaling.
func (s SiteSettingSrxApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingSrxApp.
// It customizes the JSON unmarshaling process for SiteSettingSrxApp objects.
func (s *SiteSettingSrxApp) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingSrxApp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    return nil
}

// tempSiteSettingSrxApp is a temporary struct used for validating the fields of SiteSettingSrxApp.
type tempSiteSettingSrxApp  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
