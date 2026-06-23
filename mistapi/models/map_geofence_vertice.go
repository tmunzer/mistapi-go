// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MapGeofenceVertice represents a MapGeofenceVertice struct.
// Vertex coordinate for a map geofence polygon
type MapGeofenceVertice struct {
	// Geofence vertex X coordinate in map units
	X *float64 `json:"X,omitempty"`
	// Geofence vertex Y coordinate in map units
	Y                    *float64               `json:"Y,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapGeofenceVertice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapGeofenceVertice) String() string {
	return fmt.Sprintf(
		"MapGeofenceVertice[X=%v, Y=%v, AdditionalProperties=%v]",
		m.X, m.Y, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapGeofenceVertice.
// It customizes the JSON marshaling process for MapGeofenceVertice objects.
func (m MapGeofenceVertice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"X", "Y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapGeofenceVertice object to a map representation for JSON marshaling.
func (m MapGeofenceVertice) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.X != nil {
		structMap["X"] = m.X
	}
	if m.Y != nil {
		structMap["Y"] = m.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapGeofenceVertice.
// It customizes the JSON unmarshaling process for MapGeofenceVertice objects.
func (m *MapGeofenceVertice) UnmarshalJSON(input []byte) error {
	var temp tempMapGeofenceVertice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "X", "Y")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.X = temp.X
	m.Y = temp.Y
	return nil
}

// tempMapGeofenceVertice is a temporary struct used for validating the fields of MapGeofenceVertice.
type tempMapGeofenceVertice struct {
	X *float64 `json:"X,omitempty"`
	Y *float64 `json:"Y,omitempty"`
}
