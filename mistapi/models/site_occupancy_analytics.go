// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteOccupancyAnalytics represents a SiteOccupancyAnalytics struct.
// Occupancy Analytics settings
type SiteOccupancyAnalytics struct {
    // Indicate whether named BLE assets should be included in the zone occupancy calculation
    AssetsEnabled             *bool                  `json:"assets_enabled,omitempty"`
    // Indicate whether connected Wi-Fi clients should be included in the zone occupancy calculation
    ClientsEnabled            *bool                  `json:"clients_enabled,omitempty"`
    // Minimum duration
    MinDuration               *int                   `json:"min_duration,omitempty"`
    // Indicate whether SDK clients should be included in the zone occupancy calculation
    SdkclientsEnabled         *bool                  `json:"sdkclients_enabled,omitempty"`
    // Indicate whether unconnected Wi-Fi clients should be included in the zone occupancy calculation
    UnconnectedClientsEnabled *bool                  `json:"unconnected_clients_enabled,omitempty"`
    AdditionalProperties      map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteOccupancyAnalytics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteOccupancyAnalytics) String() string {
    return fmt.Sprintf(
    	"SiteOccupancyAnalytics[AssetsEnabled=%v, ClientsEnabled=%v, MinDuration=%v, SdkclientsEnabled=%v, UnconnectedClientsEnabled=%v, AdditionalProperties=%v]",
    	s.AssetsEnabled, s.ClientsEnabled, s.MinDuration, s.SdkclientsEnabled, s.UnconnectedClientsEnabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteOccupancyAnalytics.
// It customizes the JSON marshaling process for SiteOccupancyAnalytics objects.
func (s SiteOccupancyAnalytics) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "assets_enabled", "clients_enabled", "min_duration", "sdkclients_enabled", "unconnected_clients_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteOccupancyAnalytics object to a map representation for JSON marshaling.
func (s SiteOccupancyAnalytics) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AssetsEnabled != nil {
        structMap["assets_enabled"] = s.AssetsEnabled
    }
    if s.ClientsEnabled != nil {
        structMap["clients_enabled"] = s.ClientsEnabled
    }
    if s.MinDuration != nil {
        structMap["min_duration"] = s.MinDuration
    }
    if s.SdkclientsEnabled != nil {
        structMap["sdkclients_enabled"] = s.SdkclientsEnabled
    }
    if s.UnconnectedClientsEnabled != nil {
        structMap["unconnected_clients_enabled"] = s.UnconnectedClientsEnabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteOccupancyAnalytics.
// It customizes the JSON unmarshaling process for SiteOccupancyAnalytics objects.
func (s *SiteOccupancyAnalytics) UnmarshalJSON(input []byte) error {
    var temp tempSiteOccupancyAnalytics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "assets_enabled", "clients_enabled", "min_duration", "sdkclients_enabled", "unconnected_clients_enabled")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AssetsEnabled = temp.AssetsEnabled
    s.ClientsEnabled = temp.ClientsEnabled
    s.MinDuration = temp.MinDuration
    s.SdkclientsEnabled = temp.SdkclientsEnabled
    s.UnconnectedClientsEnabled = temp.UnconnectedClientsEnabled
    return nil
}

// tempSiteOccupancyAnalytics is a temporary struct used for validating the fields of SiteOccupancyAnalytics.
type tempSiteOccupancyAnalytics  struct {
    AssetsEnabled             *bool `json:"assets_enabled,omitempty"`
    ClientsEnabled            *bool `json:"clients_enabled,omitempty"`
    MinDuration               *int  `json:"min_duration,omitempty"`
    SdkclientsEnabled         *bool `json:"sdkclients_enabled,omitempty"`
    UnconnectedClientsEnabled *bool `json:"unconnected_clients_enabled,omitempty"`
}
