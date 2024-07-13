package models

import (
    "encoding/json"
)

// ApStatsRadioStat represents a ApStatsRadioStat struct.
type ApStatsRadioStat struct {
    // radio stat
    Band24               *ApRadioStats  `json:"band_24,omitempty"`
    // radio stat
    Band5                *ApRadioStats  `json:"band_5,omitempty"`
    // radio stat
    Band6                *ApRadioStats  `json:"band_6,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsRadioStat.
// It customizes the JSON marshaling process for ApStatsRadioStat objects.
func (a ApStatsRadioStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsRadioStat object to a map representation for JSON marshaling.
func (a ApStatsRadioStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Band24 != nil {
        structMap["band_24"] = a.Band24.toMap()
    }
    if a.Band5 != nil {
        structMap["band_5"] = a.Band5.toMap()
    }
    if a.Band6 != nil {
        structMap["band_6"] = a.Band6.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsRadioStat.
// It customizes the JSON unmarshaling process for ApStatsRadioStat objects.
func (a *ApStatsRadioStat) UnmarshalJSON(input []byte) error {
    var temp apStatsRadioStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band_24", "band_5", "band_6")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Band24 = temp.Band24
    a.Band5 = temp.Band5
    a.Band6 = temp.Band6
    return nil
}

// apStatsRadioStat is a temporary struct used for validating the fields of ApStatsRadioStat.
type apStatsRadioStat  struct {
    Band24 *ApRadioStats `json:"band_24,omitempty"`
    Band5  *ApRadioStats `json:"band_5,omitempty"`
    Band6  *ApRadioStats `json:"band_6,omitempty"`
}
