// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingSkyatp represents a SiteSettingSkyatp struct.
type SiteSettingSkyatp struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Whether to send IP-MAC mapping to SkyATP
    SendIpMacMapping     *bool                  `json:"send_ip_mac_mapping,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingSkyatp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingSkyatp) String() string {
    return fmt.Sprintf(
    	"SiteSettingSkyatp[Enabled=%v, SendIpMacMapping=%v, AdditionalProperties=%v]",
    	s.Enabled, s.SendIpMacMapping, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingSkyatp.
// It customizes the JSON marshaling process for SiteSettingSkyatp objects.
func (s SiteSettingSkyatp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "send_ip_mac_mapping"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingSkyatp object to a map representation for JSON marshaling.
func (s SiteSettingSkyatp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "send_ip_mac_mapping")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.SendIpMacMapping = temp.SendIpMacMapping
    return nil
}

// tempSiteSettingSkyatp is a temporary struct used for validating the fields of SiteSettingSkyatp.
type tempSiteSettingSkyatp  struct {
    Enabled          *bool `json:"enabled,omitempty"`
    SendIpMacMapping *bool `json:"send_ip_mac_mapping,omitempty"`
}
