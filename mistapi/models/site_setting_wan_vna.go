package models

import (
	"encoding/json"
)

// SiteSettingWanVna represents a SiteSettingWanVna struct.
type SiteSettingWanVna struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingWanVna.
// It customizes the JSON marshaling process for SiteSettingWanVna objects.
func (s SiteSettingWanVna) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingWanVna object to a map representation for JSON marshaling.
func (s SiteSettingWanVna) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Enabled = temp.Enabled
	return nil
}

// tempSiteSettingWanVna is a temporary struct used for validating the fields of SiteSettingWanVna.
type tempSiteSettingWanVna struct {
	Enabled *bool `json:"enabled,omitempty"`
}
