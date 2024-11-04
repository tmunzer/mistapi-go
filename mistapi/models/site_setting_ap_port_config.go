package models

import (
    "encoding/json"
)

// SiteSettingApPortConfig represents a SiteSettingApPortConfig struct.
type SiteSettingApPortConfig struct {
    // Property key is the AP model (e.g "AP32")
    ModelSpecific        map[string]ApPortConfig `json:"model_specific,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApPortConfig.
// It customizes the JSON marshaling process for SiteSettingApPortConfig objects.
func (s SiteSettingApPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApPortConfig object to a map representation for JSON marshaling.
func (s SiteSettingApPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ModelSpecific != nil {
        structMap["model_specific"] = s.ModelSpecific
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingApPortConfig.
// It customizes the JSON unmarshaling process for SiteSettingApPortConfig objects.
func (s *SiteSettingApPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingApPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "model_specific")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ModelSpecific = temp.ModelSpecific
    return nil
}

// tempSiteSettingApPortConfig is a temporary struct used for validating the fields of SiteSettingApPortConfig.
type tempSiteSettingApPortConfig  struct {
    ModelSpecific map[string]ApPortConfig `json:"model_specific,omitempty"`
}
