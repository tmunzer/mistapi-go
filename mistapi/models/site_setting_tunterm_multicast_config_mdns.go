package models

import (
    "encoding/json"
)

// SiteSettingTuntermMulticastConfigMdns represents a SiteSettingTuntermMulticastConfigMdns struct.
type SiteSettingTuntermMulticastConfigMdns struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    VlanIds              []int          `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingTuntermMulticastConfigMdns.
// It customizes the JSON marshaling process for SiteSettingTuntermMulticastConfigMdns objects.
func (s SiteSettingTuntermMulticastConfigMdns) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingTuntermMulticastConfigMdns object to a map representation for JSON marshaling.
func (s SiteSettingTuntermMulticastConfigMdns) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingTuntermMulticastConfigMdns.
// It customizes the JSON unmarshaling process for SiteSettingTuntermMulticastConfigMdns objects.
func (s *SiteSettingTuntermMulticastConfigMdns) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingTuntermMulticastConfigMdns
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

// tempSiteSettingTuntermMulticastConfigMdns is a temporary struct used for validating the fields of SiteSettingTuntermMulticastConfigMdns.
type tempSiteSettingTuntermMulticastConfigMdns  struct {
    Enabled *bool `json:"enabled,omitempty"`
    VlanIds []int `json:"vlan_ids,omitempty"`
}
