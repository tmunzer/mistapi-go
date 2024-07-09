package models

import (
    "encoding/json"
)

// MapWallPath represents a MapWallPath struct.
// a JSON blob for wall definition (same format as wayfinding_path)
type MapWallPath struct {
    Coordinate           *string        `json:"coordinate,omitempty"`
    Nodes                []MapNode      `json:"nodes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapWallPath.
// It customizes the JSON marshaling process for MapWallPath objects.
func (m MapWallPath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapWallPath object to a map representation for JSON marshaling.
func (m MapWallPath) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for MapWallPath.
// It customizes the JSON unmarshaling process for MapWallPath objects.
func (m *MapWallPath) UnmarshalJSON(input []byte) error {
    var temp mapWallPath
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

// mapWallPath is a temporary struct used for validating the fields of MapWallPath.
type mapWallPath  struct {
    Coordinate *string   `json:"coordinate,omitempty"`
    Nodes      []MapNode `json:"nodes,omitempty"`
}
