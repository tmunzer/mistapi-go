// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MapSitesurveyPathItems represents a MapSitesurveyPathItems struct.
type MapSitesurveyPathItems struct {
	Coordinate *string `json:"coordinate,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id                   *uuid.UUID             `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Nodes                []MapNode              `json:"nodes,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapSitesurveyPathItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapSitesurveyPathItems) String() string {
	return fmt.Sprintf(
		"MapSitesurveyPathItems[Coordinate=%v, Id=%v, Name=%v, Nodes=%v, AdditionalProperties=%v]",
		m.Coordinate, m.Id, m.Name, m.Nodes, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapSitesurveyPathItems.
// It customizes the JSON marshaling process for MapSitesurveyPathItems objects.
func (m MapSitesurveyPathItems) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"coordinate", "id", "name", "nodes"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapSitesurveyPathItems object to a map representation for JSON marshaling.
func (m MapSitesurveyPathItems) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Coordinate != nil {
		structMap["coordinate"] = m.Coordinate
	}
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.Nodes != nil {
		structMap["nodes"] = m.Nodes
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapSitesurveyPathItems.
// It customizes the JSON unmarshaling process for MapSitesurveyPathItems objects.
func (m *MapSitesurveyPathItems) UnmarshalJSON(input []byte) error {
	var temp tempMapSitesurveyPathItems
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coordinate", "id", "name", "nodes")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Coordinate = temp.Coordinate
	m.Id = temp.Id
	m.Name = temp.Name
	m.Nodes = temp.Nodes
	return nil
}

// tempMapSitesurveyPathItems is a temporary struct used for validating the fields of MapSitesurveyPathItems.
type tempMapSitesurveyPathItems struct {
	Coordinate *string    `json:"coordinate,omitempty"`
	Id         *uuid.UUID `json:"id,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Nodes      []MapNode  `json:"nodes,omitempty"`
}
