package models

import (
	"encoding/json"
)

// SiteSettingSkyatp represents a SiteSettingSkyatp struct.
type SiteSettingSkyatp struct {
	Enabled *bool `json:"enabled,omitempty"`
	// whether to send IP-MAC mapping to SkyATP
	SendIpMacMapping     *bool          `json:"send_ip_mac_mapping,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingSkyatp.
// It customizes the JSON marshaling process for SiteSettingSkyatp objects.
func (s SiteSettingSkyatp) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingSkyatp object to a map representation for JSON marshaling.
func (s SiteSettingSkyatp) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.SendIpMacMapping != nil {
		structMap["send_ip_mac_mapping"] = s.SendIpMacMapping
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingSkyatp.
// It customizes the JSON unmarshaling process for SiteSettingSkyatp objects.
func (s *SiteSettingSkyatp) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingSkyatp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "send_ip_mac_mapping")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Enabled = temp.Enabled
	s.SendIpMacMapping = temp.SendIpMacMapping
	return nil
}

// tempSiteSettingSkyatp is a temporary struct used for validating the fields of SiteSettingSkyatp.
type tempSiteSettingSkyatp struct {
	Enabled          *bool `json:"enabled,omitempty"`
	SendIpMacMapping *bool `json:"send_ip_mac_mapping,omitempty"`
}
