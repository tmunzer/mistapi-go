package models

import (
	"encoding/json"
)

// SiteSettingAnalytic represents a SiteSettingAnalytic struct.
type SiteSettingAnalytic struct {
	// enable Advanced Analytic feature (using SUB-ANA license)
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingAnalytic.
// It customizes the JSON marshaling process for SiteSettingAnalytic objects.
func (s SiteSettingAnalytic) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingAnalytic object to a map representation for JSON marshaling.
func (s SiteSettingAnalytic) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingAnalytic.
// It customizes the JSON unmarshaling process for SiteSettingAnalytic objects.
func (s *SiteSettingAnalytic) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingAnalytic
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

// tempSiteSettingAnalytic is a temporary struct used for validating the fields of SiteSettingAnalytic.
type tempSiteSettingAnalytic struct {
	Enabled *bool `json:"enabled,omitempty"`
}
