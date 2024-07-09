package models

import (
    "encoding/json"
)

// SiteSettingStatusPortal represents a SiteSettingStatusPortal struct.
type SiteSettingStatusPortal struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    Hostnames            []string       `json:"hostnames,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingStatusPortal.
// It customizes the JSON marshaling process for SiteSettingStatusPortal objects.
func (s SiteSettingStatusPortal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingStatusPortal object to a map representation for JSON marshaling.
func (s SiteSettingStatusPortal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp siteSettingStatusPortal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "hostnames")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    s.Hostnames = temp.Hostnames
    return nil
}

// siteSettingStatusPortal is a temporary struct used for validating the fields of SiteSettingStatusPortal.
type siteSettingStatusPortal  struct {
    Enabled   *bool    `json:"enabled,omitempty"`
    Hostnames []string `json:"hostnames,omitempty"`
}
