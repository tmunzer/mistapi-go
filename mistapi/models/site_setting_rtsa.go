package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingRtsa represents a SiteSettingRtsa struct.
// Managed mobility
type SiteSettingRtsa struct {
    AppWaking             *bool                  `json:"app_waking,omitempty"`
    DisableDeadReckoning  *bool                  `json:"disable_dead_reckoning,omitempty"`
    DisablePressureSensor *bool                  `json:"disable_pressure_sensor,omitempty"`
    Enabled               *bool                  `json:"enabled,omitempty"`
    // Asset tracking related
    TrackAsset            *bool                  `json:"track_asset,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingRtsa,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingRtsa) String() string {
    return fmt.Sprintf(
    	"SiteSettingRtsa[AppWaking=%v, DisableDeadReckoning=%v, DisablePressureSensor=%v, Enabled=%v, TrackAsset=%v, AdditionalProperties=%v]",
    	s.AppWaking, s.DisableDeadReckoning, s.DisablePressureSensor, s.Enabled, s.TrackAsset, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingRtsa.
// It customizes the JSON marshaling process for SiteSettingRtsa objects.
func (s SiteSettingRtsa) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "app_waking", "disable_dead_reckoning", "disable_pressure_sensor", "enabled", "track_asset"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingRtsa object to a map representation for JSON marshaling.
func (s SiteSettingRtsa) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSiteSettingRtsa
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "app_waking", "disable_dead_reckoning", "disable_pressure_sensor", "enabled", "track_asset")
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

// tempSiteSettingRtsa is a temporary struct used for validating the fields of SiteSettingRtsa.
type tempSiteSettingRtsa  struct {
    AppWaking             *bool `json:"app_waking,omitempty"`
    DisableDeadReckoning  *bool `json:"disable_dead_reckoning,omitempty"`
    DisablePressureSensor *bool `json:"disable_pressure_sensor,omitempty"`
    Enabled               *bool `json:"enabled,omitempty"`
    TrackAsset            *bool `json:"track_asset,omitempty"`
}
