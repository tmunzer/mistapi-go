package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// MapSitesurveyPathItems represents a MapSitesurveyPathItems struct.
type MapSitesurveyPathItems struct {
    Coordinate           *string        `json:"coordinate,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Nodes                []MapNode      `json:"nodes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapSitesurveyPathItems.
// It customizes the JSON marshaling process for MapSitesurveyPathItems objects.
func (m MapSitesurveyPathItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapSitesurveyPathItems object to a map representation for JSON marshaling.
func (m MapSitesurveyPathItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp mapSitesurveyPathItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "coordinate", "id", "name", "nodes")
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

// mapSitesurveyPathItems is a temporary struct used for validating the fields of MapSitesurveyPathItems.
type mapSitesurveyPathItems  struct {
    Coordinate *string    `json:"coordinate,omitempty"`
    Id         *uuid.UUID `json:"id,omitempty"`
    Name       *string    `json:"name,omitempty"`
    Nodes      []MapNode  `json:"nodes,omitempty"`
}
