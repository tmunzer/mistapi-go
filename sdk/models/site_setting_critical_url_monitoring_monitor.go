package models

import (
    "encoding/json"
)

// SiteSettingCriticalUrlMonitoringMonitor represents a SiteSettingCriticalUrlMonitoringMonitor struct.
type SiteSettingCriticalUrlMonitoringMonitor struct {
    Url                  *string        `json:"url,omitempty"`
    VlanId               *int           `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingCriticalUrlMonitoringMonitor.
// It customizes the JSON marshaling process for SiteSettingCriticalUrlMonitoringMonitor objects.
func (s SiteSettingCriticalUrlMonitoringMonitor) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingCriticalUrlMonitoringMonitor object to a map representation for JSON marshaling.
func (s SiteSettingCriticalUrlMonitoringMonitor) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Url != nil {
        structMap["url"] = s.Url
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingCriticalUrlMonitoringMonitor.
// It customizes the JSON unmarshaling process for SiteSettingCriticalUrlMonitoringMonitor objects.
func (s *SiteSettingCriticalUrlMonitoringMonitor) UnmarshalJSON(input []byte) error {
    var temp siteSettingCriticalUrlMonitoringMonitor
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "url", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Url = temp.Url
    s.VlanId = temp.VlanId
    return nil
}

// siteSettingCriticalUrlMonitoringMonitor is a temporary struct used for validating the fields of SiteSettingCriticalUrlMonitoringMonitor.
type siteSettingCriticalUrlMonitoringMonitor  struct {
    Url    *string `json:"url,omitempty"`
    VlanId *int    `json:"vlan_id,omitempty"`
}
