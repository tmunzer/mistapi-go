package models

import (
    "encoding/json"
)

// SiteSettingWiredVna represents a SiteSettingWiredVna struct.
type SiteSettingWiredVna struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingWiredVna.
// It customizes the JSON marshaling process for SiteSettingWiredVna objects.
func (s SiteSettingWiredVna) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingWiredVna object to a map representation for JSON marshaling.
func (s SiteSettingWiredVna) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingWiredVna.
// It customizes the JSON unmarshaling process for SiteSettingWiredVna objects.
func (s *SiteSettingWiredVna) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingWiredVna
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

// tempSiteSettingWiredVna is a temporary struct used for validating the fields of SiteSettingWiredVna.
type tempSiteSettingWiredVna  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
