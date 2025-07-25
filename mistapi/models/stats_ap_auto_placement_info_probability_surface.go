// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsApAutoPlacementInfoProbabilitySurface represents a StatsApAutoPlacementInfoProbabilitySurface struct.
// Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result
type StatsApAutoPlacementInfoProbabilitySurface struct {
    // The radius representing placement uncertainty, measured in pixels
    Radius               *float64               `json:"radius,omitempty"`
    // The radius representing placement uncertainty, measured in meters
    RadiusM              *float64               `json:"radius_m,omitempty"`
    // Y-coordinate of the potential placement’s center, measured in pixels
    X                    *float64               `json:"x,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApAutoPlacementInfoProbabilitySurface,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApAutoPlacementInfoProbabilitySurface) String() string {
    return fmt.Sprintf(
    	"StatsApAutoPlacementInfoProbabilitySurface[Radius=%v, RadiusM=%v, X=%v, AdditionalProperties=%v]",
    	s.Radius, s.RadiusM, s.X, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoPlacementInfoProbabilitySurface.
// It customizes the JSON marshaling process for StatsApAutoPlacementInfoProbabilitySurface objects.
func (s StatsApAutoPlacementInfoProbabilitySurface) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "radius", "radius_m", "x"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoPlacementInfoProbabilitySurface object to a map representation for JSON marshaling.
func (s StatsApAutoPlacementInfoProbabilitySurface) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Radius != nil {
        structMap["radius"] = s.Radius
    }
    if s.RadiusM != nil {
        structMap["radius_m"] = s.RadiusM
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApAutoPlacementInfoProbabilitySurface.
// It customizes the JSON unmarshaling process for StatsApAutoPlacementInfoProbabilitySurface objects.
func (s *StatsApAutoPlacementInfoProbabilitySurface) UnmarshalJSON(input []byte) error {
    var temp tempStatsApAutoPlacementInfoProbabilitySurface
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "radius", "radius_m", "x")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Radius = temp.Radius
    s.RadiusM = temp.RadiusM
    s.X = temp.X
    return nil
}

// tempStatsApAutoPlacementInfoProbabilitySurface is a temporary struct used for validating the fields of StatsApAutoPlacementInfoProbabilitySurface.
type tempStatsApAutoPlacementInfoProbabilitySurface  struct {
    Radius  *float64 `json:"radius,omitempty"`
    RadiusM *float64 `json:"radius_m,omitempty"`
    X       *float64 `json:"x,omitempty"`
}
