package models

import (
    "encoding/json"
)

// StatsApAutoPlacement represents a StatsApAutoPlacement struct.
type StatsApAutoPlacement struct {
    // Additional information about auto placements AP data
    Info                 *StatsApAutoPlacementInfo `json:"info,omitempty"`
    // Flag to represent if AP is recommended as an anchor by auto placement service
    RecommendedAnchor    *bool                     `json:"recommended_anchor,omitempty"`
    // Basic Placement Status
    Status               *string                   `json:"status,omitempty"`
    // Additional info about placement status
    StatusDetail         *string                   `json:"status_detail,omitempty"`
    // Flag to represent if auto_placement values are currently utilized
    UseAutoPlacement     *bool                     `json:"use_auto_placement,omitempty"`
    // X Autoplaced Position in pixels
    X                    *float64                  `json:"x,omitempty"`
    // X Autoplaced Position in meters
    XM                   *float64                  `json:"x_m,omitempty"`
    // Y Autoplaced Position in pixels
    Y                    *float64                  `json:"y,omitempty"`
    // X Autoplaced Position in meters
    YM                   *float64                  `json:"y_m,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoPlacement.
// It customizes the JSON marshaling process for StatsApAutoPlacement objects.
func (s StatsApAutoPlacement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoPlacement object to a map representation for JSON marshaling.
func (s StatsApAutoPlacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Info != nil {
        structMap["info"] = s.Info.toMap()
    }
    if s.RecommendedAnchor != nil {
        structMap["recommended_anchor"] = s.RecommendedAnchor
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.StatusDetail != nil {
        structMap["status_detail"] = s.StatusDetail
    }
    if s.UseAutoPlacement != nil {
        structMap["use_auto_placement"] = s.UseAutoPlacement
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    if s.XM != nil {
        structMap["x_m"] = s.XM
    }
    if s.Y != nil {
        structMap["y"] = s.Y
    }
    if s.YM != nil {
        structMap["y_m"] = s.YM
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApAutoPlacement.
// It customizes the JSON unmarshaling process for StatsApAutoPlacement objects.
func (s *StatsApAutoPlacement) UnmarshalJSON(input []byte) error {
    var temp tempStatsApAutoPlacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "info", "recommended_anchor", "status", "status_detail", "use_auto_placement", "x", "x_m", "y", "y_m")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Info = temp.Info
    s.RecommendedAnchor = temp.RecommendedAnchor
    s.Status = temp.Status
    s.StatusDetail = temp.StatusDetail
    s.UseAutoPlacement = temp.UseAutoPlacement
    s.X = temp.X
    s.XM = temp.XM
    s.Y = temp.Y
    s.YM = temp.YM
    return nil
}

// tempStatsApAutoPlacement is a temporary struct used for validating the fields of StatsApAutoPlacement.
type tempStatsApAutoPlacement  struct {
    Info              *StatsApAutoPlacementInfo `json:"info,omitempty"`
    RecommendedAnchor *bool                     `json:"recommended_anchor,omitempty"`
    Status            *string                   `json:"status,omitempty"`
    StatusDetail      *string                   `json:"status_detail,omitempty"`
    UseAutoPlacement  *bool                     `json:"use_auto_placement,omitempty"`
    X                 *float64                  `json:"x,omitempty"`
    XM                *float64                  `json:"x_m,omitempty"`
    Y                 *float64                  `json:"y,omitempty"`
    YM                *float64                  `json:"y_m,omitempty"`
}
