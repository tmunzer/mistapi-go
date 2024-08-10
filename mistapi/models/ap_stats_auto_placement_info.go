package models

import (
    "encoding/json"
)

// ApStatsAutoPlacementInfo represents a ApStatsAutoPlacementInfo struct.
// Additional information about auto placements AP data
type ApStatsAutoPlacementInfo struct {
    // All APs sharing a given cluster number can be placed relative to each other
    ClusterNumber        *int                                        `json:"cluster_number,omitempty"`
    // The orientation of an AP
    OrientationStats     *int                                        `json:"orientation_stats,omitempty"`
    // Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result
    ProbabilitySurface   *ApStatsAutoPlacementInfoProbabilitySurface `json:"probability_surface,omitempty"`
    AdditionalProperties map[string]any                              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsAutoPlacementInfo.
// It customizes the JSON marshaling process for ApStatsAutoPlacementInfo objects.
func (a ApStatsAutoPlacementInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsAutoPlacementInfo object to a map representation for JSON marshaling.
func (a ApStatsAutoPlacementInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ClusterNumber != nil {
        structMap["cluster_number"] = a.ClusterNumber
    }
    if a.OrientationStats != nil {
        structMap["orientation_stats"] = a.OrientationStats
    }
    if a.ProbabilitySurface != nil {
        structMap["probability_surface"] = a.ProbabilitySurface.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsAutoPlacementInfo.
// It customizes the JSON unmarshaling process for ApStatsAutoPlacementInfo objects.
func (a *ApStatsAutoPlacementInfo) UnmarshalJSON(input []byte) error {
    var temp tempApStatsAutoPlacementInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cluster_number", "orientation_stats", "probability_surface")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ClusterNumber = temp.ClusterNumber
    a.OrientationStats = temp.OrientationStats
    a.ProbabilitySurface = temp.ProbabilitySurface
    return nil
}

// tempApStatsAutoPlacementInfo is a temporary struct used for validating the fields of ApStatsAutoPlacementInfo.
type tempApStatsAutoPlacementInfo  struct {
    ClusterNumber      *int                                        `json:"cluster_number,omitempty"`
    OrientationStats   *int                                        `json:"orientation_stats,omitempty"`
    ProbabilitySurface *ApStatsAutoPlacementInfoProbabilitySurface `json:"probability_surface,omitempty"`
}
