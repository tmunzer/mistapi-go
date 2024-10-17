package models

import (
	"encoding/json"
)

// SiteSettingTuntermMulticastConfigSsdp represents a SiteSettingTuntermMulticastConfigSsdp struct.
type SiteSettingTuntermMulticastConfigSsdp struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	VlanIds              []int          `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingTuntermMulticastConfigSsdp.
// It customizes the JSON marshaling process for SiteSettingTuntermMulticastConfigSsdp objects.
func (s SiteSettingTuntermMulticastConfigSsdp) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingTuntermMulticastConfigSsdp object to a map representation for JSON marshaling.
func (s SiteSettingTuntermMulticastConfigSsdp) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.VlanIds != nil {
		structMap["vlan_ids"] = s.VlanIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingTuntermMulticastConfigSsdp.
// It customizes the JSON unmarshaling process for SiteSettingTuntermMulticastConfigSsdp objects.
func (s *SiteSettingTuntermMulticastConfigSsdp) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingTuntermMulticastConfigSsdp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "vlan_ids")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Enabled = temp.Enabled
	s.VlanIds = temp.VlanIds
	return nil
}

// tempSiteSettingTuntermMulticastConfigSsdp is a temporary struct used for validating the fields of SiteSettingTuntermMulticastConfigSsdp.
type tempSiteSettingTuntermMulticastConfigSsdp struct {
	Enabled *bool `json:"enabled,omitempty"`
	VlanIds []int `json:"vlan_ids,omitempty"`
}
