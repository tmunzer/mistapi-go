package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingWanVna represents a SiteSettingWanVna struct.
type SiteSettingWanVna struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingWanVna,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingWanVna) String() string {
    return fmt.Sprintf(
    	"SiteSettingWanVna[Enabled=%v, AdditionalProperties=%v]",
    	s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingWanVna.
// It customizes the JSON marshaling process for SiteSettingWanVna objects.
func (s SiteSettingWanVna) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingWanVna object to a map representation for JSON marshaling.
func (s SiteSettingWanVna) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingWanVna.
// It customizes the JSON unmarshaling process for SiteSettingWanVna objects.
func (s *SiteSettingWanVna) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingWanVna
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

// tempSiteSettingWanVna is a temporary struct used for validating the fields of SiteSettingWanVna.
type tempSiteSettingWanVna  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
