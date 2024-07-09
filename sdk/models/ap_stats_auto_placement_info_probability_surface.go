package models

import (
    "encoding/json"
)

// ApStatsAutoPlacementInfoProbabilitySurface represents a ApStatsAutoPlacementInfoProbabilitySurface struct.
// Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result
type ApStatsAutoPlacementInfoProbabilitySurface struct {
    // The radius representing placement uncertainty, measured in pixels
    Radius               *float64       `json:"radius,omitempty"`
    // The radius representing placement uncertainty, measured in meters
    RadiusM              *float64       `json:"radius_m,omitempty"`
    // Y-coordinate of the potential placement’s center, measured in pixels
    X                    *float64       `json:"x,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsAutoPlacementInfoProbabilitySurface.
// It customizes the JSON marshaling process for ApStatsAutoPlacementInfoProbabilitySurface objects.
func (a ApStatsAutoPlacementInfoProbabilitySurface) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsAutoPlacementInfoProbabilitySurface object to a map representation for JSON marshaling.
func (a ApStatsAutoPlacementInfoProbabilitySurface) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Radius != nil {
        structMap["radius"] = a.Radius
    }
    if a.RadiusM != nil {
        structMap["radius_m"] = a.RadiusM
    }
    if a.X != nil {
        structMap["x"] = a.X
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsAutoPlacementInfoProbabilitySurface.
// It customizes the JSON unmarshaling process for ApStatsAutoPlacementInfoProbabilitySurface objects.
func (a *ApStatsAutoPlacementInfoProbabilitySurface) UnmarshalJSON(input []byte) error {
    var temp apStatsAutoPlacementInfoProbabilitySurface
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "radius", "radius_m", "x")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Radius = temp.Radius
    a.RadiusM = temp.RadiusM
    a.X = temp.X
    return nil
}

// apStatsAutoPlacementInfoProbabilitySurface is a temporary struct used for validating the fields of ApStatsAutoPlacementInfoProbabilitySurface.
type apStatsAutoPlacementInfoProbabilitySurface  struct {
    Radius  *float64 `json:"radius,omitempty"`
    RadiusM *float64 `json:"radius_m,omitempty"`
    X       *float64 `json:"x,omitempty"`
}
