package models

import (
    "encoding/json"
)

// RrmBandMetricInterference represents a RrmBandMetricInterference struct.
type RrmBandMetricInterference struct {
    Radar                *float64       `json:"radar,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmBandMetricInterference.
// It customizes the JSON marshaling process for RrmBandMetricInterference objects.
func (r RrmBandMetricInterference) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RrmBandMetricInterference object to a map representation for JSON marshaling.
func (r RrmBandMetricInterference) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Radar != nil {
        structMap["radar"] = r.Radar
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmBandMetricInterference.
// It customizes the JSON unmarshaling process for RrmBandMetricInterference objects.
func (r *RrmBandMetricInterference) UnmarshalJSON(input []byte) error {
    var temp tempRrmBandMetricInterference
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "radar")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Radar = temp.Radar
    return nil
}

// tempRrmBandMetricInterference is a temporary struct used for validating the fields of RrmBandMetricInterference.
type tempRrmBandMetricInterference  struct {
    Radar *float64 `json:"radar,omitempty"`
}
