// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MapWayfindingPath represents a MapWayfindingPath struct.
// JSON blob for wayfinding (using Dijkstra’s algorithm)
type MapWayfindingPath struct {
	Coordinate           *string                `json:"coordinate,omitempty"`
	Nodes                []MapNode              `json:"nodes,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapWayfindingPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapWayfindingPath) String() string {
	return fmt.Sprintf(
		"MapWayfindingPath[Coordinate=%v, Nodes=%v, AdditionalProperties=%v]",
		m.Coordinate, m.Nodes, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapWayfindingPath.
// It customizes the JSON marshaling process for MapWayfindingPath objects.
func (m MapWayfindingPath) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"coordinate", "nodes"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapWayfindingPath object to a map representation for JSON marshaling.
func (m MapWayfindingPath) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Coordinate != nil {
		structMap["coordinate"] = m.Coordinate
	}
	if m.Nodes != nil {
		structMap["nodes"] = m.Nodes
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapWayfindingPath.
// It customizes the JSON unmarshaling process for MapWayfindingPath objects.
func (m *MapWayfindingPath) UnmarshalJSON(input []byte) error {
	var temp tempMapWayfindingPath
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coordinate", "nodes")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Coordinate = temp.Coordinate
	m.Nodes = temp.Nodes
	return nil
}

// tempMapWayfindingPath is a temporary struct used for validating the fields of MapWayfindingPath.
type tempMapWayfindingPath struct {
	Coordinate *string   `json:"coordinate,omitempty"`
	Nodes      []MapNode `json:"nodes,omitempty"`
}
