package models

import (
    "encoding/json"
)

// SiteSettingSrxApp represents a SiteSettingSrxApp struct.
type SiteSettingSrxApp struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingSrxApp.
// It customizes the JSON marshaling process for SiteSettingSrxApp objects.
func (s SiteSettingSrxApp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingSrxApp object to a map representation for JSON marshaling.
func (s SiteSettingSrxApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
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
