package models

import (
    "encoding/json"
)

// ApStatsAutoPlacement represents a ApStatsAutoPlacement struct.
type ApStatsAutoPlacement struct {
    // Additional information about auto placements AP data
    Info                 *ApStatsAutoPlacementInfo `json:"info,omitempty"`
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

// MarshalJSON implements the json.Marshaler interface for ApStatsAutoPlacement.
// It customizes the JSON marshaling process for ApStatsAutoPlacement objects.
func (a ApStatsAutoPlacement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsAutoPlacement object to a map representation for JSON marshaling.
func (a ApStatsAutoPlacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Info != nil {
        structMap["info"] = a.Info.toMap()
    }
    if a.RecommendedAnchor != nil {
        structMap["recommended_anchor"] = a.RecommendedAnchor
    }
    if a.Status != nil {
        structMap["status"] = a.Status
    }
    if a.StatusDetail != nil {
        structMap["status_detail"] = a.StatusDetail
    }
    if a.UseAutoPlacement != nil {
        structMap["use_auto_placement"] = a.UseAutoPlacement
    }
    if a.X != nil {
        structMap["x"] = a.X
    }
    if a.XM != nil {
        structMap["x_m"] = a.XM
    }
    if a.Y != nil {
        structMap["y"] = a.Y
    }
    if a.YM != nil {
        structMap["y_m"] = a.YM
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsAutoPlacement.
// It customizes the JSON unmarshaling process for ApStatsAutoPlacement objects.
func (a *ApStatsAutoPlacement) UnmarshalJSON(input []byte) error {
    var temp tempApStatsAutoPlacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "info", "recommended_anchor", "status", "status_detail", "use_auto_placement", "x", "x_m", "y", "y_m")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Info = temp.Info
    a.RecommendedAnchor = temp.RecommendedAnchor
    a.Status = temp.Status
    a.StatusDetail = temp.StatusDetail
    a.UseAutoPlacement = temp.UseAutoPlacement
    a.X = temp.X
    a.XM = temp.XM
    a.Y = temp.Y
    a.YM = temp.YM
    return nil
}

// tempApStatsAutoPlacement is a temporary struct used for validating the fields of ApStatsAutoPlacement.
type tempApStatsAutoPlacement  struct {
    Info              *ApStatsAutoPlacementInfo `json:"info,omitempty"`
    RecommendedAnchor *bool                     `json:"recommended_anchor,omitempty"`
    Status            *string                   `json:"status,omitempty"`
    StatusDetail      *string                   `json:"status_detail,omitempty"`
    UseAutoPlacement  *bool                     `json:"use_auto_placement,omitempty"`
    X                 *float64                  `json:"x,omitempty"`
    XM                *float64                  `json:"x_m,omitempty"`
    Y                 *float64                  `json:"y,omitempty"`
    YM                *float64                  `json:"y_m,omitempty"`
}
