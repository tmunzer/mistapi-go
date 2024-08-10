package models

import (
    "encoding/json"
)

// ApStatsRadioConfig represents a ApStatsRadioConfig struct.
type ApStatsRadioConfig struct {
    Band24               *ApStatsRadioConfigBand `json:"band_24,omitempty"`
    Band24Usage          Optional[string]        `json:"band_24_usage"`
    Band5                *ApStatsRadioConfigBand `json:"band_5,omitempty"`
    Band6                *ApStatsRadioConfigBand `json:"band_6,omitempty"`
    ScanningEnabled      *bool                   `json:"scanning_enabled,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsRadioConfig.
// It customizes the JSON marshaling process for ApStatsRadioConfig objects.
func (a ApStatsRadioConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsRadioConfig object to a map representation for JSON marshaling.
func (a ApStatsRadioConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Band24 != nil {
        structMap["band_24"] = a.Band24.toMap()
    }
    if a.Band24Usage.IsValueSet() {
        if a.Band24Usage.Value() != nil {
            structMap["band_24_usage"] = a.Band24Usage.Value()
        } else {
            structMap["band_24_usage"] = nil
        }
    }
    if a.Band5 != nil {
        structMap["band_5"] = a.Band5.toMap()
    }
    if a.Band6 != nil {
        structMap["band_6"] = a.Band6.toMap()
    }
    if a.ScanningEnabled != nil {
        structMap["scanning_enabled"] = a.ScanningEnabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsRadioConfig.
// It customizes the JSON unmarshaling process for ApStatsRadioConfig objects.
func (a *ApStatsRadioConfig) UnmarshalJSON(input []byte) error {
    var temp tempApStatsRadioConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band_24", "band_24_usage", "band_5", "band_6", "scanning_enabled")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Band24 = temp.Band24
    a.Band24Usage = temp.Band24Usage
    a.Band5 = temp.Band5
    a.Band6 = temp.Band6
    a.ScanningEnabled = temp.ScanningEnabled
    return nil
}

// tempApStatsRadioConfig is a temporary struct used for validating the fields of ApStatsRadioConfig.
type tempApStatsRadioConfig  struct {
    Band24          *ApStatsRadioConfigBand `json:"band_24,omitempty"`
    Band24Usage     Optional[string]        `json:"band_24_usage"`
    Band5           *ApStatsRadioConfigBand `json:"band_5,omitempty"`
    Band6           *ApStatsRadioConfigBand `json:"band_6,omitempty"`
    ScanningEnabled *bool                   `json:"scanning_enabled,omitempty"`
}
