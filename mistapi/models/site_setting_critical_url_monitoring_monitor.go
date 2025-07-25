// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingCriticalUrlMonitoringMonitor represents a SiteSettingCriticalUrlMonitoringMonitor struct.
type SiteSettingCriticalUrlMonitoringMonitor struct {
    Url                  *string                `json:"url,omitempty"`
    VlanId               *VlanIdWithVariable    `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingCriticalUrlMonitoringMonitor,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingCriticalUrlMonitoringMonitor) String() string {
    return fmt.Sprintf(
    	"SiteSettingCriticalUrlMonitoringMonitor[Url=%v, VlanId=%v, AdditionalProperties=%v]",
    	s.Url, s.VlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingCriticalUrlMonitoringMonitor.
// It customizes the JSON marshaling process for SiteSettingCriticalUrlMonitoringMonitor objects.
func (s SiteSettingCriticalUrlMonitoringMonitor) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "url", "vlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingCriticalUrlMonitoringMonitor object to a map representation for JSON marshaling.
func (s SiteSettingCriticalUrlMonitoringMonitor) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Url != nil {
        structMap["url"] = s.Url
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingCriticalUrlMonitoringMonitor.
// It customizes the JSON unmarshaling process for SiteSettingCriticalUrlMonitoringMonitor objects.
func (s *SiteSettingCriticalUrlMonitoringMonitor) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingCriticalUrlMonitoringMonitor
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "url", "vlan_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Url = temp.Url
    s.VlanId = temp.VlanId
    return nil
}

// tempSiteSettingCriticalUrlMonitoringMonitor is a temporary struct used for validating the fields of SiteSettingCriticalUrlMonitoringMonitor.
type tempSiteSettingCriticalUrlMonitoringMonitor  struct {
    Url    *string             `json:"url,omitempty"`
    VlanId *VlanIdWithVariable `json:"vlan_id,omitempty"`
}
