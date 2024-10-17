package models

import (
	"encoding/json"
)

// SiteSettingApMatching represents a SiteSettingApMatching struct.
type SiteSettingApMatching struct {
	Enabled              *bool                       `json:"enabled,omitempty"`
	Rules                []SiteSettingApMatchingRule `json:"rules,omitempty"`
	AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApMatching.
// It customizes the JSON marshaling process for SiteSettingApMatching objects.
func (s SiteSettingApMatching) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApMatching object to a map representation for JSON marshaling.
func (s SiteSettingApMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "rules")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Enabled = temp.Enabled
	s.Rules = temp.Rules
	return nil
}

// tempSiteSettingApMatching is a temporary struct used for validating the fields of SiteSettingApMatching.
type tempSiteSettingApMatching struct {
	Enabled *bool                       `json:"enabled,omitempty"`
	Rules   []SiteSettingApMatchingRule `json:"rules,omitempty"`
}
