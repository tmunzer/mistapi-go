package models

import (
    "encoding/json"
)

// SiteSettingCriticalUrlMonitoring represents a SiteSettingCriticalUrlMonitoring struct.
// you can define some URLs that's critical to site operaitons the latency will be captured and considered for site health
type SiteSettingCriticalUrlMonitoring struct {
    Enabled              *bool                                     `json:"enabled,omitempty"`
    Monitors             []SiteSettingCriticalUrlMonitoringMonitor `json:"monitors,omitempty"`
    AdditionalProperties map[string]any                            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingCriticalUrlMonitoring.
// It customizes the JSON marshaling process for SiteSettingCriticalUrlMonitoring objects.
func (s SiteSettingCriticalUrlMonitoring) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingCriticalUrlMonitoring object to a map representation for JSON marshaling.
func (s SiteSettingCriticalUrlMonitoring) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Monitors != nil {
        structMap["monitors"] = s.Monitors
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingCriticalUrlMonitoring.
// It customizes the JSON unmarshaling process for SiteSettingCriticalUrlMonitoring objects.
func (s *SiteSettingCriticalUrlMonitoring) UnmarshalJSON(input []byte) error {
    var temp siteSettingCriticalUrlMonitoring
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "monitors")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    s.Monitors = temp.Monitors
    return nil
}

// siteSettingCriticalUrlMonitoring is a temporary struct used for validating the fields of SiteSettingCriticalUrlMonitoring.
type siteSettingCriticalUrlMonitoring  struct {
    Enabled  *bool                                     `json:"enabled,omitempty"`
    Monitors []SiteSettingCriticalUrlMonitoringMonitor `json:"monitors,omitempty"`
}
