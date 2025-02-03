package models

import (
    "encoding/json"
    "fmt"
)

// SiteRogue represents a SiteRogue struct.
// Rogue site settings
type SiteRogue struct {
    // Whether or not rogue detection is enabled
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Whether or not honeypot detection is enabled
    HoneypotEnabled      *bool                  `json:"honeypot_enabled,omitempty"`
    // Minimum duration for a bssid to be considered rogue
    MinDuration          *int                   `json:"min_duration,omitempty"`
    // Minimum RSSI for an AP to be considered rogue (ignoring APs that’s far away)
    MinRssi              *int                   `json:"min_rssi,omitempty"`
    // list of BSSIDs to whitelist. Ex: "cc-:8e-:6f-:d4-:bf-:16", "cc-8e-6f-d4-bf-16", "cc-73-*", "cc:82:*"
    WhitelistedBssids    []string               `json:"whitelisted_bssids,omitempty"`
    // List of SSIDs to whitelist
    WhitelistedSsids     []string               `json:"whitelisted_ssids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteRogue,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteRogue) String() string {
    return fmt.Sprintf(
    	"SiteRogue[Enabled=%v, HoneypotEnabled=%v, MinDuration=%v, MinRssi=%v, WhitelistedBssids=%v, WhitelistedSsids=%v, AdditionalProperties=%v]",
    	s.Enabled, s.HoneypotEnabled, s.MinDuration, s.MinRssi, s.WhitelistedBssids, s.WhitelistedSsids, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteRogue.
// It customizes the JSON marshaling process for SiteRogue objects.
func (s SiteRogue) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "honeypot_enabled", "min_duration", "min_rssi", "whitelisted_bssids", "whitelisted_ssids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteRogue object to a map representation for JSON marshaling.
func (s SiteRogue) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.HoneypotEnabled != nil {
        structMap["honeypot_enabled"] = s.HoneypotEnabled
    }
    if s.MinDuration != nil {
        structMap["min_duration"] = s.MinDuration
    }
    if s.MinRssi != nil {
        structMap["min_rssi"] = s.MinRssi
    }
    if s.WhitelistedBssids != nil {
        structMap["whitelisted_bssids"] = s.WhitelistedBssids
    }
    if s.WhitelistedSsids != nil {
        structMap["whitelisted_ssids"] = s.WhitelistedSsids
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteRogue.
// It customizes the JSON unmarshaling process for SiteRogue objects.
func (s *SiteRogue) UnmarshalJSON(input []byte) error {
    var temp tempSiteRogue
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "honeypot_enabled", "min_duration", "min_rssi", "whitelisted_bssids", "whitelisted_ssids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.HoneypotEnabled = temp.HoneypotEnabled
    s.MinDuration = temp.MinDuration
    s.MinRssi = temp.MinRssi
    s.WhitelistedBssids = temp.WhitelistedBssids
    s.WhitelistedSsids = temp.WhitelistedSsids
    return nil
}

// tempSiteRogue is a temporary struct used for validating the fields of SiteRogue.
type tempSiteRogue  struct {
    Enabled           *bool    `json:"enabled,omitempty"`
    HoneypotEnabled   *bool    `json:"honeypot_enabled,omitempty"`
    MinDuration       *int     `json:"min_duration,omitempty"`
    MinRssi           *int     `json:"min_rssi,omitempty"`
    WhitelistedBssids []string `json:"whitelisted_bssids,omitempty"`
    WhitelistedSsids  []string `json:"whitelisted_ssids,omitempty"`
}
