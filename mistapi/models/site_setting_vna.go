package models

import (
	"encoding/json"
)

// SiteSettingVna represents a SiteSettingVna struct.
type SiteSettingVna struct {
	// enable Virtual Network Assistant (using SUB-VNA license). This applied to AP / Switch / Gateway
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingVna.
// It customizes the JSON marshaling process for SiteSettingVna objects.
func (s SiteSettingVna) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingVna object to a map representation for JSON marshaling.
func (s SiteSettingVna) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingVna.
// It customizes the JSON unmarshaling process for SiteSettingVna objects.
func (s *SiteSettingVna) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingVna
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

// tempSiteSettingVna is a temporary struct used for validating the fields of SiteSettingVna.
type tempSiteSettingVna struct {
	Enabled *bool `json:"enabled,omitempty"`
}
