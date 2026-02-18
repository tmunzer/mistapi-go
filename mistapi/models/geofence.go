// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// Geofence represents a Geofence struct.
type Geofence struct {
	// Name of the geofence
	Name *string `json:"name,omitempty"`
	// List of vertices defining the geofence
	Vertices             []Vertex               `json:"vertices,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Geofence,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g Geofence) String() string {
	return fmt.Sprintf(
		"Geofence[Name=%v, Vertices=%v, AdditionalProperties=%v]",
		g.Name, g.Vertices, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Geofence.
// It customizes the JSON marshaling process for Geofence objects.
func (g Geofence) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"name", "vertices"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the Geofence object to a map representation for JSON marshaling.
func (g Geofence) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Name != nil {
		structMap["name"] = g.Name
	}
	if g.Vertices != nil {
		structMap["vertices"] = g.Vertices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Geofence.
// It customizes the JSON unmarshaling process for Geofence objects.
func (g *Geofence) UnmarshalJSON(input []byte) error {
	var temp tempGeofence
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "vertices")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Name = temp.Name
	g.Vertices = temp.Vertices
	return nil
}

// tempGeofence is a temporary struct used for validating the fields of Geofence.
type tempGeofence struct {
	Name     *string  `json:"name,omitempty"`
	Vertices []Vertex `json:"vertices,omitempty"`
}
