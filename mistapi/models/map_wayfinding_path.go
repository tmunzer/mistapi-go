package models

import (
    "encoding/json"
)

// MapWayfindingPath represents a MapWayfindingPath struct.
// a JSON blob for wayfinding (using Dijkstra’s algorithm)
type MapWayfindingPath struct {
    Coordinate           *string        `json:"coordinate,omitempty"`
    Nodes                []MapNode      `json:"nodes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapWayfindingPath.
// It customizes the JSON marshaling process for MapWayfindingPath objects.
func (m MapWayfindingPath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapWayfindingPath object to a map representation for JSON marshaling.
func (m MapWayfindingPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "coordinate", "nodes")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Coordinate = temp.Coordinate
    m.Nodes = temp.Nodes
    return nil
}

// tempMapWayfindingPath is a temporary struct used for validating the fields of MapWayfindingPath.
type tempMapWayfindingPath  struct {
    Coordinate *string   `json:"coordinate,omitempty"`
    Nodes      []MapNode `json:"nodes,omitempty"`
}
