package models

import (
    "encoding/json"
)

// StatsApAutoPlacementInfo represents a StatsApAutoPlacementInfo struct.
// Additional information about auto placements AP data
type StatsApAutoPlacementInfo struct {
    // All APs sharing a given cluster number can be placed relative to each other
    ClusterNumber        *int                                        `json:"cluster_number,omitempty"`
    // The orientation of an AP
    OrientationStats     *int                                        `json:"orientation_stats,omitempty"`
    // Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result
    ProbabilitySurface   *StatsApAutoPlacementInfoProbabilitySurface `json:"probability_surface,omitempty"`
    AdditionalProperties map[string]any                              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoPlacementInfo.
// It customizes the JSON marshaling process for StatsApAutoPlacementInfo objects.
func (s StatsApAutoPlacementInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoPlacementInfo object to a map representation for JSON marshaling.
func (s StatsApAutoPlacementInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ClusterNumber != nil {
        structMap["cluster_number"] = s.ClusterNumber
    }
    if s.OrientationStats != nil {
        structMap["orientation_stats"] = s.OrientationStats
    }
    if s.ProbabilitySurface != nil {
        structMap["probability_surface"] = s.ProbabilitySurface.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApAutoPlacementInfo.
// It customizes the JSON unmarshaling process for StatsApAutoPlacementInfo objects.
func (s *StatsApAutoPlacementInfo) UnmarshalJSON(input []byte) error {
    var temp tempStatsApAutoPlacementInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cluster_number", "orientation_stats", "probability_surface")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ClusterNumber = temp.ClusterNumber
    s.OrientationStats = temp.OrientationStats
    s.ProbabilitySurface = temp.ProbabilitySurface
    return nil
}

// tempStatsApAutoPlacementInfo is a temporary struct used for validating the fields of StatsApAutoPlacementInfo.
type tempStatsApAutoPlacementInfo  struct {
    ClusterNumber      *int                                        `json:"cluster_number,omitempty"`
    OrientationStats   *int                                        `json:"orientation_stats,omitempty"`
    ProbabilitySurface *StatsApAutoPlacementInfoProbabilitySurface `json:"probability_surface,omitempty"`
}