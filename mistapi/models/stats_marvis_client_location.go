// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// StatsMarvisClientLocation represents a StatsMarvisClientLocation struct.
// Last known location fix for a Marvis Client device
type StatsMarvisClientLocation struct {
	// UUID of the floor-plan map
	MapId *uuid.UUID `json:"map_id,omitempty"`
	// UUID of the site the device was located in
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Timestamp of the location fix, in epoch seconds
	Timestamp *int `json:"timestamp,omitempty"`
	// X coordinate on the floor-plan map, in pixels
	X *float64 `json:"x,omitempty"`
	// Y coordinate on the floor-plan map, in pixels
	Y                    *float64               `json:"y,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMarvisClientLocation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMarvisClientLocation) String() string {
	return fmt.Sprintf(
		"StatsMarvisClientLocation[MapId=%v, SiteId=%v, Timestamp=%v, X=%v, Y=%v, AdditionalProperties=%v]",
		s.MapId, s.SiteId, s.Timestamp, s.X, s.Y, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMarvisClientLocation.
// It customizes the JSON marshaling process for StatsMarvisClientLocation objects.
func (s StatsMarvisClientLocation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"map_id", "site_id", "timestamp", "x", "y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMarvisClientLocation object to a map representation for JSON marshaling.
func (s StatsMarvisClientLocation) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.MapId != nil {
		structMap["map_id"] = s.MapId
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.Timestamp != nil {
		structMap["timestamp"] = s.Timestamp
	}
	if s.X != nil {
		structMap["x"] = s.X
	}
	if s.Y != nil {
		structMap["y"] = s.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMarvisClientLocation.
// It customizes the JSON unmarshaling process for StatsMarvisClientLocation objects.
func (s *StatsMarvisClientLocation) UnmarshalJSON(input []byte) error {
	var temp tempStatsMarvisClientLocation
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "map_id", "site_id", "timestamp", "x", "y")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.MapId = temp.MapId
	s.SiteId = temp.SiteId
	s.Timestamp = temp.Timestamp
	s.X = temp.X
	s.Y = temp.Y
	return nil
}

// tempStatsMarvisClientLocation is a temporary struct used for validating the fields of StatsMarvisClientLocation.
type tempStatsMarvisClientLocation struct {
	MapId     *uuid.UUID `json:"map_id,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Timestamp *int       `json:"timestamp,omitempty"`
	X         *float64   `json:"x,omitempty"`
	Y         *float64   `json:"y,omitempty"`
}
