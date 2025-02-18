package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingApMatching represents a SiteSettingApMatching struct.
type SiteSettingApMatching struct {
    Enabled              *bool                       `json:"enabled,omitempty"`
    Rules                []SiteSettingApMatchingRule `json:"rules,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingApMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingApMatching) String() string {
    return fmt.Sprintf(
    	"SiteSettingApMatching[Enabled=%v, Rules=%v, AdditionalProperties=%v]",
    	s.Enabled, s.Rules, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApMatching.
// It customizes the JSON marshaling process for SiteSettingApMatching objects.
func (s SiteSettingApMatching) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "rules"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApMatching object to a map representation for JSON marshaling.
func (s SiteSettingApMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Rules != nil {
        structMap["rules"] = s.Rules
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingApMatching.
// It customizes the JSON unmarshaling process for SiteSettingApMatching objects.
func (s *SiteSettingApMatching) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingApMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "rules")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.Rules = temp.Rules
    return nil
}

// tempSiteSettingApMatching is a temporary struct used for validating the fields of SiteSettingApMatching.
type tempSiteSettingApMatching  struct {
    Enabled *bool                       `json:"enabled,omitempty"`
    Rules   []SiteSettingApMatchingRule `json:"rules,omitempty"`
}
