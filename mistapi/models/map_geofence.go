// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MapGeofence represents a MapGeofence struct.
type MapGeofence struct {
	// Name of the geofence
	Name *string `json:"name,omitempty"`
	// List of vertices defining the geofence
	Vertices             []MapGeofenceVertice   `json:"vertices,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapGeofence,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapGeofence) String() string {
	return fmt.Sprintf(
		"MapGeofence[Name=%v, Vertices=%v, AdditionalProperties=%v]",
		m.Name, m.Vertices, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapGeofence.
// It customizes the JSON marshaling process for MapGeofence objects.
func (m MapGeofence) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"name", "vertices"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapGeofence object to a map representation for JSON marshaling.
func (m MapGeofence) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.Vertices != nil {
		structMap["vertices"] = m.Vertices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapGeofence.
// It customizes the JSON unmarshaling process for MapGeofence objects.
func (m *MapGeofence) UnmarshalJSON(input []byte) error {
	var temp tempMapGeofence
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "vertices")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Name = temp.Name
	m.Vertices = temp.Vertices
	return nil
}

// tempMapGeofence is a temporary struct used for validating the fields of MapGeofence.
type tempMapGeofence struct {
	Name     *string              `json:"name,omitempty"`
	Vertices []MapGeofenceVertice `json:"vertices,omitempty"`
}
