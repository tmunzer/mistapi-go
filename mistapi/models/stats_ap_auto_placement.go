// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApAutoPlacement represents a StatsApAutoPlacement struct.
type StatsApAutoPlacement struct {
	// Additional information about auto placements AP data
	Info *StatsApAutoPlacementInfo `json:"info,omitempty"`
	// Flag to represent if AP is recommended as an anchor by auto placement service
	RecommendedAnchor *bool `json:"recommended_anchor,omitempty"`
	// Basic Placement Status
	Status *string `json:"status,omitempty"`
	// Additional info about placement status
	StatusDetail *string `json:"status_detail,omitempty"`
	// X Autoplaced Position in pixels
	X *float64 `json:"x,omitempty"`
	// X Autoplaced Position in meters
	XM *float64 `json:"x_m,omitempty"`
	// Y Autoplaced Position in pixels
	Y *float64 `json:"y,omitempty"`
	// X Autoplaced Position in meters
	YM                   *float64               `json:"y_m,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApAutoPlacement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApAutoPlacement) String() string {
	return fmt.Sprintf(
		"StatsApAutoPlacement[Info=%v, RecommendedAnchor=%v, Status=%v, StatusDetail=%v, X=%v, XM=%v, Y=%v, YM=%v, AdditionalProperties=%v]",
		s.Info, s.RecommendedAnchor, s.Status, s.StatusDetail, s.X, s.XM, s.Y, s.YM, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoPlacement.
// It customizes the JSON marshaling process for StatsApAutoPlacement objects.
func (s StatsApAutoPlacement) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"info", "recommended_anchor", "status", "status_detail", "x", "x_m", "y", "y_m"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoPlacement object to a map representation for JSON marshaling.
func (s StatsApAutoPlacement) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "info", "recommended_anchor", "status", "status_detail", "x", "x_m", "y", "y_m")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Info = temp.Info
	s.RecommendedAnchor = temp.RecommendedAnchor
	s.Status = temp.Status
	s.StatusDetail = temp.StatusDetail
	s.X = temp.X
	s.XM = temp.XM
	s.Y = temp.Y
	s.YM = temp.YM
	return nil
}

// tempStatsApAutoPlacement is a temporary struct used for validating the fields of StatsApAutoPlacement.
type tempStatsApAutoPlacement struct {
	Info              *StatsApAutoPlacementInfo `json:"info,omitempty"`
	RecommendedAnchor *bool                     `json:"recommended_anchor,omitempty"`
	Status            *string                   `json:"status,omitempty"`
	StatusDetail      *string                   `json:"status_detail,omitempty"`
	X                 *float64                  `json:"x,omitempty"`
	XM                *float64                  `json:"x_m,omitempty"`
	Y                 *float64                  `json:"y,omitempty"`
	YM                *float64                  `json:"y_m,omitempty"`
}
