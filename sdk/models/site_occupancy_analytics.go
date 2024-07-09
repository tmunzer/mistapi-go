package models

import (
    "encoding/json"
)

// SiteOccupancyAnalytics represents a SiteOccupancyAnalytics struct.
// Occupancy Analytics settings
type SiteOccupancyAnalytics struct {
    // indicate whether named BLE assets should be included in the zone occupancy calculation
    AssetsEnabled             *bool          `json:"assets_enabled,omitempty"`
    // indicate whether connected WiFi clients should be included in the zone occupancy calculation
    ClientsEnabled            *bool          `json:"clients_enabled,omitempty"`
    // minimum duration
    MinDuration               *int           `json:"min_duration,omitempty"`
    // indicate whether SDK clients should be included in the zone occupancy calculation
    SdkclientsEnabled         *bool          `json:"sdkclients_enabled,omitempty"`
    // indicate whether unconnected WiFi clients should be included in the zone occupancy calculation
    UnconnectedClientsEnabled *bool          `json:"unconnected_clients_enabled,omitempty"`
    AdditionalProperties      map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteOccupancyAnalytics.
// It customizes the JSON marshaling process for SiteOccupancyAnalytics objects.
func (s SiteOccupancyAnalytics) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteOccupancyAnalytics object to a map representation for JSON marshaling.
func (s SiteOccupancyAnalytics) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp siteOccupancyAnalytics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "assets_enabled", "clients_enabled", "min_duration", "sdkclients_enabled", "unconnected_clients_enabled")
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

// siteOccupancyAnalytics is a temporary struct used for validating the fields of SiteOccupancyAnalytics.
type siteOccupancyAnalytics  struct {
    AssetsEnabled             *bool `json:"assets_enabled,omitempty"`
    ClientsEnabled            *bool `json:"clients_enabled,omitempty"`
    MinDuration               *int  `json:"min_duration,omitempty"`
    SdkclientsEnabled         *bool `json:"sdkclients_enabled,omitempty"`
    UnconnectedClientsEnabled *bool `json:"unconnected_clients_enabled,omitempty"`
}
