package models

import (
    "encoding/json"
)

// SiteSettingRtsa represents a SiteSettingRtsa struct.
// managed mobility
type SiteSettingRtsa struct {
    AppWaking             *bool          `json:"app_waking,omitempty"`
    DisableDeadReckoning  *bool          `json:"disable_dead_reckoning,omitempty"`
    DisablePressureSensor *bool          `json:"disable_pressure_sensor,omitempty"`
    Enabled               *bool          `json:"enabled,omitempty"`
    // asset tracking related
    TrackAsset            *bool          `json:"track_asset,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingRtsa.
// It customizes the JSON marshaling process for SiteSettingRtsa objects.
func (s SiteSettingRtsa) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingRtsa object to a map representation for JSON marshaling.
func (s SiteSettingRtsa) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AppWaking != nil {
        structMap["app_waking"] = s.AppWaking
    }
    if s.DisableDeadReckoning != nil {
        structMap["disable_dead_reckoning"] = s.DisableDeadReckoning
    }
    if s.DisablePressureSensor != nil {
        structMap["disable_pressure_sensor"] = s.DisablePressureSensor
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.TrackAsset != nil {
        structMap["track_asset"] = s.TrackAsset
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingRtsa.
// It customizes the JSON unmarshaling process for SiteSettingRtsa objects.
func (s *SiteSettingRtsa) UnmarshalJSON(input []byte) error {
    var temp siteSettingRtsa
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "app_waking", "disable_dead_reckoning", "disable_pressure_sensor", "enabled", "track_asset")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AppWaking = temp.AppWaking
    s.DisableDeadReckoning = temp.DisableDeadReckoning
    s.DisablePressureSensor = temp.DisablePressureSensor
    s.Enabled = temp.Enabled
    s.TrackAsset = temp.TrackAsset
    return nil
}

// siteSettingRtsa is a temporary struct used for validating the fields of SiteSettingRtsa.
type siteSettingRtsa  struct {
    AppWaking             *bool `json:"app_waking,omitempty"`
    DisableDeadReckoning  *bool `json:"disable_dead_reckoning,omitempty"`
    DisablePressureSensor *bool `json:"disable_pressure_sensor,omitempty"`
    Enabled               *bool `json:"enabled,omitempty"`
    TrackAsset            *bool `json:"track_asset,omitempty"`
}
